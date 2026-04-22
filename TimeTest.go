// Golang program to compare times
package main

import "fmt"

// importing time module
import "time"

// Main function
func oldmain() {

	today := time.Now()
	tomorrow := today.Add(24 * time.Hour)
	sameday := tomorrow.Add(-12 * time.Hour)

	fmt.Printf("Today:\t\t%s\nTomorrow:\t%s\nSameday:\t%s\n\n", today.Format("2006-01-02"), tomorrow.Format("2006-01-02"), sameday.Format("2006-01-02"))

	if today != tomorrow {
		fmt.Println("today is not tomorrow")
	}

	if sameday == today {
		fmt.Println("sameday is today")
	}

	// using Equal function
	if today.Equal(sameday) || ( today.Year() == sameday.Year() && today.YearDay() == sameday.YearDay() ) {
		fmt.Println("today is sameday")
	}

}

func timemain() {
	now := time.Now()
	for i := 0; i < 7; i++ {
		fmt.Println(now.Add(-time.Duration(i)*24*time.Hour).Format("2006-01-02"))
	}
}
