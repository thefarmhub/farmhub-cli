package arduino

import (
	"errors"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestPrepareSketch(t *testing.T) {
	Fs = afero.NewMemMapFs()

	tests := []struct {
		name          string
		setup         func() string // function to set up the test environment
		expectedError bool
		checkOutput func(output string) error
	}{
		{
			name: "valid directory with .ino file",
			setup: func() string {
				dirPath := "/valid_sketch"
				_ = Fs.Mkdir(dirPath, 0755)
				inoFilePath := filepath.Join(dirPath, "valid_sketch.ino")
				afero.WriteFile(Fs, inoFilePath, []byte("test content"), 0644)
				return dirPath
			},
			expectedError: false,
			checkOutput: func(out string) error {
				if out != "/valid_sketch" {
					return errors.New("expected output to be /valid_sketch but received: "+out)
				}

				return nil
			},
		},
		{
			name: "directory without .ino file",
			setup: func() string {
				dirPath := "/invalid_sketch"
				_ = Fs.Mkdir(dirPath, 0755)
				return dirPath
			},
			expectedError: true,
			checkOutput: func(out string) error {
				if out != "" {
					return errors.New("expected output to be blank")
				}

				return nil
			},
		},
		{
			name: "ino file outside of valid directory",
			setup: func() string {
				dirPath := "/invalid-folder"
				_ = Fs.Mkdir(dirPath, 0755)
				inoFilePath := filepath.Join(dirPath, "valid_sketch.ino")
				afero.WriteFile(Fs, inoFilePath, []byte("test content"), 0644)
				return inoFilePath
			},
			expectedError: false,
			checkOutput: func (out string) error {
				matchPattern := `/sketch\d+$`

				r, err := regexp.Compile(matchPattern)
				if err != nil {
					return err
				}

				if !r.MatchString(out) {
					return errors.New("expected output to match a temporary file pattern but received: " + out)
				}

				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := tt.setup()

			output, err := PrepareSketch(path)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.NoError(t, tt.checkOutput(output))
		})
	}
}
