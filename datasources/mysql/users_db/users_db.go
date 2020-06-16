package users_db

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

const (
	// mysql_users_username = "root"
	// mysql_users_password = "NewdividE1*"
	// mysql_users_host = "127.0.0.1:3306"
	// mysql_users_schema = "users_db"

	mysql_users_username = "mysql_users_username"
	mysql_users_password = "mysql_users_password"
	mysql_users_host = "mysql_users_host"
	mysql_users_schema = "mysql_users_schema"

)

var (
	Client *sql.DB

	username = os.Getenv(mysql_users_username)
	password = os.Getenv(mysql_users_password)
	host = os.Getenv(mysql_users_host)
	schema = os.Getenv(mysql_users_schema)
	
)
func init(){
	// var databasePass string
	// databasePass= os.Getenv("DATABASE_PASS")
    // fmt.Printf("Database Password: %s\n", databasePass)
	var username = "root"
	var password = "NewdividE1*"
	var host = "127.0.0.1:3306"
	var schema = "users_db"
	datasourceName:=fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",username,password,host,schema)
	var err error
	Client,err =sql.Open("mysql",datasourceName)

	if err !=nil{
		
		panic(err)
		
	}
	if err = Client.Ping(); err !=nil{
		panic(err)
	}
	
	
	log.Println("database successfully congifured")
}