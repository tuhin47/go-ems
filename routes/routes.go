package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/vivasoft-ltd/go-ems/controllers"
)

type Route struct {
	echo      *echo.Echo
	eventCtrl *controllers.EventController
}

func New(echo *echo.Echo, eventCtrl *controllers.EventController) *Route {
	return &Route{
		echo:      echo,
		eventCtrl: eventCtrl,
	}
}

func (r *Route) Init() {
	e := r.echo

	// Define your routes here
	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})

	g := e.Group("/v1")
	g.POST("/events", r.eventCtrl.CreateEvent)
	g.GET("/events/:id", r.eventCtrl.ReadEventByID)
}
