package RenderPool

// A RenderOption sets an option on a RenderPool.
type RenderOption func(*RenderPool)

type RenderPool struct {
	RenderServerHost string
}

func NewRenderPool(options ...RenderOption) *RenderPool {
	r := &RenderPool{}
	r.Init()

	for _, f := range options {
		f(r)
	}

	return r
}

func (r *RenderPool) Init() {
	r.RenderServerHost = `ws://127.0.0.1:3717`

}
