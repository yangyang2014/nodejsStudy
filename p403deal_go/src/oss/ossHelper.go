package oss

import (
	"path/filepath"
	"strings"
)

func UploadDir(dirname string, remoteDir string, failedList []interface{}) {
	return
}

func getOption() {
	return
}

func ToRelativePath(path1 string) string {
	var index = strings.Index(path1, ".aliyuncs.com")
	if index >= 0 {
		return filepath.Base(path1)
	}
	return path1
}

func DeleteDir(dir string) {
	// log.info(`will delete ${dir} on oss`);
	// let result = await list( dir)
	// if(!result) return ;
	// let files=result.map(elem=>elem.name);
	// return await deleteMulti(files);
}
