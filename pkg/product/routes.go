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

	// Public routes
	routes := r.Group("/products")
	routes.GET("/", svc.ListProduct)

	// Must authorized routes
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateProduct)
	routes.DELETE("/:id", svc.DeleteProduct)

}

func (svc *ServiceClient) CreateProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, svc.Client)
}

func (svc *ServiceClient) ListProduct(ctx *gin.Context) {
	routes.ListProduct(ctx, svc.Client)
}

func (svc *ServiceClient) DeleteProduct(ctx *gin.Context) {
	routes.DeleteProduct(ctx, svc.Client)
}
