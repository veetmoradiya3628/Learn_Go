package main

import (
	"bufio"
	"fmt"
	"os"
)

func printLastNLines(lines []string, num int) []string {
	var printLastNLines []string
	for i := len(lines) - num; i < len(lines); i++ {
		printLastNLines = append(printLastNLines, lines[i])
	}
	return printLastNLines
}
func main() {
	filepath := "log.txt"
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error in opening file: ", err)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file : ", err)
	}
	printLastNLines := printLastNLines(lines, 5)
	for _, line := range printLastNLines {
		fmt.Println(line)
	}
}
