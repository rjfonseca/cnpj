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
BenchmarkValidate-6     22892857                50.64 ns/op            0 B/op          0 allocs/op
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

Links:

- https://www.gov.br/receitafederal/pt-br/acesso-a-informacao/acoes-e-programas/programas-e-atividades/cnpj-alfanumerico
- https://hom.nfe.fazenda.gov.br/portal/exibirArquivo.aspx?conteudo=P3TfrfqQ38U=