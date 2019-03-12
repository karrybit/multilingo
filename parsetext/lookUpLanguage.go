package parsetext

import "fmt"

// LookUpLanguage -
func LookUpLanguage(id string) (string, error) {
	switch id {
	case "AG6LQER0B":
		return "swift", nil
	default:
		return "", fmt.Errorf("No language corresponding to %s was found", id)
	}
}
