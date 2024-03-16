package chainofresponsibility

type Handler interface {
	HandleFunc(string) string
}

type HandlerA struct {
	next Handler
}

type HandlerB struct {
	next Handler
}

type HandlerC struct {
	next Handler
}

func (h *HandlerA) HandleFunc(request string) string {
	if request == "A" {
		return "response from A handler"
	} else if h.next != nil {
		return h.next.HandleFunc(request)
	}
	return ""
}

func (h *HandlerB) HandleFunc(request string) string {
	if request == "B" {
		return "response from B handler"
	} else if h.next != nil {
		return h.next.HandleFunc(request)
	}
	return ""
}

func (h *HandlerC) HandleFunc(request string) string {
	if request == "C" {
		return "response from C handler"
	} else if h.next != nil {
		return h.next.HandleFunc(request)
	}
	return ""
}
