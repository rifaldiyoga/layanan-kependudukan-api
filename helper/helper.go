package helper

import (
	"context"
	"fmt"
	"log"

	"strconv"
	"strings"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

const PENDUDUK = "PENDUDUK"
const RT = "RT"
const RW = "RW"
const KELURAHAN = "KELURAHAN"
const ADMIN = "ADMIN"

type response struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) response {

	res := response{
		Message: message,
		Code:    code,
		Status:  status,
		Data:    data,
	}

	return res
}

func FormatValidationError(err error) []string {
	var errors []string

	_, ok := err.(validator.ValidationErrors)
	if ok {
		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, e.Error())
		}
	} else {
		errors = append(errors, "Format salah!")
	}

	return errors
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT, UPDATE, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func FormatDateToString(time time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}

func FormatStringToDate(date string) time.Time {
	parse, err := time.Parse("2006-01-02", date)
	if err != nil {
		return time.Now()
	}
	return parse
}

func GetType(Type string) string {
	myMap := map[string]string{
		"LU":  "Layanan Umum",
		"LP":  "Layanan Pindah",
		"LN":  "Layanan Nikah",
		"LKK": "Layanan Kematian & Kelahiran",
		"LT":  "Layanan Pertanahan",
	}
	return myMap[Type]
}

func GenerateNIK(date time.Time, count int) string {
	counter := fmt.Sprintf("%04d", count)
	return "357901" + date.Format("020106") + counter
}

func GenerateNoKK(count int) string {
	counter := fmt.Sprintf("%04d", count)
	return "357901" + time.Now().Format("020106") + counter
}

func GenerateKodeSurat(codeLayanan string, lastCode string) string {
	count := 0
	if lastCode != "" {
		split := strings.Split(lastCode, "/")[1]
		num, err := strconv.Atoi(split)

		if err != nil {
			fmt.Println("Conversion error:", err)
			return ""
		}
		count = num

	}
	count = count + 1
	counter := fmt.Sprintf("%03d", count)
	return codeLayanan + "/" + counter + "/422.310.2/" + time.Now().Format("2006")
}

func FormatFileName(filename string) string {
	// Convert the string to lowercase
	processed := strings.ToLower(filename)

	// Replace spaces with underscores
	processed = strings.ReplaceAll(processed, " ", "_")

	return processed
}

func SendNotification(app *firebase.App, token string, title string, body string) {
	if token != "" {
		ctx := context.Background()
		client, err := app.Messaging(ctx)
		if err != nil {
			log.Fatalf("error getting Messaging client: %v\n", err)
		}

		registrationToken := token

		message := &messaging.Message{
			Notification: &messaging.Notification{
				Title: title,
				Body:  body,
			},
			Token: registrationToken,
		}

		response, err := client.Send(ctx, message)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Successfully sent message:", response)
	}

}

func CapitalizeEachWord(input string) string {
	words := strings.Fields(input)

	var capitalizedWords []string
	for _, word := range words {
		capitalizedWord := strings.Title(word)
		capitalizedWords = append(capitalizedWords, capitalizedWord)
	}
	return strings.Join(capitalizedWords, " ")
}
