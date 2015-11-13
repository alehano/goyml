package goyml

import (
	"encoding/xml"
	"io"
	"os"
)

func ExportToWriter(cat Catalog, w io.Writer, pretty bool) error {
	w.Write([]byte(Header))
	xmlEncoder := xml.NewEncoder(w)
	if pretty {
		xmlEncoder.Indent("", "\t")
	}
	err := xmlEncoder.Encode(cat)
	if err != nil {
		return err
	}
	defer xmlEncoder.Flush()
	return nil
}

func ExportToFile(cat Catalog, filename string, pretty bool) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return ExportToWriter(cat, file, pretty)
}
