package factory

import "fmt"

type Phone interface {
  ShowValue()
}

// IPhone 手机
type IPhone struct {
  Price int
}

func NewIPhone(value int) Phone {
  return &IPhone{
    Price: value,
  }
}

func (i *IPhone) ShowValue() {
  fmt.Printf("Price of IPhone: %v\n", i.Price)
}

// OPPO 手机
type OPPO struct {
  Price int
}

func NewOPPO(value int) Phone {
  return &OPPO{
    Price: value,
  }
}

func (o *OPPO) ShowValue() {
  fmt.Printf("Price of OPPO: %v\n", o.Price)
}

// VIVO 手机
type VIVO struct {
  Price int
}

func NewVIVO(value int) Phone {
  return &VIVO{
    Price: value,
  }
}

func (v *VIVO) ShowValue() {
  fmt.Printf("Price of VIVO: %v\n", v.Price)
}