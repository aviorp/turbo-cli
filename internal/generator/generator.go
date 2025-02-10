package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"turbo-cli/internal/templates"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

type Generator struct{}

type EntityData struct {
	Name      string
	Package   string
	Timestamp string
}

func New() *Generator {
	return &Generator{}
}

// Custom template functions
func createTemplateFuncs() template.FuncMap {
	return template.FuncMap{
		"title": strings.Title,
		"upper": strings.ToUpper,
	}
}

func getOutputFileName(entityType, name string) string {
	switch entityType {
	case "app":
		return "index.ts"
	case "package":
		return "index.ts"
	default:
		return fmt.Sprintf("%s.%s.ts", name, entityType)
	}
}
func (g *Generator) createApp(name, cwd, emoji string) error {
	success := color.New(color.FgGreen, color.Bold)
	info := color.New(color.FgCyan)
	warning := color.New(color.FgYellow)

	// Create spinner
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Prefix = "🔨 Creating app structure "
	s.Start()
	defer s.Stop()

	// Determine output directory
	var outputDir string
	if strings.Contains(cwd, "apps") {
		outputDir = filepath.Join(cwd, name)
	} else {
		outputDir = filepath.Join(cwd, "apps", name)
	}

	// Create base directory
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("❌ failed to create directory: %w", err)
	}

	// Prepare template data
	data := EntityData{
		Name:      name,
		Package:   filepath.Base(outputDir),
		Timestamp: time.Now().Format(time.RFC3339),
	}

	// Define app file structure with correct paths
	appFiles := []struct {
		templatePath string
		outputPath   string
	}{
		{"app/src/routes/v1/router.ts.tmpl", "src/routes/v1/router.ts"},
		{"app/src/index.ts.tmpl", "src/index.ts"}, // At the same level as src
		{"app/Dockerfile.tmpl", "Dockerfile"},
		{"app/eslint.config.js.tmpl", "eslint.config.js"},
		{"app/package.json.tmpl", "package.json"},
		{"app/sonar-project.properties.tmpl", "sonar-project.properties"},
		{"app/tsconfig.json.tmpl", "tsconfig.json"},
		{"app/vitest.config.ts.tmpl", "vitest.config.ts"},
		{"controller.tmpl", "src/controllers/" + name + "/" + name + ".controller.ts"},
		{"service.tmpl", "src/services/" + name + "/" + name + ".service.ts"},
	}

	// Create and process each file
	for _, file := range appFiles {
		// Create directory if needed
		dir := filepath.Join(outputDir, filepath.Dir(file.outputPath))
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("❌ failed to create directory %s: %w", dir, err)
		}

		// Read and parse template
		templateContent, err := templates.Templates.ReadFile(file.templatePath)
		if err != nil {
			return fmt.Errorf("❌ failed to read template %s: %w", file.templatePath, err)
		}

		tmpl, err := template.New(filepath.Base(file.templatePath)).
			Funcs(createTemplateFuncs()).
			Parse(string(templateContent))
		if err != nil {
			return fmt.Errorf("❌ failed to parse template %s: %w", file.templatePath, err)
		}

		// Create output file
		outputFile := filepath.Join(outputDir, file.outputPath)
		f, err := os.Create(outputFile)
		if err != nil {
			return fmt.Errorf("❌ failed to create file %s: %w", outputFile, err)
		}

		// Execute template
		if err := tmpl.Execute(f, data); err != nil {
			f.Close()
			return fmt.Errorf("❌ failed to execute template %s: %w", file.templatePath, err)
		}
		f.Close()

		info.Printf("📝 Created %s\n", file.outputPath)
	}

	s.Stop()
	success.Printf("\n✨ Successfully created app %s\n", emoji)
	warning.Printf("📁 App location: %s\n\n", outputDir)
	info.Printf("To get started:\n")
	info.Printf("  cd %s\n", name)
	info.Printf("  pnpm install\n")
	info.Printf("  pnpm dev\n\n")

	return nil
}

func (g *Generator) CreateEntity(entityType, name string) error {
	success := color.New(color.FgGreen, color.Bold)
	info := color.New(color.FgCyan)
	warning := color.New(color.FgYellow)
	emoji := getEmoji(entityType)

	info.Printf("\n🚀 Creating new %s: ", entityType)
	success.Printf("%s %s\n", name, emoji)

	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("❌ failed to get current directory: %w", err)
	}

	// Convert name to proper format
	name = strings.ToLower(name)

	if entityType == "app" {
		return g.createApp(name, cwd, emoji)
	}

	// Create spinner for file generation
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Prefix = "🔨 Preparing " + entityType + " "
	s.Start()

	// Determine the output directory based on entity type and current location
	var outputDir string
	switch entityType {
	case "app":
		if strings.Contains(cwd, "apps") {
			outputDir = filepath.Join(cwd, name)
		} else {
			outputDir = filepath.Join(cwd, "apps", name)
		}
	case "package":
		if strings.Contains(cwd, "packages") {
			outputDir = filepath.Join(cwd, name, "src")
		} else {
			outputDir = filepath.Join(cwd, "packages", name, "src")
		}
	default:
		// For controllers, services, middleware - create in current directory under src
		if strings.Contains(cwd, "src") {
			outputDir = filepath.Join(cwd, entityType+"s", name)
		} else {
			outputDir = filepath.Join(cwd, "src", entityType+"s", name)
		}
	}

	// Create output directory
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		s.Stop()
		return fmt.Errorf("❌ failed to create directory: %w", err)
	}
	info.Printf("\n📁 Created directory: %s\n", outputDir)

	// Prepare template data
	data := EntityData{
		Name:      filepath.Base(name),
		Package:   filepath.Base(outputDir),
		Timestamp: time.Now().Format(time.RFC3339),
	}

	// Read template content
	templateContent, err := templates.Templates.ReadFile(entityType + ".tmpl")
	if err != nil {
		s.Stop()
		return fmt.Errorf("❌ failed to read template: %w", err)
	}

	// Create template with custom functions
	tmpl, err := template.New(entityType + ".tmpl").
		Funcs(createTemplateFuncs()).
		Parse(string(templateContent))
	if err != nil {
		s.Stop()
		return fmt.Errorf("❌ failed to parse template: %w", err)
	}

	// Get the appropriate output filename
	outputFileName := getOutputFileName(entityType, filepath.Base(name))
	outputFile := filepath.Join(outputDir, outputFileName)

	// Create output file
	f, err := os.Create(outputFile)
	if err != nil {
		s.Stop()
		return fmt.Errorf("❌ failed to create file: %w", err)
	}
	defer f.Close()

	// Execute template
	if err := tmpl.Execute(f, data); err != nil {
		s.Stop()
		return fmt.Errorf("❌ failed to execute template: %w", err)
	}

	s.Stop()
	success.Printf("\n✨ Successfully created %s %s\n", entityType, emoji)
	info.Printf("📝 File created at: ")
	warning.Printf("%s\n\n", outputFile)

	return nil
}

func getEmoji(entityType string) string {
	switch entityType {
	case "app":
		return "📱"
	case "package":
		return "📦"
	case "controller":
		return "🎮"
	case "service":
		return "⚙️"
	case "middleware":
		return "🔗"
	default:
		return "📄"
	}
}
