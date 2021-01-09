package routes

import "github.com/labstack/echo/v4"

type Endpoint struct {
	Method  string
	Group   string
	Path    string
	Handler echo.HandlerFunc
}

func GetRoutes() []Endpoint {
	var routes []Endpoint

	routes = append(routes, statusRoutes...)
	routes = append(routes, userRoutes...)
	routes = append(routes, accountRoutes...)

	return routes
}
