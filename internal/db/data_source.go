package db

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/emp/internal/db/model"
	"github.com/google/uuid"
)

type UserService interface {
	GetUsers() ([]model.User, error)
	GetUser(id string) (model.User, error)
	AddUser(newUser model.User) (*model.User, error)
	DeleteUser(id string) (bool, error)
}

type UserDataService struct {
	FileName string
}

func (u *UserDataService) GetUsers() ([]model.User, error) {

	records, err := readFile(u.FileName)
	if err != nil {
		return nil, err
	}

	var users []model.User

	for _, record := range records {

		age, _ := strconv.Atoi(record[2])
		user := model.User{
			Id:    record[0],
			Name:  record[1],
			Age:   age,
			Phone: record[3],
		}

		fmt.Println(user.String())
		users = append(users, user)
	}

	return users, nil
}

func (u *UserDataService) GetUser(id string) (model.User, error) {
	var user model.User

	records, err := readFile(u.FileName)
	if err != nil {
		return user, err
	}

	for _, record := range records {

		if record[0] == id {
			age, _ := strconv.Atoi(record[2])
			user = model.User{
				Id:    record[0],
				Name:  record[1],
				Age:   age,
				Phone: record[3],
			}

			fmt.Println(user.String())
			break
		}
	}

	return user, nil

}

func (u *UserDataService) AddUser(user model.User) (*model.User, error) {
	if !user.IsValid() {
		return nil, errors.New("Invalid user")
	}
	records, err := readFile(u.FileName)
	if err != nil {
		fmt.Println("Failed to read.", err)
		return nil, fmt.Errorf("Failed to read file. %s", err)
	}

	var users []model.User
	for _, record := range records {

		age, _ := strconv.Atoi(record[2])
		createdAt, _ := strconv.ParseInt(record[4], 10, 64)
		user := model.User{
			Id:        record[0],
			Name:      record[1],
			Age:       age,
			Phone:     record[3],
			CreatedAt: createdAt,
		}

		fmt.Println(user.String())
		users = append(users, user)
	}

	user.Id = uuid.NewString()
	user.CreatedAt = time.Now().Unix()
	users = append(users, user)

	writeFile(u.FileName, users)

	return &user, nil

}

func (u *UserDataService) DeleteUser(id string) (bool, error) {
	var user model.User

	records, err := readFile(u.FileName)
	if err != nil {
		return false, err
	}

	var users []model.User
	var found bool
	for _, record := range records {

		if record[0] == id {
			found = true
			continue
		} else {
			age, _ := strconv.Atoi(record[2])
			createdAt, _ := strconv.ParseInt(record[4], 10, 64)
			user = model.User{
				Id:        record[0],
				Name:      record[1],
				Age:       age,
				Phone:     record[3],
				CreatedAt: createdAt,
			}

			fmt.Println(user.String())
			users = append(users, user)
		}
	}

	if found {
		writeFile(u.FileName, users)
	}

	return found, nil

}

func readFile(FileName string) ([][]string, error) {
	// Open the file

	var f, err = os.OpenFile(FileName, os.O_RDONLY, 0644)

	if os.IsNotExist(err) {
		// handle the case where the file doesn't exist
		f, err = os.Create(FileName)
	}

	defer f.Close()
	r := csv.NewReader(f)

	// Read the file
	records, err := r.ReadAll()

	if err != nil {
		return nil, err
	}

	return records, nil
}

func writeFile(FileName string, users []model.User) error {
	// Open the file
	f, err := os.Create(FileName)

	if err != nil {
		return err
	}

	defer f.Close()

	var records [][]string

	for _, record := range users {

		s := []string{
			record.Id,
			record.Name,
			strconv.Itoa(record.Age),
			record.Phone,
			strconv.FormatInt(record.CreatedAt, 10),
		}
		records = append(records, s)
	}

	w := csv.NewWriter(f)
	err = w.WriteAll(records) // calls Flush internally

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
