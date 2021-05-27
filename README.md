# golang 
## 编写测试程序
- 源码文件以``_test``结尾：``xxx_test.go``
- 测试方法名以``Test``开头：``func TestXXX(t *testing.T) {...}``

## 基本数据类型
- bool
- string
- int int8 int16 int32 int64 
- uint uint8 uint16 uint32 uint64 uintptr(无符号整型，用于存放一个指针)
- byte // alias for uint8
- rune // alias for int32,represents a Unicode code point
- float32 float64
- complex64(32 位实数和虚数) complex128(64 位实数和虚数)