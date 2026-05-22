package fs

import (
	"reflect"
	"testing"
)

func TestParseProcMounts(t *testing.T) {
	input := "/dev/sda1 /boot/efi vfat rw 0 0\n" +
		"tmpfs /run/user/1000/gvfs\\040mount fuse rw 0 0\n"

	want := []mountEntry{
		{source: "/dev/sda1", target: "/boot/efi"},
		{source: "tmpfs", target: "/run/user/1000/gvfs mount"},
	}
	if got := parseProcMounts(input); !reflect.DeepEqual(got, want) {
		t.Fatalf("parseProcMounts() = %#v, want %#v", got, want)
	}
}

func TestParseMountOutput(t *testing.T) {
	input := "/dev/disk3s1s1 on / (apfs, sealed, local, read-only, journaled)\n" +
		"map auto_home on /System/Volumes/Data/home (autofs, automounted, nobrowse)\n" +
		"not a mount line\n"

	want := []mountEntry{
		{source: "/dev/disk3s1s1", target: "/"},
		{source: "map auto_home", target: "/System/Volumes/Data/home"},
	}
	if got := parseMountOutput(input); !reflect.DeepEqual(got, want) {
		t.Fatalf("parseMountOutput() = %#v, want %#v", got, want)
	}
}

func TestDarwinBlockdevicesFromNames(t *testing.T) {
	names := []string{"disk0", "disk0s1", "rdisk0", "diskarbitrationd", "disk3s1s1"}

	want := []blockdevice{
		{Kname: "disk0", Path: "/dev/disk0"},
		{Kname: "disk0s1", Path: "/dev/disk0s1"},
		{Kname: "disk3s1s1", Path: "/dev/disk3s1s1"},
	}
	if got := darwinBlockdevicesFromNames(names); !reflect.DeepEqual(got, want) {
		t.Fatalf("darwinBlockdevicesFromNames() = %#v, want %#v", got, want)
	}
}
