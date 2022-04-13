package bizresources

import (
	"dataStructLearningWeb/dm/dmresources"
	"dataStructLearningWeb/lib/liboss"
	"strings"

	"github.com/astaxie/beego/logs"
)

// 从阿里云OSS获取指定的资源列表
func QueryResourcesList(prefix string, limit int) ([]*dmresources.Resources, error) {
	if limit == 0 {
		limit = 100
	}
	objectPropertiesList, err := liboss.GetResourcesWithPrefix(prefix, limit)
	if err != nil {
		logs.Error("[QueryResourcesList] err: %v, prefix: %v, limit: %v", err, prefix, limit)
		return nil, err
	}
	resourcesList := make([]*dmresources.Resources, 0, len(objectPropertiesList))
	for _, objectProperties := range objectPropertiesList {
		arr := strings.Split(objectProperties.Key, "/")
		if len(arr) == 2 && arr[1] != "" {
			resources := dmresources.NewResources()
			signedURL, err := liboss.GetSignedURL(objectProperties.Key)
			if err != nil {
				logs.Error("[QueryResourcesList] err: %v, prefix: %v, limit: %v", err, prefix, limit)
				continue
			}
			resources.Name = arr[1]
			resources.Url = signedURL
			resourcesList = append(resourcesList, resources)
		} 
	}
	return resourcesList, nil
}