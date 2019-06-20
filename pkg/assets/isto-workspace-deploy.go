// Package assets Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// deploy/istio-workspace/crds/istio_v1alpha1_session_crd.yaml
// deploy/istio-workspace/operator.yaml
// deploy/istio-workspace/role.yaml
// deploy/istio-workspace/role_binding.yaml
// deploy/istio-workspace/service_account.yaml
package assets

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

var _deployIstioWorkspaceCrdsIstio_v1alpha1_session_crdYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\x3d\x4e\x04\x31\x0c\x85\xfb\x9c\xc2\x27\x18\x34\x1d\x4a\x0b\x1d\x88\x02\x24\x7a\x6f\xc6\xbb\x6b\x6d\x26\xb6\x62\x67\x84\x84\xb8\x3b\xca\x64\xf9\x29\xb6\xcc\x97\x4f\xef\xbd\x04\x95\xdf\xa9\x1a\x4b\x89\x80\xca\xf4\xe1\x54\xfa\xc9\xa6\xcb\xbd\x4d\x2c\x77\xdb\x7c\x20\xc7\x39\x5c\xb8\x2c\x11\x1e\x9a\xb9\xac\xaf\x64\xd2\x6a\xa2\x47\x3a\x72\x61\x67\x29\x61\x25\xc7\x05\x1d\x63\x00\x28\xb8\x52\x04\x23\x1b\x41\x6c\xce\x32\x89\x52\xb1\x33\x1f\x7d\x4a\xb2\x06\x53\x4a\x5d\x3d\x55\x69\x1a\xe1\x96\x32\x72\xac\x5b\x00\xa3\xfd\x6d\x44\xee\x24\xb3\xf9\xd3\x7f\xfa\xcc\xe6\xfb\x8d\xe6\x56\x31\xff\x0d\xd8\xa1\x71\x39\xb5\x8c\xf5\x17\x07\x00\x4b\xa2\x14\xe1\xa5\xd7\x28\x26\x5a\x02\xc0\xf6\xf3\x19\xdb\x8c\x59\xcf\x38\x77\xaf\x1d\xea\xf5\xc5\xd7\x39\xe6\xe8\xcd\x22\x7c\x7e\x85\xef\x00\x00\x00\xff\xff\x66\x1b\x53\x70\x41\x01\x00\x00")

func deployIstioWorkspaceCrdsIstio_v1alpha1_session_crdYamlBytes() ([]byte, error) {
	return bindataRead(
		_deployIstioWorkspaceCrdsIstio_v1alpha1_session_crdYaml,
		"deploy/istio-workspace/crds/istio_v1alpha1_session_crd.yaml",
	)
}

func deployIstioWorkspaceCrdsIstio_v1alpha1_session_crdYaml() (*asset, error) {
	bytes, err := deployIstioWorkspaceCrdsIstio_v1alpha1_session_crdYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "deploy/istio-workspace/crds/istio_v1alpha1_session_crd.yaml", size: 321, mode: os.FileMode(436), modTime: time.Unix(1554454542, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _deployIstioWorkspaceOperatorYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x94\xcf\x6e\xe2\x30\x10\xc6\xef\x3c\xc5\x08\xf5\x5a\xda\x5e\x7d\x8b\x20\xcb\xa2\xb6\x10\x85\x6c\x57\x7b\x42\x83\x33\x10\x2f\x8e\xed\xb5\x1d\x10\xaa\x78\xf7\x55\xfe\x50\x68\x08\xe0\x5b\xc6\xe3\xef\xfb\x8d\x33\xe3\x8d\x50\x29\x83\x84\x72\x23\xd1\x53\x0f\x8d\xf8\x20\xeb\x84\x56\x0c\x7c\x13\x1c\x68\x43\xca\x65\x62\xe5\x07\x42\x3f\x6d\x5f\x7a\x06\x2d\xe6\xe4\xc9\x3a\xd6\x03\x78\x04\x85\x39\x31\x98\xbc\x86\x8b\xd1\x6c\xf8\x1a\xc6\x8b\x38\x1c\x4f\xe6\x49\xfc\xa7\x07\x00\x90\x92\xe3\x56\x18\x5f\x69\xf6\x47\x9a\x6f\xc8\x82\xa5\xb5\x70\xde\xee\x61\x97\x91\x25\x48\xc9\x48\xbd\xa7\x14\x44\x8e\x6b\x02\xe1\x00\xb7\x28\x24\x2e\x25\xf5\x2b\x11\x4b\xff\x0a\x61\x29\x65\xe0\x6d\x41\x55\x68\x8b\xb2\x20\x06\x69\x25\x38\x10\xfa\x1a\x4a\x34\x9b\x4f\x92\x59\x27\x4c\x4c\x46\x3b\xe1\xb5\xdd\x83\x50\xb0\xcb\x04\xcf\xc0\x67\xd4\x50\x70\x54\xb0\x24\x58\xe9\x42\xa5\xf7\x28\xd0\x49\xdc\x6c\x54\xe1\x1d\xa9\x16\xc8\xe4\x3d\x18\x87\x8b\x69\xf0\x1e\x76\x10\x24\x19\x55\xa9\xa0\x57\x67\xce\x3b\xe1\x33\x10\x1b\x82\xa5\x50\x68\xf7\xf7\xcc\x85\xf3\x42\x3f\xee\xb4\xdd\x38\x83\x9c\x3a\xfd\x93\x60\x7c\xc5\xde\xe3\xfa\xbb\xbb\xd7\x65\xd9\x85\xa3\xbb\x55\x97\xdd\xe1\xfc\x99\x5f\x12\xbe\x85\x51\x1c\xce\xc3\xe9\x30\x5c\x7c\x84\xf1\x7c\x32\x9b\x76\xd8\xfe\x52\x29\x59\xb9\x17\x6a\x0d\x09\x49\x32\x96\x1c\x29\x4e\xb0\xad\x5b\xaf\xf2\x2e\x31\xea\xb6\x00\xe1\x5d\x8d\xe6\xca\x76\x51\xe0\x76\x68\x4c\x79\xb8\xde\xcf\x49\x95\x09\xaa\x2a\x81\xcb\xc2\x79\xb2\xf7\xc8\xfb\xcf\x83\x97\xe7\xe7\x7e\x4f\x2f\xff\x12\xf7\x4d\x1f\xd7\xb3\x30\xfa\x12\xad\x0e\x9c\x4f\x04\x1a\xe3\xca\x01\x28\xe3\x39\x79\x4c\xd1\x23\xab\xbe\xa0\xb9\x80\xcb\x5f\x01\xe0\x0c\xf1\x63\x96\x25\x23\x05\x47\xc7\xe0\xa5\x89\x38\x92\xc4\xbd\xb6\xc7\x0c\x80\x1c\x3d\xcf\xde\x70\x49\xd2\x9d\x82\xb7\x0c\xe0\x6b\x54\xcf\x44\x5a\x78\xe5\x92\x17\x9a\xb7\x55\xbf\xa3\xd7\xb0\x76\x2b\x38\x05\x9c\xeb\x42\xf9\xe9\xcd\xb3\x00\x5c\x2b\x8f\x42\x35\xef\xc4\x69\x3d\xde\x71\xad\x57\xf5\xc7\x19\x3c\x7c\x76\x3c\x2c\x87\xa7\x56\xf8\x38\xe4\xc7\x8d\xd3\xd0\x1d\xd8\x79\x24\x09\xc6\x87\x96\x0f\xd7\x79\x8e\x2a\x65\xad\x70\x89\x29\x36\x6d\x28\xb4\x6b\xd7\x95\x59\x5e\x4c\x67\x01\x51\x21\x65\xa4\xa5\xe0\x7b\x06\x81\xdc\xe1\xde\xb5\xb2\x48\x6d\xbb\x04\xeb\x1b\xfa\x1d\x24\xc3\x9f\x55\x19\xf3\x28\x18\x86\x17\x79\xa7\x7e\xee\x5f\xd5\x88\x66\xa3\xd3\xeb\xd3\x71\xf8\x87\xd5\xf9\x25\x41\xb9\x56\x82\x64\x1a\xd3\xaa\x7b\xb7\xd9\x8f\xd0\x67\xec\xab\xdd\x06\xa5\xe7\x55\x94\x59\x14\xc6\x41\x32\x8b\x6f\xf2\x30\xe8\xb7\x1a\xe3\x7a\x6d\x57\xdf\x9b\x6e\xdd\x87\xcf\xae\x03\x87\x7e\xef\x7f\x00\x00\x00\xff\xff\x07\x79\xe3\xa2\x06\x07\x00\x00")

func deployIstioWorkspaceOperatorYamlBytes() ([]byte, error) {
	return bindataRead(
		_deployIstioWorkspaceOperatorYaml,
		"deploy/istio-workspace/operator.yaml",
	)
}

func deployIstioWorkspaceOperatorYaml() (*asset, error) {
	bytes, err := deployIstioWorkspaceOperatorYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "deploy/istio-workspace/operator.yaml", size: 1798, mode: os.FileMode(436), modTime: time.Unix(1560431092, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _deployIstioWorkspaceRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x52\x4f\x6b\xfb\x30\x0c\xbd\xe7\x53\x98\x5e\x0a\x3f\xa8\x7f\xec\x36\x72\xdd\x61\xf7\x31\x76\x57\x1d\xb5\x15\xb5\x2d\x23\xc9\x19\xdb\xa7\x1f\x49\xc3\xea\x42\xdb\xd1\x9c\xe2\x17\x3f\xbd\x3f\x11\x14\xfa\x40\x51\xe2\xdc\x3b\xd9\x42\xf0\x50\xed\xc0\x42\xdf\x60\xc4\xd9\x1f\x9f\xd5\x13\xff\x1f\x9f\xba\x23\xe5\xa1\x77\x2f\xb1\xaa\xa1\xbc\x71\xc4\x2e\xa1\xc1\x00\x06\x7d\xe7\x5c\x10\x9c\x09\xef\x94\x50\x0d\x52\xe9\x5d\xae\x31\x76\xce\x65\x48\xd8\x3b\x52\x23\xde\x7c\xb2\x1c\xb5\x40\xc0\x4e\x6a\x44\x9d\x88\x1b\x07\x85\x5e\x85\x6b\x99\x8f\xd3\xb3\x71\xab\xd5\xfc\x2a\xa8\x5c\x25\x60\xf3\xa5\xf0\xa0\xbf\x07\x45\x19\x29\xe0\x19\xc0\x3c\x14\xa6\x6c\x67\xa4\x4c\xd9\xd4\x30\xdb\xc8\xb1\x26\x0c\x11\x28\x35\x84\x11\xdb\xdb\x81\xf3\x8e\xf6\x09\x4a\xab\x11\x04\x97\x2b\x23\xca\xb6\xf1\xb2\xfe\xb7\x7e\x3c\xc0\x54\xc7\x5c\xc1\xd5\x91\x7b\xb4\x5b\x23\xa1\x2c\xae\xae\x0c\x1d\xb0\x44\xfe\x4a\x17\x59\x06\xc0\xc4\x59\xb1\x81\x04\x4b\xa4\x00\x17\x98\x1a\x18\xee\x6a\xd4\xc7\x43\x4e\x8e\x3c\x17\xcc\x7a\xa0\x9d\x79\xe2\xbf\xed\x9d\x0a\x7e\x54\x28\x71\x26\x63\xa1\xbc\xf7\x81\x05\x59\x7d\xe0\x74\x4b\x6c\x59\x8a\x85\x73\xa7\xe5\xe5\x97\x4f\x8b\x8b\xb7\x94\xe7\xb5\x6d\x32\xde\xd1\x3d\xf9\xbf\x1a\xeb\x27\x00\x00\xff\xff\xde\xf7\xfc\x19\x64\x03\x00\x00")

func deployIstioWorkspaceRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_deployIstioWorkspaceRoleYaml,
		"deploy/istio-workspace/role.yaml",
	)
}

func deployIstioWorkspaceRoleYaml() (*asset, error) {
	bytes, err := deployIstioWorkspaceRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "deploy/istio-workspace/role.yaml", size: 868, mode: os.FileMode(436), modTime: time.Unix(1560416801, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _deployIstioWorkspaceRole_bindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x90\x41\x6f\xdb\x30\x0c\x85\xef\xfe\x15\x44\xb0\x6b\x3c\xec\x36\xe8\x96\x05\xc1\x4e\x1b\x86\x24\xd8\x9d\x96\xe8\x98\x8b\x2d\x6a\x14\x95\x00\x2d\xfa\xdf\x8b\x28\x4e\x5b\xb4\x40\xaa\x9b\xf8\x88\xc7\xef\xbd\x23\xc7\xe0\x60\x4f\x53\x1a\xd1\xa8\xc1\xc4\x7f\x49\x33\x4b\x74\x60\xf3\xb0\x95\x44\x31\x0f\xdc\x5b\xcb\xf2\xf5\xf4\xad\x49\xa8\x38\x91\x91\x66\xd7\x00\x2c\x21\xe2\x44\x0e\x7e\xaf\x7e\x6d\x76\x7f\x56\xeb\x4d\x03\x00\x10\x28\x7b\xe5\x64\xd5\x69\xb1\x47\x3d\x90\xd5\xc5\x9c\xd0\x13\xf4\xa2\x70\x1e\xd8\x0f\xa0\x32\x12\x74\x1c\x03\xc7\x03\xe4\x41\xca\x18\xa0\x23\x08\xd4\x73\xa4\xb0\xa8\x66\x4a\xff\x0b\x2b\x05\x07\xa6\x85\xea\xe8\x84\x63\x21\x07\x9c\x8d\x65\x79\x16\x3d\x56\xdf\xa5\x24\x52\x34\xd1\x46\xba\x7f\xe4\x6d\x06\xbc\x86\x5c\x8f\x25\x1b\xe9\x56\x46\xfa\x71\xbd\x57\x9d\xde\x46\xd6\x0e\x7d\x8b\xc5\x06\x51\x7e\xc0\x0b\x7c\x7b\xfc\x9e\xe7\xd8\x97\xe5\x89\x0c\x03\x1a\xba\xfa\x83\x39\xfa\x3b\x8a\xaa\xe5\xf2\x4a\x70\x79\x37\x8a\x1d\xe9\x89\x3d\xad\xbc\x97\x12\x6d\x16\xef\x19\xdd\xd4\x3a\x71\xf0\xe5\xf1\xa5\xe8\xa7\x6b\x39\x32\xd2\x96\xfa\xdb\x9d\x0f\x59\x3f\x25\xad\x15\xfc\x54\x29\xe9\x4e\x01\xcd\x73\x00\x00\x00\xff\xff\x5a\x3b\x0c\xba\x29\x02\x00\x00")

func deployIstioWorkspaceRole_bindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_deployIstioWorkspaceRole_bindingYaml,
		"deploy/istio-workspace/role_binding.yaml",
	)
}

func deployIstioWorkspaceRole_bindingYaml() (*asset, error) {
	bytes, err := deployIstioWorkspaceRole_bindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "deploy/istio-workspace/role_binding.yaml", size: 553, mode: os.FileMode(436), modTime: time.Unix(1560431092, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _deployIstioWorkspaceService_accountYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\x31\x0e\x80\x30\x08\x00\xc0\x9d\x57\xf0\x01\x07\x57\x36\xdf\x60\xe2\x4e\x28\x03\x69\x0a\x4d\xc1\xfa\x7d\x8f\xa7\x3d\xba\xd2\xc2\x09\xf7\x09\xdd\xbc\x11\xde\xba\xb6\x89\x5e\x22\xf1\x7a\xc1\xd0\xe2\xc6\xc5\x04\x88\xce\x43\x09\x2d\xcb\xe2\xf8\x62\xf5\x9c\x2c\x0a\x7f\x00\x00\x00\xff\xff\x94\xa0\xb7\x3f\x46\x00\x00\x00")

func deployIstioWorkspaceService_accountYamlBytes() ([]byte, error) {
	return bindataRead(
		_deployIstioWorkspaceService_accountYaml,
		"deploy/istio-workspace/service_account.yaml",
	)
}

func deployIstioWorkspaceService_accountYaml() (*asset, error) {
	bytes, err := deployIstioWorkspaceService_accountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "deploy/istio-workspace/service_account.yaml", size: 70, mode: os.FileMode(436), modTime: time.Unix(1554454542, 0)}
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
	"deploy/istio-workspace/crds/istio_v1alpha1_session_crd.yaml": deployIstioWorkspaceCrdsIstio_v1alpha1_session_crdYaml,
	"deploy/istio-workspace/operator.yaml":                        deployIstioWorkspaceOperatorYaml,
	"deploy/istio-workspace/role.yaml":                            deployIstioWorkspaceRoleYaml,
	"deploy/istio-workspace/role_binding.yaml":                    deployIstioWorkspaceRole_bindingYaml,
	"deploy/istio-workspace/service_account.yaml":                 deployIstioWorkspaceService_accountYaml,
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
	"deploy": &bintree{nil, map[string]*bintree{
		"istio-workspace": &bintree{nil, map[string]*bintree{
			"crds": &bintree{nil, map[string]*bintree{
				"istio_v1alpha1_session_crd.yaml": &bintree{deployIstioWorkspaceCrdsIstio_v1alpha1_session_crdYaml, map[string]*bintree{}},
			}},
			"operator.yaml":        &bintree{deployIstioWorkspaceOperatorYaml, map[string]*bintree{}},
			"role.yaml":            &bintree{deployIstioWorkspaceRoleYaml, map[string]*bintree{}},
			"role_binding.yaml":    &bintree{deployIstioWorkspaceRole_bindingYaml, map[string]*bintree{}},
			"service_account.yaml": &bintree{deployIstioWorkspaceService_accountYaml, map[string]*bintree{}},
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