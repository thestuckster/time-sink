package internal

import (
	"github.com/fluentassert/verify"
	"regexp"
	"testing"
	"time"
)

func TestToStandardDateFormat(t *testing.T) {
	now := time.Now()
	formattedDate := ToStandardDateFormat(now)

	dateReg := regexp.MustCompile("`^\\d{4}/\\d{2}/\\d{2}$`\n")
	verify.True(dateReg.MatchString(formattedDate))
}

func TestFormatTimeToHHMMSS(t *testing.T) {

	now := time.Now()
	formattedTimestamp := FormatTimeToHHMMSS(now)

	timestampReg := regexp.MustCompile("`^([01]\\d|2[0-3]):[0-5]\\d:[0-5]\\d$`\n")
	verify.True(timestampReg.MatchString(formattedTimestamp))
}

func TestParseHHMMSS(t *testing.T) {
	now := time.Now()
	timeStamp := FormatTimeToHHMMSS(now)
	parsedTime := ParseHHMMSS(timeStamp)
	verify.True(parsedTime == now)

}

func TestParseAmericanDateFormat(t *testing.T) {
	now := time.Now()
	standardDate := ToStandardDateFormat(now)
	parsedTime := ParseStandardDateFormat(standardDate)
	verify.True(parsedTime == now)
}
