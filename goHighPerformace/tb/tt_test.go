/*
Using []byte as a map key
It is very common to use a string as a map key, but often you have a []byte.

The compiler implements a specific optimisation for this case

var m map[string]string
v, ok := m[string(bytes)]
This will avoid the conversion of the byte slice to a string for the map lookup. This is very specific, it won't work if you do something like

key := string(bytes)
val, ok := m[key]

经下面测试 使用v, ok := m[string(bytes)]性能更高
*/
/*
Reduce allocations
Make sure your APIs allow the caller to reduce the amount of garbage generated.
Consider these two Read methods
 func (r *Reader) Read() ([]byte, error)
 func (r *Reader) Read(buf []byte) (int, error)
The Òrst Read method takes no arguments and returns some data as a []byte. The second
takes a []byte bu×er and returns the amount of bytes read.
The Òrst Read method will always allocate a bu×er, putting pressure on the GC. The second
Òlls the bu×er it was given.
*/
package tb

import (
	"strconv"
	"testing"
)

var m = make(map[string]int)
var bb [][]byte

func init() {
	for i := 0; i < 2000; i++ {
		bb = append(bb, []byte(strconv.Itoa(i)))
	}
	for i := 0; i < 2000; i++ {
		s := string(bb[i])
		m[s] = i
	}
}
func MstringOne() {
	for i := 0; i < 2000; i++ {
		_, ok := m[string(bb[i])]
		if ok {

		}
	}
}
func MstringTwo() {
	for i := 0; i < 2000; i++ {
		key := string(bb[i])
		_, ok := m[key]
		if ok {

		}
	}
}
func BenchmarkMapString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MstringOne()
	}
}
func BenchmarkMapStringTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MstringTwo()
	}
}
