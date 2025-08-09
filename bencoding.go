package bencoding

import (
	"fmt"
	"strconv"
	"unicode"
)

func decodeBencodedString(bencodedString string) (interface{}, interface{}, error) {
	var firstColonIndex int

	for i := 0; i < len(bencodedString); i++ {
		if bencodedString[i] == ':' {
			firstColonIndex = i
			break
		}
	}

	lengthStr := bencodedString[:firstColonIndex]
	length, err := strconv.Atoi(lengthStr)
	if err != nil {
		return "", "", err
	}

	remainingToBeDecoded := bencodedString[firstColonIndex+1+length:]

	return bencodedString[firstColonIndex+1 : firstColonIndex+1+length], remainingToBeDecoded, nil
}

func decodeBencodedInteger(bencodedString string) (interface{}, interface{}, error) {
	var extractedNumber int

	for i := 1; i < len(bencodedString); i++ {
		if bencodedString[i] == 'e' {
			extractedNumber, _ = strconv.Atoi(bencodedString[1:i])
			break
		}
	}

	remainingToBeDecoded := bencodedString[len(strconv.Itoa(extractedNumber))+2:]
	return extractedNumber, remainingToBeDecoded, nil
}

func decodeBencodedList(bencodedString string) (interface{}, interface{}, error) {
	decodedList := make([]interface{}, 0)
	remaining := bencodedString[1:]

	for len(remaining) > 0 && remaining[0] != 'e' {
		decodedPart, rem, err := DecodeBencode(remaining)
		if err != nil {
			return "", "", err
		}
		decodedList = append(decodedList, decodedPart)
		remaining = rem.(string)
	}

	return decodedList, remaining[1:], nil
}

func decodeBencodedDictionary(bencodedString string) (interface{}, interface{}, error) {
	decodedDict := make(map[string]interface{})
	remaining := bencodedString[1:]

	for len(remaining) > 0 && remaining[0] != 'e' {
		decodedKeyPart, remKey, errKey := DecodeBencode(remaining)
		if errKey != nil {
			return "", "", errKey
		}
		remaining = remKey.(string)

		decodedValuePart, remValue, errValue := DecodeBencode(remaining)
		if errValue != nil {
			return "", "", errValue
		}
		remaining = remValue.(string)

		decodedDict[decodedKeyPart.(string)] = decodedValuePart
	}

	return decodedDict, remaining, nil
}

func DecodeBencode(bencodedString string) (interface{}, interface{}, error) {
	if unicode.IsDigit(rune(bencodedString[0])) {
		return decodeBencodedString(bencodedString)
	} else if bencodedString[0] == 'i' {
		return decodeBencodedInteger(bencodedString)
	} else if bencodedString[0] == 'l' {
		return decodeBencodedList(bencodedString)
	} else if bencodedString[0] == 'd' {
		return decodeBencodedDictionary(bencodedString)
	}
	return "", "", fmt.Errorf("Only strings are supported at the moment")
}
