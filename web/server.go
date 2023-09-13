package web

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"http-win-notice/utils/log"
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
			log.Debug("Error starting server: %v\n", err)
		}
	}()

	log.Debug("Server started on :" + setting.PortStr())

	return server
}

func closeServer(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Debug("web server shutdown error")
	} else {
		log.Debug("web server stop")
	}
}

func ServerManage(serverCh chan bool) {
	var server *http.Server
	for {
		if <-serverCh == true {
			log.Debug("start web server")
			server = runServer()
		} else {
			log.Debug("stop web server")
			closeServer(server)
		}

	}
}
