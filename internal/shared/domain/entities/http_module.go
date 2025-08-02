package entities

import "github.com/labstack/echo/v4"

type HTTPModule interface {
	GetControllers() []Controller
}

type Controller interface {
	GetBind() ControllerBind
	Execute(ctx echo.Context) error
}

type ControllerBind struct {
	Method       string
	Version      string
	RelativePath string
}

func (cb ControllerBind) GetFullPath() string {
	return cb.Version + cb.RelativePath
}
