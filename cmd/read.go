package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"
)

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read a note",
	Long:  "Read a note in the current notebook.",
	Run:   read,
}

func init() {
	rootCmd.AddCommand(readCmd)
}

func read(cmd *cobra.Command, args []string) {
	name := args[0]
	content, err := os.ReadFile(filepath.Join("notes", strings.ToLower(name+".md")))
	if err != nil {
		log.Fatal(err)
	}

	out, err := glamour.Render(string(content), "dark")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(out)
}
