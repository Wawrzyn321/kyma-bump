package requirements

type RequirementFunc func () error

func Check(list ...RequirementFunc) error {
	for _, f := range list {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}
