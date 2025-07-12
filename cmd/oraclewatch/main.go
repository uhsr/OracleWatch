// cmd/oraclewatch/main.go
package main

import (
"flag"
"log"
"os"

"oraclewatch/internal/oraclewatch"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := oraclewatch.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
