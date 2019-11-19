package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func ranNum(size int) string {
	s := ""
	for i := 0; i < size; i++ {
		s += strconv.Itoa(rand.Intn(10))
	}
	return s
}

func punch() string {
	timeUnix := time.Now().UnixNano()
	rand.Seed(timeUnix)
	v := url.Values{}
	v.Set("appKey", "bf925ea0-a5e0-11e4-a67c-00163e024631")
	v.Set("enterpriseId", "1023080")
	v.Set("identification", "89403f4e9d65d7815be019c30d58535d")
	v.Set("latitude", "22.5505"+ranNum(10))
	v.Set("longitude", "113.9521"+ranNum(9))
	v.Set("timeStamp", strconv.FormatInt(timeUnix/1e6, 10))
	body := ioutil.NopCloser(strings.NewReader(v.Encode()))
	client := &http.Client{}
	reqest, err := http.NewRequest("POST", "http://oa.zuolin.com/evh/techpark/punch/punchClock", body)
	if err != nil {
		log.Fatal("Fatal error ", err)
	}
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	reqest.Header.Set("x-everhomes-device", "7D158FC8A2FC4CB8AD1C28E2396C55F7-101")
	reqest.Header.Set("cookie", "token=ziBuZPRrsMhFHYuPac5na6i02NGT76RlQ7-6hj_o7agyeCA1tVtgTBUrPZ_I4SN9je_LMI7geEo_B_IYnzhwyi3PMBVHG6v44xEMvPuBw0McFdkQptWMFTBlabaN_Ke-")
	resp, err := client.Do(reqest) //发送请求
	defer resp.Body.Close()        //一定要关闭resp.Body
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Fatal error ", err)
	}

	return string(content)
}

// w表示response对象，返回给客户端的内容都在对象里处理
// r表示客户端请求对象，包含了请求头，请求参数等等
func index(w http.ResponseWriter, r *http.Request) {

	// 往w里写入内容，就会在浏览器里输出
	fmt.Fprintf(w, punch())
	//fmt.Fprintf(w, ranNum(9))
	//fmt.Println(string(b.Bytes()[:]))
}

func main() {
	// 设置路由，如果访问/，则调用index方法
	http.HandleFunc("/punch", index)

	// 启动web服务，监听9090端口
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
