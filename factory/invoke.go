package factory

// Invoke 使用工厂
func Invoke() {
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