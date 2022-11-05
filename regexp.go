package ctool

import "regexp"

// IsMatch 判断正则表达式是否匹配
func IsMatch(reg *regexp.Regexp, content string) bool {
	if len(content) == 0 {
		return false
	}
	return reg.MatchString(content)
}

