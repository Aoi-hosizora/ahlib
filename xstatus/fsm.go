package xstatus

// FsmStatus represents a status value for finite status machine. Actually this is a dummy type.
type FsmStatus uint64

const (
	FsmNone    FsmStatus = iota      // None
	FsmInState                       // In state
	FsmFinal                         // Final
	FsmTagA    FsmStatus = iota + 98 // Tag a
	FsmTagB                          // Tag b
	FsmTagC                          // Tag c
	FsmTagD                          // Tag d
	FsmTagE                          // Tag e
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
	case FsmTagD:
		return "fsm-tag-d"
	case FsmTagE:
		return "fsm-tag-e"
	default:
		return "fsm-?"
	}
}
