package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//创建一个路由引擎
	r := gin.Default()
	//请求方式+路径+行为
	r.GET("/json", func(c *gin.Context) {

		//定义json格式返回内容
		//方法一 使用map
		//data := map[string]interface{}{
		//	"name": "zhangsan",
		//	"age":  18,
		//}
		//
		data := gin.H{
			"name": "zhangsan",
			"age":  18,
		}
		//方法二 使用结构体 可以使用tag来灵活化
		type msg struct {
			Name string `json:"name"`
			Age  int
		}
		data1 := msg{Name: "Lisi", Age: 18}

		c.JSON(http.StatusOK, data)
	})
	r.Run(":8088")
}
