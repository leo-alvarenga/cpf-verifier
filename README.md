# CPF verifier

A small Go package to facilitate CPF* validation.

(*) _Cadastro de Pessoa Física_, a brazilian counterpart to the american Social Security Number

## Usage
To import and use the package, first download it:
```bash
go get github.com/leo-alvarenga/cpf-verifier@v1.0.0
```

### Implementation

CPF Verifier exposes two functions `GenerateCPF`, responsible for generating a single, pseudo-random, valid CPF, represented as a `string`. The second function, `Verify`, takes a `string` containing CPF, and returns values based on whether or not the argument is a valid CPF.

```go
func GenerateCPF() string
```

```go
func Verify(cpf string) (bool, error)
```
