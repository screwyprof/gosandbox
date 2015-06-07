package main

import (
	"github.com/screwyprof/gosandbox/application/bootstrap"
)

func main() {
	app := bootstrap.Bootstrap()
	app.Run()
}
