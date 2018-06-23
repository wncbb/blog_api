package captcha

import (
	api_define "blog_api/api/define"
	"fmt"
	"strconv"
	"time"

	"blog_api/util/gostrgen"

	"blog_api/config"

	"github.com/go-redis/redis"
	"github.com/mojocn/base64Captcha"
	"github.com/pkg/errors"
)

//customizeRdsStore An object implementing Store interface
type customizeRdsStore struct {
	redisClient *redis.Client
}

// customizeRdsStore implementing Set method of  Store interface
func (s *customizeRdsStore) Set(id string, value string) {
	err := s.redisClient.Set(id, value, time.Minute*10).Err()
	if err != nil {
		panic(err)
	}
}

// customizeRdsStore implementing Get method of  Store interface
func (s *customizeRdsStore) Get(id string, clear bool) (value string) {
	val, err := s.redisClient.Get(id).Result()
	if err != nil {
		// TODO 这里区分不存在还是出错
		return ""
	}
	if clear {
		err := s.redisClient.Del(id).Err()
		if err != nil {
			panic(err)
		}
	}
	return val
}

func Init() {
	redisConfig := config.GetTempRedisConfig()
	//create redis client
	db, err := strconv.Atoi(redisConfig.Db)
	if err != nil {
		panic(fmt.Sprintf("Init temp redis err:%v", err))
	}
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Ip + ":" + redisConfig.Port,
		Password: redisConfig.Password,
		DB:       db,
	})
	//init redis store
	customeStore := customizeRdsStore{client}

	base64Captcha.SetCustomStore(&customeStore)

}

//ConfigJsonBody json request body.
type ConfigJsonBody struct {
	Id              string
	CaptchaType     string
	VerifyValue     string
	ConfigAudio     base64Captcha.ConfigAudio
	ConfigCharacter base64Captcha.ConfigCharacter
	ConfigDigit     base64Captcha.ConfigDigit
}

type GetCaptchaData struct {
	CaptchaData string
	CaptchaId   string
}

func GetCaptcha() (*GetCaptchaData, error) {
	var err error
	postParameters := &ConfigJsonBody{}
	postParameters.Id, err = gostrgen.RandGen(10, gostrgen.LowerUpperDigit, "", "")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	postParameters.ConfigDigit = base64Captcha.ConfigDigit{
		Height:     80,
		Width:      200,
		CaptchaLen: 6,
		MaxSkew:    4,
		DotCount:   10,
	}

	var config interface{}
	switch postParameters.CaptchaType {
	case "audio":
		config = postParameters.ConfigAudio
	case "character":
		config = postParameters.ConfigCharacter
	default:
		config = postParameters.ConfigDigit
	}
	captchaId, captcaInterfaceInstance := base64Captcha.GenerateCaptcha(postParameters.Id, config)
	base64blob := base64Captcha.CaptchaWriteToBase64Encoding(captcaInterfaceInstance)

	getCaptchaData := &GetCaptchaData{
		CaptchaData: base64blob,
		CaptchaId:   captchaId,
	}
	return getCaptchaData, nil
}

func VerifyCaptcha(captchaId string, captchaCode string) (string, error) {

	verifyRst := base64Captcha.VerifyCaptcha(captchaId, captchaCode)
	if !verifyRst {
		return "验证码错误", nil
	}

	return api_define.ShowMsgSuccess, nil

}
