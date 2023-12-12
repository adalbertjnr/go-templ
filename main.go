package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func doHTTPCall(withFormat bool, date ...string) (*Users, error) {
	uri := "https://fakerapi.it/api/v1/persons?_quantity=10&_gender=male&_birthday_start=2005-01-01"

	if withFormat {
		uri = fmt.Sprintf("https://fakerapi.it/api/v1/persons?_quantity=%s&_gender=male&_birthday_start=%s&_birthday_end=%s", date[2], date[0], date[1])
	}

	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}

	userResp := Users{}

	if err := json.NewDecoder(resp.Body).Decode(&userResp); err != nil {
		return nil, err
	}
	return &userResp, nil
}

func tGet(c *fiber.Ctx) error {

	allUsers, err := doHTTPCall(false)
	if err != nil {
		return err
	}

	length := len(allUsers.Data)
	allImages, err := imageSeparator(length)
	if err != nil {
		return err
	}

	return adaptor.HTTPHandler(templ.Handler(Main(*allUsers, allImages)))(c)
}

func imageSeparator(length int) ([]string, error) {

	var (
		r      = rand.New(rand.NewSource(time.Now().UnixNano()))
		images = make([]string, 0)
	)

	entries, err := os.ReadDir("./public/assets/people")
	if err != nil {
		return nil, err
	}

	for i := 0; i < length; i++ {
		randomIndex := r.Intn(len(entries))
		if len(images) < length {
			images = append(images, entries[randomIndex].Name())
		}
	}
	return images, nil
}

func tPost(c *fiber.Ctx) error {

	var (
		start_date    = c.FormValue("start_date")
		end_date      = c.FormValue("end_date")
		people_amount = c.FormValue("people_amount")
	)

	allUsers, err := doHTTPCall(true, start_date, end_date, people_amount)
	if err != nil {
		return err
	}

	length := len(allUsers.Data)

	allImages, err := imageSeparator(length)
	if err != nil {
		return err
	}

	return adaptor.HTTPHandler(templ.Handler(Main(*allUsers, allImages)))(c)
}

func tDelete(c *fiber.Ctx) error {
	return nil
}

func main() {

	var (
		app = fiber.New(fiber.Config{Views: nil})
	)

	app.Static("/", "./public")

	app.Get("/", tGet)
	app.Post("/", tPost)
	app.Delete("/dummyDelete", tDelete)

	if err := app.Listen(":8000"); err != nil {
		log.Fatal(err)
	}

}
