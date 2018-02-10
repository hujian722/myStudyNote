package main

import (
	"bytes"
)

/*
https://golangbot.com/arrays-and-slices/
如果slice很大  我们只对部分数据有兴趣  那么用copy
*/

//========================================================
//========================================================

转换成string  
use fmt.Sprint(i) (Slow)
use strconv.Itoa(int(i)) (Fast)
use strconv.FormatInt(int64(i), 10) (Faster)
也可以自己写一个转换函数

//========================================================
//========================================================

//byte rune
func main() {
	var a byte = 2
	var b uint8 = 2
	var c int32 = 2
	var d rune = 2
	fmt.Printf("%T %T %T %T:\n", a, b, c, d)
	//out: uint8 uint8 int32 int32
}


//========================================================
//========================================================

go test -bench=. -benchmem -benchtime=20s
go test -test.bench=".*" -count=5

trace.Start(os.Stdout)
	defer trace.Stop()
go tool trace px.trace

Each benchmark is run for a minimum of 1 second by default.
每次性能测试至少运行1s，如果时间没到1s，b.N 增长 1, 2, 5, 10, 20, 50, …
BenchmarkFib40        50         944501481 ns/op
					最后一次的b.N  每次平均时间
is the average run time of the function under test for the final value of b.N iterations

1. 小对象组合成大对象
2. bytes.Buffer 
3. 预先分配map slice大小
4. 避免频繁创建临时对象 临时大对象
5. 高并发用池 sync.pool
6. []byte和string转换
7. 拼接字符串 +  。。。
6. 减少不必要的指针引用
7. 当一个对象不包含任何指针（注意：strings，slices，maps 和chans包含隐含的指针），时，对gc的扫描影响很小。 比如，1GB byte 的slice事实上只包含有限的几个object，不会影响垃圾收集时间。 因此，我们可以尽可能的减少指针的引用。
8. 不需要大类型  const int8
//========================================================
//========================================================

cd "my project in GOPATH"
govendor init
# Add existing GOPATH files to vendor.
govendor add +external
# View your work.
govendor list
# Look at what is using a package
govendor list -v fmt

//========================================================
//========================================================				
清空slice
fruitslice[:0]
fruitslice[:0:0]
//输出数组slice对应的数组地址
var fruitslice=[]int{1,2,3}
fmt.Printf("%p \n", &fruitslice)  输出变量地址
fmt.Printf("%p \n", fruitslice)
fmt.Printf("%p \n", &fruitslice[0])

Arrays are value types
func main() {  
    a := [...]string{"USA", "China", "India", "Germany", "France"}
    b := a // a copy of a is assigned to b
    b[0] = "Singapore"
    fmt.Println("a is ", a)
    fmt.Println("b is ", b) 
}
//========================================================
//========================================================	
golang defer精析
return xxx会被改写成:
返回值 = xxx
调用defer函数
空的return

func f() (result int) { 
    defer func() { 
        result++ 
    }() 
    return 0
}

func f() (r int) { 
    t := 5 
    defer func() { 
        t = t + 5 
    }() 
    return t
}

func f() (r int) { 
    defer func(r int) { 
        r = r + 5 
    }(r) 
    return 1
}