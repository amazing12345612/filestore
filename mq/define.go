package mq

import (
	cmn "filestore/common"
)

// TransferData : 将要写到rabbitmq的数据的结构体
type TransferData struct {
	FileHash      string
	CurLocation   string //临时存储地址
	DestLocation  string
	DestStoreType cmn.StoreType
}
