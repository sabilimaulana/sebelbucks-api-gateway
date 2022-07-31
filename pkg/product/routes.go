package product

import (
	"github.com/gin-gonic/gin"
	"github.com/sabilimaulana/sebelbucks-api-gateway/pkg/auth"
	"github.com/sabilimaulana/sebelbucks-api-gateway/pkg/config"
	"github.com/sabilimaulana/sebelbucks-api-gateway/pkg/product/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	// a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	// Public routes
	// Products
	productRoutes := r.Group("/products")
	productRoutes.GET("/", svc.ListProduct)
	productRoutes.GET("/:id", svc.DetailProduct)

	// Variants
	variantRoutes := r.Group("/variants")
	variantRoutes.GET("/", svc.ListVariant)

	// Must authorized routes
	// r.Use(a.AuthRequired)
}

func (svc *ServiceClient) ListProduct(ctx *gin.Context) {
	routes.ListProduct(ctx, svc.Client)
}

func (svc *ServiceClient) DetailProduct(ctx *gin.Context) {
	routes.DetailProduct(ctx, svc.Client)
}

func (svc *ServiceClient) ListVariant(ctx *gin.Context) {
	routes.ListVariant(ctx, svc.Client)
}
