package routes

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sabilimaulana/sebelbucks-api-gateway/pkg/product/pb"
)

func DetailProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)

	res, err := c.DetailProduct(ctx, &pb.DetailProductRequest{ProductId: id})
	if err != nil {
		ctx.AbortWithError(int(res.Status), nil)
		return
	}

	ctx.JSON(int(res.Status), res)
}
