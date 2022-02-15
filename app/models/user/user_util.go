package user

import (
	"fmt"
	"gohub/pkg/database"
	"math/rand"
)

// IsEmailExist 判断 Email 已被注册
func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// IsPhoneExist 判断手机号已被注册
func IsPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}

// IsLottery 返回当前抽奖的奖品ID，不会重复抽奖
func IsLottery() int {
	var count int64
	var choule int64
	status := 1
	allLotters := 22

	lotteryNum := rand.Intn(allLotters + 1)
	if lotteryNum == 0 {
		lotteryNum = lotteryNum + 1
	}
	database.DB.Model(Lottery{}).Where("status = ?", 1).Count(&choule)
	fmt.Println("总共抽了:", choule)
	if choule >= int64(allLotters) {
		return 100
	}
	database.DB.Model(Lottery{}).Where("id = ? AND status = ?", lotteryNum, status).Count(&count)

	for {
		if count == 0 {
			// fmt.Println("show the count:", count)
			database.DB.Model(Lottery{}).Where("id = ?", lotteryNum).Update("status", 1)
			break //如果count>=10则退出
		}
		lotteryNum = rand.Intn(allLotters + 1)
		if lotteryNum == 0 {
			lotteryNum = lotteryNum + 1
		}
		database.DB.Model(Lottery{}).Where("id = ? AND status = ?", lotteryNum, status).Count(&count)
		// fmt.Println("get again:", count)
	}

	return lotteryNum
}

// GetByPhone 通过手机号来获取用户
func GetByPhone(phone string) (userModel User) {
	database.DB.Where("phone = ?", phone).First(&userModel)
	return
}

// GetByMulti 通过 手机号/Email/用户名 来获取用户
func GetByMulti(loginID string) (userModel User) {
	database.DB.
		Where("phone = ?", loginID).
		Or("email = ?", loginID).
		Or("name = ?", loginID).
		First(&userModel)
	return
}
