package browser

import (
	"github.com/andrewarrow/feedback/wasm"
)

var Global *wasm.Global
var Document *wasm.Document

func RegisterEvents() {
	afterRegister := func(id int64) {
		Global.Location.Set("href", "/ava/start")
	}
	afterLogin := func(id int64) {
		Global.Location.Set("href", "/ava/start")
	}
	if Global.Start == "play.html" {
		b := Document.Id("board")
		canvas := b.JValue
		width := canvas.Get("width").Int()
		height := canvas.Get("height").Int()
		ctx := canvas.Call("getContext", "2d")
		ctx.Call("clearRect", 0, 0, width, height)
		ctx.Set("strokeStyle", "rgba(0, 0, 0, 1.0)")
		ctx.Call("beginPath")
		ctx.Call("moveTo", 9, 9)
		ctx.Call("lineTo", 9, 9+300)
		ctx.Call("stroke")
	} else if Global.Start == "login.html" {
		Global.AutoForm("login", "ava", nil, afterLogin)
	} else if Global.Start == "register.html" {
		Global.AutoForm("register", "ava", nil, afterRegister)
	}
}
