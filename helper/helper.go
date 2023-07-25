package helper

import "reflect"

type UserWlCommission struct {
	WlCommMember float64
	WlCommAgent  float64
	WlCommMaster float64
	WlCommSenior float64
	WlCommSuper  float64
	WlCommWeb    float64
}

type CalculateWlCommissionInput struct {
	Stake float64

	CommMember float64
	CommAgent  float64
	CommMaster float64
	CommSenior float64
	CommSuper  float64
	CommWeb    float64

	ShAgent  float64
	ShMaster float64
	ShSenior float64
	ShSuper  float64
	ShWeb    float64
}

func GetStructValueByKeyName(s interface{}, name string) interface{} {
	value := reflect.ValueOf(s)
	if value.Kind() != reflect.Struct {
		return nil
	}
	field := value.FieldByName(name)
	if field.IsValid() {
		return field.Interface()
	}
	return nil
}

func CalculateWlCommission(input CalculateWlCommissionInput) (float64, float64, float64, float64, float64, float64) {
	agentCommDiff := input.CommAgent - input.CommMember
	masterCommDiff := input.CommMaster - agentCommDiff - input.CommMember
	seniorCommDiff := input.CommSenior - masterCommDiff - agentCommDiff - input.CommMember
	superCommDiff := input.CommSuper - seniorCommDiff - masterCommDiff - agentCommDiff - input.CommMember

	wlCommMember := input.Stake * input.CommMember / 100
	wlCommAgent := (agentCommDiff * (input.ShMaster + input.ShSenior + input.ShSuper + input.ShWeb) * (input.Stake / 100)) - (wlCommMember * input.ShAgent)
	WlCommMaster := (masterCommDiff * (input.ShSenior + input.ShSuper + input.ShWeb) * (input.Stake / 100)) - ((input.Stake * input.CommAgent / 100) * input.ShMaster)
	WlCommSenior := (seniorCommDiff * (input.ShSuper + input.ShWeb) * (input.Stake / 100)) - ((input.Stake * input.CommMaster / 100) * input.ShSenior)
	WlCommSuper := (superCommDiff * (input.ShWeb) * (input.Stake / 100)) - ((input.Stake * input.CommSenior / 100) * input.ShSuper)
	WlCommWeb := 0 - ((input.Stake * input.CommSuper / 100) * input.ShWeb)

	return wlCommMember, wlCommAgent, WlCommMaster, WlCommSenior, WlCommSuper, WlCommWeb
}
