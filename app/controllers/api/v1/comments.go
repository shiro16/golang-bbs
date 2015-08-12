package controllers

import (
	"github.com/revel/revel"
)

type ApiV1Comments struct {
	ApiV1Controller
}

func (c ApiV1Comments) Index() revel.Result {
	r := Response{"index"}
	return c.RenderJson(r)
}

func (c ApiV1Comments) Show(id int) revel.Result {
	r := Response{"show"}
	return c.RenderJson(r)
}

func (c ApiV1Comments) Create() revel.Result {
	r := Response{"create"}
	return c.RenderJson(r)
}

func (c ApiV1Comments) Delete(id int) revel.Result {
	r := Response{"delete"}
	return c.RenderJson(r)
}
