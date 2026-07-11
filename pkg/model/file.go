package model

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Input: Financial Account
// Output: financial_account
func normalizeWithUnderline(name string) string {
	name = strings.TrimSpace(name)
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, "-", "_")
	name = strings.ReplaceAll(name, " ", "_")

	return name
}

// Input: Financial Account
// Output: financialaccount
func normalizeNoWithUnderline(name string) string {
	name = strings.TrimSpace(name)
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, "-", "")
	name = strings.ReplaceAll(name, " ", "")
	name = strings.ReplaceAll(name, "_", "")

	return name
}

func toPascalCase(value string) string {
	value = strings.TrimSpace(value)
	value = strings.ToLower(value)
	value = strings.ReplaceAll(value, "-", "_")
	value = strings.ReplaceAll(value, " ", "_")

	parts := strings.Split(value, "_")

	var builder strings.Builder

	for _, part := range parts {
		part = strings.TrimSpace(part)

		if part == "" {
			continue
		}

		builder.WriteString(strings.ToUpper(part[:1]))
		builder.WriteString(part[1:])
	}

	return builder.String()
}

func buildModelContent(model, usageId string) string {
	modelFileName := normalizeWithUnderline(model)
	packageName := strings.ReplaceAll(modelFileName, "_", "") + "model"
	structName := toPascalCase(model) + "Model"

	content := fmt.Sprintf(`package %s

import "time"

type %s struct {
	%s
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
	`, packageName, structName, usageId)

	return content
}

func createModelFile(file File, option Options) error {
	var modelId string

	if option.Command["uuid_use"] {
		modelId = "Uuid      string"
	} else {
		modelId = "ID        int"
	}

	modelFileName := fmt.Sprintf("%s_model.go", normalizeWithUnderline(file.Name))
	modelFileContet := buildModelContent(file.Name, modelId)

	if err := createFileWithContent(modelFileName, modelFileContet, file.FilePaths["model"]); err != nil {
		return fmt.Errorf("erro ao criar o arquivo do model: %w", err)
	}

	return nil
}

// 20060102150405_create_tableName.up.sql
// 20060102150405_create_tableName.down.sql
func buildMigrationContent(fileName, usageId string) (
	upFileName,
	downFileName,
	upContent,
	downContent string,
	err error,
) {
	version := time.Now().Format("20060102150405")

	upFileName = fmt.Sprintf("%s_create_%s.up.sql", version, fileName)
	downFileName = fmt.Sprintf("%s_create_%s.down.sql", version, fileName)

	upContent = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
	%s,
	created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
	deleted_at TIMESTAMPTZ NULL
);`, fileName, usageId)

	downContent = fmt.Sprintf(`DROP TABLE IF EXISTS %s;`, fileName)

	return upFileName, downFileName, upContent, downContent, nil
}

func createMigrationFile(file File, option Options) error {
	var migrationId string

	if option.Command["uuid_use"] {
		migrationId = "id UUID PRIMARY KEY DEFAULT gen_random_uuid()"
	} else {
		migrationId = "id BIGSERIAL PRIMARY KEY"
	}

	migrationUpFileName, migrationDownFileName, migrationUpContent, migrationDownContent, err := buildMigrationContent(file.Name, migrationId)
	if err != nil {
		return err
	}

	if err := createFileWithContent(migrationUpFileName, migrationUpContent, file.FilePaths["migration"]); err != nil {
		return fmt.Errorf("erro ao criar o arquivo .up da migration: %w", err)
	}

	if err := createFileWithContent(migrationDownFileName, migrationDownContent, file.FilePaths["migration"]); err != nil {
		return fmt.Errorf("erro ao criar o arquivo .down da migration: %w", err)
	}

	return nil
}

func buildControllerContent(controller string) string {
	packageName := normalizeNoWithUnderline(controller) + "controller"
	pascalController := toPascalCase(controller)

	pascalControllerWithService := pascalController + "Service"
	pascalControllerWithController := pascalController + "Controller"

	content := fmt.Sprintf(`package %s

import (
	"net/http"
)

type %s interface {
}

type %s struct {
	service %s
}

func New%s(service %s) *%s {
	return &%s{
		service: service,
	}
}

func (c *%s) GetAll(w http.ResponseWriter, r *http.Request) {
}

func (c *%s) Create(w http.ResponseWriter, r *http.Request) {
}

func (c *%s) FindByID(w http.ResponseWriter, r *http.Request) {
}

func (c *%s) Update(w http.ResponseWriter, r *http.Request) {
}

func (c *%s) Delete(w http.ResponseWriter, r *http.Request) {
}
`,
		packageName,
		pascalControllerWithService,
		pascalControllerWithController,
		pascalControllerWithService,
		pascalControllerWithController,
		pascalControllerWithService,
		pascalControllerWithController,
		pascalControllerWithController,

		pascalControllerWithController,
		pascalControllerWithController,
		pascalControllerWithController,
		pascalControllerWithController,
		pascalControllerWithController,
	)

	return content
}

func createControllerFile(file File, option Options) error {
	controllerFileName := fmt.Sprintf("%s_controller.go", normalizeWithUnderline(file.Name))
	controllerFileContent := buildControllerContent(file.Name)

	if err := createFileWithContent(controllerFileName, controllerFileContent, file.FilePaths["controller"]); err != nil {
		return fmt.Errorf("erro ao criar o arquivo do controller: %w", err)
	}

	return nil
}

func buildRequestContent(request string) string {
	packageName := normalizeNoWithUnderline(request) + "request"
	pascalRequest := toPascalCase(request)
	pascalRequestWithRequest := pascalRequest + "Request"

	content := fmt.Sprintf(`package %s

	type %s struct {
	}

	func (r %s) ValidatePayload() error {

		return nil
	}
	
`, packageName, pascalRequestWithRequest, pascalRequestWithRequest)

	return content
}

func createRequestFile(file File, option Options) error {
	requestFileName := fmt.Sprintf("%s_request.go", normalizeWithUnderline(file.Name))
	requestFileContent := buildRequestContent(file.Name)

	if err := createFileWithContent(requestFileName, requestFileContent, file.FilePaths["requests"]); err != nil {
		return fmt.Errorf("erro ao criar o arquivo do controller: %w", err)
	}

	return nil
}

func buildSeedContent(seed string) string {
	packageName := normalizeNoWithUnderline(seed) + "seed"
	pascalSeed := toPascalCase(seed)
	pascalSeedWithSeed := pascalSeed + "Seed"

	contentSql := `
		INSERTO INTO your_table_name ()
		VALUES ()
	`
	content := fmt.Sprintf(`package %s

	import (
		"context"

		"github.com/jackc/pgx/v5/pgxpool"
	)

	type %s struct {
		db *pgxpool.Pool
		ctx context.Context
	}

	func (s %s) %s() error {
		for i := 0; i < 50; i++ {
			if _, err := s.db.Exec(
				s.ctx,
				%s,
				args,
			); err != nil {
				return err
			}

		}

		return nil
	}
	
`, packageName, pascalSeedWithSeed, pascalSeedWithSeed, pascalSeedWithSeed, contentSql)

	return content
}

func createSeedFile(file File, _ Options) error {
	seedFileName := fmt.Sprintf("%s_seed.go", normalizeWithUnderline(file.Name))
	seedFileContent := buildSeedContent(file.Name)

	if err := createFileWithContent(seedFileName, seedFileContent, file.FilePaths["seed"]); err != nil {
		return fmt.Errorf("erro ao criar o arquivo da seed: %w", err)
	}

	return nil
}

func buildResourceContent(seed string, usageId string) string {
	packageName := normalizeNoWithUnderline(seed) + "response"
	pascalResource := toPascalCase(seed)
	pascalWithResource := pascalResource + "Response"

	createdAtWithJson := fmt.Sprintf("CreatedAt time.Time `json:%s`", `"created_at"`)
	updatedAtWithJson := fmt.Sprintf("UpdatedAt time.Time `json:%s`", `"updated_at"`)
	deletedAtWithJson := fmt.Sprintf("DeletedAt time.Time `json:%s`", `"deleted_at"`)

	content := fmt.Sprintf(`package %s

import "time"

type %s struct {
	%s
	%s
	%s
	%s
}
	`, packageName, pascalWithResource, usageId, createdAtWithJson, updatedAtWithJson, deletedAtWithJson)

	return content
}

func createResourceFile(file File, option Options) error {
	var usageId string

	if option.Command["uuid_use"] {
		usageId = fmt.Sprintf("Uuid      string    `json:%s`", `"uuid"`)
	} else {
		usageId = fmt.Sprintf("ID        int `json:%s`", `"id"`)
	}

	resourceFileName := fmt.Sprintf("%s_response.go", normalizeWithUnderline(file.Name))
	resourceFileContent := buildResourceContent(file.Name, usageId)

	if err := createFileWithContent(resourceFileName, resourceFileContent, file.FilePaths["seed"]); err != nil {
		return fmt.Errorf("erro ao criar o arquivo do resource: %w", err)
	}

	return nil
}

func createFileWithContent(fileName, fileContent, filePath string) error {
	fullPath := filepath.Join(filePath, fileName)
	osFile, err := os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		return err
	}

	defer osFile.Close()

	if _, err := osFile.WriteString(fileContent); err != nil {
		return err
	}

	return nil
}

func createFiles(file File, option Options) error {
	if err := createModelFile(file, option); err != nil {
		return err
	}

	fileCreators := map[string]func(File, Options) error{
		"migration":  createMigrationFile,
		"controller": createControllerFile,
		"requests":   createRequestFile,
		"seed":       createSeedFile,
		"resource":   createResourceFile,
	}

	for key, enabled := range option.Command {
		if !enabled {
			continue
		}

		creator, exists := fileCreators[key]
		if !exists {
			continue
		}

		if err := creator(file, option); err != nil {
			return err
		}
	}

	return nil
}
