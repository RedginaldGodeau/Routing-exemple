package main

import (
	"api/ent"
	"api/internal/route"
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	client, err := ent.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD")))
	if err != nil {
		log.Println("Database error:", err)
		return
	}

	r := echo.New()

	route.Route(client, r)

	err = r.Start(":8080")
	if err != nil {
		log.Println("Routing error:", err)
		return
	}
}
