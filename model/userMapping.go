package model

import "go/chat/database"

func SetUser(userID string, SERVERID string) {
	query := "INSERT INTO user_mapping(user_id,server_id)VALUES(?,?)"
	database.ExecuteQuery(query, userID, SERVERID)
}

func GetServerId(userID string) string {
	var ServerId string
	// var ids []string
	query := `SELECT server_id FROM user_mapping WHERE user_id =?`
	iter := database.Connection.Session.Query(query, userID).Iter()
	iter.Scan(&ServerId)

	return ServerId
}