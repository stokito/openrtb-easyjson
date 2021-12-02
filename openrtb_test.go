package openrtb

import (
	"testing"
)

func TestDeviceTypeString(t *testing.T) {
	if got := DeviceTypeUnknown.String(); got != "0" {
		t.Errorf("DeviceTypeString() = %v, want %v", got, "0")
	}
	if got := DeviceTypeMobile.String(); got != "1" {
		t.Errorf("DeviceTypeString() = %v, want %v", got, "1")
	}
	if got := DeviceTypePC.String(); got != "2" {
		t.Errorf("DeviceTypeString() = %v, want %v", got, "2")
	}
	if got := DeviceTypeTV.String(); got != "3" {
		t.Errorf("DeviceTypeString() = %v, want %v", got, "3")
	}
	if got := DeviceTypePhone.String(); got != "4" {
		t.Errorf("DeviceTypeString() = %v, want %v", got, "4")
	}
	if got := DeviceTypeTablet.String(); got != "5" {
		t.Errorf("DeviceTypeString() = %v, want %v", got, "5")
	}
	if got := DeviceTypeConnected.String(); got != "6" {
		t.Errorf("DeviceTypeString() = %v, want %v", got, "6")
	}
	if got := DeviceTypeSetTopBox.String(); got != "7" {
		t.Errorf("DeviceTypeString() = %v, want %v", got, "7")
	}
}
