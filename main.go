package main

import (
	"NewBlog/pkg/setting"
	"NewBlog/routers"
	"fmt"
	"net/http"
)

func main() {

	router := routers.InitRouter()

	s := &http.Server{
		Addr: fmt.Sprintf(":%d", setting.HTTPPort),
		//http句柄，实质为ServerHTTP，用于处理程序响应HTTP请求
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
