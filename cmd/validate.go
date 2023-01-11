package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/devops-kung-fu/common/util"
	"github.com/spf13/cobra"

	"github.com/devops-kung-fu/meadow/lib"
)

var (
	validateCmd = &cobra.Command{
		Use:     "validate",
		Example: "  meadow validate main.glade",
		Short:   "Checks that a glade file has all of the annotations needed to be rendered to valid go.",
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				util.PrintErr(errors.New("please provide a glade file"))
				fmt.Println()
				_ = cmd.Usage()
				os.Exit(1)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			util.PrintInfo("Validating:", args[0])
			fmt.Println()
			issues := lib.ValidateFile(Afs, args[0])
			if len(issues) == 0 {
				util.PrintSuccess(args[0], "is a valid and properly formatted glade file")
				util.PrintSuccess("Done!")
			} else {
				for _, issue := range issues {
					util.PrintErr(errors.New(issue.Description))
				}
				fmt.Println()
				util.PrintErr(errors.New("errors exist and the file may not generate valid go syntax"))
			}
			log.Println("Finished")
		},
	}
)

func init() {
	rootCmd.AddCommand(validateCmd)
}
