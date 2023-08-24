package web

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"http-win-notice/utils/setting"
	"net/http"
	"time"
)

func runServer() *http.Server {
	r := gin.Default()
	InitRouters(r)
	//r.Run("0.0.0.0:" + setting.PortStr())
	server := &http.Server{
		Addr:    ":" + setting.PortStr(),
		Handler: r,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("Error starting server: %v\n", err)
		}
	}()

	fmt.Println("Server started on :" + setting.PortStr())
	
	return server
}

func closeServer(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("server shutdown error")
	} else {
		fmt.Println("server stop")
	}
}

func ServerManage(serverCh chan bool) {
	var server *http.Server
	for {
		if <-serverCh == true {
			fmt.Println("start")
			server = runServer()
		} else {
			fmt.Println("stop")
			closeServer(server)
		}

	}
}
