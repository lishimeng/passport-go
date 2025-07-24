package sdk

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func printStruct(v interface{}) string {
	rv := reflect.ValueOf(v)
	rt := reflect.TypeOf(v)

	// 如果是指针，解引用
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
		rt = rt.Elem()
	}

	// 必须是结构体
	if rv.Kind() != reflect.Struct {
		return fmt.Sprintf("not a struct: %T", v)
	}

	var lines []string
	lines = append(lines, fmt.Sprintf("struct %s {", rt.Name()))

	for i := 0; i < rv.NumField(); i++ {
		field := rt.Field(i)
		value := rv.Field(i)

		// 跳过未导出字段（首字母小写）
		if !value.CanInterface() {
			continue
		}

		// 格式：FieldName: value
		line := fmt.Sprintf("  %s: %v", field.Name, value.Interface())
		lines = append(lines, line)
	}

	lines = append(lines, "}")
	return strings.Join(lines, "\n")
}

func TestPasswordAuth(t *testing.T) {
	resp, err := New(
		WithAuth("c75834b904c0d3c39cdfa25bd0919ac75ebf8d8a9c2a828795cfd08ba29ba009",
			"5efefff775513dfe7713f51ef71a652e0968d342960454f2478f5adb894107b8"),
		WithHost("http://127.0.0.1:80"),
	).AccessToken(WithPassword("admin", "pwd"))
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(printStruct(resp))
}

func TestURL(t *testing.T) {
	url, err := New(
		WithAuth("c75834b904c0d3c39cdfa25bd0919ac75ebf8d8a9c2a828795cfd08ba29ba009",
			"5efefff775513dfe7713f51ef71a652e0968d342960454f2478f5adb894107b8"),
		WithHost("http://127.0.0.1:80"),
	).Authorize("http://127.0.0.1:82/api/user/authCode", "read,passport")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(url)
}

func TestCodeAuth(t *testing.T) {
	resp, err := New(
		WithAuth("c75834b904c0d3c39cdfa25bd0919ac75ebf8d8a9c2a828795cfd08ba29ba009",
			"5efefff775513dfe7713f51ef71a652e0968d342960454f2478f5adb894107b8"),
		WithHost("http://127.0.0.1:80"),
	).AccessToken(WithCode("5MOBQanEOqCCv5H8Oi9wXyLDTAn9bGUUaGcJNdOqiuQ"))
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(printStruct(resp))
}

func TestRefreshToken(t *testing.T) {
	resp, err := New(
		WithAuth("c75834b904c0d3c39cdfa25bd0919ac75ebf8d8a9c2a828795cfd08ba29ba009",
			"5efefff775513dfe7713f51ef71a652e0968d342960454f2478f5adb894107b8"),
		WithHost("http://127.0.0.1:80"),
	).AccessToken(WithRefreshToken("-9YHacEMCDQM71Kw0ZVav_p7BAWUpIT3qJFwnTOK-Ns"))
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(printStruct(resp))
}
