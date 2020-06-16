package app
import (
	"fmt"
	"net"
	"os/exec"
	"time"
	"os"
    "os/signal"
	"github.com/robfig/cron"
	"log"
)

	

type countdown struct {
	t int
	d int
	h int
	m int
	s int
}
type Month struct{
	CheckMonth string
}
func getTimeRemaining(t time.Time) countdown {
	currentTime := time.Now()
	difference := t.Sub(currentTime)

	total := int(difference.Seconds())
	days := int(total / (60 * 60 * 24))
	hours := int(total / (60 * 60) % 24)
	minutes := int(total/60) % 60
	seconds := int(total % 60)

	return countdown{
		t: total,
		d: days,
		h: hours,
		m: minutes,
		s: seconds,
	}
}
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
	// c := cron.New()
    // c.AddFunc("*/5 * * * * *", RunEverySecond)
    // go c.Start()
    // sig := make(chan os.Signal)
    // signal.Notify(sig, os.Interrupt, os.Kill)
	// <-sig

	// fmt.Println("Program is starting now.....")
	// stop := Every(5*time.Minute, func(time.Time) bool {
	// 	c := cron.New()
    // c.AddFunc("*/5 * * * * *", RunEverySecond)
    // go c.Start()
    // sig := make(chan os.Signal)
    // signal.Notify(sig, os.Interrupt, os.Kill)
	// <-sig
    //     log.Println("ticker")
    //     return true
    // })

    // time.Sleep(168 * time.Hour)
    // log.Println("stopping ticker")
    // stop <- true

	today := time.Now()
 
	if today.Year() ==2020 && today.Month() == time.June && today.Day()<=22 {
			fmt.Println("Program is starting now.....")
	


	stop := Every(5*time.Minute, func(time.Time) bool {
	
		c := cron.New()
    c.AddFunc("*/30 * * * * *", RunEverySecond)
    go c.Start()
    sig := make(chan os.Signal)
    signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig

        log.Println("ticker")
        return true
    })

    time.Sleep(168 * time.Hour)
    log.Println("stopping ticker")
    stop <- true

	}
	

}
func Every(duration time.Duration, work func(time.Time) bool) chan bool {
    ticker := time.NewTicker(duration)
    stop := make(chan bool, 1)

    go func() {
        defer log.Println("ticker stopped")
        for {
            select {
            case time := <-ticker.C:
                if !work(time) {
                    stop <- true
                }
            case <-stop:
                return
            }
        }
    }()

    return stop
}
