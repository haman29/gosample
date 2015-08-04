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

	// if
	// {} は省略不可、三項演算子はナイ
	a2, b2 := 10, 100
	if a2 > b2 {
		fmt.Println("a2 is lager than b2")
	} else if a2 < b2 {
		fmt.Println("a2 is smaller than b2")
	} else {
		fmt.Println("a2 equals b2")
	}

	// for
	// while, loopはナイ
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 無限ループ
	// for {
		// something()
	// }
	
	// break, continue
	n := 0
	for {
		n++
		if n > 10 {
			break // ループを抜ける
		}
		if n%2 ==0 {
			continue
		}
		fmt.Println(n)
	}

	func1()
	func2()
	func3()
}
func func1() {
	n := 10
	// go の switch は明示的にbreakする必要ナシ (ruby と一緒)
	switch n {
	case 15:
		fmt.Println("FizzBuzz")
	case 5,10:
		fmt.Println("Buzz")
	case 3,6,9:
		fmt.Println("Fizz")
	default:
		fmt.Println(n)
	}
}
func func2() {
	n := 3
	// 逆に次の処理を評価したいときに `fallthrough` を使う(ruby と一緒)
	switch n {
	case 3:
		n = n - 1
		fallthrough
	case 2:
		n = n - 1
		fallthrough
	case 1:
		n = n - 1
		fmt.Println(n) // => 0
	}
}

func func3() {
	n := 10
	// caseには値だけじゃなくて式も指定できる
	switch {
	case n%15 == 0:
		fmt.Println("FizzBuzz")
	case n%5 == 0:
		fmt.Println("Buzz")
	case n%3 == 0:
		fmt.Println("Fizz")
	default:
		fmt.Println(n)
	}
}
