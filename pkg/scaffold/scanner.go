package scaffold

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	defaultRootModelDir     = "models"
	defaultRootRequestDir   = "requests"
	defaultRootResourceDir  = "resource"
	defaultRootSeedDir      = "seed"
	defaultRootMigrationDir = "migration"

	// Repo Pattern
	defaultRootRepositoryDir = "repository"
	defaultRootRoutesDir     = "routes"
	defaultRootControllerDir = "controller"
	defaultRootServiceDir    = "service"
)

var commandRootDirs = map[string]string{
	"m":        defaultRootModelDir,
	"M":        defaultRootMigrationDir,
	"requests": defaultRootRequestDir,
	"resource": defaultRootResourceDir,
	"seed":     defaultRootSeedDir,

	"repo":       defaultRootRepositoryDir,
	"routes":     defaultRootRoutesDir,
	"controller": defaultRootControllerDir,
	"service":    defaultRootServiceDir,
}

var technicalCommands = map[string]bool{
	"uuid_use":         true,
	"id_use":           true,
	"create_repo_path": true, // Add to pass in command validation,
}

func validatePathByKey(path string) (string, error) {
	fullPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	if err := createInformedPath(fullPath, 0755); err != nil {
		return "", err
	}

	return fullPath, nil
}

func resolveFileDir(commands map[string]bool, rootPath string) (map[string]string, error) {
	allPaths := make(map[string]string)

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
			fullPath = filepath.Join(rootPath, rootDir)
		}

		validatedPath, err := validatePathByKey(fullPath)
		if err != nil {
			return nil, fmt.Errorf("erro ao validar o diretório de %s: %w", key, err)
		}

		allPaths[key] = validatedPath
	}

	return allPaths, nil
}

func getModuleName() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	fileByte, err := os.ReadFile(filepath.Join(wd, "go.mod"))
	if err != nil {
		return "", err
	}

	firstLine := strings.TrimSpace(string(bytes.SplitN(fileByte, []byte("\n"), 2)[0]))
	if firstLine == "" {
		return "your_module", nil
	}

	parts := strings.Fields(firstLine)

	if len(parts) < 2 || parts[0] != "module" {
		return "your_module", nil
	}

	return parts[1], nil
}
