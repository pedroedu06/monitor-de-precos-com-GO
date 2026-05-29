package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pedroedu06/monitor-de-precos-com-GO/internal/controller/model/request"
	"github.com/pedroedu06/monitor-de-precos-com-GO/internal/model"
	"github.com/pedroedu06/monitor-de-precos-com-GO/internal/repository"
	"github.com/pedroedu06/monitor-de-precos-com-GO/rest_err"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"menssage": "pong",
	})
}

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := resterr.NewBadRequestErr(
			fmt.Sprintf("There are some incorret fields, error=%s\n", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}
	user := model.UserDomain{
		Telefone: userRequest.Telefone,
	}

	if err := user.Validate(); err != nil {
		restErr := resterr.NewBadRequestErr(err.Error())
		c.JSON(restErr.Code, restErr)
		return
	}

	repo := repository.NewUserReposity()
	createdUser, err := repo.Create(user)
	if err != nil {
		restErr := resterr.NewBadRequestErr(
			fmt.Sprintf("Erro ao criar usuario, error=%s\n", err.Error()))
		c.JSON(restErr.Code, restErr)
		return
	}        

	c.JSON(201, createdUser)
}

func LoginUser(c *gin.Context) {

}

func CreateProduct(c *gin.Context) {

}

func ListProduct(c *gin.Context) {

}

func DeleteProduct(c *gin.Context) {

}

func PauseProduct(c *gin.Context) {

}

func HistoricPrices(c *gin.Context) {

}

func ListNotifications(c *gin.Context) {

}
