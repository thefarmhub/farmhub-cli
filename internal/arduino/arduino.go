package arduino

import (
	"bytes"
	"context"

	"github.com/arduino/arduino-cli/commands/compile"
	"github.com/arduino/arduino-cli/commands/lib"
	"github.com/arduino/arduino-cli/commands/upload"
	"github.com/arduino/arduino-cli/rpc/cc/arduino/cli/commands/v1"
	rpc "github.com/arduino/arduino-cli/rpc/cc/arduino/cli/commands/v1"
	"github.com/thefarmhub/farmhub-cli/internal/arduino/cli/feedback"
	"github.com/thefarmhub/farmhub-cli/internal/arduino/cli/instance"
)

type Arduino struct {
	instance *commands.Instance
	fbqn     string
}

func NewArduino(fbqn string) *Arduino {
	return &Arduino{
		instance: instance.CreateAndInit(),
		fbqn:     fbqn,
	}
}

func (a *Arduino) InstallLibrary(name, version string) error {
	req := &rpc.LibraryInstallRequest{
		Instance: a.instance,
		Name:     name,
		Version:  version,
	}

	return lib.LibraryInstall(context.Background(), req, feedback.ProgressBar(), feedback.TaskProgress())
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
