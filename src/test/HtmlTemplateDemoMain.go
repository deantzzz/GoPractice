//**练习 4.14：** 创建一个web服务器，查询一次GitHub，然后生成BUG报告、里程碑和对应的用户信息。
package main

import (
	"gitgo/src/test/githubissues"
	"gitgo/src/test/html"
	"log"
	"net/http"
)

//!+
func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		result, err := githubissues.SearchIssues([]string{"commenter:gopherbot", "json", "encoder"})
		if err != nil {
			log.Fatal(err)
		}
		//将execute结果写入response
		if err := html.IssueList.Execute(writer, result); err != nil {
			log.Fatal(err)
		}
	})
	http.ListenAndServe("localhost:8000", nil)
}
