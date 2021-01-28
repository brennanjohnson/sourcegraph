// Code generated by go-bindata. DO NOT EDIT.
// sources:
// 1000000000_init.down.sql (19B)
// 1000000000_init.up.sql (19B)
// 1000000001_initial_schema.down.sql (471B)
// 1000000001_initial_schema.up.sql (4.428kB)

package migrations

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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
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

var __1000000000_initDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\x48\xcd\x2d\x28\xa9\x54\xc8\xcd\x4c\x2f\x4a\x2c\xc9\xcc\xcf\xe3\x02\x04\x00\x00\xff\xff\x32\x4d\x68\xbd\x13\x00\x00\x00")

func _1000000000_initDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1000000000_initDownSql,
		"1000000000_init.down.sql",
	)
}

func _1000000000_initDownSql() (*asset, error) {
	bytes, err := _1000000000_initDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1000000000_init.down.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x9c, 0x46, 0xd1, 0x31, 0xb9, 0x68, 0x19, 0xcc, 0x70, 0xb6, 0x7, 0x20, 0x2e, 0x6a, 0x4d, 0xf1, 0xce, 0xd0, 0xc8, 0xda, 0x50, 0xce, 0x8c, 0xee, 0x52, 0x36, 0x80, 0xd0, 0x5a, 0xd2, 0x7a, 0x82}}
	return a, nil
}

var __1000000000_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\x48\xcd\x2d\x28\xa9\x54\xc8\xcd\x4c\x2f\x4a\x2c\xc9\xcc\xcf\xe3\x02\x04\x00\x00\xff\xff\x32\x4d\x68\xbd\x13\x00\x00\x00")

func _1000000000_initUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1000000000_initUpSql,
		"1000000000_init.up.sql",
	)
}

func _1000000000_initUpSql() (*asset, error) {
	bytes, err := _1000000000_initUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1000000000_init.up.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x9c, 0x46, 0xd1, 0x31, 0xb9, 0x68, 0x19, 0xcc, 0x70, 0xb6, 0x7, 0x20, 0x2e, 0x6a, 0x4d, 0xf1, 0xce, 0xd0, 0xc8, 0xda, 0x50, 0xce, 0x8c, 0xee, 0x52, 0x36, 0x80, 0xd0, 0x5a, 0xd2, 0x7a, 0x82}}
	return a, nil
}

var __1000000001_initial_schemaDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x8f\xc1\xaa\xc2\x30\x10\x45\xf7\xf9\x8a\xfc\x47\x56\xaf\xaf\x51\x02\xb6\x15\x9b\x45\x77\x61\xa4\x43\x08\x98\x54\xd3\x89\xf8\xf9\x22\x62\x05\x65\x50\x77\xb3\xb8\xe7\x9e\x3b\x95\x5e\x9b\x56\x09\x51\xef\xba\xad\x34\x6d\xad\x07\x69\x56\x52\x0f\xa6\xb7\xbd\xf4\x50\x3c\x3a\x3c\x63\xa2\xd9\x65\x3c\x4e\x2e\x8c\x6e\x4f\x19\x51\x7d\x0b\x24\x88\xf8\x13\x35\xe5\xe0\x43\x82\x03\x8f\xdb\xbf\x6a\xa3\x19\x9c\xfb\x64\x29\x9b\xef\x95\x25\x85\x53\xb9\x35\x5f\x98\x4d\xaf\x00\x65\x1f\x19\xff\x33\xca\xd9\x23\x12\x8c\x40\xe0\x96\xe3\xa3\xff\x1d\xf1\x21\x31\x03\x1e\x11\x25\xc4\x7f\xd7\x34\xc6\x2a\x71\x0d\x00\x00\xff\xff\x0d\x60\xa7\x10\xd7\x01\x00\x00")

func _1000000001_initial_schemaDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1000000001_initial_schemaDownSql,
		"1000000001_initial_schema.down.sql",
	)
}

func _1000000001_initial_schemaDownSql() (*asset, error) {
	bytes, err := _1000000001_initial_schemaDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1000000001_initial_schema.down.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x46, 0x11, 0x7d, 0x48, 0x3a, 0xe0, 0x31, 0xc5, 0x3a, 0x2e, 0xd9, 0xa8, 0x64, 0xed, 0xe9, 0x73, 0xfc, 0x74, 0x4d, 0xdc, 0x8e, 0x2f, 0x32, 0xf6, 0xcf, 0x71, 0xfc, 0x39, 0x9b, 0xc1, 0x92, 0xca}}
	return a, nil
}

var __1000000001_initial_schemaUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\x5f\x73\xdb\xb8\x11\x7f\xd7\xa7\xd8\xb9\x87\x5a\x9a\xa1\xe9\x77\xbb\x73\xa9\x22\x31\x39\xf6\x6c\x2a\x95\x98\x69\xae\x9d\x8e\x06\x22\x57\x24\x2e\x20\xc0\x00\xa0\x64\xf5\xcf\x77\xef\x2c\x40\x52\xa4\x2c\x27\x99\x4c\xeb\x17\x6b\x88\xc5\x6f\x7f\xfb\x7f\xf1\x36\x7a\x1f\x27\x0f\x93\xc9\x62\x1d\xcd\xd3\x08\xa2\x4f\x69\x94\x6c\xe2\x55\x02\xf1\x3b\x48\x56\x29\x44\x9f\xe2\x4d\xba\x01\xcb\x2b\x34\x19\x13\x98\xef\x1e\xbe\x25\x5b\x17\x5b\xab\x8b\xea\x9b\x72\x19\xb7\xf8\x6c\x1f\x26\x93\xdb\x5b\x58\x63\xa6\x74\x6e\x40\x63\xad\x0c\xb7\x4a\x9f\x40\xb2\x0a\x4d\x00\x3b\x65\x4b\x28\xb9\xb1\x4a\xf3\x8c\x09\x60\x32\x87\x5a\xa3\x41\x69\x03\x68\x0c\x97\x05\x30\x68\x24\xff\xd2\xe0\xf0\xf6\x96\xae\x6f\x21\x5e\x12\xfa\xb4\x91\x1a\x05\xb3\x98\x83\x55\x60\xcb\x91\x64\xbc\x0c\x67\x1d\xd5\x74\xfe\xf6\x31\x72\x87\xee\xbe\x81\xe9\x04\x00\xe0\xf6\x16\xd2\xf2\x3a\x7c\xe8\x04\x78\x0e\x3b\x5e\x18\xd4\x9c\x09\x67\x62\xf2\xf1\xf1\x11\x3e\xac\xe3\xa7\xf9\xfa\x37\xf8\x35\xfa\x2d\x98\x0c\x81\xe8\x72\x00\x56\xf3\x42\xb3\xea\x96\xcb\x1c\x9f\x31\x87\xbd\xd2\xb0\x67\xc6\x02\x86\x45\x08\x1a\x0b\x7c\xae\x61\xcf\x85\x45\xcd\x65\xe1\x15\xd1\xcd\xd6\x71\xbd\x9e\x16\x7b\xb1\x4a\x36\xe9\x7a\x1e\x27\x29\x64\x25\x66\x9f\x3d\x45\xa9\x24\x56\xb5\x3d\xc1\xe2\x97\x68\xf1\x2b\x4c\xa7\x0e\x61\xf5\x21\x5a\xcf\xd3\xd5\x7a\xfa\xc7\x9f\x67\x70\x73\x73\x7f\xef\x21\x67\xb3\xc9\xcc\xc7\x23\x92\x7b\xa5\x33\x04\x5b\x32\xeb\x23\x01\x4c\x63\xeb\xe7\xb0\x73\xd7\xc7\x24\xfe\xcb\xc7\x08\xe2\x64\x19\x7d\x1a\x78\xcd\x6b\xf6\xb2\x5b\x9e\x3f\xc3\x2a\x19\x9c\x3a\x06\xad\x9a\x85\x46\x66\xb1\xf3\x04\x78\x4f\x18\xe7\x89\x8b\x4c\x38\x3b\x02\x76\xcc\x60\x0e\x4a\x0e\xfd\x64\x7a\x4e\xd7\xc9\x50\x42\x8e\x69\xc0\xc7\x4d\x9c\xbc\x87\x82\x4b\x98\x0a\x75\x44\xed\x5d\x33\xbb\xbf\x77\x9e\xa0\x03\x77\x6b\xab\x6a\x43\x6c\x87\x59\xca\xf4\x8e\x5b\xcd\xf4\x09\x2a\xb4\x2c\x67\x96\x01\xdb\xa9\xc6\x02\x1e\x50\x5a\x13\xc2\xc6\x2a\x8d\x39\x70\x09\x0c\x0c\xd6\x4c\x3b\x2b\xd9\x4e\x20\x30\x03\xdc\x02\x37\xa0\xf6\x16\x25\x11\x22\x17\xe4\x04\x4f\x66\x57\x8d\xb0\xbc\x16\xd8\x41\x8d\x33\xb3\x57\xf7\xbf\xcf\xcb\xa7\xb1\x25\xb6\xe4\xc6\x73\x08\xfc\xef\x8c\x49\xd8\x21\x30\x79\x1a\x98\xff\xe7\xcd\x2a\x39\x93\x3a\x96\x3c\x2b\xe1\xc8\x85\x20\x49\x8d\xb6\xd1\x12\xf3\x4e\xc1\xb1\x44\x09\x5f\x1a\xd4\x27\x0a\xa2\x37\x2f\x70\xe5\xdc\x42\xfb\x08\xfb\xd0\xd2\xe7\x42\xab\xa6\xc6\xbc\x2d\xf2\xdf\x8d\x92\x3b\x50\x35\x6a\x66\x95\x36\xf0\x26\x80\x37\x7f\x08\xe0\xcd\xbf\x83\x4e\x01\xdd\xf9\xd3\xcf\x21\xa4\x44\xd7\x94\xaa\x11\x39\xc1\x9a\x8a\x09\x01\x8e\xa0\x92\xe2\x14\x40\xad\x79\x45\xe4\x1b\x83\x90\x31\x83\x14\x0c\x2f\x24\xb8\xb1\x06\x4c\x93\x95\xc0\xcc\x7d\x8b\xdb\xc1\xc3\xbf\x7e\xfa\x9d\x1d\xd8\xf6\x80\xda\x70\x25\xcd\x4f\xf7\xf0\xf7\x30\x0c\xff\xf1\x9f\x81\x80\x60\xb2\x68\x58\x81\x74\x48\x7f\x2f\x04\xea\x46\x88\xad\xc6\x2f\x0d\x1a\x7b\x15\x81\x49\xa9\x2c\xb3\xad\x82\x0b\x04\xf7\xaf\x77\xb7\xf7\x48\x17\xd5\xab\x95\xdb\xcb\x72\xf3\xd5\xda\xed\xe4\xb6\xfd\x8f\x71\xf5\x76\x9f\xa7\xdd\x8f\x56\x59\x4c\x05\x7b\xd6\x62\x15\xa8\xda\xf2\x8a\xff\x13\xe1\xaf\xbf\x44\xeb\x08\x32\xc1\x1a\x83\x06\x8e\xdc\x96\x2d\xe1\x73\xe0\xda\x88\x9d\x83\x7a\x51\xc4\x2f\x59\x51\xb5\x0e\x33\xce\xd7\xf0\xfb\x38\x81\x4b\x66\x5d\xa9\xfa\x3c\x03\x75\x40\xed\x86\x18\x30\x63\x54\xc6\xdd\x2c\x70\xa4\xd8\xb0\x7c\xa6\x4a\x03\x75\xcc\x00\x78\x88\x21\x14\x42\xed\x98\x10\xa7\x19\x25\xaf\x46\x2a\x66\x2e\x0b\x81\xa4\x40\x36\x15\xfa\x89\x74\x60\xa2\x71\x49\x54\x28\x37\x8d\xda\xea\xe0\xe2\x04\x4d\xed\x6c\xcc\xd5\x51\x86\x93\xdb\x5b\x4f\xac\xd7\xd6\x51\xe1\x4a\xd2\xf5\xbe\xaf\xb9\x91\x37\x9a\x50\x0e\x85\x4a\x3b\x74\x05\x1f\x2f\xbb\x9a\x69\x8c\x9f\x69\x1a\xf7\x64\xa0\x22\x0d\x0c\x4c\x8d\x19\xdf\xf3\x6c\x00\x12\x80\xd2\x20\x94\xfa\xdc\xd4\x6e\x00\x66\x8d\xd6\x28\x7d\x6f\x07\xb5\x1f\xbb\x81\xed\x2d\x6a\x6a\x53\x25\x33\xb0\x43\xec\x5b\x2d\x49\xe7\x64\x49\x3f\xc6\x5e\x23\xe2\x94\x74\xe0\x17\x13\x97\x59\xf7\xc5\x85\xa3\x3d\x75\x61\xba\x31\x90\xd1\x40\xe0\x4a\x06\x5d\x3f\xc4\x67\x56\x51\x3b\x24\x44\xcd\x5c\x5e\x23\x64\x25\x93\x05\xfa\xf6\x5a\xb0\xa6\x40\xd8\xb1\xec\x33\xc9\x8c\xcc\xd8\x21\xc5\xa3\x67\x3d\xea\xa4\xee\xda\xb6\xcd\x8e\x51\x37\x75\xab\x8e\x65\x55\x7d\x66\x4e\xa9\x84\xb9\x27\xe9\x1b\xab\xe3\x9e\xc6\x4f\xd1\x26\x9d\x3f\x7d\x48\xff\x76\x39\x8c\x5b\xac\xbd\x50\xcc\x12\x89\x5a\x71\x69\xdb\x4c\x79\xcd\x7c\x8f\xec\x65\x72\xd5\xd0\xb0\xa8\x35\x66\x9c\xba\xcd\x15\xfc\xf9\x39\x91\xfb\x7a\x20\x8f\x0d\x3b\x37\xdf\x53\xc3\x0e\x47\x7d\x63\xcb\x69\x30\x59\x2c\x50\x8f\xc9\x8e\x33\x6e\xba\xd7\xaa\x72\xe4\x2a\x46\x8e\xae\x6b\xc1\x33\x9f\xab\xcb\xb7\xb3\x91\x11\xbd\x05\x70\x64\x6d\x08\x31\x0f\x21\x51\x16\x3b\x7c\xd7\x8d\x2e\xd2\xa0\x62\x27\x90\x0a\x84\x92\x05\x52\xa0\xb9\xb1\x70\x47\xb9\x74\x60\x82\xe7\xa4\xc1\x4d\x0b\xa7\x23\x80\x52\x1d\xf1\x80\x3a\xbc\xe8\xca\xb2\x11\x82\xcc\x1c\x73\x90\xca\x3a\x5f\x74\x05\x3b\xaa\x71\x57\xda\xac\x2d\x6e\x9f\x08\x33\x0f\xeb\xb6\x83\xd7\xdc\x53\x29\x63\x29\x17\x50\x5a\x71\x82\xcf\x52\x1d\x65\xbb\x99\x38\xa7\xe3\xa8\xd8\x9a\x3a\x77\x91\xa9\x51\x73\x95\x53\x9b\x10\x27\x97\x9f\x59\xa6\x1a\xe9\xc9\x51\x4d\x75\x0a\x06\xfc\x7c\xbe\x9a\x10\xe2\x17\x85\x43\xa6\xe5\x28\xd0\x62\xde\x4e\x65\x1a\x5d\x96\x06\xae\xbd\xce\xb0\x77\x13\xf5\x8e\xff\xbb\xeb\xdc\xf2\xf1\x1d\xe9\xe5\xdc\xe6\x17\x21\x52\xe9\x7d\xf9\x1d\x19\x15\x5b\x97\x34\x25\x3b\xa0\x6f\x4b\x6d\x6d\x77\x6a\x0c\x97\x59\x6b\xa6\xd2\xbc\xe0\x92\xd1\xb0\xfd\x1a\xb1\x48\x9a\x46\x23\x79\x41\x49\x4f\x71\xd4\x95\xf7\x1c\x45\xee\x9c\xec\x3b\x2a\xf9\x9d\x16\x05\xa6\x5b\x35\x2f\x96\x6e\xa7\xce\x5d\x33\xdb\xf6\xd2\x79\xf9\x76\x57\xe8\x6f\x3a\xed\x52\x2d\xde\xb8\xb2\x9e\xc1\x3c\x59\xc2\x74\x44\x76\x7c\x74\xdd\xa0\x4e\x66\xd6\x43\xaf\xd6\xaf\x68\x69\x1b\xc8\x6b\x9a\xc6\xc7\xaf\x6b\xeb\xe4\xbc\xc6\x59\xeb\xca\x77\xab\x75\x14\xbf\x4f\x68\xa5\x3c\x8f\xe2\x2d\xcf\x67\xb0\x8e\xde\x45\xeb\x28\x59\x44\x9b\xf3\x16\x41\xdf\x57\x09\x2c\xa3\xc7\x28\x8d\x60\x31\xdf\x2c\xe6\xcb\x08\x96\x24\xb9\xa6\xee\x1c\xbc\xc4\x1c\xf2\x18\x81\x0e\x1e\x16\x3f\x00\x7b\xd5\xce\x1f\xc5\x9f\x8c\x5f\x35\xe5\xa9\x46\xed\x96\xfe\x00\x6a\xa6\x2d\xa7\x94\x3a\xef\xbe\xb0\xf3\xad\xcd\x8d\xd3\x0d\x22\x94\xd6\xd6\xe6\xfe\xee\x2e\x57\x99\x09\xfb\x07\x77\x98\xa9\xea\x8e\xde\xae\xc6\xde\xb9\x4d\xf8\x76\xf0\x16\xbf\x3b\xeb\x30\x93\x4d\xf4\x18\x2d\xd2\xb6\x5a\xb6\xe7\x93\xe9\xcd\x70\xd8\xdd\x04\x70\x43\x08\x37\x63\xb2\x3b\xab\x11\x5f\x7b\x80\x0d\x1e\xa1\xa3\x05\x6d\x88\xbb\x6d\x73\x6d\xeb\x91\x56\xc9\x78\xc4\xfa\x4d\xcd\x9f\x75\x69\x39\x7b\xf8\x16\x5c\x1b\x90\xef\xc6\xec\x02\xf8\x35\xe0\xab\x21\xff\x1e\x0d\xd7\x73\xe5\x61\x32\x59\xac\x9e\x9e\xe2\xf4\x61\xf2\xdf\x00\x00\x00\xff\xff\x04\xb2\x25\xad\x4c\x11\x00\x00")

func _1000000001_initial_schemaUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1000000001_initial_schemaUpSql,
		"1000000001_initial_schema.up.sql",
	)
}

func _1000000001_initial_schemaUpSql() (*asset, error) {
	bytes, err := _1000000001_initial_schemaUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1000000001_initial_schema.up.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x92, 0xb4, 0x34, 0xea, 0xec, 0x2b, 0x56, 0x72, 0xb, 0x7c, 0xa2, 0xb2, 0x6d, 0xae, 0x73, 0x94, 0x87, 0x93, 0x31, 0x3b, 0x53, 0xce, 0x4d, 0x8d, 0x96, 0xd1, 0x35, 0xdc, 0xdf, 0xa, 0x86, 0x2e}}
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
	"1000000000_init.down.sql":           _1000000000_initDownSql,
	"1000000000_init.up.sql":             _1000000000_initUpSql,
	"1000000001_initial_schema.down.sql": _1000000001_initial_schemaDownSql,
	"1000000001_initial_schema.up.sql":   _1000000001_initial_schemaUpSql,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

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
	"1000000000_init.down.sql":           {_1000000000_initDownSql, map[string]*bintree{}},
	"1000000000_init.up.sql":             {_1000000000_initUpSql, map[string]*bintree{}},
	"1000000001_initial_schema.down.sql": {_1000000001_initial_schemaDownSql, map[string]*bintree{}},
	"1000000001_initial_schema.up.sql":   {_1000000001_initial_schemaUpSql, map[string]*bintree{}},
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
