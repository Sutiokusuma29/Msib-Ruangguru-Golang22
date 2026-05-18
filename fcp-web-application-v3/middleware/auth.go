package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if isPublicRoute(ctx.Request.URL.Path) {
			ctx.Next()
			return
		}

		cookie, err := ctx.Cookie("session_token")
		if err != nil {
			handleUnauthorized(ctx, err)
			return
		}

		tokenClaims := model.Claims{}
		token, err := jwt.ParseWithClaims(cookie, &tokenClaims, func(token *jwt.Token) (interface{}, error) {
			return model.JwtKey, nil
		})
		if err != nil || !token.Valid {
			handleUnauthorized(ctx, err)
			return
		}

		ctx.Set("email", tokenClaims.Email)
		ctx.Next()
	}
}

func isPublicRoute(path string) bool {
	publicRoutes := []string{"/user/register", "/user/login"}
	for _, route := range publicRoutes {
		if path == route {
			return true
		}
	}
	return false
}


func handleUnauthorized(ctx *gin.Context, err error) {
	if err == http.ErrNoCookie {
		if ctx.Request.Header.Get("Content-Type") == "application/json" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
		} else {
			ctx.Redirect(http.StatusSeeOther, "/login")
		}
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
	}
}
