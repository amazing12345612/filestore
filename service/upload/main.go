package main

import (
	cfg "filestore/config"
	"filestore/handler"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	pwd, _ := os.Getwd()
	fmt.Println(pwd + " " + os.Args[0])
	http.Handle("/static/", http.FileServer(http.Dir(filepath.Join(pwd, "./"))))

	http.HandleFunc("/file/upload", handler.HTTPInterceptor(handler.UploadHandler))
	http.HandleFunc("/file/upload/suc", handler.HTTPInterceptor(handler.UploadSucHandler))
	http.HandleFunc("/file/meta", handler.HTTPInterceptor(handler.GetFileMetaHandler))
	http.HandleFunc("/file/query", handler.HTTPInterceptor(handler.FileQueryHandler))
	http.HandleFunc("/file/download", handler.HTTPInterceptor(handler.DownloadHandler))
	http.HandleFunc("/file/update", handler.HTTPInterceptor(handler.FileMetaUpdateHandler))
	http.HandleFunc("/file/delete", handler.HTTPInterceptor(handler.FileDeleteHandler))

	// 秒传接口
	http.HandleFunc("/file/fastupload", handler.HTTPInterceptor(handler.TryFastUploadHandler))

	// 分块上传接口
	http.HandleFunc("/file/mpupload/init", handler.HTTPInterceptor(handler.InitialMultipartUploadHandler))
	http.HandleFunc("/file/mpupload/uppart", handler.HTTPInterceptor(handler.UploadPartHandler))
	http.HandleFunc("/file/mpupload/complete", handler.HTTPInterceptor(handler.CompleteUploadHandler))

	// 用户相关接口

	http.HandleFunc("/", handler.SignInHandler)
	http.HandleFunc("/user/signup", handler.SignUpHandler)
	http.HandleFunc("/user/signin", handler.SignInHandler)
	http.HandleFunc("/user/info", handler.HTTPInterceptor(handler.UserInfoHandler))

	fmt.Printf("服务开始启动，监听[%s]中...\n", cfg.UploadServiceHost)
	// 启动服务并监听端口
	// err := http.ListenAndServe(":8080", nil)
	err := http.ListenAndServe(cfg.UploadServiceHost, nil)
	if err != nil {
		fmt.Printf("Failed to start server, err:%s", err.Error())
	}

}
