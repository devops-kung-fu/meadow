package lib

import (
	"fmt"

	"github.com/spf13/afero"

	"github.com/devops-kung-fu/meadow/models"
)

// Validate checks that a glade file has all of the annotations needed to be rendered to valid go
func Validate(glade models.Glade) (issues []models.Issue) {
	if glade.Requires.Version != "3.24" {
		issues = append(issues, models.Issue{
			Description: fmt.Sprintf("unsupported Gtk version detected (%s)", glade.Requires.Version),
		})
	}
	for _, object := range glade.Object {
		issues = append(issues, validateObject(object)...)
	}
	return
}

// ValidateFile checks to see if a source file is valid and has all of the correct annotations needed to be rendered to valid go
func ValidateFile(afs *afero.Afero, sourcefile string) (issues []models.Issue) {
	glade, err := Load(afs, sourcefile)
	if err != nil {
		return
	}
	return Validate(glade)
}

func validateObject(object models.Object) (issues []models.Issue) {
	if object.Class == "GtkImageMenuItem" {
		issues = append(issues, models.Issue{
			Description: "GtkImageMenuItem object found which is not supported in gotk3",
		})
	}
	if object.ID == "" {
		issues = append(issues, models.Issue{
			Description: fmt.Sprintf("%s object found with no identifier (Critical)", object.Class),
		})
	}
	for _, child := range object.Child {
		issues = append(issues, validateChild(child)...)
	}
	return
}

func validateChild(child models.Child) (issues []models.Issue) {
	for _, object := range child.Object {
		issues = append(issues, validateObject(object)...)
	}
	return
}
