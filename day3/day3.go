package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type gridT struct {
	x int
	y int
}

type rectT struct {
	x int
	y int
	w int
	h int
}

func main() {

	coords, _ := fileToLines("day3data.txt")
	var coordStr string
	cutOuts := getCutOuts(coords)
	var wasted int
	var found bool
	var squaresMap map[string]bool
	squaresMap = make(map[string]bool)

	for i := range cutOuts {
		x := cutOuts[i].x
		y := cutOuts[i].y
		for xw := 0; xw < cutOuts[i].w-1; xw++ {
			for yh := 0; yh < cutOuts[i].w-1; yh++ {
				coordStr = fmt.Sprintf("%d:%d", x+xw, y+yh)
				if i > 0 {
					if _, ok := squaresMap[coordStr]; ok {
						if squaresMap[coordStr] == false {
							wasted++
							found = true
							delete(squaresMap, coordStr)
							squaresMap[coordStr] = true
							fmt.Printf("%4d - %7s - Wasted square inches: %d\n", i, coordStr, wasted)
						}
					}
				}
				if !found {
					squaresMap[coordStr] = false
					found = false
				}
			}
		}
	}
	fmt.Printf("Wasted in total: %d\n", wasted)

	/*
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
	*/
}

func getCutOuts(in []string) []rectT {
	var out []rectT
	var outTemp rectT
	for i := range in {
		op1 := strings.Split(in[i], "@")
		op2 := strings.Replace(op1[1], " ", "", -1)
		coords := strings.Split(op2, ":")
		xy := strings.Split(coords[0], ",")
		wl := strings.Split(coords[1], "x")
		outTemp.x, _ = strconv.Atoi(xy[0])
		outTemp.y, _ = strconv.Atoi(xy[1])
		outTemp.w, _ = strconv.Atoi(wl[0])
		outTemp.h, _ = strconv.Atoi(wl[1])
		out = append(out, outTemp)
	}
	return out
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
