package main

import (
	"docker-dashboard-cli/cmd"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "dashboard" {
		cmd.Dashboard()
	} else {
		fmt.Println("Usage: dockinfo dashboard") 
	}
}