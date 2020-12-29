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
		"default",
		"GET",
		"/health",
		defHealthCheck,
	},
	Route{
		"customhealth",
		"GET",
		"/actuator/health",
		custHealthCheck,
	},

	Route{
		"callecho",
		"GET",
		"/callecho",
		callecho,
	},

	Route{
		"printenv",
		"GET",
		"/printenv",
		printenv,
	},
}
