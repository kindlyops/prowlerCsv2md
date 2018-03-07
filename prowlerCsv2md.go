// prowlerCsv2md - Convert a csv to md
// prowlerCsv2md.go: The main app.
//
// Author: Christopher Mundus <chris@kindlyops.com>
// PROFILE ACCOUNT_NUM REGION TITLE_ID RESULT SCORED LEVEL TITLE_TEXT NOTES

package main

import (
        "bufio"
        "strconv"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"time"
	"os"
	"io"
)

type Report struct {
	Profile string
	Account string
	Region  string
	TitleID	string
	Result	string
	Scored	string
	Level	string
	Title	string
	Notes	string
}

func main() {
	var header string
	header = "| Category   | Result                                                                              |\n" +
	"|-------------|-------------------------------------------------------------------------------------|\n"
	filePtr := flag.String("input", "", "Input csv file to convert")
	namePtr := flag.String("name", "", "Name of the report to process")
	flag.Parse()
	csvInput := *filePtr
	// Get csv file to convert
	if _, err := os.Stat(csvInput); err == nil {
		data, err := os.Open(csvInput)
		if err != nil {
			panic(err)
		}
        	r := csv.NewReader(bufio.NewReader(data))
		var report []Report
		for {
			line, error := r.Read()
			if error == io.EOF {
			    break
			} else if error != nil {
			    log.Fatal(error)
			}
			report = append(report, Report{
				Profile:	line[0],
				Account:  	line[1],
				Region:  	line[2],
				TitleID:	line[3],
				Result: 	line[4],
				Scored: 	line[5],
				Level: 		line[6],
				Title: 		line[7],
				Notes: 		line[8],
			})
		}
		fmt.Println("New Report")
		var outputFile string
		now := time.Now()
		outputFile = *namePtr + "-" + now.Format("01-02-2006") +
					      strconv.Itoa(now.Hour()) +
					      strconv.Itoa(now.Minute()) +
					      strconv.Itoa(now.Second()) + ".md"
		f, err := os.Create(outputFile)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		var body string
		for i := 1; i < len(report); i++ {
			body += header
			body += "| TITLE  | " + report[i].Title +	" |\n"
			body += "| TITLE_ID  | " + report[i].TitleID +	" |\n"
			body += "| RESULT  | " + report[i].Result +	" |\n"
			body += "| SCORED  | " + report[i].Scored +	" |\n"
			body += "| LEVEL  | " + report[i].Level +	" |\n"
			body += "| NOTES  | " + report[i].Notes	+       " |\n"
			body += "\n\n"
		}
		var info string
		info = "^[**Profile:** " + report[1].Profile + "\n"
		info += "**Account ID:** " + report[1].Account + "\n"
		info += "**Region:** " + report[1].Region + "]\n\n"

		n4, err := f.WriteString(info+body)
		fmt.Printf("wrote %d bytes\n", n4)
		f.Sync()
	}

}
