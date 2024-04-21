package main

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/unit"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	file, err := os.Create("testFile")
	check(err)
	defer file.Close()

	go func() {
		window := new(app.Window)
		err := run(window)
		check(err)
		os.Exit(0)
	}()
	app.Main()

}

func run(window *app.Window) error {
	var ops op.Ops
	theme := material.NewTheme()

	editor := new(widget.Editor)
	editor.SingleLine = true
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			layout.Flex{Axis: layout.Vertical}.Layout(
				gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.UniformInset(unit.Dp(theme.TextSize)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return material.Editor(theme, editor, "Enter Text ...").Layout(gtx)
					})
				}),
			)
			e.Frame(gtx.Ops)
		}
	}
}

func test(window *app.Window) error {
	theme := material.NewTheme()
	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			// This graphics context is used for managing the rendering state.
			bcolor := color.NRGBA{R: 252, G: 2, B: 197, A: 255}
			draw(&ops, e)

			theme.Bg = bcolor
			material.NewTheme().Bg = bcolor
			// Define an large label with an appropriate text:
		}
	}
}

func addColorOperation(ops *op.Ops) {
	red := color.NRGBA{R: 0xFF, A: 0xFF}
	paint.ColorOp{Color: red}.Add(ops)
}

func draw(ops *op.Ops, e app.FrameEvent) {
	theme := material.NewTheme()
	gtx := app.NewContext(ops, e)
	title := material.H1(theme, "hello, World")
	maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
	title.Color = maroon
	title.Alignment = text.Middle
	title.Layout(gtx)
	e.Frame(gtx.Ops)

}
