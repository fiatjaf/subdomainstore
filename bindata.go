// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// static/index.html
package main

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
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
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

// Mode return file modify time
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

var _staticIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x56\x4f\x6f\xdb\x36\x14\xbf\xeb\x53\xfc\x56\x0f\x53\x03\xc4\x92\xb3\x01\x3b\xb8\x8a\xb3\x60\x0b\xd0\x43\xb3\x0d\x58\x0e\x39\x96\xa6\x9e\x2c\x36\x14\x9f\x4a\x3e\x45\xd1\xba\x7c\xf7\x81\xb2\xec\x3a\x59\xd1\x35\x01\xda\x83\x75\x10\x1f\x7f\xff\xe4\xc7\xc7\xa2\xf5\x04\x53\x9e\xbe\x68\x28\x04\xb5\xa1\x17\xab\x22\x6f\x3d\xad\x92\xa4\x08\xda\x9b\x56\x56\x09\x30\xad\x65\xc6\x39\xf2\xaf\xaf\x2e\xdf\xe0\x14\x6f\x93\x19\xd6\xdd\x00\x85\xd0\xad\x4b\x6e\x94\x71\xe0\x0a\xdf\x7f\xb0\xac\x95\x18\x76\x59\xcd\x41\x9c\x6a\xe8\xfe\xbb\x44\x77\xde\x22\x3d\x58\x6b\x3d\x0b\x6b\xb6\xf7\xcb\x3c\x7f\xb4\xe5\x3e\x5f\x77\x43\x5e\xec\x51\x57\x67\xaa\xe1\xce\xc9\xe9\xc9\x62\xb1\x58\xa4\xf8\x07\xef\xde\x63\xee\x91\x66\xad\x4f\x93\x64\x86\xbe\x26\x87\x56\x0d\xc6\x6d\x20\x35\xc1\xb8\x5b\x36\x9a\x8e\x21\xea\x86\xe0\x58\x28\x0a\x8b\x2b\xad\x1a\x1a\x72\x82\xd6\x33\x57\xd1\xa6\x69\xa2\xad\x64\x06\x23\xe8\x8d\xb5\x58\x13\x06\xee\x3c\x54\x27\x35\x7b\xf3\xf7\x28\x0b\xc2\x37\xe4\x62\x99\xaa\x84\xfc\xc4\x75\x8c\x2e\xd0\x08\x5b\xb1\xb5\xdc\x47\x7a\xcd\x4d\xa3\x5c\x19\x20\x9c\x65\x59\x14\x67\x4d\x90\x2d\xa4\x27\xcd\xbe\x0c\xcb\x6d\x18\xf3\xd7\x48\xcf\x0f\x49\x96\x28\x76\x82\x56\xe9\x97\x67\x75\x90\x53\x3e\x31\x6c\x23\x8a\xe4\xda\x93\x12\x82\x9a\xb8\xbf\x8c\x7a\x7e\x8d\x3f\xff\xf8\xeb\xea\x79\x12\x7e\xfd\xfd\xfc\xf2\x22\xef\xfb\x3e\x0f\xdc\x50\x5f\x93\xa7\x4c\x73\x93\x62\xb6\xcd\x77\xaf\x68\x2c\x9c\x74\xa1\xf2\xdc\xa0\xef\xfb\xec\x00\x2a\x7b\xcc\x03\x61\x3c\x00\xfd\xfa\x6e\x2e\xaf\x73\x6a\x94\xb1\x79\x7c\x64\xf1\x11\xc8\xdf\x92\x8f\xec\x67\xad\x37\xec\x8d\x0c\xa7\x27\x8b\xff\xd8\x73\xb8\xbc\x7e\x60\x6e\x84\xf9\x5f\x7b\x9f\xa0\x41\x6f\xa4\xc6\x8e\x0b\x27\x8b\xaf\xef\xfa\x3c\xff\x25\x3f\xf9\xf1\xa7\xec\xe0\x77\xd6\x7a\xbe\x1b\x4e\xc5\x77\xf4\x09\xaf\xe7\x0f\xac\xc6\x8e\x70\xea\x86\xca\x83\x83\xe1\xe5\xe7\x9c\x1f\x45\xeb\x8f\x08\xa1\x5c\x89\x91\x34\xe2\x05\x82\xa7\xf7\x1d\x05\x09\x90\xda\x73\xb7\xa9\xa1\x2d\x77\x65\x65\x95\x27\xbc\x1c\x0c\xd9\x32\xf6\x2c\x37\x4a\x8c\x46\xe5\x89\x50\x8b\xb4\x01\x9a\xbc\x98\xca\x68\x25\x14\x8e\x62\x4f\x94\x64\xe9\x39\x3d\xf1\xdb\xc5\x9b\x8b\xab\x8b\xe7\x25\x5a\xc8\xd0\xd2\x2a\x2f\xe2\x69\xb8\xda\x07\xb8\x53\x62\xed\xee\x6c\xd8\x7e\xec\x8d\xb9\x25\x87\xb8\x65\x4c\x21\x6e\xfa\x76\x22\x9f\xa0\xee\x1b\x88\xfa\x8c\x9a\xf8\x2d\x3d\xe9\x5a\xf9\xcd\x74\x66\xd7\xdc\xf9\x80\x96\xbc\xe1\xd2\x68\x65\xed\x80\xc0\x71\x09\x25\xbb\x54\x60\x39\x4c\x95\x7b\x82\x27\x0e\xa6\x1d\xdf\x53\xa6\x53\xe4\xd7\xca\x41\xd9\xc0\xe3\xc4\x88\xff\xbc\x2e\xa8\xb5\x25\x58\xd7\x79\x3b\x6f\xd5\x00\x55\x96\x9e\x42\x88\x82\x8d\xc0\x04\x04\x61\x4f\x25\x8c\xdb\x2a\xee\x95\xb5\x24\x71\x02\xb9\x72\x0f\x19\x3d\xc6\x72\xb5\x89\x3d\x66\x55\x9c\x4d\xc2\xfb\x54\x92\xd9\xd8\x8b\x1f\x49\xc6\xfa\x31\xce\x51\x4c\xa8\xb9\x1f\xb1\x1e\x66\x82\x20\x4a\xba\x10\x67\x57\xdc\x79\x32\x0e\xb2\xb7\x99\xa7\xd6\x2a\x4d\x2f\xf3\x22\xdf\x1c\x23\xfd\xc1\xca\xab\xf4\xe8\xe3\xdb\xd5\xf6\xed\x26\xbe\x4d\x8a\x7c\x77\x77\x48\x8a\x20\x83\xa5\x78\x89\x98\x4d\xb7\x08\x7c\x48\x00\xa0\x51\x77\xf3\xde\x94\x52\x2f\xf1\xf3\x62\xd1\xde\xbd\x4a\x80\xfb\x24\x01\x32\xcf\xfd\x54\x53\x9a\xd0\x5a\x35\x2c\x51\x59\x1a\x0b\x80\x77\x5d\x10\x53\x0d\x73\xcd\x4e\xc8\xc9\x12\xa1\x55\x9a\xe6\xca\x73\xe7\xca\x2d\x46\x91\x4f\x9c\xff\x06\x00\x00\xff\xff\x29\x3d\xb1\xa0\xdc\x08\x00\x00")

func staticIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_staticIndexHtml,
		"static/index.html",
	)
}

func staticIndexHtml() (*asset, error) {
	bytes, err := staticIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/index.html", size: 2268, mode: os.FileMode(420), modTime: time.Unix(1600483989, 0)}
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
	"static/index.html": staticIndexHtml,
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
	"static": &bintree{nil, map[string]*bintree{
		"index.html": &bintree{staticIndexHtml, map[string]*bintree{}},
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
