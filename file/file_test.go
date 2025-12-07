package file_test

import (
	"testing"

	"github.com/lfhy/baidu-pan-client/file"
	"github.com/lfhy/baidu-pan-client/test"
)

func TestFileList(t *testing.T) {
	test.TestSetEnv(t)
	res, err := file.List(&file.ListReq{
		Dir: "/",
	})
	test.PrintRes(res, err)
}
