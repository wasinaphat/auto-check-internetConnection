package users
import (
	// "encoding/json"
	"fmt"
	"strconv"
	// "io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wasinapl/bookstore_users-api/domain/users"
	"github.com/wasinapl/bookstore_users-api/services"
	"github.com/wasinapl/bookstore_users-api/utils/errors"
)

func getUserId(userIdParam string)(int64, *errors.RestErr){
	userId, userErr := strconv.ParseInt(userIdParam,10,64)
	if userErr !=nil{
		return 0, errors.NewBadRequestError("user id should be a number")
	
	}
	return userId,nil
}
func Create(c *gin.Context){
	var user users.User
	// c.ShouldBindJSON is insert into user model
	if err :=c.ShouldBindJSON(&user); err !=nil{
		restErr :=errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}


	result ,saveErr := services.UsersService.CreateUser(user)
	if saveErr !=nil{
		c.JSON(saveErr.Status,saveErr)
		return
	}
	
	fmt.Println(result)
	c.JSON(http.StatusCreated,result.Marshall(c.GetHeader("X-Public")=="true"))
	
}
func Get(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr !=nil{
		c.JSON(idErr.Status,idErr)
		return
	}

	user,getErr := services.UsersService.GetUser(userId)
	if getErr !=nil{
		c.JSON(getErr.Status,getErr)
		return
	}
	

	c.JSON(http.StatusCreated,user.Marshall(c.GetHeader("X-Public")=="true"))
	

}

func SearchUser(c *gin.Context){
	c.String(http.StatusNotImplemented,"implement me!")
	
}

func Update(c *gin.Context)  {
	userId, userErr := strconv.ParseInt(c.Param("user_id"),10,64)
	if userErr !=nil{
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status,err)
	}
	var user users.User
	// c.ShouldBindJSON is insert into user model
	if err :=c.ShouldBindJSON(&user); err !=nil{
		restErr :=errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch


	result, err :=services.UsersService.UpdateUser(isPartial, user); 
	if err !=nil{
		c.JSON(err.Status,err)
		return
	}
	c.JSON(http.StatusOK,result.Marshall(c.GetHeader("X-Public")=="true"))
	
}
func Delete(c *gin.Context){

	userId, idErr := getUserId(c.Param("user_id"))
	if idErr !=nil{
		c.JSON(idErr.Status,idErr)
		return
	}


	if err := services.UsersService.DeleteUser(userId); err !=nil{
		c.JSON(err.Status,err)
		return
	}
	c.JSON(http.StatusOK,map[string]string{"status":"deleted"})

	
}
func Search(c *gin.Context) {
	status := c.Query("status")

	users,err := services.UsersService.SearchUser(status)
	if err !=nil{
		c.JSON(err.Status,err)
		return
	}



	c.JSON(http.StatusOK,users.Marshall(c.GetHeader("X-Public")=="true"))
}