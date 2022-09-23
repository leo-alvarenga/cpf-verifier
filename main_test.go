package main

import "testing"

func TestVerifier(t *testing.T) {
	type data struct {
		cpf      string
		expected bool
	}

	d := []data{
		{
			"77791351598",
			false,
		},
		{
			"00000",
			false,
		},
		{
			"11111111111",
			false,
		},
		{
			"01234567890",
			false,
		},
		{
			"38721606722",
			true,
		},
	}

	for _, test := range d {
		res, _ := Verify(test.cpf)

		if test.expected != res {
			t.Errorf("Expected result of '%s' to be '%t' but got '%t\n", test.cpf, test.expected, res)
		}
	}
}

func TestGenerate(t *testing.T) {
	for i := 0; i < 1000; i++ {
		cpf := GenerateCPF()
		res, err := Verify(cpf)

		if !res || err != nil {
			t.Errorf("generated an invalid cpf! '%s'", cpf)
			break
		}
	}
}
