package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	var result int64
	var numString string
	var freqMap map[int64]int64
	freqMap = make(map[int64]int64)

	freqChange, _ := fileToLines("day1data.txt")

	for {
		for i := range freqChange {
			numString = strings.Replace(freqChange[i], "+", "", -1)
			numString = strings.Replace(freqChange[i], "-", "", -1)
			number, _ := strconv.ParseInt(numString, 10, 64)
			if strings.Contains(freqChange[i], "+") {
				result += number
			} else {
				result -= number
			}
			if i > 0 {
				if _, ok := freqMap[result]; ok {
					fmt.Printf("- Frequency %v repeated!\n", result)
					close()
				}
			}
			freqMap[result] = number
			fmt.Printf("%4d : %6s : %6d\n", i, freqChange[i], result)
		}
	}
}

func fileToLines(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return linesFromReader(f)
}

func linesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func close() {
	os.Exit(0)
}
