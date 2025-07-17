// cmd/artificialintuitionenginesolutionspro/main.go
package main

import (
"flag"
"log"
"os"

"artificialintuitionenginesolutionspro/internal/artificialintuitionenginesolutionspro"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := artificialintuitionenginesolutionspro.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
