package util

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"filestore/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)



func GenToken2(username, password string) (aToken, rToken string, err error) {
	calims := UserClaim{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(100)).Unix(), // 过期时间
			//Issuer:    "my-project",                                      // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, calims)
	// 生成 aToken
	aToken, err = token.SignedString([]byte(config.JwtKey))
	if err != nil {
		fmt.Println(err)
	}

	// rToken 不需要存储任何自定义数据
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * time.Duration(200)).Unix(), // 过期时间
		//Issuer:    "my-project",                                      // 签发人
	}).SignedString([]byte(config.JwtKey))
	return aToken, rToken, nil
}

// AnalyzeToken
// Token 解析
func AnalyzeToken(token string) (*UserClaim, error) {
	uc := new(UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return uc, errors.New("token is invalid")
	}
	return uc, err
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(config.JwtKey), nil
}

func ParasToken2(access_token string) (claims *UserClaim, err error) {
	var token *jwt.Token
	claims = new(UserClaim)
	token, err = jwt.ParseWithClaims(access_token, claims, keyFunc)
	if err != nil {
		return nil, err
	}
	if !token.Valid { // token 是否有效
		err = errors.New("invalid token")
	}
	return claims, nil
}

func RefreshToken(aToken, rToken string) (newToken, newrToken string, err error) {
	// 第一步 : 判断 rToken 格式对的，没有过期的
	if _, err := jwt.Parse(rToken, keyFunc); err != nil {
		return "", "", err
	}

	// todo 第二步：从旧的 aToken 中解析出 cliams 数据   过期了还能解析出来吗
	var claims UserClaim
	_, err = jwt.ParseWithClaims(aToken, &claims, keyFunc)
	v, _ := err.(*jwt.ValidationError)

	// 当 access token 是过期错误，并且 refresh token 没有过期就创建一个新的 access token
	if v.Errors == jwt.ValidationErrorExpired {
		return GenToken2(claims.Username, claims.Password)
	}
	return "", "", err
}

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 假设 token 的是在 url 中存放的
		authHeader := c.Request.Header.Get("Authorization")
		parts := strings.Split(authHeader, " ")
		atoken := parts[0]
		ftoken := parts[1]
		if atoken == "" || ftoken == "" {
			c.JSON(200, gin.H{
				"msg":  "没有携带 atoken 或者 ftoken",
				"code": 400,
			})
			c.Abort()
			return
		}
		parasToken, err := ParasToken2(atoken) // 解析 access_token
		if err == nil {                        // 当前的 access_token 格式对，没有过期
			c.JSON(200, gin.H{
				"msg":  "atoken 和 ftoken 没有过期",
				"data": parasToken,
				"code": 400,
			})
			c.Next()
			return
		}
		atoken, rToken, err := RefreshToken(atoken, ftoken)
		if err != nil {
			c.JSON(200, gin.H{
				"msg":  "您需要重新登录",
				"code": 400,
			})
			c.Abort()
			return
		} else {
			c.JSON(200, gin.H{
				"msg":    "atoken 和 ftoken 没有过期",
				"atoken": atoken,
				"rToken": rToken,
				"code":   400,
			})
			c.Next()
			return
		}
	}
}
