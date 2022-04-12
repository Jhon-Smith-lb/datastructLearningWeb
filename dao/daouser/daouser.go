package daouser

import (
	"dataStructLearningWeb/dao"
	"dataStructLearningWeb/lib"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

func AddUser(req *dao.User, o orm.Ormer) (int64, error) {
	userId, err := o.Insert(req)
	if err != nil {
		logs.Error("[AddUser] o.Insert, err: %v, req: %v", err, lib.PointerToString(req))
		return 0, err
	}
	return userId, nil
}

func QueryUserList(req *dao.QueryUserOption) ([]*dao.User, int64, error) {
	o := orm.NewOrm()
	table := &dao.User{}
	resp := make([]*dao.User, 0)
	qs := o.QueryTable(&table)

	if req.Number != "" {
		qs = qs.Filter("number__contains", req.Number)
	}

	if req.Username != "" {
		qs = qs.Filter("username__contains", req.Username)
	}

	qs.Offset(req.Offset).Limit(req.Limit)
	_, err := qs.All(&resp)
	if err != nil {
		logs.Error("[QueryUserList] qs.All, err: %v, req: %v", err, lib.PointerToString(req))
		return nil, 0, err
	}
	return resp, 0, nil
}

func UpdateUser(req *dao.User, o orm.Ormer) error {
	mp := make(map[string]interface{}, 0)
	mp["username"] = req.Username
	mp["number"] = req.Number
	mp["password"] = req.Password
	mp["status"] = req.Status
	mp["is_admin"] = req.IsAdmin
	mp["is_del"] = req.IsDel
	table := &dao.User{}
	_, err := o.QueryTable(table).Filter("id", req.Id).Update(mp)
	if err != nil {
		logs.Error("[UpdateUser] o.Update, err: %v, req: %v", err, lib.PointerToString(req))
		return err
	}
	return nil
}

func GetUserById(userId int64) (*dao.User, error) {
	o := orm.NewOrm()
	user := &dao.User{Id: userId}
	if err := o.Read(user); err != nil {
		logs.Error("[GetUserById] o.Read, err: %v, userId: %v", err, userId)
		return nil, err
	}
	return user, nil
}

// 通过学号/工号获取用户
func GetUserByNumber(number string) (*dao.User, error) {
	o := orm.NewOrm()
	user := &dao.User{}
	if err := o.QueryTable(user).Filter("number", number).One(user); err != nil {
		logs.Error("[GetUserByNumber] o.Read, err: %v, number: %v", err, number)
		return nil, err
	}
	return user, nil
}
