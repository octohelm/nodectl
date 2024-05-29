/*
Package wd GENERATED BY gengo:runtimedoc 
DON'T EDIT THIS FILE
*/
package wd

// nolint:deadcode,unused
func runtimeDoc(v any, names ...string) ([]string, bool) {
	if c, ok := v.(interface {
		RuntimeDoc(names ...string) ([]string, bool)
	}); ok {
		return c.RuntimeDoc(names...)
	}
	return nil, false
}

func (Dir) RuntimeDoc(names ...string) ([]string, bool) {
	return []string{}, true
}
func (v OSInfo) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "OSVersion":
			return []string{}, true
		case "Platform":
			return []string{}, true
		case "Home":
			return []string{}, true

		}
		if doc, ok := runtimeDoc(v.OSVersion, names...); ok {
			return doc, ok
		}

		return nil, false
	}
	return []string{}, true
}

func (OptionFunc) RuntimeDoc(names ...string) ([]string, bool) {
	return []string{}, true
}
func (v Options) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "BasePath":
			return []string{}, true
		case "Env":
			return []string{}, true
		case "User":
			return []string{}, true
		case "Stdout":
			return []string{}, true
		case "Stderr":
			return []string{}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v Platform) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Architecture":
			return []string{
				"Architecture field specifies the CPU architecture, for example",
				"`amd64` or `ppc64le`.",
			}, true
		case "OS":
			return []string{
				"OS specifies the operating system, for example `linux` or `windows`.",
			}, true
		case "OSVersion":
			return []string{
				"OSVersion is an optional field specifying the operating system",
				"version, for example on Windows `10.0.14393.1066`.",
			}, true
		case "OSFeatures":
			return []string{
				"OSFeatures is an optional field specifying an array of strings,",
				"each listing a required OS feature (for example on Windows `win32k`).",
			}, true
		case "Variant":
			return []string{
				"Variant is an optional field specifying a variant of the CPU, for",
				"example `v7` to specify ARMv7 when architecture is `arm`.",
			}, true

		}

		return nil, false
	}
	return []string{}, true
}
