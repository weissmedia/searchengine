package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/weissmedia/searchengine/version"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of the application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %s\n", version.Version)
		fmt.Printf("Git Commit: %s\n", version.GitCommit)
		fmt.Printf("Build Date: %s\n", version.BuildDate)
		fmt.Printf("Go Version: %s\n", version.GoVersion)
		fmt.Printf("OS/Arch: %s\n", version.OsArch)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
