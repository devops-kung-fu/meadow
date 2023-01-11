package lib

import (
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	afs := &afero.Afero{Fs: afero.NewOsFs()}
	filename, err := filepath.Abs("../test/ui.glade")
	assert.NoError(t, err)

	b, err := afs.ReadFile(filename)
	assert.NoError(t, err)

	mafs := &afero.Afero{Fs: afero.NewMemMapFs()}
	filename, _ = filepath.Abs("test.glade")
	assert.NoError(t, err)

	err = mafs.WriteFile(filename, b, 0777)
	assert.NoError(t, err)

	_, err = Generate(mafs, "NOFILE", "ui", "glade.go", "ui")
	assert.Error(t, err)

	r, err := Generate(mafs, "test.glade", "ui", "glade.go", "ui")
	assert.NoError(t, err)
	assert.NotEqual(t, "", r)

	testFile, err := filepath.Abs("ui/glade.go")
	assert.NoError(t, err)

	test, err := mafs.ReadFile(testFile)
	assert.NoError(t, err)
	assert.NotNil(t, test)
}

func Test_renderModel(t *testing.T) {
	afs := &afero.Afero{Fs: afero.NewOsFs()}

	filename, err := filepath.Abs("../test/ui.glade")
	assert.NoError(t, err)

	glade, err := Load(afs, filename)
	assert.NoError(t, err)

	renderModel := buildModel(glade)

	assert.Len(t, renderModel, 4)

}

func Test_genTemplate(t *testing.T) {
	name := "glade"
	template := genTemplate(name)

	assert.NotNil(t, template)
	assert.Equal(t, name, template.Name(), "Template should have a name `glade`")
}
