package middleware

import (
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/yangle520/gin-vue-demo/server/model/system"
)

var AuthMiddleware *jwt.GinJWTMiddleware

func InitJWT() {
	var err error
	AuthMiddleware, err = jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "gin-vue-demo",
		Key:         []byte("secret-key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour * 24,
		IdentityKey: "id",

		// 登录验证
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals struct {
				UserName string `json:"username"`
				Password string `json:"password"`
			}
			if err := c.ShouldBindJSON(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			// TODO 去数据库验证用户名密码
			if loginVals.UserName == "admin" && loginVals.Password == "admin" {
				return &system.User{ID: 1, LoginName: "admin"}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},

		// 载荷生成
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*system.User); ok {
				return jwt.MapClaims{
					"id":        v.ID,
					"loginName": v.LoginName,
				}
			}
			return jwt.MapClaims{}
		},
	})

	if err != nil {
		panic("JWT初始化失败：" + err.Error())
	}
}
