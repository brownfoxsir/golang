package extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}
type Dog struct {
	p *Pet
}

// 匿名嵌入类型
type Cat struct {
	Pet
}

func (p *Pet) Speak() {
	fmt.Println("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)

}

func (d *Dog) Speak() {
	fmt.Println("wang!")
}

func (d *Dog) SpeakTo(host string) {
	d.Speak()
	fmt.Println(" ", host)

}

//自定义了speak方法，但是没有覆盖掉父类的方法
func (c *Cat) Speak() {
	fmt.Println("miao")
}

func TestDog(t *testing.T) {
	dog := &Dog{}
	dog.SpeakTo("wang")

	cat := new(Cat)
	cat.SpeakTo("miao")

}
