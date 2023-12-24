package flasher

type Flasher interface {
	SetPort(port string) Flasher
	SetPath(path string) Flasher
	Init() error
	Upload() error
}
