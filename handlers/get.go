package handlers

import routing "github.com/rmnoff/fasthttp-routing/v3"

func Get(c *routing.Context) error {
	return c.Write("hello")
}
