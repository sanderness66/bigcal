// BIGCAL -- like cal(1), but bigger
//
// SvM 21-JAN-2021

package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

const sep = "+----------+----------+----------+----------+----------+----------+----------+"
const hdr = "|Sunday    |Monday    |Tuesday   |Wednesday |Thursday  |Friday    |Saturday  |"
const fil = "|          |          |          |          |          |          |          |"

func main() {
	var mon, year int
	do_year := false

	switch len(os.Args) {
	case 3:
		mon, _ = strconv.Atoi(os.Args[1])
		if mon < 1 || mon > 12 {
			fmt.Printf("month %d is silly\n", mon)
			os.Exit(1)
		}
		year, _ = strconv.Atoi(os.Args[2])
	case 2:
		do_year = true
		year, _ = strconv.Atoi(os.Args[1])
	case 1:
		t := time.Now()
		mon = int(t.Month())
		year = int(t.Year())
	}

	fmt.Println(sep)

	if do_year {
		for mon = 1; mon < 13; mon++ {
			do_month(mon, year)
		}
	} else {
		do_month(mon, year)
	}
}

func do_month(mon int, year int) {
	// which weekday is the first day of this month?
	t := time.Date(year, time.Month(mon), 1, 0, 0, 0, 0, time.UTC)
	day1 := int(t.Weekday())

	// and how many days are there in this month? It's "day 0" of
	// the following month!
	t = time.Date(year, time.Month(mon+1), 0, 0, 0, 0, 0, time.UTC)
	days := int(t.Day())

	// mtable[] holds all the days of the month in position, where
	// mtable[1] = Sunday, etc. Think of mtable[] as all the
	// output of a cal(1) invocation merged into one long line. A
	// month might cover (part of) 6 weeks, so that's 7*6
	// positions. And since the author gets confused by 0-based
	// arrays with 1-based dates, we use mtable[] as if it were
	// 1-based too, so add one extra for that as well.
	var mtable [(7 * 6) + 1]int

	// this puts the day of month in the appropriate place
	for i := 1; i <= days; i++ {
		mtable[i+day1] = i
	}

	// print a nice header
	fmt.Println(hdr)
	fmt.Println(sep)

	// put each week's worth of mtable[] in week[] (1-based again)
	// and pretty-print it
	var week [8]int

	for i := 1; i <= (7 * 6); i++ {
		// we want to fill week[1..7], not week[1..6 plus 0]
		j := i % 7
		if j == 0 {
			j = 7
		}
		week[j] = mtable[i]

		// print out what we have after each 7th day
		if i%7 == 0 {
			// but don't print empty lines
			if sum(week[1:]) != 0 {
				pprint(mon, week)
			}
		}
	}
}

func pprint(mon int, week [8]int) {
	m := []string{"", "JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC"}

	for _, d := range week[1:] {
		switch d {
		case 0:
			fmt.Printf("|          ")
		case 1:
			fmt.Printf("|%3s    %2d ", m[mon], d)
		default:
			fmt.Printf("|       %2d ", d)
		}
	}
	fmt.Println("|")

	fmt.Println(fil)
	fmt.Println(fil)
	fmt.Println(sep)
}

func sum(x []int) int {
	ret := 0
	for _, val := range x {
		ret += val
	}
	return ret
}
