package singleton

type instance struct {
	Id   string
	Name string
}

var (
	singleton *instance
)

func GetInstance(id, name string) *instance {
	if singleton == nil {
		singleton = &instance{Id: id, Name: name}
	}
	return singleton
}
