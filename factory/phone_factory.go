package factory

import "strings"

type PhoneFactory struct{}

func NewPhoneFactory() *PhoneFactory {
  return &PhoneFactory{}
}

// GeneratePhone 生产手机方法，空实现
func (p *PhoneFactory)GeneratePhone(brand string) Phone {
  if(strings.ToLower(brand) == "iphone"){
    return NewIPhone(8000)
  }else if(strings.ToLower(brand) == "oppo"){
    return NewOPPO(6000)
  }else if(strings.ToLower(brand) == "vivo"){
    return NewVIVO(4000)
  }else{
    return nil
  }
}

// GenerateAnimal 生产动物方法，空实现
func (p *PhoneFactory)GenerateAnimal(species string) Animal{
  return nil
}