package server

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/romycode/bank-manager/server/routes"
)

func Serve(port string) {
	api := echo.New()

	registerRoutes(api)

	api.Logger.Fatal(api.Start(fmt.Sprintf(":%s", port)))
}

func registerRoutes(api *echo.Echo) {
	var endpoints []routes.Endpoint
	endpoints = append(endpoints, routes.GetRoutes()...)

	groups := make(map[string]*echo.Group)

	for _, route := range endpoints {
		groupName := route.Group

		switch groupName {
		case "":
			api.Add(route.Method, route.Path, route.Handler)
		default:
			apiGroup, ok := groups[groupName]
			if !ok {
				groups[route.Group] = api.Group(groupName)
				apiGroup = groups[route.Group]
			}
			apiGroup.Add(route.Method, route.Path, route.Handler)
		}
	}
}
