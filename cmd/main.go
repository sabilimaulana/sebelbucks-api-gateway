package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sabilimaulana/sebelbucks-api-gateway/pkg/auth"
	"github.com/sabilimaulana/sebelbucks-api-gateway/pkg/config"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	auth.RegisterRoutes(r, &c)

	r.Run(c.Port)
}
