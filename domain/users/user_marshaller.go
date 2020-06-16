package users

import (
	// "github.com/gin-gonic/gin/internal/json"
	"encoding/json"
)

type PublicUser struct{
	Id int64 `json:"id"`
	// FirstName string `json:"first_name"`
	// LastName string `json:"last_name"`
	// Email string `json:"email"`
	DateCreated string `json:"date_created"`
	Status string `json:"status"`
	// Password string `json:"password"`
}



type PrivateUser struct{
	Id int64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	DateCreated string `json:"date_created"`
	Status string `json:"status"`
	// Password string `json:"password"`
}

func (users Users) Marshall(isPublic bool) []interface{} {
	result :=make([]interface{},len(users))
	for index, user:= range users{
		result[index]= user.Marshall(isPublic)
	}
	return result
}

func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic{
		return PublicUser{
			Id: user.Id,
			DateCreated: user.DateCreated,
			Status: user.Status,
		}
	}
	userJson,_ := json.Marshal(user)
	var PrivateUser PrivateUser
	json.Unmarshal(userJson,&PrivateUser)
	return PrivateUser

	
}