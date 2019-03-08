package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Arguments map[string]string

func main() {
	err := Perform(parseArgs(), os.Stdout)
	if err != nil {
		panic(err)
	}
}

type User struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func parseArgs() Arguments {
	operation := flag.String("operation", "", "Operation")
	filename := flag.String("fileName", "", "Filename")
	item := flag.String("item", "", "Item of User")
	id := flag.String("id", "", "User id")
	flag.Parse()
	return Arguments{
		"operation": *operation,
		"fileName":  *filename,
		"item":      *item,
		"id":        *id,
	}
}

func Perform(args Arguments, writer io.Writer) (err error) {
	operation := args["operation"]
	if operation == "" {
		return fmt.Errorf("-operation flag has to be specified")
	}
	filename := args["fileName"]
	if filename == "" {
		return fmt.Errorf("-fileName flag has to be specified")
	}
	switch operation {
	case "add":
		item := args["item"]
		if item == "" {
			return fmt.Errorf("-item flag has to be specified")
		}
		var user User
		err = json.Unmarshal([]byte(item), &user)
		if err != nil {
			return
		}
		var users []User
		var bytes []byte
		bytes, err = ioutil.ReadFile(filename)
		if err != nil {
			return
		}
		err = json.Unmarshal(bytes, &users)
		if err != nil {
			return
		}
		for _, userFromFile := range users {
			if userFromFile.Id == user.Id {
				errorMassange := "Item with id " + user.Id + " already exists"
				_, err = writer.Write([]byte(errorMassange))
				return
			}
		}
		users = append(users, user)
		bytes, err = json.Marshal(users)
		if err != nil {
			return
		}
		return ioutil.WriteFile(filename, bytes, 128)

	case "findById":
		id := args["id"]
		if id == "" {
			return errors.New("-id flag has to be specified")
		}
		var users []User
		var bytesUsers []byte
		bytesUsers, err = ioutil.ReadFile(filename)
		if err != nil {
			return err
		}
		err = json.Unmarshal(bytesUsers, &users)
		if err != nil {
			return err
		}
		var user User
		isUserExist := false
		for _, value := range users {
			if value.Id == id {
				user = value
				isUserExist = true
				break
			}
		}
		bytesUsers = []byte("")
		if isUserExist {
			bytesUsers, err = json.Marshal(user)
			if err != nil {
				return err
			}
		}
		_, err = writer.Write(bytesUsers)

	case "remove":
		id := args["id"]
		if id == "" {
			return errors.New("-id flag has to be specified")
		}
		var users []User
		var bytesUsers []byte
		bytesUsers, err = ioutil.ReadFile(filename)
		if err != nil {
			return err
		}
		err = json.Unmarshal(bytesUsers, &users)
		if err != nil {
			return err
		}
		var usersAfterRemove []User
		isUserRemoved := false
		for i, value := range users {
			if value.Id == id {
				usersAfterRemove = append(users[:i], users[i+1:]...)
				isUserRemoved = true
				break
			}
		}
		if !isUserRemoved {
			errStr := "Item with id " + id + " not found"
			_, err = writer.Write([]byte(errStr))
			return err
		}
		bytesUsers, err = json.Marshal(usersAfterRemove)
		if err != nil {
			return err
		}
		return ioutil.WriteFile(filename, bytesUsers, 128)

	case "list":
		var usersFromFile []byte
		usersFromFile, err = ioutil.ReadFile(filename)
		if err != nil {
			return err
		}
		_, err = writer.Write(usersFromFile)

	default:
		return fmt.Errorf("Operation %s not allowed!", operation)
	}
	return
}
