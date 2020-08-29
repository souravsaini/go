package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	f, err := os.Open("Environmental_Data_Deep_Moor_2015.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rdr := csv.NewReader(f)
	rdr.Comma = '\t'
	rdr.TrimLeadingSpace = true
	rows, err := rdr.ReadAll()
	if err != nil {
		panic(err)
	}

	start := time.Now()
	fmt.Println("Total records", len(rows)-1)
	fmt.Printf("Mean Air Temp: %.2f, Median Air Temp: %.2f\n", mean(rows, 1), median(rows, 1))
	fmt.Printf("Mean Barometric: %.2f, Median Barometric: %.2f\n", mean(rows, 2), median(rows, 2))
	fmt.Printf("Mean Wind Speed: %.2f, Median Wind Speed: %.2f\n", mean(rows, 7), median(rows, 7))
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("Abstracted: %s\n", delta)

}

func mean(rows [][]string, index int) float64 {
	var total float64
	for i, row := range rows {
		if i != 0 {
			val, _ := strconv.ParseFloat(row[index], 64)
			total += val
		}
	}
	return total / float64(len(rows)-1)
}

func median(rows [][]string, index int) float64 {
	var sorted []float64

	// populate the sorted []float64
	for i, row := range rows {
		if i != 0 {
			val, _ := strconv.ParseFloat(row[index], 64)
			sorted = append(sorted, val)
		}
	}

	// Sorting the []float64
	sort.Float64s(sorted)

	//even no. of items
	if len(sorted)%2 == 0 {
		middle := len(sorted) / 2
		higher := sorted[middle]
		lower := sorted[middle-1]
		return (higher + lower) / 2
	}

	//odd no. of items
	middle := len(sorted) / 2
	return sorted[middle]
}
