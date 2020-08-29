package xstatus

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"testing"
)

func TestDbStatus(t *testing.T) {
	xtesting.Equal(t, fmt.Sprintf("%v", DbSuccess), "db-success")
	xtesting.Equal(t, DbSuccess.String(), "db-success")
	xtesting.Equal(t, DbNotFound.String(), "db-not-found")
	xtesting.Equal(t, DbExisted.String(), "db-existed")
	xtesting.Equal(t, DbFailed.String(), "db-failed")
	xtesting.Equal(t, DbTagA.String(), "db-tag-a")
	xtesting.Equal(t, DbTagB.String(), "db-tag-b")
	xtesting.Equal(t, DbTagC.String(), "db-tag-c")
	xtesting.Equal(t, DbTagD.String(), "db-tag-d")
	xtesting.Equal(t, DbTagE.String(), "db-tag-e")
	xtesting.Equal(t, DbStatus(20).String(), "db-?")
}

func TestFsmStatus(t *testing.T) {
	xtesting.Equal(t, fmt.Sprintf("%v", FsmNone), "fsm-none")
	xtesting.Equal(t, FsmNone.String(), "fsm-none")
	xtesting.Equal(t, FsmInState.String(), "fsm-in-state")
	xtesting.Equal(t, FsmFinal.String(), "fsm-final")
	xtesting.Equal(t, FsmStatus(20).String(), "fsm-?")
}

func TestJwtStatus(t *testing.T) {
	xtesting.Equal(t, fmt.Sprintf("%v", JwtSuccess), "jwt-success")
	xtesting.Equal(t, JwtSuccess.String(), "jwt-success")
	xtesting.Equal(t, JwtExpired.String(), "jwt-expired")
	xtesting.Equal(t, JwtNotValid.String(), "jwt-not-valid")
	xtesting.Equal(t, JwtIssuer.String(), "jwt-issuer")
	xtesting.Equal(t, JwtSubject.String(), "jwt-subject")
	xtesting.Equal(t, JwtInvalid.String(), "jwt-invalid")
	xtesting.Equal(t, JwtBlank.String(), "jwt-blank")
	xtesting.Equal(t, JwtNotFound.String(), "jwt-not-found")
	xtesting.Equal(t, JwtUserErr.String(), "jwt-user-err")
	xtesting.Equal(t, JwtFailed.String(), "jwt-failed")
	xtesting.Equal(t, JwtTagA.String(), "jwt-tag-a")
	xtesting.Equal(t, JwtTagB.String(), "jwt-tag-b")
	xtesting.Equal(t, JwtTagC.String(), "jwt-tag-c")
	xtesting.Equal(t, JwtStatus(20).String(), "jwt-?")
}
