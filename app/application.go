package app
import (
	"github.com/gin-gonic/gin"
	_"github.com/wasinapl/bookstore_users-api/logger"
	"fmt"
    "log"
    "os/exec"
    "github.com/play175/wifiNotifier"
)
var (
	router = gin.Default()
)
func StartApplication() {

	mapUrls()
	logger.Info("about to start the application....")
	router.Run(":8080")
	// wifiNotifier.SetWifiNotifier(func(ssid string) {
    //     log.Println("onWifiChanged,current ssid:" + ssid)
        

    // })

    // log.Println("current ssid:" + wifiNotifier.GetCurrentSSID())
    // if err := wifiNotifier.GetCurrentSSID(); err ==""{
    //     fmt.Println("Internet Disconnect")
    //     if err := exec.Command("cmd", "/C", "shutdown", "/s").Run(); err != nil {
    //         fmt.Println("Failed to initiate shutdown:", err)
    //     }
    // }
    // for {

    // }

}