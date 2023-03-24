package utilities

import (
	"testing"
)

var dateTests = []TestSuite[uint64, string]{
	{Args: 1110000000 * 1000, Want: "05 Mar 2005"},
	{Args: 1500000000 * 1000, Want: "14 Jul 2017"},
	{Args: 1679676572 * 1000, Want: "24 Mar 2023"},
}

var durationTests = []TestSuite[uint64, string]{
	{Args: 66, Want: "01:06"},
	{Args: 2, Want: "00:02"},
	{Args: 3601, Want: "01:00:01"},
}

var humanReadableTests = []TestSuite[int64, string]{
	{Args: 1000, Want: "1.0 k"},
	{Args: 12000, Want: "12.0 k"},
	{Args: 3_300_000, Want: "3.3 M"},
}

func TestFormatDate(t *testing.T) {
	Test(dateTests, func(args uint64) string {
		return FormatDate(args)
	}, func(got, want string) {
		t.Errorf("got %s, want %s", got, want)
	})
}


func TestFormatDuration(t *testing.T) {
	Test(durationTests, func(args uint64) string {
		return FormatDuration(args)
	}, func(got, want string) {
		t.Errorf("got %s, want %s", got, want)
	})
}


func TestFormatHumanReadable(t *testing.T) {
	Test(humanReadableTests, func(args int64) string {
		return FormatHumanReadable(args)
	}, func(got, want string) {
		t.Errorf("got %s, want %s", got, want)
	})
}
