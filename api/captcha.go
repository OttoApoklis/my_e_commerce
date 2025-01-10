package api

import "github.com/gin-gonic/gin"

var store = base64Captcha.DefaultMemStore

type BaseApi struct{}

// 验证码相关 TODO
func (b *BaseApi) Captcha(c *gin.Context) {

}

// 类型转换
func interfaceToInt(v interface{}) (i int) {
	switch v := v.(type) {
	case int:
		i = v
	default:
		i = 0
	}
	return
}
