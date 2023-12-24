package flasher

type Flasher interface {
	SetPort(port string)
	SetPath(path string)
	Init() error
	Upload() error
}
