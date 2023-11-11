package http

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/vlad-marlo/todoer/internal/config"
	"github.com/vlad-marlo/todoer/internal/controller"
	"go.uber.org/zap"
	"net/http"
)

type Controller struct {
	log    *zap.Logger
	cfg    *config.Server
	srv    controller.Service
	server *http.Server
	router *echo.Echo
}

func New(
	logger *zap.Logger,
	cfg *config.Server,
	service controller.Service,
) (*Controller, error) {
	srv := &Controller{
		router: echo.New(),
		log:    logger,
		cfg:    cfg,
		srv:    service,
	}
	if logger == nil || cfg == nil || service == nil {
		return nil, ErrNilReference
	}
	srv.configure()
	logger.Info("successful initialized server")
	return srv, nil
}

func (ctrl *Controller) configure() {
	ctrl.configureMiddleWares()
	ctrl.configureRoutes()
}

func (ctrl *Controller) configureMiddleWares() {

}

func (ctrl *Controller) configureRoutes() {
	apiV1 := ctrl.router.Group("/api/v1")
	tasks := apiV1.Group("/tasks")
	{
		tasks.POST("/", ctrl.HandleTasksCreate)
		tasks.GET("/", ctrl.HandleTasksGetMany)
		tasks.GET("/:task_id", ctrl.HandleTasksGetOne)
		tasks.PATCH("/:task_id", ctrl.HandleTasksUpdate)
		tasks.DELETE("/:task_id", ctrl.HandleTasksDelete)
		tasks.POST("/:task_id/status", ctrl.HandleTasksSetStatus)
		tasks.POST("/file", ctrl.HandleTasksUploadFromFile)
		tasks.POST("/file", ctrl.HandleTasksDownloadToFile)
	}
}

func (ctrl *Controller) Start(_ context.Context) error {
	go func() {
		ctrl.log.Error("starting http server", zap.Error(ctrl.router.Start(fmt.Sprintf("%s:%d", ctrl.cfg.BindAddr, ctrl.cfg.BindPort))))
	}()
	ctrl.log.Info("starting http server", zap.String("bind_addr", fmt.Sprintf("%s:%d", ctrl.cfg.BindAddr, ctrl.cfg.BindPort)))
	return nil
}

func (ctrl *Controller) Stop(ctx context.Context) error {
	ctrl.log.Info("stopping http server", zap.String("bind_addr", fmt.Sprintf("%s:%d", ctrl.cfg.BindAddr, ctrl.cfg.BindPort)))
	return ctrl.router.Shutdown(ctx)
}
