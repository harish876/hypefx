package main

import (
    "context"
    "log"
    "os"
    "github.com/harish876/hypefx/docs/views/layout"
    "github.com/harish876/hypefx/docs/views/installation"
)

//RUN THIS FILE TO REGENERATE THE STATIC PAGES

func main() {
    f_0, err := os.Create("public/home.html")
    if err != nil {
        log.Fatalf("failed to create output file 'public/home.html': %v", err)
    }
    defer f_0.Close()

    err = layout.Base().Render(context.Background(), f_0)
    if err != nil {
        log.Fatalf("failed to write output file 'public/home.html': %v", err)
    }
    log.Printf("public/home.html created successfully.")
    f_1, err := os.Create("public/installation.html")
    if err != nil {
        log.Fatalf("failed to create output file 'public/installation.html': %v", err)
    }
    defer f_1.Close()

    err = installation.Installation().Render(context.Background(), f_1)
    if err != nil {
        log.Fatalf("failed to write output file 'public/installation.html': %v", err)
    }
    log.Printf("public/installation.html created successfully.")
}
