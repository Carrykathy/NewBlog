package main

import (
	"NewBlog/models"
	"NewBlog/pkg/logging"
	"NewBlog/pkg/setting"
	"NewBlog/routers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()

	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	router := routers.InitRouter()

	// 使用标准 net/http server 替代 endless
	server := &http.Server{
		Addr:           endPoint,
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("Start server at %s", endPoint)
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server error: %v", err)
	}
}
