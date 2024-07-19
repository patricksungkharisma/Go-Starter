package entity

type GetExampleRequest struct {
	ID int64
}

type GetExampleResponse struct {
	ID       int64  `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	IsActive bool   `json:"is_active" db:"is_active"`
}