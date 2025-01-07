package internal

import (
	"github.com/fluentassert/verify"
	"regexp"
	"testing"
	"time"
)

func TestToRealDate(t *testing.T) {

	now := time.Now()
	timestamp := float64(now.UnixNano() / 1e9)

	verify.True(ToRealDate(now) == timestamp)
}

func TestFromRealDate(t *testing.T) {

	now := time.Now()
	timestamp := float64(now.UnixNano() / 1e9)
	verify.True(FromRealDate(timestamp) == now)
}

func TestToStandardDateFormat(t *testing.T) {
	now := time.Now()
	formattedDate := ToStandardDateFormat(now)

	dateReg := regexp.MustCompile("`^\\d{4}/\\d{2}/\\d{2}$`\n")
	verify.True(dateReg.MatchString(formattedDate))
}
