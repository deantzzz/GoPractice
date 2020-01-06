//**练习 4.10：** 修改issu	es程序，根据问题的时间进行分类，比如不到一个月的、不到一年的、超过一年。
package main

import (
	"fmt"
	"gitgo/src/test/githubissues"
	"log"
	"os"
	"time"
)

func main() {
	result, err := githubissues.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	dateGroup := make(map[time.Month][]githubissues.Issue)
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		dateGroup[item.CreatedAt.Month()] = append(dateGroup[item.CreatedAt.Month()], *item)
	}
	for date := range dateGroup {
		fmt.Printf("%v\n", date)
		for _, item := range dateGroup[date] {
			fmt.Printf("#%-5d\t %9.9s\t %v\t %.55s\t \n",
				item.Number, item.User.Login, item.CreatedAt, item.Title)
		}
	}
}
