package main

import (
	"ai_firewall/internal/api"
)

func main() {
	r := api.SetupRouter()
	r.Run(":3300")
}