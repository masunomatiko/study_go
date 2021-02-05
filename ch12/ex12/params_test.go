package params

import (
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"testing"
)

func email(v interface{}) bool {
	return check_regexp(``, v.(string))
}

func check_regexp(reg, str string) bool {
	return regexp.MustCompile(reg).Match([]byte(str))
}

func TestUnpack(t *testing.T) {
	type Params struct {
		post string `http:e,check:email`
	}
	tests := []struct {
		input  *http.Request
		expect Params
	}{
		{
			&http.Request{Form: url.Values{"p": []string{"120"}}},
			Params{"example.email.com"},
		},
		{
			&http.Request{Form: url.Values{"p": []string{"80"}}},
			Params{"example.email.com"},
		},
	}
	validation := map[string]Validation{
		"email": email,
	}
	for _, test := range tests {
		var actual Params
		err := Unpack(test.input, &actual, validation)
		if err != nil {
			t.Errorf("Unpack(%v): %s", test.input, err)
		}
		if reflect.DeepEqual(test.expect, actual) {
			t.Errorf("Unpack(%v), got %v, want %v", test.input, actual, test.expect)
		}
	}
}
