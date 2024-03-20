package adapter

type Usb interface {
	ConnectToComputer() string
}

type CardReader struct {
	SdCard *SdCard
}

func (c *CardReader) ConnectToComputer() (result string) {
	result += "Connection ...\n"
	result += c.SdCard.ConnectToDevice()
	result += c.SdCard.CopyData()
	return
}

type SdCard struct{}

func (f *SdCard) ConnectToDevice() string {
	return "Connection is done\n"
}

func (f *SdCard) CopyData() string {
	return "Data was copied\n"
}
