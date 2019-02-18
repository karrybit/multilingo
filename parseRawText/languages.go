package parserawtext

import "fmt"

func lookUp(id string) (string, error) {
	switch id {
	case "UG6LTEJBV":
		return "swift", nil
	default:
		return "", fmt.Errorf("No language corresponding to %s was found", id)
	}
}
