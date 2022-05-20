// controllers/books.go

package controllers

import (
	models "DT/Models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Divide
func Divide(c *gin.Context) {
	fmt.Println("error")
	var requests models.Requests
	c.BindJSON(&requests)
	divide_result_1 := requests.A / requests.B
	divide_result_2 := requests.B / requests.A
	c.JSON(http.StatusOK, gin.H{
		"R1": divide_result_1,
		"R2": divide_result_2, // cast it to string before showing
	})
}
