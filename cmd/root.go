// Package cmd contains all of the commands that may be executed in the cli
package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/devops-kung-fu/common/github"
	"github.com/gookit/color"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var (
	version = "0.1.0"
	//Afs stores a global OS Filesystem that is used throughout meadow
	Afs = &afero.Afero{Fs: afero.NewOsFs()}
	//Verbose determines if the execution of hing should output verbose information
	debug   bool
	rootCmd = &cobra.Command{
		Use:     "meadow generate [flags] [glade file]",
		Example: "  meadow generate --output-path ui main.glade",
		Short:   "Generates go code from Glade files.",
		Version: version,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if !debug {
				log.SetOutput(io.Discard)
			}

			log.Println("Start")
			color.Style{color.FgGreen, color.OpBold}.Println("                         _                ")
			color.Style{color.FgGreen, color.OpBold}.Println(" _ __    ___   __ _   __| |  ___  __ __ __")
			color.Style{color.FgGreen, color.OpBold}.Println("| '  \\  / -_) / _` | / _` | / _ \\ \\ V  V /")
			color.Style{color.FgGreen, color.OpBold}.Println("|_|_|_| \\___| \\__,_| \\__,_| \\___/  \\_/\\_/ ")
			fmt.Println()
			fmt.Println("DKFM - DevOps Kung Fu Mafia")
			fmt.Println("https://github.com/devops-kung-fu/meadow")
			fmt.Printf("Version: %s\n", version)
			fmt.Println()

			latestVersion, _ := github.LatestReleaseTag("devops-kung-fu", "meadow")
			if !strings.Contains(latestVersion, version) {
				color.Yellow.Printf("A newer version of meadow is available (%s)\n\n", latestVersion)
			}

		},
	}
)

// Execute creates the command tree and handles any error condition returned
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Displays debug level log messages.")
}
