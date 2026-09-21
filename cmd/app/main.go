// A Cobra CLI. NOT A SERVICE: nothing listens on $PORT, so START_CMD is empty
// and bin/run stops at the start step by design.
package main

import (
	"fmt"
	"os"

	"github.com/Qode-Platform/qode-cobra-template-v1/cmd"
)

func main() {
	if err := cmd.NewRootCmd(os.Stdout).Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
