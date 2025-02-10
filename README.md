<div align="center">
  
  
  <img src="https://media1.tenor.com/m/hfI7Dbn7OqQAAAAd/goku-flying-nimbus.gif" alt="Turbo CLI" width="600px" />
  <h1>🚀 Turbo CLI</h1>

  <p>A powerful CLI tool for managing Turborepo monorepo projects. Quickly scaffold applications, packages, controllers, services, and middleware with consistent structure and best practices.</p>

</div>

## ✨ Features

- 📱 Create new applications with full setup
- 📦 Generate packages with proper structure
- 🎮 Scaffold controllers with Fastify setup
- ⚙️ Generate services with proper patterns
- 🔗 Create middleware with error handling
- 🎯 Interactive mode for easier creation
- 🎨 Beautiful CLI interface with colors and emojis

## 🛠 Installation

1. Clone the repository:

```bash
git clone https://github.com/aviorp/turbo-cli.git
```

2. Run the installation script:

```bash
chmod +x install-cli.sh
./install-cli.sh
```

3. Source your shell configuration:

```bash
source ~/.zshrc
```

## 📚 Usage

### Command Line Interface

### Create new entities using the following commands:

```bash
# Create a new application
tc create app my-app

# Create a new package
tc create package shared-ui

# Create a new controller
tc create controller users

# Create a new service
tc create service auth

# Create a new middleware
tc create middleware logging

```

## 📚 Quick Aliases

The CLI provides convenient aliases for faster development:

```bash
tca my-app          # Create app
tcs my-service      # Create service
tcc my-controller   # Create controller
tcm my-middleware   # Create middleware
```

## 🎯 Interactive Mode

```bash
tc new # Interactive mode
```

Navigate using:
• ↑/↓ arrows to move
• Enter to select
• ? for help
• Ctrl+C to cancel

🤝 Contributing
Contributions, issues and feature requests are welcome! Feel free to check issues page.
🌟 Show your support
Give a ⭐️ if this project helped you!
📫 Contact
• GitHub: @yourusername
• Twitter: @yourhandle
🙏 Acknowledgments
• Built with Cobra
• Inspired by Turborepo
• Color support by fatih/colorx
• Interactive prompts by manifoldco/promptui
