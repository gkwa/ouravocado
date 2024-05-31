package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/gkwa/ouravocado/core"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	verbose           bool
	ignorePaths       []string
	includeExtensions []string
)

// indexCmd represents the index command
var indexCmd = &cobra.Command{
	Use:   "index",
	Short: "Scan directories and generate file information",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide at least one directory path")
			os.Exit(1)
		}

		err := core.ProcessDirectories(args, viper.GetBool("verbose"), viper.GetStringSlice("ignore-path"), viper.GetStringSlice("ext"))
		if err != nil {
			fmt.Printf("Error processing directories: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(indexCmd)

	var err error

	indexCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose mode")
	err = viper.BindPFlag("verbose", indexCmd.PersistentFlags().Lookup("verbose"))
	if err != nil {
		slog.Error("error binding verbose flag", "error", err)
		os.Exit(1)
	}

	indexCmd.PersistentFlags().StringSliceVar(&ignorePaths, "ignore-path", []string{}, "substrings from paths to ignore")
	err = viper.BindPFlag("ignore-path", indexCmd.PersistentFlags().Lookup("ignore-path"))
	if err != nil {
		slog.Error("error binding ignore-path flag", "error", err)
		os.Exit(1)
	}

	indexCmd.PersistentFlags().StringSliceVar(&includeExtensions, "ext", core.DefaultIncludeExtensions, "file extensions to include")
	err = viper.BindPFlag("ext", indexCmd.PersistentFlags().Lookup("ext"))
	if err != nil {
		slog.Error("error binding ext flag", "error", err)
		os.Exit(1)
	}
}
