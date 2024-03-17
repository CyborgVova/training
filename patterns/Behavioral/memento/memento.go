package memento

type Object struct {
	Field int
}

type States struct {
	objects []Object
}

func NewStates(object Object) *States {
	return &States{objects: []Object{object}}
}

func (s *States) Save(object Object) {
	s.objects = append(s.objects, object)
}

func (s *States) GetLast() Object {
	len := len(s.objects)
	if len == 1 {
		return s.objects[0]
	}
	object := s.objects[len-1]
	s.objects = s.objects[:len-1]
	return object
}
