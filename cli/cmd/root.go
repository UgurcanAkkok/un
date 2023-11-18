package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func rootRun(cmd *cobra.Command, args []string) {
	fmt.Println("This is un!")
}

var rootCmd = &cobra.Command{
	Use: "un",
	Run: rootRun,
}

func Execute() {
	rootCmd.AddCommand(taskCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("There was an error executing the root cmd:", err)
		os.Exit(1)
	}

}
