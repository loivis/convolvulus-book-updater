package update

import (
	"testing"
	"time"
)

func TestParseDate(t *testing.T) {
	for _, tc := range []struct {
		desc string
		in   string
		out  time.Time
	}{
		{
			desc: "DateHourMinuteInThePast",
			in:   "01-01 12:34",
			out:  time.Date(time.Now().Year(), 1, 1, 4, 34, 0, 0, time.UTC),
		},
		{
			desc: "DateHourMinuteInTheFuture",
			in:   "12-31 23:45",
			out:  time.Date(time.Now().AddDate(-1, 0, 0).Year(), 12, 31, 15, 45, 0, 0, time.UTC),
		},
		{
			desc: "MinutesBack",
			in:   "43分钟前",
			out:  time.Now().Add(-43 * time.Minute).UTC().Truncate(time.Minute),
		},
		{
			desc: "HoursBack",
			in:   "2小时前",
			out:  time.Now().Add(-2 * time.Hour).UTC().Truncate(time.Minute),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			if got, want := ParseDate(tc.in), tc.out; got != want {
				t.Fatalf("got = %q, want %q", got, want)
			}
		})
	}
}
