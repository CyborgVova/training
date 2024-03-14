package abstractfactory

type AbstractComputerFactory interface {
	CreateOS(family string) AbstractOS
	CreateHardware(memory, processor, motherboard, videoplate string) AbstractHardware
}

type AbstractOS interface {
	GetFamily() string
}

type AbstractHardware interface {
	InstallOS(os AbstractOS)
	GetOSFamily() string
	GetMemory() string
	GetProcessor() string
	GetMotherboard() string
	GetVideoplate() string
}

type ComputerFactory struct{}

func NewComputer() AbstractComputerFactory {
	return &ComputerFactory{}
}

type OS struct {
	Family string
}

func (c *ComputerFactory) CreateOS(family string) AbstractOS {
	return &OS{Family: family}
}

type Hardware struct {
	OS          AbstractOS
	Memory      string
	Processor   string
	Motherboard string
	Videoplate  string
}

func (o *OS) GetFamily() string {
	return o.Family
}

func (c *ComputerFactory) CreateHardware(memory, processor, motherboard, videoplate string) AbstractHardware {
	return &Hardware{
		Memory:      memory,
		Processor:   processor,
		Motherboard: motherboard,
		Videoplate:  videoplate,
	}
}

func (h *Hardware) InstallOS(os AbstractOS) {
	h.OS = os
}

func (h *Hardware) GetOSFamily() string {
	return h.OS.GetFamily()
}

func (h *Hardware) GetMemory() string {
	return h.Memory
}

func (h *Hardware) GetProcessor() string {
	return h.Processor
}

func (h *Hardware) GetMotherboard() string {
	return h.Motherboard
}

func (h *Hardware) GetVideoplate() string {
	return h.Videoplate
}
