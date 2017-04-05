package structures

type file struct {
	fd       int    // file discriptor
	fileName string // path and name of the file.
}

// NewFile is the constructor that is used to get file type
// since the file is private to this package we will be forced to make file from the constructor.
func NewFile(fd int, fileName string) *file {
	if fd < 0 {
		return nil
	}
	return &file{fd, fileName}
}
