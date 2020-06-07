package app

import (
	"github.com/fedtorres/bookstore_oauth-api/src/clients/cassandra"
	"github.com/fedtorres/bookstore_oauth-api/src/http"
	"github.com/fedtorres/bookstore_oauth-api/src/repository/db"
	access_token2 "github.com/fedtorres/bookstore_oauth-api/src/services/access_token"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session := cassandra.GetSession()
	session.Close()

	atHandler := http.NewHandler(access_token2.NewService(db.NewRepository()))
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8080")
}
