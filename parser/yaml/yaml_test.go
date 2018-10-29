package yaml

import (
	"bytes"
	"testing"
)

func TestDBConfParsing(t *testing.T) {
	parser := New()
	data := `
	db:
		driver: postgres
		host: localhost
		database: postgres
		user: user
		password: 123`
	b := bytes.NewBufferString(data)
	conf, err := parser.ParseConfig(b)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if conf.DB.Host != "localhost" {
		t.Errorf("host was not parsed correctly, expected %s, got %s", "localhost", conf.DB.Host)
	}
	if conf.DB.Driver != "postgres" {
		t.Errorf("driver was not parsed correctly, expected %s, got %s", "postgres", conf.DB.Host)
	}
	if conf.DB.Database != "postgres" {
		t.Errorf("database was not parsed correctly, expected %s, got %s", "postgres", conf.DB.Host)
	}
	if conf.DB.User != "user" {
		t.Errorf("user was not parsed correctly, expected %s, got %s", "user", conf.DB.Host)
	}
	if conf.DB.Password != "123" {
		t.Errorf("password was not parsed correctly, expected %s, got %s", "123", conf.DB.Host)
	}
}
