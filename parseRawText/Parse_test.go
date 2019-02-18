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
	programMap := map[string]string{"```print(\"hello world\")```": "print(\"hello world\")"}
	for k, v := range programMap {
		program := parseProgram(k)
		if program != v {
			t.Errorf("parseProgram not %s, got= %s", v, program)
		}
	}
}
