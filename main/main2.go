// main2.go
// インポートのオプション

package main

import (
	// f.Println()でアクセスする
	f "fmt"

	// 使用しないパッケージ
	_ "github.com/haman29/gosample"

	// パッケージ名の省略 ex) strings.ToUpper() が ToUpper()
	. "strings"
)

func main() {
	f.Println(ToUpper("hello world"))
	// => HELLO WORLD
}


