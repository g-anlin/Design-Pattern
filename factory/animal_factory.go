package factory

import (
  "strings"
)

// AnimalFactory 动物工厂
type AnimalFactory struct {}
// NewAnimalFactory 新建动物工厂实例
func NewAnimalFactory() *AnimalFactory{
  return &AnimalFactory{}
}
// GenerateAnimal 生产动物方法
func (a *AnimalFactory) GenerateAnimal(species string) Animal {
  if strings.ToLower(species) == "dog" {
    return NewDog()
  }else if strings.ToLower(species) == "cat" {
    return NewCat()
  }else if strings.ToLower(species) == "pig" {
    return NewPig()
  }else {
      return nil
  }
}

// 生产手机方法，空实现
func (a *AnimalFactory) GeneratePhone(brand string) Phone {
  return nil
}