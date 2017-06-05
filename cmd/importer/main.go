package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/danjac/sightings/config"
	"github.com/danjac/sightings/models"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var columns = map[int]string{
	0: "id",
	1: "occurred_at",
	2: "reported_at",
	3: "location",
	4: "shape",
	5: "duration",
	6: "description",
	7: "latitude",
	8: "longitude",
}

var entities = map[string]string{
	"&apos;": "'",
	"&quot;": "\"",
}

const dateFormat = "20060102"

func parseCoord(value string) (float64, error) {
	return strconv.ParseFloat(strings.Replace(value, "\"", "", -1), 64)
}

func makeSighting(record map[string]string) (*models.Sighting, error) {

	var (
		occurredAt time.Time
		reportedAt time.Time
		latitude   float64
		longitude  float64
		err        error
	)

	if occurredAt, err = time.Parse(dateFormat, record["occurred_at"]); err != nil {
		return nil, err
	}

	if reportedAt, err = time.Parse(dateFormat, record["reported_at"]); err != nil {
		return nil, err
	}

	if latitude, err = parseCoord(record["latitude"]); err != nil {
		return nil, err
	}

	if longitude, err = parseCoord(record["longitude"]); err != nil {
		return nil, err
	}

	description := record["description"]

	for k, v := range entities {
		description = strings.Replace(description, k, v, -1)
	}

	s := &models.Sighting{
		OccurredAt:  occurredAt,
		ReportedAt:  reportedAt,
		Latitude:    latitude,
		Longitude:   longitude,
		Description: description,
		Shape:       record["shape"],
		Duration:    record["duration"],
		Location:    record["location"],
	}

	return s, nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please include the filename")
	}

	filename := os.Args[1]

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	cfg, err := config.Configure()

	defer cfg.Close()

	r := csv.NewReader(bufio.NewReader(f))
	r.Comma = '\t'
	r.LazyQuotes = true

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if len(record) < 8 {
			continue
		}

		row := make(map[string]string)

		for value := range record[:9] {
			row[columns[value]] = record[value]
		}

		s, err := makeSighting(row)
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}

		if err := cfg.Repo.Insert(s); err != nil {
			fmt.Printf("%v\n", err)
			continue
		}

		fmt.Printf("%v\n", s)
		fmt.Println("-----------------------------------------")
	}

}
