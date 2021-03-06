package model

type Mapping struct {
	Name        string
	Aliases     []string
	FilePath    string
	YamlPath    string
	RegistryUrl string
	Dir			string
}

type Mappings []Mapping

func (m *Mappings) ResolveName(image string) *string {
	for _, mapping := range *m {
		for _, alias := range mapping.Aliases {
			if alias == image {
				return &mapping.Name
			}
		}
	}
	return nil
}

func (m *Mappings) FindByName(name string) *Mapping {
	for _, mapping := range *m {
		if name == mapping.Name {
			return &mapping
		}
	}
	return nil
}
