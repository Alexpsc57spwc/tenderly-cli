package util

import (
	"io"
	"os"
	"path/filepath"

	"github.com/tenderly/tenderly-cli/typescript"
	"github.com/tenderly/tenderly-cli/userError"
)

func MustSaveTsConfig(directory string, config *typescript.TsConfig) {
	err := typescript.SaveTsConfig(directory, config)
	if err != nil {
		userError.LogErrorf(
			"unexpected error writing tsconfig.json",
			userError.NewUserError(err, "Unexpected error writing tsconfig.json."))
		os.Exit(1)
	}
}

func MustLoadTsConfig(directory string) *typescript.TsConfig {
	tsconfig, err := typescript.LoadTsConfig(directory)
	if err != nil {
		userError.LogErrorf("failed to load tsconfig.json: %s",
			userError.NewUserError(err, "Failed to load tsconfig.json."),
		)
		os.Exit(1)
	}

	return tsconfig
}

func TsConfigExists(directory string) bool {
	return ExistFile(filepath.Join(directory, typescript.TsConfigFile))
}

func TsFilesExists(directory string) (bool, error) {
	found := false
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == typescript.TsFileExt {
			found = true
			// return EOF because we just want to find at least one file with .ts extension
			return io.EOF
		}
		return nil
	})
	if err == io.EOF {
		err = nil
	}
	if err != nil {
		return false, err
	}

	return found, nil
}

func MustSavePackageJSON(directory string, config *typescript.PackageJson) {
	err := typescript.SavePackageJson(directory, config)
	if err != nil {
		userError.LogErrorf(
			"unexpected error writing package.json",
			userError.NewUserError(err, "Unexpected error writing package.json."))
		os.Exit(1)
	}
}

func MustLoadPackageJSON(directory string) *typescript.PackageJson {
	packageJSON, err := typescript.LoadPackageJson(directory)
	if err != nil {
		userError.LogErrorf("failed to load package.json: %s",
			userError.NewUserError(err, "Failed to load package.json."),
		)
		os.Exit(1)
	}

	return packageJSON
}

func PackageJSONExists(directory string) bool {
	return ExistFile(filepath.Join(directory, typescript.PackageJsonFile))
}
