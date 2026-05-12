package eopkg

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadPackageEntries(t *testing.T) {
	path := filepath.Join(t.TempDir(), "eopkg-index.xml")
	content := []byte(`<PISI>
  <Package>
    <Name>curl</Name>
    <Summary> command line tool for transferring data </Summary>
    <PartOf>network.clients</PartOf>
  </Package>
  <Package>
    <Name>zlib</Name>
    <Description>
      compression library
    </Description>
    <PartOf>system.base</PartOf>
  </Package>
</PISI>`)

	if err := os.WriteFile(path, content, 0o644); err != nil {
		t.Fatal(err)
	}

	entries := readPackageEntries(path)
	if len(entries) != 2 {
		t.Fatalf("expected 2 package entries, got %d", len(entries))
	}

	if entries[0].Name != "curl" || firstNonEmpty(entries[0].Summary) != "command line tool for transferring data" || entries[0].PartOf != "network.clients" {
		t.Fatalf("unexpected first package entry: %#v", entries[0])
	}
	if entries[1].Name != "zlib" || firstNonEmpty(entries[1].Summary, entries[1].Description) != "compression library" || entries[1].PartOf != "system.base" {
		t.Fatalf("unexpected second package entry: %#v", entries[1])
	}
}
