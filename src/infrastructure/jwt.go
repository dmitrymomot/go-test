package infrastructure

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	uuid "github.com/satori/go.uuid"
)

// Predefined constants
const (
	JWTContextKey = "JWT"
	JWTAuthScheme = "Bearer"
	JWTUserID     = "user_id"
	JWTScope      = "scope"
	JWTExpiresAt  = "exp"
	JWTID         = "jti"
)

// JWTMiddleware is the default JWT middleware with custom config
func JWTMiddleware(key []byte) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: middleware.AlgorithmHS256,
		SigningKey:    key,
		ContextKey:    JWTContextKey,
		TokenLookup:   "header:" + echo.HeaderAuthorization,
		AuthScheme:    JWTAuthScheme,
		Claims:        jwt.MapClaims{},
	})
}

// NewJWTHandler is a factory function,
// returns a new instance of JWTHandler structure
func NewJWTHandler(sk string, lt int64) *JWTHandler {
	return &JWTHandler{sk, lt}
}

// JWTHandler structure
type JWTHandler struct {
	signingKey string
	lifeTime   int64
}

// Make generates JWT
func (h *JWTHandler) Make(userID int, scope string) (string, error) {
	exp := time.Now().Add(time.Duration(h.lifeTime) * time.Second)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims[JWTUserID] = userID
	claims[JWTScope] = scope
	claims[JWTExpiresAt] = exp.Unix()
	claims[JWTID] = uuid.NewV1().String()

	ts, err := token.SignedString([]byte(h.signingKey))
	if err != nil {
		return "", err
	}

	return ts, nil
}
