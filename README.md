# CPF verifier

A small Go package to facilitate CPF* validation.

(*) _Cadastro de Pessoa Física_, a brazilian counterpart to the american Social Security Number

## Usage
To import and use the package, first retrieve it:
```bash
go get github.com/leo-alvarenga/cpf-verifier@v1.0.1
```

## Implementation


#### `Verify`
`Verify` takes a CPF as a `string`, and returns values based on whether or not the argument is a valid CPF.

```go
func Verify(cpf string) (bool, error)
```

Possible return values:

|`bool` value |`error` value |    meaning   |
|-------------|--------------|--------------|
|   `true`    |    `nil`     | The CPF passed as argument is valid |
|   `false`   |   `error`    | The CPF passed as argument is *NOT* valid; Check the return of `.Error()` to understand what is wrong |


#### `GenerateCPF`
`GenerateCPF` is responsible for generating a single, pseudo-random, valid CPF, represented as a `string`.

```go
func GenerateCPF() string
```

## Example

The example bellow implements a simple `http` server with an endpoint to verify a CPF sent in the request body. The `Verify` function is used in the endpoint handler to veridy the CPF:

```go
package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/leo-alvarenga/cpf-verifier"
)

func spinUpServer(port string) {
	http.HandleFunc("/api/verify", handleVerify)

	log.Println("Serving dashboard on port", port)
	http.ListenAndServe(":"+port, nil)
}

func handleVerify(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	defer req.Body.Close()
	// just to give an example, DO NOT push stuff like this into prd
	type icpf struct {
		Cpf string
	}

	c := new(icpf)
	err := json.NewDecoder(req.Body).Decode(c)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"status":"` + err.Error() + `"}`))
		return
	}

	_, err = cpf.Verify(c.Cpf)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(`{"reason":"` + err.Error() + `"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	go func() {
		// 1st is valid, 2nd is not
		cpfs := []string{"38721606722", "00000000"}

		log.Println("Waiting for the server to be up...")
		time.Sleep(5 * time.Second)

		for _, c := range cpfs {
			body, _ := json.Marshal(map[string]string{
				"cpf": c,
			})

			reqBody := bytes.NewBuffer(body)
			res, err := http.Post("http://127.0.0.1:8080/api/verify", "application/json", reqBody)

			if err != nil {
				log.Panicln(err.Error())
			}

			defer res.Body.Close()
			log.Println("Sent:", c, "Got:", res.Status)
		}
	}()

	spinUpServer("8080")
}

```
