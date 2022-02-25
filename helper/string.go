package helper

import "strings"

// 下划线写法转为驼峰写法
func TranCamel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

// 首字母大写
func UpperFirst(name string) string{
	return strings.ToUpper(name[:1]) + name[1:]
}

// 首字母小写
func LowerFirst(name string) string{
	return strings.ToLower(name[:1]) + name[1:]
}
