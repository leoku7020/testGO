package def

import "fmt"
//defer 特性 先進後出
func main() {
	fmt.Println("修改函數中的命名返回值")
	fmt.Println(c())
	fmt.Println("宣告後不影響")
	a()
	fmt.Println("先進後出")
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
//defer 特性 修改函數中的命名返回值
func c() (i int) {
    defer func() { i++ }()
    return 1
}
//defer 特性 defer中的變量在宣告的時候就已經被確立了，不會受到宣告以後的影響
func a() {
    i := 0
    defer fmt.Println(i)
    i++
    return
}