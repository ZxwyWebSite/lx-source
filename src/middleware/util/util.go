package util

import "github.com/gin-gonic/gin"

// 将路由参数转为Map
func ParaMap(c *gin.Context) map[string]string {
	parmlen := len(c.Params)
	parms := make(map[string]string, parmlen)
	for i := 0; i < parmlen; i++ {
		parms[c.Params[i].Key] = c.Params[i].Value
	}
	return parms
}
