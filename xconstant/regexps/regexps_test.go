package regexps

import (
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"regexp"
	"testing"
)

func TestRegexps(t *testing.T) {
	for _, re := range []*regexp.Regexp{
		AlphaRegex, AlphaNumericRegex, AlphaUnicodeRegex, AlphaUnicodeNumericRegex, NumericRegex, NumberRegex, HexadecimalRegex, HexColorRegex, RgbRegex, RgbaRegex,
		HslRegex, HslaRegex, E164Regex, EmailRegex, Base64Regex, Base64URLRegex, ISBN10Regex, ISBN13Regex, UUID3Regex, UUID4Regex, UUID5Regex, UUIDRegex, UUID3RFC4122Regex,
		UUID4RFC4122Regex, UUID5RFC4122Regex, UUIDRFC4122Regex, ULIDRegex, Md4Regex, Md5Regex, Sha256Regex, Sha384Regex, Sha512Regex, Ripemd128Regex, Ripemd160Regex,
		Tiger128Regex, Tiger160Regex, Tiger192Regex, ASCIIRegex, PrintableASCIIRegex, MultibyteRegex, DataURIRegex, LatitudeRegex, LongitudeRegex, SSNRegex, HostnameRegexRFC952,
		HostnameRegexRFC1123, FqdnRegexRFC1123, BtcAddressRegex, BtcUpperAddressRegexBech32, BtcLowerAddressRegexBech32, EthAddressRegex, EthAddressRegexUpper,
		EthAddressRegexLower, URLEncodedRegex, HTMLEncodedRegex, HTMLRegex, JWTRegex, SplitParamsRegex, BicRegex, SemverRegex, DnsRegexRFC1035Label,
	} {
		xtesting.NotBlankString(t, re.String())
	}
}
