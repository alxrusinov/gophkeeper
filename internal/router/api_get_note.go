package router

import (
	"fmt"
	"net/http"

	"github.com/kataras/iris/v12"
)

func (r *Router) GetNote(ctx iris.Context) {
	fmt.Printf("USER - %#v", ctx.Values())
	ctx.StatusCode(http.StatusOK)
}
