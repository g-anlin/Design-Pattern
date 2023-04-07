package factory

// Factory 工厂的抽象接口，规定工厂的行为
type Factory interface{
  GenerateAnimal(species string) Animal
}

