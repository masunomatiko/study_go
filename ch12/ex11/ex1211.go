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
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		vals.Add(name, fmt.Sprintf("%v", v.Field(i)))
	}
	return url.URL{RawQuery: vals.Encode()}, nil
}
