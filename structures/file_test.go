package structures

import (
	"testing"
	"unsafe"
)

var files = []struct {
	fd   int
	path string
	f    *file
}{
	{-11, "./test.java", nil},
	{1012, "./test.md", NewFile(1012, "./test.md")},
}

// this test checks if the factory method on the file struct behaves properly for both positive and negative scenarios.
func TestFileFactoryMethod(t *testing.T) {
	t.Logf("Size of file struct : %d\n", unsafe.Sizeof(file{}))

	for _, file := range files {
		newFile := NewFile(file.fd, file.path)
		if file.f != nil {
			if newFile.fd != file.f.fd || newFile.fileName != file.f.fileName {
				t.Errorf("There was a problem creating file expected : %v actual : %v \n", file.f, newFile)
			}
		} else {
			// means that the file.f == nil
			if newFile != file.f {
				t.Errorf("The file generated should give nil as the fd < 0")
			}
		}
	}
}
