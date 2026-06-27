package fs

import (
	"bytes"
	"testing"

	_ "embed"

	"github.com/micromdm/plist"
)

//go:embed blockdevice-diskutil.plist
var blockdeviceDiskutil []byte

func TestBlockDevicesDiskutil(t *testing.T) {
	var parsed blockdevicesDiskutil
	err := plist.NewXMLDecoder(bytes.NewReader(blockdeviceDiskutil)).Decode(&parsed)
	if err != nil {
		t.Error(err)
	}

	devices := parsed.toBlockdevices()
	if len(devices) == 0 {
		t.Error("expected blockdevices to be non-empty")
	}

	expectedFirst := blockdevice{
		Kname:        "disk0",
		Path:         "/dev/disk0",
		Size:         "113G",
		Type:         "disk",
		Parttypename: "GUID_partition_scheme",
	}
	if devices[0] != expectedFirst {
		t.Errorf("expected %v, got %v", expectedFirst, devices[0])
	}
}

func TestFormatBlockSize(t *testing.T) {
	tests := []struct {
		bytes    int
		expected string
	}{
		{500, "500B"},
		{1024, "1K"},
		{1048576, "1M"},
		{1073741824, "1G"},
		{1099511627776, "1T"},
		{121332826112, "113G"},
		{209715200, "200M"},
	}
	for _, tt := range tests {
		got := formatBlockSize(tt.bytes)
		if got != tt.expected {
			t.Errorf("formatBlockSize(%d) = %q, want %q", tt.bytes, got, tt.expected)
		}
	}
}
