/*
Package dagger GENERATED BY gengo:runtimedoc 
DON'T EDIT THIS FILE
*/
package dagger

// nolint:deadcode,unused
func runtimeDoc(v any, names ...string) ([]string, bool) {
	if c, ok := v.(interface {
		RuntimeDoc(names ...string) ([]string, bool)
	}); ok {
		return c.RuntimeDoc(names...)
	}
	return nil, false
}

func (Conn) RuntimeDoc(names ...string) ([]string, bool) {
	return []string{}, true
}
func (v EnvVariable) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Name":
			return []string{}, true
		case "Value":
			return []string{}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v Hosts) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Default":
			return []string{}, true
		case "Platformed":
			return []string{}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v ImageConfig) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "WorkingDir":
			return []string{}, true
		case "Env":
			return []string{}, true
		case "Labels":
			return []string{}, true
		case "Entrypoint":
			return []string{}, true
		case "Cmd":
			return []string{}, true
		case "User":
			return []string{}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v Label) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Name":
			return []string{}, true
		case "Value":
			return []string{}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v PiperRunnerHost) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Name":
			return []string{}, true
		case "RunnerHost":
			return []string{}, true
		case "Platforms":
			return []string{}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v Scope) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Platform":
			return []string{}, true
		case "ID":
			return []string{}, true

		}

		return nil, false
	}
	return []string{}, true
}
