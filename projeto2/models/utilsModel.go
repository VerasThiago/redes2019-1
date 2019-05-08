package models

import (
	"os"
)

type Utils struct {
}

func (u Utils) WriteOnFile(html string) error {

	// open file to write html
	f, err := os.Create("views/find.html")
	if err != nil {
		return err
	}

	// write on file
	_, err = f.WriteString(html)
	if err != nil {
		f.Close()
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}
	return nil
}
