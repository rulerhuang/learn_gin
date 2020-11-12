package main

import (
	"fmt"
	"log"
	"net/http"

	"learn.gin/pkg/logging"
	"learn.gin/pkg/setting"
	"learn.gin/routers"
)

func main() {
	/*
		type Handler interface {
			ServeHTTP(ResponseWriter, *Request)
		}
	*/

	logging.Info("will start http server")
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        routers.InitRouter(),
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("call ListenAndServe() faild, err=%v\n", err)
	}
}
