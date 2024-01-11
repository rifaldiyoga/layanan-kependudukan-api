package helper

import (
	"context"
	"fmt"
	"log"
	"net/smtp"

	"strconv"
	"strings"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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

var num = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

var numInv = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

var maxTable = []int{
	1000,
	900,
	500,
	400,
	100,
	90,
	50,
	40,
	10,
	9,
	5,
	4,
	1,
}

func highestDecimal(n int) int {
	for _, v := range maxTable {
		if v <= n {
			return v
		}
	}
	return 1
}

// ToRoman is to convert decimal number to roman numeral
func ToRoman(n int) string {
	out := ""
	for n > 0 {
		v := highestDecimal(n)
		out += numInv[v]
		n -= v
	}
	return out
}

func GenerateKodeSurat(codeLayanan string, codeKelurahan string, lastCode string) string {
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
	_, month, _ := time.Now().Date()
	num := int(month)

	count = count + 1
	counter := fmt.Sprintf("%03d", count)
	return codeLayanan + "/" + counter + "/" + codeKelurahan + "/" + ToRoman(num) + "/" + time.Now().Format("2006")
}

func FormatFileName(filename string) string {
	// Convert the string to lowercase
	processed := strings.ToLower(filename)

	// Replace spaces with underscores
	processed = strings.ReplaceAll(processed, " ", "_")

	return processed
}

func SendNotification(app *firebase.App, email string, token string, title string, body string) {
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
		sendMail(email, title, body)
	}

}

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "Kelurahan Ngaglik <satu.indonesia001@gmail.com>"
const CONFIG_AUTH_EMAIL = "satu.indonesia001@gmail.com"
const CONFIG_AUTH_PASSWORD = "sembarang007"

func sendMail(to string, subject, message string) error {
	body := "From: " + CONFIG_SENDER_NAME + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD, CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

	err := smtp.SendMail(smtpAddr, auth, CONFIG_AUTH_EMAIL, []string{to}, []byte(body))
	if err != nil {
		return err
	}

	return nil
}

func CapitalizeEachWord(input string) string {
	words := strings.Fields(input)

	var capitalizedWords []string
	for _, word := range words {
		capitalizedWord := cases.Title(language.Und).String(word)
		capitalizedWords = append(capitalizedWords, capitalizedWord)
	}
	return strings.Join(capitalizedWords, " ")
}
