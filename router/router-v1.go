package router

import (
	"github.com/gin-gonic/gin"
)



func InitiateRouterV1(r *gin.RouterGroup) *gin.RouterGroup {

	GetTaskRoutes(r)

	return r

}
