package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	boxID, _ := fileToLines("day2data.txt")
	//part1(boxID)
	part2(boxID)
}

func part1(boxID []string) {
	var resultTwo, resultThree int

	for i := range boxID {
		getTwo, getThree := getRepeatsFromID(boxID[i])
		if getTwo {
			resultTwo++
		}
		if getThree {
			resultThree++
		}
		fmt.Printf("%3d - %s - twos: %3d threes: %3d\n", i, boxID[i], resultTwo, resultThree)
	}
	checkSum := resultTwo * resultThree
	fmt.Printf("Twos: %3d - Threes: %3d - CheckSum: %d\n", resultTwo, resultThree, checkSum)
}

func part2(boxID []string) {
	result := getCommonIDs(boxID)
	if result != "" {
		fmt.Printf("common letters: %s\n", result)
	}
}

func getCommonIDs(boxID []string) string {
	for i := 0; i < len(boxID[0]); i++ {
		commonIDs := make(map[string]struct{}, len(boxID))
		for _, curID := range boxID {
			curSeq := curID[:i] + curID[i+1:]
			if _, ok := commonIDs[curSeq]; ok {
				return curSeq
			}
			commonIDs[curSeq] = struct{}{}
		}
	}
	return ""
}

func getRepeatsFromID(in string) (bool, bool) {
	var twos bool
	var threes bool
	for i := range in {
		if strings.Count(in, string(in[i])) == 2 {
			twos = true
		}
		if strings.Count(in, string(in[i])) == 3 {
			threes = true
		}
	}
	return twos, threes
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
