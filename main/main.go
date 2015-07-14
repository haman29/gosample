// main.go
// hello world の説明

// パッケージ
// package main で`main()`関数が呼ばれる
package main

// インポート
// - 使用するパッケージを列挙する
// - 使用していないパッケージがあるとコンパイルエラー
import (
	"fmt"
	"github.com/haman29/gosample"
)

func main() {
	// ドット(`.`)でパッケージ内の関数・変数にアクセスする
	fmt.Println(gosample.Message)
	// => hello world
}
