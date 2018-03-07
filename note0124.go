package main

import (
	"bytes"
)

/*
https://golangbot.com/arrays-and-slices/
如果slice很大  我们只对部分数据有兴趣  那么用copy
var aa = []*A{a0, a1, a2, a3}
bb := aa[:len(aa)-1]
此时我们只用a0 a1 a2 但是 a3不会被回收
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

. concat string in a loop, bytes.Buffer is best choose, don’s use += or x = x + y
. direct concat fixed number param and param type is confirmed, direct use +
. fixed number but complex to find param type, see fmt.Sprintf or fmt.Fprintf
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

//========================================================
//========================================================	
zero values  类型的0值
bool    → false    pointers → nil  
numbers → 0        slices → nil    
string  → ""       maps → nil      
                   channels → nil  
                   functions → nil 
                   interfaces → nil

zero values for struct types

type Person struct {
AgeYears int
Name string
Friend []Person
}
var p Person // Person{0, "", nil}

nil is a predeclared identifier representing the zero value for
a pointer,channel, func, interface,map, or slice type

var s fmt.Stringer // Stringer (nil, nil)
fmt.Println(s == nil) // true
var p *Person // nil of type *Person
var s fmt.Stringer = p // Stringer (*Person, nil)
fmt.Println(s == nil) // false  此时interface里面已经有了type 所以不为nil

type doError struct {
	err string
}

func (d *doError) Error() string {
	return d.err
}

func (d *doError) String() string {
	return d.err
}
func do() *doError {
	return nil
}
func wrapDo() error {
	return do()
}

func main() {
	err := wrapDo()
	fmt.Println(err == nil) //false
	var d *doError
	var s fmt.Stringer = d
	fmt.Println(s == nil) //false
}
//========================================================
//========================================================	

for v = range aChannel {
	// use v
}
is equivalent to
for {
	v, ok = <-aChannel
	if !ok {
		break
	}
	// use v
}
	     				 Nil Channel	  		 Closed Channel	   Active Channel
Close	 					 panic	             	panic			succeed to close
Send Value To			block for ever				panic			block or succeed to send
Receive Value From		block for ever		      never block	    block or succeed to receive
//========================================================
//========================================================	
All structs in Golang are of the same kind, but not the same type

//========================================================
//========================================================	

那么 Pool 都适用于什么场景呢？从它的特点来说，适用与无状态的对象的复用，而不适用与如连接池之类的。在 fmt 包中有一个很好的使用池的例子，它维护一个动态大小的临时输出缓冲区。

官方例子：
package main

import (
    "bytes"
    "io"
    "os"
    "sync"
    "time"
)

var bufPool = sync.Pool{
    New: func() interface{} {
        return new(bytes.Buffer)
    },
}

func timeNow() time.Time {
    return time.Unix(1136214245, 0)
}

func Log(w io.Writer, key, val string) {
    // 获取临时对象，没有的话会自动创建
    b := bufPool.Get().(*bytes.Buffer)
    b.Reset()
    b.WriteString(timeNow().UTC().Format(time.RFC3339))
    b.WriteByte(' ')
    b.WriteString(key)
    b.WriteByte('=')
    b.WriteString(val)
    w.Write(b.Bytes())
    // 将临时对象放回到 Pool 中
    bufPool.Put(b)
}

func main() {
    Log(os.Stdout, "path", "/search?q=flowers")
}

打印结果：
2006-01-02T15:04:05Z path=/search?q=flowers

//========================================================
//========================================================	
监控gc ====》》》  set GODEBUG=gctrace=1
逃逸分析 go build -gcflags "-m -l"
//========================================================
//========================================================	

// Allocate an object of size bytes.
// Small objects are allocated from the per-P cache's free lists.
// Large objects (> 32 kB) are allocated straight from the heap.
func mallocgc(size uintptr, typ *_type, needzero bool) unsafe.Pointer {}
如果堆上分配大于32k  将触发gc
//========================================================
//========================================================	
Insert

s = append(s, 0)
copy(s[i+1:], s[i:])
s[i] = x

//========================================================
//========================================================	


//========================================================
//========================================================	



//========================================================
//========================================================	


//========================================================
//========================================================	


//========================================================
//========================================================	


//========================================================
//========================================================	


//========================================================
//========================================================	



//========================================================
//========================================================	


//========================================================
//========================================================	
//========================================================
//========================================================	


//========================================================
//========================================================	


//========================================================
//========================================================	



//========================================================
//========================================================	


//========================================================
//========================================================	


//========================================================
//========================================================	


//========================================================
//========================================================	


//========================================================
//========================================================	



//========================================================
//========================================================	


//========================================================
//========================================================	


//========================================================
//========================================================	


//========================================================
//========================================================	


//========================================================
//========================================================	



//========================================================
//========================================================	


//========================================================
//========================================================	