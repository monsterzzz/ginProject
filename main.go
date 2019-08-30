package main

import (
	"fmt"
	"ginProject/pkg/setting"
	"ginProject/routers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	gin.SetMode(setting.RunMode)

	r := routers.InitRouters()

	s := &http.Server{
		Addr: fmt.Sprintf(":%d",setting.HttpPort),
		Handler: r,
		ReadTimeout: setting.ReadTimeOut,
		WriteTimeout: setting.WriteTimeOuT,
		MaxHeaderBytes: 1 << 20,

	}

	s.ListenAndServe()


}