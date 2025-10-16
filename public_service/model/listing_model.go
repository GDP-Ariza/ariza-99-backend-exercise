package model

type ListingModel struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	ListingType string `json:"listing_type"`
	Price       int64  `json:"price"`
	CreatedAt   int64  `json:"created_at"` // in microseconds
	UpdatedAt   int64  `json:"updated_at"` // in microseconds
}

type Listing struct {
	ID          int    `json:"id"`
	ListingType string `json:"listing_type"`
	Price       int64  `json:"price"`
	CreatedAt   int64  `json:"created_at"` // in microseconds
	UpdatedAt   int64  `json:"updated_at"` // in microseconds
	User        User   `json:"user"`
}

type ListingRequest struct {
	UserID      int    `json:"user_id"`
	ListingType string `json:"listing_type"`
	Price       int64  `json:"price"`
}

type ListingListResponse struct {
	Result   bool      `json:"result"`
	Listings []Listing `json:"listings"`
}

type ListingResponse struct {
	Result  bool    `json:"result"`
	Listing Listing `json:"listing"`
}
