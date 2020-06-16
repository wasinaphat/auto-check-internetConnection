package users

import (
	"fmt"
	_ "strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/wasinapl/bookstore_users-api/datasources/mysql/users_db"
	"github.com/wasinapl/bookstore_users-api/logger"
	_ "github.com/wasinapl/bookstore_users-api/utils/date_utils"
	"github.com/wasinapl/bookstore_users-api/utils/errors"
	"github.com/wasinapl/bookstore_users-api/utils/mysql_utils"
)
const (

	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser = " id, first_name,last_name,email,date_created,status  FROM users WHERE id=?;"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=? , email=? WHERE id=?;"
	queryDeleteUser = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus = "SELECT id,first_name, last_name, email , date_created , status FROM users WHERE status=?"
)


func something() {
	user := User{}
	if err := user.Get(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(user.FirstName)
}
func (user *User) Get() *errors.RestErr {

	stmt,err :=users_db.Client.Prepare(queryGetUser)
	if err !=nil{
		logger.Error("error when trying to prepare get user statement",err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	
	result := stmt.QueryRow(&user.Id)
	// fmt.Println(result)

     if getErr:= result.Scan(&user.Id,&user.FirstName,&user.LastName,&user.Email,&user.DateCreated,&user.Status); getErr!=nil{
		// return mysql_utils.ParseError(getErr)
		logger.Error("error when trying to prepare get user by id",err)
		return errors.NewInternalServerError("database error")

	}
	// fmt.Println(user)
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt,err:= users_db.Client.Prepare(queryInsertUser)
	if err !=nil {
		logger.Error("error when trying to prepare get user statement",err)

		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	// user.DateCreated = date_utils.GetNowDBFormat()

	insertResult, saveErr := stmt.Exec(user.FirstName,user.LastName,user.Email,user.DateCreated, user.Status,user.Password )
	if saveErr!=nil{
		logger.Error("error when trying to prepare save user",err)
		return errors.NewInternalServerError("database error")
	
	}

	userId ,err := insertResult.LastInsertId()

	if err !=nil{
		logger.Error("error when trying to get last insert id after creating a new user",err)
		return errors.NewInternalServerError("database error")
	}
	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr{
	
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err !=nil{
		logger.Error("error when trying prepare update user statement",err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_,err = stmt.Exec(user.FirstName,user.LastName,user.Email,user.Id)
	if err !=nil{
		logger.Error("error when trying prepare update user ",err)
		return errors.NewInternalServerError("database error")
	}
	
	return nil
}

func (user *User)Delete() *errors.RestErr {
	stmt, err:= users_db.Client.Prepare(queryDeleteUser)
	if err != nil{
		logger.Error("error when trying prepare delete user statement ",err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()


	if 	_,err = stmt.Exec(user.Id); err !=nil {
		logger.Error("error when trying prepare delete user",err)
		return errors.NewInternalServerError("database error")
	}

	return nil
}
func (user *User) FindByStatus(status string) ([]User,*errors.RestErr)  {

	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err !=nil{
		logger.Error("error when trying prepare find user by status",err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	rows, err :=stmt.Query(status)
	if err !=nil{
		logger.Error("error when trying prepare find user by status",err)
		return nil,errors.NewInternalServerError("database error")
	}
	defer rows.Close()


	results := make([]User,0)
	for rows.Next(){
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName ,&user.Email ,&user.DateCreated ,&user.Status); err!=nil{
			logger.Error("error when scan user row into user struct",err)
			return nil,mysql_utils.ParseError(err)
		}
		results = append(results,user)
	}
	if len(results) == 0 {
		return nil,errors.NewNotFoundError(fmt.Sprintf("no users matching status %s",status))
	}
	return results,nil


}