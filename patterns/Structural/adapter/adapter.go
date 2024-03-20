package adapter

type Usb interface {
	ConnectToComputer() string
}

type CardReader struct {
	MemoryCard MemoryCard
}

func (c *CardReader) ConnectToComputer() (result string) {
	result += "Connection ...\n"
	result += c.MemoryCard.ConnectToDevice()
	result += c.MemoryCard.CopyData()
	return
}

type MemoryCard interface {
	ConnectToDevice() string
	CopyData() string
}

type SdCard struct{}

func (f *SdCard) ConnectToDevice() string {
	return "SdCard is connected\n"
}

func (f *SdCard) CopyData() string {
	return "Data was copied\n"
}

type MicroSd struct{}

func (f *MicroSd) ConnectToDevice() string {
	return "MicroSd is connected\n"
}

func (f *MicroSd) CopyData() string {
	return "Data was copied\n"
}
