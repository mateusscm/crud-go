package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mateusscm/crud-go/src/config/rest_err"
	"github.com/mateusscm/crud-go/src/controller/model/request"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("Invalid JSON body, error=%s", err.Error()))
		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)

}
