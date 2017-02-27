// Code generated by go-bindata.
// sources:
// schemas/AppDefinition.json
// schemas/Group.json
// DO NOT EDIT!

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

var _appdefinitionJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x7d\xfb\x6f\x1b\x37\xf2\xf8\xef\xf9\x2b\x06\xfe\x04\x70\xdc\x93\x65\xa7\x87\x7e\xbe\x68\x80\xc3\xc1\x75\xdc\xd6\x68\x12\xfb\xfc\xb8\x7e\x0f\x71\x2e\xa1\x76\x47\x12\xcf\xbb\xe4\x96\xe4\x4a\x56\xdb\xfc\xef\x1f\xcc\x90\xdc\x97\x56\x0f\x3b\x4e\xdb\x3b\x5c\x50\xd4\x92\x96\x8f\xe1\x70\x38\xef\xe1\xfe\xf2\x04\x60\xe7\xa9\x4d\xa6\x98\x8b\x9d\x17\xb0\x33\x75\xae\x78\x71\x70\xf0\x2f\xab\xd5\xbe\xff\x75\xa8\xcd\xe4\xc0\x7f\xfc\x9f\x9d\x01\x35\x17\x69\x2a\x9d\xd4\x4a\x64\xe7\x46\x17\x68\x9c\x44\xbb\xf3\x02\xc6\x22\xb3\xc8\x0d\x52\x1c\x4b\xc5\x4d\xe8\x77\x9a\x02\x60\xa7\x10\x6e\x7a\xb5\x28\xb0\xfa\x05\x60\xc7\xf9\xef\x3b\xd6\x19\xa9\x26\x3c\x7a\x6c\xeb\xd0\x28\x7a\xf4\xcf\x67\x37\x37\x07\x7f\x7d\xf6\xec\xe6\x66\xf8\xcb\x97\x1f\xf7\x7e\x7d\xf6\x56\xec\xff\x7c\xb8\xff\xf5\xbb\xf0\xf7\xe6\x66\x7f\xf8\xee\x8b\xf8\xe3\x9f\x1a\x0d\xbe\xd8\xdb\x7b\xf6\xf4\xd7\x9b\x9b\x83\xbd\xbd\x3f\x3d\xad\xc7\xce\xa5\x7a\x85\x6a\xe2\xa6\x3b\x2f\xe0\x39\xff\xf8\xf1\x09\xc0\x47\x06\x5c\x69\x57\x03\x2c\xb2\xec\x6c\xbc\xf3\x02\xde\x86\x9e\x11\x6a\x80\x1d\x83\x3f\x95\xd2\x60\xda\x78\xca\xbf\x27\x79\xba\x53\x7d\x7f\x17\x3e\x7d\x1c\x6c\x3f\x80\x30\x13\xdb\x33\xc2\x93\xf8\xcd\x83\x59\x34\xd1\x1e\xa1\x4d\x12\x2c\x1c\xa6\x17\x68\x75\x69\x12\xbc\xd0\x59\xe3\x31\xc0\x8e\x74\x98\x37\x7f\x58\x46\x7f\x17\xde\xea\xb9\x30\x46\x2c\x6a\x0c\xa6\x68\x13\x23\x0b\xda\x5f\x7a\x7a\x56\x78\x62\x18\xc2\x11\x64\xd2\x3a\xd0\x63\x30\x01\x0a\x30\x04\xc6\x10\x5e\x0b\x23\xdc\x54\x2b\x48\xb4\xb2\x32\x45\x63\x41\xab\x6c\x51\xb7\xd3\xe3\x31\xfd\x38\x97\x6e\xea\xfb\x80\x54\xe0\xa6\xd2\xfa\x21\xc7\xda\x40\x26\x4a\x95\x4c\xa5\x9a\x80\x13\xf6\xd6\xd2\x34\xdc\x40\x14\xc5\x10\x4e\xc7\xb0\xd0\x25\xa4\x1a\x94\x76\x60\x0b\x4c\xe4\x78\xc1\xcf\x07\x7d\x93\x8b\x2c\x5b\x37\xb7\x9b\x0a\x07\x53\x31\x43\x18\x21\x72\xbf\xb1\x9c\x94\x06\x53\x18\xd1\xa0\x08\x1f\xf6\xf7\x53\x1c\x8b\x32\x73\xef\x23\xe2\xdf\xc7\xf1\xde\xf3\x18\x1f\x20\xd1\x79\x2e\x54\x0a\x99\x54\x08\xe3\x4c\x4c\x18\x4a\xa5\xb7\xea\x3c\x17\x16\x26\x72\x86\x0a\xb4\x02\xeb\x84\x71\x65\xb1\xed\x4a\x86\x70\xa5\xc1\xe0\x44\x5a\x87\xa6\xee\x43\x38\x14\xbc\xc0\x01\xe3\x4a\x21\xa6\xe0\x74\x03\x59\xbc\xae\x1c\xad\xb6\x0c\x46\xcf\x12\x1a\xd0\x54\x38\x9f\x0b\xe5\x68\x1c\x61\xad\x9c\xa8\x16\x40\xbc\x49\x02\x6c\x46\xa8\xa4\x26\x8d\xe9\x13\xa1\xa0\xb4\xd8\x45\xa7\x9f\x58\x98\x49\x99\xa3\x72\x30\x9f\x62\x98\x91\x36\xbe\x2c\xb8\x39\x8f\x57\xcd\xcf\xeb\x10\x90\x6b\x83\x30\x96\x0a\xf7\x27\x46\x48\x85\x69\xb5\x6d\x82\xe8\xb3\x77\xd2\x0a\xce\xc6\x8c\x0d\x8c\x54\x78\xb5\x53\x61\xd0\x42\x81\x86\x17\x30\x84\x4b\x44\x78\x4b\x83\xbc\x26\x6c\x81\x70\xce\xc8\x51\xe9\x10\x08\x5b\xf5\xea\x53\x9d\xf0\xa0\x0c\xc1\xbb\x67\x81\xb9\x32\x86\x87\xa2\x10\xc9\x14\x99\xbb\xb6\x9a\x1d\x64\xc2\xa1\x75\x07\xd5\x98\xb6\x06\xf3\x60\x8f\x77\x31\x45\x27\x64\x66\x87\xfe\xc4\x86\xf3\xea\x79\xc7\x6f\x71\xe0\x8f\x14\x70\x03\xda\x5c\x3f\x52\x38\x31\x06\x0b\x83\x16\x95\xb3\x20\x88\x10\x88\x87\x0b\x27\x67\x08\xb9\x4e\x91\x9b\x7b\xcc\xf2\x21\x9e\x62\x45\x5f\x4e\x83\x29\xd5\x10\xae\xe8\x3c\x13\xe9\xe7\xda\xc9\x99\x70\xfe\xc0\x59\x31\x46\x28\xad\x98\xf0\x10\x89\x56\x8e\xf6\xd7\xc8\x9f\xd1\xc0\x18\x85\x2b\x69\x6f\x32\x79\x8b\x20\x20\x29\xad\xd3\x39\xbc\xd4\xc9\x2d\x1a\x38\x79\x73\x75\xf1\x8f\xf3\xb3\xd3\x37\x57\x61\x6c\xc2\x11\x8c\x25\x66\x29\xe4\x62\x01\x23\x1a\x17\x53\xe2\x35\x45\x26\x12\x3f\x7e\x9e\x02\xd2\xc1\x63\xd2\x2b\x6d\x84\x35\x10\x68\x05\x33\xde\x61\x52\x3a\x6d\xc2\xc8\xc9\x54\xa8\x09\x42\x2e\x8d\xd1\xc6\xc2\xd1\xf9\x29\xd3\x82\xc5\x5c\x28\x27\x93\xd8\x20\xf0\xb5\x48\x38\xc7\x7e\xb0\x53\x35\xd6\x50\x18\xed\xf4\xa8\x1c\x43\x8e\x96\x17\x5b\x91\x3d\x73\xa6\x19\x1a\x2b\xb5\x82\x0f\x87\xc3\x2f\x0f\x87\x87\x1f\x86\x00\x27\xd2\x4d\xd1\xc0\x87\x24\x4f\x3f\x80\x36\xf0\x81\x96\xf7\x01\xf2\xd2\x3a\x5a\x9a\x2d\x8b\x22\x93\x98\x0e\xe1\xd4\x81\xa4\x99\x67\x22\x93\xfe\xcc\xd3\xa3\x05\x8c\xb4\x9b\x86\xee\x04\x6c\xe8\x1f\x20\xb4\x22\x47\x66\xae\x6d\x32\x1b\x89\xe4\x56\x8f\xc7\xdf\x8a\xc4\x69\xd3\xa4\xb7\x5c\x2a\x99\x97\x39\x89\xd5\xe1\xe1\x12\x49\xa9\x32\x1f\xa1\x59\x49\x53\xc7\x91\xc5\x5a\xc0\xbb\x42\x2b\x54\x4e\x8a\x0c\xc2\x64\x30\xc2\xa9\x98\x49\x6d\xfc\xa6\xd4\x92\xa0\xd0\xce\xb7\xcc\x16\x60\x65\x72\x4b\xf0\xda\xb0\x23\x85\xa1\x6d\x74\x16\xac\x50\xe9\x48\xdf\xa1\x25\x16\xa5\x13\xc9\x74\xc5\x38\x25\x36\x4a\xbb\x28\x67\x98\x2d\x60\x2c\x64\x56\x8b\x97\xb1\xd1\x39\x8c\x65\x96\x35\x18\xcf\x54\x98\x14\x52\x69\x6f\x89\x0f\xfa\x1d\x64\x5e\xc4\x53\x62\x05\x6d\x81\x46\xea\x94\x50\x9e\x97\x99\x93\xbc\x09\x51\x72\x8c\x19\x6f\x7c\x8c\x51\x24\x2d\x10\x18\x80\xd2\x20\x94\xca\xc9\x0c\x24\x1d\x28\xe2\x11\x16\x72\x71\xf7\x8a\xd7\xfc\x12\x33\xb1\xb8\xc4\x44\xab\x34\x2e\x53\xf0\x26\x93\x20\xb0\x9a\xb6\xd6\x43\xcf\xe7\x51\x18\x84\x5b\x99\x65\x98\x42\x5a\x32\x03\xae\x96\xa8\x35\xe4\x42\x2d\x60\x8a\x22\x23\x44\x4c\x31\xb9\xed\x32\x94\xb0\x9c\x30\x5d\x73\xab\xff\xbb\x79\x7f\xc4\xcd\x1b\x2c\x9f\xc4\xe5\x73\x28\x95\xc3\x09\x9a\xf6\x4e\x93\xde\xba\x7a\x7b\xaf\x9a\x9c\x9a\x40\x93\x36\x70\x3f\xe2\x2d\x7e\x21\x33\x91\x95\x48\x0f\xe6\x46\x14\x85\xc7\x98\xc7\xf1\x4c\x0a\xf8\x70\x30\x92\xea\xc0\x4e\x61\x3f\x81\xa7\xbf\x10\x4f\x49\xf2\xf4\xe3\x6f\xc9\xc2\x96\xf0\xd0\x35\x39\x7a\xcc\x82\x88\x1c\xad\xac\x23\x95\xc2\xd9\x3e\xdb\xa5\x23\x2b\xd7\x48\xdd\x76\xcb\x46\xdb\xa6\x09\xf0\x4b\xe3\xf3\x4a\x89\xdd\x92\xda\x9b\x7b\x0d\xda\x4f\x51\x31\x6d\xbc\xdd\xb9\x7e\x73\xfa\xb7\xeb\x93\x9d\x01\xec\x1c\xbf\xba\xbe\xbc\x3a\xb9\xa0\x8f\xdf\x5d\x9c\x5d\x9f\xbf\xff\xe6\x1f\xf4\xf9\xd5\xe9\x0f\xfc\xf8\xfa\x4d\xfc\xf4\xfa\xe8\xff\xbf\x3f\x3f\xb9\xd8\x79\xf7\x30\x48\x9a\xbd\x6a\x33\xa7\x81\x92\x5c\xaa\xd3\x80\x95\x2f\x1b\x3f\xd7\x26\x67\x7c\xca\xd6\xe6\x92\x02\xd3\xa1\xdc\xbf\x33\xb1\xd4\x1b\x08\x64\x37\x09\x47\x02\x9a\x8e\x97\x56\x2c\xf1\x3d\x1a\x06\x10\x90\x30\x80\x88\x82\x01\xd0\xb2\x07\xe0\x97\x3f\x80\xb0\xf8\x0e\x9b\xac\x14\x92\x26\x75\x6c\xb2\x91\xb9\x51\x8f\x1d\xe7\x57\xc1\xda\x4b\xeb\xb7\x2d\x87\x5c\x37\xac\x07\xd6\x60\xea\x99\xee\xd2\xb3\xc6\x6e\xe9\xd1\xbf\x30\x71\x1d\xba\xe9\xe1\xfb\xd5\x60\xac\xdc\x97\x6e\x4a\xdf\x12\xe1\xd0\x73\x68\x56\x9b\xbc\x26\xe6\xcd\x11\xb3\x18\x2e\x0f\xba\xed\xba\x36\xad\x2d\x3c\x97\x2a\x91\x45\xef\xf2\x60\xe3\xd1\xe8\x5f\xe6\x79\x1c\xf3\x1e\xab\x5c\x1a\xf6\xe3\xf2\x4c\x3b\x16\x13\x83\xee\x11\x21\xbd\xe4\x01\x3f\x09\xcc\x27\x1b\xc0\x6e\xb9\x2e\x9a\xf8\x86\x77\x4f\xd6\x74\xdc\x19\x6b\x93\xe0\x79\x99\x65\xa7\xb9\x98\xe0\x3a\xe2\x1b\x69\x9d\xa1\x50\x1b\xa9\xcf\x8b\xa5\x70\xf6\x60\x2e\xb3\x8c\x24\x46\x51\x92\xc8\x1c\xd0\x1a\x85\x49\x33\xb4\x16\xe4\x98\xc4\x31\x49\xdb\xcc\xa0\x48\x17\x20\x66\x42\x66\x62\x94\x11\x07\x60\xbc\x64\x3a\x11\x19\xd8\x85\x75\x98\x77\x90\xd2\x5d\x87\xdc\x04\xfe\x8a\xed\x6a\x0b\x98\x6d\xd6\xa6\x48\x78\xb1\x93\xa3\xda\x38\x9e\x9c\xb6\xb7\xb4\xb8\x01\x4e\x85\x6e\xae\xcd\xed\x43\x20\xed\x83\xc5\x8f\x46\xda\x07\x99\x72\x03\xef\x79\xa9\xb1\x6f\xa7\xba\xcc\xd2\xc0\x5e\x11\xa4\x1a\xc2\x99\x67\xae\xdf\x5c\x9c\xbe\xfc\xee\xe4\xd7\xef\xcf\x2e\xaf\x7e\x7d\x73\xf6\xe6\xe4\xd7\xeb\x4b\x92\x33\xdd\x39\x2b\x99\xe4\xdb\x93\xa8\xa1\x2e\xf4\x97\x7a\xb1\x10\xba\xec\xc8\x9d\xe5\x45\x17\xc2\x88\x1c\x1d\x9a\x3e\xee\xb0\x52\x0e\xf7\x2f\xfb\x28\xcb\xf4\x9c\x16\x2c\xcc\x48\x3a\x23\xcc\x02\xea\xe1\x69\x0f\x88\xd8\x84\xb5\xde\x89\x12\x36\xe8\xf8\xd5\xe9\x10\xde\x68\x87\x41\x87\x53\x0b\xe7\x55\xdd\xaa\x21\x23\xce\x5b\xa0\xd2\xb2\xb3\x6a\x52\x0a\x23\x94\x0b\xde\x98\xa0\xf4\x68\x43\x7a\x6e\xae\x67\xd4\x7b\xac\xcd\x5c\x98\x74\x00\xc2\xc2\x9c\x6c\xcc\xc9\xd4\xd1\x33\x04\x31\x17\x0b\xaf\xec\x36\xa8\x84\x80\x58\x5e\xde\xb2\x62\xb2\x84\x98\x15\x6c\xff\x9e\x3c\x7a\x33\x97\x06\xd8\xb9\xc5\xc5\x8a\x47\xdb\x71\xbe\x4d\x07\xaa\x7f\x4f\x7f\xc0\x45\xe5\x36\xac\x36\x73\x99\x09\xf6\x72\x6b\x80\x1d\x56\x72\x3f\x15\xec\x65\x15\xa5\xc4\x2d\x81\xda\x4a\xa8\xb4\xf9\x33\xe1\x79\x10\x21\xef\xb0\xe8\xee\x88\x4b\xa7\x49\x1b\xf7\x5a\x14\x85\x54\x93\x87\x9c\xa7\x3f\x16\xc1\x55\xbc\xea\x5c\x9b\x55\x22\xb7\xc7\x4c\xda\x6e\x13\x2f\x70\x1c\x98\x02\x1d\x43\xc2\x1b\x7f\x60\xeb\x2e\x61\xaf\x1e\xbb\xb1\x51\x71\x1b\xc9\xbe\xdb\xc8\xda\x2b\xc0\xa2\x8d\xa3\x83\x3f\x9d\xed\x99\xe0\x75\xe2\x7e\x87\x43\xf8\xb6\x61\x79\xd6\xcb\xf1\x02\x5e\x04\x23\x4c\x8f\xe1\xb0\xf6\xfa\x7a\xaf\xac\x6d\x4f\x75\x1e\x21\x64\x03\xc9\x77\x13\xbe\x8d\x6f\x8f\x29\x4c\xb5\x75\xd4\x2e\x58\xaa\x64\xf3\xb1\xf7\x8e\xcd\xf5\xd2\xe2\xb8\xcc\xbc\x4b\xb9\x28\xa2\xc5\x9a\xce\x68\x0f\x82\x7b\x35\x62\x61\xc1\xaa\xb6\x5f\x3e\x31\x33\x12\xb9\x1e\x4f\xba\x74\x8c\x88\xb9\x36\x59\xca\x63\x9d\x7f\x79\xce\x96\x66\xa9\x02\xda\x86\xf0\xa3\x74\x53\x5d\xba\x36\xec\x2f\xe0\xd0\x0f\x3d\x67\xb1\x83\xc6\x68\x85\xba\xb4\xd9\xa2\x0d\x84\x34\x50\x18\x76\x24\x36\x64\x15\x03\x36\x9f\xca\x64\x4a\x8b\x2a\x6d\xc9\x2b\x22\x4e\x5c\x9b\x8c\x1e\x17\x78\xc7\x1e\x4c\x7a\x3c\x93\x56\x92\xc2\x40\x58\xe1\x11\x96\x39\x6c\x20\x8d\x5c\xdc\x05\xe3\xfb\x7f\xbf\xfa\xea\xcf\x5f\xad\x66\x5e\xd1\x44\xdf\x9a\xfd\xc4\x1d\x79\x7c\xea\x25\xcc\xf8\x25\x3b\x23\xe2\xc1\x83\x1c\x05\xef\x98\x54\x15\x39\x0d\x6a\xc4\x09\x30\x42\xa5\x3a\xf7\xf8\xac\x64\x90\x61\x37\xa8\x54\x49\x56\xa6\xde\xb3\x5a\x7b\x3b\xdb\x71\x0a\xef\x76\x31\x68\xcb\x8c\xbd\x9c\x15\x6a\x6d\xed\x60\x71\xc2\xde\x32\xfd\xe0\x5d\xa1\x49\x82\xce\xa4\xf0\x70\xd2\x83\xe0\x07\x8f\x93\x5c\x9c\x5c\x5e\x55\x8e\x57\x9e\x35\x1e\x82\x39\x8e\xe0\xfa\x74\x58\xd1\x74\xf3\x98\x0d\xe1\x54\x05\x2d\xc5\xbb\xa9\xa5\x6b\x9f\x3a\xd2\x1f\x33\x1c\x3b\x28\x95\x77\x5f\x7b\x7f\x84\x02\x52\x48\x7c\x17\xa1\x9a\x0f\xeb\x69\x52\x8d\x5e\xc6\x8b\x8c\x54\x4c\x87\x20\x1a\xe8\x12\x01\x2d\x1e\x1b\xbf\x3d\x3d\x65\x62\x84\xd9\x2a\x86\xb9\x0d\x97\xee\x23\x26\xef\x0c\x17\xaa\xf2\xad\x93\x31\x92\xa6\x90\xa3\x13\xa9\x70\x22\xe8\x38\x44\xa1\xa6\x30\x18\x5c\xfc\xf1\xa4\x35\x59\xa6\x05\x5b\x26\x53\x3a\x8b\x63\x69\x70\x2e\xb2\xcc\xae\x44\xd2\x0a\x71\xd1\xbf\xb0\xb5\x6e\x89\x06\xd2\xb6\x46\x24\x69\xea\x8f\xac\x16\xbc\x69\x28\xff\x16\xcd\x4c\x26\x9e\xf7\x60\xea\x6d\x16\xd2\x15\x98\x39\x9f\x8e\xa1\x30\x7a\x26\x53\x32\x7a\xa4\xab\x3c\x67\xa5\x92\x3f\x91\x30\x98\xa1\xe1\x90\x1c\x93\x5d\x1e\xa4\xf9\x4a\x3c\x36\x03\xef\x6f\xc5\xfe\xcf\x31\xca\xbe\xff\xee\x8b\xa7\xdb\xeb\x49\x1c\xc9\x48\xf4\x2a\x1b\x7c\x6b\x9c\x54\x46\x81\x4b\x0a\xd2\x62\xca\x34\xfe\x19\xd0\x2f\xef\xb6\x43\xe5\x79\x80\x26\xa2\x93\x31\xf1\x2c\xf8\x7e\xde\xee\xba\xa4\xd8\x1d\xc0\x6e\x99\x16\xbb\xef\x40\x1b\xfe\x44\xc3\xef\x32\x17\x1a\x69\x37\xdd\x1b\xc2\xcb\x06\x43\xe0\x1e\x3d\xa6\xf3\x2a\x6c\x84\xed\xfb\x1c\x9c\xfb\x94\xd8\xf0\x14\xb3\x22\x8a\x35\xea\xae\x88\xef\x72\x98\x50\x13\x5b\x8d\xd4\x93\x4a\x9b\x10\x35\x2c\x42\x4c\x4b\xc0\x1c\xb3\x6c\xff\x56\xe9\xb9\xf2\x9d\x69\x94\xd0\xda\xf3\xe6\x4a\x17\x68\x2c\xa1\xf6\xf4\x12\x5f\xa3\x13\x7e\xd0\x39\xcb\x15\xd7\x95\xce\x62\x36\x86\x51\xe9\xbc\x49\x13\xb8\x01\x73\x85\xd1\x02\x32\x2d\x52\x18\x89\x4c\xa8\x84\x4c\x5b\x35\x36\xc2\x3a\x53\x26\xae\x34\x21\xc2\x7a\x19\x40\x7f\x59\x81\xfe\x8a\xfa\x7c\xc3\x7d\x68\x0d\xa9\x4e\xa0\x10\x93\x00\x6e\x13\xca\x4a\x91\xde\xa4\x51\xbd\x92\xb7\x58\x31\xec\x01\x1d\x27\xa2\x91\x6a\x91\x87\x83\x8e\xa0\x8b\x7e\x86\x88\x1b\x3e\x81\xa2\x1f\x43\x15\xfe\x9a\x58\x71\x53\xac\xe5\x69\x6d\xfc\x05\x1d\x4a\xba\x86\x2f\x3d\x9c\x61\x91\x18\x6d\x83\x06\x97\x95\xd6\x45\xd1\xc9\xed\xbc\xb0\x0c\x10\xc6\xbd\xf6\x52\x94\xe4\x66\x10\x8d\x5e\x2a\xbf\x65\x3f\xc7\x7b\x7a\xfa\x3e\x97\x6a\x00\xcd\xef\xe2\xee\x1d\xcc\xa7\x68\x10\xda\xad\x18\x69\xed\x86\x3c\x72\x2b\x37\xc0\x63\x38\xe4\x50\xc4\x48\x69\x80\x4f\x8f\xe1\xf9\xe1\xe1\xe1\x21\x0f\xf4\x25\x7d\x1a\x90\xd4\x2f\x30\xf1\xa1\x9b\xcf\x2b\xf9\x36\x7b\xb8\xd6\xdb\x40\x46\xce\x64\x86\x13\x4c\x1f\xc3\x65\x75\x51\x06\xde\xdd\xf2\xe7\x48\x05\xf5\x34\xac\x50\x74\x7d\x3b\xab\x9c\xed\xab\xf2\x88\x2a\x3f\x55\xe3\xb7\x16\xb7\xec\xca\xf6\x27\x3d\xc3\xef\x88\xa2\x48\x3e\x8b\x13\xfa\xf7\xf1\xa1\x1d\x15\xc5\xf1\x7d\x3c\x68\x72\xed\x96\x6f\x03\xe2\xff\xdb\xe8\x63\x52\x01\xa0\xd3\x97\x5e\xa5\xf6\xa3\x46\x88\xc7\xda\xe4\xc2\xc1\xee\x54\xd8\xe9\x3e\x9f\xa6\xdd\x41\x38\xa3\xfc\xdb\x2e\x75\xf2\xa1\x4d\x3b\x05\x91\x4d\xb4\x91\x6e\x9a\x7b\x0e\x4b\xa7\x6d\xd7\x77\xaa\x9a\xe1\xdd\x3e\xaa\x44\x93\x80\x48\xe5\x04\xad\x1b\xc2\x71\x69\x0c\x2a\x97\xf9\x00\x27\x27\x65\x15\x68\x72\xe9\x88\x9b\x77\xc6\x95\x16\xec\x54\x7c\xf5\xfc\xcb\x0d\x98\x5b\xa9\x58\xde\x37\xc0\x10\x53\xcb\xc0\x8f\x38\x84\xcb\x72\x42\x60\x63\x1a\x7e\x79\x01\xbb\x21\x2f\x82\x64\xb8\xb6\xbb\x03\xbf\x6e\x61\x92\xe9\xee\xf6\xb1\x86\x35\x2e\x8a\x7e\x0d\x71\x3d\xd3\xf8\xb7\xf5\x75\x6f\xc5\x67\xe2\xf9\x7d\x00\x47\x71\xed\xfc\xcf\x3e\x44\xb7\xc6\x5c\xce\x33\x08\x18\x40\x35\x21\xb1\x43\x7d\x89\x28\xa2\xbf\xb4\xf1\xb3\x05\xe1\x0d\xfa\x5c\x73\x5a\x17\x89\xac\x97\x67\xc7\x3f\x9c\x5c\x30\x85\xbc\x3e\xb9\x3c\xbb\x6c\x13\x48\xad\x6e\xf2\x43\xd2\x34\x7d\x87\x86\xd3\xb9\xb9\x98\x99\xce\xca\x7c\x89\x7c\x56\xb8\xbd\xee\xe1\xd0\x5a\xef\xcc\x6a\x38\xb2\x04\x33\x99\xcd\xb4\xdb\xe3\x2f\xeb\x21\xaf\x42\xb8\x69\x64\x3c\x7e\x69\x51\x79\xa8\x03\x9c\x3d\x23\x2d\xc7\xcd\xeb\x7f\xcb\xc1\x23\x56\xb4\x3e\x3b\xdc\x81\xd2\x69\xb2\x4f\x07\xb9\x20\xfe\x62\x1d\xaa\x7e\xed\xfd\xf1\x3d\x95\x56\xfe\xbc\x85\x35\x79\x3f\x43\x81\xd5\x63\xf9\x73\x25\x0b\xeb\x45\x35\xf6\xfa\xb5\xfc\x66\x8d\xaf\xfd\x7e\x4a\x56\xcf\xca\x57\x31\x88\x15\x3d\x76\xa2\x4f\xe0\x3f\x08\xed\x95\x9b\xe3\x31\x90\xfe\x5b\xb9\x22\xba\x7a\x94\x87\xfd\x5e\xde\x80\x99\x4c\x97\xf2\x19\x3e\x0f\x58\x50\xcd\xb6\x35\x7c\xc1\x76\x79\x64\x3f\xd8\x79\x80\x63\x3f\xb8\x05\x93\x08\x60\x2b\x2b\x39\x1a\x4e\x0f\x0f\xf7\xdc\xf7\x54\x91\x79\xf1\x48\xdc\xf7\x5c\x5b\xef\x14\x0f\x86\x1e\x89\xd8\x8b\x33\xb6\x47\x2f\x50\xa4\x67\xa4\x4b\x92\xac\xbd\xf8\xb1\xfa\xed\xe0\x47\x23\x1d\xf6\x8d\x5d\x09\xdf\x8b\x33\x92\xbc\x17\x3f\xee\xdc\x2f\x58\xb5\x1a\x0f\x1f\x7b\xd5\x14\x1f\xb4\x7a\xd2\x1d\x7c\x39\x17\xba\x39\x60\x95\xf2\x53\x94\xbd\xb9\x60\x1b\x92\x5c\x99\x64\xb9\x09\x11\xed\xf1\xf9\x75\x4c\x2c\x77\x55\x2e\x60\x88\x16\x29\xc4\xd4\xe7\x9b\x4b\x65\x9d\x50\xde\x23\x23\x6d\xec\x5e\x39\x96\xa7\x21\xb3\x3e\xb8\x54\x27\x68\x06\xec\x70\x09\x1e\x58\x01\x63\x23\x12\x8e\xa4\xf4\xa6\x07\xb6\x56\x95\x62\x81\x2a\x45\x95\xb4\xb9\xe4\xb6\x59\xe1\x55\xf1\x47\xf0\x42\x58\x28\x0b\xad\x42\xc0\x60\x69\x85\x7e\x32\x3b\x84\x23\x05\xda\xa4\xde\x51\x93\xa2\x91\x33\x4c\x1b\x11\xed\x06\x48\x4c\x44\x05\x1a\xb2\x8a\xd8\xb1\xe5\x84\x71\x07\xd6\xe9\x82\xa9\xac\x2c\x26\x46\xd4\x61\xb5\xc6\x5c\x21\x6c\x76\x27\xf2\x22\xc3\x01\xa7\xa6\x37\x00\x39\x10\x60\x90\x93\x30\x83\xda\x50\x81\x7f\x30\x8a\xd1\x0e\xef\xc8\xaa\x9b\x1d\x24\x5c\x63\xc1\x10\xb0\x77\xf5\xcf\x2d\xc7\xf5\x00\xc6\xd2\x58\x07\x07\x09\x9b\x4c\xd4\x0a\x39\x75\x52\xd1\x98\xfe\xaf\x18\xae\xcd\x19\x7c\x6a\x70\x4c\x58\xfd\x9f\x83\x46\x31\xd5\x41\x55\x43\xf5\xa4\x49\xb5\xd5\xfe\x49\x7b\xfb\x00\xaa\xfc\x5e\xcf\x21\x2f\x93\xa9\xcf\x8d\xb5\x85\x48\xbc\x60\xfa\x86\x9d\x7c\x88\xd1\x9d\xd8\xdd\xc2\xcd\x14\x29\xd4\xa7\x13\x25\xaa\x59\x73\x4d\xc1\x41\xbd\xc2\x7a\xdb\x19\x7e\xd1\x55\xc7\xb5\xc2\x56\x15\x97\xff\xf7\xcb\x12\xbb\xeb\x32\x94\x4f\x35\x5d\x8f\x14\xa0\x9a\x49\xa3\x15\x9b\x1f\x33\x61\x24\x5b\x61\x36\xe4\x67\x81\x4f\xfe\xea\xb1\x51\x37\xa5\xb8\x3d\x7a\xd6\x18\xb1\xa5\xdb\x98\x19\x81\x01\xb0\xb0\x85\xde\x95\x60\x63\x70\xb9\xdf\x57\xb7\x56\xa3\xbe\x7f\x4e\x59\x58\x20\x2c\xf9\xf7\xd7\x2a\x7d\x2b\x2d\xd9\x77\x0f\x64\xf2\xb1\xc0\x63\x9b\x82\xc5\x1e\x8c\xc6\xee\xc1\xd9\x44\x7f\x7c\xea\x7b\xcf\x41\x7a\x29\xc7\x63\x34\x44\x27\xb1\x97\x97\xa6\x95\xf1\x1e\x1c\xdc\x92\x38\x18\x31\x59\x85\xf0\x2c\x46\x39\xa3\xbb\x55\x8e\x41\xd1\x03\x19\x2a\xc7\xf6\xe8\xd3\xc1\x41\x92\xa7\x31\x6a\xeb\xc4\x2d\x06\x3f\x72\xee\x1d\x44\x21\x8d\x3b\x38\x9f\x23\x03\x9c\x62\x96\x41\x86\x33\xcc\x86\x2b\x6a\x32\x7f\xbd\xb9\x39\xb8\xb9\xa1\xb1\xe9\xd3\x5f\xdf\xfe\xf3\xe6\xe6\xe0\xdd\x9f\x9e\xdd\xdc\x1c\x84\x8f\x7b\x5f\xec\x3d\x6d\xa3\x73\x8c\x2e\x99\xde\x5f\xac\x04\x3d\x2a\x85\xeb\x8b\x53\x8f\x94\x3a\x97\xca\x87\x4f\x79\x60\x34\x90\xeb\xb4\xcc\xea\x0a\xac\x6c\xe6\xc3\xcf\xa6\x54\x4e\xe6\xb8\x9e\xdd\xae\x3c\xdc\x5b\xda\x19\xab\x0f\xed\x4e\x69\xe4\xb2\x37\x60\x7d\x82\x76\x07\x05\xd7\x17\xa7\xe1\x34\xfa\xa5\x36\x92\xeb\xdb\x6b\x5f\xad\xf5\xf8\x7d\x26\x4a\x5a\x0d\x4b\xbf\x2f\x6a\x39\xc5\xd4\x55\x60\x08\xe3\xe4\x58\x24\x8e\x18\x44\x63\x86\x35\x50\x38\x62\xfc\x9f\x0a\xc2\x89\x1f\x66\x19\x0c\x39\x6e\xe4\xce\xdd\x1b\x47\x89\x48\xa6\x9f\x8c\x9e\x63\x1a\xe4\xb1\x21\xd3\xa5\x2b\x4a\xf7\xad\x5c\xb7\x7b\x5b\x51\xd2\x05\xb2\xc5\xb4\x04\x5e\x95\x20\xec\x4b\x4e\xd9\x88\x6c\x42\xf3\xa4\x07\xae\x0e\xd3\x26\x32\xef\x14\x2c\xc7\x93\xef\xcb\x54\x8e\xb9\x4a\x65\xcb\x02\xc5\xcf\x70\x0e\x43\xbc\x6a\x35\x06\x7b\x05\xfb\xca\x2c\xb6\xc7\x4d\x5b\x5f\x97\x59\xf8\x50\x91\xde\xac\x0e\x6a\x56\x5a\x06\x9e\xbf\x4d\xf2\xf7\x4a\x99\xda\xa6\xcf\x89\x11\x09\x9e\x73\xdd\xd4\x72\xb1\x58\x67\x11\xfd\x0e\x94\x25\x9d\xb4\x09\x7b\xa8\xac\x0a\x01\xcd\x89\xd2\x26\x94\x80\xc5\x12\xf1\xda\xa6\xb2\x7e\xfa\xa8\xca\x70\x96\xd0\x08\x2b\x6b\x01\x53\xd0\x26\x54\x68\x35\x9e\x27\x3a\x47\x1b\xf0\xb5\x08\xba\x2e\x06\x1d\xbe\x2d\x3b\x02\xb4\xfd\x9e\x99\x36\x52\x3c\xa0\xdf\x3b\x57\x3c\xbf\xbb\xfb\x54\xb6\x72\xca\x83\xc1\xf7\x57\x57\xe7\xf0\xfc\xee\x8e\xe3\xa7\x5a\x59\xb4\xc3\xd5\x5c\x83\xc3\xf2\x33\x91\x3d\xd2\x96\xbc\x59\x42\xb2\xd3\x30\x17\xd2\xc1\x08\xdd\x1c\x51\xad\x2a\x47\xbb\x07\xce\x72\x71\x77\x5c\x97\xd4\x7d\x1b\xf6\xfd\xf1\x20\x6f\xd6\xeb\x4d\xfb\x49\x6c\xec\xd0\x54\x36\x2b\x42\xa9\x22\x59\x30\xad\x84\x3c\xf7\x51\x2c\xcd\x7b\xd0\x2a\x8b\x3e\x07\xf8\xfd\x18\xf9\xb9\x20\x86\xad\x01\x55\x5a\x68\xc9\x7a\xa3\xcf\xaf\x08\x85\x8a\x0c\x2c\xeb\x74\x1c\x2b\x0a\x4e\xb1\xb8\x66\xeb\x84\x2b\xed\x10\x4e\xbc\x49\xfc\x02\xd8\xb0\x3c\x70\xfa\xc0\x37\xf0\x29\xec\x2f\x7c\x34\xd0\x17\x48\x8f\x21\x26\xf5\xc0\x5f\xfe\xc2\x74\xf8\xf6\xf2\xdd\x1a\xe2\x2b\xfa\xb2\x5d\xee\xb9\x6f\xac\xf2\x46\xef\x99\x4f\x27\xd5\xb4\x85\x0a\x13\xfa\xc8\x39\x78\x89\xb0\x6c\xf2\xa7\x0b\x25\xf2\xd0\xcc\x0e\xc0\xa2\x4f\x82\x38\x55\x29\xde\x2d\x6f\xd2\xba\xc4\x82\xed\xb6\x30\x8e\xfd\xa9\x6b\xe4\x41\xaa\xbb\x2e\x44\x51\xec\xda\x2a\x7b\xc3\x88\x45\xd3\x0a\x23\xbe\x14\x36\x90\x44\x2f\x5a\xe7\xbd\x27\xd2\x0f\x61\x7d\x2b\xab\xa3\xfb\xa3\xba\xec\xa0\x91\xbe\x62\x07\xbe\x50\xfe\xed\xe1\x00\xe8\xbf\x77\x3e\xd5\xb2\x76\x8d\xa4\xa1\x04\x35\x89\x74\x1e\xf9\x26\xab\x08\xbc\x09\x7d\x66\x6d\xa8\xbf\x7f\x7a\x7e\x76\x71\xf5\xfc\x61\x87\x62\x55\xce\xd8\x3d\x0f\x46\x27\xd9\x2b\x62\x2a\x16\x6c\x78\xf7\x11\xa6\x55\x7d\x0a\x91\xf2\x80\xff\x7f\x39\x80\xab\xe3\x73\x12\x10\xc7\x67\xaf\x5f\x1f\xbd\x79\xb9\xb4\x8e\xca\x47\x49\xcd\xb9\x44\x85\xba\xd1\x87\xab\x63\xfe\x1e\x3a\xae\xae\x8e\xdc\x21\x71\xa2\x4b\xf7\xd9\x78\x72\x93\x7f\x89\x36\x8f\xf3\x05\x3b\x7c\x6d\x08\x5f\x98\x11\x0b\x96\x1b\xd1\xe3\x0a\x69\x5e\xba\x6c\xbb\x91\x5b\x18\xd8\x0d\x77\x6b\x54\x0b\x5b\xf9\x15\x9b\x1c\x5e\xab\xcc\xc3\x6b\x9f\x22\x25\xb9\x0e\x71\x2c\xd1\x54\xe2\x9b\x4f\x00\xad\xd7\xba\x90\x52\xc1\x59\x5a\xd2\x67\x24\x91\x8a\x6b\xc1\x62\x21\x4c\x75\xcd\x44\x26\xec\x14\x89\x2d\x8a\x64\xea\x43\x0b\x31\x99\x52\x38\xc8\x50\x58\x07\xcf\x21\x99\x0a\x32\x3e\xd0\xf0\x61\xc9\xc5\xc2\x73\xc9\x10\x23\x85\x54\x4e\xa4\xb3\xf0\xec\xc3\xe1\xfe\xd7\x1f\xf6\x06\x90\xf2\x98\xf0\xec\xc3\x3e\x7f\xd3\xfc\x6c\x48\x9f\x7d\x5a\xd5\x1c\x0d\xb3\xb0\x0c\x1d\xd7\x15\x3d\xfb\x20\xf6\x7f\xfe\xb0\xe7\xed\x7d\x0f\x83\xf0\xc9\xe9\x23\x9c\x48\x45\xd4\x89\x2a\x8d\x09\xff\x34\x7a\xa7\x1a\x35\xfa\x8b\x7b\x3d\xb8\x5d\xaa\x5a\xef\xa0\xae\x86\x6a\x5e\xb5\x53\xf9\x4d\x5d\xf0\x7f\x0e\xe1\x9c\x70\x83\x04\x24\xbe\x68\x29\x64\xc1\xdf\xe7\x6f\xbe\x48\x41\xa8\x05\xab\x53\x64\x2e\x06\xbf\x22\x0d\x92\x88\x0c\x97\x5c\xb6\x9b\x1d\x83\x4b\x49\x26\xab\x14\xf9\x25\xcf\x9c\x73\xc2\x5f\x40\xd0\xcc\x49\xe6\x32\x86\x3a\x6f\x79\x5c\x66\x5e\x57\x26\xb1\x0a\xb5\xae\x0f\x52\xf9\xac\x9c\x80\x02\xcd\xe5\xeb\xd1\x65\x3c\x20\xf6\x3b\xe7\xf6\xe8\x9d\x33\x62\x24\x33\xe9\x98\x8f\xfb\x1b\x46\x3c\xd8\xa0\x95\x9f\xf1\x59\x8b\xff\xf2\x15\x35\xcc\x79\xb9\x99\x6f\x62\x9d\x98\x70\xc2\xa6\x4a\x49\x08\xa7\x65\x12\x27\xcf\x85\xb9\xad\xdd\xd5\x5e\xf4\x4b\x03\x85\xb6\x0c\x6d\x8c\xdf\x17\xb2\xc0\x4c\x2a\xdc\x6b\x60\x75\x63\x22\xcc\xaa\xbb\x62\x5a\x7b\xd0\x7b\xb5\xc1\x43\x08\xef\xbf\x37\x48\xfc\x61\x6f\x90\x68\x6f\x38\xe6\x0f\x8c\x7b\x89\x5c\x97\x8a\xc3\x43\x39\xe6\xda\x2c\x42\x88\x21\xde\x27\xd1\x8a\x33\xb4\x2b\xa7\x5a\x51\xb0\xcd\x20\x4e\x5a\xb1\xb9\x55\x14\xb8\x1e\xc4\xef\xce\xaf\x21\xd1\xa6\x4a\xca\xbd\x3f\x7c\x2b\xc1\x93\xc5\x51\x9a\x1a\xb4\x0f\xe1\x5c\xa7\xe3\x6e\xf0\xaa\x96\x95\xb1\xe4\xc6\xfb\x77\x77\xab\x69\x76\x7d\xb1\xe9\xa0\x93\x77\xcc\x66\x41\x50\x8e\x40\xd0\x1a\xf6\xd9\x66\x38\x3d\xf7\x94\xcf\xd4\x3d\x84\xa3\x4a\x44\x7a\xf5\xf1\xa0\x59\x13\xd8\x09\x65\xfb\x24\x43\x12\x81\x56\x64\x24\xd9\xd2\xa6\xf7\x78\xd5\x55\x07\x31\xb5\x7b\x55\x5a\x58\x8f\x57\x66\x59\x87\xae\x59\x72\xa3\x30\x6d\x29\xf1\xbd\x93\xee\xb5\x2e\x43\x95\x17\xfb\x08\x35\xc4\xf1\x9a\x2c\xbf\x0a\x0b\xfe\xf9\x28\xde\x29\xe5\x35\xfd\x86\xed\x56\x15\x3d\x3d\x76\x01\xef\x52\xe8\xde\xf4\x36\xdb\x9c\x26\x13\x8e\xfa\xca\x64\x89\x7b\x27\x6d\xaf\x68\xf1\xf0\x84\x1b\x36\x52\x3c\x98\xf7\xa8\xd2\xf8\xcc\x35\x3c\x6b\x4b\x06\x7f\xab\xaa\x9b\xbe\x7a\x9a\xc7\x2f\xa2\xd9\xeb\xc7\xfa\xa7\x78\x1b\x9f\xf4\xac\x7c\x67\x62\x74\x59\x2c\xb9\x79\x57\x1f\xd0\x95\x87\x33\x5c\x6e\x00\x7e\x40\x6f\x19\x9e\x9e\x93\x12\x47\x0c\x34\x28\x85\x99\x66\xb9\x09\x3f\xa3\xd1\xa4\x85\xf3\x75\x83\xdd\x9e\x44\x7e\x32\x45\xe1\x79\xa1\x68\x54\xad\xfa\x61\xc2\x9d\x97\xad\x02\xde\x50\x34\x51\xb9\x9c\x16\x3e\x29\xa4\xb6\x39\x08\xc9\xad\x89\x3a\x7c\x6c\x45\xc2\xe9\xbd\xcc\xe6\x6e\x32\x55\x67\xbe\x8d\x7b\xd1\x9b\xf1\x7d\x1f\x1e\xfe\x03\x2e\x42\xdd\x4b\x21\xa4\xd7\x8f\x48\x3b\x07\x56\xcf\x39\x2e\x29\x5c\x05\x15\x7b\x36\xc7\x22\xc1\xb5\x1c\xbd\x6d\x78\xaf\x50\x78\x9d\x29\xb1\x77\x45\x61\xae\x37\xcb\x6c\xe1\x3e\x19\xcb\x5c\x9d\xc7\x57\x22\x0e\x20\xc5\x24\x0b\xd9\x3e\xfd\xc8\xe6\x35\xb6\x12\x6e\xa3\xbf\xf1\x5f\x5a\x2a\x76\x77\x15\x9c\x38\x91\x60\x48\xf2\x26\xbd\xd4\xfb\x7b\x4e\xcf\x8f\x5e\xf3\x3d\x14\x98\xb0\xa3\x4f\xcf\x7d\xa5\x79\xa8\xcc\x6a\x5c\x3b\x31\xec\x49\x7a\x6a\x69\x28\x5d\xd9\xf7\x80\x5b\x21\x63\x68\xc8\xb3\x89\xc6\x6d\xa0\xde\x20\x11\x13\x54\x8e\xd3\x73\x83\x9d\x5b\x9b\x9c\x44\xcb\xd1\x6e\xf1\xc3\xa5\xe8\xd0\xe4\x52\xa1\xe5\x55\xb1\xda\xda\xf2\xf6\xf9\x48\x79\x28\xb0\xf5\x2a\x1a\x72\xb1\x18\xcb\x51\x9f\xfb\xc3\xdf\x43\x45\xba\x57\xe8\xe8\x1c\x0f\x40\xc0\x24\xd3\x23\x5f\xd2\xee\x3d\x09\xcf\x42\x89\xd5\xfe\x5c\xa6\xb8\x17\xaa\xea\x1a\xa5\x5c\xc1\x0a\xf3\x31\x69\xc1\x37\x36\xb8\x46\xce\x51\x53\x23\x6b\xf8\xee\xa4\x6a\x96\xbb\x71\xe9\x5a\xed\xe5\x58\x1f\x9a\x5e\x29\x50\xfb\x85\x68\xff\x45\x6f\x7d\x7b\xfc\xb2\xe7\x96\xe6\xff\xb4\xdd\x6e\xec\x46\xed\xc0\x0c\x00\x6c\x49\x01\xcd\x9a\xba\xdf\x97\x12\x3e\x63\x70\xf4\x31\xdc\xf5\xb2\xa2\x84\x73\xde\x28\x69\x63\xb2\x13\x31\xd9\x60\x80\x54\x95\xcb\x71\x17\x62\x2d\x0b\x23\x67\x89\x62\xe0\xcc\x4d\xd1\xcc\xa5\x45\x90\xb5\x43\xb7\x67\x0a\x5e\x6a\x28\xfc\xef\xcc\x10\x4a\x55\xfd\x76\x33\x10\xd2\xb5\xa7\xe5\xfc\xa6\xd6\x3e\xd3\x40\xa2\x45\x78\x1b\x00\xe5\x32\x50\xcf\x94\x03\x44\x91\xbc\x36\x0e\xc2\xeb\xf4\x13\x06\x0a\x6c\x81\xd2\x2d\x38\x6d\x16\x93\x7e\x86\xc8\xc6\x8a\xda\xad\x0d\x51\xf2\xdf\xf3\x12\x80\x6d\x8b\xba\xd6\x96\x74\xad\x0e\x6e\xf7\x1a\x06\xf7\xd3\xae\x3e\x4f\x31\x7f\x2f\xdf\x08\x00\x6c\x6b\x51\x3c\x4e\x0c\x66\x0b\xbb\xe2\x93\xec\x89\x0d\xe5\xf7\x3d\xf1\x87\x96\xb0\x33\x28\x52\x92\x28\xf6\xb7\x49\x41\xe9\xac\xf4\x6f\x25\x4b\xa7\x29\x5a\x62\x5e\x01\x92\xe0\x71\xf3\x1a\x5b\x90\x78\xc4\xe1\x84\x8f\xdc\x4a\x0b\xbe\x7a\xcf\x69\x26\x98\x3a\x8e\xd5\xdc\xe9\x35\xfc\xfc\x11\x88\xb6\x32\x09\xe2\x01\x0e\x31\x16\x7f\xab\x7e\x77\x29\xeb\x42\xc1\xbf\x7f\x64\x6f\xdb\x48\xde\xea\xb8\xdd\xef\x17\xb3\xaf\x11\xbd\x31\x6c\xbf\x3e\x1c\xdf\x63\xca\x7c\x0a\x23\x8b\x01\xf9\x9f\x98\xbe\x45\xe5\xdb\xaa\x2f\xfd\xe9\x28\x99\x0d\xc0\xa7\xce\x15\xfb\xa2\x90\x7f\xa4\xec\x95\xee\xd9\x6c\x73\x9d\x3f\x1f\xc6\x8e\x0f\x09\x6e\xff\x7b\x45\x7d\xe1\x75\xbc\x7d\x39\x17\x59\x86\xc6\x57\x2c\x74\x36\xa4\x8d\x9e\xe7\x15\x7a\x1e\x14\xfc\x27\x7a\xb8\x64\xea\x3e\xd6\x29\xda\x6f\xb5\xb9\x20\xf6\xd7\x73\x73\xc0\xaa\x6c\xbc\x4d\xee\xc2\xe6\xed\xf4\x87\x4b\x3e\xc7\x86\xde\xf4\xf5\xd7\x5f\xb7\x15\x83\xcd\x5c\x92\x78\xc7\x33\xbb\x17\xce\x27\x24\x3a\x45\xaf\xf5\xa2\xe0\xd4\xd8\x5d\xe6\xe5\x4b\xe5\xe8\xab\x4a\x91\x96\xf9\x27\xb2\x0c\x78\x25\xac\xbb\x08\x7b\xf4\xc9\xf9\x64\x63\xd6\x38\xd9\x1f\x45\x8a\x35\x2b\xe8\x71\x22\x5f\x26\x2e\xac\xeb\x9e\x89\x3a\xed\x2c\x24\x8f\x73\x70\xd4\x4d\x31\x8f\x47\x9e\x6f\xdf\xaa\x4d\x21\x01\x29\x16\x99\x5e\xe4\xa8\xdc\x70\x93\xcc\xde\x26\x67\xc0\x20\x51\xb2\x4a\x16\x0f\x88\x9e\xfc\x58\xbf\xcf\xc1\xd7\xc0\x2f\x15\xc1\x06\xa7\x5d\x41\x8b\xf1\xaf\x98\x51\xa4\xb8\x92\x92\x6f\x07\x41\x84\x87\x5a\x37\x56\xbe\xbc\x9f\xa5\x8a\xa9\x4c\x85\x4a\xf9\xed\x31\x2c\xcf\x45\xc6\x04\x11\xa3\xe7\xd4\x97\x07\xdd\x22\x28\x62\xd0\x87\x38\x4f\x6c\x22\x32\x56\x83\xaf\xd6\x31\x90\xb5\xc4\xdf\x8b\x85\xa0\x66\x78\x64\x34\xd0\xe0\xf1\x12\x91\x91\x08\xe5\x13\x0f\x68\xdf\x63\xfa\x64\xd3\xf0\x91\x6e\xd7\xfa\x37\xe4\x14\x52\x29\xd6\x12\x06\x9d\x10\x93\x33\x8b\x6e\xe5\x03\x4f\xad\x15\x08\xe5\x83\xe7\x8a\x6f\x50\x63\x0e\xe6\x9f\xfb\xb5\x76\xf8\xef\xff\x1e\x1e\x7a\xfd\x70\xaa\x4b\xb3\xd7\x71\xfc\xad\x08\x27\x6c\x57\x2f\xd2\xba\xa0\x40\xd8\xdb\x57\xda\xba\x6f\x42\xc0\xf9\x13\x5c\x7f\x3f\xb6\xe2\x6d\x06\x13\x94\x33\x3a\x38\x70\x75\x74\xf9\xc3\xfb\x57\x67\x97\x57\x91\x61\x94\x45\xea\x2f\xcd\x4d\xd9\xe8\xe1\xd8\x50\xf0\x04\x7a\x3c\x9b\x52\xa9\x18\x32\x8a\xfa\xe1\x44\xab\x78\x0d\x6f\x20\xa3\x45\x20\x4a\x0b\xf3\x29\x32\x66\xdb\x5b\x51\xed\x01\xf6\x6f\x81\xa6\xbf\x1d\xb4\xff\x78\x74\x7a\xf5\xfe\xdb\xb3\x8b\x93\xbf\x77\x2e\xed\xad\x35\xa8\x76\x13\xd8\xb9\x38\x79\x75\x74\xfd\xe6\xf8\xfb\xf7\x47\xdf\x5e\x9d\x5c\xbc\xbf\x3a\x7d\x7d\x72\x76\x7d\xd5\xb6\x07\xb6\xdc\x98\x25\xce\xb0\xb1\x5b\xcd\x27\x6a\x2f\x41\x1f\xab\xe8\x72\xc9\x25\x11\xab\x0d\xc9\xbe\xc5\xa0\xb6\xd0\xbd\xf7\x49\xf3\x3b\x87\x4c\xe0\x0f\xec\x8b\x2a\x9d\xce\x85\x93\x09\xfb\x90\xea\x6b\xa1\xbc\x0d\xac\x8d\xe7\x99\x69\x75\x3b\x68\xcb\x81\xe1\x5d\xde\xde\x79\x51\xdd\xac\xd8\x28\xea\x69\xbf\xe4\x88\x63\x0e\x89\x56\xce\x68\x7f\x97\x55\x7c\xf5\x52\x7c\x61\x11\x83\xd6\x80\x56\x2a\x10\xe9\x4c\xa8\xa4\xf1\xce\x25\x8b\xae\x0d\x43\x70\xd0\x54\x6f\xdf\x59\x34\x42\x93\xd4\x89\x0b\xfa\xea\xcb\x0d\x69\xcd\xd1\x6b\x52\x4f\x45\xbd\x49\xda\x65\x56\xf3\xbd\x91\x81\x95\xd6\x2f\xac\xa2\x99\x93\x29\x72\x89\x4e\xb8\x67\x34\xe6\x62\x44\x5e\x9b\x2d\x88\x26\x69\xcc\xe6\x1b\xb8\xdc\xb4\x39\x7b\xf0\x01\x56\x75\x51\x6d\xf1\xe0\x2b\xc7\x1e\x94\x14\x04\xb9\x28\xbc\x8f\x90\x8c\x9e\x34\x56\xc1\x79\xa7\xbd\x08\x0a\x6c\xac\x94\x8b\xc9\x96\x4e\x83\x41\xae\xdc\x4a\xb0\xbe\x37\x2c\x24\x8f\xf8\xa3\xc5\x59\x3f\xb5\x93\xb3\xe5\x17\x5c\xae\xae\xda\xbe\xae\xf1\x1e\xb1\x95\xa3\x9e\xc5\xdc\x23\x16\xee\xdd\xb9\x8f\x74\x91\x78\x75\x01\x68\xa3\xd4\x70\x37\x9c\x03\x8f\xdf\x70\xe9\x51\xa8\x0f\xae\x8b\x72\x79\x05\xd6\xe9\x48\x7e\xb1\x8a\xa6\x27\x52\xbe\xba\x1c\x71\xdb\x1b\x6e\xc2\x9a\xe1\x53\xb8\x56\x9b\x32\x09\xf0\x6b\xd3\x9f\xb0\xb6\xde\xd5\x7e\x7d\xf1\x6a\xb7\x75\x20\x74\x10\xc9\xbe\xc4\x8d\xdd\xd0\x25\x5f\xfd\x17\x52\x7f\x63\x55\x0f\xcf\x39\x80\x11\x8e\x09\x69\x15\xdf\x8f\x0e\xc5\x20\xd2\xb7\xf3\x39\x2f\x6d\x71\x2f\x92\xdb\x6b\xa6\xd9\x7e\x90\x59\xf6\xdd\xba\xf2\x90\x87\xa4\x88\xb9\x56\x14\x21\x1a\x41\xd1\x90\xc3\xa0\x35\xa9\x89\x3f\x87\x97\xa7\xdf\x5d\x9d\x5c\x70\x74\xec\xf2\xf4\xbb\x1f\x4e\x5f\xbd\x0a\x6f\x95\x93\x13\x25\xb2\x46\xca\x17\xe1\xcf\xab\x6e\x0e\x87\x70\x1d\xde\xfc\x45\xb2\x96\x16\x10\xd2\xba\x06\xa1\x71\x88\xcc\x05\xb7\x83\x2f\x61\xcf\x16\x60\xa7\xa5\x83\x54\xcf\x15\xc8\x3c\xc7\x94\x18\x5c\xb6\xf0\x85\xf0\x5e\x05\xa0\x41\x03\x44\x5b\x24\x33\x85\xa2\xf6\x4b\x67\x84\xc3\xc9\x43\xb4\xde\x97\xa5\xf1\xc9\x83\x55\x81\xbc\xc8\xb2\x76\x6a\x67\x27\xa7\x68\xc2\x02\x82\xf9\x16\x9f\x32\x01\x0a\xe7\xf1\xf5\x63\xfe\x8c\x76\xc0\x8a\x32\xc9\xb6\xf5\x61\xeb\x74\x61\x41\x67\x69\xec\xec\xcd\x06\xaf\x87\xa0\x6d\x0e\x6b\x37\xa7\x26\x6e\xf7\xea\x94\x60\xc8\x9d\xcd\xd0\x1c\x8b\x42\x24\xd2\xad\xcc\x2c\xea\x24\xad\xf5\x21\xef\x28\x92\x59\xa4\x2d\x7f\x89\xe1\xf3\x5a\x56\x37\x32\xfc\xaa\x22\xba\x88\x5c\x48\x74\xa9\x1a\xf7\x67\xf3\x6d\x54\x1e\xbe\x06\xf9\xb6\x72\x4b\xe3\xae\x04\x1c\xa5\xe0\x5f\x22\x00\xde\x61\x44\x5c\x53\xe6\x08\xa9\xdf\x53\x57\xef\x04\xe9\x80\x09\x5a\xdb\xd5\x8b\x2b\xb3\xb6\xf1\x4a\xb6\x2e\xc1\x0d\x0f\x7b\x55\xe1\xd0\xc2\x57\x62\x7d\x4e\x5c\xc6\x88\xec\xbd\x51\xe9\x01\x6c\xa0\x32\x96\xea\x90\x3a\x1b\x38\x66\x7c\x21\xa8\x48\x8c\x1c\xcb\xc4\xfb\xd0\xe9\x04\xb4\x32\xf8\x4a\x43\x96\xec\xb0\xa3\x30\xe7\xe2\x16\xc1\x96\xc4\x43\x57\x63\x7c\x10\x72\x2d\xfb\xf6\xa9\x53\x9c\x16\xc1\xab\xb7\x99\x75\xa9\x6e\x3e\xc7\xbd\x37\xad\x97\xf9\x96\x46\x3e\x48\xd6\x9c\x46\xdb\x36\x05\xbe\x68\x50\x98\x5a\xd6\xb4\x64\x49\x87\x65\x54\xf7\x5c\x0c\xe3\x1d\xaa\xcd\x06\x53\xae\x10\x0e\xe1\x9d\xe6\xa5\x1e\x83\x98\xf2\x12\xeb\x4d\x9a\x93\x7f\x3e\xd1\x54\x5a\x7c\x68\x5d\x3f\x75\x6d\xd4\xf4\x9b\x52\x55\xa2\xb5\x1d\xf2\xee\xe8\xa6\x81\xd3\x3d\x70\xd6\xf8\xf6\xc7\x98\x89\xdf\xab\x44\x7a\xd5\x89\xba\x90\x41\xb9\x4f\x34\xd8\x0b\xc3\xa9\x1a\xeb\x87\x88\x12\xbe\x6e\x1d\x2b\x76\xde\xcc\x85\x7f\x2c\xfe\x9d\x09\xeb\x2e\x13\x41\x52\xf9\xa8\x1b\x91\x7e\xc0\x55\x81\xe1\x52\x7b\x3a\x8c\xd6\x89\xbc\x88\x8a\x27\x7b\xb7\xc2\x2b\x3b\xe3\x4b\x59\x48\x37\xf7\x7c\x7d\x2e\x7c\xc8\x98\xef\x5b\x20\x01\x68\x3d\x44\xe0\xdf\xa0\x8b\xd5\x8b\x39\x5b\x09\xb0\x9d\x63\xbc\x72\x33\x60\x29\x77\xca\x3a\xaf\xdc\x1c\x33\x40\xbf\xd5\xba\x7d\xc2\x30\x2f\x23\xbc\x9a\x94\x6b\xb8\x1a\x6b\x6d\xac\x34\x1c\xe8\x21\x5c\x4a\x62\xc7\x64\x67\xd6\x42\xdd\x84\xa9\x46\xc2\x62\x1a\x9c\x12\x64\xfc\x56\x57\x8b\xe2\x4f\xa5\x7f\x7b\x56\xd4\x22\xee\x83\xaa\xf6\x01\xae\x5e\x82\xbd\x74\xf5\xed\x8e\xe4\x97\x6f\xb3\xca\xbe\x74\xb1\xc7\xc7\x27\xff\x17\x00\x00\xff\xff\xd7\xe1\xca\xc4\xea\x7c\x00\x00")

func appdefinitionJsonBytes() ([]byte, error) {
	return bindataRead(
		_appdefinitionJson,
		"AppDefinition.json",
	)
}

func appdefinitionJson() (*asset, error) {
	bytes, err := appdefinitionJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "AppDefinition.json", size: 31978, mode: os.FileMode(420), modTime: time.Unix(1487783921, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _groupJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x54\x4d\x6f\x1c\x37\x0c\xbd\xfb\x57\x10\x63\x1f\x76\x13\xcf\x28\x49\x4f\xd9\x4b\x61\xa0\x1f\x97\x1e\x0a\xd4\x3d\x65\x5d\x98\x2b\x71\x76\x14\xcc\x48\x0a\xc9\xb1\xbb\x49\xf6\xbf\x17\xd2\xcc\x7e\x18\x75\x10\x23\x97\x5d\x0d\xf5\x44\x3e\xf2\x91\xfc\x72\x01\x50\x5d\x89\xed\x68\xc0\x6a\x05\x55\xa7\x9a\x56\xc6\x7c\x94\x18\xea\xc9\xda\x44\xde\x9a\xe9\x78\x59\x5d\x67\xb8\x77\x07\xa4\xac\x8c\x61\x7c\x6c\xb6\x5e\xbb\x71\x33\x0a\xb1\x8d\x41\x29\x68\x63\xe3\x60\x06\x92\x28\xa9\x23\x26\x33\x20\xa3\x76\x31\x98\x01\x45\x89\x8d\x8b\x56\xa6\x1f\x26\xd1\x1a\x93\x37\x69\xdc\xf4\xde\x9a\x7c\x7c\x78\x37\x07\x34\xbf\x73\x1c\x53\x93\xd9\x4c\xa1\xd1\x39\xaf\x3e\x06\xec\xff\xe4\x98\x88\xd5\x93\x54\x2b\x68\xb1\x17\x2a\x00\x47\xad\x0f\x05\x92\xed\x39\x3b\x80\x2a\xa1\x76\xb7\xbb\x44\x47\x0b\x40\xa5\xd3\x77\x25\xca\x3e\x6c\x8b\xf7\x03\x56\x89\x43\xbe\xfa\x67\xb1\x5e\x9b\x9f\x17\x8b\xf5\xba\xf9\xf2\x6e\xbf\xfc\xba\xf8\x80\xf5\xe7\x37\xf5\xfb\xbb\xf9\x7f\xbd\xae\x9b\xbb\x57\x07\xe3\xeb\x33\xc0\xab\xe5\x72\x71\xf5\x75\xbd\x36\xcb\xe5\xeb\xab\x93\xef\xc1\x87\x3f\x28\x6c\xb5\xab\x56\xf0\xb6\x18\xf7\x17\x00\xfb\x42\xfc\xc0\x27\x6e\x3e\x92\xd5\x29\xdb\x74\x9e\xe3\x9c\x4b\x29\xfe\x77\xb3\xb8\x62\x6a\xb3\xfd\xd2\x9c\x15\xc4\x1c\xeb\x70\xc4\x39\x12\xcb\x3e\xe5\xeb\x0c\xff\x3b\xf8\x4f\x23\x81\x77\x14\xd4\xb7\x9e\x18\xda\xc8\xa0\x1d\x01\xa6\x04\x36\x06\xf1\xa2\x3e\x6c\x21\xb6\x80\x20\xc4\x9e\x24\x9f\x03\x0e\x24\x20\x94\xb2\xcc\xe4\x60\xb3\x03\xe9\x51\x3a\x92\x06\x7e\x45\xdb\x15\x00\x0c\xa3\x28\x6c\x08\x50\xa1\x27\x14\x85\xb7\x60\x3b\x64\xb4\x4a\x0c\x18\x1c\x0c\xb8\x83\x18\xfa\x5d\x8e\xa4\xe8\x03\x38\xbf\xf5\x2a\xb0\xb8\x7f\x53\xbf\xbf\x5f\x5e\x83\x2b\x3e\x61\x71\x5f\x97\xaf\x58\xee\x9a\x7c\xce\xcf\xfb\xf8\x48\x6c\x51\x08\x7a\xca\x22\xe6\x4b\xac\x3f\xdf\x2f\x1b\xb8\xed\x68\xe6\x80\x3b\x08\x31\xd3\xd8\xfa\x00\x91\x81\x82\x83\x47\xaf\x1d\x60\xf1\xde\x54\x93\x2e\x53\x85\x2a\x4c\x49\x9e\xab\x37\x32\xe3\xee\x9b\x65\xcc\xd1\x7a\x2f\x9a\x4b\x73\x93\xd2\x2f\x27\x09\xc0\x07\xd0\xce\x0b\x6c\x4b\x63\xc3\x5f\x44\x4f\x11\xa5\xd7\x8f\x55\x9f\x27\xf0\x14\xc8\x2b\x0d\xe7\x84\xce\xa4\x6e\xcc\xff\x1d\x55\x33\x6c\xff\x24\xa9\x12\xfb\x07\xd2\x2a\xc3\x28\x60\x31\xc0\x66\xf4\xbd\x03\x04\x65\xa2\x59\xe2\xe2\xb5\x5c\x1e\xd4\x93\x71\x53\x4f\xb1\x26\x01\x4e\xdf\x80\x4c\x50\x3a\x93\x1c\xe4\x05\xf1\xb2\x14\x2f\x9f\xcf\xc7\x51\xa2\xe0\x28\xd8\xf3\x41\x79\x71\x56\x37\x47\xa9\x84\xf8\xc1\x5b\x12\x18\x53\x0c\xf0\xd8\x79\xdb\x4d\x62\x61\x4a\xbd\xb7\x98\x5f\xc0\x14\x4c\x1a\xb8\xc9\xed\xe3\x88\xc1\x0b\x38\x62\xff\x40\x0e\x5a\x8e\x43\x11\xee\x9c\x52\x51\x33\x11\xb7\x91\x87\x3c\x3b\xa2\xc8\x6a\x44\x63\x2a\x5d\x3b\xa6\x2d\xa3\xa3\x4c\x60\x1e\xb4\x43\xac\x06\x7e\xcb\x0d\xfa\x2f\x0e\xa9\xa7\xdc\xe2\x4f\x88\x18\x04\xa6\xbe\x0c\x60\x98\x9a\xe5\x40\xdf\x6c\x66\xf2\x5e\x85\xfa\xf6\x0c\x66\x6c\x03\xb7\x71\x62\x00\xd8\xf7\xf0\xd3\xb9\x4b\xb9\x86\xd6\xb3\x28\x18\x9b\x93\x2a\x28\x72\xa0\x1d\x86\xec\x73\xfa\x7f\x61\x37\x7e\x63\xf1\x3c\xab\xdf\x03\xb1\x4c\x62\x7c\x77\xaf\x3d\x33\x68\xf3\xeb\xa9\x7e\x45\x8b\xe3\x08\x9c\xde\xe5\xda\xa3\xe6\x27\x0e\x95\x6a\xf5\xc3\xcc\xe5\xb4\x80\x99\x3e\x8d\x9e\x29\xaf\xd7\x0f\xc7\x55\x7b\x01\x70\x77\xb1\xbf\xf8\x2f\x00\x00\xff\xff\x42\x45\xef\xb4\x26\x07\x00\x00")

func groupJsonBytes() ([]byte, error) {
	return bindataRead(
		_groupJson,
		"Group.json",
	)
}

func groupJson() (*asset, error) {
	bytes, err := groupJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "Group.json", size: 1830, mode: os.FileMode(420), modTime: time.Unix(1487783921, 0)}
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
	"AppDefinition.json": appdefinitionJson,
	"Group.json": groupJson,
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
	"AppDefinition.json": &bintree{appdefinitionJson, map[string]*bintree{}},
	"Group.json": &bintree{groupJson, map[string]*bintree{}},
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

