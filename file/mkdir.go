package file

import "github.com/lfhy/baidu-pan-client/types"

func Mkdir(path string) (*CreateRes, error) {
	return Create(&CreateReq{
		Path:  path,
		IsDir: types.BoolIntTrue,
	})
}
