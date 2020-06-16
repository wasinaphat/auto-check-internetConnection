package app
import (
	"fmt"
	"net"
	"os/exec"
	"time"
	"os"
    "os/signal"
	"github.com/robfig/cron"
	
)
func RunEverySecond() {
	fmt.Printf("%v\n", time.Now())
	hostName := "10.5.87.12"
	portNum := "80"
	seconds := 5
	timeOut := time.Duration(seconds) * time.Second
  
	conn, err := net.DialTimeout("tcp", hostName+":"+portNum, timeOut)
  
	if err != nil {
	   fmt.Println(err)
	   if err := exec.Command("cmd", "/C", "shutdown", "/r").Run(); err != nil {
		fmt.Println("Failed to initiate restart:", err)
	}
	   return
	}
  
	fmt.Printf("Connection established between %s and localhost with time out of %d seconds.\n", hostName, int64(seconds))
	fmt.Printf("Remote Address : %s \n", conn.RemoteAddr().String())
	fmt.Printf("Local Address : %s \n", conn.LocalAddr().String())
}
func StartApplication() {
	c := cron.New()
    c.AddFunc("*/5 * * * * *", RunEverySecond)
    go c.Start()
    sig := make(chan os.Signal)
    signal.Notify(sig, os.Interrupt, os.Kill)
    <-sig
	
}
