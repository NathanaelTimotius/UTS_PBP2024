package models

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Account struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type RoomOri struct {
	ID        int    `json:"id"`
	Room_name string `json:"room_name"`
	ID_game   int    `json:"id_game"`
}

type RoomsResponse2 struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    RoomOri `json:"data"`
}

type Room struct {
	ID        int    `json:"id"`
	Room_name string `json:"room_name"`
}

type RoomResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Room   `json:"data, omitempty"`
}

type RoomsResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Room `json:"data"`
}

type DetailRoom struct {
	ID           int           `json:"id"`
	Room_name    string        `json:"room_name"`
	Participants []Participant `json:"participants"`
}

type DetailRoomResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    DetailRoom `json:"data"`
}

type Participant struct {
	ID         int    `json:"id"`
	ID_account int    `json:"id_account"`
	Username   string `json:"username"`
}
