// Package swagger Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// v1alpha/tracker/tracker.swagger.json
// v1alpha/extractor/extractor.swagger.json
// v1alpha/schema/schema.swagger.json
// v1alpha/deps/deps.swagger.json
// v1alpha/store/store.swagger.json
package swagger

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"net/http"
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


type assetFile struct {
	*bytes.Reader
	name            string
	childInfos      []os.FileInfo
	childInfoOffset int
}

type assetOperator struct{}

// Open implement http.FileSystem interface
func (f *assetOperator) Open(name string) (http.File, error) {
	var err error
	if len(name) > 0 && name[0] == '/' {
		name = name[1:]
	}
	content, err := Asset(name)
	if err == nil {
		return &assetFile{name: name, Reader: bytes.NewReader(content)}, nil
	}
	children, err := AssetDir(name)
	if err == nil {
		childInfos := make([]os.FileInfo, 0, len(children))
		for _, child := range children {
			childPath := filepath.Join(name, child)
			info, errInfo := AssetInfo(filepath.Join(name, child))
			if errInfo == nil {
				childInfos = append(childInfos, info)
			} else {
				childInfos = append(childInfos, newDirFileInfo(childPath))
			}
		}
		return &assetFile{name: name, childInfos: childInfos}, nil
	} else {
		// If the error is not found, return an error that will
		// result in a 404 error. Otherwise the server returns
		// a 500 error for files not found.
		if strings.Contains(err.Error(), "not found") {
			return nil, os.ErrNotExist
		}
		return nil, err
	}
}

// Close no need do anything
func (f *assetFile) Close() error {
	return nil
}

// Readdir read dir's children file info
func (f *assetFile) Readdir(count int) ([]os.FileInfo, error) {
	if len(f.childInfos) == 0 {
		return nil, os.ErrNotExist
	}
	if count <= 0 {
		return f.childInfos, nil
	}
	if f.childInfoOffset+count > len(f.childInfos) {
		count = len(f.childInfos) - f.childInfoOffset
	}
	offset := f.childInfoOffset
	f.childInfoOffset += count
	return f.childInfos[offset : offset+count], nil
}

// Stat read file info from asset item
func (f *assetFile) Stat() (os.FileInfo, error) {
	if len(f.childInfos) != 0 {
		return newDirFileInfo(f.name), nil
	}
	return AssetInfo(f.name)
}

// newDirFileInfo return default dir file info
func newDirFileInfo(name string) os.FileInfo {
	return &bindataFileInfo{
		name:    name,
		size:    0,
		mode:    os.FileMode(2147484068), // equal os.FileMode(0644)|os.ModeDir
		modTime: time.Time{}}
}

// AssetFile return a http.FileSystem instance that data backend by asset
func AssetFile() http.FileSystem {
	return &assetOperator{}
}

var _v1alphaTrackerTrackerSwaggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x4d\x6f\xb4\x36\x10\xbe\xef\xaf\xb0\xdc\x1e\xa3\x90\xa6\xb7\xdc\xaa\x56\x95\x2a\xf5\xbd\x34\xb9\x55\x51\xe5\xc0\x2c\xf1\x5b\xb0\x1d\x7b\x48\xb5\x7d\xb5\xff\xbd\x02\x43\xf8\x5a\xd8\x05\x03\x09\x29\x2b\x45\xd9\xac\x3d\x93\xf9\x78\x1e\xcf\x78\x96\x6f\x3b\x42\xa8\xf9\x87\x85\x21\x68\x7a\x47\xe8\xed\xf5\x0d\xbd\x4a\x3f\xe3\x62\x2f\xe9\x1d\x49\xd7\x09\xa1\xc8\x31\x82\x74\xfd\xf5\x07\x16\xa9\x67\xe6\xa1\x66\xfe\xdf\xa0\x8b\xdf\xd7\x4a\x4b\x94\x99\x24\x21\xf4\x15\xb4\xe1\x52\x64\xfb\xed\x5b\x22\x24\x12\x03\x48\x77\x84\x1c\x33\xfd\xbe\x14\x26\x89\xc1\xd0\x3b\xf2\xa7\x95\x62\x4a\x45\xdc\x67\xc8\xa5\xf0\xbe\x1a\x29\xd2\xbd\x8f\xd9\x5e\xa5\x65\x90\xf8\x17\xee\x65\xf8\x6c\x4a\xc3\xbd\xc2\xe0\x50\x33\xf5\xec\x7d\x8b\x98\x08\x13\x16\xc2\xd1\x0b\x40\x81\x08\x40\xf8\x1c\xca\xfd\x84\xd0\x10\xb0\xf2\x27\x21\x54\x2a\xd0\xd9\x7f\xfa\x2d\x48\x3d\xfa\x9d\x1b\xfc\xa5\x2a\x7b\x55\x6e\xd5\x60\x94\x14\xa6\xa6\x30\x5b\xb8\xbd\xb9\x69\x7c\x44\x08\x0d\xc0\xf8\x9a\x2b\xcc\x43\xf5\x13\x31\x89\xef\x83\x31\xfb\x24\x22\x85\xa6\xeb\x8a\xfa\x4c\xc8\xf8\xcf\x10\xb3\x96\x32\x42\xe8\xf7\x1a\xf6\xa9\x9e\xef\xbc\x00\xf6\x5c\xf0\x54\xaf\x29\x12\xd4\xb4\xfa\x8f\x5c\x3f\xad\x69\x39\xee\x4e\xbd\x3f\x56\x3c\x54\x4c\xb3\x18\x10\x74\x99\x0c\xfb\x6a\xf8\x26\x58\x9c\xe1\xa5\x88\x77\xd3\x0d\x9e\xb9\x9c\x26\xab\xb9\xa2\xe1\x25\xe1\x1a\xd2\x58\xa3\x4e\xa0\xb1\x8a\x07\x95\xe9\x35\xa8\xb9\x08\xab\xd6\x1f\xaf\xce\x5b\x23\x75\xc8\x04\xff\x37\xcb\xe6\x69\x8b\x5e\x12\xd0\x87\x1e\x93\xf6\x2c\x32\xd3\xda\x14\xcb\x20\x89\x3a\xe2\x33\xb1\x35\x6f\xef\x1f\x2b\x19\x45\x16\x36\x73\x49\xdf\x90\x72\xb8\x07\xfd\xca\xfd\x0a\x4c\x1e\x77\x55\x65\xb9\x83\x17\xd2\xcc\x43\xa9\x64\x24\xc3\x83\x0b\xdf\x1e\x0a\x1d\x1b\xef\xc8\xc6\xbb\xcf\xc5\xbb\x02\xdb\xf3\xb0\xce\x43\x0e\xd6\x7e\x67\xf2\x3d\x58\x4d\xab\xa4\xa0\xb5\x7d\x23\xe2\x46\xc4\x7c\xff\x12\x44\xc4\xd1\x3d\x26\xae\xb0\xc3\xc4\xad\xce\x6d\xf4\x2a\xf6\x2f\xd3\x5f\xa2\x7b\x77\x89\x2b\xee\x2d\x37\xc6\x6d\x8c\x7b\xdb\xbf\x4c\x41\x9b\xaa\xaf\xc4\xd5\x77\x95\xb8\xf5\x94\x1b\x05\x67\xa3\xa0\xf5\x65\x70\xff\xb8\x1e\x1a\x7d\xc9\x1c\x5c\x8a\x3c\xaa\x93\x38\x2e\xc0\xe0\x02\x21\x04\xdd\x94\xde\x4b\x1d\x33\xcc\x37\xfc\x78\x3b\x14\xc6\xbe\x4c\x04\x7e\x08\x63\x2f\x46\xb9\x4d\xe6\x48\x8c\x7b\x31\x13\x2c\x1c\x5e\x4a\xbe\xe4\x62\xeb\x81\xbc\x35\x78\x29\xcc\x27\x3a\xfa\x58\x67\xa1\x1b\x4a\x8c\x4c\xb4\x0f\x43\x41\x72\x9f\x49\xad\xe8\x36\x9d\x1b\xfc\x31\xba\x8a\xe5\x6b\xf8\xd6\x57\xd8\xd7\x3c\x5c\x32\x39\x1b\x3e\x6d\x57\x61\xd9\xb3\x75\x15\x6d\x5b\x57\xd8\x55\xd8\x64\x8e\xc4\xb8\x45\x45\x15\xe9\x4a\x9a\x7e\xa8\x3f\x64\x12\xab\xc0\x7a\x66\xea\x52\x30\x7f\x92\x41\x0b\x1e\x16\x39\xa7\x56\xfa\x6f\x9d\x23\xfd\x2d\x78\xfd\x92\x80\xc1\x4b\xfc\x9d\x06\x5b\x6f\x0f\xae\x54\x4c\x2a\x1f\x33\x09\x40\x99\x72\xa8\x69\xbb\xbb\x18\x04\xfe\xca\xa3\x5a\x9f\x52\x70\x45\x3e\x7d\x05\xbf\xe4\x20\x55\x3a\x05\x20\xf2\x06\xbe\xca\x9a\xdc\x40\x5d\x57\xdd\xa8\xe6\xd6\x1c\x0c\x42\x3c\x46\xb2\x56\x79\x47\xc8\xe7\x55\x72\x84\x64\xf9\xe4\xd0\x60\xd1\x8e\xe7\x79\x6a\xf2\x4c\x6b\x56\x47\x29\xe5\x08\x71\x73\x7f\x27\x04\xf3\x83\xa5\x9e\xed\xd3\x27\xda\xf1\xe4\xd1\x64\x31\x6f\x65\xcd\x7b\x01\x23\x0f\xf2\x5f\xbe\x14\x06\x35\xe3\x02\x47\xc1\xcb\x97\x6a\x8a\x48\x5f\xd2\x02\xf5\x05\xd3\x92\xed\xdd\x82\x39\x9e\x65\x43\xb1\xde\x17\x83\x26\xe1\x16\x0d\xc1\xc2\xc7\x45\x4f\x1c\xee\x5b\x17\xc3\xa1\x71\x48\x6f\xca\x4e\x86\x9c\x7b\xee\xc7\xc1\xb6\x05\x4f\xb8\xdc\x8b\x71\x87\xdc\x65\xcf\x5d\x38\x04\x02\xb9\x6d\x58\x66\x89\x40\xd1\x56\x55\xbe\x93\x70\x74\x1e\x27\xcd\x3e\xae\x2d\xf7\xad\x6f\x47\x3e\x7f\xe6\x9b\xe3\x3d\x07\x8f\xdb\xd3\xff\x39\x7c\xce\x0d\xce\xeb\xc8\x48\xa7\xeb\x63\x7c\x07\x9f\x55\x4f\x21\x3a\x75\xcb\xec\xbe\x63\x56\xcb\x8c\xbd\xf6\x4e\xae\x76\xe6\x04\xd5\x0a\xfc\xb8\xc4\x34\x26\x21\xff\x97\xc4\xb4\x27\x5c\x33\x24\x26\xef\x38\x5c\x12\x33\x49\x75\x98\xd9\xd9\xfa\x31\x31\xda\xe9\xfa\x31\xe3\x72\x2a\xb6\x9a\xfe\xb3\x99\x2a\xee\x09\x03\x5a\xd0\x21\x8c\xbc\xc4\x6f\xf7\xf6\xb4\xf5\xcd\xc7\x59\x3b\x9b\xb9\xaa\x79\x3d\x45\x14\xfb\xfc\xae\xcf\x6a\x3e\x98\xdf\xc5\x7c\x66\x36\xd2\xf4\x0e\x85\x86\x93\xa7\xd6\x97\x38\xf6\x50\x2b\x68\x21\xeb\x73\x4d\x17\x7f\x53\x45\xe9\xcd\xad\xc3\xe7\x27\x29\x23\x60\xa2\xab\xc2\x14\xcb\x67\xac\x3e\x3d\x1a\x72\x30\xfb\xdd\x87\x70\x3f\xaf\x7c\x3c\xd4\x89\x45\xe7\x4b\xd8\xa0\x13\xb3\x18\xf6\x2d\x53\x77\x76\xe9\xcf\x71\xf7\x5f\x00\x00\x00\xff\xff\x4b\x1f\x46\x5f\xd7\x39\x00\x00")

func v1alphaTrackerTrackerSwaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_v1alphaTrackerTrackerSwaggerJson,
		"v1alpha/tracker/tracker.swagger.json",
	)
}

func v1alphaTrackerTrackerSwaggerJson() (*asset, error) {
	bytes, err := v1alphaTrackerTrackerSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v1alpha/tracker/tracker.swagger.json", size: 14807, mode: os.FileMode(420), modTime: time.Unix(1586631428, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v1alphaExtractorExtractorSwaggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x56\x4d\x6f\xdb\x30\x0c\xbd\xfb\x57\x08\xda\x8e\x45\x92\xf5\x98\xdb\xd0\x75\xc0\x0e\x05\x8a\x5d\x87\x1c\x58\x89\x71\x54\xd8\x92\x2a\xd2\xdd\xb2\x21\xff\x7d\x90\x62\xd7\x1f\x8d\x8b\xc4\x19\x16\x0c\xcb\xc9\xb1\xf8\x28\xf2\xe9\xf9\x51\xbf\x32\x21\x24\x7d\x87\x3c\xc7\x20\x97\x42\x5e\xcf\x16\xf2\x2a\xbe\x33\x76\xed\xe4\x52\xc4\x75\x21\x24\x1b\x2e\x30\xae\x3f\x7f\x80\xc2\x6f\x60\x8e\x3f\x38\x80\x62\x17\xda\xa7\x99\x0f\x8e\x5d\x42\x0b\x21\x9f\x31\x90\x71\x36\x61\xf6\x8f\xc2\x3a\x16\x84\x2c\x33\x21\x76\x69\x0f\xe5\x2c\x55\x25\x92\x5c\x8a\x6f\x7b\x14\x78\x5f\x18\x05\x6c\x9c\x9d\x3f\x92\xb3\x31\x76\x95\x62\x7d\x70\xba\x52\x47\xc6\x02\x6f\xa8\x2d\x7e\xde\x14\xad\xd1\xa3\xd5\x68\x95\x41\x6a\xea\x7e\x09\x8b\x38\x47\xdd\xff\x42\x48\xe7\x31\xa4\x1d\xbe\xe8\xd8\xc9\x6d\x8d\xb9\x6a\x23\x02\x92\x77\x96\x90\x7a\x40\x21\xe4\xf5\x62\x31\x78\x25\x84\xd4\x48\x2a\x18\xcf\x35\x33\x1f\x05\x55\x4a\x21\xd1\xba\x2a\x44\x93\x69\xd6\x49\x9f\x40\xa4\x36\x58\xc2\xab\x64\x42\xc8\xf7\x01\xd7\x31\xcf\xbb\xb9\xc6\xb5\xb1\x26\xe6\xa5\xf6\x44\xea\x72\xbf\xd6\x89\x65\x0f\xbe\xcb\x0e\x3d\xef\x3a\xad\x79\x08\x50\x22\x63\x68\x49\xdf\xff\x06\x4d\x59\x28\x93\x36\x1e\x9c\xde\x0e\x6b\x37\x76\x6c\x25\xe0\x53\x65\x02\x46\x5e\x39\x54\xf8\xc7\x7b\x7e\xaa\x90\xf8\x98\x96\x57\x9d\x96\x19\xf2\x61\xb3\xf2\x53\xa3\x9a\xed\x6d\xb3\x49\x9b\x76\x95\x75\xd3\xd5\xec\x8d\x28\xae\x04\x56\x9b\x93\xf4\x76\x97\x10\xff\x88\xda\x52\xb1\xff\x95\xd6\xea\x8e\x2f\xa2\xb4\x17\x13\xed\x54\xd6\x5a\x9e\x46\x4f\x6d\xb6\xae\xe6\x78\xeb\x13\x81\xee\xe1\x11\x3b\x4e\x16\xfd\xd5\x63\x60\x33\x90\x96\x74\x21\x07\x6b\x7e\x42\x2d\xa2\x9e\xe8\x9a\x5c\xc4\xc1\xd8\x5c\x1e\x3c\xd8\xd2\xe9\x2a\x8d\x8e\x93\x91\xf5\xd8\xb8\x71\x96\x38\x80\xb1\x3c\x25\x09\x29\xe7\x5f\x7f\x2d\x0d\x12\x42\x80\xbe\x58\xa4\x61\x2c\x87\xf1\x6f\xec\xd5\xd3\xf3\x41\x2f\xe8\x1f\xc5\x1d\x58\xc8\xb1\x44\xcb\x9f\x4d\x8f\x96\x93\x0f\xa6\x00\x9b\x57\x90\x4f\xa2\x96\xb6\xc4\x58\x4e\x41\x5e\x5c\x0e\x53\xa0\x5d\x17\x3e\x5f\x0a\x87\x4d\x61\xf0\xc5\x9d\x20\x90\xb1\xd1\x75\x86\x36\x08\xa3\x9f\x46\xfb\x98\xc0\xd6\xda\x14\x78\xe3\x2c\xa3\xe5\x51\xb6\x06\x85\xa4\x35\xd0\x3a\x71\x01\xc5\xfd\xe1\xb2\xde\xac\x60\x6c\x30\x54\xa1\x38\xba\x89\x23\xe9\xad\x27\xd4\x19\xfc\x96\xbd\xaf\xf8\x2f\x69\x6a\x60\x1d\x53\x14\xd6\x1b\x58\x97\xd2\x57\xff\x76\x3e\x99\xb2\xe9\x8e\x3c\x72\x63\x39\x4b\x0f\xac\x36\xa8\xef\x2f\xd6\x59\xbc\x0d\x64\xbb\xec\x77\x00\x00\x00\xff\xff\x94\xe2\x84\xeb\xce\x0d\x00\x00")

func v1alphaExtractorExtractorSwaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_v1alphaExtractorExtractorSwaggerJson,
		"v1alpha/extractor/extractor.swagger.json",
	)
}

func v1alphaExtractorExtractorSwaggerJson() (*asset, error) {
	bytes, err := v1alphaExtractorExtractorSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v1alpha/extractor/extractor.swagger.json", size: 3534, mode: os.FileMode(420), modTime: time.Unix(1586631428, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v1alphaSchemaSchemaSwaggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xcd\x31\xce\x83\x30\x0c\x05\xe0\x3d\xa7\xb0\x3c\xff\x82\xbf\x1d\xb9\x4a\xd5\x21\x0a\x01\x5c\x81\x1d\xc5\x86\x0e\x88\xbb\x57\x49\xd4\xbd\x53\xa2\xf7\xbe\x97\x9c\x0e\x00\xf5\xed\xe7\x39\x66\x1c\x00\xef\xdd\x3f\xfe\x95\x8c\x78\x12\x1c\xa0\xf4\x00\x68\x64\x6b\x2c\xfd\x71\xf3\x6b\x5a\x7c\xaf\x61\x89\xdb\xf7\xe8\x52\x16\x93\xba\x03\xc0\x23\x66\x25\xe1\xaa\xdb\x15\x58\x0c\x34\x1a\x3a\x80\xab\xbe\x1e\x84\x75\xdf\xa2\xe2\x00\x8f\xb6\xf2\x29\xad\x14\xbc\x91\x70\xff\x52\xe1\x62\x9f\xd5\xa6\x2c\xe3\x1e\x7e\xb4\xde\x96\x02\xcf\xf6\xcd\x18\x27\x62\x2a\xae\x85\xee\x72\x9f\x00\x00\x00\xff\xff\x6a\xfa\xd5\xb5\xf1\x00\x00\x00")

func v1alphaSchemaSchemaSwaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_v1alphaSchemaSchemaSwaggerJson,
		"v1alpha/schema/schema.swagger.json",
	)
}

func v1alphaSchemaSchemaSwaggerJson() (*asset, error) {
	bytes, err := v1alphaSchemaSchemaSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v1alpha/schema/schema.swagger.json", size: 241, mode: os.FileMode(420), modTime: time.Unix(1586631429, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v1alphaDepsDepsSwaggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xcd\x31\xce\x83\x30\x0c\x05\xe0\x3d\xa7\xb0\x3c\xff\x82\xbf\x1d\xb9\x4a\xd5\x21\x02\x03\xae\xa8\x1d\xc5\x86\x0e\x88\xbb\x57\x49\x2e\xd0\x25\x8a\xde\xfb\x5e\x72\x06\x00\xb4\x4f\x5c\x16\xca\x38\x00\xde\xbb\x7f\xfc\x2b\x19\xcb\xac\x38\x40\xe9\x01\xd0\xd9\x37\x2a\xfd\x71\x8b\x5b\x5a\x63\x3f\x51\xb2\x7a\x74\x29\xab\x6b\xdd\x00\xe0\x41\xd9\x58\xa5\xca\x76\x05\x51\x07\x23\xc7\x00\x70\xd5\x97\x47\x15\xdb\xdf\x64\x38\xc0\xa3\xad\x62\x4a\x1b\x8f\xd1\x59\xa5\x7f\x99\x4a\xb1\xcf\x6a\x53\xd6\x69\x1f\x7f\xb4\xd1\xd7\x02\xcf\xf6\xcd\x44\x33\x0b\x17\xd7\xc2\x70\x85\x6f\x00\x00\x00\xff\xff\x36\xf0\x74\x80\xed\x00\x00\x00")

func v1alphaDepsDepsSwaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_v1alphaDepsDepsSwaggerJson,
		"v1alpha/deps/deps.swagger.json",
	)
}

func v1alphaDepsDepsSwaggerJson() (*asset, error) {
	bytes, err := v1alphaDepsDepsSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v1alpha/deps/deps.swagger.json", size: 237, mode: os.FileMode(420), modTime: time.Unix(1586631429, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v1alphaStoreStoreSwaggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x54\xc1\x6e\xe2\x30\x10\xbd\xe7\x2b\x2c\xef\x1e\x11\x2c\x1c\xb9\xad\x04\xbb\x6a\x55\xb5\x88\x56\xea\xa1\xe2\x60\x92\x49\x30\x4d\x3c\x96\x3d\xa1\x42\x28\xff\x5e\xd9\x69\x42\x80\xd0\xa6\x82\x5e\x50\x34\xf3\xde\xf3\x7b\x1e\x33\xbb\x80\x31\x6e\xdf\x44\x92\x80\xe1\x63\xc6\x47\xfd\x3f\xbc\xe7\x6a\x52\xc5\xc8\xc7\xcc\xf5\x19\xe3\x24\x29\x05\xd7\xdf\x0c\x45\xaa\x57\x62\x60\x09\x0d\x94\xbf\x7d\x6d\x90\xd0\xb3\x18\xe3\x1b\x30\x56\xa2\xf2\xd8\xf2\x93\x29\x24\x66\x81\x78\xc0\x58\xe1\xb5\x43\x54\x36\xcf\xc0\xf2\x31\x7b\x29\x59\x42\xeb\x54\x86\x82\x24\xaa\xc1\xda\xa2\x72\xd8\x85\xc7\x6a\x83\x51\x1e\x76\xc4\x0a\x5a\x39\xe0\xae\x3c\x26\x82\x58\x2a\xe9\x70\x76\x9f\xc4\x5b\x9e\x40\x0a\x04\x73\xb0\x1a\x95\x85\xba\xe9\x82\x6e\xb5\xcf\x89\xcb\x35\x84\xde\xf2\x87\xe9\x8a\xfa\x4f\xaa\xa8\x03\xb1\x57\xd5\xb5\x41\x0d\x86\x24\xd8\x06\xda\x7b\x95\xe6\xb0\xd4\x10\x11\xc6\x88\x6d\xad\xe1\x5b\x92\x20\x3b\xc6\x33\xc6\x7f\x1b\x88\x1d\xe3\xd7\xa0\x91\xb6\x9c\xcb\x7f\x23\xf4\xea\x86\x20\x9b\x09\x69\x78\x83\x56\x04\xc7\x5f\xc5\x69\xce\x9a\x7d\x49\xc8\xa4\x12\x79\x2a\x89\xad\x61\x2d\x19\xa9\x92\xbd\xc1\x62\x9f\x9b\xbf\x0e\xbf\x60\x1d\xdc\x51\x8c\x26\x13\xe4\xba\xcb\x2d\xc1\x19\xc5\xd1\xb5\x15\x41\x85\x18\x39\xe6\x91\x6e\x97\xd1\x4c\x2b\x6e\xab\x72\x7d\x7d\x13\x41\xe2\x2a\xb6\x3b\xcc\x7b\x7a\x1a\xe7\xdc\x59\x1c\x54\x9e\xd5\x7f\x4c\x5f\x99\xff\x7d\x6e\x58\xe1\xb7\x8f\x0f\xf7\xd5\xf1\x8b\x9a\x16\x41\x2c\xf2\xd4\xfb\x73\xf8\x4f\xcc\xf8\xa7\x7b\xc1\x03\x84\x28\x39\x79\x77\x5d\xe6\xd2\x3e\x0f\x85\xd1\xa5\x6a\xe7\x07\x70\x27\x2d\x5d\x63\xb1\xb4\x2d\x8a\x1f\x5a\x2c\xdf\x5e\x2a\xb3\xbc\x4b\xc4\xa0\x62\x17\x41\x11\xbc\x07\x00\x00\xff\xff\xc3\xd9\xf5\x92\xa1\x06\x00\x00")

func v1alphaStoreStoreSwaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_v1alphaStoreStoreSwaggerJson,
		"v1alpha/store/store.swagger.json",
	)
}

func v1alphaStoreStoreSwaggerJson() (*asset, error) {
	bytes, err := v1alphaStoreStoreSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v1alpha/store/store.swagger.json", size: 1697, mode: os.FileMode(420), modTime: time.Unix(1586631429, 0)}
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
	"v1alpha/tracker/tracker.swagger.json":     v1alphaTrackerTrackerSwaggerJson,
	"v1alpha/extractor/extractor.swagger.json": v1alphaExtractorExtractorSwaggerJson,
	"v1alpha/schema/schema.swagger.json":       v1alphaSchemaSchemaSwaggerJson,
	"v1alpha/deps/deps.swagger.json":           v1alphaDepsDepsSwaggerJson,
	"v1alpha/store/store.swagger.json":         v1alphaStoreStoreSwaggerJson,
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
	"v1alpha": &bintree{nil, map[string]*bintree{
		"deps": &bintree{nil, map[string]*bintree{
			"deps.swagger.json": &bintree{v1alphaDepsDepsSwaggerJson, map[string]*bintree{}},
		}},
		"extractor": &bintree{nil, map[string]*bintree{
			"extractor.swagger.json": &bintree{v1alphaExtractorExtractorSwaggerJson, map[string]*bintree{}},
		}},
		"schema": &bintree{nil, map[string]*bintree{
			"schema.swagger.json": &bintree{v1alphaSchemaSchemaSwaggerJson, map[string]*bintree{}},
		}},
		"store": &bintree{nil, map[string]*bintree{
			"store.swagger.json": &bintree{v1alphaStoreStoreSwaggerJson, map[string]*bintree{}},
		}},
		"tracker": &bintree{nil, map[string]*bintree{
			"tracker.swagger.json": &bintree{v1alphaTrackerTrackerSwaggerJson, map[string]*bintree{}},
		}},
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
