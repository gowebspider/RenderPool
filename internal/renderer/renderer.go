package renderer

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	configs "github.com/gowebspider/renderpool/configs/dev"
)

type Renderer struct {
	configs.RenderClient
	l       *launcher.Launcher
	browser *rod.Browser
	pool    rod.PagePool
}

func (r *Renderer) init() {
	r.connect()
	r.browser = rod.New().Client(r.l.MustClient()).MustConnect()
	router := r.browser.HijackRequests()
	defer router.MustStop()
	router.MustAdd(`*.bmp|*.jpg|*.jpeg|*.png|*.gif`, blockImage)
	go router.Run()
	r.pool = rod.NewPagePool(r.RenderPoolSize)
}

func (r *Renderer) connect() {
	r.l = launcher.MustNewManaged(r.RenderServerURI)
	r.l.Set(`disable-gpu`).
		Delete(`disable-gpu`).
		Headless(true)

}

func blockImage(hijack *rod.Hijack) {
	if hijack.Request.Type() == proto.NetworkResourceTypeImage {
		hijack.Response.Fail(proto.NetworkErrorReasonBlockedByClient)
		return
	}
	hijack.ContinueRequest(&proto.FetchContinueRequest{})
}
