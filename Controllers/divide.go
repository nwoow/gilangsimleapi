// controllers
package controllers

import (
	models "TSLC/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Divide...
func Divide(c *gin.Context) {
	var requests models.Requests

	//Check the error if valid inputs are provided
	if err := c.BindJSON(&requests); err != nil {
		c.JSON(http.StatusBadRequest, "Please provide the valid input")
	}
	//Divide the inputs a/b
	divide_result_1 := requests.A / requests.B

	//Divide the inputs b/a
	divide_result_2 := requests.B / requests.A
	c.JSON(http.StatusOK, gin.H{
		"Result 1": divide_result_1,
		"Result 2": divide_result_2,
	})

}
