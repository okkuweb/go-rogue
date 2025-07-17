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

type styler struct{}
func (sty styler) GetStyle(cst gruid.Style) tc.Style {
	st := tc.StyleDefault
	return st
}

func main() {
	opt := InitOptions()
	gd := gruid.NewGrid(opt.UIWidth, opt.UIHeight)
	md := &model{gd: gd, g: &game{}, opt: opt}

	st := styler{}
	dr := tcell.NewDriver(tcell.Config{StyleManager: st})
	//dr.PreventQuit()
	
	app := gruid.NewApp(gruid.AppConfig{
		Driver: dr,
		Model: md,
	})

	if err := app.Start(context.Background()); err != nil {
		fmt.Println(err)
	}
}

