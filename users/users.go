package users

import (
	"sort"
	"slices"
)

func TrueUser(name string, validNames []string) bool {
	for _, validName := range validNames {
		if name == validName {
			return true
		}
	}
	return false
}

func DeleteName(validNames []string, nameToDelete string) [] string {
	for i, validName := range validNames {
		if nameToDelete == validName {
			validNames = slices.Delete(validNames, i, i+1)
			break
		}
	}
	return validNames
}

func AddName(validNames []string, newName string) [] string {
	if TrueUser(newName, validNames) {
		return validNames
	}
	
	validNames = append(validNames, newName)
	sort.Strings(validNames)
	return validNames
}

func Initialize(names []string) []string {
	sort.Strings(names)
	return names
}