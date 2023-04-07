package factory

import "strings"

// Factory 工厂接口，接口中包含所有工厂的生产方法
// 抽象工厂模式的缺点就是在此，每增加一个生产新类型产品的工厂，就要修改原有的所有工厂类
type Factory interface{
  GenerateAnimal(species string) Animal
  GeneratePhone(brand string) Phone
}

// AbstructFactory 抽象工厂接口
type AbstructFactory interface {
  GenerateFactory(factory string) Factory
}

// FactoryProducer 工厂生产类
type FactoryProducer struct {}

// NewFactoryProducer 创建工厂生产实例
func NewFactoryProducer() AbstructFactory {
  return &FactoryProducer{}
}

// GenerateFactory 根据传入的工厂类型生成工厂
func (p *FactoryProducer) GenerateFactory(factory string) Factory{
  if strings.ToLower(factory) == "animal" {
    return NewAnimalFactory()
  }else if strings.ToLower(factory) == "phone"{
    return NewPhoneFactory()
  }else{
    return nil
  }
}