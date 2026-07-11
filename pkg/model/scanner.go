package model

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	defaultRootModelDir      = "models"
	defaultRootRequestDir    = "requests"
	defaultRootResourceDir   = "resource"
	defaultRootSeedDir       = "seed"
	defaultRootMigrationDir  = "migration"
	defaultRootControllerDir = "controller"
)

var commandRootDirs = map[string]string{
	"model":      defaultRootModelDir,
	"migration":  defaultRootMigrationDir,
	"requests":   defaultRootRequestDir,
	"resource":   defaultRootResourceDir,
	"seed":       defaultRootSeedDir,
	"controller": defaultRootControllerDir,
}

// func existsPath(path string) (bool, error) {
// 	_, err := os.Stat(path)

// 	if err == nil {
// 		return true, nil
// 	}

// 	if errors.Is(err, os.ErrNotExist) {
// 		return false, nil
// 	}

// 	return false, err
// }

func validatePathByKey(path string) (string, error) {
	fullPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	if err := os.MkdirAll(fullPath, 0755); err != nil {
		return "", err
	}

	return fullPath, nil
}

func resolveFileDir(commands map[string]bool) (map[string]string, error) {
	allPaths := make(map[string]string)

	for key, enabled := range commands {
		if !enabled {
			continue
		}

		rootDir, exists := commandRootDirs[key]
		if !exists && key != "uuid_use" && key != "id_use" {
			return nil, fmt.Errorf("comando inválido: %s", key)
		}

		validatedPath, err := validatePathByKey(rootDir)
		if err != nil {
			return nil, fmt.Errorf("erro ao validar o diretório de %s: %w", key, err)
		}

		allPaths[key] = validatedPath
	}

	return allPaths, nil

}
