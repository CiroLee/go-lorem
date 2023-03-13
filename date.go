package lorem

import "time"

type multiDateStruct struct {
	zh string
	en [2]string
}

var weekMap = [7]multiDateStruct{
	{
		zh: "一",
		en: [2]string{"Monday", "Mon."},
	},
	{
		zh: "二",
		en: [2]string{"Tuesday", "Tue."},
	},
	{
		zh: "三",
		en: [2]string{"Wednesday", "Wed."},
	},
	{
		zh: "四",
		en: [2]string{"Thursday", "Thur."},
	},
	{
		zh: "五",
		en: [2]string{"Friday", "Fri."},
	},
	{
		zh: "六",
		en: [2]string{"Saturday", "Sat."},
	},
	{
		zh: "日",
		en: [2]string{"Sunday", "Sun."},
	},
}

var monthMap = [12]multiDateStruct{
	{
		zh: "一月",
		en: [2]string{"January", "jan."},
	},
	{
		zh: "二月",
		en: [2]string{"February", "Feb."},
	},
	{
		zh: "三月",
		en: [2]string{"March", "Mar."},
	},
	{
		zh: "四月",
		en: [2]string{"April", "Apr."},
	},
	{
		zh: "五月",
		en: [2]string{"May", "May."},
	},
	{
		zh: "六月",
		en: [2]string{"June", "Jun."},
	},
	{
		zh: "七月",
		en: [2]string{"July", "Jul."},
	},
	{
		zh: "八月",
		en: [2]string{"August", "Aug."},
	},
	{
		zh: "九月",
		en: [2]string{"September", "Sept."},
	},
	{
		zh: "十月",
		en: [2]string{"October", "Oct."},
	},
	{
		zh: "十一月",
		en: [2]string{"November", "Nov."},
	},
	{
		zh: "十二月",
		en: [2]string{"December", "Dec."},
	},
}

func dateRange(from, to time.Time) time.Time {
	end, _ := randomInteger(int(from.UnixNano()), int(to.UnixNano()))
	return time.Unix(0, int64(end)).UTC()
}

// return a random timestamp support second millisecond and nanosecond via format param
func Timestamp(from, to time.Time, format string) int {
	if format == "nano" {
		return int(dateRange(from, to).UnixNano())
	} else if format == "milli" {
		return int(dateRange(from, to).UnixNano() / 1e3)
	}
	return int(dateRange(from, to).Unix())
}

// return a random UTC date
func Date(from, to time.Time) time.Time {
	return dateRange(from, to)
}

// return a random week support Chinese and English
func Week(lang string, abbr bool) string {
	index, _ := randomInteger(0, 6)
	if lang == "en" {
		if abbr {
			return weekMap[index].en[1]
		}
		return weekMap[index].en[0]
	}
	return weekMap[index].zh
}

func Month(lang string, abbr bool) string {
	index, _ := randomInteger(0, 11)
	if lang == "en" {
		if abbr {
			return monthMap[index].en[1]
		}
		return monthMap[index].en[0]
	}
	return monthMap[index].zh
}
