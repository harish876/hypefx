package static

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/harish876/hypefx/internal/cli/commands"
	"github.com/harish876/hypefx/tests/generator/views/home"
)

func TestGenerartor(t *testing.T) {
	var files []commands.StaticFile
	files = append(files,
		commands.StaticFile{
			Package:   "home",
			Component: "Home",
			FileName:  "home.html",
		},
		commands.StaticFile{
			Package:   "home",
			Component: "Home",
			FileName:  "about.html",
		},
	)
	Generator(files, "github.com/harish876/hypefx/tests/generator/views", "static")
}

func TestRunner(t *testing.T) {
	os.MkdirAll("public", os.ModePerm)
	f0, err := os.Create("public/home.html")
	if err != nil {
		log.Fatalf("failed to create output file 'home.html': %v", err)
	}
	defer f0.Close()

	err = home.Home().Render(context.Background(), f0)
	if err != nil {
		log.Fatalf("failed to write output file 'home.html': %v", err)
	}
	log.Printf("home.html created successfully.")
	f1, err := os.Create("public/about.html")
	if err != nil {
		log.Fatalf("failed to create output file 'about.html': %v", err)
	}
	defer f1.Close()

	err = home.Home().Render(context.Background(), f1)
	if err != nil {
		log.Fatalf("failed to write output file 'about.html': %v", err)
	}
	log.Printf("about.html created successfully.")
}
