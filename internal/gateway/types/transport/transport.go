package transport

type Tender struct {
	Id          string
	Name        string
	Description string
	ServiceType string
	Status      string
	Version     int
	Responsible string
}

type Bids struct {
	Id          string
	Name        string
	Description string
	Status      string
	TenderId    string
	Version     int
	Responsible string
}

type BidFeedback struct {
	Id          string
	BidID       string
	Description string
	Responsible string
	CreatedAt   string
}
