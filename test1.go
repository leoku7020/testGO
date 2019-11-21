//宣告程式屬於哪個package
package main

//引入套件
import (
    "fmt"
)
//常數宣告
const ip string = "127.0.0.1"
var ip2 string = ""
//主程式
func main(){
    //使用:= 簡化變數宣告 var word string = "Hello World!!"
    word := "Hello World!!"
    //使用fmt 套件印出字串word
    fmt.Println(word)
    fmt.Println("MyIp:"+ip)
    //change my ip const can't change | var can change
    ip2 = "192.168.0.1"
    fmt.Println("change ip2 :"+ip2)
}
