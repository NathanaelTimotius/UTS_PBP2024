package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	m "UTS/models"
)

func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	params := r.URL.Query()
	gameIDStr := params.Get("gameID")
	var query string

	if gameIDStr == "" {
		sendErrorResponse(w, "game id is not defined")
		return
	}

	gameID, err := strconv.Atoi(gameIDStr)
	if err != nil {
		sendErrorResponse(w, "game id is not a number")
		return
	}

	query = `SELECT id, room_name
			FROM rooms 
            WHERE id_game = ?`
	roomRow, err := db.Query(query, gameID)
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, "query error")
		return
	}

	var room m.Room
	var rooms []m.Room
	for roomRow.Next() {
		if err := roomRow.Scan(&room.ID, &room.Room_name); err != nil {
			log.Println(err)
			sendErrorResponse(w, "Error getting room")
			return
		} else {
			rooms = append(rooms, room)
		}
	}

	var response m.RoomsResponse
	response.Status = 200
	response.Message = "Success!"
	response.Data = rooms
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetDetailRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	params := r.URL.Query()
	roomIDStr := params.Get("gameID")
	var query string

	if roomIDStr == "" {
		sendErrorResponse(w, "game id is not defined")
		return
	}

	roomID, err := strconv.Atoi(roomIDStr)
	if err != nil {
		sendErrorResponse(w, "game id is not a number")
		return
	}

	query = `SELECT r.id, r.room_name, p.id, p.id_account, a.username
			FROM rooms r
			JOIN participants p ON r.id = p.id_room
			JOIN accounts a ON p.id_account = a.id
            WHERE r.id_game = ?`

	detailRoomRow, err := db.Query(query, roomID)
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, "query error")
		return
	}

	var detailRoom m.DetailRoom
	var participants []m.Participant
	for detailRoomRow.Next() {
		var participant m.Participant
		if err := detailRoomRow.Scan(&detailRoom.ID, &detailRoom.Room_name,
			&participant.ID, &participant.ID_account, &participant.Username); err != nil {
			log.Println(err)
			sendErrorResponse(w, "Error getting room details")
			return
		}
		participants = append(participants, participant)
	}

	detailRoom.Participants = participants

	var response m.DetailRoomResponse
	response.Status = 200
	response.Message = "Success!"
	response.Data = detailRoom
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func InsertRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	params := r.URL.Query()
	gameIDStr := params.Get("gameID")
	name := params.Get("name")
	var query string

	if gameIDStr == "" {
		sendErrorResponse(w, "game id is not defined")
		return
	}

	gameID, err := strconv.Atoi(gameIDStr)
	if err != nil {
		sendErrorResponse(w, "game id is not a number")
		return
	}

	query = `SELECT COUNT(*) FROM rooms WHERE id_game = ?`
	var roomCount int
	err = db.QueryRow(query, gameID).Scan(&roomCount)
	if err != nil {
		sendErrorResponse(w, "Error getting count")
		return
	}

	query = `SELECT max_player
			FROM games
            WHERE id = ?`
	var maxPlayer int
	err = db.QueryRow(query, gameID).Scan(&maxPlayer)
	if err != nil {
		sendErrorResponse(w, "Error getting max_player")
		return
	}

	if roomCount == maxPlayer {
		sendErrorResponse(w, "Room is full")
		return
	}

	var room m.RoomOri
	room.Room_name = name
	room.ID_game = gameID
	insertQuery := "INSERT INTO rooms (room_name, id_game) VALUES (?, ?)"
	_, err = db.Exec(insertQuery, room.Room_name, room.ID_game)
	if err != nil {
		sendErrorResponse(w, "inserting error")
		return
	}

	var response m.RoomsResponse2
	response.Status = 200
	response.Message = "Success"
	response.Data = room
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func LeaveRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	idAccount, _ := strconv.Atoi(r.Form.Get("idAcc"))
	idRoom, _ := strconv.Atoi(r.Form.Get("idRoom"))

	deleteQuery := "DELETE FROM participants WHERE id_account = ? AND id_room = ?"
	_, err := db.Exec(deleteQuery, idAccount, idRoom)
	if err != nil {
		sendErrorResponse(w, "Query Error")
		return
	}

	var response m.RoomResponse
	response.Status = 200
	response.Message = "Success"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func sendErrorResponse(w http.ResponseWriter, pesan string) {
	var response m.ErrorResponse
	response.Status = 400
	response.Message = pesan
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
