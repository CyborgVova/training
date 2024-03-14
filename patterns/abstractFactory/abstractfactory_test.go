package abstractfactory

import (
	"testing"
)

func TestIntelOS(t *testing.T) {
	asus := NewComputer()

	ubuntu := asus.CreateOS("Ubuntu")

	intel := asus.CreateHardware("kingstone 4Gb", "Intel i5", "Asus PG5", "nVideo PGX1800")
	intel.InstallOS(ubuntu)

	want := ubuntu.GetFamily()
	got := intel.GetOSFamily()

	if want != got {
		t.Errorf("want: %s, got: %s\n", want, got)
	}

}

func TestAmdOS(t *testing.T) {
	gigabyte := NewComputer()

	windows := gigabyte.CreateOS("Windows")

	amd := gigabyte.CreateHardware("Gigabite 8Gb", "AMD 2200", "NoName", "ATI r530x")
	amd.InstallOS(windows)

	want := windows.GetFamily()
	got := amd.GetOSFamily()

	if want != got {
		t.Errorf("want: %s, got: %s\n", want, got)
	}
}

func TestWithoutOS(t *testing.T) {
	lenovo := NewComputer()

	whithout := lenovo.CreateOS("")

	intel := lenovo.CreateHardware("kingspec 4Gb", "Intel i3", "MSI B650M", "nVideo PGX1200")
	intel.InstallOS(whithout)

	want := whithout.GetFamily()
	got := intel.GetOSFamily()

	if want != got {
		t.Errorf("want: %s, got: %s\n", want, got)
	}

}
