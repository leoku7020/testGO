package main

import "fmt"

func main() {
    var i *int // 宣告i 是一個int的指標
    a := 10 //a 佔用了一個記憶體位置
    i = &a //將i 指到a的記憶體位置
    fmt.Println(i) // i 所指到的記憶體位置
    fmt.Println(*i) // *代表顯示該記憶體位置的值
    //動態配置記憶體
    //宣告一個空的int
    n := new(int)
    // 直接把值指向記憶體位置
    *n = 2
    fmt.Println(n)
    fmt.Println(*n)
    foo_value(*n)
    foo_point(n)
}

func foo_value(x int) {
    fmt.Println(&x) //function 內ｘ的記憶體位置
}

func foo_point(x *int) {
    fmt.Println(x) //function 內 x 的記憶體位置
}

/**
對struct內的成員進行修改時，使用指標
如果struct內有大量的成員，請使用指標節省記憶體空間
為了一致性，如果團隊內有時使用指標有時使用值會造成混亂，建議一致使用指標
**/
