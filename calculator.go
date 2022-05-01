package main

import (
	"application"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	const LAPTOP string = "LAPTOP"
	cntRec := 0
	ver := flag.Bool("v", false, "version info")
	input := flag.String("i", "sample-small.csv", "Name of sample input file")
	appid := flag.String("a", "374", "Application ID")
	showRecsInfo := flag.Bool("s", false, "by default, after the calculation,\nonly total number of valid records and\ntotal number of applications required will be shown,\nshould also show the records' info on screen? (default: false)")
	logInfo := flag.Bool("l", false, "by default, no log. (default: false)")

	flag.Parse()

	if *ver {
		fmt.Println("Application Number of Copy(ies) Calculator ver 1.1")
		return
	}
	// Prepare log file, if nedded
	if *logInfo {
		utc := time.Now().UTC()
		fileName := fmt.Sprint("log",
			utc.Year(), "-", utc.Month(), "-", utc.Day(), "_",
			utc.Hour(), "-", utc.Minute(), "-", utc.Second(),
			".txt")
		// If the file doesn't exist, create it or append to the file
		f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		log.SetOutput(f)
	}

	// Prepare Outter set for categorising records
	so := application.NewUsage()

	recs := make(chan []string)
	recs = RetrieveCSV(*input)

	// Loop through lines & turn into object
	// Only retrieve records with specified (or default) AppplicationID
	for rec := range recs {
		computerID := rec[0]
		userID := rec[1]
		applicationID := rec[2]
		computerType := rec[3]
		comment := rec[4]
		if *showRecsInfo {
			fields := fmt.Sprint(
				computerID + "\t" + userID + "\t" + applicationID +
					"\t" + computerType + "\t" + comment)
			fmt.Println(fields)
			if *logInfo {
				log.Println(fields)
			}
		}
		if !ValidRec(computerID, userID, applicationID, computerType) {
			if *showRecsInfo {
				validationResult := fmt.Sprint("Any empty ComputerID, UserID, ApplicationID or CompueterType invalidates pertaining record for calculation.")
				fmt.Println(validationResult)
				if *logInfo {
					log.Println(validationResult)
				}
			}
			continue
		}
		if applicationID == *appid {
			cntRec++
			theComputerType := strings.ToUpper(computerType)
			lowerCasecomputerTypeName := strings.ToLower(computerType)
			var ctc application.ComputerTypeCount

			if so.Contains(userID) {
				ctc1, _ := so.GetVal(userID)
				// A very tricky "Type Assertion" to do here -- from interface{} to expected type
				ctc.DesktopCount = ctc1.(application.ComputerTypeCount).DesktopCount
				ctc.LaptopCount = ctc1.(application.ComputerTypeCount).LaptopCount
			}

			if theComputerType == LAPTOP {
				ctc.LaptopCount = ctc.LaptopCount + 1
			} else {
				if !(lowerCasecomputerTypeName == computerType) {
					ctc.DesktopCount = ctc.DesktopCount + 1
				}
			}
			so.Add(userID, ctc)
		}
	}

	summary := fmt.Sprint(
		"\nSummary -",
		"\n\tInput Data File: ", *input,
		"\n\tApplicationID of valid records to be processed: ", *appid,
		"\n\tTotal number of valid records processed: ", cntRec,
		"\n\tTotal number of application copy required: ", so.CalculateCopyNumber(),
		"\n\tTotal execution time: ", time.Since(start))
	fmt.Println(summary)
	if *logInfo {
		log.Println(summary)
	}
}

func ValidRec(
	computerID, userID, applicationID, computerType string) bool {
	if len(computerID) == 0 || len(userID) == 0 || len(applicationID) == 0 || len(computerType) == 0 {
		return false
	} else {
		return true
	}
}

/**
 * Instead of reading every records as a whole into memory,
 * RetrieveCSV read 10 records at a time and communicate main program through channel
 */
func RetrieveCSV(filename string) (line chan []string) {
	line = make(chan []string, 50)
	go func() {
		// Open CSV file
		f, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		// Read File into a Variable
		r := csv.NewReader(f)
		if _, err := r.Read(); err != nil { //read header
			log.Fatal(err)
		}
		defer close(line)
		for {
			rec, err := r.Read()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal(err)
			}
			line <- rec
		}
	}()
	return
}
