package memento

type Object struct {
	Field int
}

type states struct {
	objects []Object
}

func NewStates(object Object) *states {
	return &states{objects: []Object{object}}
}

func (s *states) Save(object Object) {
	s.objects = append(s.objects, object)
}

func (s *states) GetLast() Object {
	len := len(s.objects)
	if len == 1 {
		return s.objects[0]
	}
	object := s.objects[len-1]
	s.objects = s.objects[:len-1]
	return object
}
