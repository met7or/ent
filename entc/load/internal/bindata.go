// Package internal Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// template/main.tmpl
// schema.go
package internal

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

var _templateMainTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x91\x51\x4f\xdb\x30\x14\x85\x9f\xe3\x5f\x71\x16\x31\x91\xb0\xe2\x02\x6f\x9b\xd4\x07\x04\x9d\xd4\x69\x83\x49\x45\xda\x03\x43\xc8\x75\x6e\x5a\x8b\xd4\xce\xae\x5d\xb4\xca\xca\x7f\x9f\xec\xb4\x6c\x7b\x4b\x7c\xbe\x7b\xce\xb9\x76\x8c\xd3\x33\x71\xe3\xfa\x3d\x9b\xf5\x26\xe0\xea\xe2\xf2\xe3\x79\xcf\xe4\xc9\x06\x7c\x56\x9a\x56\xce\xbd\x60\x61\xb5\xc4\x75\xd7\x21\x43\x1e\x49\xe7\x57\x6a\xa4\x78\xd8\x18\x0f\xef\x76\xac\x09\xda\x35\x04\xe3\xd1\x19\x4d\xd6\x53\x83\x9d\x6d\x88\x11\x36\x84\xeb\x5e\xe9\x0d\xe1\x4a\x5e\x1c\x55\xb4\x6e\x67\x1b\x61\x6c\xd6\xbf\x2e\x6e\xe6\x77\xcb\x39\x5a\xd3\x11\x0e\x67\xec\x5c\x40\x63\x98\x74\x70\xbc\x87\x6b\x11\xfe\x09\x0b\x4c\x24\xc5\xd9\x74\x18\x84\x88\x11\x0d\xb5\xc6\x12\xca\xad\x32\xb6\xc4\x30\x88\xe9\x14\x37\xa9\xcf\x9a\x2c\xb1\x0a\xd4\x60\xb5\xc7\x29\xd9\xa0\xdf\x8e\x4e\x25\x6e\xef\x71\x77\xff\x80\xf9\xed\xe2\x41\x8a\x5e\xe9\x17\xb5\x26\x24\x0f\x21\xcc\xb6\x77\x1c\x50\x89\xa2\x74\xbe\x14\x45\xb9\xda\x07\x4a\x1f\x31\x22\xd0\xb6\xef\x54\x20\x94\x23\xe5\x73\xa4\x28\xc8\x06\xaf\x37\xb4\x55\x88\x11\x3d\x1b\x1b\x5a\x94\xef\x7f\x95\x90\xdf\x0f\xde\xc3\x20\x6a\x21\x5e\x15\x63\x04\x3d\x66\x78\x7c\x22\x1b\xe4\xc2\x06\xe2\x56\x69\x8a\x29\xe2\x1c\xac\xec\x9a\x70\xf2\x3c\xc1\x89\x55\x5b\xc2\xa7\x19\xe4\x9d\xda\x92\x4f\x1e\xc5\xdf\x28\x99\xe0\xb7\x2c\x1f\x87\xf2\x30\x30\x0c\x93\xd1\x89\x6c\x93\x66\x06\x21\xda\x9d\xd5\x79\xbd\xaa\x46\x14\x45\xaa\xd1\x19\x4b\x1e\x8f\x4f\x8f\x4f\x69\x3f\x51\xb4\x8e\xf1\x3c\x39\xb4\x4b\xa1\x63\x8f\x63\xdb\x28\x8a\x62\x35\x01\x31\x27\xed\x9b\x62\xbf\x51\xdd\x32\x8b\xd5\xc8\xd4\xa2\x28\x4c\x9b\x89\x77\x33\x58\xd3\xe5\x99\xa2\x55\xa6\xab\x88\x39\xc9\xa9\xff\x98\x3b\x83\xea\x7b\xb2\x4d\x95\x7f\x27\x58\xd5\x22\xa9\xce\xcb\x65\x68\xdc\x2e\xc8\x1f\x6c\x02\x55\xf9\xea\xe5\x17\x67\xec\x11\x1c\xeb\x56\xe5\x4f\x5b\xd6\x75\xfd\xb6\xdb\x31\x25\xc5\x3b\xce\x4b\x8e\x5e\xc4\x3c\x7a\x2d\x03\x1b\xbb\x4e\x8c\x9c\x27\xa6\xaa\x3f\x64\x93\x0c\xce\x7f\x9b\x50\x5d\x66\xbb\xff\x5e\x79\xdc\x6c\x7c\xe4\x18\x8f\x17\xfa\x27\x00\x00\xff\xff\x54\xe7\x81\x8f\x3b\x03\x00\x00")

func templateMainTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateMainTmpl,
		"template/main.tmpl",
	)
}

func templateMainTmpl() (*asset, error) {
	bytes, err := templateMainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/main.tmpl", size: 827, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x5a\x5f\x6f\xdb\x46\x12\x7f\x96\x3e\xc5\xd4\x40\x0d\x32\x50\xe9\x5e\x51\x14\x77\xca\xe9\x80\x22\x4d\x50\x5f\x2f\x6e\xd0\x24\x7d\x31\x0c\x97\x26\x87\xd6\xc6\xe4\x52\x5d\xae\x1c\xbb\xae\xbf\xfb\x61\x67\x76\x97\x4b\x8a\x94\x14\xff\xc9\x4b\xc4\xd9\x99\xd9\x99\x1f\xe7\xdf\x2e\x7d\x74\x04\xaf\xea\xd5\xad\x12\x97\x4b\x0d\xdf\x7d\xfb\x8f\x7f\x7d\xb3\x52\xd8\xa0\xd4\xf0\x26\xcd\xf0\xa2\xae\xaf\xe0\x58\x66\x09\xfc\x58\x96\x40\x4c\x0d\x98\x75\x75\x8d\x79\x32\x3d\x3a\x82\x0f\x4b\xd1\x40\x53\xaf\x55\x86\x90\xd5\x39\x82\x68\xa0\x14\x19\xca\x06\x73\x58\xcb\x1c\x15\xe8\x25\xc2\x8f\xab\x34\x5b\x22\x7c\x97\x7c\xeb\x56\xa1\xa8\xd7\x32\x37\x2a\x84\x24\x96\xff\x1d\xbf\x7a\x7d\xf2\xfe\x35\x14\xa2\x44\x47\x53\x75\xad\x21\x17\x0a\x33\x5d\xab\x5b\xa8\x0b\xd0\xc1\x7e\x5a\x21\x26\xd3\xe9\x2a\xcd\xae\xd2\x4b\x84\xb2\x4e\xf3\xe9\x54\x54\xab\x5a\x69\x88\xa6\x93\x03\x94\x59\x9d\x0b\x79\x79\xf4\xa9\xa9\xe5\xc1\x74\x72\x50\x54\xda\xfc\xa7\xb0\x28\x31\xd3\x07\xd3\xe9\xe4\xe0\x52\xe8\xe5\xfa\x22\xc9\xea\xea\xa8\xb0\x0e\x1f\xa1\x24\xb6\x91\xa5\xa3\x26\x5b\x62\x95\xee\xe6\x38\xc2\xfc\x12\xf7\x60\x2b\x04\x96\xf9\x1e\x7c\x42\xe6\x78\x73\x30\x8d\xa7\x06\xb4\xf7\x44\x03\x85\xf6\x75\x35\x90\x4a\x40\xa9\x13\xbb\xa0\x97\xa9\x86\xcf\x69\x43\xa8\x60\x0e\x85\xaa\x2b\x48\x21\xab\xab\x55\x29\xcc\xab\x69\x50\x81\x45\x2e\x99\xea\xdb\x15\x3a\x95\x8d\x56\xeb\x4c\xc3\xdd\x74\x72\x92\x56\x08\xf6\x5f\xa3\x95\x90\x97\xd0\xff\xf7\x87\x81\x76\x7e\x20\xd3\x0a\x67\x75\x25\x34\x56\x2b\x7d\x7b\xf0\xc7\x74\xf2\xaa\x96\x85\xb0\xfc\xc6\xac\xf0\xb9\x2b\x9b\xd1\x4a\x57\xfa\x75\x7e\x89\x8d\x65\x3b\x3d\x7b\x61\x1e\x47\x76\x36\x18\x37\x5d\xe1\x37\x06\xcf\xc6\x0b\xd3\xe3\xb0\x30\x21\xdf\x93\x3e\x36\x28\xdb\xcd\x4f\xcf\x5e\xd0\xe3\xb0\xb4\x60\xce\xae\xf8\xcf\x75\x7d\x15\x58\xfe\xae\x6e\x84\x16\xb5\x1c\x10\x5f\x1a\xce\xae\xf0\xbb\xba\x14\xd9\xed\x3e\xc2\x2b\xe2\xec\x4a\xff\x28\x65\xad\x53\x23\xd0\x40\x95\xae\x4e\xf9\x95\x9d\x09\xa9\x51\x99\x78\xba\xbb\x77\xd2\x69\xcb\xd9\x51\x71\x4f\xa1\xe5\xb7\xcd\xb1\xc9\x94\xb8\xc0\x06\x52\x58\x39\xa2\xcd\x4c\x8e\x49\x1b\x39\x5e\xa2\x8d\x9d\x00\x37\x21\x35\xc0\xd1\x11\x30\xc9\xca\x13\xf4\x47\x06\x03\x28\x45\xa3\x93\xe9\xe4\xad\xb8\xc1\xfc\x98\x9c\xbd\xa8\xeb\xd2\x4a\x88\x2c\xd5\xd8\x80\x28\x82\x5d\xa1\xbe\xf8\x84\x19\x87\x77\x65\xa4\xbe\x11\x92\x15\x08\xe9\x36\xe1\x2d\x89\x04\x22\xdc\xb8\x22\x12\xef\xc9\xfe\x72\x80\x6c\x66\x12\xd3\x1f\x90\x48\x2c\x38\x9c\x47\xa3\x99\x34\x9e\x4a\xc7\xb2\xa8\x5b\xb6\x17\x84\x5c\xf2\xe1\x76\x85\x9d\x05\x2b\x6e\x0c\xe8\x8a\x7f\x48\xc3\xcd\x76\xec\xae\xd3\x5e\x26\xbe\x17\x7f\x05\xb6\xbf\x10\x52\xff\xf0\xfd\xa8\x74\x23\xfe\xea\x6d\xfe\x5a\xae\xab\xc6\xb3\x9d\x9e\x31\x28\x77\x70\x32\x83\xdf\x9d\x2d\x3e\x2c\xd1\x30\x77\xe5\x3f\x4a\xf1\xe7\xda\x1b\x40\x71\x31\xf0\xcf\xca\xaf\x89\xb9\xab\xe0\x44\x94\x65\x7a\x51\xe2\x5e\x0a\xa4\x65\xee\xaa\xf8\x75\x65\x62\x3b\x2d\xf7\x52\x51\x5b\xe6\xae\x8a\x9f\xb0\x48\xd7\xa5\xde\xcf\x8d\x9c\x99\x07\x35\xfc\x9e\x96\x06\x8e\x30\xa7\xc7\x35\x9c\x5f\x1b\xee\x41\x3d\xbf\x08\x69\x6a\xa2\xed\x84\x89\x7d\x1c\xd3\x73\x25\x64\xde\x7b\x2f\xab\x3c\xd5\xe8\xdc\xda\xf5\x5e\x88\xf9\x7c\xd0\xaf\xe3\xaa\x5a\x6b\xff\x82\x76\x28\x12\x8e\xb9\xab\xe3\xf7\xb4\x14\x79\xaa\x6b\x45\x91\x46\xb9\x3f\xae\xe3\xda\x33\xf7\x02\x5d\xd7\x2a\xbd\xc4\x5f\x90\xea\xef\x8e\x34\x69\x98\xf9\xfc\x0a\x6f\xfb\x15\x3c\x2c\xd9\x83\x15\x3c\x2c\xe2\xbc\xda\x33\x04\xa5\x21\x5f\xef\x85\x48\xe3\x98\x7b\x3a\xa8\x4e\x9a\x1a\x61\x78\x83\x66\xd0\xf1\xcb\xe9\x20\xe6\xf3\xcd\xca\x11\x36\x14\x18\x6b\x29\xbb\x7a\xca\xe4\x55\x5d\x55\xe8\xdf\xc9\x0e\x60\x33\x66\x1e\xe8\x4a\x34\x03\x6c\x16\x69\x22\x3f\xa0\x46\x93\xdc\xd3\x8c\x3a\x0e\xe6\xdd\xb2\xdb\x8b\xf3\x0e\xd9\x7e\x65\xfe\x0d\x0b\x6f\xf5\x76\x51\x85\xc5\xf9\xa6\xd9\xbf\x61\xe1\x19\x07\x27\xac\x50\x7e\xbc\x2a\x8f\x04\xe8\x96\x92\x7c\x2c\xaf\x51\x35\x5b\xc3\xdb\x4f\x58\xc4\xd9\xb7\xfb\xcf\xb5\x50\x98\xef\x16\x57\x96\x73\x3c\xd1\x5f\x98\xf1\x31\xe9\xa6\xfe\x1e\x59\xfe\x54\x93\x16\x0f\x2b\x9b\x41\xcd\xf4\x07\x44\x35\x0b\xb6\x61\x1d\xbc\x28\x0f\xd5\x96\x37\x13\x8c\xdd\xa7\xae\x54\xec\x35\x67\xf7\xb9\x87\x06\xeb\x00\x65\x1f\xae\x3b\x80\x66\x94\x4e\xf0\x33\x85\x67\xa6\x90\x66\xc1\x54\x3a\x44\x8c\x51\x0c\x0b\xfd\xe2\x79\x75\xa5\x6b\x95\x4c\x8b\xb5\xcc\x9c\x64\x84\xb9\x7d\xd3\x3f\x79\x8e\xd8\xc6\xfc\xdd\x74\x22\x11\xe6\x0b\x38\x34\x8f\x77\xd3\x89\x49\xc9\xb9\x8f\x24\xcc\x93\x0f\xe9\xe5\xcc\x90\x6f\x57\x38\x0f\xc9\x26\x97\xa7\x13\xaa\x1c\x21\xdd\x3c\x1b\x3a\x43\x3f\xf7\x74\x7e\x36\x2b\x36\xfe\xe7\x6e\xc5\x3e\x9b\x25\x17\xdb\x73\xbb\xe4\x9e\x79\xad\x68\xf7\xa2\xb5\xc2\xed\xd5\x42\x3b\xa7\xa5\xf6\xd9\xac\x06\xd1\x3a\x87\x2a\xbd\xc2\x68\x38\x66\xe3\xd9\x74\x72\x3f\x9d\x14\xb5\x82\xf3\x19\xa4\xda\xa0\xa2\x52\x79\x89\x46\x65\x18\xf2\x06\x25\x89\x49\x9a\xe7\x2d\x35\x4a\x75\x4c\xe2\xa2\x30\x53\x85\x91\x65\x1b\x5f\xd2\xe3\x57\x0b\x90\xa2\x74\x92\xa6\xf4\x2c\xfc\xdb\x51\x58\xc4\x4c\x0f\x42\x64\x01\xcc\x17\xd0\x48\xbd\x42\xbd\x56\x12\x24\xb6\xc1\xc1\xe3\xb6\x8f\x0e\x9f\x2e\x44\xa6\xe8\xe0\x9f\x43\xe1\x41\xb2\x51\x91\xbb\xb1\x3a\x0c\x90\x88\x8f\x8e\x33\x40\xa5\xcc\xf3\x1d\x39\x57\xe4\xc9\x6b\xa5\x42\x87\x9c\x49\xa2\x9c\x41\x51\x69\xb3\x5c\xab\x22\xe2\x24\x80\xaf\xff\x9c\xc3\xd7\xd7\x07\x33\x23\x48\xef\xcb\x6a\x60\xb4\x1a\x42\xea\x90\x36\xba\xeb\x47\x13\x78\x19\x8a\x9a\xa2\xee\xae\x18\xca\xac\x1f\xb0\xb4\x62\x43\x96\x86\xef\x79\xb8\x40\x94\x8d\xe8\xa4\xa5\x36\x3e\xdd\xc8\x3c\x6f\x6d\x70\x73\xf1\x74\xe2\xa7\xe1\x76\xd5\x51\xcc\xaa\x9d\x08\xe7\xad\x5e\x37\x23\x32\x60\xb4\x77\x38\x3b\xce\x69\xef\xce\x34\xd9\x72\xfa\xe1\x70\xee\x7d\xf6\x13\x60\x3f\xec\x69\xb9\x1b\xf8\xed\x5c\x48\xeb\x25\xca\xa8\xc8\x93\x96\x1a\x93\x12\x37\x41\xf9\x3d\x3c\x85\x96\xfd\x24\xe5\xf7\xf0\x94\x8d\xe4\x82\x5d\xe9\xe5\x86\xa1\x00\x1f\x4b\x19\xcd\xbd\x62\x33\xf7\x9a\x62\x3c\xf7\x9a\x82\xe2\x02\x16\xbb\xe3\xb3\x12\x4d\x63\xca\x30\x75\x0e\x61\x84\xcc\xf6\x2e\x6a\x0f\x66\x46\x97\x89\xbe\x56\xb7\x39\x0b\xce\x17\x40\x87\x40\x03\xa5\x39\x1c\xc6\x2f\x99\xfe\xd5\x02\xbe\x75\xd6\xd1\xa1\x71\x01\x87\x66\x21\x30\xcc\xbd\x60\xcb\x15\x1e\x45\x16\xfe\x28\x62\x80\xfd\xb5\x88\xda\xc8\x89\xe9\x74\x12\xb1\x15\xa6\x69\xf2\x55\x80\x3d\x4d\x00\x9d\x71\x20\x4b\x25\x5c\x20\xd0\x8d\x1f\xe6\xa0\x6b\xe2\xb9\x44\x89\x2a\xa5\x84\x37\x92\x6f\x6a\x05\x78\x93\x56\xab\x12\x67\x20\x6b\x0d\x29\x98\x3a\x40\x03\x7a\x29\xae\x10\xb4\xa8\x30\x39\xa9\x3f\x27\x64\xf1\x39\x65\xbe\x71\xd8\x74\xa9\xe4\x6d\xaa\x9a\x65\x5a\x86\x96\xbd\x24\x86\x00\xea\xd6\x2b\x3e\xa8\x2d\x82\x0c\x08\xcb\x57\x53\xcc\x8c\x4c\x5b\xc3\xb8\x71\x6f\x76\x38\xbe\xba\xa0\x22\xc6\x3f\x87\x8a\x18\x09\x47\x22\xbf\x31\xe7\xf3\x1c\x6f\xba\x6d\x8e\x55\xdf\xf9\xbd\x0f\x89\x60\xac\xa5\x76\x6f\x43\x51\xe4\x37\x34\x4b\x53\x75\xe0\xce\x3e\xf7\x0b\xfc\xdc\xaf\x1b\x66\xa5\xad\x1a\x61\x32\x9a\x95\x4e\x2a\xde\x5b\x4f\x2d\x86\xf6\xa6\x91\xdf\x16\xbd\xa9\xe0\xe6\xd2\xa7\x8c\xf9\x55\x43\x0a\xff\x7d\xff\xeb\x89\x11\xa6\x79\xc8\xbe\xe8\x1c\xf9\x45\x13\x8b\x51\xf0\xbe\x73\x33\xc4\xff\x59\x84\x3a\x9b\x46\x8d\xdb\xdb\x8c\x59\x76\xa7\x18\xa2\x0b\x38\x3d\xbb\xb8\xd5\xc8\xef\xbc\xad\xf8\x0d\xd5\x67\x96\xbd\xa3\x04\x96\x85\x70\xf5\xd6\x5e\x82\x31\x2d\x8a\x37\xc6\x01\x21\xf9\xee\x3a\xea\x05\x37\xcb\xc5\x31\xa5\x17\xcb\x7d\x61\x93\x16\x85\x8b\xcd\x26\x31\xa1\x42\x17\x5d\x4e\x2f\x87\xe5\x1e\x1d\xca\x62\xe1\x5b\x54\x63\x3b\x14\xba\xf6\xd4\x5e\x34\x07\x83\x2c\xd4\xd7\xa8\x94\xc8\xd1\x5f\xbe\x85\xab\xc9\x60\x15\xb3\x48\x05\x5e\x46\x31\x67\xcc\x78\x29\xeb\x38\xc8\x21\xf8\xf4\x1e\xf2\x5c\xeb\xf7\x4a\x0b\xa4\x2c\x70\x1b\x79\x43\x9e\x62\x2f\x8b\x0b\x86\x83\x95\x19\xb8\x19\x07\x1e\xbe\x17\x90\xae\x56\x28\xf3\xc8\x12\x66\xed\x10\x1b\xa4\x75\x14\xc7\x16\x26\x7b\x07\x1d\x3a\x60\x6f\xb0\x9f\xd3\x05\x53\x6b\xbc\x13\xd6\x06\xeb\x86\xbb\x3f\x0f\x1c\x39\x76\x46\x86\xb5\x6a\xd0\x9b\xde\x4b\xa7\xcb\xf4\xa7\x7f\xe7\xfd\x6d\xf8\xda\xfd\xe9\xf7\xb1\x82\x9d\xee\xd1\xc4\xb6\x14\x7e\x94\x55\xa7\x18\x72\x45\x6b\xb8\x6f\x89\x6b\x94\x70\xb1\x2e\x0a\x54\x40\x35\xd0\xb6\x03\x77\xeb\x4e\x75\xad\xa7\x21\xba\x58\x17\xb6\x88\x99\xd9\x95\x89\xb3\xb1\x52\xd6\x81\x81\x2c\xf4\xea\x8c\xa2\x19\x34\xdb\x81\x40\xa5\xc2\x80\x28\x82\x54\xb7\xed\x82\x44\xda\x3d\x8a\xc4\x76\xec\x26\xda\xd4\xbc\xa9\xda\xe8\x0e\xfa\x65\xd8\x2e\x7d\xbd\xa3\x5f\x8d\xbd\xd1\xd7\xb5\xfb\x3a\xc0\x07\xc2\xb0\xbe\x5b\xc0\xa2\x06\x2c\x2c\x31\xf4\x8b\x66\xbf\x21\x10\x6c\xc6\x36\xd2\xde\xc9\xaf\x4e\xad\xdd\x92\x5d\x21\x44\x62\x06\x55\x90\x32\x6c\x32\x9d\x86\xd2\xca\xce\x54\xc3\xad\xa2\xba\xf1\x6d\x62\x3a\x99\xd8\x73\x75\x68\x8d\x2d\x8c\xd5\x4d\xdc\xc2\x3d\x80\x6c\x77\xf0\x33\xbb\xfb\xb8\x95\x41\xd4\x1a\x7b\xc9\xe0\x4f\x9d\x77\x5a\xb4\x6f\x74\x62\x66\x17\xbb\x7f\x7b\x80\xea\x66\xb3\x61\x1b\x30\xe5\x4b\x6d\x21\x63\xcc\x4c\xe5\xaf\x51\x17\x70\xe8\x7e\xb3\x46\x2a\x27\xb6\xdf\x7e\x9a\x11\xc9\x7e\x47\x22\xa2\x56\x3c\x9c\x4c\x82\x8f\x43\x73\x10\xb3\x56\xb9\x0b\xd6\xa0\x5c\xd9\x69\x07\x9a\xc2\x01\x32\xd6\x24\x9e\x1a\xf4\xb1\xe6\xf0\xa0\xee\x40\x5a\xb7\xf5\x87\x67\xb0\x7e\xb4\x2f\x3c\xa6\x31\xd0\x06\xfc\xb5\x34\x74\x83\x9b\xc3\x93\xc7\x7d\x6b\x3f\x6d\xe9\xac\xe7\xef\xba\x81\xed\x3f\xb3\x41\x4f\x18\x8f\xce\x0c\xfb\x6d\x37\xf4\xd5\x76\xa8\xa7\x74\x56\x14\xc0\x1b\x75\x14\x35\x89\xfd\x06\x1d\x78\xfa\xce\xda\xd3\x73\xf5\x8b\xfd\x1a\x18\x0b\xab\x9b\x81\x91\x70\x78\x26\xec\x36\x84\x6e\x37\xb0\x39\xcc\xed\x80\x0f\xb0\x0f\x68\x07\x9d\x11\x73\xb4\x1f\x8c\x97\xe0\x2f\xee\x08\xc3\x05\x76\xbf\xfa\x3a\x1e\x04\xbe\x7d\x8e\x56\x4e\xf7\x7a\x88\x67\x57\x01\xdc\xc0\x7c\x10\xbb\x70\x52\x1b\x85\x6e\x2c\x87\xbf\x10\xb8\xa1\x0c\xdd\x37\x41\x7d\x7e\x72\x6c\xfa\x18\x2e\xd2\x92\x6f\x60\xef\xf7\x76\xb9\x33\x35\x8e\xfa\x3c\x9e\xcc\xfb\x7b\x3d\x98\xaa\xfb\x65\xea\x7e\xee\xf4\xd2\x4d\x6e\x1e\xd7\x28\x33\xb3\xb5\x52\x33\xa8\xaf\x78\x72\x0e\x12\xf7\x34\x95\x76\x46\x39\x23\x6b\xbf\xaa\xaf\xac\x8d\xc3\x4c\xc6\x66\xe9\xfd\x74\x3e\x56\x4e\xb7\xd9\x27\xb1\xf8\x24\x6f\x51\x5d\xa2\x8a\x5f\xc2\x6e\x9d\x15\x33\x47\xa9\x8c\xed\x95\x03\x7b\x8a\x7c\xd7\xff\x20\x3f\x71\x1f\x3f\xc7\x98\x1e\xe3\xe7\x16\x9d\x63\x7e\x16\xc0\x37\xd6\x0f\x72\xb4\xd8\xc7\xd1\x31\xa6\xc7\x38\xba\x45\xe7\x6e\x47\xdb\x33\x45\x9b\x72\xc6\xde\xf6\xb2\xf1\xef\xbf\xcd\xd3\xb1\x2c\xea\xe4\x64\x5d\xa1\x12\x59\x14\x1b\x62\xef\xfe\xb1\xbd\x80\x7c\x63\xb6\xe8\x1e\x77\xc8\x21\xd9\x22\x15\xde\xf1\x25\x51\x51\xd6\xa9\xfe\xe1\xfb\xb8\x83\xd4\x40\x47\x5e\x4b\xbc\x59\x61\xa6\x31\xef\x5d\x5e\xd2\x05\xac\xbf\x7b\x9d\xf3\xe5\x6b\x78\xf7\xda\x7c\x16\x3a\x5b\x82\xe6\xdd\xc9\x17\x73\x3c\x78\x49\xef\x30\x6d\x10\x34\xfc\x67\x01\xe1\x1f\x06\xe9\x7f\xc2\xe1\x21\x68\xf8\x77\x8f\xfc\xc3\xf7\x73\x82\xbc\x77\x4b\xc9\x37\xba\x06\xe5\x21\x75\x1f\xc5\xb0\xbe\x8f\x62\x54\xe1\xba\xd5\x38\xd4\xb4\xdb\xae\x09\x9f\x55\xba\x6a\xc2\x3f\x29\xb3\xf4\x54\xe6\x7c\x4c\x72\x84\x0a\xf5\xb2\xce\xe1\xb3\xd0\x4b\x50\x98\xd5\xd7\x7c\x36\x46\xd9\xac\x15\x82\xac\x61\x95\x4a\x91\x35\x20\x24\xd8\x83\xac\x90\x97\xb6\xd5\x07\x5d\xba\xc8\x83\xbf\x9e\x01\x4b\x8c\xe1\xf4\xac\xfd\x93\xaf\xfb\x18\x22\xdb\x90\x03\x72\xff\x66\x30\x47\x73\x3a\x37\xea\xed\xdc\x22\x0a\xb8\xa6\xde\xc4\xc6\x99\x63\xee\x75\xa7\x41\xd3\x65\x71\x27\x24\xbe\xfe\xe0\xbc\x63\xe3\xfd\xa7\xa2\x19\x5c\xd3\x09\xa8\x70\xcd\x99\xa2\x90\x66\x20\x73\x10\x74\xd1\x95\x27\xce\x81\x59\x0f\x5d\x3e\x2f\x6c\x80\xcb\xe4\xc7\x42\x19\x5e\x91\x85\x68\x32\xdd\x81\x49\xdf\x57\x0d\x96\x7c\x90\x69\x89\xcf\x81\x64\xc7\xbf\x0e\x98\x0c\x24\xda\xf3\xd3\x20\x8e\xa1\xf0\x26\x94\xee\xe0\xb2\x01\xa6\x5b\x78\x2c\x9c\xdd\x0b\xbb\x10\x50\xb7\xe2\x20\xe5\xbb\x7c\x83\xa9\x3b\x5c\x05\xf4\x67\x84\xd5\x79\x3a\x00\xac\xf0\xc7\xba\x6d\xd0\x7a\x47\xfa\xe0\xf2\x45\xce\x06\xb4\x4c\x7e\x2c\xb0\xdb\x2e\x78\x22\x3e\x20\x31\x7e\x6f\xdb\x4b\x9e\x67\xc1\x8f\xdd\x19\x40\x8f\x8d\xd8\x8e\x1d\x7b\xb1\x81\x1c\x0f\xbc\x1b\xc8\x31\xf9\xb1\xc8\x75\xe6\xf9\x20\x20\x99\xee\xc2\xd1\x3c\x51\x34\xf2\x20\xde\x12\x9f\x11\x4a\xf6\x6f\x00\xca\xa5\x3d\x00\x6c\x83\xd2\x9a\xdf\x87\xd2\x4e\xd2\x1b\x58\x5a\xfa\x63\xc1\xec\x9e\x14\x02\x34\xed\x42\x4c\xb1\x69\x37\x33\x70\xda\x69\xbf\xa5\x3e\x23\x9e\x76\xdb\x01\x40\x57\xee\x7c\xb1\x0d\x51\xe7\xc2\xac\x73\xb8\xf0\xb7\x99\xba\xf3\x5d\x37\xee\x3c\xd1\x69\xba\x56\xa0\xed\x07\xde\x70\x08\x7b\xa7\x69\x94\x9b\x68\x58\x80\x4e\x5e\x97\x58\x45\x9d\x51\x42\x4f\xef\xa7\xff\x0f\x00\x00\xff\xff\x0a\xd7\xd3\x32\xbb\x32\x00\x00")

func schemaGoBytes() ([]byte, error) {
	return bindataRead(
		_schemaGo,
		"schema.go",
	)
}

func schemaGo() (*asset, error) {
	bytes, err := schemaGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.go", size: 12987, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
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
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
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
	"template/main.tmpl": templateMainTmpl,
	"schema.go":          schemaGo,
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
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
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
	"schema.go": &bintree{schemaGo, map[string]*bintree{}},
	"template": &bintree{nil, map[string]*bintree{
		"main.tmpl": &bintree{templateMainTmpl, map[string]*bintree{}},
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
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
