package templatemethod

type Template interface {
	Start() string
	End() string
}

type Html struct{}

func (h *Html) Start() string {
	return "<!DOCTYPE html>\n"
}

func (h *Html) End() string {
	return "</html>\n"
}

type Head struct{}

func (h *Head) Start() string {
	return "\t<head>\n"
}

func (h *Head) End() string {
	return "\t</head>\n"
}

type Title struct{}

func (t *Title) Start() string {
	return "\t\t<title>"
}

func (t *Title) End() string {
	return "</title>\n"
}

type Body struct{}

func (b *Body) Start() string {
	return "\t<body>\n"
}

func (b *Body) End() string {
	return "\t</body>\n"
}

type Page struct {
	template Template
	Page     string
}

func NewPage() *Page {
	return &Page{template: &Html{}}
}

func (p *Page) SetTemplate(template Template) {
	p.template = template
}

func (p *Page) Start() {
	p.Page += p.template.Start()
}

func (p *Page) End() {
	p.Page += p.template.End()
}

func (p *Page) AddText(text string) {
	p.Page += p.template.Start()
	p.Page += text
	p.Page += p.template.End()
}
