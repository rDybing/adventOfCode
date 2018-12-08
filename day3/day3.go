package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type gridT struct {
	x int
	y int
}

type rectT struct {
	id int
	x  int
	y  int
	w  int
	h  int
}

func main() {

	coords, _ := fileToLines("day3data.txt")
	cutOuts := getCutOuts(coords)
	var coordStr string
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
}

func getCutOuts(in []string) []rectT {
	var out []rectT
	var ot rectT
	for i := range in {
		_, err := fmt.Sscanf(in[i], "#%d @ %d,%d: %dx%d", &ot.id, &ot.x, &ot.y, &ot.w, &ot.h)
		if err != nil {
			fmt.Printf("could not parse %s : %v\n", in[i], err)
		}
		out = append(out, ot)
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
