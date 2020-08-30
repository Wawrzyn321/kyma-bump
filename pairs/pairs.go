package pairs

import (
	"bump/mappings"
	"errors"
	"fmt"
)

type PairCollection map[string]string

func (pairs PairCollection) Dealiasize(m mappings.Mappings) (PairCollection, []error) {
	var newPairs = PairCollection{}

	var errs []error
	for image, tag := range pairs {
		name := m.ResolveName(image)
		if name == nil {
			errs = append(errs, errors.New(fmt.Sprintf("Image '%s' not found.", image)))
		} else {
			newPairs[*name] = tag
		}
	}
	return newPairs, errs
}

