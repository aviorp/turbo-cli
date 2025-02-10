package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"turbo-cli/internal/generator"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var (
	success = color.New(color.FgGreen, color.Bold)
	info    = color.New(color.FgCyan)
	warning = color.New(color.FgYellow)
)

var rootCmd = &cobra.Command{
	Use:   "turbo-cli",
	Short: "🚀 A CLI tool for managing turborepo monorepo",
	Long: `✨ A CLI tool for creating and managing apps, packages, and other entities in your turborepo monorepo setup.
	
Available Commands:
📱 app         - Create a new application
📦 package     - Create a new package
🎮 controller  - Create a new controller
⚙️  service     - Create a new service
🔗 middleware  - Create a new middleware`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(newCreateCommand())
	rootCmd.AddCommand(newSelectCommand())
}

type starter struct {
	Name        string
	Description string
	Type        string
	Emoji       string
}

func newSelectCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "new",
		Short: "🎯 Interactive creation of new entities",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runInteractiveCreation()
		},
	}
}

func runInteractiveCreation() error {
	entityTypes := []starter{
		{Name: "Application", Description: "Create a new application", Type: "app", Emoji: "📱"},
		{Name: "Package", Description: "Create a new package", Type: "package", Emoji: "📦"},
		{Name: "Controller", Description: "Create a new controller", Type: "controller", Emoji: "🎮"},
		{Name: "Service", Description: "Create a new service", Type: "service", Emoji: "⚙️"},
		{Name: "Middleware", Description: "Create a new middleware", Type: "middleware", Emoji: "🔗"},
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "➡️  {{ .Name | cyan }} {{ .Emoji }} ({{ .Description | yellow }})",
		Inactive: "   {{ .Name | white }} {{ .Emoji }} ({{ .Description | faint }})",
		Selected: "✨ Selected: {{ .Name | green }} {{ .Emoji }}",
		Details: `
--------- Entity Details ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Description:" | faint }}	{{ .Description }}
{{ "Type:" | faint }}	{{ .Type }}`,
	}

	entityPrompt := promptui.Select{
		Label:     "What would you like to create?",
		Items:     entityTypes,
		Templates: templates,
		Size:      5,
		HideHelp:  false,
	}

	i, _, err := entityPrompt.Run()
	if err != nil {
		return fmt.Errorf("❌ Prompt failed: %w", err)
	}

	selectedType := entityTypes[i]

	validate := func(input string) error {
		if len(input) < 1 {
			return fmt.Errorf("❌ name must not be empty")
		}
		if len(input) < 3 {
			return fmt.Errorf("❌ name must be at least 3 characters")
		}
		return nil
	}

	namePrompt := promptui.Prompt{
		Label:    fmt.Sprintf("Enter name for the new %s", selectedType.Name),
		Validate: validate,
	}

	name, err := namePrompt.Run()
	if err != nil {
		return fmt.Errorf("❌ Prompt failed: %w", err)
	}

	info.Printf("\n🚀 Creating %s: ", selectedType.Type)
	success.Printf("%s %s\n", name, selectedType.Emoji)

	return createEntity(selectedType.Type, name)
}

func validateContext(entityType string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	switch entityType {
	case "app", "package":
		if !isMonorepoRoot(cwd) && !strings.Contains(cwd, entityType+"s") {
			return fmt.Errorf(`
This command should be run from:
- The monorepo root directory (containing turbo.json)
- The %ss directory

Current directory: %s
`, entityType, cwd)
		}
	default:
		if !strings.Contains(cwd, "apps") {
			return fmt.Errorf(`
This command should be run from within an app directory.
Example: cd apps/my-app && turbo-cli create %s %s

Current directory: %s
`, entityType, "name", cwd)
		}
	}
	return nil
}

func isMonorepoRoot(dir string) bool {
	indicators := []string{
		"turbo.json",
		"package.json",
		"apps",
		"packages",
	}

	for _, indicator := range indicators {
		if _, err := os.Stat(filepath.Join(dir, indicator)); err == nil {
			return true
		}
	}
	return false
}

func createEntity(entityType, name string) error {
	if err := validateContext(entityType); err != nil {
		return err
	}

	generator := generator.New()
	if err := generator.CreateEntity(entityType, name); err != nil {
		return fmt.Errorf("failed to create %s '%s': %w", entityType, name, err)
	}

	return nil
}

func newCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [entity]",
		Short: "Create a new entity",
		Long:  `Create a new entity (app, package, controller, service, middleware)`,
	}

	cmd.AddCommand(newCreateAppCommand())
	cmd.AddCommand(newCreatePackageCommand())
	cmd.AddCommand(newCreateControllerCommand())
	cmd.AddCommand(newCreateServiceCommand())
	cmd.AddCommand(newCreateMiddlewareCommand())

	return cmd
}

func newCreateAppCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "app [name]",
		Short: "Create a new app",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return createEntity("app", args[0])
		},
	}
}

func newCreatePackageCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "package [name]",
		Short: "Create a new package",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return createEntity("package", args[0])
		},
	}
}

func newCreateControllerCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "controller [name]",
		Short: "Create a new controller",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return createEntity("controller", args[0])
		},
	}
}

func newCreateServiceCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "service [name]",
		Short: "Create a new service",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return createEntity("service", args[0])
		},
	}
}

func newCreateMiddlewareCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "middleware [name]",
		Short: "Create a new middleware",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return createEntity("middleware", args[0])
		},
	}
}
