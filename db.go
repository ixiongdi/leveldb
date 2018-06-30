package main

// 声明一个map
var m map[string]string

// 初始化map
func init()  {
	m = make(map[string]string)
}
// 新建，修改
func Put(key string, value string) {
	m[key] = value
}
// 删除
func Delete(key string)  {
	delete(m, key)
}
// 读取
func Get(key string) string {
	return m[key]
}

