// Code generated by go-bindata.
// sources:
// migrates/mysql/000001_gokins.down.sql
// migrates/mysql/000001_gokins.up.sql
// migrates/sqlite/000001_gokins.down.sql
// migrates/sqlite/000001_gokins.up.sql
// DO NOT EDIT!

package comm

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

var _mysql000001_gokinsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xd1\xcd\x4a\x43\x41\x0c\x05\xe0\xbd\xe0\x3b\xcc\x7b\x74\xa5\x58\xa1\x20\x28\xb6\x0b\x77\xb9\x71\x9a\x3b\x84\xce\x1f\x99\x4c\xa1\x6f\x2f\x15\xc4\x5d\x72\xd7\xe7\x63\x38\x39\xf3\xf2\xf9\xfe\x11\x4e\x4f\xcf\x6f\xfb\x70\x78\x0d\xfb\xaf\xc3\xf1\x74\x0c\x8b\x02\x8a\xf2\x8a\x51\xa1\x63\xbc\x60\xa2\x65\xf7\xf8\xe0\xda\x2b\xc9\xe0\x56\xb7\xd8\x26\x37\x8b\x7d\x4f\xce\x67\x0b\xc4\x72\x86\xcc\xd5\xec\x35\xd4\x29\x3e\x94\xba\x95\xab\x70\x4a\x24\x1b\x08\xc8\x34\xcf\x2e\x34\x86\x53\xa6\x49\x72\x62\xe8\xdc\xcd\x27\xee\xb9\x37\xca\x9f\x81\xd8\xea\xba\x09\x5e\xd1\x5c\xe0\xdf\xf9\xbf\xdf\x51\xb0\x58\x60\x0e\x7b\xee\x7b\x0e\x5c\xd7\xe6\x22\x67\xce\x5f\x53\x86\x6f\xb4\x5d\xc8\xbc\xe9\x56\x32\xf4\x3c\x13\xbb\x4a\xa9\xf4\x8c\x4a\xcb\xee\x27\x00\x00\xff\xff\x66\x57\xeb\xe4\x76\x03\x00\x00")

func mysql000001_gokinsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		_mysql000001_gokinsDownSql,
		"mysql/000001_gokins.down.sql",
	)
}

func mysql000001_gokinsDownSql() (*asset, error) {
	bytes, err := mysql000001_gokinsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mysql/000001_gokins.down.sql", size: 886, mode: os.FileMode(438), modTime: time.Unix(1628565625, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mysql000001_gokinsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\xdd\x6e\x1b\xc7\x15\xbe\x37\xa0\x77\x18\x24\x05\x48\x01\x52\x42\xca\x52\x6a\xab\x08\x50\xd9\x66\x1d\xa2\x12\x55\xc8\x54\x9b\x00\x06\xc8\xe1\xee\x88\xdc\x68\x77\x66\x31\x3b\xa4\x2c\xa0\x17\x12\x5a\xb7\x72\xdd\xc2\x76\xe3\xc4\x36\xdc\xb4\x15\x0c\xc4\x09\xe2\x44\x2e\xe0\xaa\x06\x25\xd7\x0f\x13\x2d\x29\x5f\xe5\x15\x8a\xfd\x23\xf7\x7f\x67\xa9\xa5\xe5\xa0\xe1\x0d\xa5\x99\xb3\xb3\x33\x67\xbe\xf3\xcd\xf9\xce\xf0\xe2\x4a\x69\xa1\x5a\x02\xd5\x85\x0b\x8b\x25\x50\x67\x35\x48\x99\xb4\x06\x05\x56\x53\xa1\xb0\x0e\x9b\xa8\x0e\xf2\x13\x67\x00\xa8\x4b\x62\x1d\x74\x20\x15\x5a\x90\xe6\xdf\x9b\x9d\x04\x95\xe5\x2a\xa8\xac\x2e\x2e\x4e\x99\xbd\xd0\xe8\x6e\x48\x4d\x09\xb3\xfc\x4c\x61\xd8\x0b\x16\x56\xab\xcb\xb5\x72\xe5\xe2\x4a\x69\xa9\x54\xa9\x5a\xc6\x14\xa9\xa4\xe6\x1f\xef\x52\xe9\x17\x0b\xab\x8b\xee\x31\x31\x54\xd0\xd0\xa6\x58\x28\x84\x19\x89\x92\xa6\xca\x70\xb3\xe6\x35\x9e\x99\x9b\x0b\x35\x46\x9a\x30\x34\x9a\x0b\x1f\x51\xa0\x08\x32\x24\xd6\x81\x08\x19\x62\x92\x82\x42\x6c\xda\xaa\x98\x68\x23\x22\x19\x99\x36\x86\x4f\x8a\xe1\xd3\x31\x2d\x6a\xc6\x00\x71\x43\xfd\x6a\xa5\xbc\xb4\xb0\xf2\x11\xf8\x65\xe9\x23\x90\x37\x5d\x3d\x65\xee\xc7\xa4\xd9\x6b\xb4\xd6\x55\xc3\x9d\xf9\x81\x67\x5d\x3d\x54\xc5\x8a\xbb\x6b\xca\x76\xec\xe4\xc4\x99\x49\x50\xaa\x5c\x2e\x57\x4a\xe0\x7d\x50\xc6\x98\x5c\xba\x30\x78\xf5\xc5\x0f\x16\x56\xae\x94\xaa\xe0\x7d\xd0\x66\x6b\xe7\x94\xc6\xec\xcf\x26\xce\x44\x02\xa5\x83\xa8\x26\x11\x7c\x7a\x40\xb1\x91\x9a\x15\xa4\x06\xeb\x49\xb0\xd3\x5a\x30\x19\x9e\x3c\x88\x53\x29\xea\x48\x68\x23\x06\x29\x19\x61\x92\x03\x48\x63\x87\x0b\xa1\x9b\x63\x40\x4a\x3b\x79\xef\x09\x6d\x72\x20\x44\x12\x11\x66\xd2\x9a\x84\xa8\x7b\xdb\x92\xb1\x34\x13\x49\x4f\xb0\x21\x87\xb0\x40\xae\x90\x03\x17\x97\x97\x8c\x15\x80\x5c\xef\xfe\x9e\x7e\xfb\x0b\xfd\xc5\x5f\x7b\xbb\x8f\xf2\x45\xeb\xfb\xb7\x85\xde\x37\x8f\xf4\xe7\xcf\x27\x73\x36\xda\x48\x9b\x0a\x28\x69\x52\x5c\x78\x93\x49\x93\x24\x72\xe5\x0f\x88\x06\xc1\xea\x95\x72\xe5\x32\xb8\x50\x5d\x29\x95\x4e\x0a\xd3\x46\x5b\x92\x45\x3e\x80\xaa\x92\x8a\x64\x09\x07\x79\xc7\x44\x69\x48\x98\x3b\xf6\x36\xc1\xf0\x3e\xa7\x31\xc8\xda\x9a\x8f\x6a\x02\xb6\x2e\x34\xfd\xfd\xf7\xfa\x41\xb7\xff\xa7\xfd\xde\xd6\xb6\x0d\x1e\x44\x29\xa1\x3e\x5c\xc4\x8c\xf0\xea\xee\x83\xe3\xbd\xbd\xa3\x97\xbb\xbd\xed\x3d\x67\x84\x0e\xc2\x8c\x7f\x0e\x47\xdd\x9b\x47\x07\xfb\xb9\xc1\x02\xa8\x07\x26\xf9\xf8\x87\xf5\xc3\x2d\xfd\xf1\xcd\xde\xbd\xfd\x57\xf7\x9e\xd9\x43\xac\x49\x58\xd2\x5a\x29\xc6\xe8\x1f\x7c\xd2\xfb\xfc\x1f\x9e\x31\x02\x88\x4e\x9a\xc6\xce\x43\xfd\xa0\xeb\x19\x22\x00\xf8\x84\x21\x7a\x0f\x9f\xf5\x3e\x7b\xea\x19\x22\x70\xba\x98\xf1\x17\xb7\x92\x1b\x3b\xbd\xbf\x3d\xc9\x85\x84\x40\xd6\xd8\x17\x14\xb1\x66\x00\x94\x0f\xfe\x4d\x4a\xda\x2a\x2f\x86\xcd\xb0\xe2\x07\x3c\xe2\x1e\xd8\x1f\x1c\x73\x61\x1b\x62\x13\x76\x5b\xb1\x59\xa8\x18\x69\x23\x10\x11\x71\x18\x61\x66\x46\x03\x43\xd7\x58\x1c\x65\x86\xa2\x23\x4d\x4c\xa4\x02\xff\x6b\x40\x88\xc6\xb8\x15\x41\x7a\xb6\x73\x85\xcd\xbf\xb7\x7b\x4f\x9f\xf5\xbb\x2f\x25\x31\x37\x12\x7e\xde\x00\xc2\x0c\x11\x24\x71\x5c\x73\xfb\x2f\xfa\x37\xf7\x72\x89\x9a\xe6\x44\x68\xfa\x3f\x61\x58\x8d\x50\x96\x18\xc3\x36\x92\x39\x9c\x3b\xee\x88\x42\x2a\x5f\x40\xa5\x0d\x01\x9f\x20\x4a\x0c\x82\x40\xcc\xa5\x87\x61\xb6\x31\x6f\xb9\x86\x13\xfe\x27\x0e\xf8\x13\xe4\x37\xe8\x9a\xc4\x6a\x09\x27\x87\x8b\x2b\xb6\xb6\xf4\x3f\x76\xfb\xff\x1c\x17\xd5\x24\x4d\xdd\x43\x35\x3f\x12\x47\x56\xa9\x99\xb1\x8f\x52\x13\x13\xca\x93\x3f\x28\x0a\xc4\xa2\x16\x48\x20\x36\xa0\xc4\xb4\x3a\xf8\x58\x23\xd8\x8d\x6e\x0e\x3e\x1b\x2f\x49\x31\x2a\x35\x9b\x86\x22\x3e\x0d\xdd\x1e\x2d\xb3\x3c\xaf\x65\x9b\x2a\xd2\xd2\x69\x76\xae\x9a\x4d\x84\x3c\x56\x21\x85\x8a\xb3\x57\xc1\x6e\x84\xc3\x65\xff\x69\x94\x74\xda\x56\x6d\xd0\xf8\x3a\x71\x15\xc7\x46\x42\x8d\xb6\xc7\x51\xef\x63\x71\x68\x18\x86\xdd\xf1\xe3\x2f\xf4\x5b\x77\xf4\x07\x5f\x96\x2f\xe5\x86\x18\x89\x3c\x74\x42\x8a\x3c\x78\x8d\xc4\xec\x9d\x97\x90\x47\xaf\x8f\x70\xec\x0d\xb3\xf6\x86\x65\xb1\x37\x0a\xd2\xb4\x74\x05\xfb\x0b\xe5\xcb\xe5\x4a\x35\x65\x84\xc6\xd1\xf2\xad\x3b\xaf\xb6\xb6\x8f\xb7\xae\x7f\x7f\xb8\xa3\xdf\xda\xeb\x7f\xd5\xfd\xfe\xf0\x46\xce\xd9\x5b\x26\x73\xa6\x11\x03\x7d\x25\x13\xdc\xf4\x51\x64\x48\x98\x47\x8d\x92\x42\x8d\xd9\x80\xf0\xbd\xab\x4d\xe5\xa4\x73\xf9\xf5\x15\xa6\x08\x6d\xbe\xf6\x9d\x8d\x2c\x74\x46\x58\x5a\xcc\x59\x2d\x7d\x58\x8d\xcc\x10\xdb\x0d\x59\x12\xea\xa0\x5c\xa9\xe6\xfd\x87\x59\xc1\x05\xa4\xeb\x4f\xf4\xc3\xad\x37\x2a\x39\xf0\xd7\x31\xbd\x53\x8f\xab\x64\xf2\x8a\xf5\x08\xe0\x98\x96\xe5\xca\xa5\xd2\x87\xd6\x96\xd9\x4c\x9e\x31\xb6\x6a\x06\x85\x3a\x00\x4b\x45\xda\x61\x75\xf5\x18\x7d\x90\x4a\x13\xdc\xdd\x3f\x3a\xd8\x3f\xea\x76\x53\x21\x61\x64\xa4\x05\xb7\x23\x72\x23\xec\x45\xe7\x9d\x3f\x32\xdd\x0e\x27\xe3\xe1\x8b\x77\x8e\x14\x8a\xef\x72\x34\xcd\x4d\xea\x20\x29\x33\xf8\x38\x93\xbb\x84\xa8\x7b\x82\x5c\x21\x77\x92\x6b\x82\x71\xed\x4d\x4d\x20\x78\xcd\x13\x2f\x5c\xc1\xc2\x99\xcc\x7a\x4e\x9e\x08\xa7\x42\x41\x40\x9a\x56\x63\x64\x1d\xf9\x25\x4b\xd0\x78\x53\x91\x6b\x81\x63\xd5\x7e\x97\x86\x28\xcf\x9e\x27\x44\x47\x66\x8e\x4d\x75\x9d\xcc\x03\xfe\xb6\xd2\x30\xb4\x8b\x8b\xcf\xe2\xaa\x02\xbd\x27\xbb\xbd\x4f\x9f\xba\xab\x02\xfe\xc2\x42\x5c\x45\x40\x6d\x6b\xad\xef\xb6\xb6\x55\xfa\xdd\xd6\x36\x26\x0c\x39\x32\xdb\x7d\x55\x9c\x14\x52\xe9\x02\x70\xb4\xb0\x4d\x76\x5b\xb8\x18\xe6\x48\xd6\x46\x08\x79\x26\xe1\xcd\xa8\xb0\x57\x69\x2d\x61\x07\x2d\x3b\xf3\xc2\x5a\x90\x09\x46\x35\x8e\xe8\x19\x33\x49\x18\xe2\x70\xa4\xc3\x94\xa7\x94\x93\x3a\x9f\x16\x21\x83\x81\xe4\xd6\xd8\x13\x6d\xb4\x24\x25\x53\x5f\x19\xf4\x93\x71\x62\xcb\xef\x44\x15\x6a\x1a\x9f\x0f\xb1\x24\xac\xf3\x8d\x09\x3b\x90\xc1\xc4\x82\x5e\xfa\x84\x46\x26\x4d\x09\x73\xa7\x96\x75\x28\x30\xa9\x83\x22\x8f\xd3\xb1\x68\x96\xe9\x69\x30\x1d\xf3\x99\x38\x53\xae\x5c\x29\xad\x54\x8d\x9c\x6c\xd9\x02\xa1\xb5\xfd\x13\x67\x7e\xbd\xb0\xb8\x5a\xba\x62\xb4\x99\x48\x00\xe0\x2d\x28\x2a\x12\x7e\x6b\xca\xfa\xaf\x68\x7f\xe7\x9a\x64\x5d\xc2\x5a\xce\xf9\x17\x15\x0b\x50\x14\xce\x9e\x9f\x3d\xdf\x80\x73\xe7\x61\xa3\x81\xe6\xde\x43\x85\xb9\x9f\xae\xcd\x14\xd6\xce\x9d\x3b\x8b\x06\x86\xfd\x6f\x77\xfb\xb7\xff\xa0\xdf\xb9\xef\xb4\x0c\x1c\x05\x40\x65\xf9\x37\xf9\xc9\x60\x73\xd1\xf8\x9a\xe4\x59\x15\x00\x61\xa0\xae\x19\xa2\xb2\xee\xac\x27\x1e\xdb\x06\x14\x5b\x04\x27\xd6\xa6\x8c\x23\x49\x81\x92\x9c\xf4\xbb\x13\x00\xea\x0d\x89\xb2\x96\x08\x37\xe3\x18\xd8\xa4\x4d\x05\xd2\x75\x8b\x1d\x06\x53\x41\x54\xb1\xe3\x32\xaa\x7c\xe5\x58\x99\xaa\x34\xc9\xc8\x92\x16\xd1\x56\x41\x32\x36\x5d\x7f\x62\x5a\x71\x69\xe6\x94\x3c\xcc\xaf\x8e\x53\x09\xa0\x74\x11\xef\xb8\x0f\x8a\x4a\xa2\x8c\x09\xe0\xdb\x7a\x94\x6e\x24\x3f\x79\xf8\xd9\xf1\x7f\xef\xf4\x3e\xff\xdd\xab\x07\xb7\xbd\x0f\xa3\x6b\x28\x59\x3f\xf5\x6e\x3c\x3e\xde\xfd\x73\xd8\xe3\x22\xd9\xc0\xa1\xbb\xee\xca\x99\x9e\xdf\x3c\x7e\xf1\x42\xdf\xd9\xd7\x3f\xd9\xf6\x0e\xc1\x23\xc5\x92\x54\xb1\xd7\x88\x44\xca\x35\xdf\x60\x64\x30\xda\x14\x88\x90\x77\x99\x80\x53\xd1\xfc\xe0\xe4\x38\xdb\x52\x16\xe4\x7c\x02\x1a\x80\xba\xa2\x8d\x0d\xac\x14\x41\x91\x29\x9c\xc6\xce\x95\x5d\xe8\xb5\x46\xc1\xb1\xe2\x28\xba\x8c\x50\x76\xc9\x06\x5e\x19\xc1\xc0\xd6\x6f\x27\x61\x29\xb7\x79\x94\xbf\xbd\x3a\x3d\xba\x6c\x6a\xf0\x99\x8a\x70\xf2\xcd\xb1\x63\xcd\x7d\x3b\x1c\x48\xa2\xe2\x6d\x53\xa4\x51\x01\x31\x3c\x48\x75\x23\x70\xba\x46\x91\xd6\xe2\x33\x46\xd7\x54\x89\x22\xad\x26\xe1\x68\x3f\x17\xfc\xc6\xfc\x40\x1c\x4e\x86\xff\x91\x74\x51\x69\xae\x32\x50\xdb\x36\xa1\x13\x56\xf4\x1e\x0b\xf5\x5a\x80\xca\x3b\x7f\x64\x1c\x46\xc3\xca\x01\xa4\xd9\x1d\xf7\xa1\xc7\x31\xb7\x7c\xe6\x2c\xba\x19\x3a\x1b\xca\x6d\x1e\x3b\x2b\x51\xf3\xab\x95\xd0\x59\xda\x95\xcf\xb8\x5f\x1c\xbb\xcb\x9e\x1c\x12\xef\xa4\x5b\xb4\xa9\xc8\x35\x55\x6e\x37\xa5\xd1\x99\xce\xeb\xce\x28\xa7\x47\xd7\xb9\x38\x4a\x8c\x29\x8a\x8c\x63\xf0\x59\x50\x1b\xb9\xbc\x16\xa2\x90\x06\x9a\x48\x6b\x0d\x64\x0e\x30\x3f\xd3\x40\x63\x48\x9d\x07\x5a\x0b\xc9\xf2\xcf\xb5\xd6\x55\x0c\xec\x8f\x5d\x2c\xaa\x40\x05\x19\xdd\xc3\x0e\x1c\x68\x71\x7e\xa4\x30\x3f\x6c\x32\x46\x46\x42\x8b\x00\x63\x5c\x02\x36\x08\x95\x45\xe7\xd5\x05\x97\x80\x72\x84\x53\xda\x15\xcd\x38\xcb\x68\xc0\x84\x35\x19\x06\x11\xab\xf2\x76\xe1\x90\xb6\x53\x58\xd9\x59\x67\x31\x2a\xd9\x40\xd4\x5c\x43\xec\xfa\x86\x66\x11\xab\x0c\x33\xc0\x91\x3d\x29\x57\x5c\xcc\x60\xc5\xb3\x03\x74\x26\xc1\x53\x6b\x5d\xa5\x51\x08\xf5\xf6\xe1\xb0\x46\x09\xab\x6d\x36\xef\x6e\x01\xa0\x45\x34\x36\x0f\x64\x22\x40\xd9\xfc\x73\x66\x06\x80\xb7\xfb\x5f\xef\xe9\xb7\x1e\xe9\x2f\xaf\xeb\xbb\x5f\x7b\xcd\x8d\x3c\x6c\x1e\x50\x42\x98\xb7\x5d\x85\x9a\x36\x0f\x8a\x33\x67\x67\xe7\xde\xf3\xf6\x6c\x10\xba\xae\xa9\x50\x40\xf3\xe0\x5d\xe3\xb9\x77\x19\xd2\x18\x78\xfb\xe8\x79\xb7\xff\x55\x57\x7f\xfa\xaf\xde\xfd\x3d\xf0\x93\x0f\x96\x97\x4a\xa0\x7f\xf7\xcb\xde\xce\x7f\xfa\x0f\xbf\xd5\x5f\x7c\xea\x1e\x63\xb8\x25\x74\x74\x18\x86\xd1\x2c\x43\x8a\x2a\x43\x86\x7e\x24\xda\x74\x44\x3b\xf0\x9b\xdb\x6d\x1e\x9f\x84\x2d\xdd\xbf\xe2\xd0\xe5\x99\x53\x8c\x63\xf0\xcb\x44\x86\xb8\x39\x08\x13\xbb\xf4\x3e\x0f\x8a\xef\x14\xae\xe2\x0e\xa4\x46\xe0\x9a\xbf\xe8\xb4\x22\xd8\x08\x20\xd8\x44\x76\x34\x7b\xd9\xaf\x2d\xc9\xa2\xd5\x8e\xfd\x0d\x46\xd4\x0d\x29\x80\xf3\x90\x68\x92\x69\x73\x88\xe9\x62\x80\x54\x87\x23\x1b\x1f\x84\x3b\xf3\x89\x8c\xd3\x24\xd6\x63\x40\x81\x12\x7e\xa7\x49\x46\x9d\xcc\x8c\x7f\x32\x46\xfc\xa5\x9f\x8b\x19\xb5\xd3\x9d\xab\x38\x3d\xd7\x9f\x02\x5c\x06\xc7\xe3\x12\xec\x20\xfc\x46\xa2\xe5\x63\xd8\x81\x59\xe2\x45\xe9\x60\x20\xc8\x08\xe2\x60\xb3\x64\xac\x71\x78\xc8\xa5\x9f\x60\x16\x18\x32\x26\x62\x83\xe8\x07\x01\xa1\x41\x1e\x52\x51\x95\x37\x12\x40\x58\x55\xb2\xc4\x0f\x56\x15\xef\x73\xa9\xa7\x11\x40\x89\xa9\xae\xdc\x0f\x72\x4f\xc4\x7e\x12\x44\x23\xe5\x7f\x01\x00\x00\xff\xff\xe1\x5b\xfa\x4b\x8c\x3f\x00\x00")

func mysql000001_gokinsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		_mysql000001_gokinsUpSql,
		"mysql/000001_gokins.up.sql",
	)
}

func mysql000001_gokinsUpSql() (*asset, error) {
	bytes, err := mysql000001_gokinsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mysql/000001_gokins.up.sql", size: 16268, mode: os.FileMode(438), modTime: time.Unix(1628584317, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sqlite000001_gokinsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x48\x28\x89\x4f\x2a\xcd\xcc\x49\x49\xb0\x06\x04\x00\x00\xff\xff\xaa\x8a\x92\x63\x1f\x00\x00\x00")

func sqlite000001_gokinsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		_sqlite000001_gokinsDownSql,
		"sqlite/000001_gokins.down.sql",
	)
}

func sqlite000001_gokinsDownSql() (*asset, error) {
	bytes, err := sqlite000001_gokinsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sqlite/000001_gokins.down.sql", size: 31, mode: os.FileMode(438), modTime: time.Unix(1628565625, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sqlite000001_gokinsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x5d\xdd\x8a\x24\xb7\x15\xbe\x1f\x98\x77\x10\xe5\x40\xaf\xcd\x8c\xb7\xfa\x77\x67\xca\x04\x32\xf1\xf6\x3a\x43\x76\x67\xd7\x33\xbd\xc6\x86\x85\x46\x5d\xa5\xae\x96\xa7\x4a\xaa\x48\xea\x9e\x99\x07\x48\xc8\x5d\x12\x48\xc0\x90\x2b\x93\x9b\x04\x83\x9d\x5c\x25\x18\x42\x5e\x26\xe3\x79\x8d\xa0\xaa\xea\xff\xea\x6e\x95\x4e\x75\x76\x49\x1a\x6c\x16\x49\xe7\xb4\x7e\xce\xf9\x74\x74\xf4\xa9\xe7\xf1\x07\x87\x07\xe8\x02\x4f\xa8\x8f\x15\x7a\x25\x48\x4c\xc7\x31\x7a\x8a\x15\x46\x3d\x81\x99\x1c\x12\x71\x78\x70\x78\x80\xae\xf8\x58\xf8\x04\x5d\x11\x31\x21\x02\x4d\x3f\x1e\x0a\x06\x6b\x95\xbd\xbb\x84\x64\x95\x57\x9f\x3e\xa7\x8a\xac\x35\xf8\x8c\x08\x49\x39\x43\x1e\x6a\xba\x4d\xd7\x75\xeb\x0b\x2d\xfc\x11\x89\xf1\x82\xfe\x18\x53\x96\x76\xa0\x87\x45\x48\xd4\x96\xef\x58\x6e\x50\xf4\x1d\xcf\x68\x44\x50\x97\xf9\x3c\xa0\x2c\x5c\xf8\x8e\x4e\x3b\xad\x3f\x3c\xd0\xe3\x26\x1e\xaa\x77\x1e\xbb\x27\x8f\x1b\x6e\xa3\x8e\xea\x2d\xaf\xdd\xf2\xdc\xe6\xe1\xc1\x07\x8f\x75\x8b\x57\x97\x67\x9f\xbc\x38\x43\x43\x2e\x08\x0d\x59\xff\x9a\xdc\x49\xf4\x63\x34\xc4\x91\x24\x1f\xe9\xfa\xe3\x63\x74\xbc\xe5\x93\x36\xe8\xe1\x41\x44\x90\x54\x62\xec\xab\xb1\x20\x5a\x19\x92\xe9\xb8\xfb\x31\x0d\x05\x56\x94\x33\x69\xa0\xea\xe9\xe5\xcb\x57\xa8\x77\xf6\xd3\xe7\x5d\x74\xfe\x0c\x75\x3f\x3f\xbf\xea\x5d\x21\x67\x4d\x93\xf3\xd1\xe1\xc1\xc7\x97\xdd\xb3\x5e\x37\x6f\x5c\xd0\x04\x3d\x3a\x3c\x40\x26\x1f\x67\x92\xcd\xab\x83\xc6\x94\xa9\x4e\xeb\xc8\x54\x2e\xa0\x42\xdd\x39\x68\xc0\x79\x74\x78\xf0\xbe\xf1\x5c\x5d\x12\x9f\x8b\x40\x22\x3e\xb4\x9a\xa2\xf3\x8b\xab\xee\x65\x0f\x9d\x5f\xf4\x5e\x16\x0e\xfb\xb3\xb3\xe7\xaf\xbb\x57\xe8\x51\xfd\x08\xb9\xe6\x9d\x2a\x5a\x40\xd5\xc7\x42\xd1\x21\xf6\x55\x3f\xc1\xfe\x35\x0e\x89\xf5\x0a\xae\xab\x5a\x5f\xc2\x82\x36\xc6\x6b\x88\x1c\x1a\x38\x68\x82\x85\x3f\xc2\xe2\x51\xa7\xf5\x3e\xba\x78\xd9\x43\x17\xaf\x9f\x3f\x37\x5d\x4d\xe4\x60\xad\x82\x32\x45\x42\x22\x2c\xc4\x05\x49\x78\x7f\xb5\x17\x4f\xbb\xcf\xce\x5e\x3f\x2f\xab\x8a\xe1\x98\xcc\xf5\xd4\x5d\xd7\x56\x51\x40\x65\x12\xe1\xbb\xfe\xb2\xc2\x46\xbb\x6d\xad\x90\x48\x7f\xae\xa8\x6d\xdf\x33\x5f\x10\xac\x48\xe0\xa0\x00\x2b\xa2\x68\x4c\x2c\xf5\x8c\x93\xa0\x12\x3d\x01\x89\x48\xaa\x87\x32\xf5\xa8\x6e\x3f\x3d\xa9\x96\xbe\xee\x08\xb4\x4b\xaf\x2e\xcf\x5f\x9c\x5d\x7e\x81\x7e\xde\xfd\x02\x3d\x4a\x8d\xf3\x7d\x4b\x94\xb1\xf2\xe3\x4a\x60\x23\x47\xd6\x2a\x60\x63\x0a\xd2\xdb\x60\x63\x06\xe4\xff\x9f\xb0\x91\x2f\xee\xbb\x06\x42\xb3\x55\xa9\x40\x97\x1c\xe1\x6a\x80\xb1\x2a\x1c\x4b\x04\x99\x50\x72\x03\xc4\x8d\x77\x0c\x0d\xf7\x02\x3d\xe6\x58\x50\x05\xf4\x70\x71\x57\x01\xea\x70\x71\xb7\x0d\x70\x74\xb5\x11\xd6\x80\x60\xc6\x1e\x61\x9c\x31\x14\x0a\x1c\x2e\x42\x30\x9e\x38\x34\x20\x4c\xd1\x21\x25\x62\xd1\xe7\x2c\x14\xad\xc4\x31\x36\x6e\xab\x63\x22\x6d\x39\xeb\x5b\x7d\xcd\xad\x19\x29\x90\xe9\xb1\x0e\x38\x92\x0a\x00\xc8\x89\x78\xc8\x81\x51\x5d\x05\xc8\x53\x01\xe8\x80\xa3\xaf\xaa\x02\xaf\x7d\x00\x9f\x19\x12\x41\x30\x6f\x30\xa6\x51\x00\x40\xbb\x54\xbe\x08\xe7\xb2\x8a\x1d\x08\x97\x62\x5b\xaf\xfb\x79\xcf\x10\x9a\x9c\x84\x26\x24\xa2\x2c\x8b\x53\x52\x41\xf3\x65\x9a\x0b\xe7\xdb\x89\x95\x12\xa9\xb0\x1a\xcb\xf2\x72\x44\x08\x2e\x2c\xc4\x26\x84\xa9\xf2\x62\xda\x82\xfb\x52\xe1\x38\x71\xd0\xd3\xb3\x5e\xb7\x77\xfe\xa2\x5b\x4e\x5e\x45\xa4\xfc\xd7\xc6\x44\xca\xf4\xec\x6d\x31\xab\x22\x75\x62\x9b\xce\x0e\x29\xa3\x72\x64\x2b\x3d\x43\x31\x1b\xe1\x19\x7e\xd9\x08\xcf\x82\xdb\x92\xb3\xb5\x8c\x34\x20\xa0\x31\x75\xff\xa5\xc4\xd1\xcc\xb9\xa7\xe9\xa2\x5a\xa7\x5e\xc7\xee\x89\xdb\x1e\x0e\x9f\x0c\x3a\x4f\xda\x2e\x76\xd3\x4f\xb3\x76\x84\xb2\xba\x27\x64\xb8\x52\x57\x9f\xd5\xad\xcb\x35\x74\x5d\xea\x30\xfa\x1f\xf9\x7f\xe9\xbc\x2c\xfd\xbf\xd6\x70\x1b\xf5\x63\xf7\xe4\xb8\xde\x41\x6e\xc7\x6b\x9e\x7a\x99\xd6\xf5\xe2\x46\x71\xf1\xb6\xd6\x35\x3d\xa3\x26\xc3\x3e\x6d\x35\x56\xba\xef\x1b\x0d\x7b\x5d\x6e\x60\x35\xec\x56\xcb\xab\x9f\xac\x0f\x44\x17\x9f\x16\x17\x6f\x6b\x6d\x3c\x6c\x7c\x12\x2c\x77\xbf\xde\x36\x1a\xf6\xba\x5c\xcb\x6e\xd8\xa7\x5e\xdb\x2d\x18\xc8\xa9\xd7\x2e\x58\xd6\xd6\xa9\xd7\x2a\x9a\x8d\x69\x6b\xf3\x61\x9f\xfa\x2b\xdd\x27\x66\xc3\x5e\x93\x0b\xac\x86\xdd\x76\x3d\xb7\xb5\x3e\x10\x5d\xdc\x2e\x2e\xde\xd6\xda\x7c\xd8\x3e\x5e\xee\x7e\xe3\x89\xd9\xb0\xd7\xe4\x3a\x4b\xc3\x16\x24\xe1\x88\x08\xe1\x8d\x94\x4a\x3c\x34\x66\xe4\x36\x21\xbe\x22\x01\xea\xbe\x7c\x86\x04\xc1\xe9\xd5\x84\x12\x98\x46\x44\xe4\x92\xfd\x90\xa8\xbe\x16\x34\x9c\xaf\x22\x33\x69\x37\xbc\x7a\xf1\x7c\x6d\x6d\x6d\x3c\x5f\x83\x66\x67\x79\xdc\x4d\xd7\x68\xbe\xd6\xe4\x1a\x43\x3b\x33\x69\x78\xcd\x02\x37\xd7\xc5\x05\x6e\xb0\xab\x75\x0d\x7a\x17\xe0\xc7\x41\x5f\x47\x5e\x80\x40\x73\xaa\xa2\x28\xd6\x9c\xd5\xed\x3e\x50\xdb\x9f\xa6\x9d\x50\xf0\x71\x02\x3b\xd0\x3a\xa9\xb1\x00\x75\x48\x45\xa0\xdd\x98\x06\xb1\xd6\xc7\x50\x87\x8d\xe3\xfc\xc4\x55\xf6\xc8\xe5\xf8\x3c\x20\xf6\xb2\x4c\xa5\x01\xb1\x22\xb7\xca\xa0\x39\xec\x84\x3a\x0f\x4b\x2d\xe5\xe7\x91\xa9\x9d\x02\x58\xa4\x57\xec\x8b\x79\x8c\x0e\x70\xc5\x69\x94\x5f\xe0\x89\xb3\x03\xc0\x2e\x47\xb4\xf6\x43\xbb\x8c\x16\x2c\x9b\x35\x3d\x0e\x59\xa7\x4b\xe6\x76\x1b\x71\x16\x9a\xd8\xae\xa3\xee\x12\x02\xf0\x50\xa0\xe9\x3b\x94\x0d\xb9\x34\x73\x33\x67\x2c\x22\x40\x3a\xaa\xca\xbc\x89\xb9\x6d\x43\x3c\x88\x8b\x10\xe0\x3d\x5c\x84\x45\x9e\xa3\x8b\xb7\x7a\x8d\x95\xc7\x94\xf7\x16\x7b\x4f\x81\x64\x58\xf3\x9c\x66\xa9\x83\xb0\x93\x8c\x07\x11\xf5\x1d\x1d\x89\x2d\x26\xfe\xdc\xad\x42\xf6\x7e\x01\x48\x56\x6e\x4c\x52\x6e\xef\x2b\x30\x33\x09\xf4\xac\x8d\xc6\xdf\x4f\x68\x02\xd9\x3f\xa6\x2a\x36\xb8\x41\x56\x67\x10\xca\xd9\xdd\x6d\xc0\xaf\x25\xd2\x44\x22\x50\x05\x34\x32\xb1\xb1\xfd\xf4\xb3\x17\xa3\x48\xb0\xc0\x31\xc0\x22\x52\xf9\x22\x73\xc8\x2a\x76\x65\x91\x2d\x0c\x01\x76\x45\x0d\x0c\x08\x9c\x00\x2b\x6c\xb2\xb9\xa6\x79\x5c\x69\x65\x24\x55\x6e\xab\xa6\xcb\x0b\x32\xa1\x3c\x39\x0f\xb1\xa2\x5c\x45\xa1\x21\x4d\xeb\xf6\x7a\x44\x84\xdd\x99\xae\x6e\xa0\xa5\x2f\xe5\x2a\xe3\x6c\xcd\xaf\x4a\x74\x18\x0a\xd2\x04\x05\xba\x4d\x3b\xa7\xc9\x8d\x6b\x25\x57\x7b\xfb\x38\x83\xcd\xa6\xd7\xe7\x6c\x58\x81\xc5\xa7\x7a\xb6\x99\x7d\xd6\xc0\x90\xdb\x04\xa1\x25\x2d\xdf\xd0\xd9\xb3\xa3\x96\x0f\x15\x96\xac\x43\x07\xfb\x3e\x91\xb2\xaf\xf8\x35\x61\x70\x6d\x77\x71\xd4\x2f\x7d\x90\x9b\x0e\x47\x12\x51\x09\x93\x72\x3f\xfb\xf7\xec\x66\x14\x8b\x2a\xcc\x71\x82\xc5\x56\x6b\xd4\xf5\x66\xc6\x08\xb0\x45\x38\x83\x65\x8b\x35\x97\x57\x55\xc5\xea\x3b\x13\x1c\x8d\xe1\x5a\x04\x89\xb1\xb8\x96\x60\x3d\xd3\x20\xd4\x8e\x0c\x53\x6d\x88\x52\xce\x82\xab\x71\x16\x30\x43\x75\x55\xd5\x76\xa7\x79\x1b\x0c\x55\xb8\x0f\xa1\x3c\x5d\x3c\x20\xc2\x41\x03\x1a\x6a\x63\x69\x58\x73\x26\x53\x3e\x84\xac\x9e\x0f\x0a\xe0\xb5\xcf\x16\xa8\x32\xa6\xfc\x4c\x63\xe5\x1c\xfc\xea\xf0\x0c\x15\xd2\x73\x01\x5d\xb3\xdc\x59\xab\xa4\xc0\xce\x62\x4d\x45\xd9\x9d\x25\xa8\xa1\x8c\xd9\xdb\xaf\xcc\xe2\x53\x92\xb7\x1f\x71\x46\xfa\x55\xc4\x44\x7b\x8a\x68\xa5\x82\xdd\x29\xa4\xf2\x45\xf0\x97\x55\x98\xf0\xc8\xca\x03\x5d\x31\x23\xcc\xce\x2d\xc0\x57\x7b\x6b\x97\x72\x65\x73\x12\x39\xc3\xcc\x32\x11\x0f\x8b\x4f\xaa\x39\x79\x82\xee\xdb\x60\x97\x6d\x20\x0c\x01\xd1\x58\x1d\xc9\x85\xb2\xb9\x0f\x75\x72\xd7\xb0\x9c\xec\xfd\xa0\x00\x49\x40\x20\x40\x92\x62\x0c\x20\xc9\x0e\x08\xb0\x43\x00\xa8\xd3\x66\x4b\xb0\x24\x5f\xce\x6b\x2b\xf1\x9b\x0a\x61\xcc\xc9\xa6\xda\xb2\x1f\x30\x04\x9b\x72\x5d\x6d\xa5\x6f\xa9\xea\x5b\x72\x0b\x60\xd8\x09\xca\x25\x43\x30\x0f\x04\x79\x10\xc4\x83\x00\x1e\x34\x68\xd4\x6b\x45\x43\xc6\x85\xd5\x42\xfb\x3c\x8e\x31\x0b\x4c\xae\xb6\x9d\x1b\x4c\xf5\x51\xe3\x4b\xc9\xd9\x8e\x25\xb4\x43\xf0\x7d\x20\xb0\x12\x34\x0c\x09\x24\x89\x93\x6b\x28\xc2\xe1\x69\xd5\xff\x18\xbb\x03\x9c\xb2\x04\x93\x35\x80\xef\x25\x57\x5e\x04\x95\xa7\xa7\xa4\x17\x3c\xb9\xa9\x97\x93\x24\xac\xf8\x41\xd4\x7f\x83\xa3\x02\x7b\x3c\x54\x65\xf6\xc9\xdc\xeb\x2a\xf0\xed\xbe\x18\x43\x72\x4e\x0b\x5a\xb6\xf8\x78\x5a\xfd\x4e\xbf\x4f\x54\xe0\xa7\x85\xe9\x35\x3e\xfc\xe4\x37\x9d\x86\x8c\x31\x55\xda\x85\x50\x41\xfc\xf1\xb6\xde\xe3\xed\xc1\x27\x0c\xad\x15\xe2\x17\x63\x09\xda\xf0\xb4\x78\x91\x27\xa4\xe5\x7b\x39\x75\x94\x37\x7a\x58\x88\x99\x60\x69\x9d\xec\x77\x18\xf5\xaf\xad\xbf\x19\x4f\xb0\xc2\xf6\x81\x35\x24\x44\x8d\x78\x48\x99\xf5\xcd\xaf\x83\x7d\x45\x27\xa4\xf4\xdd\x46\x95\x1e\x64\x68\xd7\x2b\xef\x11\x32\xb3\x9d\x3d\x47\xc0\x41\x4c\x59\xed\x08\xd5\x8f\x50\x2d\xe4\xd7\x94\xc9\xf4\x41\x41\xdd\xc5\x81\xdf\x3c\x6d\x9d\x0e\x70\xfb\x14\x0f\x06\xa4\xdd\x21\x6e\xfb\xc9\xb0\xe1\x0e\x4f\x4e\x9a\xe9\x9b\x96\x87\x6f\xbf\x7e\xf8\xed\xaf\xee\x7f\xf7\xd5\xec\xc1\x41\xbd\xd3\x38\x69\x9f\xb4\x3a\x27\xee\xda\x5b\x01\xf7\xd4\x6b\x3f\xa9\x55\xf0\xb3\x41\xba\xfb\x7d\x8d\xa5\x40\x8f\x4e\x75\x6c\x72\xeb\xac\xd2\x60\x7b\x03\x6c\x6e\x4e\x32\xe2\x0c\xf6\x2b\x18\x0e\x89\x31\x8d\x40\x8f\xd6\x9d\x01\x15\x6a\x14\xe0\x3b\xeb\x5d\x21\xbf\x2d\x34\xa4\xdc\x23\x27\x21\x22\xce\x8d\xd0\xea\x45\x76\xa6\x20\xa5\xc3\x02\xe4\x33\x0e\xa1\x95\x82\x0a\x9f\x57\x96\x31\x66\xb0\xd3\xc4\x12\x42\x4b\x9e\xaa\xd8\xe8\x32\xba\x6e\x7f\xa4\x4c\x20\x73\x2a\x96\x50\x4e\x27\x94\xa7\x24\x08\x0e\x54\x0c\x79\x69\x92\xa6\xce\x56\x93\x08\xbb\x09\x9d\x76\xdc\xe2\x25\xd1\x0a\xf9\x51\xf0\x1d\xcf\xd0\x8e\xc1\xee\x02\x63\xf1\x4f\x55\x6c\x74\x97\x9d\x7c\xfe\x6c\x05\xde\x8e\xbb\xc0\x29\xd0\x60\xfe\xb2\x06\x69\x1c\xc4\x16\x0c\xe6\x4c\x56\xdc\x58\x8b\x92\x5b\x62\xc3\x9c\xce\x84\x03\x7e\xc3\xec\x36\x96\xca\x1d\xc5\xcc\x82\xc1\x8e\x92\xb2\xe7\xa0\xae\x92\x51\xf0\x36\x39\x4b\x4e\xd0\x33\x48\x37\x58\x27\x0c\x32\x97\xb1\xbf\x8d\x77\x96\x99\xb0\x16\x3f\xdc\xe3\xf0\x84\x30\xfb\xab\xa2\x5c\x09\x9c\x0a\xb2\x72\xa2\xb3\x52\x01\x3b\xd7\x4d\x95\x2c\xb1\x33\x75\x78\x59\x5a\x85\x20\x43\x41\xe4\x08\xa4\x83\xdc\x26\x54\x10\xd9\xa7\xac\xd0\x3e\x76\x03\xc3\x82\x0e\xd8\x6f\x08\xcd\x87\x03\x53\x03\x4e\x04\x39\xe9\x7c\x1a\xbe\xff\xcb\x9c\xcb\xfc\xbd\xe0\x3e\xa0\xd0\x14\xa3\x20\x60\x78\x17\x47\xfd\x24\x1a\x87\x14\x02\x86\x73\x25\x45\x60\xb8\x50\xbb\x5f\x30\x5c\x46\x11\x8b\x7c\xa7\x35\x19\x1a\x44\xe8\xaf\xec\xd7\xba\xaa\x34\xc0\x52\x76\xb1\x92\xac\x59\x5c\xf0\x85\x1f\x5f\xae\xc9\x51\xed\x08\xd5\xb2\xbe\x1e\x23\xa9\x48\xe2\x21\x39\x22\x51\xf4\x13\x39\x7a\xc3\xa6\xa3\xc8\xc9\x03\x17\x38\x26\xba\x7a\x5e\xc1\xd6\x4a\xa6\xf7\x9e\xde\xbc\x48\x6b\x26\xfe\x88\x23\xad\x97\xa3\x1b\x2e\xa2\xa0\x76\x84\xdc\x2c\xdb\x53\xf0\x4b\x17\x45\x9d\x6d\x1c\xa1\xda\x00\x6f\xee\xae\xae\xdb\xd0\xe1\xe5\x2a\x56\x50\xb6\xaf\x4e\x37\x8f\x50\x2d\xe1\x37\x44\xa4\x7d\xdc\xd4\xf5\x79\x8b\x0d\x03\x28\x6a\xc0\x36\xd6\x98\x0f\xa6\x5e\x6a\x30\x2d\x6d\x2e\x5b\xec\x45\x8e\xde\x88\x4d\x26\xb3\x5c\xc7\x8a\x0a\x29\x4b\xc6\xca\x5b\x2c\x41\x68\xc4\xa5\xf2\x50\xc4\x7d\x1c\xa5\xff\x6c\x34\x10\x7a\xef\xe1\x9b\xef\xee\x7f\xf3\xa7\xfb\x7f\xfd\xf2\xfe\xeb\x6f\x96\x9b\x6b\x84\xf6\x90\xe0\x5c\x2d\x97\x27\x58\x4a\x0f\xd5\x1b\xcd\x56\xbb\xb3\x5c\x73\xc3\xc5\xb5\x4c\xb0\x4f\x3c\xf4\x58\xcb\x3d\x56\x44\x2a\xf4\xde\xbf\xff\xf1\xfd\xc3\x5f\xbe\xbf\xff\xeb\xdf\x7e\xf8\xea\x3b\xf4\xa3\x9f\xbd\x7c\xd1\x45\x0f\xbf\xff\xf3\x0f\xbf\xfe\xfb\xc3\x1f\xbf\xbd\xff\xe7\x1f\x16\x75\xcc\x67\x5b\x94\x31\x1e\xe8\xfe\xa0\x48\x9c\x44\x58\x41\x98\x90\x8b\x6a\x36\xed\x11\xb3\xfa\xfd\x3f\xa2\x00\xef\x14\xa0\x87\x33\xc0\xdd\xa2\xc2\x9f\xd5\xae\x7a\xc7\x28\x61\x29\x05\x20\x30\x37\x80\xc5\x5d\xe3\x13\x1e\x61\x16\x6a\x24\xc8\x6f\x35\x3d\x54\xff\xd0\x7d\xc3\x26\x58\x68\xd8\x49\xc9\x6a\x19\xfe\x68\x8c\xc0\x21\xc9\xb1\x68\x19\x96\xc7\x34\x0a\xb2\x72\xb6\x5a\xa0\x81\x65\x0e\x60\x86\x1b\x53\xc8\x8f\x53\x15\xc7\xf5\x35\xb4\x9f\x6b\xd6\x1f\xc2\x26\xde\x4e\xbc\x0c\x79\x26\x96\xfe\x25\x8c\x0f\x43\x6e\xdb\x99\xc6\x6a\x67\x34\xc4\x94\xef\x4b\x0a\x4c\xc7\x93\x37\x6c\xe7\x26\xb4\xbe\x64\x7a\xef\x7c\x81\x27\x84\xbd\x6b\x2b\xf6\x25\x9e\xe0\x2a\xd7\x2c\x9e\x30\xe4\x47\x04\xb3\xf5\x62\xaa\xc7\x38\xdf\x26\xcb\x77\xb0\x8a\x75\xd4\x1d\xc9\x17\xd2\x62\x19\x75\x34\x71\x91\xc4\xef\xda\x22\xb2\x24\xae\x72\x0d\x59\x12\x2f\xcb\x95\xee\xc6\xda\x4a\xa5\xcf\xc7\x16\x05\x8d\x3b\x92\x4b\x22\x9b\xcd\xfb\x9c\x05\xe4\x96\xc8\xd5\xed\x3b\xdb\xd2\x2d\xfe\xa6\x4a\xbe\x3f\xbf\xbe\x38\xff\xf4\x75\x17\x9d\x5f\x3c\xed\x7e\x3e\xe3\x38\xf6\xc7\x8c\xfe\x62\x4c\x9c\x6c\x83\x79\x79\x61\xf3\xb7\x66\xe6\x7c\xc9\xb3\xab\x8f\xe7\x4d\xb3\x01\x17\xff\xfd\x1d\x25\xc6\xe4\xa3\xc3\x83\xff\x04\x00\x00\xff\xff\x87\x69\x43\xc0\xcb\x68\x00\x00")

func sqlite000001_gokinsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		_sqlite000001_gokinsUpSql,
		"sqlite/000001_gokins.up.sql",
	)
}

func sqlite000001_gokinsUpSql() (*asset, error) {
	bytes, err := sqlite000001_gokinsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sqlite/000001_gokins.up.sql", size: 26827, mode: os.FileMode(438), modTime: time.Unix(1629279967, 0)}
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
	"mysql/000001_gokins.down.sql":  mysql000001_gokinsDownSql,
	"mysql/000001_gokins.up.sql":    mysql000001_gokinsUpSql,
	"sqlite/000001_gokins.down.sql": sqlite000001_gokinsDownSql,
	"sqlite/000001_gokins.up.sql":   sqlite000001_gokinsUpSql,
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
	"mysql": &bintree{nil, map[string]*bintree{
		"000001_gokins.down.sql": &bintree{mysql000001_gokinsDownSql, map[string]*bintree{}},
		"000001_gokins.up.sql":   &bintree{mysql000001_gokinsUpSql, map[string]*bintree{}},
	}},
	"sqlite": &bintree{nil, map[string]*bintree{
		"000001_gokins.down.sql": &bintree{sqlite000001_gokinsDownSql, map[string]*bintree{}},
		"000001_gokins.up.sql":   &bintree{sqlite000001_gokinsUpSql, map[string]*bintree{}},
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
