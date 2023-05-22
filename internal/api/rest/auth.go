package rest

import (
	"github.com/Onelvay/HL-architecture/internal/dto"
	"github.com/Onelvay/HL-architecture/internal/service"
	"github.com/gin-gonic/gin"
)

type Authorization struct {
	authService service.AuthService
}

func NewAuthorizationHandler(authService service.AuthService) *Authorization {
	return &Authorization{authService: authService}
}

func authRoutes(router *gin.Engine, a *Authorization) {
	group := router.Group("/auth")
	{
		group.GET("/sign-in", a.SignIn)
		group.POST("/sign-up", a.SignUp)
	}

}

func (a *Authorization) SignUp(c *gin.Context) {
	var req dto.SignUpRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	res, err := a.authService.SignUp(c, req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, res)
}

func (a *Authorization) SignIn(c *gin.Context) {
	var req dto.SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	res, err := a.authService.SignIn(c, req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, res)
}