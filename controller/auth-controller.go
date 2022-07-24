package controller

import (
	"example/web-service-gin/entity"
	"example/web-service-gin/request"
	"example/web-service-gin/response"
	"example/web-service-gin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

// Login implements AuthController
func (service *authController) Login(ctx *gin.Context) {
	var loginRequest request.LoginRequest
	errRequest := ctx.ShouldBind(&loginRequest)
	if errRequest != nil {
		response := response.BuildErrorResponse("Register Failure", errRequest.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := service.authService.VerifyCredential(loginRequest.Email, loginRequest.Password)
	if v, ok := authResult.(entity.User); ok {
		generatedToken := service.jwtService.GenerateToken(v.Username)
		response := response.BuildCommonReponse(true, "Generate Token Success", generatedToken)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := response.BuildErrorResponse("Login Failure", "Invalid Credential")
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

// Register implements AuthController
func (service *authController) Register(ctx *gin.Context) {
	var user request.UserRequest
	errRequest := ctx.ShouldBind(&user)
	if errRequest != nil {
		response := response.BuildErrorResponse("Register Failure", errRequest.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !service.authService.IsDupplicate(user.Email) {
		response := response.BuildErrorResponse("Email exists", "Bad Request")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
	} else {
		authResult := service.authService.CreateUser(user)
		response := response.BuildCommonReponse(true, "Register Success", authResult)
		ctx.JSON(http.StatusCreated, response)
	}
}

func NewAuthController(authService service.AuthService,
	jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}
