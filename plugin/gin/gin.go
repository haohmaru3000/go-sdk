package ginc

import (
	"flag"

	"github.com/gin-gonic/gin"
	sdk "github.com/haohmaru3000/go-sdk"
	logger "github.com/haohmaru3000/go-sdk/logger"
)

const (
	defaultPort = 3000
	defaultMode = "debug"
)

type Config struct {
	port    int
	ginMode string
}

type ginEngine struct {
	*Config
	name   string
	id     string
	logger logger.Logger
	router *gin.Engine
}

func NewGin(id string) *ginEngine {
	return &ginEngine{
		Config: new(Config),
		id:     id,
	}
}

func (gs *ginEngine) ID() string {
	return gs.id
}

func (gs *ginEngine) Activate(sv sdk.Service) error {
	gs.logger = sv.Logger(gs.id)
	gs.name = sv.Name()

	if gs.ginMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	gs.logger.Info("init engine...")
	gs.router = gin.New()

	return nil
}

func (gs *ginEngine) Stop() error {
	return nil
}

func (gs *ginEngine) InitFlags() {
	flag.IntVar(&gs.Config.port, "gin-port", defaultPort, "gin server port. Default 3000")
	flag.StringVar(&gs.Config.ginMode, "gin-mode", defaultMode, "gin mode (debug | release). Default debug")
}

func (gs *ginEngine) GetPort() int {
	return gs.port
}

func (gs *ginEngine) GetRouter() *gin.Engine {
	return gs.router
}
