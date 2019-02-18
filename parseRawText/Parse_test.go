package parserawtext

import "testing"

func TestNormalLanguage(t *testing.T) {
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
