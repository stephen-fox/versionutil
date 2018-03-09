package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/stephen-fox/versionutil"
)

const (
	strArg = "s"
)

var (
	str = flag.String(strArg, "", "The string to get a version from")
)

func main() {
	flag.Parse()

	if len(os.Args) == 1 {
		flag.PrintDefaults()
	}

	if len(strings.TrimSpace(*str)) > 0 {
		version, err := versionutil.StringToVersion(*str)
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println("Long:", version.Long())
		fmt.Println("Short:", version.Short())
		fmt.Println("Has build:", version.HasBuild)
		fmt.Println("Is set:", version.IsSet())
	}
}
