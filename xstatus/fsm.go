package xstatus

type FsmStatus int8

const (
	FsmNone FsmStatus = iota
	FsmInState
	FsmFinal
	FsmTagA
	FsmTagB
	FsmTagC
)

func (f FsmStatus) String() string {
	switch f {
	case FsmNone:
		return "fsm-none"
	case FsmInState:
		return "fsm-in-state"
	case FsmFinal:
		return "fsm-final"
	case FsmTagA:
		return "fsm-tag-a"
	case FsmTagB:
		return "fsm-tag-b"
	case FsmTagC:
		return "fsm-tag-c"
	default:
		return "fsm-?"
	}
}
