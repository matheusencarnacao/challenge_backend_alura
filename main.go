package main

import (
	"challenge-milhas/routes"
	"fmt"
)

func main() {
	fmt.Println("Initializing Milhas API")
	routes.HandleRequests()
}
