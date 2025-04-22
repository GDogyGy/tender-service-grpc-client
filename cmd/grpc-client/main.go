package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"GrpcClientForTenderService/internal/config"
	bidsGatewayFetch "GrpcClientForTenderService/internal/gateway/bids/fetch"
	tenderGatewayFetch "GrpcClientForTenderService/internal/gateway/tender/fetch"
	bidsFetch "GrpcClientForTenderService/internal/handlers/bids/fetch"
	tenderFetch "GrpcClientForTenderService/internal/handlers/tender/fetch"
	"GrpcClientForTenderService/internal/kafka"
	bidsProto "GrpcClientForTenderService/internal/protos/gen/bids/fetch"
	tenderProto "GrpcClientForTenderService/internal/protos/gen/tender/fetch"
	bidsUseCaseFetch "GrpcClientForTenderService/internal/usecases/bids/fetch"
	tenderUseCaseFetch "GrpcClientForTenderService/internal/usecases/tender/fetch"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	cfg := config.MustLoad()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	log, err := setupLogger(cfg.DebugLevel)
	if err != nil {
		log.Error("Failed to init logger", slog.Attr{Value: slog.StringValue(err.Error())})
		os.Exit(1) // nolint:gocritic
	}

	log.Info("Starting tender api server", slog.String("Env", cfg.Env))

	conn, err := grpc.NewClient(
		cfg.GrpcRemote,
		grpc.WithTransportCredentials(insecure.NewCredentials()), // Для теста без TLS
	)
	if err != nil {
		log.Error("Failed to init Grpc client", slog.Attr{Value: slog.StringValue(err.Error())})
		os.Exit(1) // nolint:gocritic
	}
	defer func() { _ = conn.Close() }()

	// <! ProtoService Bids
	bidsFetchService := bidsProto.NewBidsServiceFetchClient(conn)
	// ProtoService Bids !>

	// <! Gateway bids
	bidsFetchGateway := bidsGatewayFetch.NewGateway(log, bidsFetchService)
	// Gateway bids  !>

	// <! UseCase bids
	bidsFetchUseCase := bidsUseCaseFetch.NewService(bidsFetchGateway)
	// Gateway bids  !>

	// <! Handler Bids
	handlerBidsFetch := bidsFetch.NewHandler(log, bidsFetchUseCase)
	// Handler Bids !>

	// <! ProtoService Tender
	tenderFetchService := tenderProto.NewTenderServiceFetchClient(conn)
	// ProtoService Tender !>

	// <! Gateway tender
	tenderFetchGateway := tenderGatewayFetch.NewGateway(log, tenderFetchService)
	// Gateway tender  !>

	// <! UseCase tender
	tenderFetchUseCase := tenderUseCaseFetch.NewService(tenderFetchGateway)
	// Gateway tender  !>

	// <! Handler Tender
	handlerTenderFetch := tenderFetch.NewHandler(log, tenderFetchUseCase)
	// Handler Tender !>

	router := http.NewServeMux()

	handlerBidsFetch.Register(router)
	handlerTenderFetch.Register(router)

	// <! Kafka consumer
	messageChan := make(chan string, 256)
	var wg sync.WaitGroup

	consumer, err := kafka.NewConsumer(
		[]string{"localhost:29092"},
		"model-events",
		messageChan,
	)
	if err != nil {
		log.Error("Failed to init Kafka client Kafka", slog.Attr{Value: slog.StringValue(err.Error())})
		os.Exit(1) // nolint:gocritic
	}

	wg.Add(3)

	go func() {
		defer wg.Done()
		for {
			select {
			case msg := <-messageChan:
				log.Info("Processing message", slog.String("message", msg))
			case <-ctx.Done():
				log.Info("Stopping message processor")
				return
			}
		}
	}()

	go func() {
		consumer.Start(ctx, []string{"model-events"})
		wg.Done()
		log.Info("Kafka consumer stopped")
	}()
	// kafka consumer!>

	go func() {
		StartServerHttp(ctx, cfg, log, router)
		wg.Done()
	}()

	<-ctx.Done()

	wg.Wait()
	close(messageChan)
	log.Info("All service stopped")
}

func StartServerHttp(ctx context.Context, cfg *config.Config, log *slog.Logger, router http.Handler) {
	log.Info("http server starting", slog.String("address", cfg.Address))

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("http listen and serve returned err:", slog.Attr{Value: slog.StringValue(err.Error())})
		}
	}()

	<-ctx.Done()

	if err := srv.Shutdown(ctx); err != nil {
		log.Info("server shutdown returned an err: %v\n", slog.Attr{Value: slog.StringValue(err.Error())})
	}
	log.Info("http server stopping")
}

func setupLogger(lvl string) (*slog.Logger, error) {
	const op = "main.setupLogger"
	var sl slog.Level
	err := sl.UnmarshalText([]byte(lvl))
	if err != nil {
		return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})), fmt.Errorf("%s:%v", op, err)
	}
	return slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: sl}),
	), nil
}
