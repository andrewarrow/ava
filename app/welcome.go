package app

import (
  "github.com/andrewarrow/feedback/router"
)

func Welcome(c *router.Context, second, third string) {
	if second == "" && third == "" && c.Method == "GET" {
		handleWelcomeIndex(c)
		return
	}
	c.NotFound = true
}

func handleWelcomeIndex(c *router.Context) {

	send := map[string]any{}
	if len(c.User) == 0 {
		c.SendContentInLayout("welcome.html", send, 200)
		return
	}
}
    