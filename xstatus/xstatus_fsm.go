package xstatus

// FsmStatus represents a status value for a demo finite status machine. Actually this is a dummy type.
type FsmStatus uint64

const (
	FsmNone  FsmStatus = iota // None
	FsmState                  // State
	FsmFinal                  // Final

	FsmTagA FsmStatus = iota + 98 // Tag a, start from 101
	FsmTagB                       // Tag b
	FsmTagC                       // Tag c
	FsmTagD                       // Tag d
	FsmTagE                       // Tag e
	FsmTagF                       // Tag f
	FsmTagG                       // Tag g
)

func (f FsmStatus) String() string {
	switch f {
	case FsmNone:
		return "fsm-none"
	case FsmState:
		return "fsm-state"
	case FsmFinal:
		return "fsm-final"
	case FsmTagA:
		return "fsm-tag-a"
	case FsmTagB:
		return "fsm-tag-b"
	case FsmTagC:
		return "fsm-tag-c"
	case FsmTagD:
		return "fsm-tag-d"
	case FsmTagE:
		return "fsm-tag-e"
	case FsmTagF:
		return "fsm-tag-f"
	case FsmTagG:
		return "fsm-tag-g"
	default:
		return "fsm-?"
	}
}