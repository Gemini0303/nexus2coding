package utils

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"log"
	"os"
	"strings"
)

type Needle struct {
	Name   string
	Desc   string
	Type   string
	Format string
	URL    string
}

func ExitOption(items []Needle) []Needle {
	items = append(items, Needle{Name: "<Exit>", Desc: "exit the select"})
	return items
}

func SelectUI(items []Needle, label string) int {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "\U0001F63C {{ .Name| red }} ",
		Inactive: " {{ .Name | cyan }}",
		Selected: "\0001F638 Select:{{ .Name | green }}",
		Details: `
------- Info -------
{{ "Name:" | faint }} {{ .Name }}
{{ "Format:" | faint }} {{ .Format }}
{{ "Type:" | faint }} {{ .Type }}
{{ "Url:" | faint }} {{ .URL }}`,
	}

	searcher := func(input string, index int) bool {
		item := items[index]
		name := strings.Replace(strings.ToLower(item.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)
		if input == "q" && name == "<exit>" {
			return true
		}
		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     label,
		Items:     items,
		Templates: templates,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	if items[i].Name == "<Exit>" {
		fmt.Println("Exited.")
		os.Exit(1)
	}
	return i
}
