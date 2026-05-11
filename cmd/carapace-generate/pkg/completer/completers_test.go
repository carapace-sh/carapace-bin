package completer

import "testing"

func TestReadCompletersKcl(t *testing.T) {
	completers, err := ReadCompleters("../../../../completers", "linux")
	if err != nil {
		t.Fatal(err)
	}

	if _, ok := completers.Lookup("kcl@common"); !ok {
		t.Fatal("expected kcl common completer")
	}
}
