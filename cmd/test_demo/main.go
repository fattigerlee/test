package main

import "github.com/gin-gonic/gin"

func main() {
	test3()
}

func test3() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.Writer.Write([]byte("hello world, gin!!!"))
	})

	router.Run(":81")
}
