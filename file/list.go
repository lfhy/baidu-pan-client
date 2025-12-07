package file

import (
	"github.com/lfhy/baidu-pan-client/http"
	"github.com/lfhy/baidu-pan-client/types"
)

type ListReq struct {
	Dir       string          `query:"dir" default:"/"`      //	需要list的目录，以/开头的绝对路径, 默认为/
	Order     types.ListOrder `query:"name" default:"name"`  // 	排序字段：默认为name； time按修改时间排序；name表示按文件名称排序 size表示按文件大小排序。
	Desc      types.BoolInt   `query:"desc"`                 // 默认为升序 设置为true实现降序 （注：排序的对象是当前目录下所有文件，不是当前分页下的文件）
	Start     int             `query:"start" default:"0"`    //	起始位置，从0开始
	Limit     int             `query:"limit" default:"1000"` // 查询数目，默认为1000，建议最大不超过1000
	Web       types.BoolInt   `query:"web"`                  // 返回dir_empty属性和缩略图数据；不传该参数，则不返回缩略图地址
	Folder    types.BoolInt   `query:"folder"`               //	是否只返回文件夹且属性只返回path字段
	Showempty types.BoolInt   `query:"showempty"`            // 是否返回dir_empty属性，0 不返回，1 返回
}

type Thumbs struct {
	Url1 string `json:"url1"`
	Url2 string `json:"url2"`
	Url3 string `json:"url3"`
}

type ListRes struct {
	FsId        uint64             `json:"fs_id"`           // 文件在云端的唯一标识ID
	Path        string             `json:"path"`            // 文件的绝对路径
	Name        string             `json:"server_filename"` //	文件名称
	Size        types.SizeB        `json:"size"`            //	文件大小，单位B
	ServerMtime types.Time         `json:"server_mtime"`    //	文件在服务器修改时间
	ServerCtime types.Time         `json:"server_ctime"`    //	文件在服务器创建时间
	LocalMtime  types.Time         `json:"local_mtime"`     //	文件在客户端修改时间
	LocalCtime  types.Time         `json:"local_ctime"`     //	文件在客户端创建时间
	IsDir       types.BoolInt      `json:"isdir"`           //	是否为目录，0 文件、1 目录
	Category    types.FileCategory `json:"category"`        //	文件类型，1 视频、2 音频、3 图片、4 文档、5 应用、6 其他、7 种子
	ServerMd5   string             `json:"md5"`             //	云端哈希（非文件真实MD5），只有是文件类型时，该字段才存在
	DirEmpty    types.BoolInt      `json:"dir_empty"`       //	该目录是否存在子目录，只有请求参数web=1且该条目为目录时，该字段才存在， 0为存在， 1为不存在
	Thumbs      *Thumbs            `json:"thumbs"`          //	只有请求参数web=1且该条目分类为图片时，该字段才存在，包含三个尺寸的缩略图URL；不传web参数，则不返回缩略图地址
}

func List(req *ListReq) ([]*ListRes, error) {
	api := &http.API[*ListReq, struct {
		List []*ListRes `json:"list"`
	}]{
		AccessToken: types.AccessToken,
		BaseURL:     types.PanBaseURL,
		HTTPMethod:  http.GET,
		Method:      "list",
		Request:     req,
		Route:       ListRoute,
	}
	res, err := api.Do()
	if err != nil {
		return nil, err
	}
	return res.List, nil
}
