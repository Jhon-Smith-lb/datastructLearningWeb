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

	qs = qs.Filter("is_del", 0)

	count, err := qs.Count()
	if err != nil {
		logs.Error("[QueryUserList] qs.Count, err: %v, req: %v", err, lib.PointerToString(req))
		return nil, 0, err
	}

	qs = qs.Offset(req.Offset).Limit(req.Limit)
	_, err = qs.All(&resp)
	if err != nil {
		logs.Error("[QueryUserList] qs.All, err: %v, req: %v", err, lib.PointerToString(req))
		return nil, 0, err
	}
	return resp, count, nil
}

func UpdateUser(req *dao.User, o orm.Ormer) error {
	mp := make(map[string]interface{}, 0)
	if req.Username != "" {
		mp["username"] = req.Username
	}
	if req.Number != "" {
		mp["number"] = req.Number
	}
	if req.Password != nil {
		mp["password"] = *req.Password
	}
	if req.Status != nil {
		mp["status"] = *req.Status
	}
	if req.IsAdmin != nil {
		mp["is_admin"] = *req.IsAdmin
	}
	if req.IsDel != nil {
		mp["is_del"] = *req.IsDel
	}
	if req.CreateNews != nil {
		mp["create_news"] = *req.CreateNews
	}
	
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
