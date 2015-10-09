package pkgcloud

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _assets_distributions_json = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x9a\xcd\x72\xdb\x36\x10\xc7\xef\x7e\x0a\x8c\x2e\xbd\xa4\x1a\x81\xd4\x67\x6e\x75\x6c\xc7\x69\xe4\x38\xf5\x47\x72\xe8\x64\x32\x20\x09\x49\x88\x29\x90\x06\x49\xcb\x4c\x26\xaf\xd5\x17\xe8\x8b\x15\xa0\x2d\x8d\x3e\xa8\x5d\xd0\x93\xda\x17\x5c\x14\x8f\xb1\x06\x7f\x59\x2d\xf6\xbf\x0b\xee\x8f\x03\x42\x5a\x11\x0f\x5a\xaf\xc9\xdf\xfa\x47\x42\x7e\x54\x9f\xe6\x97\x22\x4b\x63\x56\x7e\x95\x6c\xce\xf5\x6a\xeb\x3a\x28\x64\x5e\xb4\x5e\x2d\xd7\x85\x8c\xf8\xfd\x6a\xb5\xd8\x5a\xbd\xe3\x2a\x13\x89\xcc\x56\xfb\xae\xef\xfd\xf0\xf7\x91\x5e\xa3\xaf\xd6\x7f\xb5\xfd\xc8\x6e\x9b\x76\xc8\x67\xa6\xf2\xb2\xfa\x9c\x25\xd3\xd6\x86\xfd\x26\xc2\xc2\x18\x6e\x1a\x3c\x52\x7c\x95\xc5\x3c\xe0\x6a\xb9\x65\x6b\x65\xf2\xf3\x15\xc0\xe6\x81\x6c\xbd\x76\xa7\x4b\x4e\x13\xa6\x4a\x72\xca\xa3\x29\x47\xe0\x66\xc6\x12\x83\x33\x7b\x5a\xc2\xf9\x08\x9c\x76\xdc\xa1\xe2\xfc\x7b\x49\x0e\x99\xa6\x53\x10\x5b\x50\x19\xe2\x70\xd6\x9e\xeb\x82\x70\xfd\x76\xa7\x4f\xc6\x57\x97\xe4\x88\xa5\x29\x57\xe4\x48\xb1\x1b\x0e\xf1\x45\x95\x1d\xc6\x67\xb6\xb5\xe4\xeb\x21\x7c\xda\x79\xc7\xd1\xb4\x24\xc7\x93\x1c\xe2\xd2\x5f\x3b\xea\xb5\xbe\xbd\xd7\xfa\x20\xd5\xc0\xc4\xdb\x09\x17\x99\x3e\x0c\x27\x6c\x21\x21\xb0\x49\x65\x86\xa1\x0d\xec\xa3\x6d\x80\xa0\x69\x87\xbd\x2d\xf2\xac\x24\x6f\x45\x10\x24\x20\xdb\xd4\xd8\xe1\x68\xd6\x5e\x1b\x82\x68\x43\xe3\x35\x13\x6b\xa7\x4c\x45\xe6\xa4\x2a\x98\x6e\x66\xcc\x30\xba\xa1\xbd\xe3\x46\x08\x9d\x76\xdc\x3b\x99\x2b\x9e\x8a\x88\xbc\x0b\xf8\x3d\xc4\x26\x1e\x0d\x71\x3c\x6b\xe7\xd1\x0e\xc8\x37\x32\xde\xfb\x93\xe9\xcc\x5e\xea\x7f\xc2\x1b\x16\x27\x29\x78\x52\xbf\x55\xb6\x18\xe0\xc8\xde\x7f\x14\x16\x88\x91\x71\xe0\x7b\xa6\xe6\x22\x24\xef\x13\x16\x33\x08\xee\xa6\xb2\xc3\xe1\xec\xbd\x07\x2b\x04\xed\x2c\x83\x6f\x5c\x84\xfa\xfb\x1d\x97\x12\xfc\x7e\x63\x63\x85\xe1\x55\x9b\xda\xf2\xc1\x22\xa1\xb7\xd2\xde\x3b\x63\xfa\x21\x22\xbc\x21\x67\x9c\xab\x1b\x06\x26\xbc\xf9\xa3\xad\x05\xa4\xbd\x13\x61\xb1\xa0\xd4\x38\xf1\x03\xcb\x75\x08\x7e\x60\x6a\x31\x63\x31\x44\x28\x8d\x21\x8a\x47\x1b\xf8\x10\xd6\x0a\xbd\x95\xf6\xe1\xb9\xe4\x42\xfb\x85\x9c\x87\x3c\x4e\x40\x0f\x26\x0f\x96\x16\x84\xf6\x0e\x84\x75\x83\x7a\xcb\x28\xfc\xa8\x78\x28\x32\x4e\x3e\x32\x39\x4d\x62\x01\xe6\xc1\xf4\xc1\x16\xe5\xf4\x1a\x78\x12\x16\x11\xbd\x95\xf6\xe4\x5f\x05\x93\x39\x8b\xf5\xbf\x3c\xff\x0e\x7f\xd5\xb7\x0f\xa6\x16\x88\xf6\xae\x84\xc5\x84\xfa\xc6\x95\x17\x4c\x09\x39\x25\x17\xfa\x23\x67\x02\x44\x54\x95\x29\x4a\xe8\x37\x70\x22\x2c\x28\x7a\x2b\xed\xc4\x4b\x56\x84\xa5\xfe\x8c\xd9\x9c\x69\x1e\xb0\xf6\xcb\x8c\xad\x05\xa1\x7d\xd5\x0c\x4b\x0a\xed\x2e\xc3\xf1\x4a\x15\xa6\x96\xb9\x62\x33\x10\x30\xaf\xcc\x50\xc2\x6e\x03\x1f\x76\x91\xbc\x5d\xb5\x1d\xd7\x79\x92\xea\x23\x7d\x2d\x45\x98\x28\xf0\xac\x14\x95\xa5\x05\x61\x83\x9c\x88\x04\x62\xd5\x7c\x7c\x12\x77\x5a\x55\x3e\x71\x75\xc7\xc1\x9c\x73\x67\xec\x50\xbc\x06\xbd\x07\xed\x21\x0e\xac\xda\x8f\xcf\x22\xd6\x6d\x1b\x57\x7c\x91\xc4\x13\xb0\x6f\xd3\x86\x16\x78\x1b\xde\x7b\xfc\xe9\xcb\xc1\x1a\xec\xde\xce\xf5\x88\x07\x82\xc9\x7d\x9d\x6b\xb4\xb5\x6a\xd9\xb9\x7a\x58\xeb\xda\x21\x3c\x0f\x67\x60\xff\xb0\xb3\x5e\xd7\xaf\x5a\x1f\x3c\xb4\x5f\x25\x31\x97\xb2\x04\x4b\x90\x5d\x83\xda\x2e\xd5\x16\x09\x2e\x40\x74\xc7\x46\xb2\xdb\x42\x37\x9f\x60\x51\x59\x6b\x52\xdb\xff\xd9\x62\xc1\x25\xc7\x80\x2c\x66\xbb\x0d\xf1\x56\xcc\xd6\x58\xd4\x76\x58\xb6\x48\x70\x99\x31\x24\xdf\x78\x96\x09\xb8\xf6\xae\xb1\xa8\xed\x5d\xac\x93\x10\xac\x35\x23\x92\xe9\x7e\x04\x09\xf1\x5a\x93\xda\x8e\xc0\x3a\xf7\x20\xfa\xd2\x21\x81\x16\x0c\xe4\xc2\xa3\xc6\xa2\xbe\xd4\x7e\x6a\xca\xf9\x43\x96\x44\x27\x96\xdf\x03\x96\xf1\x88\xe8\xd5\x5c\x89\xa0\xc8\xf5\x03\xda\xfb\xf2\x10\x5b\x3b\x79\xb6\x49\xc8\x87\xc3\xc6\x50\x3c\xee\xd4\x86\xfc\xc1\xd0\x43\x2f\x8b\x38\x7e\xa2\x2b\xc6\x42\x16\xf7\x67\xba\x79\xdd\xf7\x1f\x8f\x8d\xc1\x7c\xc3\xc0\xf6\xf6\xb0\x87\xd4\x94\x7d\x92\xf2\x5c\x81\x5d\x61\x8d\x41\x4d\x28\xd8\xde\x2c\xd1\x1e\xa2\xdc\x03\x72\xab\xd5\x06\x24\xaa\x31\xa8\x21\x1a\x58\x13\x21\x25\xe3\xa0\x4d\x89\xe2\x01\x0f\x43\x90\xaa\xd6\xa4\x8e\xab\x4d\xad\x1b\x17\xe4\x28\x0f\xda\x1e\x51\x6c\xc2\x38\xdc\xd7\xd7\x9a\xd4\x92\x79\xd6\x64\xb0\xb8\xeb\xad\x7c\xa2\x92\x0c\xc6\xda\x59\xaf\x65\xf2\x9f\x9a\x62\x2e\x58\x96\x42\x75\x8d\xda\x59\xb7\x3e\x55\xb0\x6c\xbf\x84\x3e\xd2\x1e\xac\xd9\x2f\x22\x90\x3d\x38\xfb\xbe\x94\x40\x22\xf7\x01\xcf\x24\x90\xfa\xf3\x8b\xf9\xfb\x96\x4a\xe7\x2d\xf4\xe5\xd2\xb1\xd4\xcf\x4b\x95\xb9\x9e\xa8\xf4\x62\x5f\x50\xf3\xb8\x79\xa1\x0e\xfb\x63\xfb\xc9\xc4\x94\xb7\x80\x6b\x7a\xbf\xb0\x3c\x86\xd5\x6b\x87\xac\x0f\x93\xf5\x7f\x5d\x85\x4c\xbb\x70\x6a\xde\x41\x1b\xc0\x68\x83\x66\x99\xa0\x51\x1e\x3c\xe1\x51\xb2\xa6\xe0\x3b\xaf\x42\x36\x57\x6d\x83\x06\x91\xf1\x2e\x19\xb3\x62\x3a\x43\x6e\xd1\x68\xd7\xe6\xde\xc2\x36\x5a\x10\x1d\xef\x91\x71\x72\xc7\xe3\x64\xfb\x7a\x76\x0b\x09\x8d\x5f\x6a\x1f\xc0\x3e\x22\xe0\x7d\x73\x3b\x21\xc1\xa4\x4c\xd1\xa8\xa5\xf6\x61\xeb\x63\xb2\x4d\x0e\x39\x9f\x94\xe4\x4c\x28\x16\xc6\x30\x17\x1a\xb2\xd4\x5e\xbd\x7c\xe4\xbe\x64\x48\x2e\xd3\x99\xb9\x59\x67\x31\x79\x93\x2c\x40\xae\x21\xca\x65\xaf\x60\x3e\xf2\x82\x60\x44\x2e\xc3\x99\xfa\xf7\x9f\x48\xc8\x29\x57\xbf\x65\xe4\x0d\xfc\x7e\x80\x8e\x50\x38\x7b\x21\xf3\x61\xc5\xf7\x3a\xe4\x94\xeb\x04\x24\x83\x02\x7c\xf1\xee\x75\x30\x26\xaf\xd3\x20\x2f\xc2\x1e\xf3\x28\x79\x48\x47\xc4\xa3\x20\x14\x45\xa1\x68\x03\x28\x58\x47\x3c\x6f\x05\xe5\x81\x50\x1e\x0a\xe5\x35\x28\x43\xe0\xb3\xe8\xf9\x2b\x28\x1f\x84\xf2\x51\x28\xff\xe9\xda\x71\x19\x0a\x2e\x73\x31\x11\x21\x5c\x76\x64\x2b\xbb\xe6\xd5\xb4\x0f\x4b\xc9\x36\xc2\x33\xd6\x1f\xd4\x87\x25\x65\x07\xed\x39\x0b\x10\x38\x7c\x76\xd0\x5e\xb0\x00\x39\xaf\xc4\x04\x0e\xa0\xa4\x79\xdd\x4a\xbb\x70\x02\x5c\x7f\xec\x73\x06\x4d\x17\xee\x7a\x36\xb0\x9e\x33\x60\xe0\x32\x7f\x03\xeb\xff\x0c\x96\x83\x65\xdf\x13\x65\x21\xde\xf7\xb8\xa1\x3a\x37\x54\xe7\x86\xea\xdc\x50\x9d\x1b\xaa\x73\x43\x75\xfb\x10\xdd\x50\x9d\x1b\xaa\x7b\xed\x86\xea\xf6\xf3\xb9\xa1\x3a\x80\xd3\x0d\xd5\xe1\x84\x6e\xa8\xae\x9e\xd0\x0d\xd5\xed\xe5\x73\x43\x75\x6e\xa8\xce\x0d\xd5\xad\x63\xb9\xa1\x3a\x0b\x24\x37\x54\xe7\x86\xea\xdc\x50\x9d\x1b\xaa\x5b\xd6\xb2\x6e\xa8\x6e\x9d\xcc\x0d\xd5\xb9\xa1\x3a\x37\x54\x87\xbc\x5c\x3a\xf8\x79\xf0\x5f\x00\x00\x00\xff\xff\x4c\x0b\xd2\x5a\xb5\x41\x00\x00")

func assets_distributions_json_bytes() ([]byte, error) {
	return bindata_read(
		_assets_distributions_json,
		"assets/distributions.json",
	)
}

func assets_distributions_json() (*asset, error) {
	bytes, err := assets_distributions_json_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "assets/distributions.json", size: 16821, mode: os.FileMode(420), modTime: time.Unix(1444416535, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	"assets/distributions.json": assets_distributions_json,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"assets": &_bintree_t{nil, map[string]*_bintree_t{
		"distributions.json": &_bintree_t{assets_distributions_json, map[string]*_bintree_t{
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

