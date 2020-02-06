package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//test1()

	//test2()

	//test3()

	test4()
}

func test1() {
	fmt.Println("hello world")
}

func test2() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello world, 世界!!!"))
	})

	server := &http.Server{}
	server.Addr = fmt.Sprintf(":80")
	server.Handler = http.DefaultServeMux

	// 优雅关闭
	sign := make(chan os.Signal)  // 系统信号
	finish := make(chan struct{}) // 结束信号
	signal.Notify(sign, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	go func() {
		res := <-sign
		fmt.Println("信号类型:", res)

		fmt.Println("收到关闭信号")

		// 设置最长等待时间
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		err := server.Shutdown(ctx)
		if err != nil {
			fmt.Println("shutdown error:", err)
		}

		fmt.Println("处理结束任务...")
		time.Sleep(time.Second * 5)
		fmt.Println("结束任务处理完成...")
		finish <- struct{}{}
	}()

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		fmt.Println("listen and server error:", err)
	}
	<-finish
}

func test3() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.Writer.Write([]byte("hello world, gin!!!"))
	})

	router.Run(":80")
}

func test4() {
	str := "hello world,你好,世界!!!"
	fmt.Println("data:", url.QueryEscape(str))
}
