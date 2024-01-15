package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"github.com/quangvu30/order-go/config"
	"github.com/quangvu30/order-go/types"
)

type UserPayload struct {
	Email string `json:"email"`
}

type AuthMiddleware struct {
	Config *config.Config
}

func NewAuthMiddleware(c *config.Config) *AuthMiddleware {
	return &AuthMiddleware{
		Config: c,
	}
}

func (a *AuthMiddleware) User() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtToken, err := extractBearerToken(c.GetHeader("Authorization"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, types.UnsignedResponse{
				Msg: err.Error(),
			})
			return
		}

		token, err := parseToken(jwtToken, a.Config.JwtSecret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, types.UnsignedResponse{
				Msg: "bad jwt token",
			})
			return
		}

		claims, OK := token.Claims.(jwt.MapClaims)
		if !OK {
			c.AbortWithStatusJSON(http.StatusInternalServerError, types.UnsignedResponse{
				Msg: "unable to parse claims",
			})
			return
		}
		c.Set("user", UserPayload{
			Email: claims["email"].(string),
		})
		c.Next()
	}
}

func parseToken(jwtToken string, secKey string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, OK := token.Method.(*jwt.SigningMethodHMAC); !OK {
			return nil, errors.New("bad signed method received")
		}
		return []byte(secKey), nil
	})

	if err != nil {
		return nil, errors.New("bad jwt token")
	}

	return token, nil
}

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("bad header value given")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}
