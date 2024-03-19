package main

import (
	"fmt"
	"os"

	"github.com/xagero/go-dotenv"
)

// import _ "github.com/xagero/go-dotenv/init"

func main() {

	dotenv.ReadFromFile("examples/.env")

	fmt.Println(os.Getenv("APP_ENV"))
	fmt.Println(os.Getenv("APP_SECRET"))
	fmt.Println(os.Getenv("DOMAIN"))
}
