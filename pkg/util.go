package util

import (
	"fmt"
	"os"
	"path/filepath"
  "text/template"
)

type TemplateVars struct {}

func fileExists(name string) bool {
  _, err := os.Stat(name)
  return err == nil
}

func CreateDayPackage(day int) {
  packageName := fmt.Sprintf("day%02d", day)
  srcName := filepath.Join(packageName, "day.go")

  if !fileExists(packageName) {
    mkdirErr := os.Mkdir(packageName, 0755)
    if mkdirErr != nil {
      panic(mkdirErr)
    }
  }

  if !fileExists(srcName) {
    t, templateErr := template.ParseFiles("template.go.tmpl")
    if templateErr != nil {
      panic(templateErr)
    }

    f, fileErr := os.Create(srcName)
    defer f.Close()
    if fileErr != nil {
      panic(fileErr)
    }

    vars := TemplateVars{}
    t.Execute(f, vars)
  }
}
