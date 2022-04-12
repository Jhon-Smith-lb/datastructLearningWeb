package daonews

import (
	"dataStructLearningWeb/dao"
	"dataStructLearningWeb/lib"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

func AddNews(req *dao.News, o orm.Ormer) (int64, error) {
	newsId, err := o.Insert(req)
	if err != nil {
		logs.Error("[AddNews] o.Insert, err: %v, req: %v", err, lib.PointerToString(req))
		return 0, err
	}
	return newsId, nil
}

func QueryNewsList(req *dao.QueryNewsOption) ([]*dao.News, int64, error) {
	o := orm.NewOrm()
	table := &dao.News{}
	resp := make([]*dao.News, 0)
	qs := o.QueryTable(&table)

	if req.Title != "" {
		qs = qs.Filter("title__contains", req.Title)
	}

	qs.Offset(req.Offset).Limit(req.Limit)
	_, err := qs.All(&resp)
	if err != nil {
		logs.Error("[QueryNewsList] qs.All, err: %v, req: %v", err, lib.PointerToString(req))
		return nil, 0, err
	}
	return resp, 0, nil
}

func UpdateNews(req *dao.News, o orm.Ormer) error {
	mp := make(map[string]interface{}, 0)
	mp["title"] = req.Title
	mp["user_id"] = req.UserId
	mp["type"] = req.Type
	mp["text"] = req.Text
	mp["updated_at"] = req.UpdatedAt
	mp["is_del"] = req.IsDel
	table := &dao.User{}
	_, err := o.QueryTable(table).Filter("id", req.Id).Update(mp)
	if err != nil {
		logs.Error("[UpdateNews] o.Update, err: %v, req: %v", err, lib.PointerToString(req))
		return err
	}
	return nil
}

func GetNewsById(newsId int64) (*dao.News, error) {
	o := orm.NewOrm()
	news := &dao.News{Id: newsId}
	if err := o.Read(news); err != nil {
		logs.Error("[GetNewsById] o.Read, err: %v, newsId: %v", err, newsId)
		return nil, err
	}
	return news, nil
}
