// Package cmd holds the CLI's command tree.
package cmd

import (
	"fmt"
	"io"
	"strings"

	"github.com/spf13/cobra"
)

// NewRootCmd builds the command tree. Taking the output writer as an argument
// is what makes the whole CLI testable without spawning a process.
func NewRootCmd(out io.Writer) *cobra.Command {
	root := &cobra.Command{
		Use:           "app",
		Short:         "A Cobra CLI template",
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	root.SetOut(out)

	var upper bool
	greet := &cobra.Command{
		Use:   "greet [name]",
		Short: "Print a greeting",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			if upper {
				name = strings.ToUpper(name)
			}
			fmt.Fprintf(cmd.OutOrStdout(), "hello, %s\n", name)
			return nil
		},
	}
	greet.Flags().BoolVar(&upper, "upper", false, "shout the name")
	root.AddCommand(greet)
	return root
}
