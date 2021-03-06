// Code generated by go-bindata. DO NOT EDIT.
// sources:
// assets/dns/cluster-role-binding.yaml (223B)
// assets/dns/cluster-role.yaml (397B)
// assets/dns/daemonset.yaml (6.121kB)
// assets/dns/metrics/cluster-role-binding.yaml (279B)
// assets/dns/metrics/cluster-role.yaml (246B)
// assets/dns/metrics/role-binding.yaml (293B)
// assets/dns/metrics/role.yaml (284B)
// assets/dns/namespace.yaml (369B)
// assets/dns/service-account.yaml (85B)
// assets/dns/service.yaml (520B)

package manifests

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _assetsDnsClusterRoleBindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xce\x31\x8e\x83\x40\x0c\x05\xd0\x7e\x4e\xe1\x0b\xc0\x6a\xbb\xd5\x74\x9b\xdc\x80\x48\xe9\xcd\x8c\x09\x0e\x60\xa3\xb1\x87\x22\xa7\x8f\x10\x4a\x45\x3a\x17\xfe\xff\xfd\x89\x25\x47\xb8\xce\xd5\x9c\x4a\xa7\x33\x5d\x58\x32\xcb\x23\xe0\xca\x77\x2a\xc6\x2a\x11\x4a\x8f\xa9\xc5\xea\xa3\x16\x7e\xa1\xb3\x4a\x3b\xfd\x59\xcb\xfa\xb3\xfd\x86\x85\x1c\x33\x3a\xc6\x00\x00\x20\xb8\x50\x04\x5d\x49\x6c\xe4\xc1\x9b\x2c\x16\xac\xf6\x4f\x4a\x6e\x31\x34\x70\x78\x37\x2a\x1b\x27\xfa\x4f\x49\xab\x78\xf8\xc4\xf6\xe7\xe3\xb6\x15\xd3\xa9\xa7\xe8\x4c\x1d\x0d\x3b\x74\x9a\x1d\xbe\xd3\xef\x00\x00\x00\xff\xff\xfa\x62\xe7\x50\xdf\x00\x00\x00")

func assetsDnsClusterRoleBindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsClusterRoleBindingYaml,
		"assets/dns/cluster-role-binding.yaml",
	)
}

func assetsDnsClusterRoleBindingYaml() (*asset, error) {
	bytes, err := assetsDnsClusterRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/cluster-role-binding.yaml", size: 223, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd9, 0xf6, 0x2a, 0x3b, 0x84, 0xd7, 0x3e, 0xc4, 0xe1, 0x70, 0x66, 0x31, 0xda, 0xc4, 0x2f, 0x53, 0x27, 0x29, 0x13, 0xfe, 0x80, 0x36, 0xc5, 0xa1, 0x70, 0xdc, 0x2d, 0xef, 0xcf, 0xe0, 0xc4, 0xeb}}
	return a, nil
}

var _assetsDnsClusterRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x90\xb1\x4e\xc4\x30\x10\x44\x7b\x7f\x85\x75\xfd\x05\xd1\xa1\xb4\x14\xf4\x14\xf4\x1b\x67\x50\x96\xe4\x76\xad\xdd\x75\x4e\xe2\xeb\x51\x2e\x57\xa0\x8b\xa0\x9b\x19\xd9\xf3\x3c\x9e\x59\xc6\x3e\xbf\x2e\xcd\x03\xf6\xae\x0b\x12\x55\xfe\x80\x39\xab\xf4\xd9\x06\x2a\x1d\xb5\x98\xd4\xf8\x9b\x82\x55\xba\xf9\xc5\x3b\xd6\xa7\xf5\x39\x5d\x10\x34\x52\x50\x9f\x72\x16\xba\xa0\xcf\x5a\x21\x3e\xf1\x67\x9c\x47\xf1\x64\x6d\x81\xf7\xe9\x9c\xa9\xf2\x9b\x69\xab\xbe\x9d\x3c\xe7\xd3\x29\xe5\x6c\x70\x6d\x56\x70\xcf\x20\x63\x55\x96\xf0\x9b\x73\xd8\xca\x05\xbb\xa9\x3a\xee\x62\x63\x78\xa5\x3d\x5f\x61\xc3\xfd\xee\xc2\x1e\x37\x71\xa5\x28\x53\x3a\x02\xb7\x01\x90\xe0\xf2\x7b\xc1\xf1\x0d\xa1\x33\xc4\xb0\x32\xae\x0f\x84\x62\xa0\xc0\x1f\xcd\x8f\x5f\x73\x2c\xf6\x36\x7c\xa1\x04\x95\x02\xf7\xff\x00\x3f\x01\x00\x00\xff\xff\x76\x1b\x55\x2e\x8d\x01\x00\x00")

func assetsDnsClusterRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsClusterRoleYaml,
		"assets/dns/cluster-role.yaml",
	)
}

func assetsDnsClusterRoleYaml() (*asset, error) {
	bytes, err := assetsDnsClusterRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/cluster-role.yaml", size: 397, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x84, 0xae, 0xd1, 0xba, 0xfa, 0x6b, 0xf8, 0x6e, 0x8d, 0x28, 0xc2, 0xa7, 0xaf, 0xc9, 0x3b, 0xc7, 0xcd, 0x80, 0xbe, 0xec, 0x98, 0xb4, 0x61, 0xa0, 0x9, 0xae, 0xa, 0xd8, 0xb2, 0x2e, 0x16, 0xf2}}
	return a, nil
}

var _assetsDnsDaemonsetYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\xdd\x53\x1b\x39\x12\x7f\xe7\xaf\xe8\x1d\xa8\x90\xec\x32\xd8\x24\x21\x9b\x9b\x84\xbd\xf5\x1a\x13\xa8\x8d\xc1\x85\x9d\xcd\x03\x45\xb9\x64\x4d\xdb\xa3\xb3\x46\xd2\x4a\x9a\x81\x29\xf0\xff\x7e\xa5\x19\x7f\xcc\x87\x61\x37\xb9\xda\xaa\xe3\xc1\xd8\xea\xee\x9f\xba\x5b\xfd\x25\xcd\x99\x08\x03\x38\x25\x18\x4b\x31\x44\xbb\x43\x14\xfb\x03\xb5\x61\x52\x04\x40\x94\x32\xad\xf4\x68\x67\x17\x04\x89\xf1\x20\xff\x34\x8a\x50\x04\x22\x42\xe0\x64\x82\xdc\x00\xd1\x08\x06\x2d\x10\x0b\x3a\x11\x96\xc5\xb8\x63\x14\xd2\x60\x07\xc0\x62\xac\x38\xb1\xe8\xbe\x03\xac\x56\xf3\xef\xa8\x53\x46\xb1\x43\xa9\x4c\x84\xbd\x24\x31\x06\x10\x0a\xb3\xa4\x2a\xcd\xa4\x66\x36\xeb\x72\x62\x4c\x41\x34\x99\xb1\x18\xfb\x42\x86\xe8\x53\xcd\x2c\xa3\x84\x2f\xb9\xa9\x14\x96\x30\x81\xda\xac\xd0\xfd\x5c\xd3\x32\x22\xc0\x2e\xb0\x98\xcc\x10\x98\xa9\x6b\xbb\xe2\xc8\xe9\x83\x84\xf3\x81\xe4\x8c\x66\x01\x5c\x4c\x2f\xa5\x1d\x68\x34\x28\xec\x9a\xcb\xa2\x8e\x99\x20\x96\x49\xd1\x47\x63\x9c\xc8\x92\xfd\x8c\x70\x3e\x21\x74\x3e\x92\x9f\xe5\xcc\x5c\x89\x9e\xd6\x52\xaf\xe5\xa8\x8c\x63\xe2\x5c\x7d\x03\x1e\x95\x1a\x43\x61\x3c\xb8\x5d\x93\x89\x9e\x99\x9c\xe6\x53\x29\xa6\xde\x01\x78\x2d\xb4\xb4\xb5\xe4\x6c\x75\xa5\xc6\x29\xe3\x58\x16\x49\x25\x4f\x62\xec\x3b\x07\xae\x2d\xdf\xd8\xee\x60\xd8\xcc\x2f\x98\xd6\x54\x80\xd8\xf1\x0f\x88\x8d\x02\x28\xef\x50\xe2\xd0\x48\xc2\x2b\xc1\xb3\x00\xac\x4e\x36\xa2\x4a\xea\xea\x3e\x6b\xbf\x0f\xa4\xb6\x01\x1c\xbf\x39\x7e\x53\x42\x69\x9e\x80\x3b\x57\x69\x25\x95\x3c\x80\x2f\xa7\x83\x6f\x47\xf2\x2d\x55\x5b\xd1\x46\xdd\x0d\x9a\xd3\x9e\x09\x34\x66\xa0\xe5\x04\x83\x12\x7f\x64\xad\xfa\x84\xb6\xbc\x04\xa0\x0a\x4f\x38\xa9\xac\x4a\xc8\x55\x79\x7f\xf4\xfe\xa8\xb2\x6c\x68\x84\x4e\x9d\xf3\xd1\x68\x50\x22\x30\xc1\x2c\x23\xfc\x14\x39\xc9\x86\x48\xa5\x08\x4d\x00\x47\xed\xb2\xb6\xa8\x99\x0c\xb7\xd3\x4c\x42\x29\x1a\x33\x8a\x34\x9a\x48\xf2\x30\x80\xf2\xa6\x53\xc2\x78\xa2\xb1\x44\x2d\x7b\xc7\x85\xb0\x4c\xec\x36\x60\xce\x52\xfc\x66\x4f\x44\x48\xb8\x8d\xb6\xb9\xa2\xfd\xbe\xfd\xdd\xae\x78\xd7\x7e\x46\xe5\xe3\xff\xc1\x15\xc7\xa5\x83\x37\x32\xd1\x14\x4d\x50\x89\xe5\x3f\x13\x34\xd6\x54\x4d\xa5\x2a\x09\xe0\xb8\x1d\x57\x16\x63\x8c\xa5\xce\x02\xf8\xb9\xdd\x67\xb5\x3a\x32\x4f\x26\xe8\xeb\x09\xa1\xbe\xd2\xf2\x3e\xfb\x86\x9a\x92\xa7\x75\x29\xd2\x7d\x9f\xcb\x99\x95\xc6\x86\xa8\x75\x65\xdd\x20\x4d\x34\xfa\x9c\x19\x8b\xc2\x27\x61\xa8\xd1\x98\x93\xe0\x5f\x47\xc7\x6f\x2b\x7c\x96\x1b\x9f\x32\x15\xa1\xf6\x4d\xc2\x2c\x9a\x93\xd1\xe7\xe1\xb8\xd7\x3d\x3d\xef\x8d\xaf\x87\x9d\xf1\xd7\x8b\xd1\xf9\xb8\xd3\x1b\x8e\x8f\x5e\xbf\x1f\x7f\xea\xf6\xc7\xc3\xf3\xce\xeb\xe3\x77\x07\x1b\xae\x5e\xf7\xf4\x2f\xf8\x1a\x38\xdd\xdf\xba\x7f\x0b\x67\x2b\xdf\x33\x68\x15\xcb\x12\x65\xac\x46\x12\x9f\xb8\xf0\x0c\x5a\xad\xa3\xd7\x3f\x1f\xb6\x0f\xdb\x87\x47\xce\x09\x6f\x5a\x4d\x2f\xa0\xb6\xbe\x2b\x8a\x27\x79\x21\xb3\xdc\xb4\x94\x66\x29\xb1\xe8\xbe\x1f\x52\x6d\x1b\x22\x4b\xba\x3f\xc7\xec\x19\xc9\x39\x66\x7f\xbb\xea\x55\xce\x67\x55\xab\x62\xb4\x9a\x51\xf3\xdd\xa1\x79\xf4\x44\x68\xbe\xdd\x84\xe6\xd3\xe5\xbf\x5e\xe0\x4b\xd6\x3d\xa5\xa8\xf3\xcd\x5f\x35\x80\x52\x4f\x2d\xba\xb0\x33\x8a\xa7\xa8\xff\x6f\x3a\x6c\x9e\x41\x6e\x6a\x90\xc2\xe2\x7d\xa5\xba\x39\xfb\x19\xc7\x19\x86\xb5\xa6\xf6\x7c\x0f\x8d\xa4\xb1\x26\x0f\x94\x67\x1a\x68\xce\x54\x72\x02\x8a\x14\x2e\x3b\xfd\xde\xb0\x77\xfd\x47\xef\x3a\x9f\x94\xba\x9f\xbf\x0c\x47\xbd\xeb\xf1\xe9\x55\xbf\x73\x71\xb9\x6d\x62\x5a\x89\xa3\x48\x9b\x6a\x38\xa4\x8b\x6e\x6f\x58\x52\x62\x17\xba\x6e\x9e\x00\xa9\xa1\x18\xc8\x0c\x2a\xa2\x89\xc5\x10\x5c\x05\x01\x39\x5d\x8d\x58\xa6\x22\x75\x79\x35\xea\x05\x70\x26\x35\x08\x79\x77\x00\x28\x4c\xa2\x11\x6c\x84\x06\x73\xb5\x34\x72\x62\x59\x8a\xc5\xa8\xf7\x01\xa6\x52\x03\x12\x1a\x55\x09\x07\x15\x4c\x22\x80\x70\x46\x0c\xdc\x31\x1b\x39\xac\xba\xbd\x26\x99\x4e\xd9\x3d\xdc\x31\xce\x81\x70\x23\x61\x82\x40\xc2\x10\xc3\xc3\x12\x4e\x4a\x78\x82\x01\x78\x79\x8c\xf8\x1a\x67\xcc\x58\x9d\x1d\x4a\x85\xc2\x44\x6c\x6a\xfd\x1a\xc1\xa4\xd4\x6b\x0c\x57\x25\xd7\xb5\x26\x4c\xb4\x26\xc4\x44\xe5\x22\x40\x4b\x3f\x1e\xcb\x46\xfc\xd0\x64\x87\xfc\x8c\xfc\x44\x82\x62\x0a\x5d\xe7\xd9\x29\xf7\x30\x4d\x14\xec\xff\x47\x4e\x0c\xf8\x0a\x1e\xe1\xde\x55\x7a\x98\x3b\x13\x1f\x1f\xf3\x18\xfb\x00\x77\x84\xd9\x0f\x80\xf7\xcc\x42\x7b\x1f\x46\xbd\xeb\x7e\x19\xe1\x6a\xd0\xbb\x1c\x9e\x5f\x9c\x8d\xc6\xfd\xce\xf5\xef\xbd\xeb\x13\x6f\x63\xeb\x0c\x05\xe6\xa7\x59\x4d\x35\xaf\x24\x7e\x7e\x35\x1c\x0d\xc7\x67\x17\x9f\x7b\x27\xde\x26\x0e\xcb\x1c\xa3\x5e\x7f\xd0\x60\x38\xb4\xb1\xf2\xca\x6a\x5c\x9c\x0d\x4f\xf6\x0f\x60\x3f\xcf\x7a\xf0\x35\xf8\x64\x1d\x3a\xf0\xf1\xe3\x47\xf0\xf6\x1e\x56\x01\xb8\xa8\x48\xee\x42\x9f\xcc\x11\x48\x3e\xe6\x4b\x4d\x74\x06\x2e\x55\x36\x61\x20\x79\x58\xa4\x50\xbe\xbe\x6f\x80\x58\xab\xd9\x24\xb1\x68\xca\x27\x4f\x15\xf8\x53\xf0\xfd\x0d\xd5\x97\x82\x67\x6e\xe3\x8d\x91\x0b\xcf\xfd\x5e\x9b\x54\xd5\xe4\x2e\x72\xfb\x16\x4e\x0f\x65\xa5\x74\x86\x48\xb9\x0b\x6c\xbf\x03\x26\xa5\x63\xa6\x4c\x85\xec\xe2\xdb\xa4\x14\x98\x70\xf0\x2b\xbb\x6f\x7e\xbd\x5d\x78\x0d\x28\x67\xf1\x19\x5a\x1a\xad\xfc\x03\x17\x03\x98\x6a\x19\x03\xe5\x89\xb1\xa8\x5d\x6d\x04\x36\x05\x55\x14\xb4\x43\xf8\x8a\x10\x3b\x17\x19\x4c\x51\x13\x0e\x56\x33\x34\x0d\x4c\x2b\x21\x94\xc0\x6c\x00\x17\x83\xf4\xed\x81\xfb\x7c\x97\x7f\xbe\x05\x99\xa2\x76\xd3\x6d\x5e\x45\xdc\xfa\x7a\xe5\x10\x46\x11\x82\xbd\x93\xc0\x89\xcb\x77\xb1\x05\xd8\xd9\xed\x0c\x0c\x51\x71\x99\xc5\x28\xec\x32\x47\x7f\x4f\x74\xa6\x41\x0a\x77\x42\xa8\xe1\x4a\xa1\x18\x5a\x42\xe7\xf0\xf2\x6a\x38\x38\x7a\xf3\x0a\x7c\xb0\x91\x34\xe8\xf4\x12\xd2\x36\x80\x4d\xa2\x5c\x5f\x74\x53\x3c\x70\x49\xc2\x09\xe1\x44\x50\xd4\x26\xd7\xd3\x35\x36\x96\xd7\x12\x42\x23\x26\x66\x70\x7a\x39\x04\x1b\x69\x99\xcc\xa2\x5c\xf5\x1a\x1e\x8d\x43\x73\xf2\x72\x3f\x64\x33\xf0\x2d\x74\xe0\x57\x6f\xef\x61\x53\x40\x17\x1e\xfc\x64\x22\xb7\x9b\x3b\xa0\x94\x2e\x0e\xf7\x1e\xaa\xf5\x65\xe1\xed\xd7\x10\x8b\xbf\x35\x62\xa7\xf3\x0f\x80\xc2\x4f\x96\xaa\x7f\x46\xd7\xef\x45\x7e\x55\x83\x76\x67\xcf\x5c\x68\xef\x3d\xfc\xe0\x9c\x7c\xf3\xe3\xed\xa2\xc6\xd2\x08\x71\x00\xa6\xcc\xc9\xcb\xbd\x97\x98\x12\xee\x36\xcb\x05\xd9\xed\xc2\x7b\x55\x87\x07\x17\xeb\x37\x37\xe0\xed\xfd\xdb\x03\x1f\xff\x84\x36\xbc\x78\xe1\x44\x76\x99\x2a\x52\x08\x7c\x81\xd0\x86\xdb\xdb\x0f\xae\x1e\x88\x2d\x96\x2f\x73\xf2\x66\x69\x95\x77\x7b\xe2\xed\x3d\xac\xc4\xb7\xf0\x4f\x34\x92\x79\x63\x7d\xca\x1a\x66\x09\xdc\x69\x2c\x54\x56\x76\xe1\x8b\x0a\x89\xc5\x52\x13\x87\xbc\xec\xb0\x29\xdc\x21\xcc\xd0\xba\x96\xc4\xc2\x52\xb2\x9b\x1a\xc0\x57\x2c\x7a\x9a\x90\x16\x92\x06\xd8\x5d\x84\xc2\x99\xad\xf3\x89\x68\x79\xcd\x5e\xa3\xc9\xc4\xba\x59\x49\x6a\x20\x8a\x41\x22\x48\x4a\x18\x27\x13\xc6\x99\xcd\x6a\xdb\x0c\x2d\xe1\x08\x28\xf2\xea\x01\x54\x26\x3c\x74\x4d\xc5\x58\x77\xb4\xa5\x0d\xd9\x34\xaf\xba\xab\x1d\x98\x81\x10\x39\x5a\x0c\x77\x9a\x67\xe6\x8b\x65\x20\xe5\xde\xff\xf1\xd6\x5f\x78\x4f\x1d\xd3\x2e\xfc\x96\x30\x1e\x02\x01\x81\x77\xa5\x7a\x5e\x94\xbe\xb2\xc1\xae\xb4\xc8\x44\x03\x4d\x8c\x95\xf1\x5a\xe3\x29\xe3\x16\x35\x86\xce\xe6\x1a\xf6\x4c\xa3\x02\x3f\x05\x6f\x17\xf6\x1e\xea\x0d\xb1\x28\xf9\x95\x16\xf0\xcb\x33\x4d\xa0\xd0\xb5\xa3\x14\xe6\x35\xa8\xe8\x98\x1b\x25\x5c\xa1\x6f\x4e\x44\xd0\xe8\x01\x3f\xac\x9c\xf2\x44\x0f\x58\xa6\x95\x2a\xf2\x6a\xc5\x5c\x84\xef\xed\x62\xab\x00\x00\xd2\x48\x42\x1e\xd9\x8b\x42\x68\xf5\xaf\x99\xc6\xf0\x84\x2b\x7e\x69\xd8\x5e\xdf\xa4\x11\xf4\xdb\xc2\xde\xf9\x68\x74\x75\x7a\x15\x6c\x09\x7f\x62\x65\xcc\x28\xe1\x3c\x73\x3d\x89\xa4\x92\x85\x40\x44\x06\x4c\x50\x29\x4c\x7e\x31\xb5\x30\xc1\x88\xa4\xac\x34\x76\xaf\x50\xaf\x51\x71\x37\x89\x6e\x8b\x88\x58\x86\x6c\xca\x30\x84\xb4\x78\x5a\x74\x51\x28\x10\xc3\x5a\x6c\xba\x5e\xa0\x6a\x66\x36\x62\xe0\xf1\x71\x39\x31\x3c\xcf\xd7\xb4\x7a\xc5\xeb\x32\xc3\xa5\xac\xc6\x58\xa6\x18\x6e\x6c\xcd\xa3\x9a\x6a\x74\xf7\xc0\x22\x75\xf2\x7e\xb6\x99\x4b\x80\x4a\x95\x01\x8d\x12\x5d\x4d\x92\x5a\xfd\x31\x1c\x51\xc1\xbb\x36\xbc\xc8\x47\xc0\x0a\x2d\x11\x6e\xaa\x6c\x8e\x22\x95\xc3\xfb\xe6\xa7\x8c\xd5\x75\x31\x14\x66\x75\x57\x3a\xc5\x29\x49\xf8\x6a\x73\x37\x46\x0e\x91\x23\xb5\x52\x6f\x00\xe6\xc9\x04\xb5\x40\x37\x8f\x31\xd9\x92\x26\x00\xce\x44\x72\x5f\x10\x97\x5c\xa5\xdb\xd8\x27\x4d\x28\x0e\x6a\xef\x58\xaf\x57\x2f\x3b\xc5\x55\xaa\xf1\x06\xbb\xfd\x1d\xb2\x58\xed\x13\x15\x94\x6e\x4e\x97\x24\x7e\xee\xf6\x08\xc0\x2c\xc6\x15\xfb\x7d\x98\x63\x16\xc0\xea\x75\x74\xcb\x7b\x56\x8d\xf4\xcc\xcd\xce\x2d\xe5\xd7\xba\x9d\x3a\xc6\x96\x6b\x1e\x80\xcd\x14\x06\x70\xd6\x84\xde\x76\xa7\xde\x75\x97\x53\x8d\xf6\x59\x0b\xad\xe4\x6e\xe8\x67\x52\xac\x6d\xdc\xcd\x67\x27\x97\x29\xc6\x85\xa9\x4e\x04\xb8\x51\x32\xbb\x73\x6d\xe5\x10\x46\x85\x04\x02\xe1\x1c\x2c\x61\x62\xad\xa1\x0f\x52\x39\x92\xd4\x01\xf4\x5c\xaf\x30\x3b\xff\x0d\x00\x00\xff\xff\x30\x68\x23\x0f\xe9\x17\x00\x00")

func assetsDnsDaemonsetYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsDaemonsetYaml,
		"assets/dns/daemonset.yaml",
	)
}

func assetsDnsDaemonsetYaml() (*asset, error) {
	bytes, err := assetsDnsDaemonsetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/daemonset.yaml", size: 6121, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x49, 0xad, 0x77, 0xc3, 0x63, 0x19, 0x9b, 0xdf, 0x2f, 0xf9, 0xf7, 0xc2, 0xd2, 0xa, 0xaa, 0xad, 0x42, 0xfd, 0x34, 0x89, 0xf, 0xab, 0x44, 0x5f, 0xab, 0x6f, 0xeb, 0x9d, 0x7e, 0x2e, 0x9, 0x7b}}
	return a, nil
}

var _assetsDnsMetricsClusterRoleBindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8f\xb1\x4a\x04\x41\x0c\x86\xfb\x79\x8a\xbc\xc0\xae\xd8\x1d\xd3\xa9\x85\xfd\x09\xf6\xb9\x99\x9c\x1b\x77\x27\x19\x92\xcc\x16\x3e\xbd\x2c\x8a\x08\xe2\xb5\x81\x7c\xdf\xff\xad\x2c\x35\xc3\xd3\x36\x3c\xc8\xce\xba\xd1\x23\x4b\x65\x79\x4b\xd8\xf9\x95\xcc\x59\x25\x83\x5d\xb0\xcc\x38\x62\x51\xe3\x0f\x0c\x56\x99\xd7\x93\xcf\xac\x77\xfb\x7d\x6a\x14\x58\x31\x30\x27\x00\xc1\x46\x19\xaa\xf8\xd4\x54\x38\xd4\x0e\x92\x8f\xcb\x3b\x95\xf0\x9c\x26\xf8\xd2\xbd\x90\xed\x5c\xe8\xa1\x14\x1d\x12\x3f\x7f\xdd\xb4\x51\x2c\x34\x7c\x5a\x4f\xfe\x7d\xf6\x8e\x85\x32\x68\x27\xf1\x85\xaf\xf1\x9b\x6c\xba\xd1\x99\xae\x87\xf9\x4f\xc7\x7f\x6b\x00\xb0\xf3\xb3\xe9\xe8\x37\xba\xd2\x67\x00\x00\x00\xff\xff\x5b\x52\x00\xaa\x17\x01\x00\x00")

func assetsDnsMetricsClusterRoleBindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsMetricsClusterRoleBindingYaml,
		"assets/dns/metrics/cluster-role-binding.yaml",
	)
}

func assetsDnsMetricsClusterRoleBindingYaml() (*asset, error) {
	bytes, err := assetsDnsMetricsClusterRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/metrics/cluster-role-binding.yaml", size: 279, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x79, 0x95, 0x6f, 0xa4, 0xd5, 0xed, 0x48, 0x27, 0x41, 0x56, 0x5c, 0xea, 0x5c, 0x89, 0xdc, 0xc1, 0x44, 0x91, 0xd4, 0xb, 0x18, 0x85, 0x79, 0x75, 0xaa, 0x6e, 0xb5, 0x98, 0xbe, 0xc6, 0x33, 0x43}}
	return a, nil
}

var _assetsDnsMetricsClusterRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\xcd\x31\x4b\x34\x41\x0c\x87\xf1\x7e\x3e\x45\xe0\xad\x77\x5f\xec\x64\x5a\x05\x3b\x0b\x05\xfb\xec\xce\xdf\xdb\x70\x3b\xc9\x90\x64\x0e\xf4\xd3\x8b\x70\xb6\x0f\x3f\x78\xfe\xd1\xd3\x39\x23\xe1\xe4\x76\x22\x48\x81\x86\x46\xdb\x17\x0d\xb7\x8e\x3c\x30\x83\xd2\x28\x76\xe7\x01\x7a\x7e\x7d\xa7\x8e\x74\xd9\x83\xa0\x6d\x98\x68\x16\x1e\xf2\x01\x0f\x31\xad\xe4\x1b\xef\x2b\xcf\x3c\xcc\xe5\x9b\x53\x4c\xd7\xeb\x63\xac\x62\xff\x6f\x0f\xe5\x2a\xda\xea\xdf\xf0\xcd\x4e\x94\x8e\xe4\xc6\xc9\xb5\x10\x29\x77\x54\x6a\x1a\x4b\x37\x95\x34\x17\xbd\x14\x9f\x27\xa2\x96\x85\x78\xc8\x8b\xdb\x1c\xf1\x4b\x17\xb2\x01\xe7\x34\x5f\x6d\x40\xe3\x90\xcf\x5c\xc5\x0a\x91\x23\x6c\xfa\x8e\x3b\x6b\x1a\x88\x42\x74\x83\x6f\xf7\x74\x41\x96\x9f\x00\x00\x00\xff\xff\x9f\xa8\x4d\x6c\xf6\x00\x00\x00")

func assetsDnsMetricsClusterRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsMetricsClusterRoleYaml,
		"assets/dns/metrics/cluster-role.yaml",
	)
}

func assetsDnsMetricsClusterRoleYaml() (*asset, error) {
	bytes, err := assetsDnsMetricsClusterRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/metrics/cluster-role.yaml", size: 246, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x64, 0xdb, 0xe0, 0x95, 0x65, 0xae, 0x53, 0x96, 0x3a, 0x5f, 0x5e, 0x8b, 0x69, 0xe2, 0x7d, 0x5, 0xbf, 0x1f, 0x3a, 0xf, 0xff, 0xd0, 0x6b, 0x23, 0x4f, 0xfd, 0x11, 0x7f, 0x57, 0xd4, 0x4a, 0x8b}}
	return a, nil
}

var _assetsDnsMetricsRoleBindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\xce\xb1\x4e\xc4\x40\x0c\x04\xd0\x7e\xbf\xc2\x3f\x90\x20\xba\xd3\x76\xd0\xd0\x1f\x12\xbd\x6f\xd7\x97\x98\x64\xed\x95\xed\x4d\xc1\xd7\x23\xa4\x48\x54\x20\x5d\x3b\x9a\xd1\x1b\xec\xfc\x41\xe6\xac\x92\xc1\x6e\x58\x66\x1c\xb1\xaa\xf1\x17\x06\xab\xcc\xdb\xc5\x67\xd6\xa7\xe3\x39\x6d\x2c\x35\xc3\x55\x77\x7a\x65\xa9\x2c\x4b\x6a\x14\x58\x31\x30\x27\x00\xc1\x46\x19\xba\x69\xa3\x58\x69\xf8\xb4\x5d\xfc\x8c\xbd\x63\xa1\x0c\xda\x49\x7c\xe5\x7b\x4c\x55\x3c\x99\xee\x74\xa5\xfb\xcf\x14\x3b\xbf\x99\x8e\xfe\x8f\x9f\x00\x7e\xf9\xbf\x34\x1f\xb7\x4f\x2a\xe1\x39\x4d\x67\xfb\x9d\xec\xe0\x42\x2f\xa5\xe8\x90\x78\xf0\x65\x53\xe1\x50\x63\x59\x20\x7d\x07\x00\x00\xff\xff\xb9\xd9\xab\x8d\x25\x01\x00\x00")

func assetsDnsMetricsRoleBindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsMetricsRoleBindingYaml,
		"assets/dns/metrics/role-binding.yaml",
	)
}

func assetsDnsMetricsRoleBindingYaml() (*asset, error) {
	bytes, err := assetsDnsMetricsRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/metrics/role-binding.yaml", size: 293, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc, 0x7d, 0xc7, 0x45, 0x33, 0xc4, 0xd8, 0xf, 0x8d, 0x89, 0x8d, 0x6, 0x47, 0xa7, 0xa, 0x6b, 0x17, 0xf5, 0x5f, 0x5a, 0x2f, 0xd8, 0xf9, 0x6, 0x71, 0xaa, 0x78, 0x8d, 0xb5, 0x7a, 0xf6, 0x99}}
	return a, nil
}

var _assetsDnsMetricsRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8e\xb1\x4e\xec\x40\x0c\x45\xfb\xf9\x0a\x6b\x5f\x9d\x7d\xa2\x5b\x4d\x8d\x44\x47\x01\x12\xbd\x77\xe6\x42\xac\x24\xe3\x91\xed\x04\xc1\xd7\xa3\xcd\x06\x89\xca\xf7\x1e\x59\x3e\xfe\x47\x2f\x3a\xc3\xa9\x01\x15\x95\xae\x5f\xd4\x4d\x17\xc4\x88\xd5\x29\x94\xbc\x18\x77\xd0\xe3\xf3\x2b\x2d\x08\x93\xe2\x84\x56\xbb\x4a\x8b\xc4\x5d\xde\x60\x2e\xda\x32\xd9\x95\xcb\x99\xd7\x18\xd5\xe4\x9b\x43\xb4\x9d\xa7\x8b\x9f\x45\xff\x6f\x0f\x69\x92\x56\xf3\x2e\x4a\x0b\x82\x2b\x07\xe7\x44\xd4\x78\x41\xfe\xe3\x1b\xa6\x8b\x1f\xd8\x3b\x17\x64\xd2\x8e\xe6\xa3\xbc\xc7\x50\x9b\x27\x5b\x67\x78\x4e\x03\x71\x97\x27\xd3\xb5\xfb\xed\xca\x40\xa7\x53\x22\x32\xb8\xae\x56\x70\x30\x87\x6d\x52\xe0\x7b\xf9\xfd\xf8\xde\xba\xd6\x5b\xd8\x60\xd7\x63\xf9\x03\xb1\xcf\x59\xfc\x1e\x3e\x39\xca\x98\x7e\x02\x00\x00\xff\xff\x29\x39\xda\x05\x1c\x01\x00\x00")

func assetsDnsMetricsRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsMetricsRoleYaml,
		"assets/dns/metrics/role.yaml",
	)
}

func assetsDnsMetricsRoleYaml() (*asset, error) {
	bytes, err := assetsDnsMetricsRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/metrics/role.yaml", size: 284, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x8c, 0xf2, 0x4e, 0x40, 0x91, 0xd8, 0x5e, 0x1c, 0x98, 0xb6, 0x2f, 0x11, 0x2a, 0x15, 0x8f, 0xe4, 0x7c, 0xfe, 0xc6, 0x31, 0xf3, 0xb2, 0xa0, 0x38, 0xb2, 0x3f, 0x15, 0x5a, 0x33, 0x12, 0xd2, 0x88}}
	return a, nil
}

var _assetsDnsNamespaceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x90\xcd\x4e\xc4\x30\x0c\x84\xef\x79\x8a\x51\x38\x2f\x3f\xd7\xbc\x03\x5c\x90\xb8\xbb\x8d\x97\x35\x4d\xed\x2a\x76\xcb\xeb\xa3\xb2\x15\xac\xb4\xc7\x68\x46\xf3\x7d\xf1\x24\x5a\x0b\xde\x68\x66\x5f\x68\xe4\x44\x8b\x7c\x70\x77\x31\x2d\xd8\x5e\xd2\xcc\x41\x95\x82\x4a\x02\x48\xd5\x82\x42\x4c\x7d\x7f\x02\xb6\xb0\xfa\x45\xce\xf1\x28\xf6\xa4\x56\xf9\xe4\xdc\x78\x0c\xeb\x05\x39\x27\x40\x69\xe6\xf2\x5f\x3b\x55\xf5\x04\x34\x1a\xb8\x1d\x13\x0f\x70\x0e\x6c\xd4\x56\x46\x18\x68\x33\xa9\xa8\xbc\xb0\x56\xd1\x4f\x98\x62\x5a\x07\x06\xd5\x59\x7c\x97\x42\x5c\x28\x8e\x82\xef\xf1\xdf\x38\x68\x11\xbf\xd7\xea\xab\x9e\x1a\x6f\xdc\x0a\xf2\x73\x3e\x98\xd4\x9a\x7d\xdf\x78\xcd\xa6\x12\xd6\x77\x62\x18\x9a\xd9\x84\xb3\x75\xbc\x73\xdf\x64\xe4\xd7\x6b\x0a\x1b\xbe\x78\x0c\x87\xec\x16\xe2\xbf\xbf\xbb\x1e\xed\x8e\x3a\xb6\xd5\x83\xfb\xcd\x70\x41\x8e\xbe\x72\x4e\x3f\x01\x00\x00\xff\xff\x82\x6d\x29\x03\x71\x01\x00\x00")

func assetsDnsNamespaceYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsNamespaceYaml,
		"assets/dns/namespace.yaml",
	)
}

func assetsDnsNamespaceYaml() (*asset, error) {
	bytes, err := assetsDnsNamespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/namespace.yaml", size: 369, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe, 0xab, 0x50, 0x84, 0x61, 0x5f, 0x41, 0xf4, 0x17, 0x3b, 0x6, 0x84, 0xc0, 0x5f, 0x4f, 0xbb, 0xd8, 0x1d, 0xae, 0x26, 0x3e, 0x1f, 0x29, 0x2c, 0x84, 0x6d, 0x5e, 0xc1, 0x87, 0x97, 0x5f, 0xc9}}
	return a, nil
}

var _assetsDnsServiceAccountYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xc9\xb1\x09\xc4\x30\x0c\x05\xd0\xde\x53\x68\x81\x2b\xae\x55\x77\x33\x1c\xa4\x17\xf2\x0f\x11\xc1\xb2\xb1\x14\xcf\x1f\x02\xe9\x1e\xbc\xd3\xbc\x32\xfd\x31\x97\x29\x7e\xaa\xfd\xf2\x2c\x32\x6c\xc3\x0c\xeb\xce\xb4\xbe\xa5\x21\xa5\x4a\x0a\x17\x22\x97\x06\xa6\xea\xf1\x3a\x86\x28\x98\xfa\x80\xc7\x61\x7b\x7e\x9e\xba\x03\x00\x00\xff\xff\x8e\x2c\xf1\x2e\x55\x00\x00\x00")

func assetsDnsServiceAccountYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsServiceAccountYaml,
		"assets/dns/service-account.yaml",
	)
}

func assetsDnsServiceAccountYaml() (*asset, error) {
	bytes, err := assetsDnsServiceAccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/service-account.yaml", size: 85, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x57, 0x12, 0x50, 0x4d, 0x67, 0x2f, 0x1b, 0x74, 0xa0, 0xa4, 0xbb, 0xa7, 0x59, 0xe9, 0x5a, 0xc6, 0xc1, 0x1a, 0xf8, 0x5f, 0xff, 0x5, 0xdb, 0xc, 0x10, 0x8b, 0xc1, 0x0, 0xcc, 0xf, 0x9f, 0x3a}}
	return a, nil
}

var _assetsDnsServiceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x91\x31\x6f\xe2\x40\x10\x85\x7b\xff\x8a\x27\xe8\x4e\xc0\x09\xdd\x51\x9c\xdb\xa3\x89\x52\x80\x14\x48\x3f\x5e\x4f\xcc\x8a\xf5\x8c\xb5\x33\x06\xf1\xef\x23\x4c\x42\x80\x14\x69\x56\xda\x7d\x9f\x3e\x3d\xbd\xdd\x47\xa9\x4b\xbc\x70\x3e\xc4\xc0\x05\x75\xf1\x95\xb3\x45\x95\x12\x87\x79\x31\x86\x50\xcb\x93\xe1\xb4\x8e\x02\x4f\x12\x55\x9c\x0c\x24\x35\x48\x44\x9d\x3c\xaa\x18\x28\x33\x8c\x1d\xe4\xc8\xbd\x78\x6c\xb9\xb0\x8e\x43\x59\x00\x63\x84\xd4\x9b\x73\x7e\x5a\xe3\x18\x53\x42\xc5\xa0\xde\xb5\x25\x8f\x81\x52\x3a\xa1\x25\xa1\x86\xeb\xd9\x00\x1b\x27\x0e\xae\x19\xd1\x1e\x8d\x40\xa7\xd9\xed\x2c\x9d\x0e\x95\x4a\xd4\x62\x05\x70\x09\x4a\x2c\xfe\x0c\x17\xa7\xdc\xb0\xaf\x87\xa7\x2b\x90\xd5\x35\x68\x2a\xb1\x5d\xae\xef\x05\x53\x0f\xdd\x8f\x92\x2f\xe8\x2a\xda\xfc\xbf\x15\xb5\xec\x39\x86\xdb\x36\xff\xe6\x8b\xbf\xdf\x54\x77\xd8\x83\x6a\x8c\xcd\x6a\xb9\x2a\xb1\x95\xa0\x6d\xcb\xe2\x38\xee\x58\x60\x97\xbf\x81\x6b\xa7\x49\x9b\x13\xde\x98\xbc\xcf\x8c\x86\x9c\xcf\x33\xb1\x50\x95\x3e\xf6\xfb\x84\x9e\xf9\x64\x97\xf5\x31\xc5\x68\xdf\x57\x9c\x85\x9d\x6d\x16\xf5\xf7\x4e\xcd\xcf\xa5\x47\xd7\xfc\xd7\xa8\x78\x0f\x00\x00\xff\xff\x82\x42\x75\xa4\x08\x02\x00\x00")

func assetsDnsServiceYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsServiceYaml,
		"assets/dns/service.yaml",
	)
}

func assetsDnsServiceYaml() (*asset, error) {
	bytes, err := assetsDnsServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/service.yaml", size: 520, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x18, 0x69, 0xc5, 0xf1, 0xe, 0xc, 0x77, 0xe5, 0x78, 0xce, 0xfc, 0xc2, 0x41, 0xf8, 0x21, 0x87, 0x8a, 0xb7, 0x67, 0xdd, 0x48, 0x94, 0x63, 0x79, 0x69, 0x4e, 0x38, 0x53, 0x3c, 0xdb, 0xc7, 0x13}}
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

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
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

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"assets/dns/cluster-role-binding.yaml": assetsDnsClusterRoleBindingYaml,

	"assets/dns/cluster-role.yaml": assetsDnsClusterRoleYaml,

	"assets/dns/daemonset.yaml": assetsDnsDaemonsetYaml,

	"assets/dns/metrics/cluster-role-binding.yaml": assetsDnsMetricsClusterRoleBindingYaml,

	"assets/dns/metrics/cluster-role.yaml": assetsDnsMetricsClusterRoleYaml,

	"assets/dns/metrics/role-binding.yaml": assetsDnsMetricsRoleBindingYaml,

	"assets/dns/metrics/role.yaml": assetsDnsMetricsRoleYaml,

	"assets/dns/namespace.yaml": assetsDnsNamespaceYaml,

	"assets/dns/service-account.yaml": assetsDnsServiceAccountYaml,

	"assets/dns/service.yaml": assetsDnsServiceYaml,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
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
	"assets": {nil, map[string]*bintree{
		"dns": {nil, map[string]*bintree{
			"cluster-role-binding.yaml": {assetsDnsClusterRoleBindingYaml, map[string]*bintree{}},
			"cluster-role.yaml":         {assetsDnsClusterRoleYaml, map[string]*bintree{}},
			"daemonset.yaml":            {assetsDnsDaemonsetYaml, map[string]*bintree{}},
			"metrics": {nil, map[string]*bintree{
				"cluster-role-binding.yaml": {assetsDnsMetricsClusterRoleBindingYaml, map[string]*bintree{}},
				"cluster-role.yaml":         {assetsDnsMetricsClusterRoleYaml, map[string]*bintree{}},
				"role-binding.yaml":         {assetsDnsMetricsRoleBindingYaml, map[string]*bintree{}},
				"role.yaml":                 {assetsDnsMetricsRoleYaml, map[string]*bintree{}},
			}},
			"namespace.yaml":       {assetsDnsNamespaceYaml, map[string]*bintree{}},
			"service-account.yaml": {assetsDnsServiceAccountYaml, map[string]*bintree{}},
			"service.yaml":         {assetsDnsServiceYaml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
