package parserawtext

import "testing"

func TestParseLanguageNormal(t *testing.T) {
	// map[string]string{argument: expected}
	langMap := map[string]string{"UG6LTEJBV": "swift"}
	for k, v := range langMap {
		lang, err := lookUp(k)
		if err != err {
			t.Error(err)
		}
		if lang != v {
			t.Errorf("lookUp not %s, got= %s", v, lang)
		}
	}
}

func TestParseProgramNormal(t *testing.T) {
	// map[string]string{argument: expected}
	programMap := map[string]string{"```print(\"hello world\")```": "print(\"hello world\")",
		"```func main() {\nprint(114514)\n}\n```": "func main() {\nprint(114514)\n}"}
	for k, v := range programMap {
		program := parseProgram(k)
		if program != v {
			t.Errorf("parseProgram not %s, got= %s", v, program)
		}
	}
}

func TestParseNormal(t *testing.T) {
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