package helm

import (
	"path/filepath"

	"k8s.io/helm/pkg/lint/rules"
	"k8s.io/helm/pkg/lint/support"
)

// Validate lints a helm chart directory
func Validate(basedir string) support.Linter {
	chartDir, _ := filepath.Abs(basedir)
	linter := support.Linter{ChartDir: chartDir}
	rules.Chartfile(&linter)
	return linter
}
