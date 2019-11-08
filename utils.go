package utils

import "time"

func ShanghaiDatetime(dateStr string) (*time.Time, error) {
	timeLayout := "2006-01-02 15:04:05.999"
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t, err := time.ParseInLocation(timeLayout, dateStr, loc)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
