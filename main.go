package main

import (
	"fmt"
	"os/exec"
	"strings"
	"net/http"
)

func main() {
	state := false
	for {
		out, _ := exec.Command("ping", "192.168.1.81", "-c 3", "-i 1", "-w 10").Output()
		if state == true && strings.Contains(string(out), "Destination Host Unreachable") {
			http.Get("http://localhost:5005/Master%20Bedroom/pause")
    		fmt.Println("Pause")
			state = false
		} else if state == false && !strings.Contains(string(out), "Destination Host Unreachable") {
			http.Get("http://localhost:5005/Master%20Bedroom/play")
			fmt.Println("Play")
			state = true
		}
	}
}
