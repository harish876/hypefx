package static

import (
    "context"
    "log"
    "os"
    "github.com/harish876/hypefx/tests/generator/views/home"
)

func main() {
    f_0, err := os.Create("home.html")
    if err != nil {
        log.Fatalf("failed to create output file 'home.html': %v", err)
    }
    defer f_0.Close()

    err = home.Home().Render(context.Background(), f_0)
    if err != nil {
        log.Fatalf("failed to write output file 'home.html': %v", err)
    }
    log.Printf("home.html created successfully.")
    f_1, err := os.Create("about.html")
    if err != nil {
        log.Fatalf("failed to create output file 'about.html': %v", err)
    }
    defer f_1.Close()

    err = home.Home().Render(context.Background(), f_1)
    if err != nil {
        log.Fatalf("failed to write output file 'about.html': %v", err)
    }
    log.Printf("about.html created successfully.")
}
