package completer

import "testing"

func TestReadCompletersBalooctl6(t *testing.T) {
	completers, err := ReadCompleters("../../../../completers", "linux")
	if err != nil {
		t.Fatal(err)
	}

	if _, ok := completers.Lookup("balooctl6@linux"); !ok {
		t.Fatal("expected balooctl6 linux completer")
	}
}
