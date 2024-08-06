package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the generated HTML site locally",
	Long:  "Start a local server to preview the generated HTML site.",
	Run:   serve,
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serve(cmd *cobra.Command, args []string) {
	htmlDir := "html"

	// Ensure the directory exists
	if _, err := os.Stat(htmlDir); os.IsNotExist(err) {
		log.Fatalf("The directory %s does not exist. Please run `notebutler build` first.", htmlDir)
	}

	// Create a file server to serve the static files
	fs := http.FileServer(http.Dir(htmlDir))

	http.Handle("/", fs)

	port := 8080
	fmt.Printf("Serving files from %s on http://localhost:%d\n", htmlDir, port)

	// Start the server
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal(err)
	}
}
