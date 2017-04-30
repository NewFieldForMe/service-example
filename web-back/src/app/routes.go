package main

import "net/http"

// Route :route structure
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes :route infomation
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"ArticleIndex",
		"GET",
		"/articles",
		ArticleIndex,
	},
	Route{
		"ArticleShow",
		"GET",
		"/articles/{articleId}",
		ArticleShow,
	},
	Route{
		"ArticleCreate",
		"POST",
		"/articles",
		ArticleCreate,
	},
}
