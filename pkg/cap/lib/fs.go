package lib

type FileHandle struct{}

type FileOpts struct{}

func Open(filepath string, opts *FileOpts) FileHandle {
	// TODO: Throwing error from Go
	return FileHandle{}
}

func (handle *FileHandle) Read() []byte {
	// TODO: Throwing error from Go
	return []byte{}
}
