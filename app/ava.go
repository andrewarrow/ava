package app

import (
	"github.com/andrewarrow/feedback/router"
)

func Ava(c *router.Context, second, third string) {
	if second == "play" && third == "" && c.Method == "GET" {
		handleAvaPlay(c)
		return
	}
	c.NotFound = true
}

func handleAvaPlay(c *router.Context) {

	send := map[string]any{}
	c.SendContentInLayout("play.html", send, 200)
}
