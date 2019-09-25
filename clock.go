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
	"flag"
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
	reqest.Header.Set("cookie", "token=D4CSFyuOtl0GHP98kPqAyTnRn6yTuLYBGb3mzy3Mw5cyeCA1tVtgTBUrPZ_I4SN9je_LMI7geEo_B_IYnzhwyi3PMBVHG6v44xEMvPuBw0M-lP9pFEyGD4TEAybjlHf7")
	resp, err := client.Do(reqest) //发送请求
	defer resp.Body.Close()        //一定要关闭resp.Body
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Fatal error ", err)
	}

	return string(content)
}

func main() {
	hour := flag.Int("hs", 9, "hours")
	minute := flag.Int("ms", 59, "minute")
	flag.Parse()
	myTicker:=time.NewTicker(time.Second * 23)		//设置时间周期
	for{
		nowTime:=<-myTicker.C		//当前时间
		if nowTime.Hour()==*hour && nowTime.Minute()==*minute{
			punch()
			break
		}
	}
}