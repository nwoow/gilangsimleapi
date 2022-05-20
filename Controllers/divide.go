// controllers/books.go

package controllers

import (
	models "DT/Models"

	"github.com/gin-gonic/gin"
)

//Divide
func Divide(c *gin.Context) {
	var person models.Requests
	c.BindJSON(&person)
	divide_result_1 := person.A / person.B
	divide_result_2 := person.B / person.A
	c.JSON(200, gin.H{
		"R1": divide_result_1,
		"R2": divide_result_2, // cast it to string before showing
	})
}
