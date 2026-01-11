package generator

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Data struct {
	Name    string
	Pascal  string
	Modules []string
}

//go:embed templates/**
var templateFS embed.FS

// ---------- helpers ----------

func run(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func fail(err error) {
	fmt.Println("[optics] Error:", err)
	os.Exit(1)
}

func singular(name string) string {
	if strings.HasSuffix(name, "s") {
		return strings.TrimSuffix(name, "s")
	}
	return name
}

// ---------- generators ----------

func GenerateApp(name string) {
	if _, err := os.Stat("go.mod"); errors.Is(err, os.ErrNotExist) {
		if err := run("go", "mod", "init", name); err != nil {
			fail(err)
		}
	}

	if err := run("go", "get", "github.com/gofiber/fiber/v2"); err != nil {
		fail(err)
	}

	dirs := []string{
		"cmd/server",
		"internal/config",
		"internal/database",
		"internal/middleware",
		"pkg/response",
	}

	for _, d := range dirs {
		if err := os.MkdirAll(d, 0755); err != nil {
			fail(err)
		}
	}

	render(
		"templates/app/main.go.tpl",
		"cmd/server/main.go",
		Data{
			Name:    name,
			Modules: []string{},
		},
	)
	render("templates/app/env.tpl", ".env", Data{Name: name})

	if err := run("go", "mod", "tidy"); err != nil {
		fail(err)
	}

	fmt.Println("[optics] " + name + " created successfully")
}

func registerModule(module string) {
	path := "internal/config/modules.go"

	content, _ := os.ReadFile(path)
	if strings.Contains(string(content), module) {
		return
	}

	f, _ := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()

	f.WriteString("\n\t\"" + module + "\",")
}

func GenerateModule(name string) {
	name = singular(name)

	pascal := cases.Title(language.Und).String(name)
	base := filepath.Join("internal", name)

	if err := os.MkdirAll(base, 0755); err != nil {
		fail(err)
	}

	data := Data{Name: name, Pascal: pascal}

	files := map[string]string{
		"controller.go.tpl": "controller.go",
		"service.go.tpl":    "service.go",
		"repository.go.tpl": "repository.go",
		"model.go.tpl":      "model.go",
		"routes.go.tpl":     "routes.go",
	}

	for tpl, out := range files {
		render(
			path.Join("templates/module", tpl),
			filepath.Join(base, name+"."+out),
			data,
		)
	}
	registerModule("internal/" + name)
	fmt.Printf("[optics] Resource '%s' generated\n", name)
}

func GenerateResource(name string) {
	GenerateModule(name)
}

// ---------- template renderer ----------

func render(tplPath, outPath string, data Data) {
	tplBytes, err := fs.ReadFile(templateFS, tplPath)
	if err != nil {
		fail(fmt.Errorf("[optics] template not found: %s", tplPath))
	}

	tpl, err := template.New(tplPath).Parse(string(tplBytes))
	if err != nil {
		fail(err)
	}

	if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
		fail(err)
	}

	file, err := os.Create(outPath)
	if err != nil {
		fail(err)
	}
	defer file.Close()

	if err := tpl.Execute(file, data); err != nil {
		fail(err)
	}
}
