package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*html"))
}

const (
	port = ":8080"
	host = "http://localhost"
)

type TemplateData struct {
	Input string
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Staring application on port%s\n", port)
	log.Println("Open:", host+port)

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("%v - internal server error", http.StatusInternalServerError)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method == "GET" {
		tpl.ExecuteTemplate(w, "index.html", nil)
	} else {
		r.ParseForm()
		input := r.FormValue("textarea_input")
		if !inputValidation(input) {
			http.Error(w, "400 - bad request", http.StatusBadRequest)
			return
		}

		fonts := r.FormValue("font")

		banner := ""

		if string(fonts) == "standard" {
			banner = "banner/standard.txt"
		} else if string(fonts) == "shadow" {
			banner = "banner/shadow.txt"
		} else if string(fonts) == "tinkertoy" {
			banner = "banner/tinkertoy.txt"
		}

		font, err := mapFonts(banner)
		if err != nil {
			http.Error(w, "400 - bad request", http.StatusBadRequest)
			return
		}

		uI, err := mapUserInput(input, font)
		if err != nil {
			http.Error(w, "400 - bad request", http.StatusBadRequest)
			return
		}

		data := TemplateData{
			Input: uI,
		}

		tpl.ExecuteTemplate(w, "index.html", data)

	}
}

func mapFonts(font string) (map[int][]string, error) {
	readFile, err := os.ReadFile(font)
	if err != nil {
		return nil, fmt.Errorf("could not read the content in the file: %v", err)
	}
	slice := strings.Split(string(readFile), "\n")

	ascii := make(map[int][]string)
	i := 31

	for _, ch := range slice {
		if ch == "" {
			i++
		} else {
			ascii[i] = append(ascii[i], ch)
		}
	}
	return ascii, nil
}

// Handle user input and map it to ascii
func mapUserInput(input string, ascii map[int][]string) (string, error) {
	lines := strings.Split(input, "\r\n")
	output := ""

	for _, line := range lines {
		characters := []rune(line)

		if line != "" {
			for i := 0; i < 8; i++ {
				for _, ch := range characters {
					if ch >= 32 && ch <= 126 {
						output = output + ascii[int(ch)][i]
					} else {
						return "", fmt.Errorf("input includes non ascii character(s), please use ascii character(s)")
					}
				}
				output = output + "\n"
			}
		} else {
			output = output + "\n"
		}
	}
	return output, nil
}

// User input validation
func inputValidation(input string) bool {
	for _, ch := range input {
		if ch < 32 || ch > 126 {
			if ch == 10 || ch == 13 {
				continue
			}
			return false
		}
	}
	return true
}
