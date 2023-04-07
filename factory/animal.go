package factory

import "fmt"

// Animal 定义动物行为的接口
type Animal interface {
  Speak()
}

// Dog 狗类
type Dog struct {}

// NewDog 新建狗实例
func NewDog() Animal {
  return &Dog{}
}

// Speak 实现 Animal 接口
func (d *Dog) Speak() {
  fmt.Println("Dog speak: wang wang wang")
}

// Cat 猫类
type Cat struct {}

// NewCat 新建猫实例
func NewCat() Animal {
  return &Cat{}
}

// Speak 实现 Animal 接口
func (c *Cat) Speak() {
  fmt.Println("Cat speak: miao miao miao")
}

// Pig 猪类
type Pig struct {}

// NewPig 新建猪实例
func NewPig() Animal {
  return &Pig{}
}

// Speak 实现 Animal 接口
func (p *Pig) Speak() {
  fmt.Println("Pig Speak: hen hen")
}