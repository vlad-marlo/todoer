package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/vlad-marlo/todoer/internal/config"
	"github.com/vlad-marlo/todoer/internal/controller"
	"github.com/vlad-marlo/todoer/internal/model"
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
	if logger == nil || cfg == nil || service == nil {
		return nil, ErrNilReference
	}

	srv := &Controller{
		router: echo.New(),
		log:    logger.With(zap.String("layer", "transport")),
		cfg:    cfg,
		srv:    service,
	}
	srv.configure()
	logger.Info("successful initialized server")
	return srv, nil
}

func (ctrl *Controller) configure() {
	ctrl.configureMiddleWares()
	ctrl.configureRoutes()
	ctrl.log.Info("configured all routes")
}

func (ctrl *Controller) configureMiddleWares() {
	ctrl.router.Use(middleware.Logger())
	ctrl.router.Use(middleware.CORS())
	ctrl.router.Use(middleware.Recover())
}

//goland:noinspection ALL
func (ctrl *Controller) configureRoutes() {
	apiV1 := ctrl.router.Group("/api/v1")
	tasks := apiV1.Group("/tasks")
	{
		tasks.POST("", ctrl.HandleTasksCreate)
		tasks.GET("", ctrl.HandleTasksGetMany)
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
		ctrl.log.Error(
			"got error while starting http server", zap.Error(
				ctrl.router.Start(fmt.Sprintf("%s:%d", ctrl.cfg.BindAddr, ctrl.cfg.BindPort)),
			),
		)
	}()
	ctrl.log.Info("starting http server", zap.String("bind_addr", fmt.Sprintf("%s:%d", ctrl.cfg.BindAddr, ctrl.cfg.BindPort)))
	return nil
}

func (ctrl *Controller) Stop(ctx context.Context) error {
	ctrl.log.Info("stopping http server", zap.String("bind_addr", fmt.Sprintf("%s:%d", ctrl.cfg.BindAddr, ctrl.cfg.BindPort)))
	return ctrl.router.Shutdown(ctx)
}

func (ctrl *Controller) handleErr(ctx echo.Context, err error, fields ...zap.Field) error {
	var errMsg model.ErrorMessage
	if !errors.As(err, &errMsg) {
		errMsg.Status = fmt.Sprintf("unknown error: %v", err)
		errMsg.Code = http.StatusInternalServerError
		fields = append(fields, zap.Error(err))
		ctrl.log.Error("got unexpected error", fields...)
	}
	return ctx.JSON(http.StatusInternalServerError, errMsg)
}
