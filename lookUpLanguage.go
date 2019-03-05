package main

import "fmt"

func lookUpLanguage(request *APIGateWayRequest) (string, error) {
	switch request.APIAppID {
	case "AG6LQER0B":
		return "swift", nil
	default:
		return "", fmt.Errorf("No language corresponding to %s was found", request.APIAppID)
	}
}

func debugLookUpLanguage(id string) (string, error) {
	switch id {
	case "AG6LQER0B":
		return "swift", nil
	default:
		return "", fmt.Errorf("No language corresponding to %s was found", id)
	}
}
