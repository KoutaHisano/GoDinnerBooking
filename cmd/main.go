package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/KoutaHisano/GoDinnerBooking/internal/db"
	"github.com/KoutaHisano/GoDinnerBooking/internal/models"
)

func main() {
	database, err := db.NewDB("godinnerbooking.db")
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 2 {
		fmt.Println("No command provided")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "create-user":
		if len(os.Args) != 4 {
			fmt.Println("Usage: cli create-user [username] [password]")
			return
		}
		username := os.Args[2]
		password := os.Args[3]

		// models.HashPassword 関数を呼び出し
		hashedPassword, err := models.HashPassword(password)
		if err != nil {
			log.Fatal(err) // ハッシュ化中のエラーを処理
		}

		newUser := models.User{
			LoginID:  username,
			Password: hashedPassword,
		}

		err = models.CreateUser(database, newUser)
		if err != nil {
			log.Fatal("Failed to create user:", err)
		}
		fmt.Println("User created successfully")

	case "get-user":
		if len(os.Args) != 3 {
			fmt.Println("Usage: cli get-user [userID]")
			return
		}

		userID, err := strconv.Atoi(os.Args[2]) // userID変数名を正しく修正
		if err != nil {
			log.Fatal("Invalid user ID:", err) // ID変換中のエラーを処理
		}

		user, err := models.GetUser(database, userID)
		if err != nil {
			log.Fatal("Failed to get user:", err) // ユーザー情報取得中のエラーを処理
		}

		fmt.Printf("ID: %d, Username: %s\n", user.ID, user.LoginID) // Printfを正しく使用

	default:
		fmt.Println("Unknown command")
	}
}
