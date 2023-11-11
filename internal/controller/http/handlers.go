package http

import (
	"github.com/labstack/echo/v4"
	"github.com/vlad-marlo/todoer/internal/model"
	"net/http"
)

func (ctrl *Controller) HandleTasksCreate(ctx echo.Context) error {
	var req model.CreateTaskRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	resp, err := ctrl.srv.Create(ctx.Request().Context(), req.Value, req.Status)
	if err != nil {
		return ctrl.handleErr(ctx, err)
	}
	return ctx.JSON(http.StatusCreated, resp)
}

func (ctrl *Controller) HandleTasksGetMany(ctx echo.Context) error {
	panic("not implemented")
}

func (ctrl *Controller) HandleTasksGetOne(ctx echo.Context) error {
	panic("not implemented")
}

func (ctrl *Controller) HandleTasksUpdate(ctx echo.Context) error {
	panic("not implemented")
}

func (ctrl *Controller) HandleTasksDelete(ctx echo.Context) error {
	panic("not implemented")
}

func (ctrl *Controller) HandleTasksSetStatus(ctx echo.Context) error {
	panic("not implemented")
}

func (ctrl *Controller) HandleTasksUploadFromFile(ctx echo.Context) error {
	panic("not implemented")
}

func (ctrl *Controller) HandleTasksDownloadToFile(ctx echo.Context) error {
	panic("not implemented")
}
