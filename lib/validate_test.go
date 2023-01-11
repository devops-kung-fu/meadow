package lib

import (
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"

	"github.com/devops-kung-fu/meadow/models"
)

func TestValidate_Bad(t *testing.T) {
	afs := &afero.Afero{Fs: afero.NewOsFs()}

	filename, err := filepath.Abs("../test/bad.glade")
	assert.NoError(t, err)

	glade, err := Load(afs, filename)
	assert.NoError(t, err)

	glade.Requires.Version = "1" //will fail

	issues := Validate(glade)
	assert.Len(t, issues, 2)
}

func TestValidate_Good(t *testing.T) {
	afs := &afero.Afero{Fs: afero.NewOsFs()}

	filename, err := filepath.Abs("../test/ui.glade")
	assert.NoError(t, err)

	glade, err := Load(afs, filename)
	assert.NoError(t, err)

	issues := Validate(glade)
	assert.Len(t, issues, 0)
}

func TestValidateFile_Good(t *testing.T) {
	afs := &afero.Afero{Fs: afero.NewOsFs()}

	filename, err := filepath.Abs("../test/ui.glade")
	assert.NoError(t, err)

	issues := ValidateFile(afs, filename)

	assert.Len(t, issues, 0)
}

func TestValidateFile_Bad(t *testing.T) {
	afs := &afero.Afero{Fs: afero.NewOsFs()}

	filename, err := filepath.Abs("../test/bad.glade")
	assert.NoError(t, err)

	issues := ValidateFile(afs, filename)

	assert.Len(t, issues, 1)
}

func TestValidateFile(t *testing.T) {
	afs := &afero.Afero{Fs: afero.NewOsFs()}

	filename, err := filepath.Abs(":\\//,.<>")
	assert.NoError(t, err)

	issues := ValidateFile(afs, filename)

	assert.Len(t, issues, 0)
}

func Test_validateObject(t *testing.T) {
	object := models.Object{
		Class: "GtkWhatever",
	}
	issues := validateObject(object)
	assert.Len(t, issues, 1)
}
