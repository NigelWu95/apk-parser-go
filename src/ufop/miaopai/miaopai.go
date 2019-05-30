package miaopai

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/qiniu/api.v6/auth/digest"
	"github.com/qiniu/log"
	"io"
	"net/http"
	"os"
	"regexp"
	"ufop"
	"ufop/apk"
)

const (
	MKZIP_MAX_FILE_LENGTH int64 = 100 * 1024 * 1024 //100MB
	MKZIP_MAX_FILE_COUNT  int   = 100               //100
	MKZIP_MAX_FILE_LIMIT  int   = 1000              //1000
)

type APKParser struct {
	mac *digest.Mac
}

type APKParserConfig struct {
	//ak & sk
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
}

type APKInfo struct {
	AppName     string `json:"app_name"`
	PackageName string `json:"package_name"`
	Version     string `json:"version"`
	VersionCode int    `json:"version_code"`
	Size        int64  `json:"size"`
	Icon        string `json:"icon"`
	MD5         string `json:"md5"`
}

func (this *APKParser) Name() string {
	return "aparser"
}

func (this *APKParser) InitConfig(jobConf string) (err error) {
	return
}

func (this *APKParser) parse(cmd string) (err error) {
	pattern := "^aparser"
	matched, _ := regexp.MatchString(pattern, cmd)
	if !matched {
		err = errors.New("invalid parser command format")
		return
	}
	return
}

func (this *APKParser) Do(req ufop.UfopRequest, ufopBody io.ReadCloser) (result interface{}, resultType int, contentType string, err error) {
	reqId := req.ReqId
	//parse command
	pErr := this.parse(req.Cmd)
	if pErr != nil {
		err = pErr
		return
	}

	resUrl := req.Url
	//download apk
	resResp, respErr := http.Get(resUrl)
	if respErr != nil || resResp.StatusCode != 200 {
		if respErr != nil {
			err = fmt.Errorf("retrieve resource apk failed, %s", respErr.Error())
		} else {
			err = fmt.Errorf("retrieve resource apk failed, %s", resResp.Status)
			if resResp.Body != nil {
				resResp.Body.Close()
			}
		}
		return
	}
	defer resResp.Body.Close()

	//save apk
	apkPath := fmt.Sprintf("./%s.apk", reqId)
	f, fErr := os.Create(apkPath)
	if fErr != nil {
		err = fmt.Errorf("create file failed, %s", fErr.Error())
		return
	}
	_, cpErr := io.Copy(f, resResp.Body)
	if cpErr != nil {
		err = fmt.Errorf("save local apk file failed, %s", cpErr.Error())
		return
	}
	defer f.Close()

	//md5
	fmd5, fmd5Err := os.Open(apkPath)
	if fmd5Err != nil {
		err = fmt.Errorf("open apk failed, %s", fmd5Err.Error())
		return
	}
	defer fmd5.Close()
	md5h := md5.New()
	_, cpErr = io.Copy(md5h, fmd5)
	if cpErr != nil {
		err = fmt.Errorf("calcul apk's md5 failed, %s", cpErr.Error())
		return
	}
	md5Str := md5h.Sum(nil)
	//size
	fInfo, fInfoErr := fmd5.Stat()
	if fInfoErr != nil {
		err = fmt.Errorf("get apk's size failed, %s", fInfoErr.Error())
		return
	}

	//获取 apk 基本信息
	pkg, pkgErr := apk.OpenFile(apkPath)
	if pkgErr != nil {
		err = fmt.Errorf("parse apk failed, %v", pkgErr)
		return
	}
	defer pkg.Close()

	label, labelErr := pkg.Label(nil)
	if labelErr != nil {
		err = fmt.Errorf("parse apk's label failed, %v", labelErr)
		return
	}
	icon := pkg.Icon(nil)
	packageName := pkg.PackageName()
	manifest := pkg.Manifest()
	version := manifest.VersionName
	versionCode := manifest.VersionCode

	result = APKInfo{
		AppName:     label,
		PackageName: packageName,
		Version:     version,
		VersionCode: versionCode,
		Size:        fInfo.Size(),
		Icon:        icon,
		MD5:         fmt.Sprintf("%x", md5Str),
	}

	defer os.Remove(apkPath)

	// result =
	resultType = ufop.RESULT_TYPE_JSON
	contentType = "application/json"
	log.Infof("[%s] apk parser success!", reqId)
	return
}
