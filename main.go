package main

import (
	"fmt"
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

	go func() {
		window := new(app.Window)
		err := run(window)
		check(err)
		os.Exit(0)
	}()
	app.Main()

}

func run(window *app.Window) error {
	file, err := os.Create("testFile.txt")
	check(err)
	defer file.Close()

	var ops op.Ops
	theme := material.NewTheme()

	editor := new(widget.Editor)
	editor.SingleLine = false
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			// input.Source.Event()
			// saveEvent := key.Event{Name: "S", Modifiers: key.ModShift, State: key.Press}

			layout.Flex{Axis: layout.Vertical}.Layout(
				gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.UniformInset(unit.Dp(theme.TextSize)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return material.Editor(theme, editor, "Enter Text ...").Layout(gtx)
					})
				}),
			)
			fmt.Println(editor.Text())
			file.WriteString(editor.Text())

			e.Frame(gtx.Ops)
		}
	}
}

//	func handleInput(window *app.Window, editor *widget.Editor) {
//		for _, e := range window.Events(key.Tag(editor)) {
//			if e, ok := e.(key.Event); ok && e.State == key.Press && e.Name == "S" && e.Modifiers.Contain(key.ModCtrl) {
//				save(editor.Text())
//			}
//		}
//	}
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
