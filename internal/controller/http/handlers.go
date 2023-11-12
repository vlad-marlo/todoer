package http

import (
	"github.com/labstack/echo/v4"
	"github.com/vlad-marlo/todoer/internal/model"
	"go.uber.org/zap"
	"net/http"
)

var (
	ErrNotImplemented error = model.ErrorMessage{
		Endpoint: "",
		Code:     http.StatusNotImplemented,
		Status:   "not implemented",
	}
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
	var (
		req = model.GetManyTasksRequest{
			Limit: 10,
		}
		resp     *model.GetTasksResponse
		statuses []model.Status
		err      error
	)

	if err = ctx.Bind(&req); err != nil {
		ctrl.log.Info("got error", zap.Error(err))
		return ctrl.handleErr(ctx, err)
	}

	ctrl.log.Info("got request", zap.Any("request", req))

	statuses, err = model.ParseManyStatuses(req.Status, ",")
	if err != nil {
		return ctrl.handleErr(ctx, err)
	}

	resp, err = ctrl.srv.GetMany(ctx.Request().Context(), req.Offset, req.Limit, req.Task, statuses...)
	if err != nil {
		return ctrl.handleErr(ctx, err)
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (ctrl *Controller) HandleTasksGetOne(ctx echo.Context) error {
	resp, err := ctrl.srv.GetOne(ctx.Request().Context(), ctx.Param("task_id"))
	if err != nil {
		return ctrl.handleErr(ctx, err)
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (ctrl *Controller) HandleTasksUpdate(ctx echo.Context) error {
	var req model.UpdateTaskRequest
	if err := ctx.Bind(&req); err != nil {
		return ctrl.handleErr(ctx, err)
	}

	status, err := model.ParseStatus(req.Status)
	if err != nil {
		return ctrl.handleErr(ctx, err)
	}

	resp, err := ctrl.srv.Update(ctx.Request().Context(), ctx.Param("task_id"), req.Task, *status)
	if err != nil {
		return ctrl.handleErr(ctx, err)
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (ctrl *Controller) HandleTasksDelete(ctx echo.Context) error {
	err := ctrl.srv.Delete(ctx.Request().Context(), ctx.Param("task_id"))
	if err != nil {
		return ctrl.handleErr(ctx, err)
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (ctrl *Controller) HandleTasksSetStatus(ctx echo.Context) error {
	var req struct {
		Status string `json:"status" query:"status" form:"status"`
		TaskID string `param:"task_id"`
	}

	if err := ctx.Bind(&req); err != nil {
		return ctrl.handleErr(ctx, err)
	}

	status, err := model.ParseStatus(req.Status)
	if err != nil {
		return ctrl.handleErr(ctx, err)
	}

	resp, err := ctrl.srv.ChangeStatus(ctx.Request().Context(), req.TaskID, *status)
	if err != nil {
		return ctrl.handleErr(ctx, err)
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (ctrl *Controller) HandleTasksUploadFromFile(ctx echo.Context) error {
	return ctrl.handleErr(ctx, ErrNotImplemented)
}

func (ctrl *Controller) HandleTasksDownloadToFile(ctx echo.Context) error {
	return ctrl.handleErr(ctx, ErrNotImplemented)
}
