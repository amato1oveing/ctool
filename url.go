package ctool

import "net/url"

//UrlEncode url编码
func UrlEncode(str string) string {
	return url.QueryEscape(str)
}

//UrlDecode url解码
func UrlDecode(str string) (string, error) {
	return url.QueryUnescape(str)
}

//UrlEncodeMap url编码map
func UrlEncodeMap(m map[string]string) string {
	v := url.Values{}
	for k, val := range m {
		v.Set(k, val)
	}
	return v.Encode()
}

//UrlDecodeMap url解码map
func UrlDecodeMap(str string) (map[string]string, error) {
	m := make(map[string]string)
	v, err := url.ParseQuery(str)
	if err != nil {
		return m, err
	}
	for k, val := range v {
		m[k] = val[0]
	}
	return m, nil
}

//UrlEncodeMap2 url编码map
func UrlEncodeMap2(m map[string]interface{}) string {
	v := url.Values{}
	for k, val := range m {
		v.Set(k, ToString(val))
	}
	return v.Encode()
}

//UrlDecodeMap2 url解码map
func UrlDecodeMap2(str string) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	v, err := url.ParseQuery(str)
	if err != nil {
		return m, err
	}
	for k, val := range v {
		m[k] = val[0]
	}
	return m, nil
}
