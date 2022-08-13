package day3

import (
	"fmt"
	"testing"
	"time"
)

func TestTime1(t *testing.T) {
	fmt.Println(time.Now())
}

func TestTimeParse(t *testing.T) {
	d1, err := time.Parse("2006-01-02 15:04:05", "2021-06-18 12:12:12")
	fmt.Println(d1, err)
}

func TestTime2(t *testing.T) {
	d1, err := time.Parse("2006-01-02 15:04:05", "2021-06-18 12:12:12")
	fmt.Println(d1, err)

	now := time.Now()
	fmt.Println(now.After(d1))
	fmt.Println(now.Before(d1))
	fmt.Println(now.Equal(d1))
}

func TestTS(t *testing.T) {
	now := time.Now()
	ts := now.Unix()
	fmt.Println(ts)

	fmt.Println(time.Unix(ts, 0))
}

func TestTimeAdd(t *testing.T) {
	now := time.Now()
	after1 := now.Add(time.Hour * 24)
	fmt.Println(after1)
}

func TestTimeSub(t *testing.T) {
	now := time.Now()
	after1 := now.Add(time.Hour * 24)
	fmt.Println(now.Sub(after1))
}

func TestTimeTruncate(t *testing.T) {
	now := time.Now()
	fmt.Println(now.Truncate(1 * time.Hour))
}

func TestTimeTimer1(t *testing.T) {
	stop := false
	time.AfterFunc(5*time.Second, func() {
		stop = true
	})

	for {
		if stop {
			fmt.Println("exit")
			break
		}

		time.Sleep(1 * time.Second)
	}
}

func TestTimeTimer2(t *testing.T) {
	m := time.NewTimer(5 * time.Second)
	fmt.Println(<-m.C)
	fmt.Println("exit")
}

func TestTicker(t *testing.T) {
	tk := time.NewTicker(2 * time.Second)

	count := 1
	for {
		if count > 2 {
			tk.Stop()
			break
		}
		fmt.Println(<-tk.C)
		count++
	}
}
