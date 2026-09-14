package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const claimsKey = "claims"

// Authenticate is a Gin middleware that validates the Bearer JWT in the
// Authorization header and injects the Claims into the context.
func Authenticate(jwtMgr *JWTManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenStr string
		header := c.GetHeader("Authorization")
		if header != "" && strings.HasPrefix(header, "Bearer ") {
			tokenStr = strings.TrimPrefix(header, "Bearer ")
		} else {
			tokenStr, _ = c.Cookie("token")
		}

		if tokenStr == "" {
			if c.Request.Method == http.MethodGet && !strings.HasPrefix(c.Request.URL.Path, "/api") {
				c.Redirect(http.StatusFound, "/sign-in")
				c.Abort()
				return
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing or malformed authorization header"})
			return
		}

		claims, err := jwtMgr.Validate(tokenStr)
		if err != nil {
			// For HTMX/Browser we could redirect to login instead of JSON error
			// Let's do it if it's a GET request and not /api
			if c.Request.Method == http.MethodGet && !strings.HasPrefix(c.Request.URL.Path, "/api") {
				c.Redirect(http.StatusFound, "/sign-in")
				c.Abort()
				return
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}

		c.Set(claimsKey, claims)
		c.Next()
	}
}

// RequireRole returns a Gin middleware that enforces the caller has one of the
// allowed roles. Must be used after Authenticate.
func RequireRole(roles ...Role) gin.HandlerFunc {
	allowed := make(map[Role]struct{}, len(roles))
	for _, r := range roles {
		allowed[r] = struct{}{}
	}

	return func(c *gin.Context) {
		raw, exists := c.Get(claimsKey)
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "not authenticated"})
			return
		}

		claims, ok := raw.(*Claims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "invalid claims type"})
			return
		}

		if _, ok := allowed[claims.Role]; !ok {
			// If HTMX/UI, maybe render unauthorized
			if !strings.HasPrefix(c.Request.URL.Path, "/api") {
				c.String(http.StatusForbidden, "Forbidden: Insufficient Permissions")
				c.Abort()
				return
			}
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "insufficient permissions"})
			return
		}

		c.Next()
	}
}

// GetClaims extracts Claims from a gin.Context (set by Authenticate middleware).
// Returns nil if not present.
func GetClaims(c *gin.Context) *Claims {
	raw, exists := c.Get(claimsKey)
	if !exists {
		return nil
	}
	claims, _ := raw.(*Claims)
	return claims
}
