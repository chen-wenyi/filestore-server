package meta

import (
	"sort"
)

type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

func UpdateFileMeta(fmeta FileMeta) {
	fileMetas[fmeta.FileSha1] = fmeta
}

func GetFileMeta(fileSha1 string) (fileMeta FileMeta, ok bool) {
	fileMeta, ok = fileMetas[fileSha1]
	return fileMeta, ok
}

func GetLatestFileMetas(count int) []FileMeta {
	fileMetaSlice := make([]FileMeta, len(fileMetas))
	i := 0
	for _, v := range fileMetas {
		fileMetaSlice[i] = v
		i++
	}
	sort.Sort(byUploadTime(fileMetaSlice))

	if len(fileMetaSlice) < count {
		return fileMetaSlice[0:]
	}
	return fileMetaSlice[:count]
}
