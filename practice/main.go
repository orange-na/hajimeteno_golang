package main // メインプログラムは必ずこうする
import (     // 読み込むパッケージ(ライブラリ)の宣言
	"fmt" // 入出力関連のパッケージ 。Scanln と Printlnを使う "strconv" // 文字列から整数などへの変換。Atoiを使う
) // 複数のパッケージを呼び込む場合は、(...)で指定すると「import」は1個で済む
func main() {
	var s int
	if s == 0 {
		s = 1_4_4
		fmt.Printf("s is %d\n", s)
	} else {
		fmt.Println("s is not empty")
	}
}