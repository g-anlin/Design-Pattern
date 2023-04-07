package factory

import "strings"

// AnimalFactory 动物工厂
type AnimalFactory struct {}

// NewAnimalFactory 新建动物工厂实例
func NewAnimalFactory() Factory{
  return &AnimalFactory{}
}

// GenerateAnimal 生产动物方法
func (f *AnimalFactory) GenerateAnimal(species string) Animal {
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