package main

import (
    _ "routers"
   "beegoapp"

)

func init() {
	// initialize the trace.

    beegoapp.TheApp.Initialize()
    beegoapp.TheApp.Run()
}
