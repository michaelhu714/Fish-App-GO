package models

type Lobby struct {
	Id         int
	Name       string
	Owner      Player
	NumPlayers int
	Players    []Player
}

type CreateLobbyReq struct {
	Name       string `json:"name" validate:"required"`
	NumPlayers int    `json:"num_players"` // default is 6
}

type CreateLobbyResp struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Owner Player `json:"owner"`
}

type JoinLobbyReq struct {
	Id   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type JoinLobbyResp struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
