package cmd

import (
	_ "embed"
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "bdl-datapool-searchengine",
		Short: "BDL Datapool SearchEngine ...",
		Long:  `BDL Datapool SearchEngine ...`,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Set GOMAXPROCS to the number of available CPUs
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Initialize configuration when the application starts
	cobra.OnInitialize()
}
