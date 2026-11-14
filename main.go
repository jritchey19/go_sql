package main

import (
	"fmt"
)

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
}

func main() {
	db, err := CreateDB("app_db.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.DBClose()

	fmt.Println("Success starting db.")
	user1, err := db.CreateUser("test1", "test1@test.com")
	if err != nil {
		fmt.Println(err)
		fmt.Println(IsEmailUniqueErr(err))
		return
	} else {
		fmt.Printf("Success adding user test1 with id: %v\n", user1)
	}

	userStct1, err := db.GetUser(user1)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Got User:\n\tid: %v\n\tname: %v\n\temail: %v\n", userStct1.ID, userStct1.Name, userStct1.Email)
	}

	err = db.UpdateUser(user1, "changed name", "test1@test.com")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		userStct1, err := db.GetUser(user1)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Printf("Got User:\n\tid: %v\n\tname: %v\n\temail: %v\n", userStct1.ID, userStct1.Name, userStct1.Email)
		}
	}
}
