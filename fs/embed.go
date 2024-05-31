package fs

import "embed"

//go:embed templates/*
var TemplatesFS embed.FS

//go:embed locales/*
var LocalesFS embed.FS

//go:embed static/*
var StaticFS embed.FS
