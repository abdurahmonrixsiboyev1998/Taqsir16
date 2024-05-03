package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "14022014"
	dbname   = "demo"
)

func main() {
	postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", postgresqlDbInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("failed to ping", err)
	}

	fmt.Println("successfully connected")

	err = sendFriendRequest(db, 4, 5)
	if err != nil {
		log.Println("Error sending friend request:", err)
	}

	err = acceptFriendRequest(db, 4, 5)
	if err != nil {
		log.Println("Error accepting friend request:", err)
	}

	err = blockUser(db, 3, 5)
	if err != nil {
		log.Println("Error blocking user:", err)
	}

	err = unblockUser(db, 3, 5)
	if err != nil {
		log.Println("Error unblocking user:", err)
	}
}

func sendFriendRequest(db *sql.DB, senderID, receiverID int) error {
	_, err := db.Exec("INSERT INTO friendships (user_id, friend_id, status) VALUES ($1, $2, $3)", senderID, receiverID, "requested")
	if err != nil {
		return err
	}
	return nil
}

func acceptFriendRequest(db *sql.DB, senderID, receiverID int) error {
	_, err := db.Exec("UPDATE friendships SET status = $1 WHERE user_id = $2 AND friend_id = $3", "accepted", receiverID, senderID)
	if err != nil {
		return err
	}
	return nil
}

func blockUser(db *sql.DB, userID, blockedUserID int) error {
	_, err := db.Exec("UPDATE friendships SET status = $1 WHERE user_id = $2 AND friend_id = $3", "blocked", userID, blockedUserID)
	if err != nil {
		return err
	}
	return nil
}

func unblockUser(db *sql.DB, userID, unblockedUserID int) error {
	_, err := db.Exec("UPDATE friendships SET status = $1 WHERE user_id = $2 AND friend_id = $3", "requested", userID, unblockedUserID)
	if err != nil {
		return err
	}
	return nil
}
