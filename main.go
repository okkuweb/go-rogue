package main

import (
	"context"
	"fmt"
	"os"
	"time"

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
type model struct {
	gd gruid.Grid // user interface grid
	g *game
	// other fields with the state of the application
}
func Print(message gruid.Msg) error {
    // Open the file in append mode, create it if it doesn't exist
    f, err := os.OpenFile(LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return fmt.Errorf("failed to open log file: %w", err)
    }
    defer f.Close()
    
    // Format the current time
    timestamp := time.Now().Format("2006-01-02 15:04:05")
    
    // Write the timestamped message
    logEntry := fmt.Sprintf("[%s] %s\n", timestamp, message)
    if _, err := f.WriteString(logEntry); err != nil {
        return fmt.Errorf("failed to write to log file: %w", err)
    }
    
    return nil
}

func (m *model) Update(msg gruid.Msg) gruid.Effect {
    switch msg := msg.(type) {
    case gruid.MsgInit:
        // Initialize your grid here if needed
		Print(msg)
        return nil
    case gruid.MsgKeyDown:
		Print(msg)
        return gruid.End()
    default:
    }
    return nil
}
func (md *model) Draw() gruid.Grid {
	max := md.gd.Size()
	for x := range max.X {
		// Top and bottom borders
		md.gd.Set(gruid.Point{X: x, Y: 2}, gruid.Cell{Rune: '-'})
		md.gd.Set(gruid.Point{X: x, Y: max.Y-1}, gruid.Cell{Rune: '-'})
	}
	
	for y := range max.Y {
		// Left and right borders
		md.gd.Set(gruid.Point{X: 0, Y: y}, gruid.Cell{Rune: '|'})
		md.gd.Set(gruid.Point{X: max.X - 1, Y: y}, gruid.Cell{Rune: '|'})
	}
	return md.gd
}




type styler struct{}
func (sty styler) GetStyle(cst gruid.Style) tc.Style {
	st := tc.StyleDefault
	return st
}



func main() {
	// Create a new 20x20 grid.
	gd := gruid.NewGrid(UIWidth, UIHeight)
	// Define a range (5,5)-(15,15).
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

