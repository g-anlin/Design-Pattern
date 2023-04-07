package factory

import "fmt"

// Invoke 使用工厂
func InvokeFactory() {
  fmt.Println("------Factory------")
  f := NewAnimalFactory()
  // 狗叫
  if a := f.GenerateAnimal("DOG"); a != nil {
    a.Speak()
  }
  // 猫叫
  if a := f.GenerateAnimal("Cat"); a != nil {
    a.Speak()
  }
  // 猪叫
  if a := f.GenerateAnimal("pIG"); a != nil {
    a.Speak()
  }
}

func InvokeAbstructFactory() {
  fmt.Println("------AbstructFactory------")
  fproducer := NewFactoryProducer()
  phoneFactory := fproducer.GenerateFactory("phone")
  if phoneFactory == nil {
    return
  }
  // iphone的价格
  if p := phoneFactory.GeneratePhone("IPhone"); p != nil {
    p.ShowValue()
  }
  // oppo的价格
  if p := phoneFactory.GeneratePhone("OPPO"); p != nil {
    p.ShowValue()
  }
  // vivo的价格
  if p := phoneFactory.GeneratePhone("VIVO"); p != nil {
    p.ShowValue()
  }
  
}