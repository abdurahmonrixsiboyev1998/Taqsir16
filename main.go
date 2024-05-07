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
	postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	num, err := sql.Open("postgres", postgresqlDbInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer num.Close()

	err = num.Ping()
	if err != nil {
		log.Fatal("failed to ping", err)
	}

	fmt.Println("successfully connected")
	send(num, 1, 2)
	// accept(num,4,5)
	// block(num,3,5)
	// unblock(num,3,5)

}

func send(num *sql.DB, jonatuvchiID int, qabul_qiluvchiID int) {
	_, err := num.Exec("INSERT INTO friendships (user_id, friend_id, status) VALUES ($1, $2, $3)", jonatuvchiID, qabul_qiluvchiID, "requested")
	if err != nil {
		fmt.Println("Error sending friend request:", err)
	}
}

func Accept(num *sql.DB, jonatuvchiID int, qabul_qiluvchiID int) {
	_, err := num.Exec("UPDATE friendships SET status = $1 WHERE user_id = $2 AND friend_id = $3", "accepted", qabul_qiluvchiID, jonatuvchiID)
	if err != nil {
		fmt.Println("Error accepting friend request:", err)
	}
}

func Block(num *sql.DB, userID int, blockedUserID int) {
	_, err := num.Exec("UPDATE friendships SET status = $1 WHERE user_id = $2 AND friend_id = $3", "blocked", userID, blockedUserID)
	if err != nil {
		fmt.Println("Error blocking user:", err)
	}
}

func Unblock(num *sql.DB, userID int, unblockedUserID int) {
	_, err := num.Exec("UPDATE friendships SET status = $1 WHERE user_id = $2 AND friend_id = $3", "requested", userID, unblockedUserID)
	if err != nil {
		fmt.Println("Error unblocking user:", err)
	}
}
