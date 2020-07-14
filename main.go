package main

import (
	"fmt"
	"net/http"

	"learn.gin/pkg/setting"
	"learn.gin/routers"
)

func main() {
	/*
		type Handler interface {
			ServeHTTP(ResponseWriter, *Request)
		}
	*/
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        routers.InitRouter(),
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
