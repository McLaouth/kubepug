package formatter

import (
	"fmt"

	"github.com/fatih/color"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/rikatz/kubepug/pkg/results"
)

type stdout struct{}

func newSTDOUTFormatter() Formatter {
	return &stdout{}
}

var (
	gvColor        = color.New(color.FgHiBlack).Add(color.Bold).SprintFunc()
	resourceColor  = color.New(color.FgRed).Add(color.Bold).SprintFunc()
	globalColor    = color.New(color.FgBlue).Add(color.Bold).SprintFunc()
	namespaceColor = color.New(color.FgCyan).Add(color.Bold).SprintFunc()
	errorColor     = color.New(color.FgWhite).Add(color.BgRed).Add(color.Bold).SprintFunc()
	locationColor  = color.New(color.FgHiMagenta).Add(color.Bold).SprintFunc()
)

func (f *stdout) Output(data results.Result) ([]byte, error) {
	s := fmt.Sprintf("%s:\n%s:\n\n", resourceColor("RESULTS"), resourceColor("Deprecated APIs"))

	for _, api := range data.DeprecatedAPIs {
		s = fmt.Sprintf("%s%s found in %s/%s\n", s, resourceColor(api.Kind), gvColor(api.Group), gvColor(api.Version))

		if api.Description != "" {
			s = fmt.Sprintf("%s\t ├─ %s\n", s, api.Description)
		}

		items := stdoutListItems(api.Items)
		s = fmt.Sprintf("%s%s\n", s, items)
	}

	s = fmt.Sprintf("%s\n%s:\n\n", s, resourceColor("Deleted APIs"))

	for _, api := range data.DeletedAPIs {
		s = fmt.Sprintf("%s%s found in %s/%s\n", s, resourceColor(api.Kind), gvColor(api.Group), gvColor(api.Version))
		s = fmt.Sprintf("%s\t ├─ %s\n", s, errorColor("API REMOVED FROM THE CURRENT VERSION AND SHOULD BE MIGRATED IMMEDIATELY!!"))

		items := stdoutListItems(api.Items)
		s = fmt.Sprintf("%s%s\n", s, items)
	}

	return []byte(s), nil
}

func stdoutListItems(items []results.Item) string {
	s := ""
	for _, i := range items {
		var fileLocation string
		if i.Location != "" {
			fileLocation = fmt.Sprintf("%s %s", locationColor("location:"), i.Location)
		}

		if i.Scope == "OBJECT" {
			if i.Namespace == "" {
				i.Namespace = metav1.NamespaceDefault
			}

			s = fmt.Sprintf("%s\t\t-> %s: %s %s %s %s\n", s, namespaceColor(i.Scope), i.ObjectName, namespaceColor("namespace:"), i.Namespace, fileLocation)
		} else {
			s = fmt.Sprintf("%s\t\t-> %s: %s %s\n", s, globalColor(i.Scope), i.ObjectName, fileLocation)
		}
	}

	return s
}
