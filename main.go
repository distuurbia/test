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

    var number string
    var letter string
    var text string

    for _, row := range records {
        if row[0] != "" {
            number = string(row[0][0])
        }

        if row[1] != "" {
            letter = string(row[1][0])
        } else {
            letter = ""
        }
        

        
        if row[2] != "" {
            text = strings.Join(row[2:], ",")
        }
        
        if letter != "" && text != "" {
            writer.Write([]string{number, letter, text})
        }
    }
}