package xstatus

type FsmStatus int8

const (
	FsmNone FsmStatus = iota
	FsmInState
	FsmFinal
)

func (f FsmStatus) String() string {
	switch f {
	case FsmNone:
		return "fsm-none"
	case FsmInState:
		return "fsm-in-state"
	case FsmFinal:
		return "fsm-final"
	default:
		return "fsm-?"
	}
}
