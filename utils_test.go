package utils

import (
	"testing"
	"time"
)

func TestShanghaiDatetime(t *testing.T) {
	str := "2019-11-01 10:25:15"
	loc, _ := time.LoadLocation("Asia/Shanghai")
	want := time.Date(2019, time.November, 1, 10, 25, 15, 0, loc)

	got, err := ShanghaiDatetime(str)
	if err != nil {
		t.Fatal(err)
	}

	if !got.Equal(want) {
		t.Errorf("wanted:%v, got: %v", want, *got)
	}
}

func TestShanghaiDatetimeWithMillisecond(t *testing.T) {
	str := "2019-11-01 10:25:15.012"
	loc, _ := time.LoadLocation("Asia/Shanghai")
	want := time.Date(2019, time.November, 1, 10, 25, 15, 12*1000000, loc)

	got, err := ShanghaiDatetime(str)
	if err != nil {
		t.Fatal(err)
	}

	if !got.Equal(want) {
		t.Errorf("wanted:%v, got: %v", want, *got)
	}
}
