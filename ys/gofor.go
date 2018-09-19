package main

import "fmt"

/*
对数组，map的循环，循环打印10个数  跟JAVAcontine,break,return 对比
*/
func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			s := i * j
			fmt.Printf("%d x %d = %d  ", j, i, s)
		}
		fmt.Print("\n")
	}
	var lalala map[string]string
	lalala = make(map[string]string)
	lalala["1"] = "12"
	lalala["2"] = "13"

	for arr := range lalala {
		fmt.Println(lalala[arr])
	}
}
