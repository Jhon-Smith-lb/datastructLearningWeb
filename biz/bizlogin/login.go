package bizlogin

import (
	"dataStructLearningWeb/dao/daouser"
	"dataStructLearningWeb/dm/dmlogin"
	"dataStructLearningWeb/lib/libredis"
	"errors"
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/redis.v5"
)

func Login(number, password string) (string, error) {
	if number == "" || password == "" {
		logs.Info("[Login] err: number == \"\" || password == \"\"")
		return "", errors.New("param is incorrect")
	}

	// 查询用户是否存在
	user, err := daouser.GetUserByNumber(number)
	if err != nil {
		logs.Error("[Login] daouser.GetUserByNumber, err: %v, number: %v", err, number)
		return "", err
	}

	// 进行登录频控
	key := fmt.Sprintf("datastructLearningWeb:login:user_id %v", user.Id)
	var count int64
	count, err = libredis.GetInt64(key)
	if err != nil && err != redis.Nil {
		logs.Error("[Login] libredis.GetStr, err: %v, key: %v", err, key)
		return "", err
	}

	if count >= 10 {
		return "", errors.New("登录过于频繁")
	}

	// 校验密码
	if *user.Password != password {
		logs.Info("[Login] *user.Password != password, err: 密码错误, number: %v", number)
		return "", errors.New("账号或密码错误")
	}

	// 校验账号状态
	if *user.Status != 0 {
		logs.Info("[Login] *user.Status != 0, err: 账号状态为不可用, number: %v", number)
		return "", errors.New("账号状态为不可用")
	}

	// 生成token
	token, err := generateToken(user.Id, user.Username)
	if err != nil {
		logs.Error("[Login] generateToken, err: %v, id: %v, username: %v", err, user.Id, user.Username)
		return "", err
	}

	var ttlTime time.Duration
	if count == 0 {
		// 第一次登录
		ttlTime = time.Minute * 5
	} else {
		ttlTime, err = libredis.TTL(key)
		if err != nil {
			logs.Error("[Login] libredis.TTL, err: %v, key: %v", err, key)
			return "", err
		}
	}

	// 增加登录次数记录
	count += 1
	countStr := fmt.Sprintf("%v", count)

	if err := libredis.SetStr(key, countStr, ttlTime); err != nil {
		logs.Error("[Login] libredis.SetStr, err: %v, key: %v, value: %v, ttlTime: %v", err, key, countStr, ttlTime)
		return "", err
	}

	return token, nil
}

func generateToken(id int64, username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(7 * 24 * time.Hour)
	issuer := dmlogin.ISSUER
	claims := dmlogin.Claims{
		ID:       id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    issuer,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(dmlogin.SIGNED_KEY))
	return token, err
}
