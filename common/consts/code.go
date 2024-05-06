package consts

import "google.golang.org/grpc/codes"

const (
	SUCCESS codes.Code = 1000

	ACCESS_TOKEN_INVALID   codes.Code = 2000
	ACCESS_EXPIRED         codes.Code = 2001
	ACCESS_DENY            codes.Code = 2002
	ACCESS_NOT_FOUND       codes.Code = 2003
	ACCESS_PWD_WRONG       codes.Code = 2004
	ACCESS_KEY_INVALID     codes.Code = 2005
	ACCOUNT_ALREADY_EXISTS codes.Code = 2006
	ACCESS_CODE_WRONG      codes.Code = 2007
	GROUP_ALREADY_EXISTS   codes.Code = 2008
	ACCESS_TOO_FAST        codes.Code = 2009
	DELETE_ADMIN_WRONG     codes.Code = 2010
	CANT_CREATE_GROUP      codes.Code = 2011
	CANT_CREATE_ACCOUNT    codes.Code = 2012
	REFRESH_EXPIRED        codes.Code = 2013

	GET_CODE_OVER_FREQUENCY         codes.Code = 2014
	ERROR_USERNAME_INVALID          codes.Code = 2015
	VALID_CODE_TIME                 codes.Code = 2016
	ERROR_INTERNAL_SERVER_OPERATION codes.Code = 2017

	NOT_FOUND codes.Code = 3001

	FAIL                         codes.Code = 4000
	WRONG_PARAM                  codes.Code = 4001
	NOT_FOUND_METHOD             codes.Code = 4004
	METADATA_NOT_FOUND           codes.Code = 4005
	AUTHORIZATION_NOT_FOUND      codes.Code = 4006
	ACCESSKEY_NOT_FOUND          codes.Code = 4007
	WRONG_CAPTCHA                codes.Code = 4008
	WECHAT_ERR_USERTOKEN_EXPIRED codes.Code = 4009
	DATA_EXIST                   codes.Code = 4010

	SERVER_WRONG codes.Code = 5000

	OPERATE_ARTICLE_STATUS_ERR codes.Code = 6000
	OPERATE_LABEL_STATUS_ERR   codes.Code = 6001

	// note: sdk error code %5d
	ERR_INIT_SDK_NOT_CLIENT  = 10001
	ERR_LOGININFO_NIL        = 10002
	ERR_JSON_MARSHAL         = 10003
	ERR_INIT_SDK_NOT_LOGINED = 10004
)

var WrongMessageEn = map[codes.Code]string{
	SUCCESS: "success",

	ACCESS_TOKEN_INVALID:   "invalid token",
	ACCESS_EXPIRED:         "user licence expired",
	REFRESH_EXPIRED:        "refresh licence expired",
	ACCESS_DENY:            "permission denied",
	ACCESS_NOT_FOUND:       "account does not exist",
	ACCESS_PWD_WRONG:       "incorrect username or password",
	ACCESS_KEY_INVALID:     "AccessKey is invalid",
	ACCOUNT_ALREADY_EXISTS: "user already exists",
	ACCESS_CODE_WRONG:      "verification code error",
	ACCESS_TOO_FAST:        "Access too fast",
	GROUP_ALREADY_EXISTS:   "user group already exists",
	CANT_CREATE_GROUP:      "Super administrator cannot create groups",
	CANT_CREATE_ACCOUNT:    "unable to create sub-account, please use root account to create one",

	GET_CODE_OVER_FREQUENCY:         "Get code too frequency",
	ERROR_USERNAME_INVALID:          "Username is invalid",
	VALID_CODE_TIME:                 "Verification code is expired",
	ERROR_INTERNAL_SERVER_OPERATION: "Internal Server Error",

	NOT_FOUND:                    "record not found",
	FAIL:                         "fail",
	NOT_FOUND_METHOD:             "request method not found",
	WRONG_PARAM:                  "param error",
	METADATA_NOT_FOUND:           "metadata not found",
	AUTHORIZATION_NOT_FOUND:      "authorization not found",
	ACCESSKEY_NOT_FOUND:          "accesskey not found",
	WRONG_CAPTCHA:                "wrong captcha",
	WECHAT_ERR_USERTOKEN_EXPIRED: "wechat user_token is expired",
	DATA_EXIST:                   "data already exists",

	DELETE_ADMIN_WRONG: "super administrator cannot be deleted",

	SERVER_WRONG: "Internal Server Error",

	OPERATE_ARTICLE_STATUS_ERR: "The article is on the shelf and cannot be operated",
	OPERATE_LABEL_STATUS_ERR:   "Tab is open and not operable",
	ERR_INIT_SDK_NOT_CLIENT:    "sdk client is nil",
	ERR_LOGININFO_NIL:          "reset time, logininfo is nil",
	ERR_JSON_MARSHAL:           "json marshal err",
	ERR_INIT_SDK_NOT_LOGINED:   "sdk client isn't logined",
}

var WrongMessageZh = map[codes.Code]string{
	SUCCESS: "请求成功",

	ACCESS_TOKEN_INVALID:   "无效token",
	ACCESS_EXPIRED:         "用户凭证过期",
	REFRESH_EXPIRED:        "刷新凭证过期",
	ACCESS_DENY:            "权限验证失败",
	ACCESS_NOT_FOUND:       "账户不存在",
	ACCESS_PWD_WRONG:       "用户名或密码不正确",
	ACCESS_KEY_INVALID:     "AccessKey无效",
	ACCOUNT_ALREADY_EXISTS: "用户已存在",
	ACCESS_TOO_FAST:        "太频繁了",
	ACCESS_CODE_WRONG:      "验证码错误",
	DELETE_ADMIN_WRONG:     "超级管理员不可删除",
	GROUP_ALREADY_EXISTS:   "用户组已存在",
	CANT_CREATE_GROUP:      "超级管理员不可创建组",
	CANT_CREATE_ACCOUNT:    "无法创建子账号,请用根账号创建",

	GET_CODE_OVER_FREQUENCY:         "获取验证码太频繁",
	ERROR_USERNAME_INVALID:          "用户名不合法",
	VALID_CODE_TIME:                 "验证码已过期",
	ERROR_INTERNAL_SERVER_OPERATION: "服务器内部错误",

	DATA_EXIST: "数据已经存在",

	NOT_FOUND: "记录未找到",

	FAIL:                         "请求失败",
	WRONG_PARAM:                  "参数错误",
	NOT_FOUND_METHOD:             "未找到请求方法",
	METADATA_NOT_FOUND:           "没找到metadata",
	AUTHORIZATION_NOT_FOUND:      "没找到验证头",
	ACCESSKEY_NOT_FOUND:          "没找到用户appid",
	WRONG_CAPTCHA:                "验证码错误",
	WECHAT_ERR_USERTOKEN_EXPIRED: "微信授权中用户的token已过期",

	SERVER_WRONG: "服务器错误",

	OPERATE_ARTICLE_STATUS_ERR: "文章处于上架状态，不可操作",
	OPERATE_LABEL_STATUS_ERR:   "标签处于开放状态，不可操作",
	ERR_INIT_SDK_NOT_CLIENT:    "客户端尚未完成初始化",
	ERR_LOGININFO_NIL:          "重置过期时间时，返回的登录信息为空",
	ERR_JSON_MARSHAL:           "json序列化错误",
	ERR_INIT_SDK_NOT_LOGINED:   "sdk尚未登录",
}
