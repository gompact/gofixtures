package main

import (
	"testing"
)

func TestReadConfig(t *testing.T) {
	c, err := ReadConfig("testdata/config.yml")
	if err != nil {
		t.Error(err.Error())
		t.Fail()
	}
	if c.DB.Driver != "postgres" {
		t.Error("Wrong db driver value")
		t.Fail()
	}
	if c.DB.Database != "test" {
		t.Error("Wrong database value")
		t.Fail()
	}
}


func TestFilesToParse(t *testing.T) {
	files, err := filesToParse("testdata")
	if err != nil {
		t.Error(err.Error())
		t.Fail()
	}

	if len(files) != 1 {
		t.Error("Wrong number of files to parse")
		t.Fail()
	}
}

func TestHandleCommandLineArguments(t *testing.T) {
	args := []string{"", "version"}

	cmd := handleCommandLineArguments(args)

	if cmd != PrintVersion {
		t.Error("Failed to read cmd arguments correctly")
		t.Fail()
	}

	args = []string{""}
	cmd = handleCommandLineArguments(args)
	if cmd != 0 {
		t.Error("handleCommandLineArguments should return 0 if user didn't enter any commands")
		t.Fail()
	}
}
