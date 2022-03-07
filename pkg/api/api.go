package api

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"simple-api-gateway/config"
	customerTrp "simple-api-gateway/pkg/api/customer"
)

func Start() {
	router := gin.Default()

	group := router.Group("/v1")

	customerTrp.NewHTTP(customerTrp.New(), group)

	router.Run(config.Env.ServerAddress)
}
