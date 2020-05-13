package golang

import (
	"fmt"
	"go/build"
	"os"
	"path/filepath"

	"github.com/VinnieApps/cicd-toolbox/internal/arrayutil"
)

// ListPackages lists all go packages in the specified path and subpath, ignoring the list of
// packages specified.
func ListPackages(path string, packagesToIgnore []string) ([]*build.Package, error) {
	packages := make([]*build.Package, 0)

	walkErr := filepath.Walk(path, func(subPath string, info os.FileInfo, err error) error {
		pkg, readErr := build.Import(fmt.Sprintf("./%s", subPath), ".", build.ImportComment)
		if readErr == nil && !arrayutil.ContainsString(packagesToIgnore, pkg.Name) {
			packages = append(packages, pkg)
		}

		return nil
	})

	if walkErr != nil {
		return nil, walkErr
	}

	return packages, nil
}
