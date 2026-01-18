package main

import (
	"context"
	"fmt"

	"codeberg.org/anaseto/gruid"
)

func main() {
	InitLogger()
	defer logFile.Close()
	// TODO: Move this grid stuff to a grid file
	opt := &options{width: MapWidth, height: MapHeight}
	gd := gruid.NewGrid(opt.width, opt.height)
	md := &model{grid: gd}
	initDriver()

	app := gruid.NewApp(gruid.AppConfig{
		Driver: driver,
		Model:  md,
	})

	if err := app.Start(context.Background()); err != nil {
		fmt.Println(err)
	}
}

