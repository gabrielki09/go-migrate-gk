package migration

import "testing"

func TestParseMigrationFileName(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		fileName      string
		wantVersion   string
		wantName      string
		wantDirection string
		wantOK        bool
	}{
		{
			name:          "valid up migration",
			fileName:      "20260711120000_create_users.up.sql",
			wantVersion:   "20260711120000",
			wantName:      "create_users",
			wantDirection: "up",
			wantOK:        true,
		},
		{
			name:          "valid down migration",
			fileName:      "20260711120000_create_users.down.sql",
			wantVersion:   "20260711120000",
			wantName:      "create_users",
			wantDirection: "down",
			wantOK:        true,
		},
		{
			name:     "missing direction",
			fileName: "20260711120000_create_users.sql",
			wantOK:   false,
		},
		{
			name:     "invalid direction",
			fileName: "20260711120000_create_users.create.sql",
			wantOK:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			version, name, direction, ok := parseMigrationFileName(tt.fileName)
			if ok != tt.wantOK {
				t.Fatalf("ok = %v, want %v", ok, tt.wantOK)
			}

			if version != tt.wantVersion {
				t.Errorf("version = %q, want %q", version, tt.wantVersion)
			}

			if name != tt.wantName {
				t.Errorf("name = %q, want %q", name, tt.wantName)
			}

			if direction != tt.wantDirection {
				t.Errorf("direction = %q, want %q", direction, tt.wantDirection)
			}
		})
	}
}
