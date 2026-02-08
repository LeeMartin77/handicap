package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leemartin77/handicap/internal/config"
	"github.com/leemartin77/handicap/internal/storage"
)

type Server interface {
	RunServer() error
}

type serverState struct {
	strg storage.Storage
}

func NewServer(ctx context.Context, cfg *config.Config) (Server, error) {
	strg, err := storage.NewStorage(ctx, cfg)
	if err != nil {
		return nil, err
	}
	return &serverState{
		strg: strg,
	}, nil
}

// RunServer implements [Server].
func (s *serverState) RunServer() error {
	// TODO: make setup part of initialisation
	router := gin.Default()
	router.Static("static", "web/static")
	router.LoadHTMLGlob("web/template/*")
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": s.strg.GetTestData(c),
		})
	})
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"Message": s.strg.GetTestData(ctx),
		})
	})
	return router.Run()
}
