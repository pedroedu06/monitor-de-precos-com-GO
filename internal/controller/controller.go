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
    userID := c.GetString("UserID")

    repo := repository.NewProdutoRepository()
    produtos, err := repo.ListByUserID(userID)
    if err != nil {
        restErr := resterr.NewInternalServerErr(
            fmt.Sprintf("error listing products, error=%s", err.Error()))
        c.JSON(restErr.Code, restErr)
        return
    }
    c.JSON(200, produtos)
}

func DeleteProduct(c *gin.Context) {
    id := c.Param("id")
    userID := c.GetString("UserID")

    repo := repository.NewProdutoRepository()
    rowsAffected, err := repo.DeleteProductById(id, userID)
    if err != nil {
        restErr := resterr.NewInternalServerErr(
            fmt.Sprintf("error deleting product, error=%s", err.Error()))
        c.JSON(restErr.Code, restErr)
        return
    }
    if rowsAffected == 0 {
        restErr := resterr.NewNotFoundErr("Product not found")
        c.JSON(restErr.Code, restErr)
        return
    }

    c.JSON(200, gin.H{"mensage": "sucess"})
}

func PauseProduct(c *gin.Context) {
    setProductActive(c, false)
}

func ResumeProduct(c *gin.Context) {
    setProductActive(c, true)
}

func setProductActive(c *gin.Context, active bool) {
    id := c.Param("id")
    userID := c.GetString("UserID")

    repo := repository.NewProdutoRepository()
    product, err := repo.SetActiveById(id, userID, active)
    if err == sql.ErrNoRows {
        restErr := resterr.NewNotFoundErr("Product not found")
        c.JSON(restErr.Code, restErr)
        return
    }
    if err != nil {
        restErr := resterr.NewInternalServerErr(
            fmt.Sprintf("error updating product state, error=%s", err.Error()))
        c.JSON(restErr.Code, restErr)
        return
    }

    c.JSON(200, product)
}

func HistoricPrices(c *gin.Context) {
    productID := c.Param("id")
    userID := c.GetString("UserID")

    repo := repository.NewPriceHistoryRepository()
    history, err := repo.ListByProductID(productID, userID)
    if err != nil {
        restErr := resterr.NewInternalServerErr(
            fmt.Sprintf("error fetching price history, error=%s", err.Error()))
        c.JSON(restErr.Code, restErr)
        return
    }

    // Normalize nil to empty slice so the client always gets [] instead of null.
    if history == nil {
        history = []model.PriceHistoryDomain{}
    }

    c.JSON(200, history)
}

func ListNotifications(c *gin.Context) {
    userID := c.GetString("UserID")

    repo := repository.NewNotificationRepository()
    notifications, err := repo.ListByUserID(userID)
    if err != nil {
        restErr := resterr.NewInternalServerErr(
            fmt.Sprintf("error fetching notifications, error=%s", err.Error()))
        c.JSON(restErr.Code, restErr)
        return
    }

    // Normalize nil to empty slice so the client always gets [] instead of null.
    if notifications == nil {
        notifications = []model.NotificationDomain{}
    }

    c.JSON(200, notifications)
}
