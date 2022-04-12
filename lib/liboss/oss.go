package liboss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/astaxie/beego/logs"
)

const (
	ACCESSKEY_ID = "LTAI5tJm78FjDDwT2WG9wean"
	ACCESSKEY_SECRET = "dSqQeCQSoYHbcrDkQLtBKzxpGl47dj"
	ENDPOINT = "oss-cn-beijing.aliyuncs.com"
	BUCKET_NAME = "datastruct-resources"
	EXPIRATION = 300
)

var client *oss.Client
var datastructResourcesBucket *oss.Bucket

func init() {
	var err error
	client, err = oss.New(ENDPOINT, ACCESSKEY_ID, ACCESSKEY_SECRET)
    if err != nil {
        logs.Error("[init] oss 初始化失败")
    }
	datastructResourcesBucket, err = client.Bucket(BUCKET_NAME)
    if err != nil {
        logs.Error("[init] client.Bucket, err: %v, bucket_name: %v", err, BUCKET_NAME)
    }
}

func GetResourcesWithPrefix(prefix string, limit int) ([]oss.ObjectProperties, error) {
    // 列举包含指定前缀的文件。列举指定个数文件。
    lsRes, err := datastructResourcesBucket.ListObjects(oss.Prefix(prefix), oss.MaxKeys(limit))
    if err != nil {
        logs.Error("[GetResourcesWithPrefix] bucket.ListObjects, err: %v, prefix: %v, limit: %v", err, prefix, limit)
		return nil, err
    }
	return lsRes.Objects, nil
}

func GetSignedURL(objectKey string) (string, error) {
	signedURL, err := datastructResourcesBucket.SignURL(objectKey, oss.HTTPGet, EXPIRATION)
	if err != nil {
		logs.Error("[GetSignURL] err: %v, objectKey: %v", err, objectKey)
		return "", err
	}
	return signedURL, nil
}