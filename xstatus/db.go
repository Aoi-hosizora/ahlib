package xstatus

type DbStatus int8

const (
	DbSuccess  DbStatus = iota // success (CRUD)
	DbNotFound                 // not found (retrieve, update, delete)
	DbExisted                  // existed (create update)
	DbFailed                   // failed (CRUD)
	DbTagA                     // tag a
	DbTagB                     // tag b
	DbTagC                     // tag c
	DbTagD                     // tag d
	DbTagE                     // tag e
)

func (d DbStatus) String() string {
	switch d {
	case DbSuccess:
		return "db-success"
	case DbNotFound:
		return "db-not-found"
	case DbExisted:
		return "db-existed"
	case DbFailed:
		return "db-failed"
	case DbTagA:
		return "db-tag-a"
	case DbTagB:
		return "db-tag-b"
	case DbTagC:
		return "db-tag-c"
	case DbTagD:
		return "db-tag-d"
	case DbTagE:
		return "db-tag-e"
	default:
		return "db-?"
	}
}
