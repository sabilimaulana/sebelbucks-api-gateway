package product

import (
	"github.com/gin-gonic/gin"
	"github.com/sabilimaulana/sebelbucks-api-gateway/pkg/auth"
	"github.com/sabilimaulana/sebelbucks-api-gateway/pkg/config"
	"github.com/sabilimaulana/sebelbucks-api-gateway/pkg/product/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/product")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateProduct)
}

func (svc *ServiceClient) CreateProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, svc.Client)
}
