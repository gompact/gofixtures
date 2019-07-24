package file

import (
	"testing"
)

func TestRead(t *testing.T) {
	f := New()
	files := []string{"testdata/fixture.yaml"}
	inputs, err := f.Read(files)
	if err != nil {
		t.Error(err.Error())
		t.Fail()
	}

	if inputs[0].Type != ".yaml" {
		t.Errorf("Expected input type to be yaml, found %s\n", inputs[0].Type)
		t.Fail()
	}
}
