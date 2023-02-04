package config

import (
	cmn "filestore/common"
)

const (
	// CurrentStoreType : 设置当前文件的存储类型
	CurrentStoreType    = cmn.StoreLocal
	PwdSalt             = "*#890"
	TokenExpireDuration = 3600
	FokenExpireDuration = 10000
	JwtKey              = "filestore-key"
)
