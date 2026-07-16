package scaffold

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	defaultRootModelDir     = "models"
	defaultRootRequestDir   = "requests"
	defaultRootResourceDir  = "resource"
	defaultRootSeedDir      = "seed"
	defaultRootMigrationDir = "migration"

	// Repo Pattern
	defaultRootControllerDir = "controller"
	defaultRootServiceDir    = "service"
	defaultRootRoutesDir     = "routes"
	defaultRootRepositoryDir = "repository"
)

var commandRootDirs = map[string]string{
	"m":         defaultRootModelDir,
	"migration": defaultRootMigrationDir,
	"requests":  defaultRootRequestDir,
	"resource":  defaultRootResourceDir,
	"seed":      defaultRootSeedDir,

	"controller": defaultRootControllerDir,
	"service":    defaultRootServiceDir,
	"routes":     defaultRootRoutesDir,
	"repository": defaultRootRepositoryDir,
}

var technicalCommands = map[string]bool{
	"uuid_use": true,
	"id_use":   true,
}

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

func resolveFileDir(commands map[string]bool, rootPath string) (map[string]string, error) {
	allPaths := make(map[string]string)

	commands["controller"] = true
	commands["service"] = true
	commands["routes"] = true
	commands["repository"] = true

	for key, enabled := range commands {
		if !enabled {
			continue
		}

		rootDir, exists := commandRootDirs[key]
		if !exists {
			if technicalCommands[key] {
				continue
			}

			return nil, fmt.Errorf("comando inválido: %s", key)
		}

		var fullPath string

		switch key {
		case "m", "seed", "migration", "requests", "resource":
			fullPath = rootDir

		default:
			fullPath = rootPath + "\\" + rootDir
		}

		validatedPath, err := validatePathByKey(fullPath)
		if err != nil {
			return nil, fmt.Errorf("erro ao validar o diretório de %s: %w", key, err)
		}

		allPaths[key] = validatedPath
	}

	return allPaths, nil
}
