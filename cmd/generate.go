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
	outputPath, outputFileName, pkg string

	// summary, detailed bool
	generateCmd = &cobra.Command{
		Use:     "generate",
		Example: "  meadow generate --output-path ui main.glade",
		Short:   "Generates Go code to connect to GTK components in Glade layouts.",
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				util.PrintErr(errors.New("please provide a glade file"))
				fmt.Println()
				_ = cmd.Usage()
				os.Exit(1)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			util.PrintInfo("Processing:", args[0])
			filename, err := lib.Generate(Afs, args[0], outputPath, outputFileName, pkg)
			if err != nil {
				util.PrintErr(err)
				os.Exit(1)
			}
			util.PrintInfo("Source file written to:", filename)
			util.PrintSuccess("Done!")
			log.Println("Finished")
		},
	}
)

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.PersistentFlags().StringVar(&outputPath, "output-path", "ui", "Output path for generated source file.")
	generateCmd.PersistentFlags().StringVar(&outputFileName, "output-file", "glade.go", "Output file name for the generated source.")
	generateCmd.PersistentFlags().StringVar(&pkg, "package", "ui", "The package name for the generated source.")
}
