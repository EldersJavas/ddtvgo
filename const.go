// Created by EldersJavas(EldersJavas&gmail.com)

package ddtvgo

/* 错误码
0	成功
-2	UID不存在
6000	登陆信息失效
6001	登陆验证失败
6002	APIsig计算校验失败
7000	操作失败
*/

type ErrCode int

const (
	CodeSuccess      ErrCode = 0
	CodeNoUID        ErrCode = -2
	CodeLoginInvalid ErrCode = 6000
	CodeLoginFail    ErrCode = 6001
	CodeSigFail      ErrCode = 6002
	CodeFailure      ErrCode = 700
)

func (e ErrCode) ToString() string {
	switch e {
	case CodeSuccess:
		return "成功"
	case CodeNoUID:
		return "UID不存在"
	case CodeLoginInvalid:
		return "登陆信息失效"
	case CodeLoginFail:
		return "登陆验证失败"
	case CodeSigFail:
		return "APIsig计算校验失败"
	case CodeFailure:
		return "操作失败"
	}
	return ""
}

/* 请求方法
1 Post
2 Get
*/

type RequestMode int

const (
	Post RequestMode = 1
	Get  RequestMode = 2
)

func (e RequestMode) ToString() string {
	switch e {
	case Post:
		return "POST"
	case Get:
		return "GET"

	}
	return ""
}

/* 请求方法
1 JSON
2 PNG
3 FileStram
*/

type RequestReturnType int

const (
	JSON      RequestReturnType = 1
	PNG       RequestReturnType = 2
	FileStram RequestReturnType = 3
)

func (e RequestReturnType) ToString() string {
	switch e {
	case JSON:
		return ""
	case PNG:
		return ""
	case FileStram:
		return ""
	}
	return ""
}

/*
方式	名称	返回内容	解释
POST	System_Info	JSON	获取系统运行情况
POST	System_Config	JSON	获取系统配置文件信息
POST	System_Resources	JSON	获取系统硬件资源使用情况
POST	System_QueryWebFirstStart	JSON	返回一个可以自行设定的初始化状态值
POST	System_SetWebFirstStart	JSON	设置初始化状态值
POST	System_QueryUserState	JSON	查询B站接口返回数据判断用户登陆状态是否有效
POST	Config_Transcod	JSON	设置自动转码总开关
POST	Config_FileSplit	JSON	根据文件大小自动切片
POST	Config_DanmuRec	JSON	弹幕录制总共开关(包括礼物、舰队、SC)
POST	Config_GetFollow	JSON	导入关注列表中的V
POST	File_GetAllFileList	JSON	获取已录制的文件列表
POST	File_GetTypeFileList	JSON	分类获取已录制的文件总列表
POST	File_GetFilePathList	JSON	根据文件树结构返回已录制的文件总列表
POST	File_GetFile	FileStram	下载对应的文件
POST	Login	JSON	WEB登陆
GET	loginqr	PNG	在提示登陆的情况下获取用于的登陆二维码
POST	Login_Reset	JSON	重新登陆哔哩哔哩账号
POST	Login_State	JSON	查询内部登陆状态
POST	Rec_RecordingInfo	JSON	获取下载中的任务情况详细情况
POST	Rec_RecordingInfo_Lite	JSON	获取下载中的任务情况简略情况
POST	Rec_RecordCompleteInfon	JSON	获取已经完成的任务详细情况
POST	Rec_RecordCompleteInfon_Lite	JSON	获取已经完成的任务简略情况
POST	Rec_CancelDownload	JSON	取消某个下载任务
POST	Room_AllInfo	JSON	获取房间详细配置信息
POST	Room_SummaryInfo	JSON	获取房间简要配置信息
POST	Room_Add	JSON	增一个加房间配置
POST	Room_Del	JSON	删除一个房间配置
POST	Room_AutoRec	JSON	修改房间自动录制配置信息
POST	Room_DanmuRec	JSON	修改房间弹幕录制配置信息
POST	User_Search	JSON	通过B站搜索搜索直播用户
*/
var (
	System_Info               = NewRequest("System_Info", Post, JSON)
	System_Config             = NewRequest("System_Config", Post, JSON)
	System_Resources          = NewRequest("System_Resources", Post, JSON)
	System_QueryWebFirstStart = NewRequest("System_QueryWebFirstStart", Post, JSON)
	System_QueryUserState     = NewRequest("System_QueryUserState", Post, JSON)

	Config_Transcod  = NewRequest("Config_Transcod", Post, JSON)
	Config_FileSplit = NewRequest("Config_FileSplit", Post, JSON)
	Config_DanmuRec  = NewRequest("Config_DanmuRec", Post, JSON)
	Config_GetFollow = NewRequest("Config_GetFollow", Post, JSON)

	File_GetAllFileList  = NewRequest("File_GetAllFileList", Post, JSON)
	File_GetTypeFileList = NewRequest("File_GetTypeFileList", Post, JSON)
	File_GetFilePathList = NewRequest("File_GetFilePathList", Post, JSON)
	File_GetFile         = NewRequest("File_GetFile", Post, FileStram)

	Login       = NewRequest("Login", Post, JSON)
	loginqr     = NewRequest("loginqr", Get, PNG)
	Login_Reset = NewRequest("Login_Reset", Post, JSON)
	Login_State = NewRequest("Login_State", Post, JSON)

	Rec_RecordingInfo            = NewRequest("Rec_RecordingInfo", Post, JSON)
	Rec_RecordingInfo_Lite       = NewRequest("Rec_RecordingInfo_Lite", Post, JSON)
	Rec_RecordCompleteInfon      = NewRequest("Rec_RecordCompleteInfon", Post, JSON)
	Rec_RecordCompleteInfon_Lite = NewRequest("Rec_RecordCompleteInfon_Lite", Post, JSON)
	Rec_CancelDownload           = NewRequest("Rec_CancelDownload", Post, JSON)

	Room_AllInfo     = NewRequest("Room_AllInfo", Post, JSON)
	Room_SummaryInfo = NewRequest("Room_SummaryInfo", Post, JSON)
	Room_Add         = NewRequest("Room_Add", Post, JSON)
	Room_Del         = NewRequest("Room_Del", Post, JSON)
	Room_AutoRec     = NewRequest("Room_AutoRec", Post, JSON)
	Room_DanmuRec    = NewRequest("Room_DanmuRec", Post, JSON)

	User_Search = NewRequest("User_Search", Post, JSON)
)
