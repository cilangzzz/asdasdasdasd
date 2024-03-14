/**
  @author: cilang
  @qq: 1019383856
  @bili: https://space.bilibili.com/433915419
  @gitee: https://gitee.com/OpencvLZG
  @github: https://github.com/OpencvLZG
  @since: 2023/10/10
  @desc: //TODO
**/

package main

import (
	_ "github.com/go-sql-driver/mysql"
	"goDnParse/util"
	http2 "goDnParse/util/http"
	"net/http"
	"os"
)

func main() {
	// 初始化
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path := dir + "/config.json"
	print(path)
	config := util.ReadJson(path)
	print(config)
	////实例化数据库
	_, err = util.GetOrmEngine()
	if err != nil {
		panic(err)
	}
	// 需要校验 下载解析
	//dnHandle := http.NewServeMux()
	//dnHandle.HandleFunc("/", http2.GenerateDnLink)
	//http.HandleFunc("/download", http2.Auth(dnHandle))
	//// 需要校验 文件上传
	//uploadHandle := http.NewServeMux()
	//uploadHandle.HandleFunc("/", http2.FileUploadHandle)
	//http.HandleFunc("/fileUpload", http2.Auth(uploadHandle))
	//// 需要校验 默认页面
	//staticPage := http.FileServer(http.Dir("./static/page"))
	//staticHandle := http.NewServeMux()
	//staticHandle.Handle("/", http.StripPrefix("/static/page/", staticPage))
	//http.HandleFunc("/static/page/", http2.AuthLoading(staticHandle))
	// 默认可以访问
	// 错误页面
	//errPage := http.FileServer(http.Dir("./static/err"))
	//http.Handle("/static/err/", http.StripPrefix("/static/err", errPage))
	// 文件上传
	filePage := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", filePage))
	// 获取密钥
	//http.HandleFunc("/generateToken", http2.GenerateToken)
	http.HandleFunc("/", http2.DefaultHandler)
	http.HandleFunc("/user/add_user", http2.AddUser)
	http.HandleFunc("/user/get_user", http2.GetUser)
	http.HandleFunc("/favicon.ico", http2.DefaultIconHandler)
	err = http.ListenAndServe(config.Server.Address+":"+config.Server.Port, nil)
	if err != nil {
		return
	}
	err = http.ListenAndServeTLS(":443", "./cert/cilang.buzz.cert", "./cert/cilang.buzz.key", nil)
	if err != nil {
		return
	}
}
