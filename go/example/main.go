package main

// Example is a simple Go program that prints its version.
import "fmt"

var version string //ビルド時にldflagsフラグ経由でバージョンを埋め込むための変数

func main() {
	fmt.Printf("Example %s\n", version)
}
