package main

import (
	"github/ludo62/bank_db/api"
)

func main() {
	server := api.NewServer(".")
	server.Start(3000)
}
