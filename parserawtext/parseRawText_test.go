package parserawtext

import "testing"

func TestParse(t *testing.T) {
	// map[string]string{argument: expected}
	textMap := map[string]Pair{"<@UG6LTEJBV>\n```print(114514)```\n": Pair{"swift", "print(114514)"}}
	for k, v := range textMap {
		lang, program, err := Parse(k)
		if err != err {
			t.Error(err)
		}

		if lang != v.lang {
			t.Errorf("Parse language %s, got= %s", v.lang, lang)
		}

		if program != v.program {
			t.Errorf("Parse program %s, got= %s", v.program, program)
		}
	}
}

type Pair struct {
	lang    string
	program string
}
