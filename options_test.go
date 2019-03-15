package cfg

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"testing"
)

func TestOptionsNoFiles(t *testing.T) {
	c := New("none")
	s := c.String("s", "default", "")
	Parse()
	if *s != "default" {
		t.Error("Expected default from no files...")
	}
}

func TestOptionsHomeFiles(t *testing.T) {
	if user, err := user.Current(); err == nil {
		ioutil.WriteFile(fmt.Sprintf("%s/.cfgflagpackagetestdata.json", user.HomeDir), []byte("{\"test\": \"value\"}"), 0644)
		ioutil.WriteFile(fmt.Sprintf("%s/.cfgflagpackagetestdata.yml", user.HomeDir), []byte("{\"test2\": \"value\"}"), 0644)
		c := New("cfgflagpackagetestdata")
		s := c.String("test", "default", "")
		s2 := c.String("test2", "default", "")
		Parse()
		if *s != "value" {
			t.Error("Expected default from home dir files...")
		}
		if *s2 != "value" {
			t.Error("Expected default from home dir files...")
		}
	}
}

func TestOptionsBadFiles(t *testing.T) {
	ioutil.WriteFile(".bad.json", []byte("{\n"), 0644)
	ioutil.WriteFile(".bad.yml", []byte("-blah"), 0644)
	c := New("bad")
	s := c.String("sbad", "default", "")
	Parse()
	if *s != "default" {
		t.Error("Expected default from no files...")
	}
}

func TestEnvFile(t *testing.T) {
	filePath := "/tmp/mysecret.yml"
	ioutil.WriteFile(filePath, []byte("value"), 0644)
	defer os.Remove(filePath)
	os.Setenv("MY_VAR_FILE", filePath)
	defer os.Unsetenv("MY_VAR_FILE")

	c := New("my")
	s := c.String("var", "default","")

	if *s != "value" {
		t.Error("Expected value from env var file path got default")
	}

}

func TestEnv(t *testing.T) {
	os.Setenv("MY_VARS", "new_value")
	defer os.Unsetenv("MY_VARS")

	d := New("my")
	x := d.String("vars", "default","")

	if *x != "new_value" {
		t.Errorf("Expected value from env var got default")
	}
}
