package params

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

func Pack(ptr interface{}) (url.URL, error) {
	v := reflect.ValueOf(ptr).Elem()
	if v.Type().Kind() != reflect.Struct {
		return url.URL{}, fmt.Errorf("Pack(%v): got %T want struct", ptr, ptr)
	}
	vals := &url.Values{}
	for i := 0; i < v.NumField(); i++ {
		path := v.Type().Field(i)
		tag := path.Tag
		param := tag.Get("http")
		if param == "" {
			param = strings.ToLower(path.Name)
		}
		vals.Add(param, fmt.Sprintf("%v", v.Field(i)))
	}
	return url.URL{RawQuery: vals.Encode()}, nil
}
