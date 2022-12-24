package app

import (
	"fmt"
	"os"

	"github.com/photowey/popctl/internal/version"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:     "popctl",
	Short:   "Auxiliary development tools for Popup Spring Cloud microservice development scaffold",
	Version: version.Now(),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Welcome to popctl cmder %s~", version.Now())
	},
}

func init() {
	cobra.OnInitialize(onInit)
	root.AddCommand(javaCmd)
}

func Run() {
	if err := root.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
