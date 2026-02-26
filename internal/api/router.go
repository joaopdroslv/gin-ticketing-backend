package api

import "github.com/gin-gonic/gin"

func Register(r *gin.Engine, dependencies Dependencies) {

	apiGroup := r.Group("/api")

	RegisterV1(apiGroup, dependencies)
}
