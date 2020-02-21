package main

import (
	"fmt"
	"getReslut/config"
	"getReslut/public"
	"getReslut/public/dbConnect"
	"getReslut/result"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	nowTime "time"

	"github.com/Unknwon/goconfig"
	"github.com/gin-gonic/gin"
)

//自開彩資料
type IssueData struct {
	Issue    int64 `form:"issue" json:"issue" binding:"exists"`
	TypGroup int64 `form:"typeGroup" json:"typeGroup" binding:"exists"`
	Type     int64 `form:"type" json:"type" binding:"exists"`
	Status   bool  `form:"status" json:"status" binding:"exists"`
}

//區塊鍊資料
type BlockChainIssueData struct {
	TypGroup int64 `form:"typeGroup" json:"typeGroup" binding:"exists"`
	Type     int64 `form:"type" json:"type" binding:"exists"`
	Issue    int64 `form:"issue" json:"issue" binding:"exists"`
	Status   bool  `form:"status" json:"status" binding:"exists"`
	Time     int64 `form:"time" json:"time" binding:"exists"`
}

var cfg *goconfig.ConfigFile

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}
func init() {
	fileAddr := "/"
	if config.GetDeBugMode() {
		fileAddr = "./"
	} else {
		fileAddr = getCurrentDirectory() + "/"
	}
	config, err := goconfig.LoadConfigFile(fileAddr + "configData.conf") //加载配置文件
	if err != nil {
		fmt.Println("SERVER get config file error")
		os.Exit(-1)
	}
	cfg = config
}

func main() {
	router := gin.Default()
	router.POST("/getResult", func(c *gin.Context) {
		var json IssueData

		c.ShouldBindJSON(&json)
		public.Println(fmt.Sprint("自開彩開獎  ", json))
		if json.Status == true {
			thisLotteryTypeGroup := fmt.Sprint(json.TypGroup)
			thisLotteryType := fmt.Sprint(json.Type)
			thisLotteryIssue := fmt.Sprint(json.Issue)
			status := json.Status
			//撈RTP設定
			rtpData := dbConnect.GetRtpSetting(thisLotteryTypeGroup, thisLotteryType)
			if rtpData["state"].(int) == 0 {
				c.JSON(http.StatusOK, gin.H{"status": false, "error": 2, "result": make([]string, 0)})
				return
			}
			thisRatio, _ := strconv.ParseFloat(rtpData["result"].(string), 64)
			//撈注單
			betOrderData := dbConnect.Run(thisLotteryTypeGroup, thisLotteryType, thisLotteryIssue)
			if betOrderData["state"].(int) == 0 {
				c.JSON(http.StatusOK, gin.H{"status": false, "error": 2, "result": make([]string, 0)})
			} else {
				if len(betOrderData["result"].([]interface{})) == 0 { //有期號 但無注單
					randResult := result.GetRandResult(status, thisLotteryTypeGroup)
					public.Println(fmt.Sprint("有期號 但無注單 -------> ", randResult))
					c.JSON(http.StatusOK, gin.H{"status": true, "error": 0, "result": randResult["thisOpenResult"]})
					return
				}
				openResult := result.Run(betOrderData, status, rtpData, thisLotteryTypeGroup, thisLotteryType, thisLotteryIssue, 0)
				amountData := dbConnect.GetRtpSetting(thisLotteryTypeGroup, thisLotteryType)
				amount := amountData["amount"].(float64)
				bonus := amountData["bonus"].(float64)
				isUpdate := true
				t := nowTime.Now()
				now := fmt.Sprintf("%4d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

				if amount >= 100000000 || bonus >= 100000000 {
					amount = (amount + openResult["amount"].(float64)) / 100
					bonus = (bonus + openResult["bonus"].(float64)) / 100

					isUpdate = dbConnect.SetAmount(amount, bonus, thisLotteryTypeGroup, thisLotteryType, thisRatio, now)
					if isUpdate == false {
						c.JSON(http.StatusOK, gin.H{"status": false, "error": 4, "result": make([]string, 0)})
						return
					}
					isUpdate = dbConnect.SetRtpRecord(openResult["amount"].(float64), openResult["bonus"].(float64), thisLotteryTypeGroup, thisLotteryType, thisRatio, now, thisLotteryIssue, bonus/amount)
					if isUpdate == false {
						c.JSON(http.StatusOK, gin.H{"status": false, "error": 4, "result": make([]string, 0)})
						return
					}
				} else {

					amount = (amount + openResult["amount"].(float64))
					bonus = (bonus + openResult["bonus"].(float64))
					isUpdate = dbConnect.UpdateAmount(amount, bonus, thisLotteryTypeGroup, thisLotteryType, thisRatio, now)
					if isUpdate == false {
						c.JSON(http.StatusOK, gin.H{"status": false, "error": 4, "result": make([]string, 0)})
						return
					}
					isUpdate = dbConnect.SetRtpRecord(openResult["amount"].(float64), openResult["bonus"].(float64), thisLotteryTypeGroup, thisLotteryType, thisRatio, now, thisLotteryIssue, bonus/amount)
					if isUpdate == false {
						c.JSON(http.StatusOK, gin.H{"status": false, "error": 4, "result": make([]string, 0)})
						return
					}
				}

				if len(openResult) == 0 {
					public.Println(fmt.Sprint("openResult -------> ", 0))
					c.JSON(http.StatusOK, gin.H{"status": false, "error": 3, "result": make([]string, 0)})
				} else {
					public.Println(fmt.Sprint("openResult -------> ", openResult))
					c.JSON(http.StatusOK, gin.H{"status": true, "error": 0, "result": openResult["thisOpenResult"]})
				}
			}
		} else {
			c.JSON(http.StatusOK, gin.H{"status": false, "error": 1, "result": make([]string, 0)})
		}
	})

	router.POST("/getBlockChainResult", func(c *gin.Context) {
		var json BlockChainIssueData

		c.ShouldBindJSON(&json)
		public.Println(fmt.Sprint("區塊鍊開獎  ", json))

		public.Println(fmt.Sprint("123 ------->   ", json.Status == true))
		public.Println(fmt.Sprint("123 ------->   ", config.GetgameCodeMap()))

		if json.Status == true {

			thisLotteryTypeGroup := fmt.Sprint(json.TypGroup)
			thisLotteryType := fmt.Sprint(json.Type)
			thisLotteryIssue := fmt.Sprint(json.Issue)

			LotteryDrawTimeData := dbConnect.GetLotteryDrawTime(thisLotteryType, thisLotteryIssue)
			if LotteryDrawTimeData["datedraw"] == "0" {
				c.JSON(http.StatusOK, gin.H{"status": false, "error": 2, "chainCode": "", "result": make([]string, 0)})
				return
			}
			datedrawTime, _ := strconv.ParseInt(LotteryDrawTimeData["datedraw"].(string), 10, 64)
			status := json.Status
			//撈RTP設定
			rtpData := dbConnect.GetRtpSetting(thisLotteryTypeGroup, thisLotteryType)

			if rtpData["state"].(int) == 0 {
				c.JSON(http.StatusOK, gin.H{"status": false, "error": 2, "chainCode": "", "result": make([]string, 0)})
				return
			}

			thisRatio, _ := strconv.ParseFloat(rtpData["result"].(string), 64)

			//撈注單
			betOrderData := dbConnect.Run(thisLotteryTypeGroup, thisLotteryType, thisLotteryIssue)
			if betOrderData["state"].(int) == 0 {
				c.JSON(http.StatusOK, gin.H{"status": false, "error": 2, "chainCode": "", "result": make([]string, 0)})
			} else {
				if len(betOrderData["result"].([]interface{})) == 0 { //有期號 但無注單
					randResult := result.GetChainRandResult(status, thisLotteryTypeGroup, datedrawTime)
					public.Println(fmt.Sprint("有期號 但無注單 -------> ", randResult))

					if len(randResult) == 0 {
						public.Println(fmt.Sprint("openResult -------> ", 0))
						c.JSON(http.StatusOK, gin.H{"status": false, "error": 3, "chainCode": "", "result": make([]string, 0)})
					} else {
						public.Println(fmt.Sprint("有期號 但無注單 openResult -------> ", randResult))
						dbConnect.SetLotteryDrawChainCode(thisLotteryType, thisLotteryIssue, randResult["chainCode"].(string), randResult["thisOpenResult"].([]string), datedrawTime)
						c.JSON(http.StatusOK, gin.H{"status": true, "error": 0, "chainCode": randResult["chainCode"].(string), "result": randResult["thisOpenResult"]})
					}

					return
				}
				public.Println(fmt.Sprint("123 ------->   ", json))

				if datedrawTime <= 0 {
					c.JSON(http.StatusOK, gin.H{"status": true, "error": 3, "chainCode": "", "result": make([]string, 0)})
					return
				}
				openResult := result.Run(betOrderData, status, rtpData, thisLotteryTypeGroup, thisLotteryType, thisLotteryIssue, datedrawTime)
				if len(openResult) > 0 {
					amountData := dbConnect.GetRtpSetting(thisLotteryTypeGroup, thisLotteryType)
					amount := amountData["amount"].(float64)
					bonus := amountData["bonus"].(float64)
					isUpdate := true
					t := nowTime.Now()
					now := fmt.Sprintf("%4d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

					if amount >= 100000000 || bonus >= 100000000 {
						amount = (amount + openResult["amount"].(float64)) / 100
						bonus = (bonus + openResult["bonus"].(float64)) / 100

						isUpdate = dbConnect.SetAmount(amount, bonus, thisLotteryTypeGroup, thisLotteryType, thisRatio, now)
						if isUpdate == false {
							c.JSON(http.StatusOK, gin.H{"status": false, "error": 4, "chainCode": "", "result": make([]string, 0)})
							return
						}
						isUpdate = dbConnect.SetRtpRecord(openResult["amount"].(float64), openResult["bonus"].(float64), thisLotteryTypeGroup, thisLotteryType, thisRatio, now, thisLotteryIssue, bonus/amount)
						if isUpdate == false {
							c.JSON(http.StatusOK, gin.H{"status": false, "error": 4, "chainCode": "", "result": make([]string, 0)})
							return
						}
					} else {

						amount = (amount + openResult["amount"].(float64))
						bonus = (bonus + openResult["bonus"].(float64))
						isUpdate = dbConnect.UpdateAmount(amount, bonus, thisLotteryTypeGroup, thisLotteryType, thisRatio, now)
						if isUpdate == false {
							c.JSON(http.StatusOK, gin.H{"status": false, "error": 4, "chainCode": "", "result": make([]string, 0)})
							return
						}
						isUpdate = dbConnect.SetRtpRecord(openResult["amount"].(float64), openResult["bonus"].(float64), thisLotteryTypeGroup, thisLotteryType, thisRatio, now, thisLotteryIssue, bonus/amount)
						if isUpdate == false {
							c.JSON(http.StatusOK, gin.H{"status": false, "error": 4, "chainCode": "", "result": make([]string, 0)})
							return
						}
					}
				}

				if len(openResult) == 0 {
					public.Println(fmt.Sprint("openResult -------> ", 0))
					c.JSON(http.StatusOK, gin.H{"status": false, "error": 3, "chainCode": "", "result": make([]string, 0)})
				} else {
					public.Println(fmt.Sprint("openResult -------> ", openResult))
					c.JSON(http.StatusOK, gin.H{"status": true, "error": 0, "chainCode": openResult["chainCode"], "result": openResult["thisOpenResult"]})
					dbConnect.SetLotteryDrawChainCode(thisLotteryType, thisLotteryIssue, openResult["chainCode"].(string), openResult["thisOpenResult"].([]string), datedrawTime)
				}
			}
		} else {
			c.JSON(http.StatusOK, gin.H{"status": false, "error": 1, "chainCode": "", "result": make([]string, 0)})
		}

	})

	// ConfigData := config.GetConfig()
	// port := ":" + ConfigData["serverPort"].(string)

	port, _ := cfg.GetValue("port", "serverPort") //读取单个值

	err := cfg.Reload() //重载配置
	if err != nil {
		fmt.Printf("reload config file error: %s", err)
	}
	portData := ":" + port
	config.ConfigInit(cfg)
	//fmt.Println("cfg", cfg)
	fmt.Println(portData)
	router.Run(portData)

}

func getResult(context *gin.Context) {
	context.String(http.StatusOK, "hello, world")
}
