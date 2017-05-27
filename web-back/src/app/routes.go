package main

import (
	"app/controller"
	"net/http"
)

// Route :route structure
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes :route infomation
type Routes []Route

const _apiPath = "/api/v1"

var routes = Routes{
	Route{
		"ArticleIndex",
		"GET",
		_apiPath + "/articles",
		controller.ArticleIndex,
	},
	Route{
		"ArticleShow",
		"GET",
		_apiPath + "/articles/{articleId}",
		controller.ArticleShow,
	},
	Route{
		"ArticleCreate",
		"POST",
		_apiPath + "/articles",
		controller.ArticleCreate,
	},
	Route{
		"SignUp",
		"POST",
		_apiPath + "/users",
		controller.SignUp,
	},
}
