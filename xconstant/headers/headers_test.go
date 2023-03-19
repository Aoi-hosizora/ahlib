package headers

import (
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"testing"
)

func TestHeaders(t *testing.T) {
	for _, header := range []string{
		// Standard header fields
		AIM, Accept, AcceptCH, AcceptCharset, AcceptDatetime, AcceptEncoding, AcceptLanguage, AcceptPatch, AcceptRanges, AccessControlAllowCredentials,
		AccessControlAllowHeaders, AccessControlAllowMethods, AccessControlAllowOrigin, AccessControlExposeHeaders, AccessControlMaxAge, AccessControlRequestHeaders,
		AccessControlRequestMethod, Age, Allow, AltSvc, Authorization, CacheControl, Connection, ContentDisposition, ContentEncoding, ContentLanguage, ContentLength,
		ContentLocation, ContentMD5, ContentRange, ContentType, Cookie, Date, DeltaBase, ETag, Expect, Expires, Forwarded, From, HTTP2Settings, Host, IM, IfMatch,
		IfModifiedSince, IfNoneMatch, IfRange, IfUnmodifiedSince, LastModified, Link, Location, MaxForwards, Origin, P3P, Pragma, Prefer, PreferenceApplied,
		ProxyAuthenticate, ProxyAuthorization, PublicKeyPins, Range, Referer, RetryAfter, Server, SetCookie, StrictTransportSecurity, TE, Tk, Trailer, TransferEncoding,
		Upgrade, UserAgent, Vary, Via, WWWAuthenticate, Warning, XFrameOptions,

		// Common non-standard header fields
		ContentSecurityPolicy, CorrelationID, DNT, ExpectCT, FrontEndHttps, NEL, PermissionsPolicy, ProxyConnection, Refresh, ReportTo, SaveData, SecGPC, Status,
		TimingAllowOrigin, UpgradeInsecureRequests, XATTDeviceId, XContentDuration, XContentSecurityPolicy, XContentTypeOptions, XCorrelationID, XCsrfToken, XForwardedFor,
		XForwardedHost, XForwardedProto, XHttpMethodOverride, XPoweredBy, XRateLimitLimit, XRateLimitRemaining, XRateLimitReset, XRealIP, XRedirectBy, XRequestID,
		XRequestedWith, XUACompatible, XUIDH, XWapProfile, XWebKitCSP, XXSSProtection,
	} {
		xtesting.NotBlankString(t, header)
	}
}
