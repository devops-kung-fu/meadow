package lib

import (
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	afs := &afero.Afero{Fs: afero.NewOsFs()}

	_, err := Load(afs, "test/ui.glade")
	assert.Error(t, err)

	filename, err := filepath.Abs("../test/ui.glade")
	assert.NoError(t, err)

	glade, err := Load(afs, filename)
	assert.NoError(t, err)
	assert.Equal(t, "3.24", glade.Requires.Version)
}
