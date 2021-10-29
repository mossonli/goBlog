package servers

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"goBlog/utils"
	"goBlog/utils/errmsg"
	"mime/multipart"
)

var AccessKey = utils.AccessKey
var SecretKey = utils.SecretKey
var Bucket = utils.Bucket
var ImgUrl = utils.QiniuServer

func UploadFile(file multipart.File, fileSize int64) (string, int) {
	mac := qbox.NewMac(AccessKey, SecretKey)

	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{
		Zone: &storage.ZoneHuabei,
		UseCdnDomains: false,
		UseHTTPS: false,
	}
	putExtra := storage.PutExtra{}

	formUpLoader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err := formUpLoader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil{
		return "", errmsg.ERROR
	}
	url := ImgUrl + ret.Key
	return url, errmsg.SUCCESS

}