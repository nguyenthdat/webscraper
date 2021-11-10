package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	fname := "data.csv"
	file, err := os.Create(fname)
	if err != nil {
		log.Fatal("Failed to create file, error: %q", err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
}
