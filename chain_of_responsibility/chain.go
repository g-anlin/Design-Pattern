package chainofresponsibility

import (
	"fmt"
)

// Passenger 旅客
type Passenger struct {
	name                 string // 姓名
	hasCheckedIn         bool   // 是否值机
	hasRegisteredLuggage bool   // 是否托运过行李
	hasSecurityChecked   bool   // 是否通过安检
	hasBoard             bool   // 是否登机
}

// Processer 处理器接口，每个处理节点需要实现该接口
type Processer interface {
	SetNextStep(p Processer)
	Process(passenger *Passenger)
}

// baseProcesser 统一的基类处理器
type baseProcesser struct {
	nextStep Processer
}

// SetNextStep 基类提供的统一的设置下一步处理者的方法
func (b *baseProcesser) SetNextStep(p Processer) {
	b.nextStep = p
}

// Process 基类提供的统一的处理方法
func (b *baseProcesser) Process(p *Passenger) {
	if b.nextStep != nil {
		b.nextStep.Process(p)
	}
}

// CheckInProcesser 值机处理器
type CheckInProcesser struct {
	// nextStep Processer
	baseProcesser // 嵌入基类
}

// SetNextStep 设置下一个处理节点，直接使用基类的方法
// func (c *CheckInProcesser) SetNextStep(p Processer) {
//   c.nextStep = p
// }

// Process 该节点的处理动作
func (c *CheckInProcesser) Process(p *Passenger) {
	// 处理值机
	if !p.hasCheckedIn {
		fmt.Printf("正在为旅客 %s 办理值机手续...\n", p.name)
		p.hasCheckedIn = true
	}
	// if c.nextStep != nil {
	//   c.nextStep.Process(p)
	// }
	// 使用基类的Process方法
	c.baseProcesser.Process(p)
}

// RegisterLuggageProcesser 托运处理器
type RegisterLuggageProcesser struct {
	// nextStep Processer
	baseProcesser
}

// SetNextStep 设置下一个处理节点
// func (r *RegisterLuggageProcesser) SetNextStep(p Processer) {
//  r.nextStep = p
// }

// Process 该节点的处理动作
func (r *RegisterLuggageProcesser) Process(p *Passenger) {
	// 处理托运
	if !p.hasRegisteredLuggage {
		fmt.Printf("正在为旅客 %s 办理托运手续...\n", p.name)
		p.hasRegisteredLuggage = true
	}
	// if r.nextStep != nil {
	//  r.nextStep.Process(p)
	//}
	r.baseProcesser.Process(p)
}

// SecurityCheck 安检处理器
type SecurityCheckProcess struct {
	// nextStep Processer
	baseProcesser
}

// SetNextStep 设置下一个处理节点
//func (s *SecurityCheckProcess) SetNextStep(p Processer) {
//  s.nextStep = p
//}

// Process 该节点的处理动作
func (s *SecurityCheckProcess) Process(p *Passenger) {
	// 处理托运
	if !p.hasSecurityChecked {
		fmt.Printf("旅客 %s 正在安检...\n", p.name)
		p.hasSecurityChecked = true
	}
	// if s.nextStep != nil {
	//  s.nextStep.Process(p)
	//}
	s.baseProcesser.Process(p)
}

// BoardProcesser 安检处理器
type BoardProcesser struct {
	// nextStep Processer
	baseProcesser
}

// SetNextStep 设置下一个处理节点
// func (b *BoardProcesser) SetNextStep(p Processer) {
//  b.nextStep = p
// }

// Process 该节点的处理动作
func (b *BoardProcesser) Process(p *Passenger) {
	// 处理登机
	if !p.hasBoard {
		fmt.Printf("旅客 %s 正在登机...\n", p.name)
		p.hasBoard = true
	}
	// if b.nextStep != nil {
	//  b.nextStep.Process(p)
	//}
	b.baseProcesser.Process(p)
}
