package arduino

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"github.com/arduino/arduino-cli/commands/compile"
	"github.com/arduino/arduino-cli/commands/core"
	"github.com/arduino/arduino-cli/commands/lib"
	climonitor "github.com/arduino/arduino-cli/commands/monitor"
	"github.com/arduino/arduino-cli/commands/upload"
	"github.com/arduino/arduino-cli/rpc/cc/arduino/cli/commands/v1"
	rpc "github.com/arduino/arduino-cli/rpc/cc/arduino/cli/commands/v1"
	"github.com/spf13/afero"
	"github.com/thefarmhub/farmhub-cli/internal/arduino/cli/feedback"
	"github.com/thefarmhub/farmhub-cli/internal/arduino/cli/instance"
)

type Arduino struct {
	instance *commands.Instance
	fbqn     string
}

var Fs afero.Fs = afero.NewOsFs()

func NewArduino() *Arduino {
	return &Arduino{
		instance: instance.CreateAndInit(),
	}
}

func (a *Arduino) SetFBQN(fbqn string) {
	a.fbqn = fbqn
}

func (a *Arduino) InstallLibrary(req *rpc.LibraryInstallRequest) error {
	req.Instance = a.instance

	return lib.LibraryInstall(context.Background(), req, feedback.ProgressBar(), feedback.TaskProgress())
}

func (a *Arduino) GitLibraryInstall(req *rpc.GitLibraryInstallRequest) error {
	req.Instance = a.instance

	err := lib.GitLibraryInstall(context.Background(), req, feedback.TaskProgress())
	if err != nil {
		if strings.Contains(err.Error(), "already installed") {
			return nil
		}

		return err
	}

	return nil
}

func (a *Arduino) PlatformInstall(req *rpc.PlatformInstallRequest) (*rpc.PlatformInstallResponse, error) {
	req.Instance = a.instance

	return core.PlatformInstall(context.Background(), req, feedback.ProgressBar(), feedback.TaskProgress())
}

func (a *Arduino) Compile(path string) error {
	ctx := context.Background()
	compileReq := &rpc.CompileRequest{
		Fqbn:       a.fbqn,
		Instance:   a.instance,
		SketchPath: path,
	}

	var outStream, errStream bytes.Buffer
	_, err := compile.Compile(ctx, compileReq, &outStream, &errStream, nil)

	return err
}

func (a *Arduino) Upload(portAddress, path string) error {
	ctx := context.Background()
	uploadReq := &rpc.UploadRequest{
		Fqbn:       a.fbqn,
		Port:       &rpc.Port{Address: portAddress},
		Instance:   a.instance,
		SketchPath: path,
	}

	var outStream, errStream bytes.Buffer
	_, err := upload.Upload(ctx, uploadReq, &outStream, &errStream)

	return err
}

func (a *Arduino) Monitor(ctx context.Context, portAddress string) error {
	monitorReq := &rpc.MonitorRequest{
		Port:     &rpc.Port{Address: portAddress},
		Instance: a.instance,
		Fqbn:     a.fbqn,
	}

	_, _, err := climonitor.Monitor(ctx, monitorReq)

	return err
}

// This prepares a sketch folder or file for compilation
// with requirements for arduino like naming conventions and folder
// structure setup
func PrepareSketch(path string) (string, error) {
	fileInfo, err := Fs.Stat(path)
	if err != nil {
		return "", err
	}

	if fileInfo.IsDir() {
		inoFile := filepath.Join(path, filepath.Base(path)+".ino")
		_, err := Fs.Stat(inoFile)
		if err != nil {
			return "", errors.New("no .ino file found in the directory")
		}

		return path, nil
	}

	tempDir, err := afero.TempDir(Fs, "", "sketch")
	if err != nil {
		return "", err
	}

	destFile := filepath.Join(tempDir, filepath.Base(tempDir)+".ino")
	err = copyFile(path, destFile)
	if err != nil {
		return "", err
	}

	fmt.Println("Copied", path, "to", destFile)

	return tempDir, nil
}

func copyFile(src, dest string) error {
	srcFile, err := Fs.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := Fs.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}
