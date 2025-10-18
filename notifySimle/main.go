package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

func every(n int, f func()) {

}

func main() {
	StartPeriodEyes := time.Now()
	StartPeriodBack := time.Now()

	for {
		now := time.Now()

		if now.Sub(StartPeriodEyes) >= 3*time.Second {
			fmt.Println(now.Format("2006-01-02 15:04:05"), ": eyes notify : ")

			e := exec.Command("notify-send", "Notify", "Glaza").Run()
			if e != nil {
				log.Println(e)
			}

			StartPeriodEyes = time.Now()
		}

		if now.Sub(StartPeriodBack) >= 5*time.Second {
			fmt.Println(now.Format("2006-01-02 15:04:05"), ": back notify : ")

			e := exec.Command("notify-send", "Stop work!", "Spina").Run()
			if e != nil {
				log.Println(e)
			}

			StartPeriodBack = time.Now()

		}

		time.Sleep(1 * time.Second)
	}
}
