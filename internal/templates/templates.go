package templates

import "embed"

//go:embed *.tmpl app/*.tmpl app/src/routes/v1/*.tmpl app/src/*.tmpl
var Templates embed.FS
