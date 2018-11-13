package config

import (
	"wonder_go_df/application/controllers"
	"wonder_go_df/system/core"
)

var RoutesVal = core.Routes{
	core.Route{
		Name:        "Index",
		Method:      "GET",
		Pattern:     "/",
		HandlerFunc: controllers.Index, // 指向要调用的控制器 en: Point to the controller to be called
	},
}
