package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/rjfonseca/cnpj"
)

func main() {
	log.SetFlags(0)
	quiet := flag.Bool("quiet", false, "Quiet mode, no output")
	pretty := flag.Bool("pretty", false, "Pretty print CNPJ")

	flag.Parse()

	cnpjs := getCNPJs()

	if len(cnpjs) == 0 {
		log.Fatal("No CNPJ provided. Please provide a CNPJ as an argument or via stdin.")
	}

	atLeastOneInvalid := false
	for _, c := range cnpjs {
		c = cnpj.Clean(c)
		err := cnpj.Validate(c)
		if *pretty {
			c = cnpj.Format(c)
		}
		if err != nil {
			atLeastOneInvalid = true
			if !*quiet {
				fmt.Println(c, "is invalid:", err)
			}
			continue
		}
		if !*quiet {
			fmt.Println(c, "is valid")
		}
	}

	if atLeastOneInvalid {
		os.Exit(1)
	}

}

func getCNPJs() []string {
	if len(flag.Args()) > 0 {
		return flag.Args()
	}
	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Error reading CNPJ from stdin: %v", err)
	}

	cnpjs := make([]string, 0)
	scanner := bufio.NewScanner(strings.NewReader(string(stdin)))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			cnpjs = append(cnpjs, line)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading CNPJ from stdin: %v", err)
	}

	return cnpjs
}
