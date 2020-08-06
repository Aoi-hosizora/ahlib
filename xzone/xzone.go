package xzone

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xnumber"
	"regexp"
	"time"
)

// Parse Timezone string to time.Location. Format: `^[+-][0-9]{1,2}([0-9]{1,2})?$`
func ParseTimeZone(zone string) (*time.Location, error) {
	// reference: https://jp.cybozu.help/general/ja/admin/list_systemadmin/list_system_time/timezone.html

	regex, err := regexp.Compile(`^([+-])([0-9]{1,2})(?::([0-9]{1,2}))?$`)
	if err != nil {
		return nil, err
	}

	wrongFmtErr := fmt.Errorf("timezone string has a wrong format")
	ok := regex.Match([]byte(zone))
	if !ok {
		return nil, wrongFmtErr
	}

	matches := regex.FindAllStringSubmatch(zone, 1)
	if len(matches) == 0 || len(matches[0][1:]) < 3 {
		return nil, wrongFmtErr
	}
	group := matches[0][1:]

	signStr := group[0]
	hourStr := group[1]
	minuteStr := group[2]
	if signStr != "+" && signStr != "-" {
		return nil, wrongFmtErr
	}
	if minuteStr == "" {
		minuteStr = "0"
	}

	sign := +1
	if signStr == "-" {
		sign = -1
	}
	hour, err1 := xnumber.ParseInt(hourStr, 10)
	minute, err2 := xnumber.ParseInt(minuteStr, 10)
	if err1 != nil || err2 != nil {
		return nil, wrongFmtErr
	}

	name := fmt.Sprintf("UTC%s%02d:%02d", signStr, hour, minute)
	offset := sign * (hour*3600 + minute*60)
	return time.FixedZone(name, offset), nil
}
