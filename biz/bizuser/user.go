package bizuser

import (
	"dataStructLearningWeb/dao"
	"dataStructLearningWeb/dao/daouser"
	"dataStructLearningWeb/dm/dmuser"
	"dataStructLearningWeb/lib"
	"errors"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

func AddUser(req *dmuser.AddUserReq) (int64, error) {
	if err := req.CheckParam(); err != nil {
		return 0, err
	}

	o := orm.NewOrm()

	dbReq := dao.NewUser()
	dbReq.Username = req.Username
	dbReq.Number = req.Number
	dbReq.Password = req.Password
	dbReq.IsAdmin = req.IsAdmin
	dbReq.CreateNews = req.CreateNews

	userId, err := daouser.AddUser(dbReq, o)
	if err != nil {
		logs.Error("[AddUser] daouser.AddUser, err: %v, dbReq: %v", err, lib.PointerToString(dbReq))
		return 0, err
	}
	return userId, nil
}

func UpdateUser(req *dmuser.UpdateUserReq) error {
	if err := req.CheckParam(); err != nil {
		logs.Info("[UpdateUser] req.CheckParam(), err: %v, req: %v", err, lib.PointerToString(req))
		return err
	}

	// 更新之前查询一下用户是否存在
	_, err := daouser.GetUserById(req.Id)
	if err != nil {
		// 查询不到
		if err == orm.ErrNoRows {
			logs.Error("[UpdateUser] daouser.GetUserById, err: %v, req: %v", err, lib.PointerToString(req))
			return errors.New("找不到用户")
		}
		return err
	}

	o := orm.NewOrm()

	dbReq := dao.NewUser()
	dbReq.Id = req.Id
	dbReq.Username = req.Username
	dbReq.Number = req.Number
	dbReq.Password = req.Password
	dbReq.Status = req.Status
	dbReq.IsAdmin = req.IsAdmin
	dbReq.IsDel = req.IsDel
	dbReq.CreateNews = req.CreateNews

	if err := daouser.UpdateUser(dbReq, o); err != nil {
		logs.Error("[UpdateUser] daouser.UpdateUser, err: %v, dbReq: %v", err, lib.PointerToString(dbReq))
		return err
	}

	return nil
}

func QueryUserList(req *dmuser.QueryUserOption) ([]*dmuser.User, int64, error) {
	if err := req.CheckParam(); err != nil {
		logs.Info("[QueryUserList] req.CheckParam(), err: %v, req: %v", err, lib.PointerToString(req))
		return nil, 0, err
	}

	dbQueryUserOption := dao.NewQueryUserOption()
	dbQueryUserOption.Username = req.Username
	dbQueryUserOption.Number = req.Number
	dbQueryUserOption.Offset = req.Offset
	dbQueryUserOption.Limit = req.Limit

	dbUserList, total, err := daouser.QueryUserList(dbQueryUserOption)
	if err != nil {
		logs.Error("[QueryUserList] daouser.QueryUserList, err: %v, dbQueryUserOption: %v", err, lib.PointerToString(dbQueryUserOption))
		return nil, 0, err
	}

	dmUserList := make([]*dmuser.User, 0, len(dbUserList))
	for _, dbUser := range dbUserList {
		dmUserList = append(dmUserList, &dmuser.User{
			Id:        dbUser.Id,
			Username:  dbUser.Username,
			Number:    dbUser.Number,
			Password:  dbUser.Password,
			Status:    dbUser.Status,
			IsAdmin:   dbUser.IsAdmin,
			CreatedAt: dbUser.CreatedAt,
			UpdatedAt: dbUser.UpdatedAt,
			IsDel:     dbUser.IsDel,
			CreateNews: dbUser.CreateNews,
		})
	}

	return dmUserList, total, nil
}

func GetUserById(userId int64) (*dmuser.User, error) {
	if userId == 0 {
		logs.Info("[GetUserById] req.CheckParam(), err: userId == 0, userId: %v", userId)
		return nil, errors.New("param is incorrect")
	}

	dbUser, err := daouser.GetUserById(userId)
	if err != nil {
		logs.Error("[GetUserById] daouser.GetUserById, err: %v, userId: %v", err, userId)
		return nil, err
	}

	dmUser := dmuser.NewUser()
	dmUser.Id = dbUser.Id
	dmUser.Username = dbUser.Username
	dmUser.Number = dbUser.Number
	dmUser.Password = dbUser.Password
	dmUser.Status = dbUser.Status
	dmUser.IsAdmin = dbUser.IsAdmin
	dmUser.CreatedAt = dbUser.CreatedAt
	dmUser.UpdatedAt = dbUser.UpdatedAt
	dmUser.IsDel = dbUser.IsDel
	dmUser.CreateNews = dbUser.CreateNews

	return dmUser, nil
}
