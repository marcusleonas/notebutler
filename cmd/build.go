package cmd

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/marcusleonas/notebutler/lib"
	"github.com/spf13/cobra"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build your notes to a html site",
	Long:  "Build your notes to a static html site, with picocss.",
	Run:   build,
}

func init() {
	rootCmd.AddCommand(buildCmd)
}

//go:embed assets/*
var embedded embed.FS

func build(cmd *cobra.Command, args []string) {
	lib.Check() // Checks if notebutler is initialised

	notesDir := "notes"
	htmlDir := "html"

	if _, err := os.Stat(htmlDir); err == nil {
		err = os.RemoveAll(htmlDir)
		if err != nil {
			log.Fatal(err)
		}
	}

	err := os.MkdirAll(htmlDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	copyCSS(htmlDir)

	err = filepath.Walk(notesDir, func(p string, i os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !i.IsDir() && strings.HasSuffix(p, ".md") {
			convert(p, htmlDir)
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Build successful!")
}

func copyCSS(htmlDir string) {
	baseCSSTemplate, err := embedded.ReadFile("assets/pico.min.css")
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create(filepath.Join(htmlDir, "pico.min.css"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.Write(baseCSSTemplate)
	if err != nil {
		log.Fatal(err)
	}
}

func convert(p string, htmlDir string) {
	content, err := os.ReadFile(p)
	if err != nil {
		log.Fatal(err)
	}

	htmlContent := convertMarkdownToHTML(content)
	if err != nil {
		log.Fatal(err)
	}
	path := filepath.Join(htmlDir, strings.Replace(filepath.Base(p), ".md", ".html", 1))
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.Write(htmlContent)
	if err != nil {
		log.Fatal(err)
	}
}

func removeBetweenCharacters(input string, startChar, endChar string) (string, error) {
	// Create a regular expression pattern to match the startChar, everything between them, and the endChar
	pattern := regexp.MustCompile(fmt.Sprintf(`%s[^%s]*%s`, regexp.QuoteMeta(startChar), regexp.QuoteMeta(endChar), regexp.QuoteMeta(endChar)))

	// Replace all matches with an empty string
	result := pattern.ReplaceAllString(input, "")
	return result, nil
}

func convertMarkdownToHTML(mdContent []byte) []byte {
	var buf strings.Builder

	mdWithoutFrontmatter, err := removeBetweenCharacters(string(mdContent), "---", "---")
	if err != nil {
		log.Fatal(err)
	}

	mdContent = []byte(mdWithoutFrontmatter)

	md := goldmark.New(
		goldmark.WithRendererOptions(html.WithUnsafe()),
	)

	if err := md.Convert(mdContent, &buf); err != nil {
		log.Fatal(err)
	}

	baseHTMLTemplate, err := embedded.ReadFile("assets/template.html")
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.New("base").Parse(string(baseHTMLTemplate))
	if err != nil {
		log.Fatal(err)
	}

	data := struct {
		Title   string
		Content template.HTML
	}{
		Title:   "Notes", // Adjust title as needed
		Content: template.HTML(buf.String()),
	}

	var finalHTML bytes.Buffer
	err = tmpl.Execute(&finalHTML, data)
	if err != nil {
		log.Fatal(err)
	}

	return finalHTML.Bytes()
}
