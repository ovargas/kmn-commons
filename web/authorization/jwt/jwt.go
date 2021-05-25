package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	config2 "github.com/ovargasmahisoft/kmn-commons/config"
	authorization2 "github.com/ovargasmahisoft/kmn-commons/web/authorization"
	"strings"
)

type Handler struct {
	pemPublicKey string
}

//AuthenticationHandler jwt token authentication handler
func (h Handler) AuthenticationHandler(ctx *gin.Context) {
	stringToken := ctx.GetHeader("Authorization")
	bearer := strings.Split(stringToken, " ")

	var principal = authorization2.Anonymous()

	if len(bearer) == 2 {
		if token, _ := jwt.Parse(bearer[1], h.verifyToken); token != nil {
			claims, _ := token.Claims.(jwt.MapClaims)

			var roles []string
			if iRoles, ok := claims["authorities"].([]interface{}); ok {
				for _, r := range iRoles {
					roles = append(roles, r.(string))
				}
			}

			userName, _ := claims["user_name"].(string)

			domain, _ := claims["domain"].(string)
			clientID, _ := claims["client_id"].(string)

			if userName == "" || domain == "" {
				principal = authorization2.NewPrincipal(roles,
					authorization2.NewClientCredentialsIdentity(clientID))
			} else {
				principal = authorization2.NewPrincipal(roles,
					authorization2.NewUserIdentity(userName, domain, clientID))
			}

		}
	}

	ctx.Set(authorization2.PrincipalContextName, principal)
	ctx.Next()
}

//verifyToken validate token signature
func (h Handler) verifyToken(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	if pem, err := jwt.ParseRSAPublicKeyFromPEM([]byte(h.pemPublicKey)); err != nil {
		return nil, err
	} else {
		return pem, nil
	}
}

func New() Handler {
	var key *string
	err := config2.Config().UnmarshalKey("security.oauth2.resource.jwt.keyValue", &key)
	if err != nil {
		panic(fmt.Errorf("error loading configuration [security.oauth2.resource.jwt.keyValue]: %v", err))
	}

	if key == nil {
		panic(fmt.Errorf("missing configuration [security.oauth2.resource.jwt.keyValue]"))
	}

	return NewWithPublicKey(*key)
}

func NewWithPublicKey(publicKey string) Handler {
	return Handler{
		pemPublicKey: publicKey,
	}
}
