package main

import (
	"github.com/NickNaskida/GoTrim/api"
)

func main() {
	router := api.SetupRouter()
	router.Run(":8080")
}
