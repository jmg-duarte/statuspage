package format

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
)

type Writer struct {
	io.Writer
	Format
}

func NewWriter(writer io.Writer, format Format) Writer {
	return Writer{Writer: writer, Format: format}
}

func (w Writer) Write(data interface{}) {
	switch w.Format {
	case JSON:
		w.writeJson(data)
	case CSV:
		switch data.(type) {
		case [][]string:
			w.writeCSV(data.([][]string))
		default:
			log.Fatal("unsupported format")
		}
	default:
		log.Fatal("unsupported format")
	}
}

func (w Writer) writeJson(data interface{}) {
	b, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	_, err = w.Writer.Write(b)
	if err != nil {
		log.Fatal(err)
	}
}

func (w Writer) writeCSV(record [][]string) {
	csvW := csv.NewWriter(w.Writer)
	err := csvW.WriteAll(record)
	if err != nil {
		log.Fatal(err)
	}
	csvW.Flush()
	err = csvW.Error()
	if err != nil {
		log.Fatal(err)
	}
}
