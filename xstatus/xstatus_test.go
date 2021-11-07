package xstatus

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"testing"
)

func TestDbStatus(t *testing.T) {
	xtesting.Equal(t, DbUnknown.String(), "db-unknown")
	xtesting.Equal(t, DbSuccess.String(), "db-success")
	xtesting.Equal(t, DbNotFound.String(), "db-not-found")
	xtesting.Equal(t, DbExisted.String(), "db-existed")
	xtesting.Equal(t, DbFailed.String(), "db-failed")
	xtesting.Equal(t, DbTagA.String(), "db-tag-a")
	xtesting.Equal(t, DbTagB.String(), "db-tag-b")
	xtesting.Equal(t, DbTagC.String(), "db-tag-c")
	xtesting.Equal(t, DbTagD.String(), "db-tag-d")
	xtesting.Equal(t, DbTagE.String(), "db-tag-e")

	xtesting.Equal(t, fmt.Sprintf("%v", DbSuccess), "db-success")
	xtesting.Equal(t, DbStatus(999).String(), "db-?")
}

func TestFsmStatus(t *testing.T) {
	xtesting.Equal(t, FsmNone.String(), "fsm-none")
	xtesting.Equal(t, FsmState.String(), "fsm-state")
	xtesting.Equal(t, FsmFinal.String(), "fsm-final")
	xtesting.Equal(t, FsmTagA.String(), "fsm-tag-a")
	xtesting.Equal(t, FsmTagB.String(), "fsm-tag-b")
	xtesting.Equal(t, FsmTagC.String(), "fsm-tag-c")
	xtesting.Equal(t, FsmTagD.String(), "fsm-tag-d")
	xtesting.Equal(t, FsmTagE.String(), "fsm-tag-e")

	xtesting.Equal(t, fmt.Sprintf("%v", FsmNone), "fsm-none")
	xtesting.Equal(t, FsmStatus(999).String(), "fsm-?")
}

func TestJwtStatus(t *testing.T) {
	xtesting.Equal(t, JwtUnknown.String(), "jwt-unknown")
	xtesting.Equal(t, JwtSuccess.String(), "jwt-success")
	xtesting.Equal(t, JwtBlank.String(), "jwt-blank")
	xtesting.Equal(t, JwtInvalid.String(), "jwt-invalid")
	xtesting.Equal(t, JwtTokenNotFound.String(), "jwt-token-not-found")
	xtesting.Equal(t, JwtUserNotFound.String(), "jwt-user-not-found")
	xtesting.Equal(t, JwtFailed.String(), "jwt-failed")
	xtesting.Equal(t, JwtTagA.String(), "jwt-tag-a")
	xtesting.Equal(t, JwtTagB.String(), "jwt-tag-b")
	xtesting.Equal(t, JwtTagC.String(), "jwt-tag-c")
	xtesting.Equal(t, JwtTagD.String(), "jwt-tag-d")
	xtesting.Equal(t, JwtTagE.String(), "jwt-tag-e")

	xtesting.Equal(t, JwtAudience.String(), "jwt-audience")
	xtesting.Equal(t, JwtExpired.String(), "jwt-expired")
	xtesting.Equal(t, JwtId.String(), "jwt-id")
	xtesting.Equal(t, JwtIssuedAt.String(), "jwt-issued-at")
	xtesting.Equal(t, JwtIssuer.String(), "jwt-issuer")
	xtesting.Equal(t, JwtNotValidYet.String(), "jwt-not-valid-yet")
	xtesting.Equal(t, JwtSubject.String(), "jwt-subject")
	xtesting.Equal(t, JwtClaimsInvalid.String(), "jwt-claims-invalid")

	xtesting.Equal(t, fmt.Sprintf("%v", JwtSuccess), "jwt-success")
	xtesting.Equal(t, JwtStatus(999).String(), "jwt-?")

	s := JwtAudience | JwtExpired | JwtId | JwtIssuedAt | JwtIssuer | JwtNotValidYet | JwtSubject | JwtClaimsInvalid
	xtesting.NotEqual(t, s&JwtAudience, 0)
	xtesting.NotEqual(t, s&JwtExpired, 0)
	xtesting.NotEqual(t, s&JwtId, 0)
	xtesting.NotEqual(t, s&JwtIssuedAt, 0)
	xtesting.NotEqual(t, s&JwtIssuer, 0)
	xtesting.NotEqual(t, s&JwtNotValidYet, 0)
	xtesting.NotEqual(t, s&JwtSubject, 0)
	xtesting.NotEqual(t, s&JwtClaimsInvalid, 0)
}
