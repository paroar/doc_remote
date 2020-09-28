package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var one float64

func main() {

	filePath := flag.String("f", "", "string csv filePath")
	flag.Float64Var(&one, "s", 0.0, "student scores separated by spaces")
	flag.Parse()

	//-f command
	if *filePath != "" {
		records, err := readCsvFile(*filePath)
		if err != nil {
			log.Fatal(err)
		}

		finalRecords := [][]string{}
		var record []string
		for _, r := range records {
			fs, err := parseArrayStrToFloat32(r[1:])
			if err != nil {
				log.Fatal("Couldn't parse Csv %s\n", err)
			} else {
				record = append(r, fmt.Sprint(avg(fs)))
				finalRecords = append(finalRecords, record)
				fmt.Println(record)
			}
		}

		f, err := os.Create("avg_scores.csv")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		w := csv.NewWriter(f)
		w.WriteAll(finalRecords)
		w.Flush()
	}

	//-s command
	if one != 0 {
		firstArg := fmt.Sprint(one)
		tailArgs := flag.Args()
		scores := append([]string{firstArg}, tailArgs...)
		fs, err := parseArrayStrToFloat32(scores)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(avg(fs))
		}
	}

}

func readCsvFile(filePath string) ([][]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("Unable to parse file as CSV: %s\n%s", filePath, err)
	}
	return records, nil
}

func parseArrayStrToFloat32(ss []string) ([]float32, error) {
	fs := []float32{}
	for _, s := range ss {
		f, err := strconv.ParseFloat(s, 32)
		if err != nil {
			return nil, err
		}
		fs = append(fs, float32(f))
	}
	return fs, nil
}

func avg(a []float32) (avg float32) {
	var sum float32 = 0.0
	length := len(a)
	for i := 0; i < length; i++ {
		sum += a[i]
	}
	avg = sum / float32(length)
	return avg
}
