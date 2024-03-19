package templatemethod

import "testing"

func TestTemplateMethod(t *testing.T) {
	html := &Html{}
	head := &Head{}
	title := &Title{}
	body := &Body{}

	page := NewPage()

	textTitle := "Page title"
	textBody := "First line\nSecond line\nThird line\n"

	want := "<!DOCTYPE html>\n" +
		"\t<head>\n\t\t<title>" +
		textTitle + "</title>\n" + "\t</head>\n" +
		"\t<body>\n" + textBody + "\t</body>\n" +
		"</html>\n"

	page.SetTemplate(html)
	page.Start()

	page.SetTemplate(head)
	page.Start()

	page.SetTemplate(title)
	page.AddText(textTitle)

	page.SetTemplate(head)
	page.End()

	page.SetTemplate(body)
	page.AddText(textBody)

	page.SetTemplate(html)
	page.End()

	got := page.Page

	if want != got {
		t.Errorf("want: %s, got: %s\n", want, got)
	}
}
