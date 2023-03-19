package regexps

import (
	"regexp"
)

// Regexps are referred from https://github.com/go-playground/validator/blob/v10.11.1/regexes.go.

const (
	alphaRegexString                 = "^[a-zA-Z]+$"
	alphaNumericRegexString          = "^[a-zA-Z0-9]+$"
	alphaUnicodeRegexString          = "^[\\p{L}]+$"
	alphaUnicodeNumericRegexString   = "^[\\p{L}\\p{N}]+$"
	numericRegexString               = "^[-+]?[0-9]+(?:\\.[0-9]+)?$"
	numberRegexString                = "^[0-9]+$"
	hexadecimalRegexString           = "^(0[xX])?[0-9a-fA-F]+$"
	hexColorRegexString              = "^#(?:[0-9a-fA-F]{3}|[0-9a-fA-F]{4}|[0-9a-fA-F]{6}|[0-9a-fA-F]{8})$"
	rgbRegexString                   = "^rgb\\(\\s*(?:(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])|(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%)\\s*\\)$"
	rgbaRegexString                  = "^rgba\\(\\s*(?:(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])|(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%)\\s*,\\s*(?:(?:0.[1-9]*)|[01])\\s*\\)$"
	hslRegexString                   = "^hsl\\(\\s*(?:0|[1-9]\\d?|[12]\\d\\d|3[0-5]\\d|360)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*\\)$"
	hslaRegexString                  = "^hsla\\(\\s*(?:0|[1-9]\\d?|[12]\\d\\d|3[0-5]\\d|360)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*,\\s*(?:(?:0.[1-9]*)|[01])\\s*\\)$"
	emailRegexString                 = "^(?:(?:(?:(?:[a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(?:\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|(?:(?:\\x22)(?:(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(?:\\x20|\\x09)+)?(?:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(\\x20|\\x09)+)?(?:\\x22))))@(?:(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
	e164RegexString                  = "^\\+[1-9]?[0-9]{7,14}$"
	base64RegexString                = "^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$"
	base64URLRegexString             = "^(?:[A-Za-z0-9-_]{4})*(?:[A-Za-z0-9-_]{2}==|[A-Za-z0-9-_]{3}=|[A-Za-z0-9-_]{4})$"
	iSBN10RegexString                = "^(?:[0-9]{9}X|[0-9]{10})$"
	iSBN13RegexString                = "^(?:(?:97(?:8|9))[0-9]{10})$"
	uUID3RegexString                 = "^[0-9a-f]{8}-[0-9a-f]{4}-3[0-9a-f]{3}-[0-9a-f]{4}-[0-9a-f]{12}$"
	uUID4RegexString                 = "^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
	uUID5RegexString                 = "^[0-9a-f]{8}-[0-9a-f]{4}-5[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
	uUIDRegexString                  = "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"
	uUID3RFC4122RegexString          = "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-3[0-9a-fA-F]{3}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"
	uUID4RFC4122RegexString          = "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$"
	uUID5RFC4122RegexString          = "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-5[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$"
	uUIDRFC4122RegexString           = "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"
	uLIDRegexString                  = "^[A-HJKMNP-TV-Z0-9]{26}$"
	md4RegexString                   = "^[0-9a-f]{32}$"
	md5RegexString                   = "^[0-9a-f]{32}$"
	sha256RegexString                = "^[0-9a-f]{64}$"
	sha384RegexString                = "^[0-9a-f]{96}$"
	sha512RegexString                = "^[0-9a-f]{128}$"
	ripemd128RegexString             = "^[0-9a-f]{32}$"
	ripemd160RegexString             = "^[0-9a-f]{40}$"
	tiger128RegexString              = "^[0-9a-f]{32}$"
	tiger160RegexString              = "^[0-9a-f]{40}$"
	tiger192RegexString              = "^[0-9a-f]{48}$"
	aSCIIRegexString                 = "^[\x00-\x7F]*$"
	printableASCIIRegexString        = "^[\x20-\x7E]*$"
	multibyteRegexString             = "[^\x00-\x7F]"
	dataURIRegexString               = `^data:((?:\w+\/(?:([^;]|;[^;]).)+)?)`
	latitudeRegexString              = "^[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?)$"
	longitudeRegexString             = "^[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$"
	sSNRegexString                   = `^[0-9]{3}[ -]?(0[1-9]|[1-9][0-9])[ -]?([1-9][0-9]{3}|[0-9][1-9][0-9]{2}|[0-9]{2}[1-9][0-9]|[0-9]{3}[1-9])$`
	hostnameRegexStringRFC952        = `^[a-zA-Z]([a-zA-Z0-9\-]+[\.]?)*[a-zA-Z0-9]$`                                                                   // https://tools.ietf.org/html/rfc952
	hostnameRegexStringRFC1123       = `^([a-zA-Z0-9]{1}[a-zA-Z0-9-]{0,62}){1}(\.[a-zA-Z0-9]{1}[a-zA-Z0-9-]{0,62})*?$`                                 // accepts hostname starting with a digit https://tools.ietf.org/html/rfc1123
	fqdnRegexStringRFC1123           = `^([a-zA-Z0-9]{1}[a-zA-Z0-9-]{0,62})(\.[a-zA-Z0-9]{1}[a-zA-Z0-9-]{0,62})*?(\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\.?$` // same as hostnameRegexStringRFC1123 but must contain a non numerical TLD (possibly ending with '.')
	btcAddressRegexString            = `^[13][a-km-zA-HJ-NP-Z1-9]{25,34}$`                                                                             // bitcoin address
	btcAddressUpperRegexStringBech32 = `^BC1[02-9AC-HJ-NP-Z]{7,76}$`                                                                                   // bitcoin bech32 address https://en.bitcoin.it/wiki/Bech32
	btcAddressLowerRegexStringBech32 = `^bc1[02-9ac-hj-np-z]{7,76}$`                                                                                   // bitcoin bech32 address https://en.bitcoin.it/wiki/Bech32
	ethAddressRegexString            = `^0x[0-9a-fA-F]{40}$`
	ethAddressUpperRegexString       = `^0x[0-9A-F]{40}$`
	ethAddressLowerRegexString       = `^0x[0-9a-f]{40}$`
	uRLEncodedRegexString            = `^(?:[^%]|%[0-9A-Fa-f]{2})*$`
	hTMLEncodedRegexString           = `&#[x]?([0-9a-fA-F]{2})|(&gt)|(&lt)|(&quot)|(&amp)+[;]?`
	hTMLRegexString                  = `<[/]?([a-zA-Z]+).*?>`
	jWTRegexString                   = "^[A-Za-z0-9-_]+\\.[A-Za-z0-9-_]+\\.[A-Za-z0-9-_]*$"
	splitParamsRegexString           = `'[^']*'|\S+`
	bicRegexString                   = `^[A-Za-z]{6}[A-Za-z0-9]{2}([A-Za-z0-9]{3})?$`
	semverRegexString                = `^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$` // numbered capture groups https://semver.org/
	dnsRegexStringRFC1035Label       = "^[a-z]([-a-z0-9]*[a-z0-9]){0,62}$"
)

var (
	AlphaRegex                 = regexp.MustCompile(alphaRegexString)
	AlphaNumericRegex          = regexp.MustCompile(alphaNumericRegexString)
	AlphaUnicodeRegex          = regexp.MustCompile(alphaUnicodeRegexString)
	AlphaUnicodeNumericRegex   = regexp.MustCompile(alphaUnicodeNumericRegexString)
	NumericRegex               = regexp.MustCompile(numericRegexString)
	NumberRegex                = regexp.MustCompile(numberRegexString)
	HexadecimalRegex           = regexp.MustCompile(hexadecimalRegexString)
	HexColorRegex              = regexp.MustCompile(hexColorRegexString)
	RgbRegex                   = regexp.MustCompile(rgbRegexString)
	RgbaRegex                  = regexp.MustCompile(rgbaRegexString)
	HslRegex                   = regexp.MustCompile(hslRegexString)
	HslaRegex                  = regexp.MustCompile(hslaRegexString)
	E164Regex                  = regexp.MustCompile(e164RegexString)
	EmailRegex                 = regexp.MustCompile(emailRegexString)
	Base64Regex                = regexp.MustCompile(base64RegexString)
	Base64URLRegex             = regexp.MustCompile(base64URLRegexString)
	ISBN10Regex                = regexp.MustCompile(iSBN10RegexString)
	ISBN13Regex                = regexp.MustCompile(iSBN13RegexString)
	UUID3Regex                 = regexp.MustCompile(uUID3RegexString)
	UUID4Regex                 = regexp.MustCompile(uUID4RegexString)
	UUID5Regex                 = regexp.MustCompile(uUID5RegexString)
	UUIDRegex                  = regexp.MustCompile(uUIDRegexString)
	UUID3RFC4122Regex          = regexp.MustCompile(uUID3RFC4122RegexString)
	UUID4RFC4122Regex          = regexp.MustCompile(uUID4RFC4122RegexString)
	UUID5RFC4122Regex          = regexp.MustCompile(uUID5RFC4122RegexString)
	UUIDRFC4122Regex           = regexp.MustCompile(uUIDRFC4122RegexString)
	ULIDRegex                  = regexp.MustCompile(uLIDRegexString)
	Md4Regex                   = regexp.MustCompile(md4RegexString)
	Md5Regex                   = regexp.MustCompile(md5RegexString)
	Sha256Regex                = regexp.MustCompile(sha256RegexString)
	Sha384Regex                = regexp.MustCompile(sha384RegexString)
	Sha512Regex                = regexp.MustCompile(sha512RegexString)
	Ripemd128Regex             = regexp.MustCompile(ripemd128RegexString)
	Ripemd160Regex             = regexp.MustCompile(ripemd160RegexString)
	Tiger128Regex              = regexp.MustCompile(tiger128RegexString)
	Tiger160Regex              = regexp.MustCompile(tiger160RegexString)
	Tiger192Regex              = regexp.MustCompile(tiger192RegexString)
	ASCIIRegex                 = regexp.MustCompile(aSCIIRegexString)
	PrintableASCIIRegex        = regexp.MustCompile(printableASCIIRegexString)
	MultibyteRegex             = regexp.MustCompile(multibyteRegexString)
	DataURIRegex               = regexp.MustCompile(dataURIRegexString)
	LatitudeRegex              = regexp.MustCompile(latitudeRegexString)
	LongitudeRegex             = regexp.MustCompile(longitudeRegexString)
	SSNRegex                   = regexp.MustCompile(sSNRegexString)
	HostnameRegexRFC952        = regexp.MustCompile(hostnameRegexStringRFC952)
	HostnameRegexRFC1123       = regexp.MustCompile(hostnameRegexStringRFC1123)
	FqdnRegexRFC1123           = regexp.MustCompile(fqdnRegexStringRFC1123)
	BtcAddressRegex            = regexp.MustCompile(btcAddressRegexString)
	BtcUpperAddressRegexBech32 = regexp.MustCompile(btcAddressUpperRegexStringBech32)
	BtcLowerAddressRegexBech32 = regexp.MustCompile(btcAddressLowerRegexStringBech32)
	EthAddressRegex            = regexp.MustCompile(ethAddressRegexString)
	EthAddressRegexUpper       = regexp.MustCompile(ethAddressUpperRegexString)
	EthAddressRegexLower       = regexp.MustCompile(ethAddressLowerRegexString)
	URLEncodedRegex            = regexp.MustCompile(uRLEncodedRegexString)
	HTMLEncodedRegex           = regexp.MustCompile(hTMLEncodedRegexString)
	HTMLRegex                  = regexp.MustCompile(hTMLRegexString)
	JWTRegex                   = regexp.MustCompile(jWTRegexString)
	SplitParamsRegex           = regexp.MustCompile(splitParamsRegexString)
	BicRegex                   = regexp.MustCompile(bicRegexString)
	SemverRegex                = regexp.MustCompile(semverRegexString)
	DnsRegexRFC1035Label       = regexp.MustCompile(dnsRegexStringRFC1035Label)
)
