package users

import (
	"errors"
	"sort"
	"slices"
)

func TrueUser(name string, validNames []string) bool {
	return slices.Contains(validNames, name)
}

func DeleteName(validNames []string, nameToDelete string) ([]string, error) {

	ErrNotFound := errors.New("ErrNotFound")

	for i, validName := range validNames {
		if nameToDelete == validName {
			validNames = slices.Delete(validNames, i, i+1)
			return validNames, nil
		}
	}
	return validNames, ErrNotFound
}

func AddName(validNames []string, newName string) ([]string, error) {

	ErrAlreadyExist := errors.New("ErrAlreadyExist")

	if TrueUser(newName, validNames) {
		return validNames, ErrAlreadyExist
	}
	
	validNames = append(validNames, newName)
	sort.Strings(validNames)
	return validNames, nil
}

func Initialize(names []string) []string {
	sort.Strings(names)
	return names
}