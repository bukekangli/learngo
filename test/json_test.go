package test

import (
	"encoding/json"
	"strconv"
	"testing"
)

type HistoryEventOptionVersion struct {
	Version   int64
	Index     int64
	Size      int64
	CreatorId int64
	CTime     int64
	Scene     int64
	TagName   string
}

func TestJsonList(t *testing.T) {
	versions := make([]*HistoryEventOptionVersion, 0)
	versions = append(versions, &HistoryEventOptionVersion{
		Version:   1,
		Index:     2,
		Size:      3,
		CreatorId: 4,
		CTime:     5,
		Scene:     6,
		TagName:   "tag_name",
	})
	versionsList := make([][]string, 0)
	for _, version := range versions {
		l := make([]string, 0)
		l = append(l, strconv.FormatInt(version.Version, 10))
		l = append(l, strconv.FormatInt(version.Index, 10))
		l = append(l, strconv.FormatInt(version.Size, 10))
		l = append(l, strconv.FormatInt(version.CreatorId, 10))
		l = append(l, strconv.FormatInt(version.CTime, 10))
		l = append(l, strconv.FormatInt(version.Scene, 10))
		l = append(l, version.TagName)
		versionsList = append(versionsList, l)
	}
	val, err := json.Marshal(versionsList)
	t.Logf("val: %s, err: %v", string(val), err)
}
