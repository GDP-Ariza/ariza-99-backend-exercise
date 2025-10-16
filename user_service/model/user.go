package model

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"` // in microseconds
	UpdatedAt int64  `json:"updated_at"` // in microseconds
}

type UserListResponse struct {
	Result bool   `json:"result"`
	Users  []User `json:"users"`
}

type UserResponse struct {
	Result bool `json:"result"`
	User   User `json:"user"`
}
