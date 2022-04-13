package biznews

import (
	"dataStructLearningWeb/dao"
	"dataStructLearningWeb/dao/daonews"
	"dataStructLearningWeb/dao/daouser"
	"dataStructLearningWeb/dm/dmnews"
	"dataStructLearningWeb/dm/dmuser"
	"dataStructLearningWeb/lib"
	"errors"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

func AddNews(req *dmnews.AddNewsReq) (int64, error) {
	if err := req.CheckParam(); err != nil {
		return 0, err
	}
	o := orm.NewOrm()

	dbReq := dao.NewNews()
	dbReq.Title = req.Title
	dbReq.UserId = req.UserId
	dbReq.Type = int16(req.Type)
	dbReq.Text = req.Text
	dbReq.CreatedAt = time.Now()
	dbReq.UpdatedAt = time.Now()

	userId, err := daonews.AddNews(dbReq, o)
	if err != nil {
		logs.Error("[AddNews] daonews.AddNews, err: %v, dbReq: %v", err, lib.PointerToString(dbReq))
		return 0, err
	}
	return userId, nil
}

func UpdateNews(req *dmnews.UpdateNewsReq) error {
	if err := req.CheckParam(); err != nil {
		logs.Info("[UpdateNews] req.CheckParam(), err: %v, req: %v", err, lib.PointerToString(req))
		return err
	}

	// 更新之前查询一下用户是否存在
	news, err := daonews.GetNewsById(req.Id)
	if err != nil {
		// 查询不到
		if err == orm.ErrNoRows {
			logs.Error("[UpdateNews] daonews.GetNewsById, err: %v, req: %v", err, lib.PointerToString(req))
			return errors.New("找不到文章")
		}
		return err
	}

	// 进行可以进行修改的权限校验
	// 是不是作者
	if req.UserId != news.UserId {
		user, err := daouser.GetUserById(req.UserId)
		if err != nil {
			if err == orm.ErrNoRows {
				logs.Error("[UpdateNews] daouser.GetUserById, err: %v, req: %v", err, lib.PointerToString(req))
				return errors.New("找不到用户")
			}
			return err
		}
		// 是不是管理员
		if user.IsAdmin != dmuser.IS_ADMIN {
			return errors.New("没有修改该文章的权限")
		}
	}

	o := orm.NewOrm()

	dbReq := dao.NewNews()
	dbReq.Id = req.Id
	dbReq.Title = req.Title
	dbReq.UserId = req.UserId
	dbReq.Type = req.Type
	dbReq.Text = req.Text
	dbReq.UpdatedAt = time.Now()
	dbReq.IsDel = req.IsDel

	if err := daonews.UpdateNews(dbReq, o); err != nil {
		logs.Error("[UpdateNews] daonews.UpdateNews, err: %v, dbReq: %v", err, lib.PointerToString(dbReq))
		return err
	}

	return nil
}

func QueryNewsList(req *dmnews.QueryNewsOption) ([]*dmnews.News, int64, error) {
	if err := req.CheckParam(); err != nil {
		logs.Info("[QueryNewsList] req.CheckParam(), err: %v, req: %v", err, lib.PointerToString(req))
		return nil, 0, err
	}

	dbQueryNewsOption := dao.NewQueryNewsOption()
	dbQueryNewsOption.Title = req.Title
	dbQueryNewsOption.Offset = req.Offset
	dbQueryNewsOption.Limit = req.Limit

	dbNewsList, total, err := daonews.QueryNewsList(dbQueryNewsOption)
	if err != nil {
		logs.Error("[QueryUserList] daouser.QueryUserList, err: %v, dbQueryUserOption: %v", err, lib.PointerToString(dbQueryNewsOption))
		return nil, 0, err
	}

	dmNewsList := make([]*dmnews.News, 0, len(dbNewsList))
	for _, dbNews := range dbNewsList {
		dmNewsList = append(dmNewsList, &dmnews.News{
			Id:        dbNews.Id,
			Title:     dbNews.Title,
			UserId:    dbNews.UserId,
			Type:      dbNews.Type,
			Text:      dbNews.Text,
			CreatedAt: dbNews.CreatedAt,
			UpdatedAt: dbNews.UpdatedAt,
			IsDel:     dbNews.IsDel,
			Offset:    dbNews.Offset,
			Limit:     dbNews.Limit,
		})
	}

	return dmNewsList, total, nil
}
