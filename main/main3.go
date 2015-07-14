package main

import (
	"fmt"
)

func main() {
	// 変数
	// - 基本形 : var 変数名 型 = 値
	// - 定義して使用していない変数はコンパイルエラー
	var message string = "hello world"
	fmt.Println(message)
	// => hello world

	// 複数 v1
	var foo, bar, buz string = "foo", "bar", "buz"
	fmt.Println(foo, bar, buz)
	// => foo bar buz

	// 複数 v2
	var (
		a string = "a"
		b = "b"
		c = "c"
	)
	fmt.Println(a, b, c)
	// => a b c

	// 関数内部での宣言と初期化 := (型推論)
	// var message2 string = "hello world" と同等の記述
	message2 := "hello world"
	fmt.Println(message2)
	// => hello world
	
	// 定数
	// - 再代入はコンパイルエラー
	// - 定義して使用してなくてもコンパイルエラーにはならない
	const Hello string = "hello"
	fmt.Println(Hello)
	// Hello = "bye" // cannot assign to Hello
	
	// ゼロ値
	// - 明示的に初期化しなかった場合に設定される値(型による)
	var i int
	fmt.Println(i) // => 0
}
