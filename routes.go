package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"ProviderIndex",
		"GET",
		"/providers",
		ProviderIndex,
	},
	Route{
		"ProviderShow",
		"GET",
		"/providers/{providerId}",
		ProviderShow,
	},
	Route{
		"ProviderCreate",
		"POST",
		"/providers",
		ProviderCreate,
	},
}
