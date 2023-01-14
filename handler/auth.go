package handler

import (
	"filestore/common"
	"filestore/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HTTPInterceptor : http请求拦截器
func HTTPInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Request.FormValue("username")
		token := c.Request.FormValue("token")

		//验证登录token是否有效
		if len(username) < 3 || !IsTokenValid(token) {
			// w.WriteHeader(http.StatusForbidden)
			// token校验失败则跳转到登录页面
			c.Abort()
			resp := util.NewRespMsg(
				int(common.StatusTokenInvalid),
				"token无效",
				nil,
			)
			c.JSON(http.StatusOK, resp)
			return
		}
		c.Next()
	}
}

// gin jwt 认证中间件
// func AuthRequired() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		tokenString := strings.TrimPrefix(ctx.GetHeader("Authorization"), "Bearer ")
// 		token, err := jwt.ParseWithClaims(tokenString, &customClaims{}, func(t *jwt.Token) (interface{}, error) { return jwtKey, nil })
// 		if err != nil {
// 			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"code": -1, "msg": fmt.Sprintf("access token parse error: %v", err)})
// 			return
// 		}
// 		if claims, ok := token.Claims.(*customClaims); ok && token.Valid {
// 			if !claims.VerifyExpiresAt(time.Now(), false) {
// 				ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"code": -1, "msg": "access token expired"})
// 				return
// 			}
// 			// 即将超过过期时间，则添加一个http header `new-token` 给前端更新
// 			if t := claims.ExpiresAt.Time.Add(-time.Minute * TOKEN_MAX_REMAINING_MINUTE); t.Before(time.Now()) {
// 				claims := customClaims{
// 					Username: claims.Username,
// 					IsAdmin:  claims.Username == "admin",
// 					RegisteredClaims: jwt.RegisteredClaims{
// 						ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(TOKEN_MAX_EXPIRE_HOUR * time.Hour)},
// 					},
// 				}
// 				token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 				tokenString, _ := token.SignedString(jwtKey)
// 				ctx.Header("new-token", tokenString)
// 			}
// 			ctx.Set("claims", claims)
// 		} else {
// 			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"code": -1, "msg": fmt.Sprintf("Claims parse error: %v", err)})
// 			return
// 		}
// 		ctx.Next()
// 	}
// }
