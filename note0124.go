package main

import (
	"path/filepath"
	"time"
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

当defer被声明时，其参数就会被实时解析
func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
	}
	上面我们说过，defer函数会在return之后被调用。那么这段函数执行完之后，是不用应该输出1呢？
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
    b.WriteByte(, ,)
    b.WriteString(key)
    b.WriteByte(,=,)
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
// Small objects are allocated from the per-P cache,s free lists.
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
(pprof) top10
Total: 2525 samples
     298  11.8%  11.8%      345  13.7% runtime.mapaccess1_fast64
     268  10.6%  22.4%     2124  84.1% main.FindLoops
     251   9.9%  32.4%      451  17.9% scanblock
     178   7.0%  39.4%      351  13.9% hash_insert
     131   5.2%  44.6%      158   6.3% sweepspan
     119   4.7%  49.3%      350  13.9% main.DFS
      96   3.8%  53.1%       98   3.9% flushptrbuf
      95   3.8%  56.9%       95   3.8% runtime.aeshash64
      95   3.8%  60.6%      101   4.0% runtime.settype_flush
      88   3.5%  64.1%      988  39.1% runtime.mallocgc
When CPU profiling is enabled, the Go program stops about 100 times per second and records a sample 
consisting of the program counters on the currently executing goroutine,s stack. The profile has 2525 samples, 
so it was running for a bit over 25 seconds. In the `go tool pprof` output, there is a row for each function that 
appeared in a sample. The first two columns show the number of samples in which the function was running 
(as opposed to waiting for a called function to return), as a raw count and as a percentage of total samples. 
The runtime.mapaccess1_fast64 function was running during 298 samples, or 11.8%. The top10 output is sorted by this sample count. 
The third column shows the running total during the listing: the first three rows account for 32.4% of the samples. 
The fourth and fifth columns show the number of samples in which the function appeared (either running or waiting for a called function to return). 
The main.FindLoops function was running in 10.6% of the samples, 
but it was on the call stack (it or functions it called were running) in 84.1% of the samples.

To sort by the fourth and fifth columns, use the -cum (for cumulative) flag:

(pprof) top5 -cum
Total: 2525 samples
       0   0.0%   0.0%     2144  84.9% gosched0
       0   0.0%   0.0%     2144  84.9% main.main
       0   0.0%   0.0%     2144  84.9% runtime.main
       0   0.0%   0.0%     2124  84.1% main.FindHavlakLoops
     268  10.6%  10.6%     2124  84.1% main.FindLoops
(pprof) top5 -cum
In fact the total for main.FindLoops and main.main should have been 100%, but each stack sample only includes the bottom 100 stack frames; during about a quarter of the samples, the recursive main.DFS function was more than 100 frames deeper than main.main so the complete trace was truncated.

//========================================================
//========================================================	

BenchmarkMatchString-4            100,000  17,380 ns/op  42,752 B/op  70 allocs/op
BenchmarkMatchStringCompiled-4  2,000,000     843 ns/op       0 B/op   0 allocs/op

BenchmarkConcatString-4    10,000,000  159 ns/op  530 B/op  0 allocs/op
BenchmarkConcatBuffer-4   200,000,000   10 ns/op    2 B/op  0 allocs/op
BenchmarkConcatBuilder-4  100,000,000   11 ns/op    2 B/op  0 allocs/op

//========================================================
//========================================================	
return xxx会被改写成:

返回值 = xxx

调用defer函数

空的return

//========================================================
//========================================================	

func find(num int, nums ...int)  
find(89, []int{nums}) 
func main() {  
    nums := []int{89, 90, 95}
    find(89, nums)
	 find(89, nums...)
}
//========================================================
//========================================================	
wrap errors with http://github.com/pkg/errors
so: errors.Wrap(err, “additional message to a given error”)

implement Stringer interface for integers const values
https://godoc.org/golang.org/x/tools/cmd/stringer

be careful with range in Go:
for i := range a and for i, v := range &a doesn,t make a copy of a
but for i, v := range a does
func main() {
	v := make([]int, 4, 10)

	for i := range v {
		v = append(v, i+10)
	}
	log.Println("over", v)
}
this rule: if you range over an array (or pointer to) and you only assign the index: then only len(a) is evaluated. 
more: https://play.golang.org/p/4b181zkB1O

don,t forget to stop ticker, unless you need a leaked channel
ticker := time.NewTicker(1 * time.Second)
defer ticker.Stop()
//========================================================
//========================================================	
http.Get()  避免使用  这个不会超时
golang json   不需要的字段导出   `json:"-"`
实现自定义json
type Month struct {
    MonthNumber int
    YearNumber int
}

func (m Month) MarshalJSON() ([]byte, error){
    return []byte(fmt.Sprintf("%d/%d", m.MonthNumber, m.YearNumber)), nil
}

func (m *Month) UnmarshalJSON(value []byte) error {
    parts := strings.Split(string(value), "/")
    m.MonthNumber = strconv.ParseInt(parts[0], 10, 32)
    m.YearNumber = strconv.ParseInt(parts[1], 10, 32)

    return nil
}
//========================================================
//========================================================	
字节对齐
type T1 struct {
	a int8
	// To make b 8-aligned on AMD64 OS and 4-aligned on i386 OS,
	// 7 bytes padded on AMD64 OS and pad 3 bytes padded on i386 OS here.
	b int64
	c int16
	// To make the size of T1 values is a multiple of the alignment of T1,
	// 6 bytes padded on AMD64 OS and pad 2 bytes padded on i386 OS here.
}

// the sizes of T1 values are 24 on AMD64 OS and 16 on i386 OS.

type T2 struct {
	a int8
	// To make c 2-aligned,
	// 1 byte padded on both AMD64 and i386 OS here.
	c int16
	// To make b 8-aligned on AMD64 OS and 4-aligned on i386 OS,
	// 4 bytes padded on AMD64 OS here. No padding on i386 OS.
	b int64
}

//========================================================
//========================================================	
func main() {
	time.Sleep(time.Second)
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	for {
		select {
		case <-time.After(3 * time.Second):
		case <-s:
			fmt.Println(" ctrl + c")
			return
		default:
		}
	}
}
time.After会产生内存问题  会释放  但是释放必须等时间到了才行
*** best 主动释放
package main

import "time"

func main() {
    for {
        t := time.NewTimer(3*time.Second)

        select {
        case <- t.C:
        default:
            t.Stop()
        }
    }
}
//========================================================
//========================================================	
http://localhost:6060/debug/pprof/
http://localhost:6060/debug/pprof/profile
http://localhost:6060/debug/pprof/trace?seconds=5
http://localhost:6060/debug/pprof/heap
go tool pprof --inuse_space http://localhost:6060/debug/pprof/heap
go tool pprof --alloc_space http://localhost:6060/debug/pprof/heap

用--inuse_space来分析程序常驻内存的占用情况;
用--alloc_space来分析内存的临时分配情况，可以提高程序的运行速度。
//========================================================
//========================================================	
GC 算法：引用计数（reference counting）、
标记-清扫（mark & sweep）、三色标记算法 白色灰色黑色
节点复制（Copying Garbage Collection），
分代收集（Generational Garbage Collection）。

//========================================================
//========================================================	
Why are map elements unaddressable?
The first reason is that, in Go, for a map m and a key value k, the read operation m[k] is always valid, even if m is a nil map or m doesn,t contain the key k. For the two circumstances, m[k] is always the zero value of the map element type. The zero value of any type should be immutable. Immutable values can,t be taken addresses, otherwise their values may be modified. Whether or not m[k] is a zero value can only be detected at run time. So, to make it safe and consistent, Go compilers always think m[k] is unaddressable.

Another reason is making map elements addressable means the address of a map element mustn,t be changed in its life cycle. This prevents Go compilers using more efficient algorithms to implement map. For the official Go compiler, the internal addresses of map elements may be changed at run time.

//========================================================
//========================================================	

What does the compile error message non-name *** on left side of := mean?
Up to now (Go 1.8), there is a mandatory rule for short variable declarations:

All items at the left side of := must be pure identifiers and at least one of them must be a new variable name.
This means container elements (x[i]), struct fields (x.f), and pointer dereferences etc, can,t appear at the left side of :=.
Currently, there is an open issue (which was merged with a more related one) for this problem. It looks Go authors want to leave this problem unresolved until Go 2.0.

var arry =[]int{1,2,3,4}
	arry[0],i:=7,8
	log.Println(arry)
	log.Println(k,i)
编译报错：non-name arry[0] on left side of :=
//========================================================
//========================================================	
在默认情况下，Go语言的运行时系统会以100 Hz的的频率对CPU使用情况进行取样。
也就是说每秒取样100次，即每10毫秒会取样一次。为什么使用这个频率呢？因为100 Hz既足够产生有用的数据，
又不至于让系统产生停顿。并且100这个数上也很容易做换算，比如把总取样计数换算为每秒的取样数。
实际上，这里所说的对CPU使用情况的取样就是对当前的Goroutine的堆栈上的程序计数器的取样。
由此，我们就可以从样本记录中分析出哪些代码是计算时间最长或者说最耗CPU资源的部分了。

//========================================================
//========================================================	
go 语言没有构造函数一说，所以一般会定义NewXXX函数来初始化相关类。
 NewXXX 函数返回接口时就是简单工厂模式，也就是说Golang的一般推荐做法就是简单工厂。
//API is interface
type API interface {
	Say(name string) string
}

//NewAPI return Api instance by type
func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &helloAPI{}
	}
	return nil
}
//========================================================
//========================================================	
var s []string
s1 := []string{}
log.Println(s == nil, s1 == nil)
//true false

//========================================================
//========================================================	
path  只处理uinix '/'
path/filepath 会对不同平台处理
//========================================================
//========================================================	

首先 interface 是一种类型，从它的定义可以看出来用了 type 关键字，
更准确的说 interface 是一种具有一组方法的类型，这些方法定义了 interface 的行为。

//========================================================
//========================================================	
一个指针变量可以指向任何一个值的内存地址 它指向那个值的内存地址，
在 32 位机器上占用 4 个字节，在 64 位机器上占用 8 个字节，并且与它所指向的值的大小无关。

//========================================================
//========================================================	
空结构体  
cs := make(chan struct{})
var a struct{}
go func() {
	cs <- a  //此处不能  cs <- struct{}  编译报错
}()
rev := <-cs
log.Println(rev)
type A struct {
	a string
}
//========    
    var s struct{}
	fmt.Printf("%p \n", &s)
	//
	a := make([]struct{}, 10)
	b := make([]struct{}, 20)
	fmt.Printf("%p \n", a)
	fmt.Printf("%p \n", b)
	fmt.Println(&a[0] == &b[1])
	// OUTPUT:
	// 0x547fa8
	// 0x547fa8
	// 0x547fa8
	// true
//========================================================
//========================================================	
为了实现多值返回，Go是使用栈空间来返回值的。而常见的C语言是通过寄存器来返回值的。

//========================================================
//========================================================	
字符串  --ptr--len  sizeof==  64位  16字节 每个8字节

//========================================================
//========================================================	
package main

import (
    "fmt"
    "unsafe"
)
func main() {
    var n int64 = 5
    var pn = &n
    var pf = (*float64)(unsafe.Pointer(pn))
    // now, pn and pf are pointing at the same memory address
    fmt.Println(*pf) // 2.5e-323
    *pf = 3.14159
	fmt.Println(n) // 4614256650576692846
	//========
	{
		a := [4]int{0, 1, 2, 3}
		p1 := unsafe.Pointer(&a[1])
		p3 := unsafe.Pointer(uintptr(p1) + 2 * unsafe.Sizeof(a[0]))
		*(*int)(p3) = 6
		fmt.Println("a =", a) // a = [0 1 2 6]
	}
}


//========================================================
//========================================================	
// 字符串值是不能更改的  但是指向能更改
var s = "1234"
up := (*reflect.StringHeader)(unsafe.Pointer(&s))
p := (*string)(unsafe.Pointer(&s))
*p = "789099999"
//此种改变了s的ptr指向  能改变s的值
var s = "1234"
up := (*reflect.StringHeader)(unsafe.Pointer(&s))
p := *(*string)(unsafe.Pointer(&s))
p = "789099999"
//此种没改变了s的ptr指向  不能改变s的值  因为p就是一个新定义的字符串
//========================================================
//========================================================	

golang   bits包
golang   rate包
golang http 持久连接示例
func PrintLocalDial(network, addr string) (net.Conn, error) {
	dial := net.Dialer{
		Timeout:   3 * time.Second,
		KeepAlive: 3 * time.Second,
	}

	conn, err := dial.Dial(network, addr)
	if err != nil {
		return conn, err
	}
	fmt.Println("connect done, use", conn.LocalAddr().String())

	return conn, err
}