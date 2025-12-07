package test

import (
	"reflect"
	"testing"

	"github.com/lfhy/flag"

	glog "log"

	"github.com/lfhy/baidu-pan-client/log"
	"github.com/lfhy/baidu-pan-client/types"
)

type TestLogger struct {
}

func (t *TestLogger) Printf(f string, v ...any) {
	glog.Printf(f, v...)
}

func (t *TestLogger) Println(v ...any) {
	glog.Println(v...)
}

func ReadConfig() {
	f := flag.NewFlagSet("test", flag.ContinueOnError)
	f.String("c", "../config.toml", "配置文件")
	f.StringConfigVar(&types.ClientId,
		"client_id",
		"auth", "client_id",
		"",
		"client_id",
	)
	f.StringConfigVar(&types.ClientSecret,
		"client_secret",
		"auth", "client_secret",
		"",
		"client_secret",
	)
	f.StringConfigVar(&types.RedirectUri,
		"redirect_uri",
		"auth", "redirect_uri",
		"",
		"redirect_uri",
	)
	f.StringConfigVar(&types.AccessToken,
		"access_token",
		"auth", "access_token",
		"",
		"access_token",
	)
	f.StringConfigVar(&types.RefreshToken,
		"refresh_token",
		"auth", "refresh_token",
		"",
		"refresh_token",
	)
	f.Parse(nil)
}

func TestSetEnv(t *testing.T) {
	ReadConfig()
	log.SetLogger(&TestLogger{})
}

// dereferencePointer 解引用指针获取实际值
func dereferencePointer(v reflect.Value) interface{} {
	// 如果是nil指针，直接返回nil
	if v.IsNil() {
		return nil
	}
	// 循环解引用直到得到非指针值
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v.Interface()
}

// formatValue 格式化值用于打印
func formatValue(v interface{}) interface{} {
	if v == nil {
		return nil
	}

	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Ptr:
		return dereferencePointer(rv)
	case reflect.Slice:
		// 处理切片类型
		if rv.Len() == 0 {
			return v
		}
		// 检查切片元素是否为指针
		if rv.Type().Elem().Kind() == reflect.Ptr {
			// 创建新的切片存储解引用后的值
			result := make([]interface{}, rv.Len())
			for i := 0; i < rv.Len(); i++ {
				result[i] = dereferencePointer(rv.Index(i))
			}
			return result
		}
		return v
	default:
		return v
	}
}

func PrintRes(res any, err error) {
	if err != nil {
		log.Println("发生错误:", err)
	} else {
		formatted := formatValue(res)
		log.Printf("响应结果:%+v", formatted)
	}
}
