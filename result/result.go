package result

import (
	"fmt"
	betcollect "getReslut/betCollect"
	"getReslut/config"
	"getReslut/public"
	"getReslut/public/redisConnect"

	"math"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/seehuhn/mt19937"
)

func Run(betOrdersData map[string]interface{}, thisStatusData bool, thisRatioData map[string]interface{}, thisLotteryTypeGroupData string, thisLotteryTypeData string, thisLotteryIssueData string, blockChaintime int64) map[string]interface{} {

	rng := rand.New(mt19937.New())
	rng.Seed(time.Now().UnixNano())

	thisRatio := 0.0            //該注單賠率
	thisOdds := 0.0             //可使用的賠率array
	var thisPrize float64 = 0.0 //獎金
	thisDrawn := 0.0            //退回
	thisWinnings := 0.0         //預設單注獎金

	AllBOBetCount := 0  //總注數
	AllPrize := 0.0     //總贏錢
	AllRealPrice := 0.0 //總下注

	// targetRatio := 0.96 //目標殺率
	// isWrongOpenResult := true //取得正確開獎數字

	//====開獎的組數====
	allResultCount := 100 //開獎的組數
	if thisLotteryTypeGroupData == "6" && blockChaintime == 0 {
		allResultCount = 216
	}
	//====開獎的組數====
	resultRandCount := float64(allResultCount)       //總權重
	resultArr := make([]interface{}, allResultCount) // 存各組開獎結果 輸贏資料
	var availableList []interface{}                  //可用來調整權重的Index 大於目標RTP的獎號

	var availableList1 []interface{} //可用來調整權重的Index  小於目標RTP的獎號

	var isOnlyTwoWayBet bool = true

	thisState := 0
	thisOpenResult := ""
	errorCode := 0
	thisLottteryTypeGroup := 0
	thisLotteryType := 0

	betOrders := make(map[string]interface{})

	targetRatio, _ := strconv.ParseFloat(thisRatioData["result"].(string), 64) //目標殺率
	nowAmount, _ := thisRatioData["amount"].(float64)
	nowBonus, _ := thisRatioData["bonus"].(float64)
	nowRatio := nowBonus / nowAmount

	//單機測試
	// jsonFile, err := os.Open("config/test.json") //讀取json檔案
	// var boData map[string]interface{}
	// if err != nil { //讀取檔案錯誤
	// 	fmt.Println(err)
	// } else {
	// 	defer jsonFile.Close()                     //關閉檔案流
	// 	byteValue, _ := ioutil.ReadAll(jsonFile)   //json string轉為byte
	// 	json.Unmarshal([]byte(byteValue), &boData) //json轉為map
	// }

	if thisStatusData == true {
		if targetRatio > 0 { //殺率
			// targetRatio = boData["ratio"].(float64)
		} else {
			targetRatio = 0
			thisState = -1
			errorCode = 1
		}

		if lottteryTypeGroup, err := strconv.Atoi(thisLotteryTypeGroupData); err == nil {
			thisLottteryTypeGroup = lottteryTypeGroup
		} else {
			thisState = -1
			errorCode = 1
		}

		if lotteryType, err := strconv.Atoi(thisLotteryTypeData); err == nil {
			thisLotteryType = lotteryType
		} else {
			thisState = -1
			errorCode = 1
		}

		thisLotteryConfig := config.Init(thisLottteryTypeGroup, thisLotteryType)

		// thisLotteryConfig := config.ConfigData

		if betOrdersData["count"].(int) <= 0 {
			thisState = -1
			errorCode = 1
		}
		if betOrdersData["count"].(int) > 0 {
			betOrders["BO"] = make(map[string]interface{})

			for iBO1 := 0; iBO1 < betOrdersData["count"].(int); iBO1++ {
				iBO := iBO1 % betOrdersData["count"].(int)
				this_BO_Mode := betOrdersData["result"].([]interface{})[iBO].(map[string]string)["BO_Mode"]
				this_BO_LotteryPlayGroup := betOrdersData["result"].([]interface{})[iBO].(map[string]string)["BO_LotteryPlayGroup"]
				this_BO_LotteryPlay := betOrdersData["result"].([]interface{})[iBO].(map[string]string)["BO_LotteryPlay"]

				this_BO_LotteryContent := betOrdersData["result"].([]interface{})[iBO].(map[string]string)["BO_LotteryContent"]
				this_BO_Unit := betOrdersData["result"].([]interface{})[iBO].(map[string]string)["BO_Unit"]
				this_BO_Multiple := betOrdersData["result"].([]interface{})[iBO].(map[string]string)["BO_Multiple"]
				this_BO_BetCount := betOrdersData["result"].([]interface{})[iBO].(map[string]string)["BO_BetCount"]

				this_BO_Price, _ := strconv.ParseFloat(betOrdersData["result"].([]interface{})[iBO].(map[string]string)["BO_Price"], 64)
				this_BO_RealPrice, _ := strconv.ParseFloat(betOrdersData["result"].([]interface{})[iBO].(map[string]string)["BO_RealPrice"], 64)
				// this_BO_PriceNew, _ := strconv.ParseFloat(betOrdersData["result"].([]interface{})[iBO].(map[string]string)["BO_Price"], 64)

				this_BO_Odds1, _ := strconv.ParseFloat(betOrdersData["result"].([]interface{})[iBO].(map[string]string)["BO_Odds"], 64)
				this_BO_Odds2, _ := strconv.ParseFloat(betOrdersData["result"].([]interface{})[iBO].(map[string]string)["BO_Odds2"], 64)
				this_BO_Odds3, _ := strconv.ParseFloat(betOrdersData["result"].([]interface{})[iBO].(map[string]string)["BO_Odds3"], 64)

				this_BO_Winnings1, _ := strconv.ParseFloat(betOrdersData["result"].([]interface{})[iBO].(map[string]string)["BO_Winnings"], 64)
				this_BO_Winnings2, _ := strconv.ParseFloat(betOrdersData["result"].([]interface{})[iBO].(map[string]string)["BO_Winnings2"], 64)
				this_BO_Winnings3, _ := strconv.ParseFloat(betOrdersData["result"].([]interface{})[iBO].(map[string]string)["BO_Winnings3"], 64)
				this_BO_WinningLimit, _ := strconv.ParseFloat(betOrdersData["result"].([]interface{})[iBO].(map[string]string)["BO_WinningLimit"], 64)
				// fmt.Println("this_BO_Mode  ", this_BO_Mode)
				data := make(map[string]interface{})
				if this_BO_Mode == "1" {
					//帶入注單內容
					betOrderListDetail := make(map[string]interface{})
					betOrderListDetail["LPG"] = this_BO_LotteryPlayGroup
					betOrderListDetail["LP"] = this_BO_LotteryPlay
					betOrderListDetail["BC"] = this_BO_LotteryContent
					betOrderListDetail["BU"] = this_BO_Unit
					betOrderListDetail["BM"] = this_BO_Multiple
					betOrderListDetail["BP"] = this_BO_Price

					betOrderListDetail["RP"] = this_BO_RealPrice

					betOrderListDetail["BBC"] = this_BO_BetCount
					betOrderListDetail["BO1"] = this_BO_Odds1
					betOrderListDetail["BO2"] = this_BO_Odds2
					betOrderListDetail["BO3"] = this_BO_Odds3

					betOrderListDetail["WIN1"] = this_BO_Winnings1
					betOrderListDetail["WIN2"] = this_BO_Winnings2
					betOrderListDetail["WIN3"] = this_BO_Winnings3
					betOrderListDetail["WinningLimit"] = this_BO_WinningLimit

					if _, ok := betOrders["BO"].(map[string]interface{})["1"]; ok { //已有資料
						data = betOrders["BO"].(map[string]interface{})["1"].(map[string]interface{})
						data["betOrderList"] = append(data["betOrderList"].([]interface{}), betOrderListDetail)
						betOrders["BO"].(map[string]interface{})["1"] = data
					} else {

						data["LPM"] = "1"
						data["betOrderList"] = make([]interface{}, 0)
						data["betOrderList"] = append(data["betOrderList"].([]interface{}), betOrderListDetail)

						betOrders["BO"].(map[string]interface{})["1"] = data

					}
					lotteryPlayGroup, _ := strconv.Atoi(this_BO_LotteryPlayGroup)
					lotteryPlay, _ := strconv.Atoi(this_BO_LotteryPlay)

					switch thisLottteryTypeGroup {
					//北京PK10  	//幸运飞艇
					case 1, 3:
						if lotteryPlayGroup != 1 {
							isOnlyTwoWayBet = false
						} else if lotteryPlayGroup == 1 && lotteryPlay == 1 {
							isOnlyTwoWayBet = false
						}
					//时时彩
					case 2:
						if lotteryPlayGroup != 1 {
							isOnlyTwoWayBet = false
						}
					//六合彩
					case 5:
						isOnlyTwoWayBet = false
					//快三
					case 6:
						if lotteryPlayGroup != 1 {
							isOnlyTwoWayBet = false
						} else if lotteryPlayGroup == 1 && lotteryPlay == 1 && this_BO_Odds1 != 1.995 {
							isOnlyTwoWayBet = false
						}
					default:

						if lotteryPlayGroup != 1 || thisLottteryTypeGroup == 4 { //不為PC蛋蛋
							isOnlyTwoWayBet = false
						}

					}

				} else if this_BO_Mode == "2" {
					isOnlyTwoWayBet = false
					data := make(map[string]interface{})

					//帶入注單內容
					betOrderListDetail := make(map[string]interface{})
					betOrderListDetail["LPG"] = this_BO_LotteryPlayGroup
					betOrderListDetail["LP"] = this_BO_LotteryPlay
					betOrderListDetail["BC"] = this_BO_LotteryContent
					betOrderListDetail["BU"] = this_BO_Unit
					betOrderListDetail["BM"] = this_BO_Multiple
					betOrderListDetail["BP"] = this_BO_Price

					betOrderListDetail["RP"] = this_BO_RealPrice
					betOrderListDetail["BBC"] = this_BO_BetCount
					betOrderListDetail["BO1"] = this_BO_Odds1
					betOrderListDetail["BO2"] = this_BO_Odds2
					betOrderListDetail["BO3"] = this_BO_Odds3

					betOrderListDetail["WIN1"] = this_BO_Winnings1
					betOrderListDetail["WIN2"] = this_BO_Winnings2
					betOrderListDetail["WIN3"] = this_BO_Winnings3
					betOrderListDetail["WinningLimit"] = this_BO_WinningLimit

					if _, ok := betOrders["BO"].(map[string]interface{})["2"]; ok { //已有資料
						data = betOrders["BO"].(map[string]interface{})["2"].(map[string]interface{})
						data["betOrderList"] = append(data["betOrderList"].([]interface{}), betOrderListDetail)
						betOrders["BO"].(map[string]interface{})["2"] = data
					} else {
						data["LPM"] = "2"
						data["betOrderList"] = make([]interface{}, 0)
						data["betOrderList"] = append(data["betOrderList"].([]interface{}), betOrderListDetail)
						betOrders["BO"].(map[string]interface{})["2"] = data
						// fmt.Println("	data[betOrderList]  ", data["betOrderList"])
					}
				}
			}

		}

		if isOnlyTwoWayBet {
			public.Println(fmt.Sprint("二面 只跑一次 阿 "))
		}
		var thisOpenResult1 []string
		var thisOpenResult2 string
		var hashCode string

		// public.Println(fmt.Sprint("二面 blockChaintime  ", blockChaintime))
		var availableChain []interface{} //從Redis拉的區塊鍊資料
		if blockChaintime != 0 {
			BlockTime := blockChaintime

			// public.Println(fmt.Sprint("二面 BlockTime  ", BlockTime))

			thisCount := 0
			for {
				t := time.Unix(BlockTime+int64(thisCount), 0)
				BlockChainKey := fmt.Sprintf("%d-%02d-%02d:%d", t.Year(), t.Month(), t.Day(), BlockTime+int64(thisCount))
				//fmt.Printf("BlockChainKey %s \n", BlockChainKey)
				chainList := redisConnect.GetBlockChain(BlockChainKey)
				availableChain = append(availableChain, chainList...)
				if len(availableChain) == 0 && thisCount == 0 {

					// public.Println(fmt.Sprint("QQQQ  "))
					return make(map[string]interface{})
				}

				//fmt.Printf("chainList %s \n", chainList)

				thisCount++
				//撈超過10秒區間或者區塊鍊筆數超過100筆
				if thisCount >= 10 || len(availableChain) >= 100 {

					break
				}

			}
			//public.Println(fmt.Sprint("availableChain -------> availableChain  ", availableChain))
			public.Println(fmt.Sprint("availableChain -------> availableChain  ", len(availableChain)))

			//===============================重新定義總筆數與權重資料長度===============================
			allResultCount = len(availableChain)
			resultRandCount = float64(allResultCount)       //總權重
			resultArr = make([]interface{}, allResultCount) // 存各組開獎結果 輸贏資料
			//===============================重新定義總筆數與權重資料長度===============================
		}
		// else {
		// 	return make(map[string]interface{})
		// }

		if (betOrders != nil) && thisState != -1 {
			//  有其他組合則跑完所有開獎結果
			for resultCount := 0; (isOnlyTwoWayBet && resultCount == 0) || (!isOnlyTwoWayBet && resultCount < allResultCount); resultCount++ {

				AllPrize = 0     //總贏錢
				AllRealPrice = 0 //總下注
				AllBOBetCount = 0

				thisOpenResult = ""
				hashCode = ""

				public.Println(fmt.Sprint("availableChain -------> resultCount  ", resultCount))
				//================================開獎號區START=======================
				if blockChaintime != 0 { //區塊鍊開獎

					hashCode = availableChain[resultCount].(map[string]interface{})["hash"].(string)
					thisOpenResult = GetHashCodeResult(thisLottteryTypeGroup, hashCode)

					if thisLottteryTypeGroup == 4 || thisLottteryTypeGroup == 5 {
						temp := strings.Split(thisOpenResult, ";")[0]
						temp1 := strings.Split(thisOpenResult, ";")[1]

						thisOpenResult1 = strings.Split(temp, ",")
						thisOpenResult2 = temp1

					} else {
						thisOpenResult1 = strings.Split(thisOpenResult, ",")
					}

				} else { //自開彩開獎

					//快3 隨機開
					if isOnlyTwoWayBet && thisLottteryTypeGroup == 6 {
						randResult := GetRandResult(true, "6")
						thisOpenResult = strings.Join(randResult["thisOpenResult"].([]string), ",")
					} else {
						thisOpenResult = getResultNum(thisLottteryTypeGroup, resultCount)
					}

					if thisLottteryTypeGroup == 4 || thisLottteryTypeGroup == 5 {
						temp := strings.Split(thisOpenResult, ";")[0]
						temp1 := strings.Split(thisOpenResult, ";")[1]

						thisOpenResult1 = strings.Split(temp, ",")
						thisOpenResult2 = temp1

					} else {
						thisOpenResult1 = strings.Split(thisOpenResult, ",")
					}

				}

				//================================開獎號區END=======================
				//	thisOpenResult1 = strings.Split(thisOpenResult, ",")
				//從規則參數中，取得對獎結果記錄格式 (從 PlayRule 取得空預設值，或從資料庫取出目前值)

				thisFullResult1 := thisLotteryConfig["LTR_ResultFormat"].(map[string]interface{})
				//從規則參數中，取得對獎所需的各項參數
				thisConfig := thisLotteryConfig["LTR_Config"].(map[string]interface{})

				thisFullResult := config.ReplaceRealResult(thisLottteryTypeGroup, thisLotteryType, thisOpenResult1, thisOpenResult2, nil, thisFullResult1, thisConfig)

				for _, v := range betOrders["BO"].(map[string]interface{}) {
					data := v.(map[string]interface{})
					thisLotteryPlayMode := 0
					if data["LPM"].(string) != "" {
						if lotteryPlayMode, err := strconv.Atoi(data["LPM"].(string)); err == nil {
							thisLotteryPlayMode = lotteryPlayMode
						} else {
							thisState = -1
							errorCode = 1
						}
					} else {
						thisLotteryPlayMode = 0
						thisState = -1
						errorCode = 1

					}

					if data["betOrderList"] != nil {
						for _, element := range data["betOrderList"].([]interface{}) {
							_element := element.(map[string]interface{})

							thisPrize = 0.0
							thisDrawn = 0

							thisLotteryPlayGroup := 0
							thisLotteryPlay := 0
							// thisOpenResult2 := []int{}
							var thisBOContent interface{}

							thisBOPrice := 0.0
							thisBORealPrice := 0.0

							thisBOUnit := 0
							thisBOMultiple := 0.0
							thisBOOdds1 := 0.0
							thisBOOdds2 := 0.0
							thisBOOdds3 := 0.0

							if _element["BC"].(interface{}) != "" {
								thisBOContent = _element["BC"]
							} else {

								thisState = -1
								errorCode = 1

							}

							if _element["LPG"].(string) != "" {
								if lotteryPlayGroup, err := strconv.Atoi(_element["LPG"].(string)); err == nil {
									thisLotteryPlayGroup = lotteryPlayGroup
								} else {
									thisState = -1
									errorCode = 1
								}
							} else {
								thisLotteryPlayGroup = 0
								thisState = -1
								errorCode = 1
							}

							if _element["LP"].(string) != "" {
								if lotteryPlay, err := strconv.Atoi(_element["LP"].(string)); err == nil {
									thisLotteryPlay = lotteryPlay

								} else {
									thisState = -1
									errorCode = 1
								}
							} else {
								thisBOPrice = 0
								thisState = -1
								errorCode = 1
							}

							if _element["BP"].(float64) != 0 {
								thisBOPrice = _element["BP"].(float64)

							} else {
								thisBOPrice = 0
								thisState = -1
								errorCode = 1
							}

							if _element["RP"].(float64) != 0 {
								thisBORealPrice = _element["RP"].(float64)

							} else {
								thisBORealPrice = 0
								thisState = -1
								errorCode = 1
							}

							if _element["BU"].(string) != "" {
								if BOUnit, err := strconv.Atoi(_element["BU"].(string)); err == nil {
									thisBOUnit = BOUnit
								} else {
									thisState = -1
									errorCode = 1
								}
							} else {
								thisBOUnit = 0
								thisState = -1
								errorCode = 1
							}

							if _element["BM"].(string) != "" {
								if BOMultiple, err := strconv.ParseFloat(_element["BM"].(string), 64); err == nil {
									thisBOMultiple = BOMultiple
								} else {
									thisState = -1
									errorCode = 1
								}
							} else {
								thisBOMultiple = 0
								thisState = -1
								errorCode = 1
							}

							thisBOOdds1 = _element["BO1"].(float64)
							thisBOOdds2 = _element["BO2"].(float64)
							thisBOOdds3 = _element["BO3"].(float64)

							thisBOWinnings1 := _element["WIN1"].(float64)
							thisBOWinnings2 := _element["WIN2"].(float64)
							thisBOWinnings3 := _element["WIN3"].(float64)

							thisWinningLimit := _element["WinningLimit"].(float64)

							thisBOContent = config.ReplaceRealContent(thisLottteryTypeGroup, thisLotteryType, "Test_Issue", thisLotteryConfig, thisLotteryPlayMode, thisLotteryPlayGroup, thisLotteryPlay, thisBOContent)
							newContent := make(map[string]interface{})
							newContent["1"] = make(map[string]interface{})
							newContent["1"] = thisBOContent
							//注數計算 直接從DB拉
							thisBOBetCount, _ := strconv.Atoi(_element["BBC"].(string))
							AllBOBetCount += thisBOBetCount
							thisBOPriceNew := thisBOPrice
							//對獎
							checkWinning := betcollect.CheckWinnings(thisLottteryTypeGroup, thisLotteryType, thisLotteryPlayMode, thisLotteryPlayGroup, thisLotteryPlay, thisBOContent, thisFullResult, thisConfig)
							//單位換算

							if thisBOUnit == 1 {
								thisBOPrice = thisBOPrice / 100
								thisBOPriceNew = thisBOPriceNew / 100
							} else if thisBOUnit == 2 {
								thisBOPrice = thisBOPrice / 10
								thisBOPriceNew = thisBOPriceNew / 10
							}

							AllRealPrice += thisBORealPrice
							//取得是否中獎 (0:沒中獎|1:中獎|2:和局退本金)

							thisState = checkWinning["Message"].(int)

							//取得中獎的注數 (0:沒中獎)
							thisRatioList := checkWinning["Ratio"].([]float64)
							thisRatioCalc := len(checkWinning["Ratio"].([]float64))

							thisOddsList := []float64{thisBOOdds1, thisBOOdds2, thisBOOdds3}
							thisBOWinList := []float64{thisBOWinnings1, thisBOWinnings2, thisBOWinnings3}

							switch thisState {
							//中獎
							case 1:
								{
									for iRatio := 0; iRatio < thisRatioCalc; iRatio++ {
										if thisRatioList[iRatio] > 0 {
											thisRatio = thisRatioList[iRatio]
											//依判斷結果轉換為正確的賠率
											if thisOddsList[iRatio] > 0 {
												thisOdds = thisOddsList[iRatio]
											} else {
												thisOdds = thisOddsList[0]
											}

											//依判斷結果轉換為正確的預設單注獎金
											if thisBOWinList[iRatio] > 0 {
												thisWinnings = thisBOWinList[iRatio]
											} else {
												thisWinnings = thisBOWinList[0]
											}
											//計算獎金
											if thisLotteryPlayMode == 2 && thisWinnings > 0 {
												thisPrize = thisPrize + thisWinnings*thisRatio
											} else {
												thisPrize = thisPrize + thisBOPriceNew*thisBOMultiple*thisOdds*thisRatio
											}

											//檢查獎金是否超過最大獎金限制  未做

											if thisWinningLimit != 0 && thisWinningLimit < thisPrize {
												thisPrize = thisWinningLimit
											}
										}
									}

								}
								//和局退本金
							case 2:
								{
									thisPrize = 0
									thisDrawn = thisBOPrice
									break
								}
								//2:和局退本金
							default:
								{
									thisPrize = 0
									thisDrawn = 0
									break
								}

							}
							AllPrize += thisPrize
							AllRealPrice -= thisDrawn
						}

					} else {
						errorCode = 1
					}

				}

				resultData := make(map[string]interface{})
				tempRatio := 0.0
				if AllRealPrice > 0.0 {
					tempRatio = AllPrize / AllRealPrice
				}

				resultData["targetRatio"] = tempRatio
				resultData["weights"] = 1.0
				resultData["allPrize"] = AllPrize
				resultData["allRealPrice"] = AllRealPrice

				resultData["thisOpenResult"] = thisOpenResult1
				resultData["chainCode"] = hashCode
				if thisLottteryTypeGroup == 5 { //六合彩
					resultData["thisOpenResult"] = append(resultData["thisOpenResult"].([]string), thisOpenResult2)
				}
				resultArr[resultCount] = resultData

				//計算有幾組大於 目標RTP
				if tempRatio > targetRatio {
					availableList = append(availableList, resultCount)
				} else if tempRatio <= targetRatio && tempRatio != 0 {
					availableList1 = append(availableList1, resultCount)
				}

			}

		}

		//計算所有組別獎號 基本權重後的平均RTP
		nowRTP := 0.0
		allWeightsCount := resultRandCount //總權重

		targetIndex := -1
		smallestRatioIndex := 0 //找出最小RTP獎號
		biggestRatioIndex := 0  //找出最大RTP獎號
		// fmt.Println(" 計算所有組別獎號 基本權重後的平均RTP resultArr ", len(resultArr))
		// fmt.Println(" 計算所有組別獎號 基本權重後的平均RTP resultArr ", resultArr)
		if isOnlyTwoWayBet == true {
			targetIndex = 0
		} else {
			for idx, item := range resultArr {
				// fmt.Println(" idx ", idx)
				// fmt.Println(" item ", item)
				if item.(map[string]interface{})["targetRatio"].(float64) > 0.0 {
					nowRTP += item.(map[string]interface{})["weights"].(float64) / float64(allWeightsCount) * item.(map[string]interface{})["targetRatio"].(float64)
				}

				if item.(map[string]interface{})["targetRatio"].(float64) > resultArr[biggestRatioIndex].(map[string]interface{})["targetRatio"].(float64) {
					biggestRatioIndex = idx
				} else if item.(map[string]interface{})["targetRatio"].(float64) < resultArr[smallestRatioIndex].(map[string]interface{})["targetRatio"].(float64) {
					smallestRatioIndex = idx
				}
			}

			if len(availableList) <= 0 { // 獎號RTP均小於目標RTP
				//找出最大RTP獎號
				public.Println(fmt.Sprint("獎號RTP均小於目標RTP "))
				targetIndex = biggestRatioIndex
			} else if len(availableList) == allResultCount { // 獎號RTP均大於目標RTP
				public.Println(fmt.Sprint("獎號RTP均大於目標RTP "))

				//找出最小RTP獎號
				targetIndex = smallestRatioIndex
			} else { //有可調整 機率的獎號RTP
				cont := 0
				cont1 := 0
				index := len(availableList)

				resultDataIndex := 0
				resultDataIndex1 := 0
				thisWeights := 0.0

				for nowRTP > targetRatio+0.0003 || nowRTP < targetRatio-0.0003 { //目前RTP過大或過小
					cont++
					// resultDataIndex := rng.Intn(len(resultArr))
					// resultDataIndex1 = resultDataIndex
					if nowRTP > targetRatio+0.003 { //往下調整RTP
						// fmt.Println("往下調整RTP")
						cont1++
						resultDataIndex = availableList[rng.Intn(index)].(int)
						thisWeights = resultArr[resultDataIndex].(map[string]interface{})["weights"].(float64)
						num := (rng.Intn(10) + 10)

						resultArr[resultDataIndex].(map[string]interface{})["weights"] = resultArr[resultDataIndex].(map[string]interface{})["weights"].(float64) * (1 / float64(num))

						// fmt.Println("往下調整RTP, weights", resultArr[resultDataIndex].(map[string]interface{})["weights"], resultDataIndex)
						allWeightsCount = allWeightsCount - thisWeights + resultArr[resultDataIndex].(map[string]interface{})["weights"].(float64)
						//
						// fmt.Println("往下調整RTP, allWeightsCount", allWeightsCount)

					} else { //往上調整RTP
						//resultDataIndex1 = availableList1[rng.Intn(index)].(int)
						resultDataIndex1 = availableList[rng.Intn(index)].(int)
						thisWeights = resultArr[resultDataIndex1].(map[string]interface{})["weights"].(float64)
						num := 1 + float64(rng.Intn(7)+1)/10
						resultArr[resultDataIndex1].(map[string]interface{})["weights"] = resultArr[resultDataIndex1].(map[string]interface{})["weights"].(float64) * num
						allWeightsCount = allWeightsCount - thisWeights + resultArr[resultDataIndex1].(map[string]interface{})["weights"].(float64)
					}

					nowRTP = 0
					for _, item := range resultArr {
						nowRTP += item.(map[string]interface{})["weights"].(float64) / float64(allWeightsCount) * item.(map[string]interface{})["targetRatio"].(float64)
					}
				}
				public.Println(fmt.Sprint("nowRTP  ", nowRTP))

				targetIndex = rng.Intn(allResultCount)
				tempAmount := resultArr[targetIndex].(map[string]interface{})["allRealPrice"].(float64) + nowAmount
				tempBonus := resultArr[targetIndex].(map[string]interface{})["allPrize"].(float64) + nowBonus
				if nowRatio < 1 && tempBonus/tempAmount > targetRatio+0.04 { //超過太多重rand一次
					targetIndex = rng.Intn(allResultCount)
				}

			}
		}

		if errorCode != 0 {
			return make(map[string]interface{})
		}
		resultInfo := make(map[string]interface{})
		// fmt.Println(" item      ", resultArr)
		// fmt.Println(" allWeightsCount      ", allWeightsCount)
		if isOnlyTwoWayBet {
			nowRTP += resultArr[0].(map[string]interface{})["weights"].(float64) / float64(allWeightsCount) * resultArr[0].(map[string]interface{})["targetRatio"].(float64)
		} else {
			for _, item := range resultArr {
				nowRTP += item.(map[string]interface{})["weights"].(float64) / float64(allWeightsCount) * item.(map[string]interface{})["targetRatio"].(float64)
			}
		}

		public.Println(fmt.Sprint("All_BO_BetCount ", AllBOBetCount))
		public.Println(fmt.Sprint("All_Prize ", fmt.Sprintf("%f", resultArr[targetIndex].(map[string]interface{})["allPrize"].(float64))))
		public.Println(fmt.Sprint("All_RealPrice ", AllRealPrice))
		fmt.Println("  targetIndex  ", targetIndex)
		fmt.Println("    ", resultArr[targetIndex])
		resultInfo["chainCode"] = resultArr[targetIndex].(map[string]interface{})["chainCode"].(string)
		resultInfo["amount"] = resultArr[targetIndex].(map[string]interface{})["allRealPrice"].(float64)
		resultInfo["bonus"] = resultArr[targetIndex].(map[string]interface{})["allPrize"].(float64)
		resultInfo["thisOpenResult"] = resultArr[targetIndex].(map[string]interface{})["thisOpenResult"].([]string)
		return resultInfo

	}

	// return true
	return make(map[string]interface{})
}

// func findDataIndex(List []interface{}, status string, targetRatio float64) []interface{} {
// 	rng := rand.New(mt19937.New())
// 	rng.Seed(time.Now().UnixNano())
// 	var availList []interface{}

// 	for idx, item := range List {

// 		if status == "up" {
// 			if targetRatio >= item.(map[string]interface{})["targetRatio"].(float64) {
// 				item.(map[string]interface{})["index"] = idx
// 				availList = append(availList, item)
// 			}
// 		} else if status == "down" {
// 			if targetRatio <= item.(map[string]interface{})["targetRatio"].(float64) {
// 				item.(map[string]interface{})["index"] = idx
// 				availList = append(availList, item)
// 			}
// 		}
// 	}
// 	availListLen := len(availList)
// 	if availListLen == 0 {

// 		index := len(List)
// 		return rng.Intn(index)
// 	} else {
// 		return availList[rng.Intn(availListLen)].(map[string]interface{})["index"].(int)
// 	}

// }
func getPartOfHashCode(hashCode string, start int, length int) string {
	return hashCode[start : start+length]
}

func hexToBigInt(hex string) uint64 { //16轉10

	n, err := strconv.ParseUint(hex, 16, 64)
	if err != nil {
		panic(err)
	}

	nn := uint64(n)

	return nn
}
func GetHashCodeResult(LottteryTypeGroup int, hashCode string) string {

	newHashCode := ""
	thisOpenResult := "" //獎號
	var n uint64
	var nn float64
	switch LottteryTypeGroup {
	case 1, 3: //北京PK10 //幸运飞艇 // 1分PK10 // 3分PK10  PK10
		var tempNumArray []string
		var resultArray []string
		for i := 0; i < 10; i++ {
			tempNumArray = append(tempNumArray, fmt.Sprintf("%d", i))
		}

		for i := 0; i < 9; i++ {
			newHashCode = getPartOfHashCode(hashCode, 0+i*6, 16)
			n = hexToBigInt(newHashCode)
			nn = float64(n) / math.Pow(2, 64)

			index := uint64(nn*10000000000) % uint64(10-i)
			resultArray = append(resultArray, tempNumArray[index])
			// if i != 10-1 {
			tempNumArray = append(tempNumArray[:index], tempNumArray[index+1:]...)
			// }
		}
		resultArray = append(resultArray, tempNumArray[0])
		thisOpenResult = strings.Join(resultArray, ",")

	case 2:
		newHashCode = getPartOfHashCode(hashCode, 0, 16)
		n = hexToBigInt(newHashCode)
		nn := float64(n) / math.Pow(2, 64)
		thisOpenResult = strings.Join(strings.Split(fmt.Sprintf("%5.0F", nn*100000), ""), ",")

	case 4: //PC蛋蛋 // 1分 PC蛋蛋 // 3分 PC蛋蛋
		var resultArray []int
		var tempNumArray []int
		for i := 1; i <= 80; i++ {
			tempNumArray = append(tempNumArray, i)
		}
		for i := 0; i < 20; i++ {
			shift := 0
			if i%2 == 1 {
				shift = 3 + ((i-1)/2)*5
			} else {
				shift = 0 + ((i)/2)*5
			}
			newHashCode = getPartOfHashCode(hashCode, shift, 16)
			n = hexToBigInt(newHashCode)
			nn = float64(n) / math.Pow(2, 64)

			index := uint64(nn*1000000000) % uint64(80-i)
			resultArray = append(resultArray, tempNumArray[index])
			tempNumArray = append(tempNumArray[:index], tempNumArray[index+1:]...)
		}

		thisOpenResultPCEgg := [3]int{0, 0, 0}

		for i := 0; i < 3; i++ {
			for j := 0; j < 6; j++ {
				thisOpenResultPCEgg[i] += resultArray[6*i+j]
			}
			thisOpenResultPCEgg[i] = thisOpenResultPCEgg[i] % 10
		}

		thisOpenResult = strconv.Itoa(thisOpenResultPCEgg[0]) + "," + strconv.Itoa(thisOpenResultPCEgg[1]) + "," + strconv.Itoa(thisOpenResultPCEgg[2]) + ";" + fmt.Sprintf("%d", thisOpenResultPCEgg[0]+thisOpenResultPCEgg[1]+thisOpenResultPCEgg[2])
	case 5: //六合彩 // 1分六合彩 // 5分六合彩
		var resultArray []string
		var tempNumArray []string
		for i := 1; i <= 49; i++ {
			tempNumArray = append(tempNumArray, fmt.Sprintf("%d", i))
		}

		for i := 0; i < 7; i++ {
			newHashCode = getPartOfHashCode(hashCode, 0+i*8, 16)
			n = hexToBigInt(newHashCode)
			nn = float64(n) / math.Pow(2, 64)

			index := uint64(nn*1000000000) % uint64(49-i)

			resultArray = append(resultArray, tempNumArray[index])
			tempNumArray = append(tempNumArray[:index], tempNumArray[index+1:]...)
		}

		special := resultArray[len(resultArray)-1]
		resultArray = resultArray[:len(resultArray)-1] // Truncate
		thisOpenResult = strings.Join(resultArray, ",")

		thisOpenResult = thisOpenResult + ";" + special

	case 6: //快三
		var resultArray []string
		for i := 0; i < 3; i++ {
			newHashCode = getPartOfHashCode(hashCode, 0+i*24, 16)

			n = hexToBigInt(newHashCode)
			nn = float64(n) / math.Pow(2, 64)

			num := uint64(nn*1000000000)%uint64(6) + 1
			resultArray = append(resultArray, strconv.Itoa(int(num)))
		}
		thisOpenResult = strings.Join(resultArray, ",")

	default:

	}
	public.Println("thisOpenResult " + thisOpenResult)
	return thisOpenResult

}
func getResultNum(LottteryTypeGroup int, count int) string {
	// public.Println(fmt.Sprint("getResultNum ", 1))
	rng := rand.New(mt19937.New())
	rng.Seed(time.Now().UnixNano())
	thisOpenResult := ""
	drawNum := 5

	if LottteryTypeGroup == 2 {
		for i := 0; i < drawNum; i++ {
			//number := strconv.Itoa(rand.Intn(10)) //亂數產生
			number := ""

			switch LottteryTypeGroup {
			//时时彩
			case 2:
				number = strconv.Itoa(rng.Intn(10))
			default:
				number = strconv.Itoa(rng.Intn(10))
			}
			if i == (drawNum - 1) {
				thisOpenResult += number
			} else {
				thisOpenResult += (number + ",") //資料整合
			}
		}

	} else if LottteryTypeGroup == 4 { //PC蛋蛋
		drawNum = 20
		var arr []int
		var thisOpenResultInt []int

		for i := 0; i < 80; i++ {
			arr = append(arr, i+1)
		}

		//亂數產生數字
		for i := 0; i < drawNum; i++ {
			index := rng.Intn(len(arr))
			number := arr[index]
			arr = append(arr[:index], arr[index+1:]...)
			thisOpenResultInt = append(thisOpenResultInt, number)
		}
		//排序
		sort.Ints(thisOpenResultInt)
		thisOpenResultPCEgg := [4]int{0, 0, 0, 0}

		for i := 0; i < 3; i++ {
			for j := 0; j < 6; j++ {
				thisOpenResultPCEgg[i] += thisOpenResultInt[6*i+j]
			}
			thisOpenResultPCEgg[i] = thisOpenResultPCEgg[i] % 10
		}

		thisOpenResultPCEgg[3] = thisOpenResultPCEgg[0] + thisOpenResultPCEgg[1] + thisOpenResultPCEgg[2]

		//轉成輸出格式
		for i, v := range thisOpenResultPCEgg {
			if len(thisOpenResult) > 0 {
				if i == len(thisOpenResultPCEgg)-1 {
					thisOpenResult += ";"
				} else {
					thisOpenResult += ","
				}

			}
			thisOpenResult += strconv.Itoa(v)
		}
		//public.Println(fmt.Sprint("thisOpenResult ", thisOpenResult))

	} else if LottteryTypeGroup == 5 { //六合彩
		drawNum = 7
		var arr []int
		for i := 0; i < 49; i++ {
			arr = append(arr, i+1)
		}

		for i := 0; i < drawNum; i++ {
			index := rng.Intn(len(arr))
			number := arr[index]

			arr = append(arr[:index], arr[index+1:]...)

			if i == (drawNum - 1) {
				thisOpenResult += strconv.Itoa(number)
			} else if i == (drawNum - 2) {
				thisOpenResult += (strconv.Itoa(number) + ";") //資料整合
			} else {
				thisOpenResult += (strconv.Itoa(number) + ",") //資料整合
			}
		}
	} else if LottteryTypeGroup == 6 { //快三
		cunt := count + 1
		x1 := math.Floor(float64(cunt-1)/36) + 1.0
		x2 := math.Floor(float64((cunt-1)%36)/6.0) + 1.0
		x3 := ((cunt - 1) % 36 % 6) + 1

		thisOpenResult = strconv.Itoa(int(x1)) + "," + strconv.Itoa(int(x2)) + "," + strconv.Itoa(int(x3))

	} else if LottteryTypeGroup == 1 || LottteryTypeGroup == 3 {
		maxNum := 10
		drawNum := 10
		var Lotto []int
		for i := 0; i < maxNum; i++ {
			Lotto = append(Lotto, i+1)
		}

		for i := 0; i < drawNum; i++ {
			pos := rand.Intn(maxNum)
			t := Lotto[i]
			Lotto[i] = Lotto[pos]
			Lotto[pos] = t
		}

		for i := 0; i < drawNum; i++ {
			num := ""
			if Lotto[i] == 10 {
				num = strconv.Itoa(Lotto[i])
			} else {
				num = "0" + strconv.Itoa(Lotto[i])
			}

			if i == (drawNum - 1) {
				thisOpenResult += num
			} else {
				thisOpenResult += (num + ",") //資料整合
			}

		}

	}
	return thisOpenResult

}

func GetChainRandResult(thisStatusData bool, thisLotteryTypeGroupData string, blockChaintime int64) map[string]interface{} {
	rng := rand.New(mt19937.New())
	rng.Seed(time.Now().UnixNano())
	LottteryTypeGroup, _ := strconv.Atoi(thisLotteryTypeGroupData)
	thisOpenResult := ""
	var thisOpenResult1 []string

	thisOpenResult = getResultNum(LottteryTypeGroup, 0)
	var availableChain []interface{} //從Redis拉的區塊鍊資料

	if blockChaintime != 0 {
		BlockTime := blockChaintime

		thisCount := 0
		for {
			t := time.Unix(BlockTime+int64(thisCount), 0)
			BlockChainKey := fmt.Sprintf("%d-%02d-%02d:%d", t.Year(), t.Month(), t.Day(), BlockTime+int64(thisCount))
			chainList := redisConnect.GetBlockChain(BlockChainKey)
			availableChain = append(availableChain, chainList...)
			thisCount++
			//撈超過10秒區間或者區塊鍊筆數超過100筆
			if thisCount >= 10 || len(availableChain) >= 1 {
				break
			}
		}

		if len(availableChain) == 0 {
			return make(map[string]interface{})
		}

	} else {
		return make(map[string]interface{})
	}

	hashCode := availableChain[0].(map[string]interface{})["hash"].(string)
	thisOpenResult = GetHashCodeResult(LottteryTypeGroup, hashCode)
	if LottteryTypeGroup == 4 {
		temp := strings.Split(thisOpenResult, ";")[0]
		thisOpenResult1 = strings.Split(temp, ",")
	} else if LottteryTypeGroup == 5 {
		temp := strings.Split(thisOpenResult, ";")[0]
		temp1 := strings.Split(thisOpenResult, ";")[1]

		thisOpenResult1 = strings.Split(temp, ",")
		thisOpenResult1 = append(thisOpenResult1, temp1)

	} else {
		thisOpenResult1 = strings.Split(thisOpenResult, ",")
	}

	public.Println(fmt.Sprint("thisOpenResult1 ", thisOpenResult1))

	resultInfo := make(map[string]interface{})
	resultInfo["amount"] = 0
	resultInfo["bonus"] = 0
	resultInfo["thisOpenResult"] = thisOpenResult1
	resultInfo["chainCode"] = hashCode
	return resultInfo

}
func GetRandResult(thisStatusData bool, thisLotteryTypeGroupData string) map[string]interface{} {
	rng := rand.New(mt19937.New())
	rng.Seed(time.Now().UnixNano())
	LottteryTypeGroup, _ := strconv.Atoi(thisLotteryTypeGroupData)
	thisOpenResult := ""
	thisOpenResult = getResultNum(LottteryTypeGroup, 0)

	var thisOpenResult1 []string
	if LottteryTypeGroup == 6 {
		x1 := rng.Intn(6) + 1
		x2 := rng.Intn(6) + 1
		x3 := rng.Intn(6) + 1

		thisOpenResult = strconv.Itoa(int(x1)) + "," + strconv.Itoa(int(x2)) + "," + strconv.Itoa(int(x3))
		thisOpenResult1 = strings.Split(thisOpenResult, ",")
		fmt.Println("  thisOpenResult1  ", thisOpenResult1)

	}
	if LottteryTypeGroup == 4 {
		temp := strings.Split(thisOpenResult, ";")[0]
		thisOpenResult1 = strings.Split(temp, ",")
	} else if LottteryTypeGroup == 5 {
		temp := strings.Split(thisOpenResult, ";")[0]
		temp1 := strings.Split(thisOpenResult, ";")[1]

		thisOpenResult1 = strings.Split(temp, ",")
		thisOpenResult1 = append(thisOpenResult1, temp1)

	} else {
		thisOpenResult1 = strings.Split(thisOpenResult, ",")
	}

	resultInfo := make(map[string]interface{})
	resultInfo["amount"] = 0
	resultInfo["bonus"] = 0
	resultInfo["thisOpenResult"] = thisOpenResult1
	return resultInfo
}

func print_map(m map[string]interface{}) {
	for k, v := range m {
		switch value := v.(type) {
		case nil:
			fmt.Println(k, "is nil", "null")
		case string:
			fmt.Println(k, "is string", value)
		case int:
			fmt.Println(k, "is int", value)
		case float64:
			fmt.Println(k, "is float64", value)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range value {
				fmt.Println(i, u)
			}
		case map[string]interface{}:
			fmt.Println(k, "is an map:")
			print_map(value)
		default:
			fmt.Println(k, "is unknown type", fmt.Sprintf("%T", v))
		}
	}
}
