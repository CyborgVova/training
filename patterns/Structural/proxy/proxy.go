package proxy

type Common interface {
	Walking() string
	Run() string
}

type Object struct{}

func (o *Object) Walking() string {
	return "Walking training\n"
}

func (o *Object) Run() string {
	return "Run training\n"
}

type Proxy struct {
	Object Common
}

func NewProxy(common Common) *Proxy {
	return &Proxy{Object: common}
}

func (p *Proxy) Walking() string {
	return "Start " + p.Object.Walking() + "End " + p.Object.Walking()
}

func (p *Proxy) Run() string {
	return "Start " + p.Object.Run() + "End " + p.Object.Run()
}
