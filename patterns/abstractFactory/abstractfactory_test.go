package abstractfactory

import (
	"testing"
)

func TestIntelOS(t *testing.T) {
	comp := NewComputer()

	ubuntu := comp.CreateOS("Ubuntu")

	intel := comp.CreateHardware("kingstone 4Gb", "Intel i5", "Asus PG5", "nVideo PGX1800")
	intel.InstallOS(ubuntu)

	want := ubuntu.GetFamily()
	got := intel.GetOSFamily()

	if want != got {
		t.Errorf("want: %s, got: %s\n", want, got)
	}

}

func TestAmdOS(t *testing.T) {
	comp := NewComputer()

	windows := comp.CreateOS("Windows")

	amd := comp.CreateHardware("Gigabite 8Gb", "AMD 2200", "NoName", "ATI r530x")
	amd.InstallOS(windows)

	want := windows.GetFamily()
	got := amd.GetOSFamily()

	if want != got {
		t.Errorf("want: %s, got: %s\n", want, got)
	}
}
