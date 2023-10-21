package helper

import (
	"fmt"
	"time"

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
