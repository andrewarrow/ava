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
		ctx := b.JValue.Call("getContext", "2d")
		ctx.Call("beginPath")
		ctx.Call("moveTo", 9, 9)
		ctx.Call("lineTo", 9, 9+300)
		ctx.Set("strokeStyle", "white")
		ctx.Call("stroke")
	} else if Global.Start == "login.html" {
		Global.AutoForm("login", "ava", nil, afterLogin)
	} else if Global.Start == "register.html" {
		Global.AutoForm("register", "ava", nil, afterRegister)
	}
}
