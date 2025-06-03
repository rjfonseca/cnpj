# Alphanumeric CNPJ

My take on the brazilian alphanumeric CNPJ validation.

Goals:
- Zero allocation
- Zero dependencies
- No regular expressions
- Fast
- 100% test coverage

Benchmark:

```shell
$ go test -benchmem -run='^$' -bench '^Benchmark*' .
goos: linux
goarch: amd64
pkg: github.com/rjfonseca/cnpj
cpu: Intel(R) Core(TM) i5-8600K CPU @ 3.60GHz
BenchmarkValidate-6     25791698                45.71 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/rjfonseca/cnpj       1.217s
```

## Usage

```go
import (
	"github.com/rjfonseca/cnpj"
)

// Not a proper working code, just some examples 
// on how to call the functions on this package
func example(c string) {
    // Cleans the input, if needed
	c = cnpj.Clean(c)

    // Validate, receiving error
	err := cnpj.Validate(c)
    if err != nil {
        //handle error
    }

    // Validate, receiving boolean
    if cnpj.IsValid(c) {
        // handle valid
    }

    // Format as XX.XXX.XXX/XXXX-XX
	c = cnpj.Format(c)
	
    }
// <more code>
```

## CLI

There is also a CLI available that validates and generates CNPJs em batches.

```shell
go install github.com/rjfonseca/cnpj/cmd/cnpj@latest
```

Examples:

Generating random CNPJs

```shell
$ cnpj -generate 5
2WR35R661WPC43
083HL2707A0837
4RKKYLDQKMS813
BW6Y2694366053
5F80W4024LC727

$ cnpj -pretty -generate 1 
FW.29Z.3ES/5W3V-57
```

Validating from `stdin`

```shell
$ cnpj -generate 5 | cnpj -pretty
2Z.509.REZ/1HQE-12 is valid
S5.D29.A73/1C4Z-61 is valid
7W.Y06.193/526Z-26 is valid
99.3K0.BO4/AU28-83 is valid
53.9K7.4MC/OM2F-75 is valid
```

Validating from command-line args:

```shell
$ cnpj 53.9K7.4MC/OM2F-75 9TRPPM4BPC5114 9TRPPM4BPC5113 banana
53.9K7.4MC/OM2F-75 is valid
9TRPPM4BPC5114 is valid
9TRPPM4BPC5113 is invalid: unexpected digit at position 13: expected 4, got 3
banana is invalid: invalid CNPJ length, must be 14 characters
```

Quiet mode for scripting:

```shell
$ cnpj -quiet banana && echo "valid" || echo "invalid"
invalid
```

## Links:

- https://www.gov.br/receitafederal/pt-br/acesso-a-informacao/acoes-e-programas/programas-e-atividades/cnpj-alfanumerico
- https://hom.nfe.fazenda.gov.br/portal/exibirArquivo.aspx?conteudo=P3TfrfqQ38U=