package config

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
)

// holds application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InProduction  bool
	Session       *scs.SessionManager
}
