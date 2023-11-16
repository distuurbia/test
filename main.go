package main

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

func main() {
    inputFile, err := os.Open("./input.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer inputFile.Close()

    outputFile, err := os.Create("./output.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer outputFile.Close()

    writer := csv.NewWriter(outputFile)
    defer writer.Flush()

    reader := csv.NewReader(inputFile)
    records, err := reader.ReadAll()
    if err != nil {
        log.Fatal(err)
        return
    }

    for _, row := range records {
        number := string(row[0][0])
        letter := string(row[1][0])
        text := strings.Join(row[2:], ",")
        writer.Write([]string{number, letter, text})
    }
}