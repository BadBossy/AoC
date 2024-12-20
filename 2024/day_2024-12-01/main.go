package main

import (
  "bufio"
  "fmt"
  "math"
  "os"
  "sort"
  "strconv"
  "strings"
)

func readInputFile(filename string) ([]int, []int, error) {
  file, err := os.Open(filename)

  if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var left, right []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("invalid line format: %s", line)
		}

		leftValue, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid number in left column: %s", parts[0])
		}
		rightValue, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid number in right column: %s", parts[1])
		}

		left = append(left, leftValue)
		right = append(right, rightValue)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return left, right, nil
}


func calculateTotalDistance(left, right []int) int {
  sort.Ints(left)
  sort.Ints(right)

  totalDistance := 0

  for i := 0; i < len(left); i++ {
    totalDistance += int(math.Abs(float64(left[i] - right[i])))
  }

  return totalDistance
}

func main() {

left, right, err := readInputFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	if len(left) != len(right) {
		fmt.Println("Error: The left and right lists have different lengths.")
		return
	}


  totalDistance := calculateTotalDistance(left, right)
  fmt.Printf("Total Distance is: %d\n", totalDistance)
  
}
