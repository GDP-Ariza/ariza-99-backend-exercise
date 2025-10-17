package adapter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"public_service/model"
	"time"
)

type UserAdapter interface {
	CreateUser(userRequest model.UserRequest) (model.User, error)
	GetUser(id int) (model.User, error)
}

type userAdapter struct {
	httpClient *http.Client
	baseURL    string
}

func NewUserAdapter() UserAdapter {
	baseURL := os.Getenv("USER_SERVICE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:7001" // Default user service URL
	}

	return &userAdapter{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		baseURL: baseURL,
	}
}

func (u *userAdapter) CreateUser(userRequest model.UserRequest) (model.User, error) {
	// Prepare form data
	formData := url.Values{}
	formData.Set("name", userRequest.Name)

	// Create POST request
	req, err := http.NewRequest("POST", u.baseURL+"/users", bytes.NewBufferString(formData.Encode()))
	if err != nil {
		return model.User{}, fmt.Errorf("failed to create request: %w", err)
	}

	// Set content type
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Make the request
	resp, err := u.httpClient.Do(req)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusCreated {
		return model.User{}, fmt.Errorf("user service returned status %d", resp.StatusCode)
	}

	// Parse response
	var userResponse struct {
		Result bool       `json:"result"`
		User   model.User `json:"user"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&userResponse); err != nil {
		return model.User{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return userResponse.User, nil
}

func (u *userAdapter) GetUser(id int) (model.User, error) {
	// Create GET request
	url := fmt.Sprintf("%s/users/%d", u.baseURL, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to create request: %w", err)
	}

	// Make the request
	resp, err := u.httpClient.Do(req)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		return model.User{}, fmt.Errorf("user service returned status %d", resp.StatusCode)
	}

	// Parse response
	var userResponse struct {
		Result bool       `json:"result"`
		User   model.User `json:"user"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&userResponse); err != nil {
		return model.User{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return userResponse.User, nil
}
