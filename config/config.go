package config

import (
	"encoding/json"
	"fmt"
	"getReslut/betCount"
	"getReslut/public"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/Unknwon/goconfig"
)

var DEBUG = false
var cfg *goconfig.ConfigFile
var gameCodemap map[string]interface{}

func GetDeBugMode() bool {
	return DEBUG
}

func GetgameCodeMap() map[string]interface{} {
	gameCodeMap := make(map[string]interface{})

	gameCode := make(map[string]interface{})
	//北京PK10 //幸运飞艇 // 區塊鏈 PK10 // 1分PK10 // 3分PK10
	gameCode["ltg"] = []int{1, 3}
	gameCode["lt"] = []int{1, 5, 31, 34, 35}
	gameCodeMap["pk10"] = gameCode
	// 30 區塊鏈時時彩
	gameCode = make(map[string]interface{})
	gameCode["ltg"] = []int{2}
	gameCode["lt"] = []int{2, 3, 4, 26, 27, 28, 29, 30}
	gameCodeMap["shishi."] = gameCode
	//PC蛋蛋 //區塊鏈 PC蛋蛋 // 1分 PC蛋蛋 // 3分 PC蛋蛋
	gameCode = make(map[string]interface{})
	gameCode["ltg"] = []int{4}
	gameCode["lt"] = []int{6, 33, 40, 41}
	gameCodeMap["pcegg"] = gameCode
	//六合彩 //區塊鏈六合彩 // 1分六合彩 // 5分六合彩
	gameCode = make(map[string]interface{})
	gameCode["ltg"] = []int{5}
	gameCode["lt"] = []int{7, 32, 38, 39}
	gameCodeMap["markSixTrad"] = gameCode

	gameCode = make(map[string]interface{})
	gameCode["ltg"] = []int{6}
	gameCode["lt"] = []int{8, 9, 10, 36, 37, 42}
	gameCodeMap["kuaiThree"] = gameCode

	return gameCodeMap
}

func ConfigInit(config *goconfig.ConfigFile) {
	// fileAddr := "/"
	// if DEBUG {
	// 	fileAddr = "./"
	// } else {
	// 	fileAddr = getCurrentDirectory() + "/"
	// }
	// config, err := goconfig.LoadConfigFile(fileAddr + "configData.conf") //加载配置文件
	// if err != nil {
	// 	fmt.Println("get config file error")
	// 	os.Exit(-1)
	// }

	cfg = config
}

func GetConfig() map[string]interface{} {

	ConfigData := make(map[string]interface{})

	ip, _ := cfg.GetValue("db", "ip")
	poolname, _ := cfg.GetValue("db", "poolname")
	port, _ := cfg.GetValue("db", "port")
	user, _ := cfg.GetValue("db", "user")
	password, _ := cfg.GetValue("db", "password")
	database, _ := cfg.GetValue("db", "database")
	charset, _ := cfg.GetValue("db", "charset")

	redisIp, _ := cfg.GetValue("redis", "redisIp")
	redisPort, _ := cfg.GetValue("redis", "redisPort")
	redisPassword, _ := cfg.GetValue("redis", "redisPassword")
	redisDB, _ := cfg.GetValue("redis", "redisDB")
	redisDBInt, _ := strconv.Atoi(redisDB)

	//台北
	// ConfigData["ip"] = "192.168.88.181"
	// ConfigData["poolname"] = "LC_MEM"
	// ConfigData["port"] = "3306"
	// ConfigData["user"] = "lc_db"
	// ConfigData["password"] = "1qaz!QAZ"
	// ConfigData["database"] = "LC_MEM"
	// ConfigData["charset"] = "utf8"

	// ConfigData["redisIp"] = "192.168.88.118"
	// ConfigData["redisPort"] = "6379"
	// ConfigData["redisPassword"] = ""
	// ConfigData["redisDB"] = 3

	//???
	// ConfigData["ip"] = "192.168.88.181"
	// ConfigData["poolname"] = "LC_MEM"
	// ConfigData["port"] = "3306"
	// ConfigData["user"] = "lc_db"
	// ConfigData["password"] = "1qaz!QAZ"
	// ConfigData["database"] = "LC_MEM"
	// ConfigData["charset"] = "utf8"

	// ConfigData["redisIp"] = "192.168.88.118"
	// ConfigData["redisPort"] = "6379"
	// ConfigData["redisPassword"] = ""
	// ConfigData["redisDB"] = 3

	// 台中
	ConfigData["ip"] = ip
	ConfigData["poolname"] = poolname
	ConfigData["port"] = port
	ConfigData["user"] = user
	ConfigData["password"] = password
	ConfigData["database"] = database
	ConfigData["charset"] = charset

	ConfigData["redisIp"] = redisIp
	ConfigData["redisPort"] = redisPort
	ConfigData["redisPassword"] = redisPassword
	ConfigData["redisDB"] = redisDBInt

	return ConfigData
}
func GetRedisConfig() map[string]interface{} {
	redisConfigData := GetConfig()
	ConfigData := make(map[string]interface{})

	ConfigData["ip"] = redisConfigData["redisIp"]
	ConfigData["port"] = redisConfigData["redisPort"]
	ConfigData["password"] = redisConfigData["redisPassword"]
	ConfigData["database"] = redisConfigData["redisDB"]

	return ConfigData
}

func GetDBConfig() string {

	dbConfigData := GetConfig()

	connStr := dbConfigData["user"].(string) + ":" + dbConfigData["password"].(string) +
		"@tcp(" + dbConfigData["ip"].(string) + ":" + dbConfigData["port"].(string) + ")/" +
		dbConfigData["database"].(string) + "?charset=" + dbConfigData["charset"].(string)
	return connStr
}

func Init(lotteryTypeGroup int, lotteryType int) map[string]interface{} {
	var ConfigData map[string]interface{}
	fileAddr := "/"
	if DEBUG {
		fileAddr = "./"
	} else {
		fileAddr = getCurrentDirectory() + "/"
	}

	switch lotteryTypeGroup {
	case 1, 3:
		switch lotteryType {
		case 1, 5, 31, 34, 35: //北京PK10 //幸运飞艇 // 區塊鏈 PK10 // 1分PK10 // 3分PK10
			fileAddr = fileAddr + "config/configPk10.json"
		}
	case 2:
		switch lotteryType {
		case 2, 3, 4, 26, 27, 28, 29, 30: // 30 區塊鏈時時彩
			fileAddr = fileAddr + "config/configShiShi.json"
		default:
			fileAddr = fileAddr + "config/configShiShi.json"
		}
	case 4:
		switch lotteryType {
		case 6, 33, 40, 41: //PC蛋蛋 //區塊鏈 PC蛋蛋 // 1分 PC蛋蛋 // 3分 PC蛋蛋
			fileAddr = fileAddr + "config/configPCEgg.json"
		default:
			fileAddr = fileAddr + "config/configPCEgg.json"
		}
	case 5:
		switch lotteryType {
		case 7, 32, 38, 39: //六合彩 //區塊鏈六合彩 // 1分六合彩 // 5分六合彩
			fileAddr = fileAddr + "config/configMarkSixTrad.json"
		default:
			fileAddr = fileAddr + "config/configShiShi.json"
		}
	case 6: //快三
		switch lotteryType {
		//江苏快三 //安徽快三 //广西快三// 1分快三 //3分快三 //區域鏈快三
		case 8, 9, 10, 36, 37, 42:
			fileAddr = fileAddr + "config/configKuaiThree.json"
		}

	default:
		fileAddr = fileAddr + "config/configShiShi.json"
	}

	// jsonFile, err := os.Open(getCurrentDirectory() + fileAddr) //讀取json檔案
	jsonFile, err := os.Open(fileAddr) //讀取json檔案..

	if err != nil { //讀取檔案錯誤
		fmt.Println(" 讀取檔案錯誤 ")
		fmt.Println(err)
	} else {
		defer jsonFile.Close()                         //關閉檔案流
		byteValue, _ := ioutil.ReadAll(jsonFile)       //json string轉為byte
		json.Unmarshal([]byte(byteValue), &ConfigData) //json轉為map
	}

	return ConfigData
}

func getCurrentDirectory() string {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}

func ReplaceRealContent(LotteryTypeGroup int,
	LotteryType int,
	DrawIssue string,
	playRule map[string]interface{},
	LotteryMode int,
	LotteryPlayGroup int,
	LotteryPlay int,
	LotteryContent interface{}) map[string]interface{} {

	data := make(map[string]interface{})
	ContentType := playRule["LTR_Config"].(map[string]interface{})["LTR_ContentType"].(map[string]interface{})[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].(float64)

	// fmt.Println("  ContentType ", ContentType)
	switch ContentType {
	case 1:
		{
			//直接把字串多餘的字元拿掉
			newLotteryContent := strings.Replace(LotteryContent.(string), "[\"", "", -1)
			newLotteryContent = strings.Replace(newLotteryContent, "\"]", "", -1)
			newLotteryContent = strings.Replace(newLotteryContent, "\", \"", ",", -1)

			_data2 := strings.Split(newLotteryContent, ",")
			_data := make([]interface{}, 0)
			for i := 0; i < len(_data2); i++ {
				_data = append(_data, _data2[i])
			}

			for i := 0; i < len(_data2); i++ {
				//check "PlayRule" for real content by content cells with items

				if Content, ok := playRule["LTR_ContentToResult"].(map[string]interface{})[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].(map[string]interface{}); ok {
					if Content1, ok1 := Content[_data2[i]].(string); ok1 {
						_data[i] = Content1
					} else if Content2, ok1 := Content[_data2[i]].([]interface{}); ok1 {
						_data[i] = Content2
					} else if Content3, ok1 := Content[_data2[i]].(float64); ok1 {
						_data[i] = fmt.Sprintf("%.0f", Content3)
					}
				}
			}
			newData := make(map[string]interface{})
			if len(_data2) == 1 {
				for i := 0; i < len(_data2); i++ {
					newData[strconv.Itoa(i)] = _data[i]
				}
			} else {
				newData["1"] = _data
			}

			data = newData
		}
	case 2:
		{
			jsonStr := LotteryContent.(string)
			newLotteryContent := make(map[string]interface{})
			err := json.Unmarshal([]byte(jsonStr), &newLotteryContent)

			if err != nil {
				panic(err)
			}
			data = newLotteryContent["1"].(map[string]interface{})

			for key, value := range data {
				valueArr := value.([]interface{})

				for j := 0; j < len(valueArr); j++ {

					if Content, ok := playRule["LTR_ContentToResult"].(map[string]interface{})[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]interface{}); ok {
						item, err := strconv.Atoi(valueArr[j].(string))
						if err != nil {
							panic(err)
						}
						if realContent, ok := Content[item].(string); ok {
							data[key].([]interface{})[j] = realContent
						}
					} else if Content2, ok := playRule["LTR_ContentToResult"].(map[string]interface{})[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].(map[string]interface{}); ok {
						item := valueArr[j].(string)
						if realContent, ok := Content2[item].(string); ok {
							data[key].([]interface{})[j] = realContent
						} else if realContent1, ok := Content2[item].([]interface{}); ok {
							data[key].([]interface{})[j] = realContent1
						}
					}
				}
			}
		}
	case 3:
		{
			jsonStr := LotteryContent.(string)
			newLotteryContent := make(map[string]interface{})
			err := json.Unmarshal([]byte(jsonStr), &newLotteryContent)

			if err != nil {
				panic(err)
			}

			data = newLotteryContent["1"].(map[string]interface{})
		}
	case 4:
		{
			jsonStr := LotteryContent.(string)
			newLotteryContent := make(map[string]interface{})
			err := json.Unmarshal([]byte(jsonStr), &newLotteryContent)

			if err != nil {
				panic(err)
			}
			data = newLotteryContent["1"].(map[string]interface{})
			rule := playRule["LTR_ContentToResult"].(map[string]interface{})[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].(map[string]interface{})

			for key, value := range rule {
				data["text"] = strings.Replace(data["text"].(string), key, value.(string), -1)
			}

		}
	default:
		data = LotteryContent.(map[string]interface{})
		break
	}

	return data
}
func ReplaceRealResult(lotteryTypeGroup int,
	lotteryType int,
	lotteryResult1 []string,
	lotteryResult2 interface{},
	fullResult interface{},
	emptyResult map[string]interface{},
	config map[string]interface{}) map[string]interface{} {

	result := make(map[string]interface{})
	result["Status"] = true
	result["Info"] = ""

	result["FullResult"] = make(map[string]interface{})

	switch lotteryTypeGroup {
	case 1, 3:
		switch lotteryType {
		case 1, 5, 31, 34, 35: //北京PK10 //幸运飞艇 // 區塊鏈 PK10 // 1分PK10 // 3分PK10
			result["FullResult"] = ReplaceModulePKTen(lotteryTypeGroup, lotteryType, lotteryResult1, lotteryResult2, emptyResult, config)
		}
	case 2:
		{
			//彩種 (因為有對應 TS_LotteryType.LT_Code，所以需要跟資料庫一起變動...ry)
			switch lotteryType {
			//重庆时时彩 //新疆时时彩 //天津时时彩 //精彩1分彩 //精彩3分彩 //精彩5分彩 //精彩秒秒彩
			case 2, 3, 4, 26, 27, 28, 29, 30: // 30 區塊鏈時時彩
				result["FullResult"] = ReplaceModuleShiShi(lotteryTypeGroup, lotteryType, lotteryResult1, lotteryResult2, emptyResult, config)
			//彩種錯誤
			default:
				result["Status"] = false
				break
			}
			break
		}
	case 4:
		switch lotteryType {
		case 6, 33, 40, 41: //PC蛋蛋 //區塊鏈 PC蛋蛋 // 1分 PC蛋蛋 // 3分 PC蛋蛋
			result["FullResult"] = ReplaceModulePCEgg(lotteryTypeGroup, lotteryType, lotteryResult1, lotteryResult2, emptyResult, config)
		//彩種錯誤
		default:
			result["Status"] = false
			break
		}
		break
	case 5:
		switch lotteryType {
		case 7, 32, 38, 39: //六合彩 //區塊鏈六合彩 // 1分六合彩 // 5分六合彩
			lotteryResult2t := lotteryResult2.(string)
			result["FullResult"] = ReplaceModuleMarkSix(lotteryTypeGroup, lotteryType, lotteryResult1, lotteryResult2t, emptyResult, config)
		//彩種錯誤
		default:
			result["Status"] = false
			break
		}
		break
	case 6:

		//彩種 (因為有對應 TS_LotteryType.LT_Code，所以需要跟資料庫一起變動...ry)
		switch lotteryType {
		//江苏快三 //安徽快三 //广西快三// 1分快三 //3分快三 ////區域鏈快三
		case 8, 9, 10, 36, 37, 42:
			result["FullResult"] = ReplaceModuleKuaiThree(lotteryTypeGroup, lotteryType, lotteryResult1, lotteryResult2, emptyResult, config)
		//彩種錯誤
		default:
			result["Status"] = false
			break
		}
		break
	default:
		result["Status"] = false
		break //彩種群組錯誤
	}

	if result["Status"] == true {

	} else {
		result["Info"] = "BetCollect Gateway Error"
	}

	return result["FullResult"].(map[string]interface{})
}

func ReplaceModuleMarkSix(lotteryTypeGroup int,
	lotteryType int,
	OpenResult []string,
	OpenResult2 string,
	Result1 map[string]interface{},
	config map[string]interface{}) map[string]interface{} {

	Result := Result1

	OpenResultInt := []int{}
	OpenResultFloat64 := []float64{}
	for _, num := range OpenResult {
		numInt, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		OpenResultInt = append(OpenResultInt, numInt)
		OpenResultFloat64 = append(OpenResultFloat64, float64(numInt))
	}
	//var OpenResult2Int int
	OpenResult2Int, err := strconv.Atoi(OpenResult2)
	if err != nil {
		panic(err)
	}
	OpenResult2Float64 := float64(OpenResult2Int)
	// for _, num := range OpenResult2 {
	// 	numInt, err := strconv.Atoi(num)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	OpenResult2Int = append(OpenResult2Int, numInt)
	// }
	//初始化
	// State := true
	LotteryPlayMode := 0
	LotteryPlayGroup := 0
	LotteryPlay := 0

	//取得大小的判斷基準 (LTG: 1, 2, 4, 7, 10)
	// BigOrSmallBase := int(config["LTR_BigOrSmallBase"].(float64))
	// 取得龙虎的判斷基準 (LTG: 1, 3, 6, 10)
	// DragonOrTigerBase := int(config["LTR_DragonOrTigerBase"].(float64))
	//取得总和計算的元素個數 (LTG: 1, 3, 6, 7)
	// TotalBase := int(config["LTR_TotalBase"].(float64))
	//取得总和大小的判斷基準 (LTG: 1, 2, 3, 5, 6, 7, 10)
	TotalBigOrSmallBase := int(config["LTR_TotalBigOrSmallBase"].(float64))
	//取得極大的判斷基準 (LTG: 4) 尾大尾小的的判斷基準 (LTG: 10)
	ExtremeBigBase := int(config["LTR_ExtremeBigBase"].(float64))
	//取得極小的判斷基準 (LTG: 4) 尾大尾小的的判斷基準 (LTG: 10)
	// ExtremeSmallBase := config["LTR_ExtremeSmallBase"]
	//取得上下盘的判斷基準 (LTG: 8)
	// TopOrUnderBase := config["LTR_TopOrUnderBase"]
	//取得特碼大小的判斷基準 (LTG: 5)
	UniqueBigOrSmallBase := int(config["LTR_UniqueBigOrSmallBase"].(float64))
	//取得全部位數和大小的判斷基準 (LTG: 5)
	DigitBigOrSmallBase := int(config["LTR_DigitBigOrSmallBase"].(float64))
	//取得三個位數和大小的判斷基準 (LTG: 2)
	// TriplexBigOrSmallBase := config["LTR_TriplexBigOrSmallBase"]
	//取得和局的判斷基準 (LTG: 5,7)
	DrawBase := int(config["LTR_DrawBase"].(float64))
	//取得總和和局的判斷基準 (LTG: 7,8,10)
	// TotalDrawBase := config["LTR_TotalDrawBase"]
	//取得波色的分類資料 (LTG: 5)
	Wave := config["LTR_Wave"].(map[string]interface{})
	// 取得生肖的分類資料 (LTG: 5)
	Zodiac := config["LTR_Zodiac"].(map[string]interface{})
	// 取得五行的分類資料 (LTG: 5)
	Elements := config["LTR_FiveElements"].(map[string]interface{})

	switch lotteryTypeGroup {
	//六合彩
	case 5:
		{
			//彩種 (因為有對應 TS_LotteryType.LT_Code，所以需要跟資料庫一起變動...ry)
			switch lotteryType {
			case 7, 32, 38, 39: //六合彩 //區塊鏈六合彩 // 1分六合彩 // 5分六合彩
				LotteryPlayMode = 1 //信用模式 (传统)

				// /***********************************************************************
				//  *	開獎結果預先處理 (头数|尾数|头数尾数和|色波|生肖|五行) - [START] */

				OpenResultTens := make([]interface{}, len(OpenResult)) //正码头数
				var OpenResult2Tens []interface{}                      //特码头数

				OpenResultUnits := make([]interface{}, len(OpenResult)) //正码尾数
				var OpenResult2Units []interface{}                      //特码尾数
				var OpenResultAllUnits []interface{}                    //全码尾数

				OpenResultDigitsTotal := make([]interface{}, len(OpenResult)) //正码头数尾数和
				var OpenResult2DigitsTotal []interface{}                      //特码头数尾数和

				OpenResultWave := make([]interface{}, len(OpenResult)) //正码色波
				var OpenResult2Wave []interface{}                      //特码色波
				var OpenResultAllWave []interface{}                    //全码色波

				OpenResultZodiac := make([]int, len(OpenResult)) //正码生肖
				var OpenResult2Zodiac []int                      //特码生肖
				var OpenResultAllZodiac []int                    //全码生肖

				OpenResultElement := make([]interface{}, len(OpenResult)) //正码五行
				var OpenResult2Element []interface{}                      //特码五行
				var OpenResultAllElement []interface{}                    //全码五行

				//取得生肖與五行的判斷基準(取得立春的陽曆日期來做跨年度的判斷依據)
				t1 := time.Now()
				Year := t1.Year()
				//Month := t.Month()
				if public.GetDateForSpringBeginning() {
					Zodiac = config["LTR_Zodiac"].(map[string]interface{})[strconv.Itoa((Year-1)%12)].(map[string]interface{})
					Elements = config["LTR_FiveElements"].(map[string]interface{})[strconv.Itoa((Year-1)%30)].(map[string]interface{})
				} else {
					Zodiac = config["LTR_Zodiac"].(map[string]interface{})[strconv.Itoa((Year)%12)].(map[string]interface{})
					Elements = config["LTR_FiveElements"].(map[string]interface{})[strconv.Itoa((Year)%30)].(map[string]interface{})
				}

				for i := 0; i < len(OpenResult); i++ {
					tempResult := fmt.Sprintf("%02.0f", float64(OpenResultInt[i]))
					tempResultInt1, _ := strconv.Atoi(strings.Split(tempResult, "")[0])
					tempResultInt2, _ := strconv.Atoi(strings.Split(tempResult, "")[1])
					//正码头数

					OpenResultTens[i] = tempResultInt2
					//正码尾数

					OpenResultUnits[i] = tempResultInt1
					//正码头数尾数和

					OpenResultDigitsTotal[i] = tempResultInt1 + tempResultInt2

					//正码色波

					if betCount.In_array(OpenResultFloat64[i], Wave["Red"]) {
						OpenResultWave[i] = "RW"
					} else if betCount.In_array(OpenResultFloat64[i], Wave["Cyan"]) {
						OpenResultWave[i] = "CW"
					} else if betCount.In_array(OpenResultFloat64[i], Wave["Green"]) {
						OpenResultWave[i] = "GW"
					}

					//正码生肖
					if betCount.In_array(OpenResultFloat64[i], Zodiac["Rat"]) {
						OpenResultZodiac[i] = 1
					} else if betCount.In_array(OpenResultFloat64[i], Zodiac["Ox"]) {
						OpenResultZodiac[i] = 2
					} else if betCount.In_array(OpenResultFloat64[i], Zodiac["Tiger"]) {
						OpenResultZodiac[i] = 3
					} else if betCount.In_array(OpenResultFloat64[i], Zodiac["Rabbit"]) {
						OpenResultZodiac[i] = 4
					} else if betCount.In_array(OpenResultFloat64[i], Zodiac["Dragon"]) {
						OpenResultZodiac[i] = 5
					} else if betCount.In_array(OpenResultFloat64[i], Zodiac["Snake"]) {
						OpenResultZodiac[i] = 6
					} else if betCount.In_array(OpenResultFloat64[i], Zodiac["Horse"]) {
						OpenResultZodiac[i] = 7
					} else if betCount.In_array(OpenResultFloat64[i], Zodiac["Goat"]) {
						OpenResultZodiac[i] = 8
					} else if betCount.In_array(OpenResultFloat64[i], Zodiac["Monkey"]) {
						OpenResultZodiac[i] = 9
					} else if betCount.In_array(OpenResultFloat64[i], Zodiac["Rooster"]) {
						OpenResultZodiac[i] = 10
					} else if betCount.In_array(OpenResultFloat64[i], Zodiac["Dog"]) {
						OpenResultZodiac[i] = 11
					} else if betCount.In_array(OpenResultFloat64[i], Zodiac["Pig"]) {
						OpenResultZodiac[i] = 12
					}

					//正码五行
					if betCount.In_array(OpenResultFloat64[i], Elements["Metal"]) {
						OpenResultElement[i] = "5M"
					} else if betCount.In_array(OpenResultFloat64[i], Elements["Tree"]) {
						OpenResultElement[i] = "5T"
					} else if betCount.In_array(OpenResultFloat64[i], Elements["Water"]) {
						OpenResultElement[i] = "5W"
					} else if betCount.In_array(OpenResultFloat64[i], Elements["Fire"]) {
						OpenResultElement[i] = "5F"
					} else if betCount.In_array(OpenResultFloat64[i], Elements["Earth"]) {
						OpenResultElement[i] = "5E"
					}

				}

				tempResult2 := fmt.Sprintf("%02.0f", OpenResult2Float64)
				tempResult2Int1, _ := strconv.Atoi(strings.Split(tempResult2, "")[0])
				tempResult2Int2, _ := strconv.Atoi(strings.Split(tempResult2, "")[1])

				//tempResult2 := strings.Split(fmt.Sprintf("%02d", OpenResult2), "")
				// Result2Tens := strings.Split(tempResult2, "")[0]

				//特码头数

				OpenResult2Tens = append(OpenResult2Tens, tempResult2Int1)
				// OpenResult2Tens[0] = strings.Split(tempResult2, "")[0]
				//特码尾数
				//Result2Units, _ := strconv.Atoi(fmt.Sprintf("%02d", OpenResult2)[1:1])
				OpenResult2Units = append(OpenResult2Tens, tempResult2Int2)
				// OpenResult2Units[0] = strings.Split(tempResult2, "")[1]
				//特码头数尾数和

				// DigitsTotal1, _ := strconv.Atoi(fmt.Sprintf("%02d", OpenResult2)[0:1])
				// DigitsTotal2, _ := strconv.Atoi(fmt.Sprintf("%02d", OpenResult2)[1:1])
				// OpenResult2DigitsTotal[0] = DigitsTotal1 + DigitsTotal2
				OpenResult2DigitsTotal = append(OpenResult2DigitsTotal, tempResult2Int1+tempResult2Int2)

				//特码色波
				if betCount.In_array(OpenResult2Float64, Wave["Red"]) {
					OpenResult2Wave = append(OpenResult2Wave, "RW")
				} else if betCount.In_array(OpenResult2Float64, Wave["Cyan"]) {
					OpenResult2Wave = append(OpenResult2Wave, "CW")
				} else if betCount.In_array(OpenResult2Float64, Wave["Green"]) {
					OpenResult2Wave = append(OpenResult2Wave, "GW")
				}

				//特码生肖
				if betCount.In_array(OpenResult2Float64, Zodiac["Rat"]) {
					OpenResult2Zodiac = append(OpenResult2Zodiac, 1)
				} else if betCount.In_array(OpenResult2Float64, Zodiac["Ox"]) {
					OpenResult2Zodiac = append(OpenResult2Zodiac, 2)
				} else if betCount.In_array(OpenResult2Float64, Zodiac["Tiger"]) {
					OpenResult2Zodiac = append(OpenResult2Zodiac, 3)
				} else if betCount.In_array(OpenResult2Float64, Zodiac["Rabbit"]) {
					OpenResult2Zodiac = append(OpenResult2Zodiac, 4)
				} else if betCount.In_array(OpenResult2Float64, Zodiac["Dragon"]) {
					OpenResult2Zodiac = append(OpenResult2Zodiac, 5)
				} else if betCount.In_array(OpenResult2Float64, Zodiac["Snake"]) {
					OpenResult2Zodiac = append(OpenResult2Zodiac, 6)
				} else if betCount.In_array(OpenResult2Float64, Zodiac["Horse"]) {
					OpenResult2Zodiac = append(OpenResult2Zodiac, 7)
				} else if betCount.In_array(OpenResult2Float64, Zodiac["Goat"]) {
					OpenResult2Zodiac = append(OpenResult2Zodiac, 8)
				} else if betCount.In_array(OpenResult2Float64, Zodiac["Monkey"]) {
					OpenResult2Zodiac = append(OpenResult2Zodiac, 9)
				} else if betCount.In_array(OpenResult2Float64, Zodiac["Rooster"]) {
					OpenResult2Zodiac = append(OpenResult2Zodiac, 10)
				} else if betCount.In_array(OpenResult2Float64, Zodiac["Dog"]) {
					OpenResult2Zodiac = append(OpenResult2Zodiac, 11)
				} else if betCount.In_array(OpenResult2Float64, Zodiac["Pig"]) {
					OpenResult2Zodiac = append(OpenResult2Zodiac, 12)
				}

				//特码五行
				if betCount.In_array(OpenResult2Float64, Elements["Metal"]) {
					OpenResult2Element = append(OpenResult2Element, "5M")
				} else if betCount.In_array(OpenResult2Float64, Elements["Tree"]) {
					OpenResult2Element = append(OpenResult2Element, "5T")
				} else if betCount.In_array(OpenResult2Float64, Elements["Water"]) {
					OpenResult2Element = append(OpenResult2Element, "5W")
				} else if betCount.In_array(OpenResult2Float64, Elements["Fire"]) {
					OpenResult2Element = append(OpenResult2Element, "5F")
				} else if betCount.In_array(OpenResult2Float64, Elements["Earth"]) {
					OpenResult2Element = append(OpenResult2Element, "5E")
				}
				//全码尾数
				OpenResultAllUnits = OpenResultUnits
				OpenResultAllUnits = append(OpenResultAllUnits, OpenResult2Units[0])
				//全码色波
				OpenResultAllWave = OpenResultWave
				OpenResultAllWave = append(OpenResultAllWave, OpenResult2Wave[0])
				//全码生肖
				OpenResultAllZodiac = OpenResultZodiac
				OpenResultAllZodiac = append(OpenResultAllZodiac, OpenResult2Zodiac[0])
				//全码五行
				OpenResultAllElement = OpenResultElement
				OpenResultAllElement = append(OpenResultAllElement, OpenResult2Element[0])
				/*	開獎結果預先處理 (头数|尾数|头数尾数和|色波|生肖|五行) - [END]
				***********************************************************************/

				LotteryPlayGroup = 1 //玩法群組 (特码)
				LotteryPlay = 1      //玩法 (特码)

				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResult2

				//玩法 (两面)
				LotteryPlay = 2
				s := ""
				t := make([]interface{}, 0)
				if OpenResult2Int != DrawBase {
					//大 or 小 (B|S)
					if OpenResult2Int > UniqueBigOrSmallBase {
						s = "B"
					} else {
						s = "S"
					}
					t = append([]interface{}{s}, t...)
					//单 or 双 (O|E)
					if OpenResult2Int%2 > 0 {
						s = "O"
					} else {
						s = "E"
					}
					t = append([]interface{}{s}, t...)
					//Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = t
				} else {
					s = "N"
					t = append([]interface{}{s}, t...)
					//Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = t
					// array_push($MR_Result->{$MR_LotteryPlayMode}->{$MR_LotteryPlayGroup}->{$MR_LotteryPlay},"N");

				}
				t = append([]interface{}{OpenResult2Wave[0]}, t...)
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = t
				//红波 or 绿波 or 蓝波 (RW|GW|CW)

				//玩法群組 (两面)
				LotteryPlayGroup = 2
				//玩法 (两面)
				LotteryPlay = 1
				s = ""
				t = make([]interface{}, 0)
				if OpenResult2Int != DrawBase {
					//特大 or 特小 (B|S)
					if OpenResult2Int > UniqueBigOrSmallBase {
						s = "B"
					} else {
						s = "S"
					}
					t = append([]interface{}{s}, t...)
					//特单 or 特双 (O|E)
					if OpenResult2Int%2 > 0 {
						s = "O"
					} else {
						s = "E"
					}
					t = append([]interface{}{s}, t...)
					//特尾大 or 特尾小 (FB|FS)
					if OpenResult2Units[0].(int) > ExtremeBigBase {
						s = "FB"
					} else {
						s = "FS"
					}
					t = append([]interface{}{s}, t...)

					//特大单 or 特大双 (LO|LE)
					if OpenResult2Int > UniqueBigOrSmallBase {
						if OpenResult2Int%2 > 0 {
							s = "LO"
						} else {
							s = "LE"
						}
						t = append([]interface{}{s}, t...)
					}

					//特小单 or 特小双 (MO|ME)
					if OpenResult2Int < UniqueBigOrSmallBase {
						if OpenResult2Int%2 > 0 {
							s = "MO"
						} else {
							s = "ME"
						}
						t = append([]interface{}{s}, t...)
					}
					//特合大 or 特合小 (PB|PS)
					if OpenResult2DigitsTotal[0].(int) > DigitBigOrSmallBase {
						s = "PB"
					} else {
						s = "PS"
					}
					t = append([]interface{}{s}, t...)

					//特合单 or 特合双 (PO|PE)
					if OpenResult2DigitsTotal[0].(int)%2 > 0 {
						s = "PO"
					} else {
						s = "PE"
					}
					t = append([]interface{}{s}, t...)

					if betCount.In_array(OpenResult2Zodiac[0], []int{2, 4, 5, 7, 9, 12}) {
						t = append([]interface{}{"TZ"}, t...)
					} else if betCount.In_array(OpenResult2Zodiac[0], []int{1, 3, 6, 8, 10, 11}) {
						t = append([]interface{}{"UZ"}, t...)
					}

					if betCount.In_array(OpenResult2Zodiac[0], []int{1, 2, 3, 4, 5, 6}) {
						t = append([]interface{}{"AZ"}, t...)
					} else if betCount.In_array(OpenResult2Zodiac[0], []int{7, 8, 9, 10, 11, 12}) {
						t = append([]interface{}{"FZ"}, t...)
					}

					if betCount.In_array(OpenResult2Zodiac[0], []int{2, 7, 8, 10, 11, 12}) {
						t = append([]interface{}{"HZ"}, t...)
					} else if betCount.In_array(OpenResult2Zodiac[0], []int{1, 3, 4, 5, 6, 9}) {
						t = append([]interface{}{"WZ"}, t...)
					}
				} else {
					s = "N"
					t = append([]interface{}{s}, t...)
				}

				OpenResultTotal := 0
				for i := 0; i < len(OpenResultInt); i++ {
					OpenResultTotal = OpenResultTotal + OpenResultInt[i]
				}
				OpenResultTotal = OpenResultTotal + OpenResult2Int

				//总和单 or 总和双 (TO|TE)
				if OpenResultTotal%2 > 0 {
					s = "TO"
				} else {
					s = "TE"
				}
				t = append([]interface{}{s}, t...)

				//总和大 or 总和小 (TB|TS)
				if OpenResultTotal > TotalBigOrSmallBase {
					s = "TB"
				} else {
					s = "TS"
				}
				t = append([]interface{}{s}, t...)
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = t

				//玩法群組 (正码)
				LotteryPlayGroup = 3
				//玩法 (正码)
				LotteryPlay = 1

				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResult

				//玩法 (两面)
				LotteryPlay = 2
				s = ""
				t = make([]interface{}, 0)
				//总和单 or 总和双 (TO|TE)
				if OpenResultTotal%2 > 0 {
					s = "TO"
				} else {
					s = "TE"
				}
				t = append([]interface{}{s}, t...)

				//总和大 or 总和小 (TB|TS)
				if OpenResultTotal > TotalBigOrSmallBase {
					s = "TB"
				} else {
					s = "TS"
				}
				t = append([]interface{}{s}, t...)
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = t

				//玩法群組 (特码1-6)
				LotteryPlayGroup = 4
				OpenResultIndex := 0
				s = ""

				//玩法 (正码一|正码二|正码三|正码四|正码五|正码六)

				for LotteryPlay := 1; LotteryPlay <= 6; LotteryPlay++ {
					t = make([]interface{}, 0)
					OpenResultIndex = LotteryPlay - 1
					if OpenResultInt[OpenResultIndex] != DrawBase {
						//大 or 小 (B|S)
						if OpenResultInt[OpenResultIndex] > UniqueBigOrSmallBase {
							s = "B"
						} else {
							s = "S"
						}
						t = append([]interface{}{s}, t...)

						//单 or 双 (O|E)
						if OpenResultInt[OpenResultIndex]%2 > 0 {
							s = "O"
						} else {
							s = "E"
						}
						t = append([]interface{}{s}, t...)

						//尾大 or 尾小 (FB|FS)

						if OpenResultUnits[OpenResultIndex].(int) > ExtremeBigBase {
							s = "FB"
						} else {
							s = "FS"
						}
						t = append([]interface{}{s}, t...)

						//合大 or 合小 (PB|PS)
						if OpenResultDigitsTotal[OpenResultIndex].(int) > DigitBigOrSmallBase {
							s = "PB"
						} else {
							s = "PS"
						}
						t = append([]interface{}{s}, t...)
						//合单 or 合双 (PO|PE)
						if OpenResultDigitsTotal[OpenResultIndex].(int)%2 > 0 {
							s = "PO"
						} else {
							s = "PE"
						}
						t = append([]interface{}{s}, t...)
					} else {
						t = append([]interface{}{"N"}, t...)
					}
					t = append([]interface{}{OpenResultWave[OpenResultIndex]}, t...)

					Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = t
				}

				//玩法群組 (正特)
				LotteryPlayGroup = 5
				//玩法 (正1特|正2特|正3特|正4特|正5特|正6特)
				for LotteryPlay = 1; LotteryPlay <= 12; LotteryPlay = LotteryPlay + 2 {
					OpenResultIndex = (LotteryPlay - 1) / 2
					Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResult[OpenResultIndex]
				}
				//玩法 (正1特双面|正2特双面|正3特双面|正4特双面|正5特双面|正6特双面)

				for LotteryPlay = 2; LotteryPlay <= 12; LotteryPlay = LotteryPlay + 2 {
					s = ""
					t = make([]interface{}, 0)
					OpenResultIndex = (LotteryPlay - 2) / 2
					if OpenResultInt[OpenResultIndex] != DrawBase {
						//大 or 小 (B|S)
						if OpenResultInt[OpenResultIndex] > UniqueBigOrSmallBase {
							s = "B"
						} else {
							s = "S"
						}
						t = append([]interface{}{s}, t...)

						//单 or 双 (O|E)
						if OpenResultInt[OpenResultIndex]%2 > 0 {
							s = "O"
						} else {
							s = "E"
						}
						t = append([]interface{}{s}, t...)

						//尾大 or 尾小 (FB|FS)
						if OpenResultUnits[OpenResultIndex].(int) > ExtremeBigBase {
							s = "FB"
						} else {
							s = "FS"
						}
						t = append([]interface{}{s}, t...)

						//合大 or 合小 (PB|PS)
						if OpenResultDigitsTotal[OpenResultIndex].(int) > DigitBigOrSmallBase {
							s = "PB"
						} else {
							s = "PS"
						}
						t = append([]interface{}{s}, t...)
						//合单 or 合双 (PO|PE)
						if OpenResultDigitsTotal[OpenResultIndex].(int)%2 > 0 {
							s = "PO"
						} else {
							s = "PE"
						}
						t = append([]interface{}{s}, t...)
					} else {
						t = append([]interface{}{"N"}, t...)
					}
					t = append([]interface{}{OpenResultWave[OpenResultIndex]}, t...)
					Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = t
				}

				//玩法群組 (连码)
				LotteryPlayGroup = 6
				//玩法 (三全中)
				LotteryPlay = 1
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResult

				//玩法 (三中二)
				LotteryPlay = 2
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResult

				//玩法 (二全中)
				LotteryPlay = 3
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResult

				//玩法 (二中特)
				LotteryPlay = 4
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = []interface{}{OpenResult, OpenResult2}

				//玩法 (特串)
				LotteryPlay = 5
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = []interface{}{OpenResult, OpenResult2}

				//玩法 (四全中)
				LotteryPlay = 6
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResult

				//玩法群組 (色波)
				LotteryPlayGroup = 7
				//玩法 (特码色波)
				LotteryPlay = 1
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResult2Wave[0]

				//玩法 (7色波)
				LotteryPlay = 2
				OpenResultWaveCount := map[string]float64{
					"RW": 0,
					"CW": 0,
					"GW": 0,
				}
				OpenResult2WaveCount := map[string]float64{
					"RW": 0.0,
					"CW": 0.0,
					"GW": 0.0,
				}
				OpenResultAllWaveCount := ""

				for i := 0; i < len(OpenResultWave); i++ {
					wave := OpenResultWave[i].(string)
					if wave == "RW" {
						OpenResultWaveCount["RW"]++
					} else if wave == "CW" {
						OpenResultWaveCount["CW"]++
					} else if wave == "GW" {
						OpenResultWaveCount["GW"]++
					}

				}

				wave2 := OpenResult2Wave[0].(string)
				if wave2 == "RW" {
					OpenResult2WaveCount["RW"] = OpenResult2WaveCount["RW"] + 1.5
				} else if wave2 == "CW" {
					OpenResult2WaveCount["CW"] = OpenResult2WaveCount["CW"] + 1.5
				} else if wave2 == "GW" {
					OpenResult2WaveCount["GW"] = OpenResult2WaveCount["GW"] + 1.5
				}

				if int(OpenResultWaveCount["RW"]) == 3 && int(OpenResultWaveCount["CW"]) == 3 && OpenResult2WaveCount["GW"] == 1.5 ||
					int(OpenResultWaveCount["RW"]) == 3 && int(OpenResultWaveCount["GW"]) == 3 && OpenResult2WaveCount["CW"] == 1.5 ||
					int(OpenResultWaveCount["CW"]) == 3 && int(OpenResultWaveCount["GW"]) == 3 && OpenResult2WaveCount["RW"] == 1.5 {
					OpenResultAllWaveCount = "N"
				} else {
					OpenResultWaveCount["RW"] = OpenResultWaveCount["RW"] + OpenResult2WaveCount["RW"]
					OpenResultWaveCount["CW"] = OpenResultWaveCount["CW"] + OpenResult2WaveCount["CW"]
					OpenResultWaveCount["GW"] = OpenResultWaveCount["GW"] + OpenResult2WaveCount["GW"]

					keys := make([]string, 0, len(OpenResultWaveCount))
					for key := range OpenResultWaveCount {
						keys = append(keys, key)
					}
					sort.Slice(keys, func(i, j int) bool { return OpenResultWaveCount[keys[i]] > OpenResultWaveCount[keys[j]] })
					OpenResultAllWaveCount = keys[0]
				}

				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResultAllWaveCount
				//玩法 (红球)
				LotteryPlay = 3
				s = ""
				t = make([]interface{}, 0)
				if OpenResult2Int != DrawBase {
					if betCount.In_array(OpenResult2Float64, Wave["Red"]) {
						//单 or 双 (ROoRE)
						if OpenResultInt[0]%2 > 0 {
							s = "RO"
						} else {
							s = "RE"
						}
						t = append([]interface{}{s}, t...)
						//大 or 小 (RBoRS)
						if OpenResultInt[0] > UniqueBigOrSmallBase {
							s = "RB"
						} else {
							s = "RS"
						}
						t = append([]interface{}{s}, t...)

						//和单 or 和双 (RTOoRTE)
						if OpenResult2DigitsTotal[0].(int)%2 > 0 {
							s = "RTO"
						} else {
							s = "RTE"
						}
						t = append([]interface{}{s}, t...)
						if OpenResultInt[0]%2 > 0 {
							//大单 or 小单 (RLO|RMO)
							if OpenResultInt[0] > UniqueBigOrSmallBase {
								s = "RLO"
							} else {
								s = "RMO"
							}
							t = append([]interface{}{s}, t...)
						} else {
							//大双 or 小双 (RLE|RME)
							if OpenResultInt[0] > UniqueBigOrSmallBase {
								s = "RLE"
							} else {
								s = "RME"
							}
							t = append([]interface{}{s}, t...)
						}

					} else if betCount.In_array(OpenResult2Float64, Wave["Cyan"]) {
						//单 or 双 (COoCE)
						if OpenResultInt[0]%2 > 0 {
							s = "CO"
						} else {
							s = "CE"
						}
						t = append([]interface{}{s}, t...)
						//大 or 小 (CBoCS)
						if OpenResultInt[0] > UniqueBigOrSmallBase {
							s = "CB"
						} else {
							s = "CS"
						}
						t = append([]interface{}{s}, t...)

						//和单 or 和双 (CTOoCTE)
						if OpenResult2DigitsTotal[0].(int)%2 > 0 {
							s = "CTO"
						} else {
							s = "CTE"
						}
						t = append([]interface{}{s}, t...)
						if OpenResultInt[0]%2 > 0 {
							//大单 or 小单 (CLO|RMO)
							if OpenResultInt[0] > UniqueBigOrSmallBase {
								s = "CLO"
							} else {
								s = "CMO"
							}
							t = append([]interface{}{s}, t...)
						} else {
							//大双 or 小双 (RLE|RME)
							if OpenResultInt[0] > UniqueBigOrSmallBase {
								s = "CLE"
							} else {
								s = "CME"
							}
							t = append([]interface{}{s}, t...)
						}

					} else if betCount.In_array(OpenResult2Float64, Wave["Green"]) {
						//单 or 双 (GOoGE)
						if OpenResultInt[0]%2 > 0 {
							s = "GO"
						} else {
							s = "GE"
						}
						t = append([]interface{}{s}, t...)
						//大 or 小 (GBoGS)
						if OpenResultInt[0] > UniqueBigOrSmallBase {
							s = "GB"
						} else {
							s = "GS"
						}
						t = append([]interface{}{s}, t...)

						//和单 or 和双 (GTOoGTE)
						if OpenResult2DigitsTotal[0].(int)%2 > 0 {
							s = "GTO"
						} else {
							s = "GTE"
						}
						t = append([]interface{}{s}, t...)
						if OpenResultInt[0]%2 > 0 {
							//大单 or 小单 (GLO|GMO)
							if OpenResultInt[0] > UniqueBigOrSmallBase {
								s = "GLO"
							} else {
								s = "GMO"
							}
							t = append([]interface{}{s}, t...)
						} else {
							//大双 or 小双 (GLE|GME)
							if OpenResultInt[0] > UniqueBigOrSmallBase {
								s = "GLE"
							} else {
								s = "GME"
							}
							t = append([]interface{}{s}, t...)
						}
					}

				} else {
					t = append([]interface{}{"N"}, t...)
				}
				//t = append([]interface{}{OpenResultWave[OpenResultIndex]}, t...)
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = t
				//玩法 (蓝球)
				LotteryPlay = 4
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay-1)]
				//玩法 (绿球)
				LotteryPlay = 5
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay-1)]

				//玩法群組 (特码头尾数)
				LotteryPlayGroup = 8
				LotteryPlay = 1
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResult2Tens[0].(int)
				LotteryPlay = 2
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResult2Units[1].(int)

				//玩法群組 (总肖)
				LotteryPlayGroup = 9

				LotteryPlay = 1
				t = make([]interface{}, 0)
				//总和计算
				OpenResultAllZodiacCount := len(betCount.Array_unique(OpenResultAllZodiac))
				t = append([]interface{}{strconv.Itoa(OpenResultAllZodiacCount)}, t...)

				//单 or 双 (O|E)
				if OpenResultAllZodiacCount%2 > 0 {
					t = append([]interface{}{"O"}, t...)

				} else {
					t = append([]interface{}{"E"}, t...)
				}
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = t

				//玩法群組 (平特一肖尾数)
				LotteryPlayGroup = 10

				//玩法 (一肖)
				LotteryPlay = 1
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResultAllZodiac
				//玩法 (尾数)
				LotteryPlay = 2
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResultUnits

				//玩法群組 (特肖)
				LotteryPlayGroup = 11
				//玩法 (特肖)
				LotteryPlay = 1
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResult2Zodiac[0]

				//玩法群組 (连肖)
				LotteryPlayGroup = 12
				//玩法 (二连肖)
				LotteryPlay = 1
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResultAllZodiac
				//玩法 (三连肖)
				LotteryPlay = 2
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResultAllZodiac
				//玩法 (四连肖)
				LotteryPlay = 3
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResultAllZodiac
				//玩法 (五连肖)
				LotteryPlay = 4
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResultAllZodiac

				//玩法群組 (合肖)
				LotteryPlayGroup = 13
				//玩法 (合肖)
				LotteryPlay = 1

				if OpenResult2Int != DrawBase {
					Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResult2Zodiac[0]
				} else {
					t = make([]interface{}, 0)
					t = append([]interface{}{"O"}, t...)
					Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = t
				}

				//玩法群組 (连尾)
				LotteryPlayGroup = 14

				//玩法 (二连尾)
				LotteryPlay = 1
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResultAllUnits
				//玩法 (三连尾)
				LotteryPlay = 2
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResultAllUnits
				//玩法 (四连尾)
				LotteryPlay = 3
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResultAllUnits
				//玩法 (五连尾)
				LotteryPlay = 4
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResultAllUnits

				//玩法群組 (正肖)
				LotteryPlayGroup = 15
				//玩法 (正肖)
				LotteryPlay = 1
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResultZodiac

				//玩法群組 (五行)
				LotteryPlayGroup = 16

				//玩法 (五行)
				LotteryPlay = 1
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResult2Element

				//玩法群組 (自选不中)
				LotteryPlayGroup = 17

				OpenResultAll := OpenResult
				OpenResultAll = append(OpenResult, OpenResultAll...)

				//玩法 (自选不中)
				LotteryPlay = 1
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResultAll
			default:
				// State = false
				break //彩種群組錯誤
			}
		}
	}
	return Result
}

func ReplaceModuleKuaiThree(lotteryTypeGroup int,
	lotteryType int,
	OpenResult []string,
	OpenResult2 interface{},
	Result1 map[string]interface{},
	config map[string]interface{}) map[string]interface{} {

	Result := Result1

	OpenResultInt := []int{}
	for _, num := range OpenResult {
		numInt, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		OpenResultInt = append(OpenResultInt, numInt)
	}
	//初始化
	// State := true
	LotteryPlayMode := 0
	LotteryPlayGroup := 0
	LotteryPlay := 0

	//取得大小的判斷基準 (LTG: 1, 2, 4, 7, 10)
	// BigOrSmallBase := int(config["LTR_BigOrSmallBase"].(float64))
	// 取得龙虎的判斷基準 (LTG: 1, 3, 6, 10)
	// DragonOrTigerBase := int(config["LTR_DragonOrTigerBase"].(float64))
	//取得总和計算的元素個數 (LTG: 1, 3, 6, 7)
	TotalBase := int(config["LTR_TotalBase"].(float64))
	//取得总和大小的判斷基準 (LTG: 1, 2, 3, 5, 6, 7, 10)
	TotalBigOrSmallBase := int(config["LTR_TotalBigOrSmallBase"].(float64))
	//取得極大的判斷基準 (LTG: 4) 尾大尾小的的判斷基準 (LTG: 10)
	// ExtremeBigBase := config["LTR_ExtremeBigBase"]
	//取得極小的判斷基準 (LTG: 4) 尾大尾小的的判斷基準 (LTG: 10)
	// ExtremeSmallBase := config["LTR_ExtremeSmallBase"]
	//取得上下盘的判斷基準 (LTG: 8)
	// TopOrUnderBase := config["LTR_TopOrUnderBase"]
	//取得特碼大小的判斷基準 (LTG: 5)
	// UniqueBigOrSmallBase := config["LTR_UniqueBigOrSmallBase"]
	//取得全部位數和大小的判斷基準 (LTG: 5)
	// DigitBigOrSmallBase := config["LTR_DigitBigOrSmallBase"]
	//取得三個位數和大小的判斷基準 (LTG: 2)
	// TriplexBigOrSmallBase := config["LTR_TriplexBigOrSmallBase"]
	//取得和局的判斷基準 (LTG: 5,7)
	// DrawBase := config["LTR_DrawBase"]
	//取得總和和局的判斷基準 (LTG: 7,8,10)
	// TotalDrawBase := config["LTR_TotalDrawBase"]
	//取得波色的分類資料 (LTG: 5)
	// Wave := config["LTR_Wave"]
	//取得生肖的分類資料 (LTG: 5)
	// Zodiac := config["LTR_Zodiac"]
	//取得五行的分類資料 (LTG: 5)
	// Elements := config["LTR_FiveElements"]

	switch lotteryTypeGroup {
	//快三
	case 6:
		{
			//彩種 (因為有對應 TS_LotteryType.LT_Code，所以需要跟資料庫一起變動...ry)
			switch lotteryType {

			//江苏快三 //安徽快三 //广西快三// 1分快三 //3分快三 //區域鏈快三
			case 8, 9, 10, 36, 37, 42:
				LotteryPlayMode = 1  //信用模式 (传统)
				LotteryPlayGroup = 1 //玩法群組 (三军)
				LotteryPlay = 1
				OpenResultTotal := 0 //总和大 or 总和小 (TB|TS)

				//OpenResultTotal = OpenResultInt[0] + OpenResultInt[1]

				for i := 0; i < TotalBase; i++ {
					OpenResultTotal = OpenResultTotal + OpenResultInt[i]
				}

				//总和大 or 总和小 (TB|TS)
				s := ""
				//t := Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]interface{})
				t := make([]interface{}, 0)

				for i := 0; i < len(OpenResult); i++ {
					t = append([]interface{}{OpenResult[i]}, t...)
				}
				//t = append([]interface{}{OpenResult}, t...)
				//冠亚和大 or 小 (B|S)
				if OpenResultTotal > TotalBigOrSmallBase {
					s = "B"
				} else {
					s = "S"
				}
				t = append([]interface{}{s}, t...)

				//冠亚和单 or 双 (O|E)
				if OpenResultTotal%2 > 0 {
					s = "O"
				} else {
					s = "E"
				}

				t = append([]interface{}{s}, t...)
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = t

				// arr := Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})

				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = t

				LotteryPlayGroup = 2 //玩法群組  (围骰)
				LotteryPlay = 1
				t = make([]interface{}, 0)

				//resultTemp := betCount.Array_unique(OpenResultTemp)

				t = append([]interface{}{OpenResult}, t...)
				resultTemp := betCount.Array_unique(OpenResultInt)

				if len(resultTemp) == 1 { //豹子
					t = append([]interface{}{"TK"}, t...)
				}
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = t

				LotteryPlayGroup = 3 //玩法群組   (点数)
				LotteryPlay = 1
				// t = make([]interface{}, 0)
				// t = append([]interface{}{OpenResultTotal}, t...)
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResultTotal

				LotteryPlayGroup = 4 //玩法群組 (长牌)
				LotteryPlay = 1
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResult
				LotteryPlayGroup = 5 //玩法群組 (短牌)
				LotteryPlay = 1
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResult

				//官方模式

				LotteryPlayMode = 2
				//玩法群組 (同号)
				LotteryPlayGroup = 1
				LotteryPlay = 1
				t = make([]interface{}, 0)
				resultTemp = betCount.Array_unique(OpenResultInt)

				if len(resultTemp) == 1 { //豹子
					t = append([]interface{}{"TK"}, t...)
				}
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = t

				LotteryPlay = 2
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResult
				LotteryPlay = 3
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResult
				LotteryPlay = 4
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResult
				LotteryPlay = 5
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResult

				//玩法群組 (三连号)
				LotteryPlayGroup = 2
				LotteryPlay = 1
				OpenResultSubInt := []int{OpenResultInt[0], OpenResultInt[1], OpenResultInt[2]}

				sort.Ints(OpenResultSubInt)

				t = make([]interface{}, 0)
				if OpenResultSubInt[0]+1 == OpenResultSubInt[1] && OpenResultSubInt[1]+1 == OpenResultSubInt[2] {

					t = append([]interface{}{"SF"}, t...)

				}
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = t
				///玩法群組 (不同号)
				LotteryPlayGroup = 3
				LotteryPlay = 1
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResult
				LotteryPlay = 2
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResult
				LotteryPlay = 3
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResult
				LotteryPlay = 4
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResult
				LotteryPlay = 5
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResult
				LotteryPlay = 6
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResult

			default:
				// State = false
				break //彩種群組錯誤
			}
		}
	}
	return Result
}
func ReplaceModulePCEgg(lotteryTypeGroup int,
	lotteryType int,
	OpenResult []string,
	OpenResult2 interface{},
	Result1 map[string]interface{},
	config map[string]interface{}) map[string]interface{} {

	Result := Result1

	OpenResultInt := []int{}
	OpenResultFloat64 := []float64{}

	for _, num := range OpenResult {
		numInt, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		OpenResultInt = append(OpenResultInt, numInt)
		OpenResultFloat64 = append(OpenResultFloat64, float64(numInt))
	}
	//初始化
	// State := true
	LotteryPlayMode := 0
	LotteryPlayGroup := 0
	LotteryPlay := 0

	//取得大小的判斷基準 (LTG: 1, 2, 4, 7, 10)
	BigOrSmallBase := int(config["LTR_BigOrSmallBase"].(float64))
	// 取得龙虎的判斷基準 (LTG: 1, 3, 6, 10)
	//DragonOrTigerBase := int(config["LTR_DragonOrTigerBase"].(float64))
	//取得总和計算的元素個數 (LTG: 1, 3, 6, 7)
	//TotalBase := int(config["LTR_TotalBase"].(float64))
	//取得总和大小的判斷基準 (LTG: 1, 2, 3, 5, 6, 7, 10)
	//TotalBigOrSmallBase := int(config["LTR_TotalBigOrSmallBase"].(float64))
	//取得極大的判斷基準 (LTG: 4) 尾大尾小的的判斷基準 (LTG: 10)
	ExtremeBigBase := int(config["LTR_ExtremeBigBase"].(float64))
	//取得極小的判斷基準 (LTG: 4) 尾大尾小的的判斷基準 (LTG: 10)
	ExtremeSmallBase := int(config["LTR_ExtremeSmallBase"].(float64))
	//取得上下盘的判斷基準 (LTG: 8)
	// TopOrUnderBase := config["LTR_TopOrUnderBase"]
	//取得特碼大小的判斷基準 (LTG: 5)
	// UniqueBigOrSmallBase := config["LTR_UniqueBigOrSmallBase"]
	//取得全部位數和大小的判斷基準 (LTG: 5)
	// DigitBigOrSmallBase := config["LTR_DigitBigOrSmallBase"]
	//取得三個位數和大小的判斷基準 (LTG: 2)
	// TriplexBigOrSmallBase := config["LTR_TriplexBigOrSmallBase"]
	//取得和局的判斷基準 (LTG: 5,7)
	// DrawBase := config["LTR_DrawBase"]
	//取得總和和局的判斷基準 (LTG: 7,8,10)
	// TotalDrawBase := config["LTR_TotalDrawBase"]
	//取得波色的分類資料 (LTG: 5)
	Wave := config["LTR_Wave"].(map[string]interface{})
	//取得生肖的分類資料 (LTG: 5)
	// Zodiac := config["LTR_Zodiac"]
	//取得五行的分類資料 (LTG: 5)
	// Elements := config["LTR_FiveElements"]n

	switch lotteryTypeGroup {
	//PC蛋蛋 (幸运28)
	case 4:
		{
			//彩種 (因為有對應 TS_LotteryType.LT_Code，所以需要跟資料庫一起變動...ry)
			switch lotteryType {
			//PC蛋蛋 (幸运28)
			case 6, 33, 40, 41: //PC蛋蛋 //區塊鏈 PC蛋蛋 // 1分 PC蛋蛋 // 3分 PC蛋蛋
				LotteryPlayMode = 1  //信用模式 (传统)
				LotteryPlayGroup = 1 //玩法群組 (混合)
				LotteryPlay = 1      //大 or 小 (B|S)
				OpenResultTotal := 0
				OpenResultTotal = OpenResultInt[0] + OpenResultInt[1] + OpenResultInt[2]

				//大 or 小 (B|S)
				s := ""
				//t := Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]interface{})
				t := make([]interface{}, 0)

				//大 or 小 (B|S)
				if OpenResultTotal > BigOrSmallBase {
					s = "B"
				} else {
					s = "S"
				}
				t = append([]interface{}{s}, t...)

				//极大 or 极小 (XB|XS)
				if OpenResultTotal > ExtremeBigBase {
					s = "XB"
					t = append([]interface{}{s}, t...)
				} else if OpenResultTotal < ExtremeSmallBase {
					s = "XS"
					t = append([]interface{}{s}, t...)
				}

				s = ""
				//单 or 双 (O|E)
				if OpenResultTotal%2 > 0 {
					s = "O"

				} else {
					s = "E"
				}
				t = append([]interface{}{s}, t...)

				//豹子 (TK)
				resultTemp := betCount.Array_unique(OpenResultInt)
				if len(resultTemp) == 1 { //豹子
					t = append([]interface{}{"TK"}, t...)
				}
				//色波

				if betCount.In_array(OpenResultFloat64, Wave["Red"]) {
					t = append([]interface{}{"RW"}, t...)
				} else if betCount.In_array(OpenResultFloat64, Wave["Cyan"]) {
					t = append([]interface{}{"CW"}, t...)
				} else if betCount.In_array(OpenResultFloat64, Wave["Green"]) {
					t = append([]interface{}{"GW"}, t...)
				} else if betCount.In_array(OpenResultFloat64, Wave["Silver"]) {
					t = append([]interface{}{"SW"}, t...)
				}
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = t

				s = ""
				t = make([]interface{}, 0)

				LotteryPlayGroup = 2 //玩法群組 (特码)
				LotteryPlay = 1      //玩法 (特码)

				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResultTotal

			default:
				// State = false
				break //彩種群組錯誤
			}
		}
	}

	return Result
}
func ReplaceModulePKTen(lotteryTypeGroup int,
	lotteryType int,
	OpenResult []string,
	OpenResult2 interface{},
	Result1 map[string]interface{},
	config map[string]interface{}) map[string]interface{} {

	Result := Result1

	OpenResultInt := []int{}
	for _, num := range OpenResult {
		numInt, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		OpenResultInt = append(OpenResultInt, numInt)
	}
	//初始化
	// State := true
	LotteryPlayMode := 0
	LotteryPlayGroup := 0
	LotteryPlay := 0

	//取得大小的判斷基準 (LTG: 1, 2, 4, 7, 10)
	BigOrSmallBase := int(config["LTR_BigOrSmallBase"].(float64))
	// 取得龙虎的判斷基準 (LTG: 1, 3, 6, 10)
	DragonOrTigerBase := int(config["LTR_DragonOrTigerBase"].(float64))
	//取得总和計算的元素個數 (LTG: 1, 3, 6, 7)
	//TotalBase := int(config["LTR_TotalBase"].(float64))
	//取得总和大小的判斷基準 (LTG: 1, 2, 3, 5, 6, 7, 10)
	TotalBigOrSmallBase := int(config["LTR_TotalBigOrSmallBase"].(float64))
	//取得極大的判斷基準 (LTG: 4) 尾大尾小的的判斷基準 (LTG: 10)
	// ExtremeBigBase := config["LTR_ExtremeBigBase"]
	//取得極小的判斷基準 (LTG: 4) 尾大尾小的的判斷基準 (LTG: 10)
	// ExtremeSmallBase := config["LTR_ExtremeSmallBase"]
	//取得上下盘的判斷基準 (LTG: 8)
	// TopOrUnderBase := config["LTR_TopOrUnderBase"]
	//取得特碼大小的判斷基準 (LTG: 5)
	// UniqueBigOrSmallBase := config["LTR_UniqueBigOrSmallBase"]
	//取得全部位數和大小的判斷基準 (LTG: 5)
	// DigitBigOrSmallBase := config["LTR_DigitBigOrSmallBase"]
	//取得三個位數和大小的判斷基準 (LTG: 2)
	// TriplexBigOrSmallBase := config["LTR_TriplexBigOrSmallBase"]
	//取得和局的判斷基準 (LTG: 5,7)
	// DrawBase := config["LTR_DrawBase"]
	//取得總和和局的判斷基準 (LTG: 7,8,10)
	// TotalDrawBase := config["LTR_TotalDrawBase"]
	//取得波色的分類資料 (LTG: 5)
	// Wave := config["LTR_Wave"]
	//取得生肖的分類資料 (LTG: 5)
	// Zodiac := config["LTR_Zodiac"]
	//取得五行的分類資料 (LTG: 5)
	// Elements := config["LTR_FiveElements"]

	switch lotteryTypeGroup {
	//北京PK10  	//幸运飞艇
	case 1, 3:
		{
			//彩種 (因為有對應 TS_LotteryType.LT_Code，所以需要跟資料庫一起變動...ry)
			switch lotteryType {
			//北京PK10  //幸运飞艇
			case 1, 5, 31, 34, 35: //北京PK10 //幸运飞艇 // 區塊鏈 PK10// 1分PK10 // 3分PK10
				LotteryPlayMode = 1  //信用模式 (传统)
				LotteryPlayGroup = 1 //======玩法群組 (两面)=====
				LotteryPlay = 1      //玩法 (冠亚和)
				OpenResultTotal := 0 //总和大 or 总和小 (TB|TS)

				OpenResultTotal = OpenResultInt[0] + OpenResultInt[1]

				//总和大 or 总和小 (TB|TS)
				s := ""
				//t := Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]interface{})
				t := make([]interface{}, 0)

				//冠亚和大 or 小 (B|S)
				if OpenResultTotal > TotalBigOrSmallBase {
					s = "B"
				} else {
					s = "S"
				}
				t = append([]interface{}{s}, t...)

				//冠亚和单 or 双 (O|E)
				if OpenResultTotal%2 > 0 {
					s = "O"
				} else {
					s = "E"
				}

				t = append([]interface{}{s}, t...)
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = t

				arr := Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})

				// 	//寫入開獎結果
				for LotteryPlay := 2; LotteryPlay <= len(arr); LotteryPlay++ {

					t := make([]interface{}, 0)
					//大 or 小 (B|S)
					if OpenResultInt[LotteryPlay-2] > BigOrSmallBase {
						s = "B"

					} else {
						s = "B"
					}
					t = append([]interface{}{s}, t...)
					//单 or 双 (O|E)
					if OpenResultInt[LotteryPlay-2]%2 > 0 {
						s = "O"

					} else {
						s = "E"
					}
					t = append([]interface{}{s}, t...)

					//龙 or 虎 (D|T)
					if OpenResultInt[LotteryPlay-2] < DragonOrTigerBase+1 {
						s = "D"

					} else {
						s = "T"
					}
					t = append([]interface{}{s}, t...)
					Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = t

				}

				LotteryPlayGroup = 2 //玩法群組 (冠亚和)
				LotteryPlay = 1      //冠、亚军 组合
				t = make([]interface{}, 0)
				OpenResultTotal = OpenResultInt[0] + OpenResultInt[1]

				t = append([]interface{}{OpenResultTotal}, t...)
				if OpenResultTotal > TotalBigOrSmallBase {
					s = "TB"
				} else {
					s = "TS"
				}
				t = append([]interface{}{s}, t...)
				if OpenResultTotal%2 > 0 {
					s = "TO"
				} else {
					s = "TE"
				}
				t = append([]interface{}{s}, t...)
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = t
				//玩法群組 (1-5名)
				LotteryPlayGroup = 3
				// {
				//寫入對獎結果
				//arr = Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})

				for i := 0; i < 5; i++ {
					//玩法 (冠军|亚军|第三名|第四名|第五名)

					LotteryPlay1 := strconv.Itoa(i + 1)

					Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[LotteryPlay1] = OpenResult[i]
				}
				//玩法群組 (6-10名)
				LotteryPlayGroup = 4
				// {
				//寫入對獎結果
				arr = Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})
				//arr = Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})
				for i := 5; i < 10; i++ {
					//玩法 (冠军|亚军|第三名|第四名|第五名)
					LotteryPlay1 := strconv.Itoa(i - 4)
					Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[LotteryPlay1] = OpenResult[i]
					// 		$MR_Result->{$MR_LotteryPlayMode}->{$MR_LotteryPlayGroup}->{$MR_LotteryPlay} = $MR_OpenResult[$i];
				}
				// }
				//官方模式
				OpenResultSub := []string{}
				LotteryPlayMode = 2
				//玩法群組 (前一)
				LotteryPlayGroup = 1
				LotteryPlay = 1 //玩法 (前一：直选复式)
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResult[0]

				//玩法群組 (前二)
				LotteryPlayGroup = 2
				LotteryPlay = 1 //玩法 (前二：直选复式)
				OpenResultSub = []string{OpenResult[0], OpenResult[1]}
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResultSub
				LotteryPlay = 2 //玩法 (前二：直选单式)
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResultSub
				//玩法群組 (前三)
				LotteryPlayGroup = 3
				LotteryPlay = 1 //玩法 (前三：直选复式)
				OpenResultSub = []string{OpenResult[0], OpenResult[1], OpenResult[2]}
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResultSub
				LotteryPlay = 2 //玩法 (前三：直选单式)
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResultSub
				//玩法群組 (定位胆)
				LotteryPlayGroup = 4
				LotteryPlay = 1 //玩法 (定位胆)
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResult
			default:
				// State = false
				break //彩種群組錯誤
			}
		}
	}
	return Result
}
func ReplaceModuleShiShi(lotteryTypeGroup int,
	lotteryType int,
	OpenResult []string,
	OpenResult2 interface{},
	Result1 map[string]interface{},
	config map[string]interface{}) map[string]interface{} {

	Result := Result1
	OpenResultInt := []int{}
	for _, num := range OpenResult {
		numInt, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		OpenResultInt = append(OpenResultInt, numInt)
	}
	//初始化
	// State := true
	LotteryPlayMode := 0
	LotteryPlayGroup := 0
	LotteryPlay := 0

	//取得大小的判斷基準 (LTG: 1, 2, 4, 7, 10)
	BigOrSmallBase := int(config["LTR_BigOrSmallBase"].(float64))
	// 取得龙虎的判斷基準 (LTG: 1, 3, 6, 10)
	DragonOrTigerBase := int(config["LTR_DragonOrTigerBase"].(float64))
	//取得总和計算的元素個數 (LTG: 1, 3, 6, 7)
	TotalBase := int(config["LTR_TotalBase"].(float64))
	//取得总和大小的判斷基準 (LTG: 1, 2, 3, 5, 6, 7, 10)
	TotalBigOrSmallBase := int(config["LTR_TotalBigOrSmallBase"].(float64))
	//取得極大的判斷基準 (LTG: 4) 尾大尾小的的判斷基準 (LTG: 10)
	// ExtremeBigBase := config["LTR_ExtremeBigBase"]
	//取得極小的判斷基準 (LTG: 4) 尾大尾小的的判斷基準 (LTG: 10)
	// ExtremeSmallBase := config["LTR_ExtremeSmallBase"]
	//取得上下盘的判斷基準 (LTG: 8)
	// TopOrUnderBase := config["LTR_TopOrUnderBase"]
	//取得特碼大小的判斷基準 (LTG: 5)
	// UniqueBigOrSmallBase := config["LTR_UniqueBigOrSmallBase"]
	//取得全部位數和大小的判斷基準 (LTG: 5)
	// DigitBigOrSmallBase := config["LTR_DigitBigOrSmallBase"]
	//取得三個位數和大小的判斷基準 (LTG: 2)
	// TriplexBigOrSmallBase := config["LTR_TriplexBigOrSmallBase"]
	//取得和局的判斷基準 (LTG: 5,7)
	// DrawBase := config["LTR_DrawBase"]
	//取得總和和局的判斷基準 (LTG: 7,8,10)
	// TotalDrawBase := config["LTR_TotalDrawBase"]
	//取得波色的分類資料 (LTG: 5)
	// Wave := config["LTR_Wave"]
	//取得生肖的分類資料 (LTG: 5)
	// Zodiac := config["LTR_Zodiac"]
	//取得五行的分類資料 (LTG: 5)
	// Elements := config["LTR_FiveElements"]

	switch lotteryTypeGroup {
	//时时彩
	case 2:
		{
			//彩種 (因為有對應 TS_LotteryType.LT_Code，所以需要跟資料庫一起變動...ry)
			switch lotteryType {
			//重庆时时彩 //新疆时时彩 //天津时时彩 //精彩1分彩 //精彩3分彩 //精彩5分彩 //精彩秒秒彩 // 區塊鏈時時彩
			case 2, 3, 4, 26, 27, 28, 29, 30:
				LotteryPlayMode = 1  //信用模式 (传统)
				LotteryPlayGroup = 1 //======玩法群組 (两面)=====
				LotteryPlay = 1      //总和值計算
				OpenResultTotal := 0 //总和大 or 总和小 (TB|TS)
				for j := 0; j < TotalBase; j++ {
					OpenResultTotal = OpenResultTotal + OpenResultInt[j]
				}
				//总和大 or 总和小 (TB|TS)
				s := ""
				//t := Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]interface{})
				t := make([]interface{}, 0)
				if OpenResultTotal > TotalBigOrSmallBase {
					s = "TB"
				} else {
					s = "TS"
				}
				t = append([]interface{}{s}, t...)
				//总和单 or 总和双 (TO|TE)
				if OpenResultTotal%2 > 0 {
					s = "TO"
				} else {
					s = "TE"
				}
				t = append([]interface{}{s}, t...)

				for k := 0; k < DragonOrTigerBase-1; k++ {

					if OpenResultInt[k] > OpenResultInt[len(OpenResult)-k-1] {
						s = "D"
					} else if OpenResultInt[k] < OpenResultInt[len(OpenResult)-k-1] {
						s = "T"
					} else {
						s = "N"
					}
					t = append([]interface{}{s}, t...)
				}
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = t

				//arrList := Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})
				//使用迴圈依球號寫入開獎結果
				for i := 0; i < len(OpenResult); i++ {
					//玩法 (第一球|第二球|第三球|第四球|第五球)
					LotteryPlay = i + 2
					t = make([]interface{}, 0)
					ResultI, err := strconv.Atoi(OpenResult[i])
					//t := Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]interface{})
					if err == nil {
					}
					if ResultI > BigOrSmallBase {
						s = "B"

					} else {
						s = "S"
					}
					t = append([]interface{}{s}, t...)

					if ResultI%2 > 0 {
						s = "O"

					} else {
						s = "E"
					}
					t = append([]interface{}{s}, t...)

					Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = t
				}
				LotteryPlayGroup = 2 //======玩法群組 (1-5球)=====
				arr := Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})
				//寫入開獎結果
				for key := range arr {
					LotteryPlay, err := strconv.ParseInt(key, 10, 64)
					if err == nil {
					}
					Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[key] = OpenResult[LotteryPlay-1]
				}

				LotteryPlayGroup = 3 //玩法群組 (前中后)
				OpenResultTemp := []int{0, 0, 0}
				//玩法 (前三|中三|后三)
				for LotteryPlay := 1; LotteryPlay <= 3; LotteryPlay++ {
					// //玩法 (前三)
					if LotteryPlay == 1 {
						OpenResultTemp = []int{OpenResultInt[0], OpenResultInt[1], OpenResultInt[2]}
					}
					// //玩法 (中三)
					if LotteryPlay == 2 {
						OpenResultTemp = []int{OpenResultInt[1], OpenResultInt[2], OpenResultInt[3]}
					}
					// //玩法 (后三)
					if LotteryPlay == 3 {
						OpenResultTemp = []int{OpenResultInt[2], OpenResultInt[3], OpenResultInt[4]}
					}
					resultTemp := betCount.Array_unique(OpenResultTemp)
					t = make([]interface{}, 0)
					//豹子
					if len(resultTemp) == 1 {
						t = append([]interface{}{"TK"}, t...)
					} else if len(resultTemp) == 2 { //對子
						t = append([]interface{}{"OP"}, t...)
					} else { //排除豹子、对子
						OpenResultTempSF := OpenResultTemp //顺子
						if OpenResultTempSF[0] == 0 && OpenResultTempSF[1] == 9 {
							OpenResultTempSF[0] = OpenResultTempSF[0] + 10
						} else if OpenResultTempSF[1] == 0 && OpenResultTempSF[0] == 9 {
							OpenResultTempSF[2] = OpenResultTempSF[2] + 10
							OpenResultTempSF[1] = OpenResultTempSF[1] + 10
						} else if OpenResultTempSF[1] == 0 && OpenResultTempSF[2] == 9 {
							OpenResultTempSF[0] = OpenResultTempSF[0] + 10
							OpenResultTempSF[1] = OpenResultTempSF[1] + 10
						} else if OpenResultTempSF[2] == 0 && OpenResultTempSF[1] == 9 {
							OpenResultTempSF[2] = OpenResultTempSF[2] + 10
						}

						if (OpenResultTempSF[2]-OpenResultTempSF[1] == 1 && OpenResultTempSF[1]-OpenResultTempSF[0] == 1) ||
							(OpenResultTempSF[2]-OpenResultTempSF[1] == -1 && OpenResultTempSF[1]-OpenResultTempSF[0] == -1) {
							t = append([]interface{}{"SF"}, t...)
						} else { //半顺 - 不依開獎順序
							sort.Ints(OpenResultTemp)
							inArrayResult, _ := in_array(9, OpenResultTemp)
							if OpenResultTemp[0] == 0 && inArrayResult {
								OpenResultTemp[0] = OpenResultTemp[0] + 10
								sort.Ints(OpenResultTemp)
							}

							if OpenResultTemp[2]-OpenResultTemp[1] == 1 || OpenResultTemp[1]-OpenResultTemp[0] == 1 || OpenResultTemp[2]-OpenResultTemp[1] == -1 || OpenResultTemp[1]-OpenResultTemp[0] == -1 {
								t = append([]interface{}{"HF"}, t...)
							} else { //杂六
								t = append([]interface{}{"HC"}, t...)
							}

						}

					}
					Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = t

				}

				LotteryPlayMode = 2  //官方模式
				LotteryPlayGroup = 1 //玩法群組 (定位胆)
				LotteryPlay = 1      //玩法 (直选复式)
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResult

				LotteryPlayGroup = 2 //玩法群組 (五星)
				for LotteryPlay := 1; LotteryPlay <= 8; LotteryPlay++ {
					//玩法 (直选复式) 1
					//玩法 (组选120) 2
					//玩法 (组选60) 3
					//玩法 (组选30) 4
					//玩法 (组选20) 5
					//玩法 (组选10) 6
					//玩法 (组选5) 7
					//玩法 (直选单式) 8
					Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResult
				}
				OpenResultSub := []string{}
				OpenResultSubInt := []int{}
				LotteryPlayGroup = 3 //玩法群組 (四星)
				OpenResultSub = []string{OpenResult[1], OpenResult[2], OpenResult[3], OpenResult[4]}
				for LotteryPlay := 1; LotteryPlay <= 6; LotteryPlay++ {
					//玩法 (直选复式)
					//玩法 (组选24)
					//玩法 (组选12)
					//玩法 (组选6)
					//玩法 (组选4)
					//玩法 (直选单式)
					Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResultSub
				}

				LotteryPlayGroup = 4 //玩法群組 (后三)
				OpenResultSub = []string{OpenResult[2], OpenResult[3], OpenResult[4]}
				OpenResultSubInt = []int{OpenResultInt[2], OpenResultInt[3], OpenResultInt[4]}
				for LotteryPlay := 1; LotteryPlay <= 14; LotteryPlay++ {
					if LotteryPlay != 9 && LotteryPlay != 10 {
						//玩法 (直选复式) 1
						//玩法 (直选和值) 2
						//玩法 (直选跨度) 3
						//玩法 (后三组合) 4
						//玩法 (组三复式) 5
						//玩法 (组六复式) 6
						//玩法 (组选和值) 7
						//玩法 (组选包胆) 8
						//玩法 (直选单式) 11
						//玩法 (組三单式) 12
						//玩法 (組六单式) 13
						//玩法 (混合组选) 14
						Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResultSub
					} else if LotteryPlay == 9 {
						//玩法 (和值尾数) 9
						TotalCount := 0
						for j := 0; j < len(OpenResultSubInt); j++ {
							TotalCount = TotalCount + OpenResultSubInt[j]
						}
						Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = strconv.Itoa(TotalCount % 10)
					} else if LotteryPlay == 10 {
						//玩法 (特殊号)

						t := make([]interface{}, 0)
						resultTemp := betCount.Array_unique(OpenResultSubInt)
						if len(resultTemp) == 1 { //豹子
							t = append([]interface{}{"TK"}, t...)
						} else if len(resultTemp) == 2 { //對子
							t = append([]interface{}{"OP"}, t...)
						}

						OpenResultTempSF := OpenResultTemp
						if OpenResultTempSF[0] == 0 && OpenResultTempSF[1] == 9 {
							OpenResultTempSF[0] = OpenResultTempSF[0] + 10
						} else if OpenResultTempSF[1] == 0 && OpenResultTempSF[0] == 9 {
							OpenResultTempSF[2] = OpenResultTempSF[2] + 10
							OpenResultTempSF[1] = OpenResultTempSF[1] + 10
						} else if OpenResultTempSF[1] == 0 && OpenResultTempSF[2] == 9 {
							OpenResultTempSF[0] = OpenResultTempSF[0] + 10
							OpenResultTempSF[1] = OpenResultTempSF[1] + 10
						} else if OpenResultTempSF[2] == 0 && OpenResultTempSF[1] == 9 {
							OpenResultTempSF[2] = OpenResultTempSF[2] + 10
						}

						if (OpenResultTempSF[2]-OpenResultTempSF[1] == 1 && OpenResultTempSF[1]-OpenResultTempSF[0] == 1) ||
							(OpenResultTempSF[2]-OpenResultTempSF[1] == -1 && OpenResultTempSF[1]-OpenResultTempSF[0] == -1) {
							t = append([]interface{}{"SF"}, t...)
						}
						Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = t
					}
				}

				LotteryPlayGroup = 5 //玩法群組 (前三)
				OpenResultSub = []string{OpenResult[0], OpenResult[1], OpenResult[2]}
				OpenResultSubInt = []int{OpenResultInt[0], OpenResultInt[1], OpenResultInt[2]}
				for LotteryPlay := 1; LotteryPlay <= 14; LotteryPlay++ {
					if LotteryPlay != 9 && LotteryPlay != 10 {
						//玩法 (直选复式) 1
						//玩法 (直选和值) 2
						//玩法 (直选跨度) 3
						//玩法 (后三组合) 4
						//玩法 (组三复式) 5
						//玩法 (组六复式) 6
						//玩法 (组选和值) 7
						//玩法 (组选包胆) 8
						//玩法 (直选单式) 11
						//玩法 (組三单式) 12
						//玩法 (組六单式) 13
						//玩法 (混合组选) 14
						Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResultSub
					} else if LotteryPlay == 9 {
						//玩法 (和值尾数) 9
						TotalCount := 0
						for j := 0; j < len(OpenResultSub); j++ {
							TotalCount = TotalCount + OpenResultSubInt[j]
						}
						Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = strconv.Itoa(TotalCount % 10)
					} else if LotteryPlay == 10 {
						//玩法 (特殊号)
						//t := Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]interface{})
						t := make([]interface{}, 0)
						resultTemp := betCount.Array_unique(OpenResultSubInt)
						if len(resultTemp) == 1 { //豹子
							t = append([]interface{}{"TK"}, t...)
						} else if len(resultTemp) == 2 { //對子
							t = append([]interface{}{"OP"}, t...)
						}

						OpenResultTempSF := OpenResultTemp
						if OpenResultTempSF[0] == 0 && OpenResultTempSF[1] == 9 {
							OpenResultTempSF[0] = OpenResultTempSF[0] + 10
						} else if OpenResultTempSF[1] == 0 && OpenResultTempSF[0] == 9 {
							OpenResultTempSF[2] = OpenResultTempSF[2] + 10
							OpenResultTempSF[1] = OpenResultTempSF[1] + 10
						} else if OpenResultTempSF[1] == 0 && OpenResultTempSF[2] == 9 {
							OpenResultTempSF[0] = OpenResultTempSF[0] + 10
							OpenResultTempSF[1] = OpenResultTempSF[1] + 10
						} else if OpenResultTempSF[2] == 0 && OpenResultTempSF[1] == 9 {
							OpenResultTempSF[2] = OpenResultTempSF[2] + 10
						}

						if (OpenResultTempSF[2]-OpenResultTempSF[1] == 1 && OpenResultTempSF[1]-OpenResultTempSF[0] == 1) ||
							(OpenResultTempSF[2]-OpenResultTempSF[1] == -1 && OpenResultTempSF[1]-OpenResultTempSF[0] == -1) {
							t = append([]interface{}{"SF"}, t...)
						}
						Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = t
					}
				}
				//玩法群組 (中三)
				LotteryPlayGroup = 12
				OpenResultSub = []string{OpenResult[1], OpenResult[2], OpenResult[3]}
				OpenResultSubInt = []int{OpenResultInt[1], OpenResultInt[2], OpenResultInt[3]}
				for LotteryPlay := 1; LotteryPlay <= 14; LotteryPlay++ {
					if LotteryPlay != 9 && LotteryPlay != 10 {
						//玩法 (直选复式) 1
						//玩法 (直选和值) 2
						//玩法 (直选跨度) 3
						//玩法 (后三组合) 4
						//玩法 (组三复式) 5
						//玩法 (组六复式) 6
						//玩法 (组选和值) 7
						//玩法 (组选包胆) 8
						//玩法 (直选单式) 11
						//玩法 (組三单式) 12
						//玩法 (組六单式) 13
						//玩法 (混合组选) 14
						Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResultSub
					} else if LotteryPlay == 9 {
						//玩法 (和值尾数) 9
						TotalCount := 0
						for j := 0; j < len(OpenResultSub); j++ {
							TotalCount = TotalCount + OpenResultSubInt[j]
						}
						Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = strconv.Itoa(TotalCount % 10)
					} else if LotteryPlay == 10 {
						//玩法 (特殊号)
						t := make([]interface{}, 0)
						//t := Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]interface{})
						resultTemp := betCount.Array_unique(OpenResultSubInt)
						if len(resultTemp) == 1 { //豹子
							t = append([]interface{}{"TK"}, t...)
						} else if len(resultTemp) == 2 { //對子
							t = append([]interface{}{"OP"}, t...)
						}

						OpenResultTempSF := OpenResultTemp
						if OpenResultTempSF[0] == 0 && OpenResultTempSF[1] == 9 {
							OpenResultTempSF[0] = OpenResultTempSF[0] + 10
						} else if OpenResultTempSF[1] == 0 && OpenResultTempSF[0] == 9 {
							OpenResultTempSF[2] = OpenResultTempSF[2] + 10
							OpenResultTempSF[1] = OpenResultTempSF[1] + 10
						} else if OpenResultTempSF[1] == 0 && OpenResultTempSF[2] == 9 {
							OpenResultTempSF[0] = OpenResultTempSF[0] + 10
							OpenResultTempSF[1] = OpenResultTempSF[1] + 10
						} else if OpenResultTempSF[2] == 0 && OpenResultTempSF[1] == 9 {
							OpenResultTempSF[2] = OpenResultTempSF[2] + 10
						}

						if (OpenResultTempSF[2]-OpenResultTempSF[1] == 1 && OpenResultTempSF[1]-OpenResultTempSF[0] == 1) ||
							(OpenResultTempSF[2]-OpenResultTempSF[1] == -1 && OpenResultTempSF[1]-OpenResultTempSF[0] == -1) {
							t = append([]interface{}{"SF"}, t...)
						}
						Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = t

					}
				}
				//玩法群組 (前二)
				LotteryPlayGroup = 6
				OpenResultSub = []string{OpenResult[0], OpenResult[1]}
				for LotteryPlay := 1; LotteryPlay <= 8; LotteryPlay++ {
					//玩法 (直选复式) 1
					//玩法 (直选和值) 2
					//玩法 (直选跨度) 3
					//玩法 (组选复式) 4
					//玩法 (组选和值) 5
					//玩法 (组选包胆) 6
					//玩法 (直选单式) 7
					//玩法 (組选单式) 8
					Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResultSub
				}
				//玩法群組 (后二)
				LotteryPlayGroup = 13

				OpenResultSub = []string{OpenResult[3], OpenResult[4]}
				for LotteryPlay := 1; LotteryPlay <= 8; LotteryPlay++ {
					//玩法 (直选复式) 1
					//玩法 (直选和值) 2
					//玩法 (直选跨度) 3
					//玩法 (组选复式) 4
					//玩法 (组选和值) 5
					//玩法 (组选包胆) 6
					//玩法 (直选单式) 7
					//玩法 (組选单式) 8
					Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResultSub
				}
				//玩法群組 (不定位)
				LotteryPlayGroup = 7
				//玩法 (前三一码)
				LotteryPlay = 1
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = []string{OpenResult[0], OpenResult[1], OpenResult[2]}
				//玩法 (前三二码)
				LotteryPlay = 2
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = []string{OpenResult[0], OpenResult[1], OpenResult[2]}
				//玩法 (后三一码)
				LotteryPlay = 3
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = []string{OpenResult[2], OpenResult[3], OpenResult[4]}
				//玩法 (后三二码)
				LotteryPlay = 4
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = []string{OpenResult[2], OpenResult[3], OpenResult[4]}
				//玩法 (前四一码)
				LotteryPlay = 5
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = []string{OpenResult[0], OpenResult[1], OpenResult[2], OpenResult[3]}
				//玩法 (前四二码)
				LotteryPlay = 6
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = []string{OpenResult[0], OpenResult[1], OpenResult[2], OpenResult[3]}
				//玩法 (后四一码)
				LotteryPlay = 7
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = []string{OpenResult[1], OpenResult[2], OpenResult[3], OpenResult[4]}
				//玩法 (后四二码)
				LotteryPlay = 8
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = []string{OpenResult[1], OpenResult[2], OpenResult[3], OpenResult[4]}
				//玩法 (五星一码)
				LotteryPlay = 9
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResult
				//玩法 (五星二码)
				LotteryPlay = 10
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResult
				//玩法 (五星三码)
				LotteryPlay = 11
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResult
				//玩法 (中三一码)
				LotteryPlay = 12
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = []string{OpenResult[1], OpenResult[2], OpenResult[3]}
				//玩法 (中三二码)
				LotteryPlay = 13
				Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = []string{OpenResult[1], OpenResult[2], OpenResult[3]}

				//玩法群組 (双面/串关)
				LotteryPlayGroup = 8
				{
					AlternativeTemp := ""
					//玩法 (前二大小单双)
					LotteryPlay = 1
					t := make([]interface{}, 2)
					t[0] = make([]interface{}, 0)
					t[1] = make([]interface{}, 0)
					k1 := t[0].([]interface{})
					k2 := t[1].([]interface{})
					if OpenResultInt[0]%2 == 1 {
						AlternativeTemp = "O"

					} else {
						AlternativeTemp = "E"
					}
					k1 = append([]interface{}{AlternativeTemp}, k1...)

					if OpenResultInt[1]%2 == 1 {
						AlternativeTemp = "O"

					} else {
						AlternativeTemp = "E"
					}
					k2 = append([]interface{}{AlternativeTemp}, k2...)

					if OpenResultInt[0] > BigOrSmallBase {
						AlternativeTemp = "B"
					} else {
						AlternativeTemp = "S"
					}
					k1 = append([]interface{}{AlternativeTemp}, k1...)

					if OpenResultInt[1] > BigOrSmallBase {
						AlternativeTemp = "B"
					} else {
						AlternativeTemp = "S"
					}
					k2 = append([]interface{}{AlternativeTemp}, k2...)
					t[0] = k1
					t[1] = k2
					Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = t

					//玩法 (后二大小单双)
					LotteryPlay = 2
					t = make([]interface{}, 2)
					t[0] = make([]interface{}, 0)
					t[1] = make([]interface{}, 0)
					k1 = t[0].([]interface{})
					k2 = t[1].([]interface{})
					if OpenResultInt[3]%2 == 1 {
						AlternativeTemp = "O"

					} else {
						AlternativeTemp = "E"
					}
					k1 = append([]interface{}{AlternativeTemp}, k1...)

					if OpenResultInt[4]%2 == 1 {
						AlternativeTemp = "O"

					} else {
						AlternativeTemp = "E"
					}
					k2 = append([]interface{}{AlternativeTemp}, k2...)

					if OpenResultInt[3] > BigOrSmallBase {
						AlternativeTemp = "B"
					} else {
						AlternativeTemp = "S"
					}
					k1 = append([]interface{}{AlternativeTemp}, k1...)

					if OpenResultInt[4] > BigOrSmallBase {
						AlternativeTemp = "B"
					} else {
						AlternativeTemp = "S"
					}
					k2 = append([]interface{}{AlternativeTemp}, k2...)
					t[0] = k1
					t[1] = k2
					Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = t

					//玩法 (前三大小单双)
					LotteryPlay = 3
					t = make([]interface{}, 3)
					t[0] = make([]interface{}, 0)
					t[1] = make([]interface{}, 0)
					t[2] = make([]interface{}, 0)

					k1 = t[0].([]interface{})
					k2 = t[1].([]interface{})
					k3 := t[2].([]interface{})

					if OpenResultInt[0]%2 == 1 {
						AlternativeTemp = "O"

					} else {
						AlternativeTemp = "E"
					}
					k1 = append([]interface{}{AlternativeTemp}, k1...)

					if OpenResultInt[1]%2 == 1 {
						AlternativeTemp = "O"

					} else {
						AlternativeTemp = "E"
					}
					k2 = append([]interface{}{AlternativeTemp}, k2...)
					if OpenResultInt[2]%2 == 1 {
						AlternativeTemp = "O"

					} else {
						AlternativeTemp = "E"
					}
					k3 = append([]interface{}{AlternativeTemp}, k3...)

					if OpenResultInt[0] > BigOrSmallBase {
						AlternativeTemp = "B"
					} else {
						AlternativeTemp = "S"
					}
					k1 = append([]interface{}{AlternativeTemp}, k1...)

					if OpenResultInt[1] > BigOrSmallBase {
						AlternativeTemp = "B"
					} else {
						AlternativeTemp = "S"
					}
					k2 = append([]interface{}{AlternativeTemp}, k2...)

					if OpenResultInt[2] > BigOrSmallBase {
						AlternativeTemp = "B"
					} else {
						AlternativeTemp = "S"
					}
					k3 = append([]interface{}{AlternativeTemp}, k3...)
					t[0] = k1
					t[1] = k2
					t[2] = k3
					Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = t

					//玩法 (后三大小单双)
					LotteryPlay = 4
					t = make([]interface{}, 3)
					t[0] = make([]interface{}, 0)
					t[1] = make([]interface{}, 0)
					t[2] = make([]interface{}, 0)

					k1 = t[0].([]interface{})
					k2 = t[1].([]interface{})
					k3 = t[2].([]interface{})
					if OpenResultInt[2]%2 == 1 {
						AlternativeTemp = "O"

					} else {
						AlternativeTemp = "E"
					}
					k1 = append([]interface{}{AlternativeTemp}, k1...)

					if OpenResultInt[3]%2 == 1 {
						AlternativeTemp = "O"

					} else {
						AlternativeTemp = "E"
					}
					k2 = append([]interface{}{AlternativeTemp}, k2...)
					if OpenResultInt[4]%2 == 1 {
						AlternativeTemp = "O"

					} else {
						AlternativeTemp = "E"
					}
					k3 = append([]interface{}{AlternativeTemp}, k3...)

					if OpenResultInt[2] > BigOrSmallBase {
						AlternativeTemp = "B"
					} else {
						AlternativeTemp = "S"
					}
					k1 = append([]interface{}{AlternativeTemp}, k1...)

					if OpenResultInt[3] > BigOrSmallBase {
						AlternativeTemp = "B"
					} else {
						AlternativeTemp = "S"
					}
					k2 = append([]interface{}{AlternativeTemp}, k2...)

					if OpenResultInt[4] > BigOrSmallBase {
						AlternativeTemp = "B"
					} else {
						AlternativeTemp = "S"
					}
					k3 = append([]interface{}{AlternativeTemp}, k3...)
					t[0] = k1
					t[1] = k2
					t[2] = k3
					Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = t
					//玩法 (串关)
					LotteryPlay = 5
					t = make([]interface{}, 5)
					t[0] = make([]interface{}, 0)
					t[1] = make([]interface{}, 0)
					t[2] = make([]interface{}, 0)
					t[3] = make([]interface{}, 0)
					t[4] = make([]interface{}, 0)
					k1 = t[0].([]interface{})
					k2 = t[1].([]interface{})
					k3 = t[2].([]interface{})
					k4 := t[3].([]interface{})
					k5 := t[4].([]interface{})
					if OpenResultInt[0]%2 == 1 {
						AlternativeTemp = "O"

					} else {
						AlternativeTemp = "E"
					}
					k1 = append([]interface{}{AlternativeTemp}, k1...)

					if OpenResultInt[1]%2 == 1 {
						AlternativeTemp = "O"

					} else {
						AlternativeTemp = "E"
					}
					k2 = append([]interface{}{AlternativeTemp}, k2...)
					if OpenResultInt[2]%2 == 1 {
						AlternativeTemp = "O"

					} else {
						AlternativeTemp = "E"
					}
					k3 = append([]interface{}{AlternativeTemp}, k3...)
					if OpenResultInt[3]%2 == 1 {
						AlternativeTemp = "O"

					} else {
						AlternativeTemp = "E"
					}
					k4 = append([]interface{}{AlternativeTemp}, k4...)

					if OpenResultInt[4]%2 == 1 {
						AlternativeTemp = "O"

					} else {
						AlternativeTemp = "E"
					}
					k5 = append([]interface{}{AlternativeTemp}, k5...)

					if OpenResultInt[0] > BigOrSmallBase {
						AlternativeTemp = "B"
					} else {
						AlternativeTemp = "S"
					}
					k1 = append([]interface{}{AlternativeTemp}, k1...)

					if OpenResultInt[1] > BigOrSmallBase {
						AlternativeTemp = "B"
					} else {
						AlternativeTemp = "S"
					}
					k2 = append([]interface{}{AlternativeTemp}, k2...)

					if OpenResultInt[2] > BigOrSmallBase {
						AlternativeTemp = "B"
					} else {
						AlternativeTemp = "S"
					}
					k3 = append([]interface{}{AlternativeTemp}, k3...)

					if OpenResultInt[3] > BigOrSmallBase {
						AlternativeTemp = "B"
					} else {
						AlternativeTemp = "S"
					}
					k4 = append([]interface{}{AlternativeTemp}, k4...)

					if OpenResultInt[4] > BigOrSmallBase {
						AlternativeTemp = "B"
					} else {
						AlternativeTemp = "S"
					}
					k5 = append([]interface{}{AlternativeTemp}, k5...)

					if IsPrime(OpenResultInt[0]) {
						AlternativeTemp = "P"
					} else {
						AlternativeTemp = "C"
					}
					k1 = append([]interface{}{AlternativeTemp}, k1...)

					if IsPrime(OpenResultInt[1]) {
						AlternativeTemp = "P"
					} else {
						AlternativeTemp = "C"
					}
					k2 = append([]interface{}{AlternativeTemp}, k2...)

					if IsPrime(OpenResultInt[2]) {
						AlternativeTemp = "P"
					} else {
						AlternativeTemp = "C"
					}
					k3 = append([]interface{}{AlternativeTemp}, k3...)

					if IsPrime(OpenResultInt[3]) {
						AlternativeTemp = "P"
					} else {
						AlternativeTemp = "C"
					}
					k4 = append([]interface{}{AlternativeTemp}, k4...)

					if IsPrime(OpenResultInt[4]) {
						AlternativeTemp = "P"
					} else {
						AlternativeTemp = "C"
					}
					k5 = append([]interface{}{AlternativeTemp}, k5...)

					t[0] = k1
					t[1] = k2
					t[2] = k3
					t[3] = k4
					t[4] = k5
					Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = t
				}
				//玩法群組 (任选二)
				LotteryPlayGroup = 9
				{
					for LotteryPlay := 1; LotteryPlay <= 6; LotteryPlay++ {
						//玩法 (直选复式) 1
						//玩法 (直选和值) 2
						//玩法 (組选复式) 3
						//玩法 (組选和值) 4
						//玩法 (直选单式) 5
						//玩法 (直选单式) 6
						Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResult
					}
				}
				//玩法群組 (任选三)
				LotteryPlayGroup = 10
				{
					for LotteryPlay := 1; LotteryPlay <= 9; LotteryPlay++ {
						//玩法 (直选复式) 1
						//玩法 (直选和值) 2
						//玩法 (组三复式) 3
						//玩法 (组六复式) 4
						//玩法 (组选和值) 5
						//玩法 (直选单式) 6
						//玩法 (組三单式) 7
						//玩法 (組六单式) 8
						//玩法 (混合组选) 9
						Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResult
					}
				}
				//玩法群組 (任选四)
				LotteryPlayGroup = 11
				{
					for LotteryPlay := 1; LotteryPlay <= 6; LotteryPlay++ {
						//玩法 (直选复式) 1
						//玩法 (组选24) 2
						//玩法 (组选12) 3
						//玩法 (组选6) 4
						//玩法 (组选4) 5
						//玩法 (直选单式) 6
						Result[strconv.Itoa(LotteryPlayMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] = OpenResult
					}
				}
			default:
				// State = false
				break //彩種錯誤
			}

		}
	default:
		// State = false
		break //彩種群組錯誤
	}

	return Result
}

func in_array(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}
	return
}

func IsPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
