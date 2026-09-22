package auth
import (
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
)
const claimsKey = "claims"
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
func GetClaims(c *gin.Context) *Claims {
	raw, exists := c.Get(claimsKey)
	if !exists {
		return nil
	}
	claims, _ := raw.(*Claims)
	return claims
}
