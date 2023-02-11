package handler

import (
	"encoding/json"
	"strings"
	//"filestore-server/store/ceph"
	cmn "filestore/common"
	cfg "filestore/config"
	dblayer "filestore/db"
	"filestore/meta"
	"filestore/mq"
	"filestore/util"
	"fmt"
	_ "fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	_ "time"
	// "github.com/xilepeng/LeidianNetdisk/util"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			io.WriteString(w, "internal server error")
			return
		}
		io.WriteString(w, string(data))
	case "POST":
		file, head, err := r.FormFile("file")
		if err != nil {
			log.Fatal(err)
			return
		}
		defer file.Close()
		fileMeta := meta.FileMeta{
			FileName: head.Filename,
			Location: "./tmp/" + head.Filename,
			UploadAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		newFile, err := os.Create(fileMeta.Location)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer newFile.Close()

		fileMeta.FileSize, err = io.Copy(newFile, file)
		if err != nil {
			log.Fatal(err)
			return
		}

		//TODO：更新用户文件记录
		newFile.Seek(0, 0)
		fileMeta.FileSha1 = util.FileSha1(newFile)

		newFile.Seek(0, 0)

		if cfg.CurrentStoreType == cmn.StoreCeph {
			// 文件写入Ceph存储
			// data, _ := ioutil.ReadAll(newFile)
			// cephPath := "/ceph/" + fileMeta.FileSha1
			// _ = ceph.PutObject("userfile", cephPath, data)
			// fileMeta.Location = cephPath
		} else if cfg.CurrentStoreType == cmn.StoreLocal {
			// 文件写入存储
			ossPath := "./oss/" + fileMeta.FileName
			// 判断写入OSS为同步还是异步
			if !cfg.AsyncTransferEnable {
				transferFile, err := os.Create(ossPath)
				if err != nil {
					log.Fatal(err)
					return
				}
				defer transferFile.Close()

				_, err = io.Copy(transferFile, newFile)
				if err != nil {
					log.Fatal(err)
					return
				}

				if err != nil {
					fmt.Println(err.Error())
					w.Write([]byte("Upload failed!"))
					return
				}
				fileMeta.Location = ossPath
			} else {
				// 写入异步转移任务队列
				data := mq.TransferData{
					FileHash:      fileMeta.FileSha1,
					CurLocation:   fileMeta.Location,
					DestLocation:  ossPath,
					DestStoreType: cmn.StoreLocal,
				}
				pubData, _ := json.Marshal(data)
				pubSuc := mq.Publish(
					cfg.TransExchangeName,
					cfg.TransOSSRoutingKey,
					pubData,
				)
				if !pubSuc {
					// TODO: 当前发送转移信息失败，稍后重试
				}
			}
		}

		meta.UpdateFileMetaDB(fileMeta)
		r.ParseForm()
		username := r.Form.Get("username")
		res := dblayer.OnUserFileUploadFinished(username, fileMeta.FileSha1, fileMeta.FileName, fileMeta.FileSize)
		if res {
			http.Redirect(w, r, "/static/view/home.html", http.StatusFound)
		}
		w.Write([]byte("Upload failed"))

	}

}

// Upload finished 上传成功！
func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Upload finished !")
}

// GetFileMetaHandler : 获取文件元信息
func GetFileMetaHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// 获取文件的filehash
	filehash := r.Form["filehash"][0]
	// fMeta := meta.GetFileMeta(filehash)
	fMeta, err := meta.GetFileMetaDB(filehash)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(fMeta) //结构体转换为json

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

// FileQueryHandler : 查询批量的文件元信息
func FileQueryHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	limitCnt, _ := strconv.Atoi(r.Form.Get("limit"))
	username := r.Form.Get("username")
	userFiles, err := dblayer.QueryUserFileMetas(username, limitCnt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(userFiles)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

// DownloadHandler : 文件下载接口
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fsha1 := r.Form.Get("filehash")
	fm, _ := meta.GetFileMetaDB(fsha1)

	f, err := os.Open(fm.Location)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octect-stream")
	// attachment表示文件将会提示下载到本地，而不是直接在浏览器中打开
	w.Header().Set("content-disposition", "attachment; filename=\""+fm.FileName+"\"")
	w.Write(data)
}

// FileMetaUpdateHandler : 更新元信息接口（重命名）
func FileMetaUpdateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	opType := r.Form.Get("op")
	fileSha1 := r.Form.Get("filehash")
	newFileName := r.Form.Get("filename")

	if opType != "0" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	curFileMeta := meta.GetFileMeta(fileSha1)
	curFileMeta.FileName = newFileName
	meta.UpdateFileMeta(curFileMeta)

	data, err := json.Marshal(curFileMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// FileDeleteHandler : 删除文件及元信息
func FileDeleteHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fsha1 := r.Form.Get("filehash")

	// 删除文件
	fMeta := meta.GetFileMeta(fsha1)
	os.Remove(fMeta.Location)

	// 删除文件元信息
	meta.RemoveFileMeta(fsha1)

	w.WriteHeader(http.StatusOK)

}

// TryFastUploadHandler : 尝试秒传接口
func TryFastUploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// 1. 解析请求参数
	username := r.Form.Get("username")
	filehash := r.Form.Get("filehash")
	filename := r.Form.Get("filename")
	filesize, _ := strconv.Atoi(r.Form.Get("filesize"))

	// 2. 从文件表中查询相同hash的文件记录
	fileMeta, err := meta.GetFileMetaDB(filehash)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 3. 查不到记录则返回秒传失败
	if fileMeta.FileSha1 == "" {
		resp := util.RespMsg{
			Code: -1,
			Msg:  "秒传失败，请访问普通上传接口",
		}
		w.Write(resp.JSONBytes())
		return
	}

	// 4. 上传过则将文件信息写入用户文件表， 返回成功
	suc := dblayer.OnUserFileUploadFinished(
		username, filehash, filename, int64(filesize))
	if suc {
		resp := util.RespMsg{
			Code: 0,
			Msg:  "秒传成功",
		}
		w.Write(resp.JSONBytes())
		return
	}
	resp := util.RespMsg{
		Code: -2,
		Msg:  "秒传失败，请稍后重试",
	}
	w.Write(resp.JSONBytes())
	return
}

// DownloadURLHandler : 生成文件的下载地址
func DownloadURLHandler(w http.ResponseWriter, r *http.Request) {
	filehash := r.Form.Get("filehash")
	// 从文件表查找记录
	row, _ := dblayer.GetFileMeta(filehash)

	// TODO: 判断文件存在OSS，还是Ceph，还是在本地
	if strings.HasPrefix(row.FileAddr.String, "./tmp") {
		username := r.Form.Get("username")
		token := r.Form.Get("token")
		tmpUrl := fmt.Sprintf("http://%s/file/download?filehash=%s&username=%s&token=%s",
			r.Host, filehash, username, token)
		w.Write([]byte(tmpUrl))
	}
}
