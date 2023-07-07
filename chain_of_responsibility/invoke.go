package chainofresponsibility

import "fmt"

func Invoke() {
	// 申明旅客
	passenger := &Passenger{
		name:                 "张三",
		hasCheckedIn:         false,
		hasRegisteredLuggage: false,
		hasSecurityChecked:   false,
		hasBoard:             false,
	}
	// 链接责任链
	boardNode := &BoardProcesser{}

	securityCheck := &SecurityCheckProcess{}
	securityCheck.SetNextStep(boardNode)

	registerLuggageNode := &RegisterLuggageProcesser{}
	registerLuggageNode.SetNextStep(securityCheck)

	checkInNode := &CheckInProcesser{}
	checkInNode.SetNextStep(registerLuggageNode)

	checkInNode.Process(passenger)

	fmt.Printf("旅客状态：%+v", passenger)
}
