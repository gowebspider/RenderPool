package main

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/go-rod/rod/lib/utils"
)

func main() {
	l := launcher.MustNewManaged("ws://150.158.3.190:30713").
		Set("disable-gpu").Delete("disable-gpu").
		//XVFB("--server-num=5", "--server-args=-screen 0 1600x900x16").
		Headless(true)

	browser := rod.New().Client(l.MustClient()).MustConnect()
	launcher.Open(browser.ServeMonitor(``))
	defer browser.Close()
	pagePool := rod.NewPagePool(2 << 4)
	page := pagePool.Get(func() *rod.Page {
		return browser.MustIncognito().MustPage()
	})

	router := page.HijackRequests()
	//router.MustAdd("*.png", func(ctx *rod.Hijack) {
	router.MustAdd("*.jpg", func(ctx *rod.Hijack) {
		// 你可以使用很多其他 enum 类型，比如 NetworkResourceTypeScript 用于 javascript
		// 这个例子里我们使用 NetworkResourceTypeImage 来阻止图片
		if ctx.Request.Type() == proto.NetworkResourceTypeImage {
			ctx.Response.Fail(proto.NetworkErrorReasonBlockedByClient)
			return
		}
		ctx.ContinueRequest(&proto.FetchContinueRequest{})
	})

	defer pagePool.Put(page)
	page.MustNavigate(`https://www.github.com`)
	utils.Pause()
}
