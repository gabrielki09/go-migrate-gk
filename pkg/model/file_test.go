package model

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNormalizeWithUnderline(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "space separated",
			in:   "Financial Account",
			want: "financial_account",
		},
		{
			name: "dash separated",
			in:   "financial-account",
			want: "financial_account",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := normalizeWithUnderline(tt.in)
			if got != tt.want {
				t.Fatalf("normalizeWithUnderline(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestToPascalCase(t *testing.T) {
	t.Parallel()

	got := toPascalCase("financial_account")
	want := "FinancialAccount"

	if got != want {
		t.Fatalf("toPascalCase(%q) = %q, want %q", "financial_account", got, want)
	}
}

func TestBuildModelContentContainsPackage(t *testing.T) {
	t.Parallel()

	content := buildModelContent("Financial Account", "ID        int")

	if !strings.Contains(content, "package financialaccountmodel") {
		t.Fatalf("model content does not contain expected package:\n%s", content)
	}
}

func TestBuildControllerContentContainsStructAndConstructor(t *testing.T) {
	t.Parallel()

	content := buildControllerContent("Financial Account")

	wantParts := []string{
		"type FinancialAccountController struct",
		"func NewFinancialAccountController(service FinancialAccountService) *FinancialAccountController",
	}

	for _, want := range wantParts {
		if !strings.Contains(content, want) {
			t.Fatalf("controller content does not contain %q:\n%s", want, content)
		}
	}
}

func TestResolveFileDirUsesRequestAndResourceDirectories(t *testing.T) {
	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	tempDir := t.TempDir()
	if err := os.Chdir(tempDir); err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err := os.Chdir(currentDir); err != nil {
			t.Errorf("failed to restore working directory: %v", err)
		}
	})

	paths, err := resolveFileDir(map[string]bool{
		"requests": true,
		"resource": true,
	})
	if err != nil {
		t.Fatal(err)
	}

	wantRequests := filepath.Join(tempDir, "requests")
	if paths["requests"] != wantRequests {
		t.Fatalf("requests path = %q, want %q", paths["requests"], wantRequests)
	}

	wantResource := filepath.Join(tempDir, "resource")
	if paths["resource"] != wantResource {
		t.Fatalf("resource path = %q, want %q", paths["resource"], wantResource)
	}
}
