package middleware

import (
	"example/web-service-gin/response"
	"example/web-service-gin/service"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			response := response.BuildErrorResponse("Can not read Authorization", "Bad Request")
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		}
		token, err := jwtService.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("AuthorizeJWT claims[username]: ", claims["username"])
			log.Println("AuthorizeJWT claims[issuer]: ", claims["issuer"])
		} else {
			log.Println(err)
			response := response.BuildErrorResponse("Token is not Valid", err.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
