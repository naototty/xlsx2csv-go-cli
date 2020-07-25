// Copyright 2011-2015, The xlsx2csv Authors.
// All rights reserved.
// For details, see the LICENSE file.

package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"

	"github.com/tealeg/xlsx"
)

func generateCSVFromXLSXFile(w io.Writer, excelFileName string, sheetIndex int, csvOpts csvOptSetter) error {
	xlFile, error := xlsx.OpenFile(excelFileName)
	if error != nil {
		return error
	}
	sheetLen := len(xlFile.Sheets)
	switch {
	case sheetLen == 0:
		return errors.New("This XLSX file contains no sheets.")
	case sheetIndex >= sheetLen:
		return fmt.Errorf("No sheet %d available, please select a sheet between 0 and %d\n", sheetIndex, sheetLen-1)
	}
	cw := csv.NewWriter(w)
	if csvOpts != nil {
		csvOpts(cw)
	}
	re3 := regexp.MustCompile("\n")
	sheet := xlFile.Sheets[sheetIndex]
	var vals []string
	for _, row := range sheet.Rows {
		if row != nil {
			vals = vals[:0]
			for _, cell := range row.Cells {
				str, err := cell.FormattedValue()
				if err != nil {
					vals = append(vals, err.Error())
				} else {
					*str = re3.ReplaceAllString(*str, "_")
					vals = append(vals, str)
				}
			}
			if err := cw.Write(vals); err != nil {
				return err
			}
		}
	}
	cw.Flush()
	return cw.Error()
}

type csvOptSetter func(*csv.Writer)

func main() {
	var (
		outFile    = flag.String("o", "-", "filename to output to. -=stdout")
		sheetIndex = flag.Int("i", 0, "Index of sheet to convert, zero based")
		delimiter  = flag.String("d", ",", "Delimiter to use between fields(TAB is tab)")
	)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `%s
	dumps the given xlsx file's chosen sheet as a CSV,
	with the specified delimiter, into the specified output.

Usage:
	%s [flags] <xlsx-to-be-read>
`, os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
	re1 := regexp.MustCompile("\\t")
	re2 := regexp.MustCompile("TAB")
	if *delimiter == "\t" {
		// re1.ReplaceAllLiteralString(*delimiter, "\u0009")
		// *delimiter = re1.ReplaceAllString(*delimiter, "\u0009")
		*delimiter = re1.ReplaceAllString(*delimiter, "\t")
	}
	if *delimiter == "TAB" {
		// re2.ReplaceAllLiteralString(*delimiter, "\t")
		*delimiter = re2.ReplaceAllString(*delimiter, "\t")
	}
	out := os.Stdout
	if !(*outFile == "" || *outFile == "-") {
		var err error
		if out, err = os.Create(*outFile); err != nil {
			log.Fatal(err)
		}
	}
	defer func() {
		if closeErr := out.Close(); closeErr != nil {
			log.Fatal(closeErr)
		}
	}()

	if err := generateCSVFromXLSXFile(out, flag.Arg(0), *sheetIndex,
		func(cw *csv.Writer) { cw.Comma = ([]rune(*delimiter))[0] },
	); err != nil {
		log.Fatal(err)
	}
}
