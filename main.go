package main

import (
	"context"
	"fmt"

	"codeberg.org/anaseto/gruid"
	tcell "codeberg.org/anaseto/gruid-tcell"
	tc "github.com/gdamore/tcell/v2"
)

type styler struct{}
func (sty styler) GetStyle(cst gruid.Style) tc.Style {
	st := tc.StyleDefault
	return st
}

func main() {
	opt := &options{width: 80, height: 24}
	gd := gruid.NewGrid(opt.width, opt.height)
	md := &model{grid: gd}
	st := styler{}
	dr := tcell.NewDriver(tcell.Config{StyleManager: st})

	app := gruid.NewApp(gruid.AppConfig{
		Driver: dr,
		Model: md,
	})

	if err := app.Start(context.Background()); err != nil {
		fmt.Println(err)
	}
}
