package xstatus

// DbStatus represents a status value for the result of database operation.
type DbStatus uint64

const (
	DbUnknown  DbStatus = iota // Unknown (?)
	DbSuccess                  // Success (CRUD)
	DbNotFound                 // Not found (RUD)
	DbExisted                  // Existed (CU)
	DbFailed                   // Failed (CRUD)

	DbTagA DbStatus = iota + 96 // Tag a, start from 101
	DbTagB                      // Tag b
	DbTagC                      // Tag c
	DbTagD                      // Tag d
	DbTagE                      // Tag e
	DbTagF                      // Tag f
	DbTagG                      // Tag g
)

func (d DbStatus) String() string {
	switch d {
	case DbUnknown:
		return "db-unknown"
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
	case DbTagF:
		return "db-tag-f"
	case DbTagG:
		return "db-tag-g"
	default:
		return "db-?"
	}
}
