package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	stdTemplate "text/template"

	"github.com/charmbracelet/huh"
	"github.com/marcusleonas/notebutler/lib"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new note",
	Long:  "Create a new note in the current notebook.",
	Run:   new,
}

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.Flags().StringP("name", "n", "", "name")
	newCmd.Flags().StringP("template", "t", "default", "template")
	newCmd.Flags().BoolP("open", "o", false, "open")
}

func new(cmd *cobra.Command, args []string) {
	name, _ := cmd.Flags().GetString("name")
	template, _ := cmd.Flags().GetString("template")
	open, _ := cmd.Flags().GetBool("open")

	if !strings.HasSuffix(template, ".md") {
		template += ".md"
	}

	lib.Check() // Checks if notebutler is initialised

	configBytes, err := os.ReadFile("notebutler.json")
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		log.Fatal(err)
	}

	templates, err := os.ReadDir(".notebutler/templates")
	if err != nil {
		log.Fatal(err)
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("Note Name").Value(&name),
			huh.NewSelect[string]().Title("Template").OptionsFunc(func() []huh.Option[string] {
				var options []huh.Option[string]
				for _, t := range templates {
					if !t.IsDir() {
						lower := strings.ToLower(t.Name())
						options = append(options, huh.Option[string]{Value: lower, Key: lower})
					}
				}
				return options
			}, &template).Value(&template),
			huh.NewConfirm().Title("Open in default editor?").Value(&open).Affirmative("Yes").Negative("No"),
		),
	)
	if name == "" || template == "" {
		err := form.Run()
		if err != nil {
			log.Fatal(err)
		}
	}

	nameWithSuffix := name + ".md"
	if _, err := os.Stat(filepath.Join("notes", strings.ToLower(nameWithSuffix))); err == nil {
		log.Fatal("Note already exists. Run `notebutler read` to read the note.")
	}

	fmt.Println(name, template)

	content, err := os.ReadFile(filepath.Join(".notebutler/templates", template))
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err := stdTemplate.New("markdown").Parse(string(content))
	if err != nil {
		log.Fatal(err)
	}

	d := time.Now().Format("01/02/2006")
	t := time.Now().Format("15:04:05")
	n := time.Now().Format("2006-01-02 15:04:05")
	data := struct {
		Name      string
		Notebook  string
		CreatedAt string
		Date      string
		Time      string
	}{
		Name:      name,
		Notebook:  config.NotebookName,
		CreatedAt: n,
		Date:      d,
		Time:      t,
	}

	var result bytes.Buffer
	err = tmpl.Execute(&result, data)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create(filepath.Join("notes", strings.ToLower(name+".md")))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(result.String())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Note created successfully!")

	if open {
		editor := os.Getenv("EDITOR")
		if editor == "" {
			log.Fatal("No editor found. Please set the EDITOR environment variable.")
		}
		cmd := exec.Command(editor, filepath.Join("./notes", strings.ToLower(name+".md")))
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}
