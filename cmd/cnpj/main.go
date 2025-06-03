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

var (
	quiet  bool
	pretty bool
)

func main() {
	log.SetFlags(0)

	flag.BoolVar(&quiet, "quiet", false, "Quiet mode, don't print output when validating CNPJ")
	flag.BoolVar(&pretty, "pretty", false, "Pretty print CNPJs (format with dots and dashes)")
	generate := flag.Int("generate", 0, "Generate N random CNPJs (0 to disable)")
	flag.Parse()

	if *generate > 0 {
		generateCNPJs(*generate)
		return
	}

	cnpjs := getCNPJsFromArgsOrStdin()

	if len(cnpjs) == 0 {
		log.Fatal("No CNPJ provided. Please provide a CNPJ as an argument or via stdin.")
	}

	if validateAllCNPJs(cnpjs) {
		os.Exit(0)
	}
	os.Exit(1)

}

func getCNPJsFromArgsOrStdin() []string {
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

func generateCNPJs(count int) {
	for range count {
		printCNPJ(cnpj.Generate())
	}
}

func validateAllCNPJs(cnpjs []string) bool {
	allCNPJsAreValid := true
	for _, c := range cnpjs {
		err := cnpj.Validate(cnpj.Clean(c))
		if pretty {
			c = cnpj.Format(c)
		}
		if err != nil {
			allCNPJsAreValid = false
			printInvalidCNPJ(c, err)
			continue
		}
		printValidCNPJ(c)
	}
	return allCNPJsAreValid
}

func printCNPJ(c string) {
	if pretty {
		c = cnpj.Format(c)
	}
	fmt.Println(c)
}

func printInvalidCNPJ(c string, err error) {
	if quiet {
		return
	}
	if pretty {
		c = cnpj.Format(c)
	}
	fmt.Println(c, "is invalid:", err)
}

func printValidCNPJ(c string) {
	if quiet {
		return
	}
	if pretty {
		c = cnpj.Format(c)
	}
	fmt.Println(c, "is valid")
}
