package ping
import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func Ping(c *gin.Context){	
	c.String(http.StatusOK, "Hello, World!")
	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "pong",
	// })
}