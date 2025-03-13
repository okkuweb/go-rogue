package main

import (
	"context"
	"fmt"

	"codeberg.org/anaseto/gruid"
	tcell "codeberg.org/anaseto/gruid-tcell"
	tc "github.com/gdamore/tcell/v2"
)

type game struct {
}

const (
	UIWidth  = 80
	UIHeight = 24
	LogFile = "./log.txt"
)


type styler struct{}
func (sty styler) GetStyle(cst gruid.Style) tc.Style {
	st := tc.StyleDefault
	return st
}

func main() {
	gd := gruid.NewGrid(UIWidth, UIHeight)
	m := &model{gd: gd, g: &game{}}

	st := styler{}
	dr := tcell.NewDriver(tcell.Config{StyleManager: st})
	//dr.PreventQuit()
	
	app := gruid.NewApp(gruid.AppConfig{
		Driver: dr,
		Model: m,
	})

	if err := app.Start(context.Background()); err != nil {
		fmt.Println(err)
	}
}

