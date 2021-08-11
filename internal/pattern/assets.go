// Code generated by go-bindata. DO NOT EDIT. @generated
// sources:
// DefaultHttpPattern.json
// DefaultChannelPattern.json
package pattern

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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _defaulthttppatternJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x58\x5f\x6f\xdb\x36\x10\x7f\xcf\xa7\xb8\x0a\x41\x53\x03\xa9\x92\xfd\x79\x32\xe0\x87\x38\x0b\xd6\x75\x0b\x56\xb4\xee\xfa\x30\xec\x81\x96\xce\x16\x63\x8a\xd4\xc8\xa3\x33\xaf\xc9\x77\x1f\x44\x49\x0e\xf5\xc7\xb2\x1d\xc3\x01\x86\xbe\x24\x11\xef\x78\xfc\xdd\xef\x8e\x77\xbc\x7c\x3d\x01\x08\x24\x4b\x31\x18\x42\xf0\x13\xce\x98\x15\xf4\x8e\x28\xfb\xc0\x88\x50\xcb\xe0\x3c\x97\x1b\xc2\xcc\x04\x43\xf8\xf3\x04\x00\xe0\xab\xfb\x09\x10\xf0\x59\xbe\xe9\x34\x8c\x94\x9c\x85\xd6\xe0\x47\x46\xf8\x1b\x4f\x39\xa1\x86\xd1\x08\x48\x5b\x74\xfb\x9d\xb2\x41\xbd\xe4\x91\x3b\xc6\xd3\x7b\x92\x73\x99\x59\x0a\x86\x6b\xeb\x00\x01\xa9\x05\xca\x7c\xc3\x5c\xa8\x29\x13\x41\x29\x79\x74\xbf\x1f\xcf\xfb\xc1\xbc\xff\x32\xe9\x05\xf1\xfe\xcb\xe4\x0f\x26\x78\xcc\x48\xed\x88\x62\x74\x1a\x66\x6c\x25\x14\x8b\xc3\x04\x59\x8c\xda\x84\x57\x96\x12\xa5\xf9\xbf\x8c\xb8\x92\x6b\x2b\xf9\x41\x7c\x2e\xb9\x9c\xdf\x22\x25\x2a\x2e\x37\x3b\x64\x77\xf7\xf4\xa9\x26\xf3\x36\x2d\x70\xd5\x50\xfd\x15\x57\xbe\x02\xb3\x4d\x5b\x57\xb6\x66\x81\x1b\xd3\x50\xf8\xc5\x98\x1a\x2e\x3b\x6d\xa2\xb1\xd3\xfd\x88\xbd\xe6\x3a\xb2\x9c\xc6\x1a\xd9\x62\x4b\xa0\xeb\xaa\x41\x8f\xf9\x37\xad\xc0\xcd\x98\x30\x08\x0f\x0f\x70\x1a\xfa\xa1\x0a\x95\xa5\xcc\x92\x09\x97\xf9\x4a\x75\xfa\x00\x5e\xbf\x86\xd2\x44\xea\x78\xcd\x25\x67\x3f\xdf\x4c\xce\x02\xe8\x42\x96\xa7\xf8\x98\x45\x0b\x94\xf1\xd5\x13\xf4\x84\x09\x2a\xd1\xf8\x0a\x21\x6a\xad\x34\xbc\x1a\x81\xe4\xc2\x1d\xf5\xca\xad\x84\xdc\x48\x24\xf7\x67\xe7\x8e\xc1\xb1\x3d\x7e\xd3\x76\xf9\xc3\xe7\xc9\x59\x61\xa3\x25\xf9\xfd\xd3\x46\xd1\xd5\xe4\xfa\xdd\xd9\xa0\x33\x86\x9e\x5b\xe3\xde\x7b\x12\x29\x49\x28\xc9\x4f\xaf\x6a\xa9\xca\xaf\x2d\x44\x8f\xf7\x26\x7a\xbc\x03\xd1\xdb\x32\xd7\x11\xd9\x17\xf1\xcd\xb9\x35\x78\x78\x80\x3e\x17\x8e\x10\xa3\x41\x77\x94\x1a\x37\xad\x2f\x4e\x2a\x43\x5d\xd4\xab\x61\x1e\x34\x2b\xa9\xba\x9a\x87\x55\x80\x1e\x1e\x47\xcf\xe6\x71\xf4\x3f\xe1\x51\xa3\x41\x6a\xb0\x78\x02\xf0\x97\x6b\x9f\x1a\x4d\xa6\xa4\xc1\x9e\x16\xea\xb5\xc4\xf5\x7d\x17\xf9\xf7\x47\x64\x51\x82\x71\xbb\xcc\x3a\x7e\x82\xa1\x5b\x5d\x2f\x16\x5b\x9b\xf7\x32\xce\x5d\xfb\xf1\xf2\x07\xaf\x0f\xc4\x8c\x58\x4d\xcd\x75\x79\x46\xd6\x54\x0d\x1a\x1c\x1c\xb8\xf9\x27\x42\x8c\x31\x86\xb7\x30\x49\x10\x4a\xaa\x60\xa5\x2c\x24\x6c\x89\xa0\xf1\x6f\x8b\x86\x30\x06\x6e\x40\x2d\x51\x03\x25\x08\x4c\x08\x75\x8f\x31\x38\x17\xc2\x60\x7d\xca\xe3\x73\x5b\x78\x91\x3b\xfd\x65\xd1\xd5\xcf\x43\x08\xfa\x6e\x0b\x41\xce\xe2\x2d\x1a\xc3\xe6\x58\x96\xb9\x4e\x44\x35\xbd\x73\xdf\xc2\xb2\xd0\xe5\x4a\xee\x60\xa6\xad\xfc\x4c\x1e\x37\xdf\xd7\xd3\xb0\x2e\x5b\x9f\x4d\x9a\x67\xd9\xcb\xa4\x5d\x65\x31\x88\x0a\x28\x30\x2d\x71\x96\x18\xf6\xf3\xfa\x88\x0f\x89\x16\x07\xce\xf2\x2e\x24\x7c\x7f\x79\xd9\x26\x21\x8f\x7b\xad\x56\x56\x80\x9c\xfc\x85\x5c\x3d\xda\x0b\xe2\x28\x1c\x8d\xf7\xe0\x68\xff\x44\xed\x3a\xff\xc6\x59\xd9\x50\xd7\xcb\x62\xd8\x51\xd6\xab\x81\xaa\x73\xd2\x89\xd1\x44\x9a\x67\x55\xeb\x70\xc5\x56\x34\x95\x34\xba\x18\xcf\x39\x25\x76\x1a\x46\x2a\xbd\xc8\xb4\xba\xc3\x88\xde\xce\x84\x9a\xab\x8b\x94\x47\x5a\xcd\x19\xe1\x3d\x5b\x5d\xb0\x88\xf8\x92\xd3\xea\x42\x33\xc2\x96\x29\x83\x44\x5c\xce\x4d\xdd\x6d\xa7\xe6\x3f\xd6\x74\x85\xb5\x9f\xd7\xca\xb5\xee\xf9\xa9\xe1\x5b\xa9\x81\x60\x54\x8a\xe0\x26\x29\x73\xa8\x8f\x77\xf7\xd4\xfd\xdc\xab\x90\x6d\xea\xe1\x0d\x6c\xd7\x8d\x72\x53\xbd\x02\x0e\x84\x57\x56\xb1\x69\xf3\xf4\xee\x28\xa4\x45\xf2\xad\x83\xe0\xbe\xbd\x3c\xa4\x44\xa3\x49\x94\xa8\xcd\x7d\x4f\x8b\xbe\x26\x4f\x51\xd9\x5a\x44\xab\x25\x4f\x2b\x43\xcd\xeb\x03\x69\xb9\xb2\x53\xcc\xbb\xa7\xa6\x06\xaf\xb7\x6c\x81\xc0\x24\x24\x44\x19\x44\x4c\x08\x20\x95\xbf\x17\x34\x4c\x8b\xad\xbb\x33\x9c\xcf\x0f\x9a\x4f\xbd\xfc\x46\x43\x5b\x29\x6d\xcd\xdc\x69\x6b\xd2\xb6\x9a\xfb\x0a\x25\xb0\xcf\x7a\xcb\xff\x18\x3a\x78\x18\x7f\x13\x3c\xe4\x25\xef\xe4\xf1\xbf\x00\x00\x00\xff\xff\x09\x07\x8f\xf1\x2c\x12\x00\x00")

func defaulthttppatternJsonBytes() ([]byte, error) {
	return bindataRead(
		_defaulthttppatternJson,
		"DefaultHttpPattern.json",
	)
}

func defaulthttppatternJson() (*asset, error) {
	bytes, err := defaulthttppatternJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "DefaultHttpPattern.json", size: 4652, mode: os.FileMode(420), modTime: time.Unix(1551218260, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _defaultchannelpatternJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x53\xc1\x8e\xda\x30\x10\xbd\xf3\x15\x53\x6b\xb5\x27\x1a\x68\xd5\x13\x12\x87\x6d\x7b\xda\x6a\xa5\x4a\x5d\x75\x0f\x55\x0f\x83\x33\x80\xc1\x78\x22\x7b\x1c\x94\xee\xf2\xef\x55\x1c\x12\xb2\x94\x05\x7a\x49\xa4\xf1\x7b\xcf\xef\xcd\x78\x9e\x07\x00\xca\xe1\x86\xd4\x04\xd4\x57\x9a\x63\xb4\xf2\x65\x89\xce\x91\xfd\x8e\x22\xe4\x9d\x1a\xd6\x90\x20\x54\x04\x35\x81\x5f\x03\x00\x80\xe7\xf4\x05\x50\x66\x5e\xf3\x6e\x32\xcd\x6e\x9e\xc5\x40\xf7\x4f\x8f\x30\x9d\x82\xf8\x48\x89\x97\x40\x81\x7c\x69\x74\xba\xe1\xfe\xe9\xf1\x27\x5a\x93\xa3\xb0\x3f\x00\x8c\x2b\xa2\xa8\x49\x27\x0b\xa0\x84\xd7\xe4\x6a\xc6\xf4\x26\x2b\xb0\xb2\x8c\x79\xb6\x24\xcc\xc9\x87\xec\x2e\xca\x92\xbd\xf9\x83\x62\xd8\x75\x2a\x00\x6a\x4d\xd5\x9e\x92\xfc\xac\xb6\xf2\x8d\x2a\xb5\x3f\xdf\xa5\xff\x6e\x78\x55\x80\x39\xda\x40\xf0\xf2\x02\x37\x59\xdf\x72\xc6\x51\x8a\x28\x21\x2b\xeb\xca\xd9\xa4\x77\xba\x76\xf7\x19\xf5\x9a\x5c\x7e\x36\xaa\x6e\xda\xdd\x77\xde\x96\x7a\xd9\x4a\xb4\x91\xfa\x98\xa6\xd0\x86\xeb\x2e\x58\xa2\xad\xf5\x55\x4a\xd0\x1c\xd7\xc9\x7f\xa7\x29\x7a\x0a\x05\xbb\x40\xff\x37\x49\xb8\xbd\xbd\xd8\x87\xe6\xba\xce\x05\x79\xcf\x5e\x4d\x12\xbd\x2b\x36\xa4\xa3\xf0\x9c\xd7\xa9\x3e\x8d\x3f\xf4\xb2\xe6\x28\xf8\x0a\xd6\x2a\x3e\x50\x08\xb8\x68\xdb\x70\xd2\xd1\x2b\xdc\xb0\xaf\x50\x36\x58\xc3\xee\x0a\x99\x7f\xc1\x9d\xd4\xee\xec\x83\x6a\x93\xa7\x86\x5c\x13\xfd\xe3\x78\x7c\x21\x7a\x3b\xb5\xda\xef\x8f\xa8\x35\x85\xf0\xee\x4d\x3b\xdd\xac\xf7\x6f\xf1\xc4\xa8\xdb\x75\x3f\xbd\x8c\x39\x05\xed\x4d\x91\x76\x6b\x02\x6a\x8f\x20\x08\xbc\x21\x48\x6b\x19\x0e\x60\x4f\xe9\xd9\x2c\x8c\x2c\xe3\x2c\xd3\xbc\x19\x15\x9e\x57\xa4\xe5\xfd\xdc\xf2\x82\x47\x1b\xa3\x3d\x2f\x50\x68\x8b\xd5\x08\xb5\x98\xd2\x48\x35\x5a\x6d\x45\x9d\x6c\x5e\xeb\xec\x8d\xe5\x39\xb2\xf6\x80\x6b\x02\x74\x80\x09\x0d\x1a\xad\x05\x61\xa8\x38\x7a\x98\x1d\x73\x2f\x39\xd5\xec\xc4\x9b\xd9\xc1\x64\xbb\x83\x87\xb6\x0e\x76\x83\xbf\x01\x00\x00\xff\xff\x9e\xe9\x59\x58\x2f\x05\x00\x00")

func defaultchannelpatternJsonBytes() ([]byte, error) {
	return bindataRead(
		_defaultchannelpatternJson,
		"DefaultChannelPattern.json",
	)
}

func defaultchannelpatternJson() (*asset, error) {
	bytes, err := defaultchannelpatternJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "DefaultChannelPattern.json", size: 1327, mode: os.FileMode(420), modTime: time.Unix(1553025640, 0)}
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
	"DefaultHttpPattern.json": defaulthttppatternJson,
	"DefaultChannelPattern.json": defaultchannelpatternJson,
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
	"DefaultChannelPattern.json": &bintree{defaultchannelpatternJson, map[string]*bintree{}},
	"DefaultHttpPattern.json": &bintree{defaulthttppatternJson, map[string]*bintree{}},
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

