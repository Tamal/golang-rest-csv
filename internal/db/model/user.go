package model

import (
	"encoding/json"
	"fmt"
	"strings"
)

type User struct {
	Id        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Age       int    `json:"age,omitempty"`
	Phone     string `json:"phone,omitempty"`
	CreatedAt int64  `json:",omitempty"`
}

func (u *User) IsValid() bool {
	return len(strings.TrimSpace(u.Name)) > 0 && len(strings.TrimSpace(u.Phone)) > 0 && u.Age > 0
}

func (u *User) String() string {
	return fmt.Sprintf("Id:%s, Name: %s, Age: %d, Phone: %s", u.Id, u.Name, u.Age, u.Phone)
}

func (u *User) ToCSV() string {
	return fmt.Sprintf("%s, %s, %d, %s", u.Id, u.Name, u.Age, u.Phone)
}

func (u *User) ToJson() (string, error) {
	if u, err := json.Marshal(u); err != nil {
		return string(u), nil
	} else {
		return "", err
	}
}
