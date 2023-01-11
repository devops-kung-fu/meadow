package lib

import (
	"encoding/xml"
	"path/filepath"

	"github.com/spf13/afero"

	"github.com/devops-kung-fu/meadow/models"
)

// Load loads a source file in glade format into a struct
func Load(afs *afero.Afero, sourcefile string) (models.Glade, error) {

	var glade models.Glade

	sourcefile, err := filepath.Abs(sourcefile)
	if err != nil {
		return glade, err
	}

	b, _ := afs.ReadFile(sourcefile)
	err = xml.Unmarshal(b, &glade)

	return glade, err
}
