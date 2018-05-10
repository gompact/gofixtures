package yaml

import (
	"bytes"
	"testing"
)

func TestDBConfParsing(t *testing.T) {
	parser := New()
	data := "driver: postgres\nhost: localhost\ndatabase: postgres\nuser: user\npassword: 123\n"
	b := bytes.NewBufferString(data)
	dbConf, err := parser.ParseDBConf(b)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if dbConf.Host != "localhost" {
		t.Errorf("host was not parsed correctly, expected %s, got %s", "localhost", dbConf.Host)
	}
	if dbConf.Driver != "postgres" {
		t.Errorf("driver was not parsed correctly, expected %s, got %s", "postgres", dbConf.Host)
	}
	if dbConf.Database != "postgres" {
		t.Errorf("database was not parsed correctly, expected %s, got %s", "postgres", dbConf.Host)
	}
	if dbConf.User != "user" {
		t.Errorf("user was not parsed correctly, expected %s, got %s", "user", dbConf.Host)
	}
	if dbConf.Password != "123" {
		t.Errorf("password was not parsed correctly, expected %s, got %s", "123", dbConf.Host)
	}
}
