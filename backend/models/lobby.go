package models

type CreateLobbyRequest struct {
	Name string `json:"name" validate:"required"`
}

type CreateLobbyResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Owner Player `json:"owner"`
}
