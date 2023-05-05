package spider

import (
	"encoding/json"
	"io"
	"strings"
)

func (s *App) OutputJsonForApi(w io.Writer, v interface{}) error {
	b, err := MarshalJsonForApi(v)
	if err != nil {
		return err
	}

	_, err = w.Write(b)
	return err
}

func MustMarshalJsonForApi(v interface{}) string {
	b, err := MarshalJsonForApi(v)
	if err != nil {
		panic(err)
	}

	return string(b)
}

func MarshalJsonForApi(v interface{}) ([]byte, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	m := map[string]interface{}{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}

	walkMap(m, jsonKeyForApi)
	bs, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return nil, err
	}

	bs = append(bs, '\n')
	return bs, nil
}

func (s *App) UnmarshalJsonForStruct(src []byte, v interface{}, path string) error {
	m := map[string]interface{}{}
	if err := json.Unmarshal(src, &m); err != nil {
		return err
	}

	walkMap(m, jsonKeyForStruct)
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}

	return unmarshalJson(b, v, path)
}

func jsonKeyForApi(s string) string {
	if len(s) == 0 {
		return s
	}

	return strings.ToLower(s[:1]) + s[1:]
}

func jsonKeyForStruct(s string) string {
	if len(s) == 0 {
		return s
	}

	return strings.ToUpper(s[:1]) + s[1:]
}

func walkMap(m map[string]interface{}, fn func(string) string) {
	for key, value := range m {
		delete(m, key)
		newKey := key
		if fn != nil {
			newKey = fn(key)
		}
		if value != nil {
			m[newKey] = value
		}

		switch value := value.(type) {
		case map[string]interface{}:
			walkMap(value, fn)
		case []interface{}:
			if len(value) > 0 {
				walkArray(value, fn)
			} else {
				delete(m, newKey)
			}
		default:
		}
	}
}

func walkArray(a []interface{}, fn func(string) string) {
	for _, value := range a {
		switch value := value.(type) {
		case map[string]interface{}:
			walkMap(value, fn)
		case []interface{}:
			walkArray(value, fn)
		default:
		}
	}
}
