package xstatus

// JwtStatus represents a status value for the result of jwt and database operation.
type JwtStatus uint64

// 0 - 13
const (
	JwtUnknown JwtStatus = 0    // Unknown
	JwtSuccess JwtStatus = iota // Success
	JwtBlank                    // Blank token
	JwtInvalid                  // Invalid token (malformed, unverifiable, invalid signature)

	JwtTokenNotFound // Token not found
	JwtUserNotFound  // User not found
	JwtFailed        // Something error
	JwtTagA          // Tag a
	JwtTagB          // Tag b
	JwtTagC          // Tag c
	JwtTagD          // Tag d
	JwtTagE          // Tag e
	JwtTagF          // Tag f
	JwtTagG          // Tag g
)

// 16 - 2048
const (
	JwtAudience      JwtStatus = 16 << iota // AUD (Audience)
	JwtExpired                              // EXP (Expires at)
	JwtId                                   // JTI (Id)
	JwtIssuedAt                             // IAT (Issued at)
	JwtIssuer                               // ISS (Issuer)
	JwtNotValidYet                          // NBF (Not before)
	JwtSubject                              // SUB (Subject)
	JwtClaimsInvalid                        // Invalid claims, generic claims error
)

func (j JwtStatus) String() string {
	switch j {
	case JwtUnknown:
		return "jwt-unknown"
	case JwtSuccess:
		return "jwt-success"
	case JwtBlank:
		return "jwt-blank"
	case JwtInvalid:
		return "jwt-invalid"

	case JwtTokenNotFound:
		return "jwt-token-not-found"
	case JwtUserNotFound:
		return "jwt-user-not-found"
	case JwtFailed:
		return "jwt-failed"
	case JwtTagA:
		return "jwt-tag-a"
	case JwtTagB:
		return "jwt-tag-b"
	case JwtTagC:
		return "jwt-tag-c"
	case JwtTagD:
		return "jwt-tag-d"
	case JwtTagE:
		return "jwt-tag-e"
	case JwtTagF:
		return "jwt-tag-f"
	case JwtTagG:
		return "jwt-tag-g"

	case JwtAudience:
		return "jwt-audience"
	case JwtExpired:
		return "jwt-expired"
	case JwtId:
		return "jwt-id"
	case JwtIssuedAt:
		return "jwt-issued-at"
	case JwtIssuer:
		return "jwt-issuer"
	case JwtNotValidYet:
		return "jwt-not-valid-yet"
	case JwtSubject:
		return "jwt-subject"
	case JwtClaimsInvalid:
		return "jwt-claims-invalid"

	default:
		return "jwt-?"
	}
}
