package c9r

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var Location *time.Location

func init() {
	Location, _ = time.LoadLocation("Asia/Shanghai")
}

func ParseDate(str string) time.Time {
	// 43分钟前
	// 2小时前
	if strings.Contains(str, "前") {
		str = strings.Replace(strings.Replace(str, "分钟前", "m", -1), "小时前", "h", -1)
		d, err := time.ParseDuration("-" + str)
		if err != nil {
			fmt.Println(err)
			return time.Time{}
		}

		return time.Now().Add(d).UTC().Truncate(time.Minute)
	}

	// 01-01 04:56
	if len(str) == 11 {
		layout := "01-02 15:04 2006 -0700"
		str += " " + strconv.Itoa(time.Now().Year()) + " +0800"
		t, err := time.Parse(layout, str)
		if err != nil {
			fmt.Println(err)
		}
		t = t.UTC()

		if t.After(time.Now()) {
			return t.AddDate(-1, 0, 0)
		}

		return t
	}

	return time.Time{}
}
