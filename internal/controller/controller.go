package controller

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pedroedu06/monitor-de-precos-com-GO/internal/controller/model/request"
	"github.com/pedroedu06/monitor-de-precos-com-GO/internal/model"
	"github.com/pedroedu06/monitor-de-precos-com-GO/internal/pkg/jwt"
	"github.com/pedroedu06/monitor-de-precos-com-GO/internal/repository"
	"github.com/pedroedu06/monitor-de-precos-com-GO/rest_err"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"menssage": "pong",
	})
}


func AuthUser(c *gin.Context) {
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

    // novo — busca antes de criar
    existingUser, err := repo.FindByTelefone(user.Telefone)
    if err == sql.ErrNoRows {
        existingUser, err = repo.Create(user)
        if err != nil {
            restErr := resterr.NewBadRequestErr(
                fmt.Sprintf("Error creating user, error=%s\n", err.Error()))
            c.JSON(restErr.Code, restErr)
            return
        }
    } else if err != nil {
        restErr := resterr.NewBadRequestErr("Error fetching user")
        c.JSON(restErr.Code, restErr)
        return
    }

    token, err := jwt.GenerateToken(existingUser.ID)
    if err != nil {
        restErr := resterr.NewInternalServerErr("Error generating token")
        c.JSON(restErr.Code, restErr)
        return
    }

    c.JSON(200, gin.H{"token": token})
}


func CreateProduct(c *gin.Context) {
	var ProductRequest request.ProductRequest
	if err := c.ShouldBindJSON(&ProductRequest); err != nil {
		restErr := resterr.NewBadRequestErr(
            fmt.Sprintf("There are some incorret fields, error=%s\n", err.Error()))
        c.JSON(restErr.Code, restErr)
        return
	}

	userid := c.GetString("UserID")

	product := model.ProdutoDomain{
		UserID: userid,
		URL: ProductRequest.URL,
		TargetPrice: ProductRequest.PrecoAlvo,
	}

	if err := product.Validate(); err != nil {
		restErr := resterr.NewBadRequestErr(err.Error())
		c.JSON(restErr.Code, restErr)
		return
	}

	repo := repository.NewProdutoRepository()
	created, err := repo.Create(product)
	if err != nil {
		restErr := resterr.NewInternalServerErr(
			fmt.Sprintf("Error creating product, error=%s\n", err.Error()))
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(201, created)
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
