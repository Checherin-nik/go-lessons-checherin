package users

import (
	"errors"
	"sort"
	"slices"
)

var (
	ErrNotFound = errors.New("ErrNotFound")
	ErrAlreadyExist = errors.New("ErrAlreadyExist")
)

type Names struct {
	validNames []string
}

func (n *Names) TrueUser(name string) bool {
	return slices.Contains(n.validNames, name)
}

func (n *Names) Delete(name string) error {
	for i, validName := range n.validNames {
		if name == validName {
			n.validNames = slices.Delete(n.validNames, i, i+1)
			return nil
		}
	}
	return ErrNotFound
}

func (n *Names) Add(name string) error {
	if n.TrueUser(name) {
		return ErrAlreadyExist
	}
	
	n.validNames = append(n.validNames, name)
	sort.Strings(n.validNames)
	return nil
}

func (n *Names) Initialize(names []string) {
	n.validNames = names
	sort.Strings(n.validNames)
}

func (n *Names) GetValidNames() []string {
	return n.validNames
}