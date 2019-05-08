package models

import (
	"os"
)

type Utils struct {
}

func (u Utils) WriteOnFile(html string, fileName string) error {

	// open file to write html
	f, err := os.Create("views/" + fileName)
	if err != nil {
		return err
	}

	// write on file
	if _, err = f.WriteString(html); err != nil {
		f.Close()
		return err
	}

	// close file
	if err = f.Close(); err != nil {
		return err
	}
	return nil
}
