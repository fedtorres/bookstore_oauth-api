package http

import (
	"github.com/fedtorres/bookstore_oauth-api/src/domain/access_token"
	access_token2 "github.com/fedtorres/bookstore_oauth-api/src/services/access_token"
	"github.com/fedtorres/bookstore_utils-go/rest_errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccessTokenHandler interface {
	GetById(c *gin.Context)
	Create(c *gin.Context)
}

type accessTokenHandler struct {
	service access_token2.Service
}

func NewHandler(service access_token2.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetById(c *gin.Context) {
	accessToken, err := handler.service.GetById(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (handler *accessTokenHandler) Create(c *gin.Context) {
	var at access_token.AccessToken
	if err := c.ShouldBindJSON(&at); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	if err := handler.service.Create(at); err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, at)
}
