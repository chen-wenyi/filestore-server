package meta

import "time"

const baseFormat = "2006-01-02 15:04:05"

type byUploadTime []FileMeta

func (fMetaSlice byUploadTime) Len() int {
	return len(fMetaSlice)
}

func (fMetaSlice byUploadTime) Less(i, j int) bool {
	iTime, _ := time.Parse(baseFormat, fMetaSlice[i].UploadAt)
	jTime, _ := time.Parse(baseFormat, fMetaSlice[j].UploadAt)
	return iTime.UnixNano() > jTime.UnixNano()
}

func (fMetaSlice byUploadTime) Swap(i, j int) {
	fMetaSlice[i], fMetaSlice[j] = fMetaSlice[j], fMetaSlice[i]
}
