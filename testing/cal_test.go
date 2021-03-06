package cal

import (
    "testing"
    "fmt"
)

func Test_Division_1(t *testing.T) {
    if i, e:= Division(6, 2); i != 3 || e != nil {
        t.Error("Fail")
    } else {
        t.Log("pass")
    }
}

func Test_Division_2(t *testing.T) {
    t.Error("Fail")
}

//測試多組案例
func Test_Division_table(t *testing.T) {
	tables := []struct {
		x float64
		y float64
	}{
		{3, 1},
		{6, 2},
		{9, 3},
		{8, 2},
	}

	for _, table := range tables {
		if i, e := Division(table.x, table.y); i != 3 || e != nil {
			t.Error("除法函數沒通過測試")
		} else {
			t.Log("測試通過")
		}
	}
}
//壓力測試
func Benchmark_Division(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Division(4, 5)
	}
}

//示例函數
func ExampleDivision() {
	i, _ := Division(6, 2)
	x, _ := Division(12, 3)
	fmt.Println(i)
	fmt.Println(x)
	// Output:
	// 3
	// 4
}
//test command
//go test -timeout 30s -run ^ExampleDivision
//覆蓋率
//$ go test -cover


