/*
Package kubepkg GENERATED BY gengo:runtimedoc 
DON'T EDIT THIS FILE
*/
package kubepkg

// nolint:deadcode,unused
func runtimeDoc(v any, names ...string) ([]string, bool) {
	if c, ok := v.(interface {
		RuntimeDoc(names ...string) ([]string, bool)
	}); ok {
		return c.RuntimeDoc(names...)
	}
	return nil, false
}

func (v Apply) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Task":
			return []string{}, true
		case "Kubeconfig":
			return []string{
				"Kubeconfig path",
			}, true
		case "Manifests":
			return []string{
				"Manifests of k8s resources",
			}, true

		}
		if doc, ok := runtimeDoc(v.Task, names...); ok {
			return doc, ok
		}

		return nil, false
	}
	return []string{
		"Apply to kubernetes",
	}, true
}

func (v Cueify) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Task":
			return []string{}, true
		case "KubePkg":
			return []string{
				"KubePkg spec",
			}, true
		case "PkgName":
			return []string{
				"pkg name",
			}, true
		case "OutDir":
			return []string{
				"OutDir for cue files",
			}, true

		}
		if doc, ok := runtimeDoc(v.Task, names...); ok {
			return doc, ok
		}

		return nil, false
	}
	return []string{}, true
}

func (v KubePkg) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "TypeMeta":
			return []string{}, true
		case "ObjectMeta":
			return []string{}, true
		case "Spec":
			return []string{}, true
		case "Status":
			return []string{}, true

		}
		if doc, ok := runtimeDoc(v.TypeMeta, names...); ok {
			return doc, ok
		}
		if doc, ok := runtimeDoc(v.ObjectMeta, names...); ok {
			return doc, ok
		}

		return nil, false
	}
	return []string{}, true
}

func (v Manifests) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Task":
			return []string{}, true
		case "KubePkg":
			return []string{
				"KubePkg spec",
			}, true
		case "Recursive":
			return []string{
				"recursively extract KubePkg in sub manifests",
			}, true
		case "Manifests":
			return []string{
				"Manifests of k8s resources",
			}, true

		}
		if doc, ok := runtimeDoc(v.Task, names...); ok {
			return doc, ok
		}

		return nil, false
	}
	return []string{
		"Manifests extract manifests from KubePkg",
	}, true
}

func (v OciTar) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Task":
			return []string{}, true
		case "KubePkg":
			return []string{
				"KubePkg spec",
			}, true
		case "Platforms":
			return []string{
				"Platforms of oci tar, if empty it will based on KubePkg",
			}, true
		case "WithAnnotations":
			return []string{
				"WithAnnotations pick annotations of KubePkg as image annotations",
			}, true
		case "OutFile":
			return []string{
				"OutFile of OciTar",
			}, true
		case "File":
			return []string{
				"File of tar created",
			}, true

		}
		if doc, ok := runtimeDoc(v.Task, names...); ok {
			return doc, ok
		}

		return nil, false
	}
	return []string{}, true
}

func (v PushOciTar) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Task":
			return []string{}, true
		case "SrcFile":
			return []string{
				"SrcFile of oci tar",
			}, true
		case "RemoteURL":
			return []string{
				"RemoteURL of container registry",
			}, true

		}
		if doc, ok := runtimeDoc(v.Task, names...); ok {
			return doc, ok
		}

		return nil, false
	}
	return []string{}, true
}