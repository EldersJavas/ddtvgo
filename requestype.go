// Created by EldersJavas(EldersJavas&gmail.com)

package ddtvgo

var (
	RqSystem_Info               = NewRequest("System_Info", Post, JSON)
	RqSystem_Config             = NewRequest("System_Config", Post, JSON)
	RqSystem_Resources          = NewRequest("System_Resources", Post, JSON)
	RqSystem_QueryWebFirstStart = NewRequest("System_QueryWebFirstStart", Post, JSON)
	RqSystem_QueryUserState     = NewRequest("System_QueryUserState", Post, JSON)

	RqConfig_Transcod  = NewRequest("Config_Transcod", Post, JSON)
	RqConfig_FileSplit = NewRequest("Config_FileSplit", Post, JSON)
	RqConfig_DanmuRec  = NewRequest("Config_DanmuRec", Post, JSON)
	RqConfig_GetFollow = NewRequest("Config_GetFollow", Post, JSON)

	RqFile_GetAllFileList  = NewRequest("File_GetAllFileList", Post, JSON)
	RqFile_GetTypeFileList = NewRequest("File_GetTypeFileList", Post, JSON)
	RqFile_GetFilePathList = NewRequest("File_GetFilePathList", Post, JSON)
	//RqFile_GetFile         = NewRequest("File_GetFile", Post, FileStram)

	RqLogin = NewRequest("Login", Post, JSON)
	//RqLoginqr     = NewRequest("loginqr", Get, PNG)
	RqLogin_Reset = NewRequest("Login_Reset", Post, JSON)
	RqLogin_State = NewRequest("Login_State", Post, JSON)

	RqRec_RecordingInfo            = NewRequest("Rec_RecordingInfo", Post, JSON)
	RqRec_RecordingInfo_Lite       = NewRequest("Rec_RecordingInfo_Lite", Post, JSON)
	RqRec_RecordCompleteInfon      = NewRequest("Rec_RecordCompleteInfon", Post, JSON)
	RqRec_RecordCompleteInfon_Lite = NewRequest("Rec_RecordCompleteInfon_Lite", Post, JSON)
	RqRec_CancelDownload           = NewRequest("Rec_CancelDownload", Post, JSON)

	RqRoom_AllInfo     = NewRequest("Room_AllInfo", Post, JSON)
	RqRoom_SummaryInfo = NewRequest("Room_SummaryInfo", Post, JSON)
	RqRoom_Add         = NewRequest("Room_Add", Post, JSON)
	RqRoom_Del         = NewRequest("Room_Del", Post, JSON)
	RqRoom_AutoRec     = NewRequest("Room_AutoRec", Post, JSON)
	RqRoom_DanmuRec    = NewRequest("Room_DanmuRec", Post, JSON)

	RqUser_Search = NewRequest("User_Search", Post, JSON)
)
