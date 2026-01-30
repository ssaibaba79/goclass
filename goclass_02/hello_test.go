package hello

import "testing"

func TestSayHello(t *testing.T) {

	scenarios := []struct {
		input  []string
		result string
	}{
		{
			result: "Hello World!",
		},
		{
			input:  []string{"Saravanan"},
			result: "Hello Saravanan!",
		},
		{
			input:  []string{"Saravanan", "Poornima"},
			result: "Hello Saravanan, Poornima!",
		},
	}

	for _, st := range scenarios {
		if s := Say(st.input); s != st.result {
			t.Errorf("input:%v wanted:%s, got :%s", st.input, st.result, s)
		}
	}
}
