package controllers

import (
	"project-skbackend/configs"
	"project-skbackend/internal/middlewares"
	"project-skbackend/packages/consttypes"

	"github.com/gin-gonic/gin"
)

type (
	manageroutes struct {
		cfg *configs.Config
	}
)

func newManageRoutes(
	rg *gin.RouterGroup,
	cfg *configs.Config,
) {
	// * change this into the correct declaration
	_ = &manageroutes{
		cfg: cfg,
	}

	gmanage := rg.Group("manages")
	gmanage.Use(middlewares.JWTAuthMiddleware(
		cfg,
		consttypes.UR_ADMIN,
	))
	{

	}
}
