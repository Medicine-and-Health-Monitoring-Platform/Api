package middleware

import (
	"Api_Gateway/config"
	"fmt"
	"net/http"

	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type casbinPermission struct {
	enforcer *casbin.Enforcer
}

func Check(c *gin.Context) {
	accessToken := c.GetHeader("Authorization")

	if accessToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Authorization header is required",
		})
		return
	}

	token, err := jwt.ParseWithClaims(accessToken, jwt.MapClaims{}, func(
		t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v",
				t.Header["alg"])
		}
		return []byte(config.Load().ACCESS_TOKEN), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Token could not be parsed",
		})
		return
	}

	if !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token provided",
		})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid claims' type",
		})
		return
	}

	_, roleOk := claims["role"].(string)
	userID, idOk := claims["user_id"].(string)
	if !roleOk || !idOk {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token claims",
		})
		return
	}

	_, err = uuid.Parse(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid user id",
		})
		return
	}

	c.Set("user_id", userID)

	//e, err := rbac.Policy()
	//if err != nil {
	//	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
	//		"error": "Policy could not be loaded",
	//	})
	//	return
	//}
	//
	//if ok, err := e.Enforce(userRole, c.Request.URL.Path, c.Request.Method); !ok || err != nil {
	//	msg := fmt.Sprintf("Access denied: %s cannot %s %s",
	//		userRole, c.Request.Method, c.Request.URL.Path,
	//	)
	//	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
	//		"error": msg,
	//	})
	//	return
	//}

	c.Next()
}

// func (casb *casbinPermission) GetRole(c *gin.Context) (string, int) {
// 	token := c.GetHeader("Authorization")
// 	if token == "" {
// 		return "unauthorized", http.StatusUnauthorized
// 	}
// 	_, role, _, err := auth.GetUserInfoFromAccessToken(token)
// 	if err != nil {
// 		return "error while reding role", 500
// 	}

// 	return role, 0
// }

// func (casb *casbinPermission) CheckPermission(c *gin.Context) (bool, error) {

// 	act := c.Request.Method
// 	sub, status := casb.GetRole(c)
// 	if status != 0 {
// 		return false, errors.New("error in get role")
// 	}
// 	obj := c.FullPath()

// 	ok := casb.enforcer.Enforce(sub, obj, act)
// 	// if err != nil {
// 	// 	c.JSON(http.StatusInternalServerError, gin.H{
// 	// 		"Error": "Internal server error",
// 	// 	})
// 	// 	c.Abort()
// 	// 	return false, err
// 	// }
// 	return ok, nil
// }

// func CheckPermissionMiddleware(enf *casbin.Enforcer) gin.HandlerFunc {
// 	casbHandler := &casbinPermission{
// 		enforcer: enf,
// 	}

// 	return func(c *gin.Context) {
// 		result, err := casbHandler.CheckPermission(c)

// 		if err != nil {
// 			c.AbortWithError(500, err)
// 		}
// 		if !result {
// 			c.AbortWithStatusJSON(401, gin.H{
// 				"message": "Forbidden",
// 			})
// 		}

// 		c.Next()
// 	}
// }
