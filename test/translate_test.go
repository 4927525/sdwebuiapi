package test

import (
	"os"
	config2 "sdwebuiapi/config"
	"testing"

	alimt20181012 "github.com/alibabacloud-go/alimt-20181012/v2/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

func TestTranslate(t *testing.T) {
	config2.InitConfig()
	config2.InitDB()
	config2.InitRDB()

	err := _main(tea.StringSlice(os.Args[1:]))
	if err != nil {
		panic(err)
	}
}

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *alimt20181012.Client, _err error) {
	accessKey := config2.Config.GetString("aliyun.accessKey")
	accessSecret := config2.Config.GetString("aliyun.accessKeySecret")
	config := &openapi.Config{
		// 您的 AccessKey ID
		AccessKeyId: &accessKey,
		// 您的 AccessKey Secret
		AccessKeySecret: &accessSecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("mt.aliyuncs.com")
	_result = &alimt20181012.Client{}
	_result, _err = alimt20181012.NewClient(config)
	return _result, _err
}

func _main(args []*string) (_err error) {
	client, _err := CreateClient(tea.String("accessKeyId"), tea.String("accessKeySecret"))
	if _err != nil {
		return _err
	}

	translateGeneralRequest := &alimt20181012.TranslateGeneralRequest{
		FormatType:     tea.String("text"),
		SourceLanguage: tea.String("zh"),
		TargetLanguage: tea.String("en"),
		SourceText:     tea.String("hello,世界"),
		Scene:          tea.String("general"),
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		_result, _err := client.TranslateGeneralWithOptions(translateGeneralRequest, runtime)
		if _err != nil {
			return _err
		}

		println(*_result.Body.Data.Translated)

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 如有需要，请打印 error
		_, _err = util.AssertAsString(error.Message)
		if _err != nil {
			return _err
		}
	}
	return _err
}
