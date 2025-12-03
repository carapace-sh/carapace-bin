package completer

import "testing"

func TestSortVariants(t *testing.T) {
	m := CompleterMap{
		"test": {
			{
				Name:    "test",
				Variant: "z",
			},
			{
				Name:    "test",
				Variant: "",
			},
		},
	}
	if m["test"][0].Variant != "z" {
		t.Error("invalid variant")
	}
	m.SortVariants()
	if m["test"][0].Variant != "" {
		t.Error("sorting failed")
	}
}
