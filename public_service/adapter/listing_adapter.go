package adapter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"public_service/model"
	"strconv"
	"time"
)

type ListingAdapter interface {
	GetListing(page, size, userId int) ([]model.ListingModel, error)
	CreateListing(request model.ListingRequest) (model.ListingModel, error)
}

type listingAdapter struct {
	httpClient *http.Client
	baseURL    string
}

func NewListingAdapter() ListingAdapter {
	baseURL := os.Getenv("LISTING_SERVICE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:6000" // Default listing service URL
	}

	return &listingAdapter{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		baseURL: baseURL,
	}
}

func (l *listingAdapter) GetListing(page, size, userId int) ([]model.ListingModel, error) {
	// Build URL with query parameters
	reqURL := fmt.Sprintf("%s/listings", l.baseURL)
	params := url.Values{}
	params.Set("page_num", strconv.Itoa(page))
	params.Set("page_size", strconv.Itoa(size))
	if userId > 0 {
		params.Set("user_id", strconv.Itoa(userId))
	}

	if len(params) > 0 {
		reqURL += "?" + params.Encode()
	}

	// Create GET request
	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Make the request
	resp, err := l.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("listing service returned status %d", resp.StatusCode)
	}

	// Parse response
	var listingResponse struct {
		Result   bool                 `json:"result"`
		Listings []model.ListingModel `json:"listings"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&listingResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return listingResponse.Listings, nil
}

func (l *listingAdapter) CreateListing(request model.ListingRequest) (model.ListingModel, error) {
	// Prepare form data
	formData := url.Values{}
	formData.Set("user_id", strconv.Itoa(request.UserID))
	formData.Set("listing_type", request.ListingType)
	formData.Set("price", strconv.FormatInt(request.Price, 10))

	// Create POST request
	req, err := http.NewRequest("POST", l.baseURL+"/listings", bytes.NewBufferString(formData.Encode()))
	if err != nil {
		return model.ListingModel{}, fmt.Errorf("failed to create request: %w", err)
	}

	// Set content type
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Make the request
	resp, err := l.httpClient.Do(req)
	if err != nil {
		return model.ListingModel{}, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	// Parse response
	var listingResponse struct {
		Result  bool               `json:"result"`
		Listing model.ListingModel `json:"listing"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&listingResponse); err != nil {
		return model.ListingModel{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return listingResponse.Listing, nil
}
