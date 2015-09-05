package main

import (
	"beegoapp"
	_ "routers"
)

func init() {
	// initialize the trace.

	beegoapp.TheApp.Initialize()
	beegoapp.TheApp.Run()
}
