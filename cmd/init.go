package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new notebook",
	Long:  "Initialize a new notebook in the current directory.",
	Run:   initFunc,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

type Config struct {
	NotebookName string `json:"name"`
}

func initFunc(cmd *cobra.Command, args []string) {
	var (
		directory string
		nbName    string // Name of the notebook
	)
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("Directory").Value(&directory).Validate(func(str string) error {
				invalidChars := "/\\:*?\"<>|"
				if strings.ContainsAny(str, invalidChars) {
					return errors.New("directory name invalid. can not contain: /\\:*?\"<>|")
				}
				return nil
			}),
			huh.NewInput().Title("Notebook Name").Value(&nbName),
		),
	).WithTheme(huh.ThemeCatppuccin())

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	action := func() {
		err := os.MkdirAll(filepath.Join(directory, ".notebutler/templates"), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}

		err = os.MkdirAll(filepath.Join(directory, "notes"), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}

		filePath := filepath.Join(directory, "notebutler.json")
		file, err := os.Create(filePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// _, err = file.WriteString("{}")
		// if err != nil {
		// 	log.Fatal(err)
		// }

		config := Config{
			NotebookName: nbName,
		}

		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")
		err = encoder.Encode(&config)
		if err != nil {
			log.Fatal(err)
		}

		templatePath := filepath.Join(directory, ".notebutler/templates/default.md")
		defaultTemplate, err := os.Create(templatePath)
		if err != nil {
			log.Fatal(err)
		}
		defer defaultTemplate.Close()

		data := `---
name: {{ .Name }}
---
# {{ .Name }}
Created at: {{ .CreatedAt }}

`
		defaultTemplate.WriteString(data)

		fmt.Println(`Notebook initialized successfully!`)
	}

	_ = spinner.New().Title("Initializing Notebutler...").Action(action).Run()

	// fmt.Println("Directory:", directory)
	// fmt.Println("Notebook Name:", nbName)
}
