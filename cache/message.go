package cache

import (
	"github.com/go-redis/redis/v7"
	"os"
	"strconv"
	"time"
)

// mobileCodeKey 获取缓存中手机验证码的 Key 值。
func mobileCodeKey(mobile string) string {
	return "rpc_user_mb_code_" + mobile
}

// CheckMobile 检查 mobile 是否可以再次接收验证码。
func CheckMobile(mobile string) (valid bool, err error) {
	err = Client.Get(mobileCodeKey(mobile)).Err()
	if err == redis.Nil {
		valid = true
		err = nil
		return
	}
	return
}

// CacheMobileCode 缓存发送的验证码。
func StoreMobileCode(mobile, code string) (err error) {
	m, err := strconv.Atoi(os.Getenv("SMS_CODE_MOBILE_RECV_MINUTE"))
	if err != nil {
		return
	}
	if err = Client.Set(mobileCodeKey(mobile), code, time.Duration(m)*time.Minute).Err(); err != nil {
		return
	}
	return
}

// VerifyMobileCode 检测手机和验证码是否匹配。
func VerifyMobileCode(mobile, code string) (valid bool, err error) {
	res, err := Client.Get(mobileCodeKey(mobile)).Result()
	if err == redis.Nil {
		err = nil
		return
	}
	if err != nil {
		return
	}
	if res == code {
		valid = true
		return
	}
	return
}
