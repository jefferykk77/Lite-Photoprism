package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 创建一个默认的 Gin 引擎 (路由器)
	// Default() 会自带两个大杀器中间件：Logger 和 Recovery
	// 以防你的机器因为一个 bug 的恐慌 (panic) 直接挂掉
	r := gin.Default()

	// 2. 注册路由组
	// 以后哪怕换版本了，加上 /api/v1 前缀也可以优雅过渡
	api := r.Group("/api")
	{
		// 3. 注册具体到端点的 Handler
		// 当 HTTP 请求的方法是 GET 且路径匹配时，就会执行这段匿名函数
		api.GET("/ping", func(c *gin.Context) {
			// 把字典数据封装成 JSON 格式响应回浏览器
			c.JSON(http.StatusOK, gin.H{
				"message": "pong from photoprism-lite",
			})
		})
	}

	// 4. 死循环监听端口
	// 这行代码底层就是我们前面说到的：绑定 0.0.0.0:8080 并开启底层 for 死循环
	log.Println("Server is running on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
}
