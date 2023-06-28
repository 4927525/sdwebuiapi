package translate

import (
	config2 "sdwebuiapi/config"

	alimt20181012 "github.com/alibabacloud-go/alimt-20181012/v2/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

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

func Cn2En(word string) (*string, error) {
	client, err := CreateClient(tea.String("accessKeyId"), tea.String("accessKeySecret"))
	if err != nil {
		return nil, err
	}

	translateGeneralRequest := &alimt20181012.TranslateGeneralRequest{
		FormatType:     tea.String("text"),
		SourceLanguage: tea.String("zh"),
		TargetLanguage: tea.String("en"),
		SourceText:     tea.String(word),
		Scene:          tea.String("general"),
	}
	runtime := &util.RuntimeOptions{}
	translate, tryErr := func() (translate *string, _e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		trans, err := client.TranslateGeneralWithOptions(translateGeneralRequest, runtime)
		if err != nil {
			return nil, err
		}

		return trans.Body.Data.Translated, nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 如有需要，请打印 error
		_, err = util.AssertAsString(error.Message)
		if err != nil {
			return nil, err
		}
	}
	return translate, err
}
