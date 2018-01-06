package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"github.com/EGR203/go-utils/textfmt"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: ./showcsv file.scv")
		os.Exit(1)
	}
	path := os.Args[1]

	file, _ := os.Open(path)
	csv_file := csv.NewReader(file)
	sizes := getSizes(csv_file)
	file.Close()

	file, _ = os.Open(path)
	csv_file = csv.NewReader(file)
	PrintCsv(csv_file, sizes)
	file.Close()

}
func PrintCsv(r *csv.Reader, sizes []int) {
	one_str_lenth := calcSizes(sizes) + len(sizes) + 1
	salute_str := textfmt.MultiString("-", one_str_lenth)
	devide_str := "|"
	for i, v := range sizes {
		if i == len(sizes)-1 {
			devide_str += textfmt.MultiString("-", v) + "|"

		} else {
			devide_str += textfmt.MultiString("-", v) + "+"
		}
	}
	fmt.Println(salute_str)
	first := true
	for line, err := r.Read(); err == nil; line, err = r.Read() {
		if first {
			first = false
		} else {
			fmt.Println(devide_str)
		}
		fmt.Print("|")
		for i, s := range line {
			fmt.Print(textfmt.TabText(s, sizes[i]) + "|")
		}
		fmt.Print("\n")
	}
	fmt.Println(salute_str)
}

func calcSizes(sizes []int) int {
	var res int
	for _, v := range sizes {
		res += v
	}
	return res
}


func getSizes(r *csv.Reader) []int {
	first_line, _ := r.Read()
	sizes := make([]int, len(first_line))
	for line, err := first_line, error(nil); err == nil; line, err = r.Read() {
		for i, v := range line {
			if sizes[i] < len([]rune(v)) {
				sizes[i] = len([]rune(v)) + 3
			}
		}
	}
	return sizes
}
