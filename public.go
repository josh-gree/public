package main

import (
	"github.com/labstack/echo"
	"github.com/josh-gree/comm"
)

var j = comm.JobMessage{}
var r = comm.ResMessage{}
var public = true // read from cml

func main(){
	e := echo.New()

	e.POST("/job", j.Recieve(public))
	e.POST("/res", r.Recieve(public))

	e.Start(":8000")
}
