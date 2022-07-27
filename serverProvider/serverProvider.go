package serverProvider

import (
	"example.com/m/provider"
	"net/http"
	"os"
	"time"

	"context"

	"github.com/sirupsen/logrus"
)

type Server struct {
	StorageProvider provider.StorageProvider
	httpServer      *http.Server
}

func SrvInit() *Server {
	sp := provider.NewStorageProvider(os.Getenv("GCP_Key"))
	return &Server{
		StorageProvider: sp,
	}
}

func (srv *Server) Start() {
	addr := ":" + os.Getenv("PORT")

	httpSrv := &http.Server{
		Addr:    addr,
		Handler: srv.SetupRoutes(),
	}

	srv.httpServer = httpSrv

	logrus.Info("Server running at PORT ", addr)
	if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.Fatalf("Start %v", err)
		return
	}
}

func (srv *Server) Stop() {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logrus.Info("closing server...")
	_ = srv.httpServer.Shutdown(ctx)
	logrus.Info("Done")
}
