// Code generated by go-bindata. (@generated) DO NOT EDIT.

// Package bindata generated by go-bindata.
// sources:
// hack/cmake/CMakeDetermineGoCompiler.cmake
// hack/cmake/CMakeGoCompiler.cmake.in
// hack/cmake/CMakeGoInformation.cmake
// hack/cmake/CMakeTestGoCompiler.cmake
// hack/cmake/golang.cmake
// hack/cmake/tars-tools.cmake
// hack/scripts/makefile.tars
// hack/scripts/makefile.tars.gomod
package bindata

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _cmakeCmakedeterminegocompilerCmake = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x54\x4d\x73\xab\x36\x14\xdd\xf3\x2b\xee\x28\x5e\xe0\x76\x6c\xb7\xcb\x66\x55\x85\xc8\x84\x29\x1f\xae\xc0\x9e\xb4\xc9\x1b\x8d\x82\x05\xd1\x18\xa4\x0c\x60\xbf\x0f\x8f\xff\xfb\x1b\xbe\x6c\x12\x93\xc7\x06\xe9\x4a\xf7\xde\xa3\x73\x8e\x74\x23\x13\x88\x75\xfe\x26\x33\x51\x80\xc8\xdf\xaa\xef\x50\xbd\x0a\x05\xa9\xa8\x60\x5f\x48\xe0\x6a\x0b\xa5\xa8\xc0\xf2\xf0\x3f\x84\xd9\x01\xb3\x02\x6f\xe5\xb8\x84\xc2\x06\x53\x43\x26\xa6\x1f\x44\xfd\xa2\x3e\x2f\x4e\x0d\x80\x6e\x6d\x42\xfc\xcd\x71\x90\x77\x82\x30\xa2\xe4\xdf\x35\x76\x01\xa1\x29\xdc\x10\x7f\x63\x0e\x96\xa7\x20\xcb\x16\x87\x01\x00\x35\x0c\x96\xc8\x4c\x28\x9e\x0b\x56\xe3\xd4\x4a\xa8\xca\xbc\x6a\xc8\x1c\xdf\x19\xeb\xb5\xa2\x81\x4d\xb1\xd7\xff\x19\xa6\x76\x78\x81\xbb\x74\xb1\x1d\x32\xe2\x6f\x9a\xf4\xa9\xd1\xb4\x94\x89\xf9\xe9\x06\x68\xbe\x52\x8c\x21\xc0\xd4\xfe\x13\xd0\xe4\xf8\x49\xf2\x09\x81\x85\xad\x07\x52\x1f\xdf\xf1\x6d\x40\x4b\x59\x94\x15\xf0\x22\xdd\xe7\x42\x55\x50\x69\xb0\xf5\x59\x0a\xd4\xf6\x12\x6a\x2b\x13\xf3\x02\xac\x26\x94\x3c\x3a\x61\x14\xc2\xa0\xd1\x3b\x12\x4e\x3d\xca\x5c\x94\x25\x4f\x85\x19\x12\xff\x9e\x11\x4a\x03\x0a\xc8\xd2\xfb\x6c\x0b\x4a\x57\x90\x48\xb5\xbd\x08\x5f\x2b\x2c\x15\x08\x75\x90\x85\x56\x0d\x9e\x03\x2f\x24\x7f\xc9\x04\x0c\xf8\xbc\x7d\x56\x57\x14\xcf\xaf\xa1\x0e\x87\xdf\x44\xbc\xaf\x04\x7b\x2b\x74\x2c\xca\xd2\xb4\x02\xcf\xc3\xfe\x3d\xa4\xba\x6e\x06\x76\xb0\xc2\xd1\x03\x04\xeb\x68\xb5\x8e\xd8\x06\x53\x07\xdf\xb9\xa4\x0b\x37\x05\xca\xaa\x90\x2a\x35\x29\xb1\xc9\x23\x50\xb2\x72\xb1\x45\x00\x3d\xab\x09\x02\x84\xfa\x7c\x34\x39\xb6\xa3\x13\x6a\x93\x44\x65\xda\x9a\xdd\x39\x3e\xab\xa3\x0d\xbe\xf3\x96\x76\xd6\x1e\xe3\x2a\x40\x83\x20\xba\x0e\x2c\x5e\xa4\x1a\x06\x2f\xa7\x6f\xa2\x8b\x7d\x59\x9c\xb7\x34\x93\x4c\xc7\x3c\x3b\x87\x1a\x4c\x43\x5b\xbd\x13\xac\xa5\x6f\xd4\x53\x9f\x8b\xdc\x79\xa9\x3d\xbd\xad\xc1\x1a\xfa\x46\x64\xa5\xe8\x4d\x53\xeb\x5c\xb3\x9f\x16\x3c\x1f\x69\xe0\x63\x8f\x84\xb5\x1a\x75\xa5\xda\x54\x03\xda\x3a\x27\x7d\x90\x10\xce\x1a\x8e\x80\x3b\x01\xa0\x83\x28\x4a\xa9\x15\x1a\x51\xd5\xc5\xbe\xcd\x36\x84\x86\x4e\xe0\x77\xf8\xda\xdb\xd0\xe9\xeb\xe1\xc8\x7a\x00\x94\xea\xa7\x3f\x66\x7f\x7d\xf9\x7d\xfe\xee\xf7\x04\x0b\x3c\xfb\x9f\xcf\x7e\xd4\xb3\xdf\x10\x74\x75\x5a\xf1\x87\x95\x4f\x9d\x23\xfb\x0b\x80\x66\x33\x88\x5e\x05\xd8\x3a\xe3\x2a\xbd\x98\x5e\x6e\x85\xaa\x64\x22\x63\x5e\x49\xad\xea\x77\x67\x72\xfc\x45\x09\xeb\x55\xc4\x3b\x48\x74\x01\x5f\x75\xb1\x93\x2a\xfd\x58\xf0\x76\x94\x91\x56\x91\xfe\x46\x9c\x07\x39\x2f\x76\x8c\x97\x8c\x6f\x0f\x5c\xc5\x62\x7b\xad\xcd\xd4\x30\x62\xad\x12\x99\xee\x0b\xd1\xbc\x81\x66\x5f\xde\x5a\x53\x4a\xfc\x88\x85\xc1\x9a\x5a\x84\xdd\x3b\xf4\xb4\x88\x73\xbe\x13\x0b\xcb\xe3\x3b\x61\xeb\xde\x0d\xf3\x26\x3a\x6f\x6c\xd8\x27\xaf\x5c\x1c\x2d\x03\xea\x31\xc7\x5f\x06\x6d\xee\x68\x16\xfc\x1d\xf8\xee\x7f\x53\xc3\x18\x7f\xed\xea\x67\x6d\x83\x29\xa0\xc1\x6d\x40\x53\xe3\x67\x00\x00\x00\xff\xff\x7c\x96\x46\x5a\x4f\x06\x00\x00")

func cmakeCmakedeterminegocompilerCmakeBytes() ([]byte, error) {
	return bindataRead(
		_cmakeCmakedeterminegocompilerCmake,
		"cmake/CMakeDetermineGoCompiler.cmake",
	)
}

func cmakeCmakedeterminegocompilerCmake() (*asset, error) {
	bytes, err := cmakeCmakedeterminegocompilerCmakeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmake/CMakeDetermineGoCompiler.cmake", size: 1615, mode: os.FileMode(420), modTime: time.Unix(1661308924, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _cmakeCmakegocompilerCmakeIn = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xce\xb1\x0a\xc2\x30\x14\x85\xe1\xbd\x4f\x71\xc9\x64\x17\x51\xf0\x01\x1a\xd2\xd3\x12\x9a\xe6\x86\xdb\xa4\xb8\xdd\x49\x1c\x3b\xe8\xfb\x23\x2e\x8a\xb1\xeb\xe1\xe3\xe7\x3c\x6e\xcf\x83\x9b\xed\x04\x1d\x37\x75\x3c\x27\x1f\x20\x64\xba\xbf\xad\x33\x6d\xb3\x8b\x35\xb0\xed\xd1\xd3\xb9\x6d\x7e\xc1\xc2\x45\x1c\x74\xf0\x01\x8a\x6b\x46\x5c\x3c\xc7\x85\xee\x5b\x15\x0a\x3e\x4e\x10\x4d\x82\x01\x82\xe8\x40\x97\x53\x45\xb8\xe4\x54\xf2\xb7\x42\xc7\x3a\x52\x0b\x15\xa4\x60\x1d\xde\xb7\xf6\x6f\x23\xae\xba\x5a\x21\x33\xf2\x67\x34\x6d\xf3\x0a\x00\x00\xff\xff\x2d\x4f\x79\xa3\x11\x01\x00\x00")

func cmakeCmakegocompilerCmakeInBytes() ([]byte, error) {
	return bindataRead(
		_cmakeCmakegocompilerCmakeIn,
		"cmake/CMakeGoCompiler.cmake.in",
	)
}

func cmakeCmakegocompilerCmakeIn() (*asset, error) {
	bytes, err := cmakeCmakegocompilerCmakeInBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmake/CMakeGoCompiler.cmake.in", size: 273, mode: os.FileMode(420), modTime: time.Unix(1661308924, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _cmakeCmakegoinformationCmake = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\xcd\xb1\xaa\x83\x30\x14\x87\xf1\x3d\x4f\xf1\xc7\x49\x07\xdf\x40\x04\x0d\x07\xf1\x1a\x93\x8b\x46\xe8\x96\xa1\x8d\x25\x34\xf5\x14\xea\xfb\x53\x5a\x87\x16\x4b\xe7\xef\x83\x5f\x98\x53\x6d\x2c\x64\x5f\x75\xe4\x1a\x76\xd2\xf4\xff\xad\x22\x67\xea\x3f\x92\x36\x13\xc0\xdd\xaf\xe9\x8f\x8c\xe4\xcc\x58\x99\x23\x8e\x7c\xbd\x85\xe8\x91\x47\xe4\x1a\x39\xa3\xd8\x8e\x12\xc5\x68\xa6\x41\x52\x89\x24\x13\x7e\x39\x85\x39\xcd\x84\xd8\xab\xaa\xd5\x9d\xa3\x03\xc9\xc9\x56\xb5\xa2\x2f\x76\xd7\xdf\x6e\x0c\xcb\xe5\xc5\xd9\x6a\x68\xe8\xc9\x6d\xee\x58\xe2\x13\x7c\x04\x00\x00\xff\xff\x66\x8f\x90\x26\xe6\x00\x00\x00")

func cmakeCmakegoinformationCmakeBytes() ([]byte, error) {
	return bindataRead(
		_cmakeCmakegoinformationCmake,
		"cmake/CMakeGoInformation.cmake",
	)
}

func cmakeCmakegoinformationCmake() (*asset, error) {
	bytes, err := cmakeCmakegoinformationCmakeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmake/CMakeGoInformation.cmake", size: 230, mode: os.FileMode(420), modTime: time.Unix(1661308924, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _cmakeCmaketestgocompilerCmake = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x4e\x2d\xd1\x70\xf6\x75\xf4\x76\x8d\x77\xcf\x8f\x77\xf6\xf7\x0d\xf0\xf4\x71\x0d\x8a\x0f\xf7\x0f\xf2\x0e\x56\x30\x54\x70\x76\x74\xf6\x70\x55\xf0\xf4\x0b\x71\x0d\xf2\x73\xf4\x51\x50\x52\xd2\xe4\x02\x04\x00\x00\xff\xff\xb5\x30\x45\x6b\x31\x00\x00\x00")

func cmakeCmaketestgocompilerCmakeBytes() ([]byte, error) {
	return bindataRead(
		_cmakeCmaketestgocompilerCmake,
		"cmake/CMakeTestGoCompiler.cmake",
	)
}

func cmakeCmaketestgocompilerCmake() (*asset, error) {
	bytes, err := cmakeCmaketestgocompilerCmakeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmake/CMakeTestGoCompiler.cmake", size: 49, mode: os.FileMode(420), modTime: time.Unix(1661308924, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _cmakeGolangCmake = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x55\xcd\x6e\xa3\x3c\x14\xdd\xf3\x14\x57\x34\x0b\xf8\xa4\xf6\x7b\x82\x2e\x1c\xf0\x50\xab\xe6\x67\x1c\xd3\x51\x57\x16\xc1\x4e\xca\x08\xf0\x28\x90\xaa\x55\xc4\xbb\x8f\x80\xa4\x24\x9d\x34\x29\xd3\xcc\xce\xc1\xc7\xf7\x9c\xfb\x77\x52\xa9\xda\x72\x7c\x74\x8f\x05\x62\xce\x1d\x79\xc0\x22\x8c\x79\x14\x73\xe1\x12\x86\x1d\x1e\xb2\x47\x98\x6c\x7a\x80\x13\x33\x86\x03\x2e\xa6\x24\x40\xec\xb1\x05\x34\xb6\x31\xbc\xa7\x64\xca\xda\xef\x7f\xfb\x9e\xc5\x01\x27\xfe\x58\x7e\xe3\x0a\x96\x1a\x96\xaa\x86\xc5\xba\x4c\xeb\x4c\x97\xc6\xee\x60\xe1\x97\x5a\xad\xca\x24\xf7\x74\xb4\xd2\x3f\x55\x5a\x0b\x24\x25\x70\xc4\x3c\xdb\x00\x48\xa4\x14\xe9\xba\xaa\x75\x21\xea\x64\xb5\x54\xb5\x35\xd9\xb4\x77\x0d\xa8\xf2\x19\xbc\x30\x42\xfc\xee\x76\xb2\xe9\x0f\xcd\x9b\x0a\x4f\x0b\x27\xf4\x23\x42\x31\x6b\x3a\xde\xc9\x06\x31\x2f\x68\x6c\x43\x95\xf2\x34\xb5\x6d\x0c\xda\x5a\xf6\xa5\x16\xea\x45\xa5\xeb\x3a\x99\xe7\x0a\x02\xe4\xe3\x56\xd7\x22\xcb\x95\xe5\xd1\x70\x0a\x5e\x28\x66\x61\xcc\x1c\x0c\x0c\x53\xc4\xc9\x03\x06\xf3\x7d\x31\x7a\x40\x57\x0c\x13\xcc\xff\x6e\x96\xda\xb4\x0d\x03\xe0\x0a\x0a\x55\x55\xc9\x52\x59\x6d\x0a\x5b\x58\x63\xbf\xbb\x39\x5d\xf9\xe6\xff\xc9\xa6\x55\xd5\x74\x11\xf7\xea\x95\xea\xa2\x48\x4a\x69\xf5\x0f\xe0\x7c\x9c\x9b\x3a\x93\xaf\x06\x00\x80\x13\xfa\x3e\x0a\xdc\xe3\xe5\x2c\xb4\x84\x03\x24\x0e\xf8\x90\xf2\x51\xa8\xd9\x61\x7f\x84\xec\x9e\x04\xde\x89\xa1\xa1\x64\xc6\x77\x23\xf3\xe5\x5c\x0a\x55\xd5\x49\xf1\xeb\x20\xa1\xcf\x0f\xcd\x7c\x9d\xe5\x12\xae\xf5\x90\xd9\xd9\x06\x98\x43\xa4\x50\x7c\xa3\xc8\x9b\xb5\xb1\x87\xbe\x7e\xa2\x60\x6f\xac\x9f\x26\x3d\xc3\xd9\x57\xde\xc5\x11\x0e\xdc\xd9\xa8\x19\xf8\x42\xb7\xde\x36\xb5\x57\x88\x28\x1d\xa7\x60\xdb\xb9\x61\x69\x0f\xb6\xf6\x8f\xa5\xdc\xdf\x58\xe4\xba\x6d\x25\xb6\x36\xd7\xad\x2b\x4c\x63\x42\x5d\xc1\x1f\xa3\x6e\x73\xb3\x85\x35\x7c\x80\x19\x67\xf8\x7b\x8c\x28\x98\x33\x8e\x38\x71\x4c\xbb\x4b\xbe\x35\xbc\x1e\xe5\x87\x2e\x86\xeb\xae\x2f\x85\x96\xea\x36\xbd\x4e\x56\xe9\x53\xf6\xac\x06\x20\x25\x53\xd1\x11\x99\x79\x36\xdf\x26\x7d\x93\x74\x91\x54\x5e\x29\xeb\x7c\xc8\xea\x29\x59\x29\xd9\xe3\xb2\x85\x85\xa2\x88\xe2\xfe\xd7\x09\x06\xf9\x9a\x67\xf3\xad\xde\x3d\x9e\x13\x2f\x2a\xbd\x83\x97\x32\x5b\x74\xf8\xdd\xe9\x32\x96\x76\x7a\x67\x87\x5e\xef\xcf\xd9\x18\xb7\x19\xe7\x37\x97\x75\x9c\x77\xea\x2f\xe1\x2e\x93\xcd\x30\x10\xcd\x81\xd7\x7c\xf4\x37\xdd\xae\xfd\xae\xb5\x23\xfd\xe6\xcb\xe2\x46\x69\x1b\x69\x4b\xc7\x47\xe3\x5f\x59\xd0\xf1\x56\x7e\x60\x37\x87\x8e\x62\x1b\xbf\x03\x00\x00\xff\xff\x4d\xb0\x1b\xde\x8c\x09\x00\x00")

func cmakeGolangCmakeBytes() ([]byte, error) {
	return bindataRead(
		_cmakeGolangCmake,
		"cmake/golang.cmake",
	)
}

func cmakeGolangCmake() (*asset, error) {
	bytes, err := cmakeGolangCmakeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmake/golang.cmake", size: 2444, mode: os.FileMode(420), modTime: time.Unix(1661308924, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _cmakeTarsToolsCmake = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\x4f\x8f\xe3\x4a\x11\x3f\xc7\x9f\xa2\xf1\xcc\xc1\x79\x22\x89\xb4\xdc\x46\x1a\x81\x27\xe9\xc9\x98\xf5\xd8\xc1\x71\x76\x76\xa5\x48\x96\xc7\xee\x24\xd6\x3a\xb6\x65\x77\xe6\xed\xbe\x51\x6e\x48\x80\x40\x7a\x20\x21\x0e\x5c\x10\x12\x02\x6e\xdc\x10\xe2\x89\x6f\xf3\x76\x97\x8f\x81\xaa\xdb\x7f\x63\x27\x93\xc9\xcb\xec\x2e\x82\x5c\x66\x6c\x57\x55\xff\xea\x4f\x57\x55\x57\x0b\x02\x79\x43\x9c\x15\x25\x56\x14\x87\x0e\x49\x12\xa9\xaf\x5f\x5f\xcb\xda\x00\xcd\x43\x44\x82\x3b\x34\xd4\x47\xb2\x79\x85\xf4\x89\x39\x9a\x98\xd6\x0b\xd9\x50\xe4\x0b\x15\xa7\xaf\xdb\x42\x42\x63\x2f\x98\x4b\x06\x1e\xe2\x97\xc8\xc0\x23\x55\xee\x63\x24\x4e\x83\x53\x11\x89\x62\xc6\x2c\x9e\xde\xf3\xff\xd6\x62\x5b\x48\x08\x95\xfa\xd7\xf2\x73\x6c\x5d\xeb\x83\x89\x8a\x2d\x46\x72\x7a\x5f\x7b\xb7\x06\x3e\xfe\xb6\x3f\x31\x0c\xac\x99\xd6\x58\x9f\x18\x7d\x6c\x0d\x14\x63\xdd\x73\x96\xf6\x6b\x22\xb6\x05\xc1\x0b\x1c\x7f\xe5\x12\xe9\x41\xe2\xde\x3c\xf4\xed\x60\xde\x65\x0f\x6d\xa1\x84\x44\x36\xfa\x57\xca\x0b\x6c\xa5\x5a\x0e\x14\x03\xf7\x4d\xdd\x78\x95\xc3\xba\x50\x34\xd9\x78\xc5\x65\xf9\xde\x6d\x59\x0b\x55\xb9\x30\xe0\xdb\x21\xbc\xc6\x44\x33\x95\xeb\x3d\xd7\xbd\xf5\x02\xd0\x76\x26\xdd\x28\xda\x0f\x9e\xb5\x85\x16\xc8\x31\x65\x63\xfc\xac\x3f\x1a\x95\x6c\x0c\x84\x3d\x6a\xc7\xc9\xb3\x79\xd8\x25\x6f\xc0\x46\x39\xa9\xb5\xe1\x8f\x06\x5a\xe2\x27\x44\xda\x4f\xfa\xbe\x92\x41\x6a\xe0\x7a\x33\x29\x35\x3a\xa3\xbf\xc1\x17\xd6\x95\x3e\x36\x21\x4e\xfa\x72\xff\x0a\xa3\xb1\x69\x28\xda\x10\x89\x09\xa1\xe8\x4b\x72\x8b\x16\x61\x42\xc5\xb6\xa0\x5c\xa2\x0d\x8e\xb1\x69\xe0\x9f\x4c\x64\x15\x89\x15\x04\x85\xc4\x05\xa5\xd1\x59\xaf\xf7\x25\xb9\xed\x02\x84\xae\x13\x2e\xc5\xb6\x80\xb5\x01\xc8\x2a\x83\x30\xf5\xe7\x58\xdb\x8e\x80\x86\xaf\x49\x00\x31\x76\x82\x72\x16\x03\xab\x58\x1e\xe3\x22\x38\xcb\x3e\x8a\x57\x41\x27\x26\x3e\xb1\x13\xd2\xcd\x02\x34\xe7\x9c\x8c\x54\x5d\x1e\x6c\x67\x5c\x45\x7e\x68\xbb\x75\x3e\x53\x36\xb6\x33\x51\x3b\xce\x39\x84\x13\x74\xa9\xa8\x58\xba\x31\x14\x13\xa3\xd3\xfb\x32\xde\x35\x12\xf1\x4b\xdc\x9f\x98\xd8\x1a\x19\x7a\x1f\x8f\xc7\xf9\x4e\xcf\x37\x0e\x7f\x5e\xa3\x0e\x46\xc4\x59\x84\x28\xd5\x04\xd9\xbe\xdf\x9e\x82\x1d\xea\xd2\xb9\x4e\x87\x08\xe7\xda\xee\x92\x6d\xca\xc6\x9a\xb9\x58\x98\xad\x02\x87\x7a\x61\x20\xcd\x49\x60\x81\x47\x91\x29\x1b\x43\x6c\xb6\x05\xa1\x35\xf3\x7c\x22\x0d\x55\xfd\xc2\x32\x70\x7f\x62\x8c\x31\x62\xcc\x8a\x36\x9a\x98\xe8\x0b\xe6\x7f\x20\x03\x6b\xea\x13\xd3\x62\x1f\xaf\x2c\x55\x19\x33\x6e\x6f\x96\x06\x17\xa3\x87\x37\xad\x59\x18\x13\xdb\x59\xf0\xd7\x80\x2a\xc3\xc3\x48\xd6\x6d\xa1\xd5\x6a\xcd\x09\xb5\x60\xe1\xc0\x5e\x12\xcb\x09\x97\x51\x18\x90\x20\x75\x97\x26\x5f\xe7\x2c\xc0\xbe\x46\xf0\xc6\xba\xc1\x0f\x72\xa6\x99\xb0\xc4\xc9\x93\x2c\xf0\xb1\xa4\x31\x31\x38\xfe\x21\xd6\x32\x3a\xbe\xd9\xd2\x07\x58\x68\xdd\x5d\x70\x0e\xdb\x75\x2d\x67\x95\xd0\x70\x09\xeb\x2c\xed\xc0\x95\xe0\x75\xab\xc5\xb3\x0d\x78\xa6\x24\x6f\xcd\xbf\xdd\xe8\xc6\x73\x45\x1b\x56\x52\x51\xb1\x0e\xa7\x29\x5c\x9b\xe5\x87\x35\xea\x84\x2b\xea\x7a\x31\x02\x73\x77\xa2\x38\xa4\xa1\x13\xfa\x10\xb5\xe9\x36\xd3\x55\xeb\x52\x95\x87\x6b\x31\x7f\xc7\x14\x14\x0b\x89\x58\x33\xb3\x6f\xbb\x44\xd6\x24\xa2\x06\x79\x03\x3c\xc2\xda\x60\x5c\x41\x58\x26\x03\x2a\x6e\x24\xdf\x4b\xa8\x24\x8f\x80\x1c\x6d\x84\xc7\xa6\x81\x18\x03\x09\xdc\x2c\x3e\xd8\x73\xc9\xc8\xd4\x8e\xe7\x84\x4a\x6c\x9d\x21\x36\xd7\x48\x56\x55\x54\x20\xd9\x90\xce\xc5\x25\x84\x5a\xae\x17\x13\x87\x86\xf1\x5b\x28\xc1\x11\x89\xa9\x47\x12\x69\x64\xe8\x23\x6c\x98\x0a\x1e\x23\x79\x30\x50\x4c\x45\xd7\x64\xd5\xe2\x3b\x49\xc5\xb2\xc6\xf4\x18\x83\xc1\x36\xe5\xc2\x7e\x69\xe5\xe9\x16\xf0\x66\x9b\x07\x12\xc4\x87\xdf\xfe\xe1\xfd\xcf\x7f\xfd\xee\x1f\x7f\x06\xab\xbe\xff\xdd\xcf\xbe\xfd\xe6\xef\x1f\x7e\xff\xd3\x77\x5f\xff\xed\xfd\x2f\xfe\xf2\xef\x3f\xfe\xea\xc3\x5f\x7f\xf9\xee\x9f\x5f\x0b\x4b\xdb\x89\x43\xb6\xd9\x12\x12\xdf\x91\x18\xc9\xa3\x51\x69\xcb\x41\x2c\xf2\x8d\x9f\xc2\x38\xbd\x97\x8d\x21\xb7\x50\x5a\x8d\x73\xad\x40\x99\xd3\xfb\x91\xa1\xff\x18\xf7\x2b\x35\x19\x88\xd3\x7c\x55\xd9\xbd\x63\xa3\x9f\x4a\x45\xe2\x17\x5d\x56\x3c\x04\x66\xe6\x79\x68\xf1\x5e\xc5\xbe\xf5\x49\x61\xe6\xf6\xde\x29\xc0\x9b\x55\xb6\x3a\xec\x45\x9e\x4d\x24\x16\x62\x65\x89\x6c\x41\x97\x44\x24\x70\x49\xe0\x70\x1d\x32\xb7\xd6\x88\x73\x63\xb7\x4e\x20\x11\x03\x01\x3a\x39\xd6\x0f\x84\xae\x12\x8a\x9c\x98\xd8\x94\x20\xba\x8c\x50\x1e\x30\xdf\x47\xf6\x5d\xe8\xb9\xc8\xf7\x82\xd5\x1b\xc4\xaa\x00\x72\xc2\x60\xe6\x7b\x0e\xfd\x9e\xd0\x1a\x63\x53\x32\x26\x1a\x44\x47\x96\x7e\x79\x3e\xdb\x55\x4c\x0a\xcd\xf2\xb2\xd2\xaa\xa4\xe6\x26\x89\x87\xe4\xff\xe5\x6b\xd8\xdc\x9d\xa8\xb1\xd7\xa1\xcb\xa8\x57\x98\x98\xd5\x07\x8e\x22\xdd\xaa\xc7\x83\x11\x2f\x51\x27\x9e\x7d\x3a\x10\x60\xe3\x22\x05\x7c\x62\x63\x38\x61\xf4\x16\x41\xe3\x56\x44\xfb\x83\x80\x7a\x1c\x91\xd0\x52\x2e\xf3\xce\xf4\xf8\xf0\xea\xc8\xa0\x5d\xdd\x1b\x5d\x0b\xab\x63\x2c\x7d\x24\x64\x8f\x40\x05\x1d\xa9\xd4\x6e\x41\x1b\x93\x56\x95\xb4\x49\x3c\xbd\x2f\x27\xd8\xf5\x93\x20\x2f\x5c\x9e\xad\xf6\x08\xe8\x15\x77\x3f\x91\x55\x0f\x81\x95\x3b\xfa\x73\xc2\x94\xb9\xb9\xdc\x3e\x6c\x73\xf4\x13\xec\x6d\x67\x01\xa9\x76\x1b\x5e\x56\xaf\x9c\xaf\x66\x77\xa8\xd8\x5b\x74\xfe\x15\xaa\xe9\xf1\x24\x06\xb5\xbc\x99\xe5\x7a\xb3\x19\x89\x49\x40\x1f\xb6\x69\x8a\xac\x4e\x55\xa1\xc8\x52\x52\x43\x0b\x5c\x34\xbf\x35\x11\x5d\xea\x2d\x49\x42\xed\x65\x24\x6c\xe9\x84\x6b\x3c\x42\xa5\x1d\xde\xd4\x70\xb4\xcd\x4c\xc2\x66\x8f\x0a\xc0\x85\x72\x23\xdc\xec\x93\x4d\xa5\x36\x5b\x4e\x28\xe0\xa5\x96\x73\x97\x8a\xf5\x40\x2b\x9d\xb9\xf6\x77\xe2\x56\x15\x33\x17\xf0\x96\x28\x3d\xec\x1d\xb1\x2b\xca\x9a\x9b\x74\xf7\x6c\xf6\x37\x59\xc3\xd9\x7c\xc6\x6e\x6a\x72\x6a\xe5\xab\xdc\xf0\x34\xac\x72\x48\xb2\x4d\x8d\x9c\x0d\x2a\xd6\x3d\x3b\xf2\x7a\x1c\x92\x65\x07\xae\x15\xad\x6e\x7d\x2f\x59\xa0\xce\x65\xb2\x4a\xc8\xf9\x8f\xaa\x61\xdf\xb9\xb4\xa3\xc8\xf7\x1c\x1b\xba\xf9\xf3\xd3\x7b\x99\x9d\x92\x2e\x97\xa1\xbb\xf2\x89\x05\x07\xca\xf3\xa2\x06\x75\x2e\x21\xe2\x49\x40\xcf\x5d\x72\x47\x7c\x38\x58\x74\xec\x15\x0d\x53\x0b\x64\xa9\xa9\xbe\xa7\x1f\xab\x6a\xe9\x14\x4a\x17\x5e\xec\x46\x76\x4c\xdf\xb2\xe9\x8f\xb3\x8a\xfd\xb4\x4a\xef\xa5\xf6\x0f\xa9\xe7\xbc\x26\xf4\x3c\x3f\xe2\x3d\xc7\xda\xfa\xbf\xc8\x16\x99\x97\xa7\x41\xe7\x78\xbf\xc6\xfe\xe5\xff\xa1\xb9\x87\xaa\x10\x7e\xff\x2b\xa1\xb7\xd5\xad\x4f\x10\x8b\x79\xdb\x0a\xd9\x9d\x8f\x0d\xde\xfd\xe6\x5f\xdf\x7e\xf3\xa7\xdd\xa5\x89\x2b\xfc\xa8\xca\x5a\x2b\x91\x50\xdf\xca\x25\x97\x1f\x79\xf3\x22\xd4\x64\xb3\x4a\x55\x4d\xcb\x10\xf7\x59\xb7\xea\x56\x3b\x70\x51\x1a\x0d\xdd\x6e\x57\xdc\x52\x1f\x0f\x98\x77\xee\x42\x97\x57\x49\xc4\xeb\x64\x36\x71\x3d\x62\xa1\x3c\x41\x59\xa9\x4c\x27\xc1\x7b\xd7\xca\x14\x4c\x43\xb1\x04\xa9\x9b\x93\xd3\xd6\x09\xda\x63\x74\x7a\x82\x0e\x1f\x9e\x3e\xc4\xbb\x7d\x7c\x0a\x9c\x8f\x1d\xa0\x02\x4f\xe9\xb2\x05\x1e\xeb\xc9\xb7\xc9\xa4\xdf\x6d\x00\xe0\x9c\x4d\xa7\xd4\x8e\x13\x36\xf0\x9c\x4e\xb3\xb9\xe7\x74\xca\x83\x16\xfe\x56\x27\x01\x05\xb0\x4a\xfe\x38\x02\x32\x7e\x3e\x8c\xb2\xf9\x48\x79\x16\xfa\x39\xc1\x4c\x8f\x65\x47\xc0\x47\xfc\x84\x78\x33\x00\xa8\xe2\x8f\xe8\xf2\x53\xac\xbd\xb8\xbf\xd2\xaf\xf1\xba\x97\x43\xee\x65\x88\x7b\x1c\x70\x6d\x00\xf4\x69\xdc\xfe\x99\x41\x6d\x74\xfd\xc1\x18\x53\xf7\x4f\x34\xe5\xe5\x47\xf4\x7e\x6f\x11\x2e\x49\x09\xe9\x67\xe5\xef\x4f\x0e\xae\xd1\xc3\x8f\x40\x95\xcd\xe6\xe1\xa1\x34\x7c\xd9\x5a\xa0\x18\xe1\xae\x36\x26\x2d\x8a\x5c\xfa\x9e\x9d\x0c\x90\x36\x1c\xf7\xe1\xf5\x96\x4e\xa6\xd1\x72\x05\x07\x6b\x66\x1c\xdb\xf7\x77\x91\x8b\x6d\x01\xa5\x3f\xc6\xd9\xd0\xcd\x1c\x72\x37\xbc\x13\x62\x6e\x79\x66\x77\x24\xb1\x8b\x7f\x7e\x9b\xd4\x16\x84\xba\x5d\xf3\xa6\x70\xdf\x9e\xb0\xc1\x5e\xe5\xbe\x8c\xdd\x81\xd7\x97\xc9\x9c\xb6\xbf\xcf\xb6\xf8\xa6\x62\xb5\x46\x8d\x78\x87\xfa\x5d\xd5\x31\x65\x76\x41\xb6\x24\x49\x62\xcf\x89\x24\x1e\xb1\x89\x2f\x7e\x62\xbb\x58\x80\x83\x2b\xae\xe7\xce\x50\xfe\xcb\x90\x97\xee\xee\xea\x9c\x85\x5a\x0d\x9c\x25\x9d\xcb\x9c\xf5\x4b\xc1\xb3\x82\xb3\xe1\xc6\xb0\x61\xd5\x89\xa2\x0e\x2c\xf3\xd5\x08\x37\xad\x9a\x7f\xac\xae\xaa\xca\xe6\xa5\x6e\x5c\x97\x38\x4a\xab\xa6\x1f\x2b\x1c\xd9\x1d\x72\x23\x47\x7e\xc1\xbc\xc9\x91\x9f\x39\xcf\x1a\x38\x8a\x03\x69\x8d\x8d\x9d\x3d\x6b\x4b\x55\x0e\xa6\x65\x9e\xa7\x8a\x0c\xe1\x3f\x01\x00\x00\xff\xff\x1c\xe0\x38\x88\x1a\x26\x00\x00")

func cmakeTarsToolsCmakeBytes() ([]byte, error) {
	return bindataRead(
		_cmakeTarsToolsCmake,
		"cmake/tars-tools.cmake",
	)
}

func cmakeTarsToolsCmake() (*asset, error) {
	bytes, err := cmakeTarsToolsCmakeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmake/tars-tools.cmake", size: 9754, mode: os.FileMode(493), modTime: time.Unix(1661308924, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _scriptsMakefileTars = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\x6d\x73\xa3\x38\x12\xfe\x0c\xbf\xa2\x8f\x25\x17\x2b\x63\xf0\x24\xa9\xba\xba\xb2\x8b\x9b\x21\x1e\xe2\xcc\x2e\x6b\x5c\xb6\xb3\x57\x57\xf1\x14\x4b\x40\xd8\xaa\x05\xc4\x81\xec\x64\x92\xf5\x7f\xbf\x92\x04\xd8\x24\x4e\x6e\xa7\x2a\xc9\x97\xc8\xad\x7e\x79\xba\xa5\x7e\xd4\xfc\x64\xbc\xed\x9f\xfa\x53\x4c\xee\x21\x5c\x52\x08\x69\x9a\x93\x04\x03\x2e\x0a\x5a\xa8\xf8\x3e\xa7\x05\x03\x77\xe8\xdb\xae\x0b\x00\x16\xe0\xcc\xbf\x9e\x99\xd7\xf3\x4b\xe3\x9f\xcd\xae\x3d\x1e\x01\x3c\xdb\x7d\x73\x8c\xea\xc8\x9b\xd8\xf3\x2b\xf8\x64\x81\xde\x29\x57\x38\x49\x60\x49\x01\x67\x1b\x90\x1b\x48\x1d\x79\x53\xcf\x9b\x1f\x54\xe0\x1b\x5c\x01\x2a\xa4\xfa\xa3\x94\x6d\x7b\xb7\x24\xeb\x2d\xa9\xfa\xf6\x78\x13\x72\x9b\x07\x6c\x65\xe9\x8f\xe5\xfa\xb6\x64\xd0\xef\x42\x57\xef\x54\x58\xb7\xea\xdc\x9e\xce\xce\x2a\x40\x7d\x8e\x38\x26\x45\xc9\xee\x68\x11\x71\xf4\x8d\x05\x34\x26\x08\x09\xac\x2c\x28\xca\x33\x0e\x98\xc4\xf8\xbf\xd0\xe9\x72\xc3\x2c\x2a\x59\x41\xb2\x25\x18\x74\xcd\x22\x52\x74\xf5\xce\xcf\x67\x23\xcf\xbf\x74\xed\x11\x42\x48\xe5\x31\x1a\x01\x00\x7c\xb0\x6a\x4d\x6b\x83\xb3\x88\x9f\x74\x16\x91\x58\x55\x27\x17\x35\x24\xe5\x47\x30\xe5\x05\x65\x34\x7c\x87\x1a\x8a\x2a\xf9\xb3\xe9\x70\x57\xa6\x3b\x92\x44\x61\x50\x44\x70\x62\xf2\x52\x20\x75\x32\xf5\x1a\x8d\x0f\x4f\x34\x04\x30\x7e\xee\x8d\xc6\x33\x27\x4b\x8a\xde\x1e\xb8\xc0\x1e\xd2\xfc\x7b\xcc\x9b\xc9\x02\x12\x03\xc3\x25\x03\xe3\x01\x34\xbd\x63\x4f\x26\x48\x83\x3f\xff\xdc\x97\xcd\xed\xe9\xc8\x99\x23\x6d\x00\x6c\x85\x33\x58\xa8\xd0\xfe\xc3\xe1\x8a\x82\x76\x73\x6c\x4f\x26\xc7\x40\x0b\x38\x96\x06\xc7\x40\x73\x46\x68\x06\xa4\x04\x9c\xe6\xec\xbb\xf9\x4d\x1b\x00\xbe\x27\x0c\x4e\x07\x6d\x2f\x0a\x4e\x4a\x0c\x0b\x55\x51\x6a\x41\x8d\xea\x6f\x60\xf0\x03\x3e\x43\x2f\x04\x57\x94\x2a\xfa\x98\x42\xb9\x0e\x57\x10\x91\xa2\xcf\xf5\xbb\x90\xd1\x3b\xb8\xc3\x10\x16\x38\x60\x18\x08\xe3\xe1\xa5\xb9\xa2\x28\x4a\xfa\x47\x44\x0a\x30\x72\xe9\x9c\x87\x56\x94\x98\xb4\x70\x35\xbe\x87\x34\xff\x0e\xa2\x5c\x7a\xe7\x14\x81\xf1\x2f\x61\xc4\xb3\x69\x69\x87\x39\x18\x1b\xa9\x22\x11\x0b\xaf\x31\x19\xa8\xaa\xed\xba\x7d\x68\x2a\xf9\x2e\x67\xda\x78\xef\x03\xbf\x99\x17\xd7\x5f\xdd\x2f\xa2\x1b\xf8\xfd\x42\xaa\xc2\x97\x08\x6e\xd7\x24\x89\xa4\x58\x68\xc8\x3e\x04\x83\x82\xfe\xf9\x7d\xee\x1a\x89\x33\x4e\x06\x7a\x87\xf3\x40\x2e\x8b\x20\xfa\x06\xa1\x2e\x52\x1b\xa8\xfd\xfd\x1d\x55\xf9\x2c\x4a\x6f\x60\xd0\x16\xf8\xe6\xfc\x7c\x70\x9a\xca\xed\xb3\x91\x87\x16\xf8\xe6\x63\x0a\x5c\xfe\x8f\xc1\x69\x0a\xfa\x63\x6d\xb7\x05\xb1\x65\x9a\xa6\xc6\xf3\xad\xf5\x61\x9f\x72\x5a\x61\xf8\xb5\xfb\x3f\x10\xb4\x8c\x02\x6f\x66\x71\xfe\x5a\x4d\x46\x4d\x52\x55\x87\xf3\x54\x26\x53\xaf\x71\x53\x8b\x0f\x25\x22\x19\xe9\x59\x12\x95\x49\x2b\x07\x69\xac\x77\x04\xf7\x21\xae\x74\x51\x27\xb2\x05\xbd\x13\x44\x51\x5e\x60\xfe\x3e\x1a\x86\xf0\xea\x0b\x56\xe7\x24\x58\xf2\xe7\x4f\xef\xf0\x3b\xbe\x43\x83\x10\x6a\x61\xd3\x3b\x31\x2d\x70\x10\xae\x44\x72\xdd\xbd\x6c\xf4\x0e\xde\x04\x09\xbc\x16\xbe\x15\xb2\x8a\xc4\xdd\xa0\xea\x1f\x52\x95\x98\x16\xb2\x6d\x48\xb6\x83\xd4\x80\xe1\x4d\x17\x51\xd1\x24\x11\x29\xb2\x20\xc5\x96\x2e\xfc\xf0\x25\xe8\xd2\x8b\xe8\xcc\x17\x11\x2c\xa9\x4f\xd7\xcc\xca\x93\xf5\x92\x64\xa5\xc5\xcf\xa9\xc8\xc3\xbe\xae\xd7\x5e\xda\x20\xf5\xb6\x73\x09\x20\xc3\xf2\x1a\xbc\x76\x7c\xfc\x12\x08\x47\xed\x5b\xf0\x1e\xed\xc2\x82\x62\x8f\x29\x40\xef\x0c\xbd\xf1\xe5\xd7\x11\x87\x42\x62\xb8\x91\x74\x58\xed\xfa\x2c\xcd\x7d\x5e\xf8\x6f\x0d\x3b\xd6\x9c\xc5\xa5\xab\xa0\xe4\x74\x5b\xb2\xfe\x33\x8b\x2e\x04\xb7\xb4\x60\xa6\x26\x49\x6a\x47\xca\x3b\x26\x96\x04\xf9\x3c\xd6\x33\x49\xaf\x91\x88\xc3\xe2\x2c\x58\xc4\x87\x52\x78\xcd\xb4\x27\x81\x84\x07\xb2\x1b\xf0\xfe\x03\xc3\xc0\xf7\x61\xb2\x8e\xb0\xa5\x99\xe5\x26\xd3\xf6\x05\xbe\x14\x84\x0f\x9b\xbd\xb8\x26\x5b\x3e\xc0\xbe\xff\x30\x02\xd3\x94\x61\x64\x25\xe3\xbd\xc7\x8d\x6b\x6b\xad\x3a\x2a\xe9\x06\x8c\x97\x1d\x9a\xbf\x47\xfc\x59\xf9\x70\xf4\x9f\xa3\xf4\x28\x3a\xba\x3a\xfa\xf5\x68\xf6\x3b\x57\x91\x11\xe4\x53\x22\x9c\xbc\x92\xb5\x70\x69\x56\xb9\x17\x69\xbb\x70\xbb\xfc\x77\xc7\xca\x2b\x11\x6e\xe2\x87\x27\xa8\x38\x5b\x08\xb5\x98\x54\x73\xc9\xbf\x9d\x0b\xff\xca\x9b\xcd\x01\xf8\xcc\xb9\x62\x2c\xef\xf7\x7a\x09\x0d\x83\x64\x45\x4b\xd6\x3f\xff\xf8\xf1\xe3\x4e\x71\xee\xfd\xe2\x8c\x85\xa2\xa6\xa9\xd7\x13\xd7\xb3\xbf\xf8\xd7\x33\x67\x2a\xde\xb5\xbd\x91\xf5\x6e\x45\x83\x94\xa0\x5a\xc5\x9b\x55\x4f\xdf\x27\x0b\x12\x92\xad\xef\xd5\x75\x9e\xd0\x20\xea\x43\x35\x7e\x8f\x3c\x6f\x66\xe9\x8f\x8d\xfa\xb6\x51\x60\x41\xd1\xf4\x56\x3b\x15\xc3\x30\x40\x8e\x21\x66\x7b\x03\xbc\x59\x5f\xcc\xc6\xde\x6c\xab\xa9\x4a\xb8\x2e\x92\x9a\xfa\xeb\x64\xb7\xbd\x20\x27\x3d\x19\xc3\x0f\xb2\xc8\xcf\xd7\xb7\x09\x29\x57\x9f\x18\x09\xff\xc0\xcc\xda\x53\x17\x29\x6f\xc1\xb8\x2c\xd7\x25\xb6\x3e\x8b\x9d\x91\x33\xdf\x4a\x0c\x97\x41\x9e\x27\x24\x0c\xf8\xf8\x62\xe9\x8f\xf6\x64\xc2\x55\x53\x1a\xad\x13\xec\x4b\xa6\xaa\x0d\xc0\xb8\x0c\x69\x9a\xe2\x8c\x59\x32\x30\x8e\x8c\xdb\xef\x46\x93\x35\xaf\xe3\xb6\xc9\x75\x91\xbd\x21\x57\x2c\x32\x4d\x55\xaf\x1c\x77\x22\x67\x4b\xbe\xf2\xe7\xf6\x14\xed\x3d\x4f\xcd\x73\xd6\x45\xaa\x3a\x73\xa6\xbf\x39\x53\x7f\x6c\xff\xea\xc8\x51\xb3\x1e\x47\x2a\x3e\x7b\x9f\x49\x33\xc1\x41\xd6\x57\xc5\xed\x16\xdd\xf4\xc5\x99\x38\xe3\x2f\xbe\x40\xe6\x5d\xfc\xcc\x59\xe1\xeb\xf8\x37\xef\x17\xc7\x3f\xb0\xe3\x7a\x43\xdb\xad\x7f\xec\x08\x45\xae\x66\x48\xba\xf3\x2f\xbf\xba\x0e\x7f\x1c\x86\xae\x63\x8f\xf9\x8f\x2d\x98\x27\x66\x64\xb2\x34\x87\x65\x4a\x33\x93\xae\x59\x83\xe0\xa4\x71\x74\xc2\x0f\xbb\x42\x18\x24\xc9\x4b\x20\xaf\xd0\x13\xc1\x70\x32\x79\x2a\x3a\x04\xf7\x6a\xe8\x48\x5b\xbe\xa8\x6c\x7e\x24\x85\x13\x93\xee\xd2\xe0\x8b\xbf\x90\xcb\x93\xdb\x20\xbc\xa1\x83\x52\xdb\x75\x91\xd4\x97\xbf\xe5\xb7\xa7\xb6\xc8\x16\xf8\xe6\x74\x70\x7e\x9e\x8a\xb2\x88\x59\xa4\xbf\x60\x0b\x76\x53\xe0\x94\x6e\xf0\x0b\x67\xf2\x4d\x53\x5b\xae\x0f\xb8\x0a\x92\xa4\xf6\x76\x23\x04\xf0\x77\x28\x52\x91\x58\x6d\x3d\xb7\xa7\xf5\x47\xf0\x9e\x35\x0b\x8a\x3d\x18\x77\x24\x49\x20\xa2\x70\xcc\xa9\xf0\x29\x37\x4f\x1d\xd7\xb1\x67\x8e\xac\xe6\xf1\x37\x4d\x55\x57\x38\xc9\xfb\x7b\x53\x98\x2c\x01\xd2\x44\xef\xbc\xf9\x37\xe1\xff\x02\x00\x00\xff\xff\xfc\xf1\x74\x31\xfd\x10\x00\x00")

func scriptsMakefileTarsBytes() ([]byte, error) {
	return bindataRead(
		_scriptsMakefileTars,
		"scripts/makefile.tars",
	)
}

func scriptsMakefileTars() (*asset, error) {
	bytes, err := scriptsMakefileTarsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/makefile.tars", size: 4349, mode: os.FileMode(493), modTime: time.Unix(1661308924, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _scriptsMakefileTarsGomod = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\xff\x73\x9b\x38\x16\xff\xd9\xfc\x15\xef\xbc\xe4\x6c\x65\x0d\xae\x93\x9b\xce\x8e\x3d\xbe\x0b\x75\x5c\x27\x3d\xc7\x78\xb0\xdb\x99\x9b\xa4\xc3\x12\x10\x58\x53\x40\x1c\xc8\x4e\xdb\x34\xff\xfb\x8d\x24\xc0\xe0\x90\x6c\xbb\xdb\xcc\x9c\x7f\x68\xd1\xd3\xd3\x47\xef\x9b\x3e\x7a\xca\x2f\xda\xcf\xfd\x29\xbf\xf8\xe4\x33\xb8\x01\x05\x97\x46\x09\x09\x31\xe0\x34\xa5\xa9\x82\x3f\x27\x34\x65\x30\x9f\xd8\xc6\x7c\x0e\x00\x63\xc0\xb1\xfd\x7e\xa5\xbf\x5f\xbf\xd5\x7e\x2b\x67\x8d\xc5\x0c\xe0\xd1\xec\x4f\xb7\x51\x99\x99\x4b\x63\x7d\x01\xff\x1a\x83\xda\xcd\x36\x38\x0c\x21\xa0\x80\xe3\x1d\xc8\x09\xa4\xcc\x4c\xcb\x34\xd7\x8d\x0a\x7c\x82\x2b\x40\x6e\xa9\x7a\x2f\x65\x0f\xfd\x5b\x12\xf7\x03\xaa\xcc\x4c\xfb\xea\x72\x61\x5a\xf6\x87\xa9\xb5\xba\x34\x17\xb0\x07\x51\xbb\x33\x13\xc1\x0e\xa7\x19\xa1\x31\x7c\x03\x77\xcb\x40\x73\x61\xf0\x0f\xad\x18\x78\x1d\xe8\x80\xe6\x0f\xf6\x63\x9d\x8f\x4f\xd0\x23\xd8\xc1\x6b\x18\xc3\xe0\xb5\xf2\xf3\xc3\x13\x92\xdb\xc4\x61\x9b\xb1\x7a\x9f\x6d\x6f\x33\x06\xc3\x1e\xf4\xb8\xe5\x22\x34\x0f\xca\xcc\xbc\x32\xcf\xdf\xcf\xa7\x0b\xe3\x6a\x3a\xdc\xfb\xb6\xc1\x8e\x07\x5a\x3c\x80\x80\xea\x11\xf5\xe0\x1b\x38\x77\x9f\xa0\x73\x9f\xa4\x24\x66\xa0\xaa\x27\x0f\x1d\xa4\x28\xc4\xc7\xff\x85\x6e\x4f\xed\xfa\x24\xf6\x32\x96\x92\x38\x00\x8d\x6e\x99\x47\xd2\x9e\xda\x7d\x77\x32\x33\xed\xb7\x73\x63\x86\x10\x52\x78\x78\x4b\x01\x00\xfc\x3a\x2e\x34\xc7\xcc\x49\x33\x2d\x49\x29\xa3\x2e\x0d\x15\x1c\x7b\xc4\x7f\x02\x3a\xa2\xde\x36\xc4\xdf\x03\x2d\x35\xc7\x3c\x9d\x7b\xff\x1e\x0a\xf0\xe5\x9b\x93\x3c\xe5\x2d\xe1\xb3\x4f\xd2\x8c\xdd\xd1\xd4\xe3\xfe\x97\x51\x82\x32\x4c\x08\x89\x72\x90\x36\xbe\x40\x8e\x94\xb5\x61\xad\xec\x95\x35\x11\x36\x09\x93\xee\x48\xe8\xb9\x4e\xea\xc1\xb1\xce\xe3\x83\x94\xa5\x65\x96\x1a\xbf\x1e\x68\x08\xc3\x44\x51\x15\x1a\x8f\x40\x02\x8a\x7e\xbe\xe1\xc2\x76\x97\x26\x5f\x7c\xce\x0d\x63\x20\x3e\x30\x9c\x31\xd0\xbe\x42\x5b\xed\x1a\xcb\x25\x6a\xc3\xb7\x6f\x55\xd9\xda\xb0\x66\xd3\x35\x6a\x8f\x80\x6d\x70\x0c\x37\x0a\xd4\x7f\xd8\xdd\x50\x68\x5f\x77\x8c\xe5\xb2\x03\x34\x85\x8e\x5c\xd0\x01\x9a\x30\x7e\xca\x48\x06\x38\x4a\xd8\x17\xfd\x63\x7b\x04\xf8\x33\x61\x30\x18\xd5\x51\x5a\x38\xcc\x30\xdc\x28\xad\x56\x21\x28\xac\xfa\x1b\x68\x3c\xc1\x27\xe8\x89\xcd\x5b\xad\x7c\xf7\x05\x85\x6c\xeb\x6e\xc0\x23\xe9\x90\xeb\xf7\x20\xa6\x77\x70\x87\xc1\x4d\xb1\xc3\x30\x10\xc6\xb7\x97\xcb\x5b\xad\x56\x2b\xfa\xe4\x91\x14\xb4\x44\x82\xf3\xad\x5b\x2d\x9f\xd4\xec\x2a\xb1\x27\x34\xf9\x02\x22\x5c\x6a\x77\x80\x40\xfb\xa7\x58\xc4\xbd\xa9\x69\xbb\x09\x68\x3b\xa9\x22\x2d\x16\xa8\x3e\x19\x29\x8a\x31\x9f\x0f\xa1\x8c\xe4\x8b\xe4\xb4\x44\x1f\x02\xaf\xcc\x37\xef\x2f\xe7\xe7\xb0\xb4\x4c\xf9\xc1\x8f\x05\x2f\x34\xa4\xb4\x24\x05\x72\x82\x60\xc4\xfb\x52\x8c\x6f\xb7\x24\xf4\xa4\x9a\x58\x21\x4f\x2a\x68\x14\xd4\xb3\x97\x29\x42\xe2\xc7\x9c\x2e\xd4\x2e\x67\x8a\x44\x46\x47\x1c\x28\x84\x7a\x48\x29\x7d\x18\x56\x67\x80\x9f\xab\x93\x80\x2a\xad\x33\x99\x9b\x9b\x57\xa7\xa7\xd7\xa7\xa7\xa3\x41\x24\xb5\x4e\x66\x26\x12\xb2\x57\x11\xc8\xb9\xd7\xa3\x41\x04\xea\x7d\x01\xf1\x00\xf9\xb4\xae\xeb\x6d\xee\x7c\xb1\x0a\xaa\x0c\x55\xdd\x53\xe1\xc5\xf9\x84\x3d\xa5\x1d\x31\x15\xa6\x89\x2a\x69\xef\xf9\x30\xf7\x30\xe7\x01\xee\x57\x91\x11\x0e\x53\x88\x41\x92\x94\x16\xe0\x58\x0b\x68\x7d\xa4\x71\xd8\x34\x71\x1b\x3c\x96\x7a\x8d\xde\xe6\xc8\x07\xce\x4a\x04\xb5\x2b\xa8\x14\x71\xb5\x37\x85\xc7\x0f\xa0\x76\x1d\xcf\x4b\x52\xcc\xbb\x07\x4d\x52\xbb\x2d\x2e\x21\xce\xa9\x19\x6f\x0e\xd4\x2e\x3f\x32\x7b\xb3\x11\x42\x95\x11\x8f\xa5\x4f\x53\xec\xb8\x1b\x11\x85\x5e\xc5\x6d\xb5\x8b\x77\x4e\x08\xcf\x6d\x5f\xdb\x32\xdf\x89\xc3\xa0\xfc\x3f\xa4\xb4\x7c\x9a\xca\x53\x48\xe2\xbd\x49\xa5\x31\xfc\x0c\x7b\x54\x9c\x39\x8f\xa4\xb1\x13\xe1\xb1\x2a\x70\xf8\x27\xa8\x12\x45\x1c\xf4\x27\x2d\x08\xa8\x4d\xb7\x6c\xac\xaa\xc5\x2a\xad\x92\x80\x47\x53\x55\x7b\xd5\xfa\x3e\xd2\x96\x18\xcb\xd2\x69\x4a\x79\xb5\x70\x04\x50\xbd\x72\x5e\xe2\xbc\x31\x27\xad\x70\x10\xa8\xdd\x89\xb9\x78\x7b\x39\xe3\xa6\x10\x1f\xae\x25\xd1\xe6\xb3\x36\x8b\x12\x9b\xe7\xe0\x63\xc9\xbb\x05\x1b\x72\xe9\xc6\xc9\x38\x91\x67\x6c\xf8\x68\x45\x0f\x9c\x5b\x9a\x32\xbd\x2d\xe9\x6f\x4f\xf7\x7b\x8e\xaf\x50\xef\xc1\xe2\x7e\x29\x11\x89\xe2\x84\x9a\xfa\x4d\x36\x3f\xb7\xb4\x2f\x77\x76\x1b\xdc\x19\xf1\x43\x0a\x9a\x86\x3f\xbb\xe1\xd6\xc3\xe3\xb6\x9e\xed\xe2\x76\x55\x60\x4b\x81\xfb\x75\x57\xd9\x57\x67\xc1\x57\xa8\xe2\xbb\x1e\xe8\xba\xdc\x46\x86\xce\xaf\xdc\x93\x5c\xbb\x5d\x0b\x5c\x2b\xda\x81\xf6\x34\xa0\xfe\xbb\xc7\x6f\xa8\x5f\x8f\xfe\x73\x14\x1d\x79\x47\x17\x47\x57\x47\xab\xdf\xb9\x8a\xdc\x41\xde\x4a\x02\xe4\x19\xaf\x05\xa4\x9e\xfb\x9e\x46\xf5\xc0\xed\xfd\xdf\xe7\x91\x47\xc2\xdd\xf9\x5f\x0f\xac\xe2\x4c\x21\xd4\x7c\xa2\x28\xca\xc5\x74\xbe\x94\x9d\x0b\xff\xb2\xd7\x86\x85\x2a\xb4\x56\xd2\x60\x0f\x29\xca\x6a\x6a\x7d\x98\x5a\x36\xef\xd9\x64\x23\x53\x5c\x76\x79\x4d\xbf\x4c\x1f\x13\x62\x27\x1e\x2a\xc2\x61\x11\xe0\xf3\xe9\x72\xba\x38\xb7\x85\x65\xe6\x9b\x77\xbc\x50\x2e\x17\x1f\xcc\x7f\x4f\xed\x86\x99\xb9\x39\x31\xe6\xc5\x60\x5f\x63\xf2\x6b\x85\x24\x9c\xfd\xf6\x72\x3e\xe5\x5c\x31\x99\x4f\x8d\x05\x1f\x3c\x80\x7e\xac\x7b\x3a\x8b\x12\x08\x22\x1a\xeb\x74\xcb\x4a\x0b\x8e\x4b\xa0\x63\x1e\xce\xdc\x42\x27\x0c\x9f\x32\xf2\x02\x1d\x08\x26\xcb\xe5\xa1\xa8\xc9\xdc\x8b\xc9\x54\xae\xe5\x1f\xf9\x9a\x1f\x71\xe1\x58\xa7\x7b\x37\xf8\xc7\x77\xf8\x72\x50\x0d\x02\x0d\x35\x4a\x8d\xf9\x1c\x49\x7d\x39\x96\x0f\xb5\xf6\x4d\x2c\xee\xa3\xc1\xe8\xf4\x34\x12\x81\xc9\xaf\xa7\xe1\x0d\xbb\x61\xd7\x29\x8e\xe8\x0e\x3f\x91\x97\x8f\x6d\xa5\x06\xdf\x08\xe7\x84\xe1\x1e\xf1\x5a\x88\xe0\xef\x90\x46\xc2\xc1\x02\x61\x6d\x58\xc5\xcb\xb1\x86\xc0\x9c\xb4\x66\xce\x1d\x09\x43\xf0\x28\x74\xf8\x49\x39\x3c\xba\xd6\x74\x3e\x35\x56\x53\x19\xd9\xce\xc7\xb6\xa2\x6c\x70\x98\x0c\x0b\x5a\xd7\x70\x1e\x0e\xd4\xbe\x89\xdb\x2f\xf0\xfa\xc8\xfb\xa0\x61\xf1\xe6\x2a\x1f\x81\x77\x1b\xe2\x6e\x8a\x36\x89\xdf\x9a\x67\xf7\xfc\x34\x67\x98\x71\xa3\x38\xad\xe6\x4f\xfd\x99\x39\x18\x0c\xe4\x3b\x6b\x4c\x63\x71\xe4\xf3\x56\xc8\x5e\x5f\x2d\xed\xf3\x4b\x8b\xdf\x9f\xd1\x27\x86\xa3\x04\x34\x4f\x30\x32\xe7\x54\xb5\xd0\x3a\xbf\xb4\xb8\x2c\xa0\xa2\x97\x24\x31\x61\xc0\xa2\x84\x8b\x04\x27\x8a\x56\xf2\xca\x78\xb7\x7f\x2f\x23\xd0\x02\xd6\x20\x1f\xbc\x46\x25\x5d\x82\x44\x24\x71\xc6\x1c\xfe\xe6\x27\x6c\xb3\xbd\xd5\x5d\x1a\xf5\xd7\x4e\x9a\x4d\x42\xba\xf5\xc4\xd7\x8c\xf6\xb9\x8f\x7d\x46\x69\x98\xf5\x73\x77\xcf\x42\x87\x3f\x1a\x46\x95\xcb\x26\xa0\x10\x60\xf6\xe7\x71\x7c\xc2\xff\x2d\xf8\x54\x3d\x88\x90\x88\xe7\x83\x92\x4b\xc7\xcd\x7f\xca\x10\xcf\xd0\xa2\x6f\x2d\x7b\xc9\xaa\x7e\x3d\x67\x45\x0f\x50\x6b\x03\x9f\x4a\x74\x4d\xe9\x07\xd3\xbd\xb4\xcc\xb5\x39\xb1\x67\xd3\x85\xfd\x87\x49\x6f\xd4\x7d\xf9\xf4\x53\x1a\x84\x58\x0f\x68\xe8\xc4\x81\x4e\xd3\x40\x3e\xe7\x6f\xb7\x7e\xdf\x8d\xbc\x7e\xcd\xf9\xb3\xdd\x40\x3f\xf9\x4d\x7f\xd5\x94\xfd\x3f\x09\x73\x90\xfc\xc6\x18\xe4\x25\x50\x9b\x7b\xae\x10\x6a\x9b\x95\xfd\x61\xc3\xda\xa6\xfc\x36\x96\x46\xd1\xa0\x7e\x57\x89\x14\xca\x7f\xad\x54\x0c\x6b\x65\x2d\x27\x3f\x56\x32\xf5\x35\xff\x2f\xcc\xd1\x18\x9c\xbf\xcc\x23\xcf\xa2\xc2\x1f\x56\x56\x3d\x54\x4d\x15\x56\xe8\x7c\x77\xa5\x95\xcf\xc8\xc7\x15\xf7\x08\xeb\xb9\xb2\xc9\x2b\xf0\x7f\x01\x00\x00\xff\xff\xf9\x62\xf0\x5d\xb5\x16\x00\x00")

func scriptsMakefileTarsGomodBytes() ([]byte, error) {
	return bindataRead(
		_scriptsMakefileTarsGomod,
		"scripts/makefile.tars.gomod",
	)
}

func scriptsMakefileTarsGomod() (*asset, error) {
	bytes, err := scriptsMakefileTarsGomodBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/makefile.tars.gomod", size: 5813, mode: os.FileMode(493), modTime: time.Unix(1661769677, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"cmake/CMakeDetermineGoCompiler.cmake": cmakeCmakedeterminegocompilerCmake,
	"cmake/CMakeGoCompiler.cmake.in":       cmakeCmakegocompilerCmakeIn,
	"cmake/CMakeGoInformation.cmake":       cmakeCmakegoinformationCmake,
	"cmake/CMakeTestGoCompiler.cmake":      cmakeCmaketestgocompilerCmake,
	"cmake/golang.cmake":                   cmakeGolangCmake,
	"cmake/tars-tools.cmake":               cmakeTarsToolsCmake,
	"scripts/makefile.tars":                scriptsMakefileTars,
	"scripts/makefile.tars.gomod":          scriptsMakefileTarsGomod,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("nonexistent") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"cmake": &bintree{nil, map[string]*bintree{
		"CMakeDetermineGoCompiler.cmake": &bintree{cmakeCmakedeterminegocompilerCmake, map[string]*bintree{}},
		"CMakeGoCompiler.cmake.in":       &bintree{cmakeCmakegocompilerCmakeIn, map[string]*bintree{}},
		"CMakeGoInformation.cmake":       &bintree{cmakeCmakegoinformationCmake, map[string]*bintree{}},
		"CMakeTestGoCompiler.cmake":      &bintree{cmakeCmaketestgocompilerCmake, map[string]*bintree{}},
		"golang.cmake":                   &bintree{cmakeGolangCmake, map[string]*bintree{}},
		"tars-tools.cmake":               &bintree{cmakeTarsToolsCmake, map[string]*bintree{}},
	}},
	"scripts": &bintree{nil, map[string]*bintree{
		"makefile.tars":       &bintree{scriptsMakefileTars, map[string]*bintree{}},
		"makefile.tars.gomod": &bintree{scriptsMakefileTarsGomod, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
