package service

import (
	"context"
	"time"

	"demo/internal/dao"
	"demo/internal/model/entity"

	jwt "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/frame/g"
)

// Global variable to store the JWT middleware instance
var authService *jwt.GfJWTMiddleware

// Function to get the JWT middleware instance for use elsewhere
func JWTAuth() *jwt.GfJWTMiddleware {
	return authService
}

// Initialize the JWT middleware with custom handler functions
func init() {
	auth := jwt.New(&jwt.GfJWTMiddleware{
		Realm:           "auth zone",
		Key:             []byte("QH7GMFykSVxPrLkc"),
		Timeout:         time.Hour * 1,
		MaxRefresh:      time.Hour * 2,
		IdentityKey:     "uid",
		TokenLookup:     "header: Authorization, query: token",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
		Authenticator:   Authenticator,   // Actual authentication handler
		Unauthorized:    Unauthorized,    // Handler for failed authentication
		PayloadFunc:     PayloadFunc,     // Build payload for JWT token
		IdentityHandler: IdentityHandler, // Extract identity from token
	})
	authService = auth
}

// PayloadFunc builds the JWT payload from user data.
// Example: If user info is {"uid": 1, "username": "admin"}, these will be stored in the token.
func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	// data is a map[string]interface{} containing user info
	if params, ok := data.(map[string]interface{}); ok {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

// IdentityHandler extracts the identity (uid) from the JWT token.
// Example: If token contains {"uid": 1}, this returns 1.
func IdentityHandler(ctx context.Context) interface{} {
	claims := jwt.ExtractClaims(ctx)
	return claims[authService.IdentityKey]
}

// Unauthorized handles failed authentication (e.g., wrong user/pass, expired token, etc.)
// Example: If authentication fails, returns {"code": code, "message": message} and stops request.
func Unauthorized(ctx context.Context, code int, message string) {
	r := g.RequestFromCtx(ctx)
	r.Response.WriteJson(g.Map{
		"code":    code,
		"message": message,
	})
	r.ExitAll()
}

// Authenticator checks user/pass from the request body and verifies with the database.
// Example: If username/password match a user in DB, returns user info for token payload.
func Authenticator(ctx context.Context) (interface{}, error) {
	r := g.RequestFromCtx(ctx)
	var login struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	// Parse the JSON body into the login struct
	if err := r.Parse(&login); err != nil {
		return nil, jwt.ErrFailedAuthentication
	}

	// Query user from database
	var user entity.User
	err := dao.User.Ctx(ctx).
		Where("username = ? AND password = ?", login.Username, login.Password).
		Scan(&user)
	if err != nil || user.Id == 0 {
		// If not found or error, return authentication error
		return nil, jwt.ErrFailedAuthentication
	}

	// Return user info to be stored in the token
	return g.Map{
		"uid": user.Id,
		// "username": user.Username,
	}, nil
}

/*
--- Example usage ---
1. Login request:
   POST /login
   Body: { "username": "admin", "password": "123456" }

2. If authentication succeeds, response contains JWT token.
3. Use token for protected routes:
   GET /api/protected
   Header: Authorization: Bearer <token>
----------------------
*/
