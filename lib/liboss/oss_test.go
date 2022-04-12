package liboss

import (
	"encoding/json"
	"testing"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/smartystreets/goconvey/convey"
)

func TestGetFileList(t *testing.T) {
	convey.Convey("TestGetFileList", t, func ()  {
		convey.Convey("success", func ()  {
			client, err := oss.New(ENDPOINT, ACCESSKEY_ID, ACCESSKEY_SECRET)
			if err != nil {
				t.Errorf("err: %v", err)
				return
			}
			bucket, err := client.Bucket(BUCKET_NAME)
			if err != nil {
				t.Errorf("err: %v", err)
				return 
			}
			
			lsRes, err := bucket.ListObjects()
			if err != nil {
				t.Errorf("err: %v", err)
				return 
			}
			
			for _, object := range lsRes.Objects {
				signedURL, err := bucket.SignURL(object.Key, oss.HTTPGet, 3600)
				if err != nil {
					t.Errorf("err: %v, objectKey: %v", err, object.Key)
					continue
				}
				objectJson, _ := json.Marshal(object)
				t.Logf("objectJson: %v, signedURL: %v\n", string(objectJson), signedURL)
			}
		})
	})
}

