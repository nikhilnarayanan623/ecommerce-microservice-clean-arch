package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/find", handler)

	r.Run(":8000")
}

func handler(ctx *gin.Context) {

	url := ctx.Query("url")

	res, err := http.Get(url)

	// http.Post()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Println("status: ", res.StatusCode)

	body := Data{}

	jsonData, err := io.ReadAll(res.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Println()

	err = json.Unmarshal(jsonData, &body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, body)

}

type Data struct {
	Error   interface{} `json:"error"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}
