package betcollect

import (
	"fmt"
	"getReslut/betCount"
	"reflect"
	"strconv"
)

func CheckWinnings(LotteryTypeGroup int,
	LotteryType int,
	LotteryMode int,
	LotteryPlayGroup int,
	LotteryPlay int,
	thisBOContent interface{},
	thisFullResult map[string]interface{},
	config map[string]interface{}) map[string]interface{} {
	// println(" replaceRealResult ")

	//初始化
	CWCheck := make(map[string]interface{})
	//對獎結果 (0：沒中獎|1:中獎|2:和局退本金)
	CWCheck["Message"] = 0
	//中獎金額倍數 (先預設為0)(簡單的說，就是嬴的注數…)
	CWCheck["Ratio"] = 0
	// //中獎賠率 (預設為1，即使用第一個)
	// CWCheck["Odds"] = 1;
	//執行狀態
	CWCheck["State"] = true
	// //執行訊息
	CWCheck["Info"] = ""

	switch LotteryTypeGroup {
	//北京PK10 	//幸运飞艇
	case 1, 3:
		switch LotteryType {
		case 1, 5, 31, 34, 35:
			// fmt.Println("====>InspectModulePKTen   ", thisBOContent)
			CWCheck = InspectModulePKTen(LotteryTypeGroup, LotteryType, LotteryMode, LotteryPlayGroup, LotteryPlay, thisBOContent, thisFullResult, config)
		default:
			CWCheck["State"] = false

		}
		break
	//时时彩
	case 2:
		//彩種 (因為有對應 TS_LotteryType.LT_Code，所以需要跟資料庫一起變動...ry)
		switch LotteryType {
		//重庆时时彩 //新疆时时彩 //天津时时彩 //精彩1分彩 //精彩3分彩 //精彩5分彩 //精彩秒秒彩
		case 2, 3, 4, 26, 27, 28, 29, 30: // 30 區塊鏈時時彩
			// fmt.Println("====>CWCheck  ", CWCheck)
			CWCheck = InspectModuleShiShi(LotteryTypeGroup, LotteryType, LotteryMode, LotteryPlayGroup, LotteryPlay, thisBOContent, thisFullResult, config)
		//彩種錯誤
		default:
			CWCheck["State"] = false
		}
		break
		//PC蛋蛋
	case 4:
		//彩種 (因為有對應 TS_LotteryType.LT_Code，所以需要跟資料庫一起變動...ry)
		switch LotteryType {
		case 6, 33, 40, 41: //PC蛋蛋 //區塊鏈 PC蛋蛋 // 1分 PC蛋蛋 // 3分 PC蛋蛋
			CWCheck = InspectModulePCEgg(LotteryTypeGroup, LotteryType, LotteryMode, LotteryPlayGroup, LotteryPlay, thisBOContent, thisFullResult, config)
		//彩種錯誤
		default:
			CWCheck["State"] = false
		}
		break
		//六合彩
	case 5:
		//彩種 (因為有對應 TS_LotteryType.LT_Code，所以需要跟資料庫一起變動...ry)
		switch LotteryType {
		case 7, 32, 38, 39: //六合彩 //區塊鏈六合彩 // 1分六合彩 // 5分六合彩
			CWCheck = InspectModuleMarkSixTrad(LotteryTypeGroup, LotteryType, LotteryMode, LotteryPlayGroup, LotteryPlay, thisBOContent, thisFullResult, config)
		//彩種錯誤
		default:
			CWCheck["State"] = false
		}
		break
		//快三
	case 6:
		//彩種 (因為有對應 TS_LotteryType.LT_Code，所以需要跟資料庫一起變動...ry)
		switch LotteryType {
		//江苏快三 //安徽快三 //广西快三// 1分快三 //3分快三 //區塊鍊快三
		case 8, 9, 10, 36, 37, 42:
			CWCheck = InspectModuleKuaiThree(LotteryTypeGroup, LotteryType, LotteryMode, LotteryPlayGroup, LotteryPlay, thisBOContent, thisFullResult, config)
		//彩種錯誤
		default:
			CWCheck["State"] = false

		}
		break

	}
	if CWCheck["State"] == true {

	} else {
		CWCheck["Info"] = "'Lottery Type Error"

	}

	return CWCheck

}
func InspectModulePCEgg(LotteryTypeGroup int,
	LotteryType int,
	LotteryMode int,
	LotteryPlayGroup int,
	LotteryPlay int,
	thisBOContent interface{},
	thisFullResult map[string]interface{},
	config map[string]interface{}) map[string]interface{} {

	//初始化
	Inspect := make(map[string]interface{})
	//對獎結果 (0：沒中獎|1:中獎|2:和局退本金)
	Inspect["Message"] = 0
	//中獎金額倍數 (先預設為0)(簡單的說，就是嬴的注數…)
	Inspect["Ratio"] = []float64{0, 0, 0, 0}
	// //中獎賠率 (預設為1，即使用第一個)
	// CWCheck["Odds"] = 1;
	//執行狀態
	Inspect["State"] = true
	// //執行訊息
	Inspect["Info"] = ""
	//最大下注內容上限 (每一注單/每一玩法)(如果投注項目超過上限，超出的項目全部無視掉 (σ′▽‵)′▽‵)σ)
	ContentMax := int(config["LTR_ContentMax"].(map[string]interface{})[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].(float64))

	switch LotteryTypeGroup {
	//PC蛋蛋 (幸运28)
	case 4:
		{
			switch LotteryType { //PC蛋蛋 (幸运28)
			case 6, 33, 40, 41: //PC蛋蛋 //區塊鏈 PC蛋蛋 // 1分 PC蛋蛋 // 3分 PC蛋蛋

				switch LotteryMode {
				case 1: //传统模式 (信用模式)
					{
						//玩法群組 (因為有對應 TS_LotteryPlayGroup.LPG_Code，所以需要跟資料庫一起變動...ry)
						// public.Println(fmt.Sprint("LotteryPlayGroup ", LotteryPlayGroup))
						switch LotteryPlayGroup {
						//玩法群組 (混合)
						case 1:
							{
								//如果投注項目比規則中的項目少，就只比對投注的項目就好 (σﾟ∀ﾟ)σﾟ∀ﾟ)σ
								if len(thisBOContent.(map[string]interface{})) < ContentMax {
									ContentMax = len(thisBOContent.(map[string]interface{}))
								}
								for i := 0; i < ContentMax; i++ {
									// public.Println(fmt.Sprint("thisBOContent ", thisBOContent.(map[string]interface{})[strconv.Itoa(i)]))
									// public.Println(fmt.Sprint("thisFullResult ", thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)]))

									if betCount.In_array(thisBOContent.(map[string]interface{})[strconv.Itoa(i)], thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)]) {
										Inspect["Ratio"].([]float64)[0]++
										Inspect["Message"] = 1
									}

								}

							}
						//玩法群組 (特码)
						case 2:
							{
								//如果投注項目比規則中的項目少，就只比對投注的項目就好 (σﾟ∀ﾟ)σﾟ∀ﾟ)σ
								if len(thisBOContent.(map[string]interface{})) < ContentMax {
									ContentMax = len(thisBOContent.(map[string]interface{}))
								}
								for i := 0; i < ContentMax; i++ {
									// public.Println(fmt.Sprint("thisBOContent ", thisBOContent.(map[string]interface{})[strconv.Itoa(i)]))
									// public.Println(fmt.Sprint("thisFullResult ", thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)]))

									if thisBOContent.(map[string]interface{})[strconv.Itoa(i)] == thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] {
										Inspect["Ratio"].([]float64)[0]++
										Inspect["Message"] = 1
									}
								}
							}

						default:
							Inspect["State"] = false
							Inspect["Info"] = "Lottery Play Group Error"
							break
						}
					}
				}
			default:
				Inspect["State"] = false
				Inspect["Info"] = "Lottery Play Mode Group Error"
			}
		}

	default:
		Inspect["State"] = false
	}

	return Inspect
}
func InspectModuleKuaiThree(LotteryTypeGroup int,
	LotteryType int,
	LotteryMode int,
	LotteryPlayGroup int,
	LotteryPlay int,
	thisBOContent interface{},
	thisFullResult map[string]interface{},
	config map[string]interface{}) map[string]interface{} {

	//初始化
	Inspect := make(map[string]interface{})
	//對獎結果 (0：沒中獎|1:中獎|2:和局退本金)
	Inspect["Message"] = 0
	//中獎金額倍數 (先預設為0)(簡單的說，就是嬴的注數…)
	Inspect["Ratio"] = []float64{0, 0, 0, 0}
	// //中獎賠率 (預設為1，即使用第一個)
	// CWCheck["Odds"] = 1;
	//執行狀態
	Inspect["State"] = true
	// //執行訊息
	Inspect["Info"] = ""
	//最大下注內容上限 (每一注單/每一玩法)(如果投注項目超過上限，超出的項目全部無視掉 (σ′▽‵)′▽‵)σ)
	ContentMax := int(config["LTR_ContentMax"].(map[string]interface{})[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].(float64))

	switch LotteryTypeGroup {
	//快三
	case 6:
		{
			switch LotteryType { //江苏快三 //安徽快三 //广西快三
			//江苏快三 //安徽快三 //广西快三// 1分快三 //3分快三
			case 8, 9, 10, 36, 37, 42: //塊鍊快三

				switch LotteryMode {
				case 1: //传统模式 (信用模式)
					{
						//玩法群組 (因為有對應 TS_LotteryPlayGroup.LPG_Code，所以需要跟資料庫一起變動...ry)

						// fmt.Println("====LotteryPlayGroup  ", LotteryPlayGroup)
						switch LotteryPlayGroup {
						//玩法群組 (三军)
						case 1:
							{

								// fmt.Println("====三军   ")
								//如果投注項目比規則中的項目少，就只比對投注的項目就好 (σﾟ∀ﾟ)σﾟ∀ﾟ)σ
								if len(thisBOContent.(map[string]interface{})) < ContentMax {
									ContentMax = len(thisBOContent.(map[string]interface{}))
								}
								// fmt.Println("====三军   ContentMax", ContentMax)

								// fmt.Println("====三军   thisBOContent", thisBOContent)
								// fmt.Println("====三军   thisFullResult", thisFullResult)
								for i := 0; i < ContentMax; i++ {
									// fmt.Println("====三军   1", reflect.TypeOf(thisBOContent.(map[string]interface{})[strconv.Itoa(i)]).String())
									// fmt.Println("====三军   1", reflect.TypeOf(thisBOContent.(map[string]interface{})[strconv.Itoa(i)]).String() == "int")
									if reflect.TypeOf(thisBOContent.(map[string]interface{})[strconv.Itoa(i)]).String() == "int" {
										ResultTemp := betCount.Dup_count(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))
										ResultSorted := betCount.SortMapByValue(ResultTemp)
										for _, v := range ResultSorted {

											// fmt.Println("====三军 v.Key  ", v.Key)
											// fmt.Println("====三军 v.Value  ", v.Value)

											// fmt.Println("====三军    ", thisBOContent.(map[string]interface{})[strconv.Itoa(i)])
											if v.Key == thisBOContent.(map[string]interface{})[strconv.Itoa(i)] && v.Value == 1 {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											} else if v.Key == thisBOContent.(map[string]interface{})[strconv.Itoa(i)] && v.Value == 2 {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											} else if v.Key == thisBOContent.(map[string]interface{})[strconv.Itoa(i)] && v.Value == 3 {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											}
										}
									} else {

										// fmt.Println("====三军 v.Key  ", thisBOContent.(map[string]interface{})[strconv.Itoa(i)])
										// fmt.Println("====三军 v.Value  ", thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{}))
										// fmt.Println("====三军 v.Value  ", thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)])

										// fmt.Println("====三军 v.Value  ", betCount.In_array(thisBOContent.(map[string]interface{})[strconv.Itoa(i)], thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)]))
										// fmt.Println("====三军   thisBOContent.(map[string]interface{})[strconv.Itoa(i)]", thisBOContent.(map[string]interface{})[strconv.Itoa(i)])
										if betCount.In_array(thisBOContent.(map[string]interface{})[strconv.Itoa(i)], thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)]) {

											//fmt.Println(thisBOContent.(map[string]interface{})[strconv.Itoa(i)])
											Inspect["Ratio"].([]float64)[0]++
											Inspect["Message"] = 1
										}

									}
								}

							}
							//玩法群組 (围骰)
						case 2:
							{
								//如果投注項目比規則中的項目少，就只比對投注的項目就好 (σﾟ∀ﾟ)σﾟ∀ﾟ)σ

								// fmt.Println("====thisBOContent  ", thisBOContent)
								if len(thisBOContent.(map[string]interface{})) < ContentMax {
									ContentMax = len(thisBOContent.(map[string]interface{}))
								}
								for i := 0; i < ContentMax; i++ {

									// fmt.Println("====thisBOContent =>>> ", reflect.TypeOf(thisBOContent.(map[string]interface{})[strconv.Itoa(i)]))
									// fmt.Println("====thisBOContent =>>> ", thisBOContent.(map[string]interface{})[strconv.Itoa(i)])
									// fmt.Println("====thisBOContent =>>> ", reflect.TypeOf(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)]))
									// fmt.Println("====thisBOContent  ", thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)])

									// fmt.Println("====thisBOContent =>>> ", thisBOContent.(map[string]interface{})[strconv.Itoa(i)] == thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]interface{})[1])
									// fmt.Println("====In_array2  ", betCount.In_array(thisBOContent.(map[string]interface{})[strconv.Itoa(i)], thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]interface{})[1]))

									// fmt.Println("====In_array  kkk ", thisBOContent.(map[string]interface{})[strconv.Itoa(i)])

									// fmt.Println("====In_array  ccc", thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]interface{})[1].([]string))

									// for k := 0; k < len(thisBOContent.(map[string]interface{})[strconv.Itoa(i)].([]interface{})); k++ {

									// 	fmt.Println("====In_array    cccc ", thisBOContent.(map[string]interface{})[strconv.Itoa(i)].([]interface{})[k])

									// 	fmt.Println("====In_array    cccc ", len(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]interface{})))
									num1 := thisBOContent.(map[string]interface{})[strconv.Itoa(i)].([]interface{})

									numString := make([]string, 0)
									for k := 0; k < len(num1); k++ {
										numString = append(numString, fmt.Sprintf("%.0f", num1[k].(float64)))
									}
									thisBOContent.(map[string]interface{})[strconv.Itoa(i)] = numString
									FullResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]interface{})
									for s := 0; s < len(FullResult); s++ {
										if reflect.TypeOf(FullResult[s]).String() == "[]string" {
											num2 := FullResult[s].([]string)
											if numString[0] == num2[0] && numString[1] == num2[1] && numString[2] == num2[2] {
												// fmt.Println("numString=================================== ", numString[0])
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1

											}

										}
									}

									// }

									// if betCount.In_array(numString, thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)]) {

									// 	fmt.Println("numString====QQQ=============================== ", numString[0])
									// 	// Inspect["Ratio"].([]float64)[0]++
									// 	// Inspect["Message"] = 1
									// }
								}
							}
						//玩法群組 (点数)   還沒做
						case 3:
							{
								//如果投注項目比規則中的項目少，就只比對投注的項目就好 (σﾟ∀ﾟ)σﾟ∀ﾟ)σ
								if len(thisBOContent.(map[string]interface{})) < ContentMax {
									ContentMax = len(thisBOContent.(map[string]interface{}))
								}

								for i := 0; i < ContentMax; i++ {

									// fmt.Println("numString====aaa=============================== ", thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)])

									// fmt.Println("numString====aaa=============================== ", thisBOContent.(map[string]interface{})[strconv.Itoa(i)] == fmt.Sprintf("%v", thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].(int)))
									if thisBOContent.(map[string]interface{})[strconv.Itoa(i)] == fmt.Sprintf("%v", thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].(int)) {
										// fmt.Println(" 1045245====QQQ=============================== ")
										Inspect["Ratio"].([]float64)[0]++
										Inspect["Message"] = 1
									}
								}
							}
							//玩法群組 (长牌)
						case 4:
							{
								//如果投注項目比規則中的項目少，就只比對投注的項目就好 (σﾟ∀ﾟ)σﾟ∀ﾟ)σ
								if len(thisBOContent.(map[string]interface{})) < ContentMax {
									ContentMax = len(thisBOContent.(map[string]interface{}))
								}
								//使用迴圈對獎，防止注單中可能不只有一項
								tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)
								for i := 0; i < ContentMax; i++ {
									// fmt.Println("numString====tempResult=============================== ", tempResult)
									// fmt.Println("numString====tempResult=============================== ", thisBOContent.(map[string]interface{})[strconv.Itoa(i)].([]interface{})[0])
									// fmt.Println("numString====tempResult=============================== ", thisBOContent.(map[string]interface{})[strconv.Itoa(i)].([]interface{})[1])

									// fmt.Println("numString====tempResult=============================== ", betCount.In_array(thisBOContent.(map[string]interface{})[strconv.Itoa(i)].([]interface{})[0], tempResult))
									if betCount.In_array(fmt.Sprintf("%v", thisBOContent.(map[string]interface{})[strconv.Itoa(i)].([]interface{})[0]), tempResult) &&
										betCount.In_array(fmt.Sprintf("%v", thisBOContent.(map[string]interface{})[strconv.Itoa(i)].([]interface{})[1]), tempResult) {

										// fmt.Println(" 1045245====QQQ=============================== ")
										Inspect["Ratio"].([]float64)[0]++
										Inspect["Message"] = 1
									}
								}

							}
							//玩法群組 (短牌)
						case 5:
							{
								//如果投注項目比規則中的項目少，就只比對投注的項目就好 (σﾟ∀ﾟ)σﾟ∀ﾟ)σ
								if len(thisBOContent.(map[string]interface{})) < ContentMax {
									ContentMax = len(thisBOContent.(map[string]interface{}))
								}
								//使用迴圈對獎，防止注單中可能不只有一項
								for i := 0; i < ContentMax; i++ {
									tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)
									if betCount.In_array(fmt.Sprintf("%v", thisBOContent.(map[string]interface{})[strconv.Itoa(i)].([]interface{})[0]), tempResult) {

										// fmt.Println(" 1045245====aaaa=============================== ", fmt.Sprintf("%v", thisBOContent.(map[string]interface{})[strconv.Itoa(i)].([]interface{})[0]))
										ResultTemp := betCount.Dup_count(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))
										ResultSorted := betCount.SortMapByValue(ResultTemp)

										// fmt.Println(" 1045245====ResultSorted==== ", ResultSorted)
										for _, v := range ResultSorted {
											if v.Key == fmt.Sprintf("%v", thisBOContent.(map[string]interface{})[strconv.Itoa(i)].([]interface{})[0]) && v.Value >= 2 {

												// fmt.Println(" 1045245====qqqqqqqq==== ")
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											}
										}

									}
								}

							}
						default:
							Inspect["State"] = false
							Inspect["Info"] = "Lottery Play Group Error"
							break

						}
					}

				case 2: //官方模式
					{
						//玩法群組 (因為有對應 TS_LotteryPlayGroup.LPG_Code，所以需要跟資料庫一起變動...ry)
						switch LotteryPlayGroup {
						//玩法群組 (同号)
						case 1:
							{
								//玩法 (因為有對應 TS_LotteryPlay.LP_Code，所以需要跟資料庫一起變動...ry)
								switch LotteryPlay {

								//玩法 (三同号通选)
								case 1:
									{
										//依下注的項目列出全部的組合
										ContentAll := thisBOContent.(map[string]interface{})["1"].([]interface{})

										for i := 0; i < len(ContentAll); i++ {
											tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]interface{})
											// fmt.Println("====tempResult==== ", tempResult)
											if betCount.In_array(ContentAll[i].(string), tempResult) {
												// fmt.Println("====yes==== ")
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											}
										}
									}
									//玩法 (三同号单选)
								case 2:
									{
										//依下注的項目列出全部的組合
										ContentAll := thisBOContent.(map[string]interface{})["1"].([]interface{})

										tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)

										for i := 0; i < len(ContentAll); i++ {
											contentArr := ContentAll[i].([]interface{})
											if contentArr[0].(string) == tempResult[0] &&
												contentArr[1].(string) == tempResult[1] &&
												contentArr[2].(string) == tempResult[2] {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1

											}
										}
									}

								//玩法 (二同号复选)
								case 3:
									{
										//依下注的項目列出全部的組合
										ContentAll := thisBOContent.(map[string]interface{})["1"].([]interface{})
										//使用迴圈對獎，依下注的組合逐筆驗證
										for i := 0; i < len(ContentAll); i++ {
											tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)

											if betCount.In_array(ContentAll[i].([]interface{})[0], tempResult) {
												// fmt.Println("====QQQQQQQQQQQQQQQQQQQQQQQQQQQQ==== ", ContentAll[i])
												ResultTemp := betCount.Dup_count(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))
												ResultSorted := betCount.SortMapByValue(ResultTemp)
												for _, v := range ResultSorted {
													if v.Key == ContentAll[i].([]interface{})[0] && v.Value >= 2 {
														Inspect["Ratio"].([]float64)[0]++
														Inspect["Message"] = 1
													}
												}

											}

										}
										break
									}
								case 4, 5:
									{
										var ContentAll []interface{}
										if LotteryPlay == 4 {
											ContentAll = betCount.KuaiThreeTwoKindSingleBet(thisBOContent.(map[string]interface{})["1"].([]interface{}), thisBOContent.(map[string]interface{})["2"].([]interface{}))
										}
										if LotteryPlay == 5 {
											ContentAll = betCount.KuaiThreeTwoKindSimpleBet(thisBOContent.(map[string]interface{})["text"].(string))
										}
										tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)

										for i := 0; i < len(ContentAll); i++ {

											if betCount.In_array(ContentAll[i].([]string)[0], tempResult) &&
												betCount.In_array(ContentAll[i].([]string)[1], tempResult) {
												ResultTemp := betCount.Dup_count(tempResult)
												ResultSorted := betCount.SortMapByValue(ResultTemp)
												for _, v := range ResultSorted {
													// fmt.Println("====QQQQQQQQQQQQQQQQQQQQQQQQQQQQ==== ", ContentAll[i].([]string)[0])
													// fmt.Println("====QQQQQQQQQQQQQQQQQQQQQQQQQQQQ==== ", thisBOContent.(map[string]interface{})[strconv.Itoa(i)])
													if v.Key == ContentAll[i].([]string)[0] && v.Value >= 2 {
														//if v.Key == thisBOContent.(map[string]interface{})[strconv.Itoa(i)].([]string)[0] && v.Value >= 2 {
														Inspect["Ratio"].([]float64)[0]++
														Inspect["Message"] = 1
													}
												}
											}
										}
									}
								//玩法錯誤
								default:
									Inspect["State"] = false
									Inspect["Info"] = "Lottery Play  Error"
								}

							}
						//玩法群組 (三连号)
						case 2:
							{

								//玩法 (因為有對應 TS_LotteryPlay.LP_Code，所以需要跟資料庫一起變動...ry)
								switch LotteryPlay {
								case 1: //玩法 (三同号通选)
									{
										//依下注的項目列出全部的組合
										ContentAll := thisBOContent.(map[string]interface{})["1"].([]interface{})
										//使用迴圈對獎，依下注的組合逐筆驗證

										// fmt.Println("====QQQQQQQQQQQQQQQQQQQQQQQQQQQQ==== ", thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)])
										for i := 0; i < len(ContentAll); i++ {
											// fmt.Println("====QQQQQQQQQQQQQQQQQQQQQQQQQQQQ===>==== ", ContentAll[i])
											// fmt.Println("====QQQQQQQQQQQQQQQQQQQQQQQQQQQQ===>==== ", betCount.In_array(ContentAll[i], thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)]))
											if betCount.In_array(ContentAll[i], thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]interface{})) {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
												// fmt.Println("====QQQQQQQQQsssssssssssssssssssssQQQQQQQQQQQQQQQQQQQ===>==== ", ContentAll[i])
											}
										}
									}
								//玩法錯誤
								default:
									Inspect["State"] = false
									Inspect["Info"] = "Lottery Play  Error"

								}
							}
							//玩法群組 (不同号)
						case 3:
							{
								//玩法 (因為有對應 TS_LotteryPlay.LP_Code，所以需要跟資料庫一起變動...ry)
								switch LotteryPlay {
								//玩法 (三不同号标准) 	//玩法 (三不同号单式)
								case 1, 5:
									{
										var ContentAll []interface{}
										if LotteryPlay == 1 {
											ContentAll = betCount.KuaiThreeThreeDifferBet(thisBOContent.(map[string]interface{})["1"].([]interface{}))
										}
										if LotteryPlay == 5 {
											ContentAll = betCount.KuaiThreeThreeDifferSimpleBet(thisBOContent.(map[string]interface{})["text"].(string))
										}

										if len(betCount.Array_unique_str(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))) >= 3 {
											for i := 0; i < len(ContentAll); i++ {
												if betCount.In_array(ContentAll[i].([]string)[0], thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)) &&
													betCount.In_array(ContentAll[i].([]string)[1], thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)) &&
													betCount.In_array(ContentAll[i].([]string)[2], thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)) {
													Inspect["Ratio"].([]float64)[0]++
													Inspect["Message"] = 1
												}
											}
										}

									}
									//玩法 (三不同号胆拖)(停用)
								case 2:
									{
										// //依下注的項目列出全部的組合
										// $MI_ContentAll = $MI_Module->KuaiThreeThreeDifferDanTuBet($MI_Content->{"1"},$MI_Content->{"2"});
										// //先檢查開獎結果是否符合目前的玩法規則
										// if (count(array_unique($MI_Result->{$MI_LotteryPlayMode}->{$MI_LotteryPlayGroup}->{$MI_LotteryPlay})) >= 3)
										// {l
										// 	//使用迴圈對獎，依下注的組合逐筆驗證
										// 	for ($i = 0; $i < count($MI_ContentAll); $i++)
										// 	{
										// 		if (in_array($MI_ContentAll[$i][0],$MI_Result->{$MI_LotteryPlayMode}->{$MI_LotteryPlayGroup}->{$MI_LotteryPlay})
										// 		 && in_array($MI_ContentAll[$i][1],$MI_Result->{$MI_LotteryPlayMode}->{$MI_LotteryPlayGroup}->{$MI_LotteryPlay})
										// 		 && in_array($MI_ContentAll[$i][2],$MI_Result->{$MI_LotteryPlayMode}->{$MI_LotteryPlayGroup}->{$MI_LotteryPlay})) { $MI_Inspect->{'Ratio'}[0] += 1; $MI_Inspect->{'Message'} = 1; }
										// 	}
										// }
									}

									//玩法 (二不同号标准)	//玩法 (二不同号单式)
								case 3, 6:
									{
										var ContentAll []interface{}
										if LotteryPlay == 3 {
											ContentAll = betCount.KuaiThreeTwoDifferBet(thisBOContent.(map[string]interface{})["1"].([]interface{}))
										}
										if LotteryPlay == 6 {
											ContentAll = betCount.KuaiThreeTwoDifferSimpleBet(thisBOContent.(map[string]interface{})["text"].(string))

										}

										if len(betCount.Array_unique_str(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))) >= 2 {
											for i := 0; i < len(ContentAll); i++ {
												if betCount.In_array(ContentAll[i].([]string)[0], thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)) &&
													betCount.In_array(ContentAll[i].([]string)[1], thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)) {
													Inspect["Ratio"].([]float64)[0]++
													Inspect["Message"] = 1
												}
											}
										}

									}
									//玩法 (二不同号胆拖)(停用)
								case 4:
									{
										ContentAll := betCount.KuaiThreeTwoDifferDanTuBet(thisBOContent.(map[string]interface{})["1"].([]interface{}), thisBOContent.(map[string]interface{})["2"].([]interface{}))

										if len(betCount.Array_unique_str(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))) >= 2 {
											for i := 0; i < len(ContentAll); i++ {
												if betCount.In_array(ContentAll[i].([]string)[0], thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)) &&
													betCount.In_array(ContentAll[i].([]string)[1], thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)) {
													Inspect["Ratio"].([]float64)[0]++
													Inspect["Message"] = 1
												}
											}

										}
									}
								//玩法錯誤
								default:
									Inspect["State"] = false
									Inspect["Info"] = "Lottery Play  Error"
								}
							}
						default:
							Inspect["State"] = false
							Inspect["Info"] = "Lottery Play Group Error"
							break
						}
					}
				}

			default:
				Inspect["State"] = false
				Inspect["Info"] = "Lottery Play Mode Group Error"
			}
		}

	default:
		Inspect["State"] = false
	}

	return Inspect
}
func InspectModuleMarkSixTrad(LotteryTypeGroup int,
	LotteryType int,
	LotteryMode int,
	LotteryPlayGroup int,
	LotteryPlay int,
	thisBOContent interface{},
	thisFullResult map[string]interface{},
	config map[string]interface{}) map[string]interface{} {

	//初始化
	Inspect := make(map[string]interface{})
	//對獎結果 (0：沒中獎|1:中獎|2:和局退本金)
	Inspect["Message"] = 0
	//中獎金額倍數 (先預設為0)(簡單的說，就是嬴的注數…)
	Inspect["Ratio"] = []float64{0, 0, 0, 0}
	// //中獎賠率 (預設為1，即使用第一個)
	// CWCheck["Odds"] = 1;
	//執行狀態
	Inspect["State"] = true
	// //執行訊息
	Inspect["Info"] = ""
	//最大下注內容上限 (每一注單/每一玩法)(如果投注項目超過上限，超出的項目全部無視掉 (σ′▽‵)′▽‵)σ)
	ContentMax := int(config["LTR_ContentMax"].(map[string]interface{})[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].(float64))

	//彩種 (因為有對應 TS_LotteryType.LT_Code，所以需要跟資料庫一起變動...ry)
	switch LotteryTypeGroup {
	//六合彩
	case 5:
		{
			//彩種 (因為有對應 TS_LotteryType.LT_Code，所以需要跟資料庫一起變動...ry)
			switch LotteryType {
			case 7, 32, 38, 39: //六合彩 //區塊鏈六合彩 // 1分六合彩 // 5分六合彩
				switch LotteryMode {
				case 1: //信用模式 (传统)
					{
						//玩法群組 (因為有對應 TS_LotteryPlayGroup.LPG_Code，所以需要跟資料庫一起變動...ry)

						// fmt.Println("====LotteryPlayGroup  ", LotteryPlayGroup)
						// fmt.Println("====LotteryPlay  ", LotteryPlay)
						// fmt.Println("====thisBOContent  ", thisBOContent)
						// fmt.Println("====thisFullResult  ", thisFullResult)
						switch LotteryPlayGroup {
						//玩法群組 (特码)
						case 1:
							{

								switch LotteryPlay {
								//玩法 (特码)
								case 1:
									{
										if len(thisBOContent.(map[string]interface{})) < ContentMax {
											ContentMax = len(thisBOContent.(map[string]interface{}))
										}
										//fmt.Println("====ContentMax  ", ContentMax)

										for i := 0; i < ContentMax; i++ {

											// fmt.Println("====thisBOContent.(map[string]interface{})[strconv.Itoa(i)]  ", thisBOContent.(map[string]interface{})[strconv.Itoa(i)])
											// fmt.Println("====thisBOContent.(map[string]interface{})[strconv.Itoa(i)]  ", thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)])
											if thisBOContent.(map[string]interface{})[strconv.Itoa(i)] == thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] {

												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											}
										}
									}
									//玩法 (两面)
								case 2:
									{
										if len(thisBOContent.(map[string]interface{})) < ContentMax {
											ContentMax = len(thisBOContent.(map[string]interface{}))
										}
										//fmt.Println("====ContentMax  ", ContentMax)
										for i := 0; i < ContentMax; i++ {
											// fmt.Println(" thisFullResult  ", thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)])
											// fmt.Println(" thisFullResult  ", reflect.TypeOf(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)]))

											if betCount.In_array(thisBOContent.(map[string]interface{})[strconv.Itoa(i)], []string{"B", "S", "O", "E", "FB", "FS", "LO", "LE", "MO", "ME", "PB", "PS", "PO", "PE", "TZ", "UZ", "AZ", "FZ", "HZ", "WZ"}) &&
												betCount.In_array("N", thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]interface{})) {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 2
											} else if betCount.In_array(thisBOContent.(map[string]interface{})[strconv.Itoa(i)], thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]interface{})) {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											}
										}
									}
									//玩法錯誤
								default:
									Inspect["State"] = false
									Inspect["Info"] = "Lottery Play  Error"

								}

							}
						//玩法 (两面)
						case 2:
							{
								//玩法 (因為有對應 TS_LotteryPlay.LP_Code，所以需要跟資料庫一起變動...ry)
								switch LotteryPlay {
								//玩法 (特码)
								case 1:
									{
										if len(thisBOContent.(map[string]interface{})) < ContentMax {
											ContentMax = len(thisBOContent.(map[string]interface{}))
										}
										for i := 0; i < ContentMax; i++ {
											if betCount.In_array(thisBOContent.(map[string]interface{})[strconv.Itoa(i)], []string{"B", "S", "O", "E", "FB", "FS", "LO", "LE", "MO", "ME", "PB", "PS", "PO", "PE", "TZ", "UZ", "AZ", "FZ", "HZ", "WZ"}) &&
												betCount.In_array("N", thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]interface{})) {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 2
											} else if betCount.In_array(thisBOContent.(map[string]interface{})[strconv.Itoa(i)], thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]interface{})) {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											}
										}

									}
									//玩法錯誤
								default:
									Inspect["State"] = false
									Inspect["Info"] = "Lottery Play  Error"
								}

							}
						//玩法群組 (正码)
						case 3:
							{
								switch LotteryPlay {
								//玩法 (正码)
								//玩法 (两面)
								case 1, 2:
									{
										if len(thisBOContent.(map[string]interface{})) < ContentMax {
											ContentMax = len(thisBOContent.(map[string]interface{}))
										}
										for i := 0; i < ContentMax; i++ {
											if betCount.In_array(thisBOContent.(map[string]interface{})[strconv.Itoa(i)], []string{"B", "S", "O", "E", "FB", "FS", "LO", "LE", "MO", "ME", "PB", "PS", "PO", "PE", "TZ", "UZ", "AZ", "FZ", "HZ", "WZ"}) &&
												betCount.In_array("N", thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]interface{})) {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 2
											} else if betCount.In_array(thisBOContent.(map[string]interface{})[strconv.Itoa(i)], thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)) {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											}
										}

									}
								default:
									Inspect["State"] = false
									Inspect["Info"] = "Lottery Play  Error"

								}
							}
						//玩法群組 (正码1-6)
						case 4:
							{
								switch LotteryPlay {
								//正码1-6
								case 1, 2, 3, 4, 5, 6:
									{
										if len(thisBOContent.(map[string]interface{})) < ContentMax {
											ContentMax = len(thisBOContent.(map[string]interface{}))
										}
										for i := 0; i < ContentMax; i++ {
											// fmt.Println("正码1-6 ====thisBOContent  ", thisBOContent)
											// fmt.Println("正码1-6 ====thisFullResult  ", thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)])

											if betCount.In_array(thisBOContent.(map[string]interface{})[strconv.Itoa(i)], []string{"B", "S", "O", "E", "FB", "FS", "LO", "LE", "MO", "ME", "PB", "PS", "PO", "PE", "TZ", "UZ", "AZ", "FZ", "HZ", "WZ"}) &&
												betCount.In_array("N", thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]interface{})) {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 2
												// fmt.Println("====CC  ")
											} else if betCount.In_array(thisBOContent.(map[string]interface{})[strconv.Itoa(i)], thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]interface{})) {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
												// fmt.Println("====DD ")
											}
										}
									}

								default:
									Inspect["State"] = false
									Inspect["Info"] = "Lottery Play  Error"
								}
							}
						case 5: //玩法群組 (正特)
							{
								//玩法 (正1特 - //玩法 (正6特) )
								switch LotteryPlay {
								//玩法 (正1特) -  (正6特)
								case 1, 3, 5, 7, 9, 11:
									{
										if len(thisBOContent.(map[string]interface{})) < ContentMax {
											ContentMax = len(thisBOContent.(map[string]interface{}))
										}
										for i := 0; i < ContentMax; i++ {
											if thisBOContent.(map[string]interface{})[strconv.Itoa(i)].(string) == thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].(string) {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											}
										}
									}
									//玩法 (正1特双面) -  (正2特双面)
								case 2, 4, 6, 8, 10, 12:
									{
										if len(thisBOContent.(map[string]interface{})) < ContentMax {
											ContentMax = len(thisBOContent.(map[string]interface{}))
										}
										for i := 0; i < ContentMax; i++ {
											if betCount.In_array(thisBOContent.(map[string]interface{})[strconv.Itoa(i)], []string{"B", "S", "O", "E", "FB", "FS", "PB", "PS", "PO", "PE"}) &&
												betCount.In_array("N", thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]interface{})) {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 2
											} else if betCount.In_array(thisBOContent.(map[string]interface{})[strconv.Itoa(i)], thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]interface{})) {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											}
										}
									}
								default:
									Inspect["State"] = false
									Inspect["Info"] = "Lottery Play  Error"

								}
							}
							//玩法群組 (连码)
						case 6:
							{
								//玩法 (因為有對應 TS_LotteryPlay.LP_Code，所以需要跟資料庫一起變動...ry)
								switch LotteryPlay {
								//玩法 (三全中)
								case 1:
									{
										//依下注的項目列出全部的組合
										ContentAll := betCount.MarkSixStraightThreeGetAllBet(thisBOContent.(map[string]interface{})["1"].([]interface{}))
										//使用迴圈對獎，依下注的組合逐筆驗證
										FullResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)
										for i := 0; i < len(ContentAll); i++ {
											if betCount.In_array(ContentAll[i].([]string)[0], FullResult) &&
												betCount.In_array(ContentAll[i].([]string)[1], FullResult) &&
												betCount.In_array(ContentAll[i].([]string)[2], FullResult) {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											}
										}
									}
								//玩法 (三中二)
								case 2:
									{
										// fmt.Println("三中二  ", thisBOContent)
										//依下注的項目列出全部的組合
										ContentAll := betCount.MarkSixStraightThreeGetTwoBet(thisBOContent.(map[string]interface{})["1"].([]interface{}))
										//使用迴圈對獎，依下注的組合逐筆驗證
										// fmt.Println("三中二  ", reflect.TypeOf(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)]))
										FullResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)]
										for i := 0; i < len(ContentAll); i++ {
											if betCount.In_array(ContentAll[i].([]string)[0], FullResult) &&
												betCount.In_array(ContentAll[i].([]string)[1], FullResult) &&
												betCount.In_array(ContentAll[i].([]string)[2], FullResult) {
												Inspect["Ratio"].([]float64)[1]++
												Inspect["Message"] = 1

											} else if betCount.In_array(ContentAll[i].([]string)[0], FullResult) &&
												betCount.In_array(ContentAll[i].([]string)[1], FullResult) {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											} else if betCount.In_array(ContentAll[i].([]string)[0], FullResult) &&
												betCount.In_array(ContentAll[i].([]string)[2], FullResult) {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											} else if betCount.In_array(ContentAll[i].([]string)[1], FullResult) &&
												betCount.In_array(ContentAll[i].([]string)[2], FullResult) {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											}
										}
									}
									//玩法 (二全中)
								case 3: //玩法 (二全中)
									{
										// fmt.Println(" thisBOContent  ", thisBOContent.(map[string]interface{}))
										// fmt.Println(" thisBOContent  ", reflect.TypeOf(thisBOContent.(map[string]interface{})))
										// fmt.Println(" thisBOContent  ", thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)])

										// fmt.Println(" thisBOContent  ", reflect.TypeOf(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)]))
										//依下注的項目列出全部的組合
										ContentAll := betCount.MarkSixStraightTwoGetAllBet(thisBOContent.(map[string]interface{})["1"].([]interface{}))
										//使用迴圈對獎，依下注的組合逐筆驗證
										FullResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)
										// fmt.Println(" ContentAll  ", ContentAll)
										for i := 0; i < len(ContentAll); i++ {
											if betCount.In_array(ContentAll[i].([]string)[0], FullResult) &&
												betCount.In_array(ContentAll[i].([]string)[1], FullResult) {

												// fmt.Println(" ContentAll[i]  ", ContentAll[i])
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											}
										}
										// //依下注的項目列出全部的組合
										// ContentAll = betCount.MarkSixStraightThreeGetAllBet(thisBOContent)
										// //使用迴圈對獎，依下注的組合逐筆驗證
										// FullResult = thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]interface{})[1].([]string)
										// for i := 0; i < len(ContentAll); i++ {

										// }
									}
								// 玩法 (二中特)
								case 4:
									{
										//依下注的項目列出全部的組合
										ContentAll := betCount.MarkSixStraightTwoGetUniqueBet(thisBOContent.(map[string]interface{})["1"].([]interface{}))
										//使用迴圈對獎，依下注的組合逐筆驗證

										FullResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]interface{})
										for i := 0; i < len(ContentAll); i++ {

											// fmt.Println(" FullResult0  ", FullResult[0])
											// fmt.Println(" ContentAll[i].([]string)[0]  ", ContentAll[i].([]string)[0])

											// fmt.Println(" FullResult1  ", FullResult[1])
											// fmt.Println(" ContentAll[i].([]string)[1]  ", ContentAll[i].([]string)[1])

											// fmt.Println(" betCount.In_array(FullResult[0], ContentAll[i].([]string)[0])  ", betCount.In_array(FullResult[0], ContentAll[i].([]string)[0]))
											// fmt.Println("FullResult[1] == ContentAll[i].([]string)[1]  ", FullResult[1] == ContentAll[i].([]string)[1])

											if betCount.In_array(ContentAll[i].([]string)[0], FullResult[0]) &&
												betCount.In_array(ContentAll[i].([]string)[1], FullResult[0]) {
												Inspect["Ratio"].([]float64)[1]++
												Inspect["Message"] = 1

											} else if betCount.In_array(ContentAll[i].([]string)[0], FullResult[0]) &&
												FullResult[1] == ContentAll[i].([]string)[1] {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											} else if FullResult[1] == ContentAll[i].([]string)[0] &&
												betCount.In_array(ContentAll[i].([]string)[1], FullResult[0]) {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											}
										}
									}
								// //玩法 (特串)
								case 5:
									{
										ContentAll := betCount.MarkSixStraightUniqueThreadBet(thisBOContent.(map[string]interface{})["1"].([]interface{}))
										//使用迴圈對獎，依下注的組合逐筆驗證
										FullResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]interface{})
										for i := 0; i < len(ContentAll); i++ {
											if betCount.In_array(ContentAll[i].([]string)[0], FullResult[0]) &&
												FullResult[1] == ContentAll[i].([]string)[1] {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											} else if FullResult[1] == ContentAll[i].([]string)[0] &&
												betCount.In_array(ContentAll[i].([]string)[1], FullResult[0]) {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											}

										}
									}
								// //玩法 (四全中)
								case 6:
									{

										ContentAll := betCount.MarkSixStraightFourGetAllBet(thisBOContent.(map[string]interface{})["1"].([]interface{}))

										// fmt.Println(" ContentAll  ", ContentAll)
										// fmt.Println(" thisBOContent  ", thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)])

										//使用迴圈對獎，依下注的組合逐筆驗證
										FullResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)
										for i := 0; i < len(ContentAll); i++ {
											if betCount.In_array(ContentAll[i].([]string)[0], FullResult) &&
												betCount.In_array(ContentAll[i].([]string)[1], FullResult) &&
												betCount.In_array(ContentAll[i].([]string)[2], FullResult) &&
												betCount.In_array(ContentAll[i].([]string)[3], FullResult) {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1

											}
										}

									}
								default:
									Inspect["State"] = false
									Inspect["Info"] = "Lottery Play  Error"

								}
							}
							//玩法群組(色波)
						case 7:
							{
								switch LotteryPlay {

								//玩法 (特码色波) //玩法 (7色波)
								case 1, 2:
									{
										if len(thisBOContent.(map[string]interface{})) < ContentMax {
											ContentMax = len(thisBOContent.(map[string]interface{}))
										}
										for i := 0; i < ContentMax; i++ {
											if thisBOContent.(map[string]interface{})[strconv.Itoa(i)].(string) != "N" && thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].(string) == "N" {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 2
											} else if thisBOContent.(map[string]interface{})[strconv.Itoa(i)].(string) == thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].(string) {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											}

										}
									}
									//玩法 (红球) //玩法 (蓝球) //玩法 (绿球)
								case 3, 4, 5:
									{
										if len(thisBOContent.(map[string]interface{})) < ContentMax {
											ContentMax = len(thisBOContent.(map[string]interface{}))
										}

										for i := 0; i < ContentMax; i++ {
											if betCount.In_array("N", thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)]) && thisBOContent.(map[string]interface{})[strconv.Itoa(i)].(string) != "N" {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 2
											} else if betCount.In_array(thisBOContent.(map[string]interface{})[strconv.Itoa(i)], thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)]) {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											}

										}
									}

								default:
									Inspect["State"] = false
									Inspect["Info"] = "Lottery Play  Error"
								}
							}
						//玩法群組 (特码头尾数)
						case 8:
							{
								//玩法 (因為有對應 TS_LotteryPlay.LP_Code，所以需要跟資料庫一起變動...ry)
								switch LotteryPlay {
								//玩法 (头数) 	//玩法 (尾数)
								case 1, 2:
									{
										if len(thisBOContent.(map[string]interface{})) < ContentMax {
											ContentMax = len(thisBOContent.(map[string]interface{}))
										}

										for i := 0; i < ContentMax; i++ {

											// fmt.Println(" thisBOContent  ", thisBOContent.(map[string]interface{})[strconv.Itoa(i)])
											// fmt.Println(" thisBOContent  ", reflect.TypeOf(thisBOContent.(map[string]interface{})[strconv.Itoa(i)]))
											// fmt.Println(" thisFullResult  ", thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{}))

											// fmt.Println(" thisBOContent  ", reflect.TypeOf(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)]))

											// fmt.Println(" In_array  ", thisBOContent.(map[string]interface{})[strconv.Itoa(i)].(string) == strconv.Itoa(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].(int)))

											if thisBOContent.(map[string]interface{})[strconv.Itoa(i)].(string) == strconv.Itoa(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].(int)) {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											}
										}
									}
								default:
									Inspect["State"] = false
									Inspect["Info"] = "Lottery Play  Error"
								}
							}
							//玩法 (总肖)
						case 9:
							{
								//玩法 (因為有對應 TS_LotteryPlay.LP_Code，所以需要跟資料庫一起變動...ry)
								switch LotteryPlay {
								//玩法 (总肖)
								case 1:
									{
										if len(thisBOContent.(map[string]interface{})) < ContentMax {
											ContentMax = len(thisBOContent.(map[string]interface{}))
										}

										// fmt.Println(" thisBOContent  ", thisBOContent.(map[string]interface{}))

										// fmt.Println(" thisFullResult  ", thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)])
										// fmt.Println(" thisBOContent  ", reflect.TypeOf(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)]))
										// fmt.Println(" ContentMax  ", len(thisBOContent.(map[string]interface{})))
										for i := 0; i < ContentMax; i++ {

											// fmt.Println("  ===================    ", thisBOContent.(map[string]interface{})[strconv.Itoa(i)].(string))
											// fmt.Println(" thisFullResult  ", thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)])
											// fmt.Println(" In_array  ", betCount.In_array(thisBOContent.(map[string]interface{})[strconv.Itoa(i)].(string), thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)]))
											if betCount.In_array(thisBOContent.(map[string]interface{})[strconv.Itoa(i)].(string), thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)]) {

												// fmt.Println("  ===================    ")
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											}
										}

									}
								default:
									Inspect["State"] = false
									Inspect["Info"] = "Lottery Play  Error"
								}
							}
							//玩法群組 (平特一肖尾数)
						case 10:
							{

								//玩法 (因為有對應 TS_LotteryPlay.LP_Code，所以需要跟資料庫一起變動...ry)
								switch LotteryPlay {
								//玩法 (一肖) 	//玩法 (尾数)
								case 1, 2:
									{
										if len(thisBOContent.(map[string]interface{})) < ContentMax {
											ContentMax = len(thisBOContent.(map[string]interface{}))
										}

										// fmt.Println(" thisBOContent  ", reflect.TypeOf(thisBOContent.(map[string]interface{})["0"]))
										// fmt.Println(" thisFullResult  ", thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)])
										// fmt.Println(" thisBOContent  ", reflect.TypeOf(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)]))

										FullResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)]
										// fmt.Println(" ==========1==")
										for i := 0; i < ContentMax; i++ {

											BOContentInt, _ := strconv.Atoi(thisBOContent.(map[string]interface{})[strconv.Itoa(i)].(string))
											// fmt.Println(" In_array QQQ ", betCount.In_array(BOContentInt, FullResult))
											if betCount.In_array(BOContentInt, FullResult) {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											}
										}
									}
								default:
									Inspect["State"] = false
									Inspect["Info"] = "Lottery Play  Error"
								}
							}
						//玩法群組 (特肖)
						case 11:
							{
								//玩法 (因為有對應 TS_LotteryPlay.LP_Code，所以需要跟資料庫一起變動...ry)
								switch LotteryPlay {
								//玩法 (特肖)
								case 1:
									{
										if len(thisBOContent.(map[string]interface{})) < ContentMax {
											ContentMax = len(thisBOContent.(map[string]interface{}))
										}

										// fmt.Println(" thisBOContent  ", reflect.TypeOf(thisBOContent.(map[string]interface{})["0"]))
										// fmt.Println(" thisFullResult  ", thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)])
										// fmt.Println(" thisBOContent  ", reflect.TypeOf(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)]))

										FullResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].(int)
										for i := 0; i < ContentMax; i++ {
											BOContentInt, _ := strconv.Atoi(thisBOContent.(map[string]interface{})[strconv.Itoa(i)].(string))
											if BOContentInt == FullResult {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											}
										}
									}
								default:
									Inspect["State"] = false
									Inspect["Info"] = "Lottery Play  Error"
								}
							}
						//玩法群組 (连肖)
						case 12:
							{
								//玩法 (因為有對應 TS_LotteryPlay.LP_Code，所以需要跟資料庫一起變動...ry)
								switch LotteryPlay {
								//玩法 (二连肖)
								case 1:
									{
										ContentAll := betCount.MarkSixStraightTwoZodiacBet(thisBOContent.([]interface{}))
										//使用迴圈對獎，依下注的組合逐筆驗證
										FullResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]interface{})
										for i := 0; i < len(ContentAll); i++ {
											if betCount.In_array(ContentAll[i].([]string)[0], FullResult) &&
												betCount.In_array(ContentAll[i].([]string)[1], FullResult) {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1

											}
										}
									}
									//玩法 (三连肖)
								case 2:
									{
										ContentAll := betCount.MarkSixStraightThreeZodiacBet(thisBOContent.([]interface{}))
										//使用迴圈對獎，依下注的組合逐筆驗證
										FullResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]interface{})
										for i := 0; i < len(ContentAll); i++ {
											if betCount.In_array(ContentAll[i].([]string)[0], FullResult) &&
												betCount.In_array(ContentAll[i].([]string)[1], FullResult) &&
												betCount.In_array(ContentAll[i].([]string)[2], FullResult) {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1

											}
										}
									}
									//玩法 (四连尾)
								case 3:
									{

									}
									//玩法 (五连尾)
								case 4:
									{
									}
								default:
									Inspect["State"] = false
									Inspect["Info"] = "Lottery Play  Error"
								}
							}
							//玩法群組 (正肖)
						case 15:
							{
								//玩法 (因為有對應 TS_LotteryPlay.LP_Code，所以需要跟資料庫一起變動...ry)
								switch LotteryPlay {
								//玩法 (正肖)
								case 1:
									{

									}
								default:
									Inspect["State"] = false
									Inspect["Info"] = "Lottery Play  Error"
								}
							}
							//玩法群組 (五行)
						case 16:
							{
								//玩法 (因為有對應 TS_LotteryPlay.LP_Code，所以需要跟資料庫一起變動...ry)
								switch LotteryPlay {
								//玩法 (正肖)
								case 1:
									{

									}
								default:
									Inspect["State"] = false
									Inspect["Info"] = "Lottery Play  Error"
								}
							}
							//玩法群組 (自选不中)
						case 17:
							{
								//玩法 (因為有對應 TS_LotteryPlay.LP_Code，所以需要跟資料庫一起變動...ry)
								switch LotteryPlay {
								//玩法 (正肖)
								case 1:
									{

									}
								default:
									Inspect["State"] = false
									Inspect["Info"] = "Lottery Play  Error"
								}
							}
						default:
							Inspect["State"] = false
							Inspect["Info"] = "Lottery Play Group  Error"
						}

					}

				case 2: //官方模式
					{
						//玩法群組 (因為有對應 TS_LotteryPlayGroup.LPG_Code，所以需要跟資料庫一起變動...ry)
						switch LotteryPlayGroup {
						//玩法群組 (同号)
						case 1:
							{
								//玩法 (因為有對應 TS_LotteryPlay.LP_Code，所以需要跟資料庫一起變動...ry)
								switch LotteryPlay {

								//玩法 (三同号通选)
								case 1:
									{
										//依下注的項目列出全部的組合
										ContentAll := thisBOContent.(map[string]interface{})["1"].([]interface{})
										// fmt.Println("====ContentAll==== ", ContentAll)
										for i := 0; i < len(ContentAll); i++ {
											tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]interface{})

											if betCount.In_array(ContentAll[i].(string), tempResult) {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											}
										}
									}
									//玩法 (三同号单选)
								case 2:
									{
										//依下注的項目列出全部的組合
										ContentAll := thisBOContent.(map[string]interface{})["1"].([]interface{})

										tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)

										for i := 0; i < len(ContentAll); i++ {
											if betCount.In_array(ContentAll[i], tempResult) {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1

											}
										}
									}

								//玩法 (二同号复选)
								case 3:
									{
										//依下注的項目列出全部的組合
										ContentAll := thisBOContent.(map[string]interface{})["1"].([]interface{})
										//使用迴圈對獎，依下注的組合逐筆驗證
										for i := 0; i < len(ContentAll); i++ {
											tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)
											if betCount.In_array(ContentAll[i].([]interface{})[0], tempResult) {
												// fmt.Println("====QQQQQQQQQQQQQQQQQQQQQQQQQQQQ==== ", ContentAll[i])
												ResultTemp := betCount.Dup_count(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))
												ResultSorted := betCount.SortMapByValue(ResultTemp)
												for _, v := range ResultSorted {
													if v.Key == ContentAll[i].([]interface{})[0] && v.Value >= 2 {
														// fmt.Println("====QQQQQQQQQQQQQQQQQccccQQQQQQQQQQQ==== ", ContentAll[i])
														Inspect["Ratio"].([]float64)[0]++
														Inspect["Message"] = 1
													}
												}
											}
										}
										break
									}
								case 4, 5:
									{
										var ContentAll []interface{}
										if LotteryPlay == 4 {
											ContentAll = betCount.KuaiThreeTwoKindSingleBet(thisBOContent.(map[string]interface{})["1"].([]interface{}), thisBOContent.(map[string]interface{})["2"].([]interface{}))
										}
										if LotteryPlay == 5 {
											ContentAll = betCount.KuaiThreeTwoKindSimpleBet(thisBOContent.(map[string]interface{})["text"].(string))
										}
										tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)
										for i := 0; i < len(ContentAll); i++ {
											// fmt.Println("====tempResult==== ", tempResult)
											// fmt.Println("====ContentAll[i]==== ", ContentAll[i])
											if betCount.In_array(ContentAll[i].([]string)[0], tempResult) &&
												betCount.In_array(ContentAll[i].([]string)[1], tempResult) {

												ResultTemp := betCount.Dup_count(tempResult)
												ResultSorted := betCount.SortMapByValue(ResultTemp)
												for _, v := range ResultSorted {
													// fmt.Println("====QQQQQQQQQQQQQQQQQQQQQQQQQQQQ==== ", ContentAll[i].([]string)[0])
													// fmt.Println("====QQQQQQQQQQQQQQQQQQQQQQQQQQQQ==== ", thisBOContent.(map[string]interface{})[strconv.Itoa(i)])
													if v.Key == ContentAll[i].([]string)[0] && v.Value >= 2 {
														//if v.Key == thisBOContent.(map[string]interface{})[strconv.Itoa(i)].([]string)[0] && v.Value >= 2 {
														Inspect["Ratio"].([]float64)[0]++
														Inspect["Message"] = 1
													}
												}
											}
										}
									}
								//玩法錯誤
								default:
									Inspect["State"] = false
									Inspect["Info"] = "Lottery Play  Error"
								}

							}
						//玩法群組 (三连号)
						case 2:
							{

								//玩法 (因為有對應 TS_LotteryPlay.LP_Code，所以需要跟資料庫一起變動...ry)
								switch LotteryPlay {
								case 1: //玩法 (三同号通选)
									{
										//依下注的項目列出全部的組合
										ContentAll := thisBOContent.(map[string]interface{})["1"].([]interface{})
										//使用迴圈對獎，依下注的組合逐筆驗證

										// fmt.Println("====QQQQQQQQQQQQQQQQQQQQQQQQQQQQ==== ", thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)])
										for i := 0; i < len(ContentAll); i++ {
											// fmt.Println("====QQQQQQQQQQQQQQQQQQQQQQQQQQQQ===>==== ", ContentAll[i])
											// fmt.Println("====QQQQQQQQQQQQQQQQQQQQQQQQQQQQ===>==== ", betCount.In_array(ContentAll[i], thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)]))
											if betCount.In_array(ContentAll[i], thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]interface{})) {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
												// fmt.Println("====QQQQQQQQQsssssssssssssssssssssQQQQQQQQQQQQQQQQQQQ===>==== ", ContentAll[i])
											}
										}
									}
								//玩法錯誤
								default:
									Inspect["State"] = false
									Inspect["Info"] = "Lottery Play  Error"

								}
							}
							//玩法群組 (不同号)
						case 3:
							{
								//玩法 (因為有對應 TS_LotteryPlay.LP_Code，所以需要跟資料庫一起變動...ry)
								switch LotteryPlay {
								//玩法 (三不同号标准) 	//玩法 (三不同号单式)
								case 1, 5:
									{
										var ContentAll []interface{}
										if LotteryPlay == 1 {
											ContentAll = betCount.KuaiThreeThreeDifferBet(thisBOContent.(map[string]interface{})["1"].([]interface{}))
										}
										if LotteryPlay == 5 {
											ContentAll = betCount.KuaiThreeThreeDifferSimpleBet(thisBOContent.(map[string]interface{})["text"].(string))
										}

										if len(betCount.Array_unique_str(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))) >= 3 {
											for i := 0; i < len(ContentAll); i++ {
												if betCount.In_array(ContentAll[i].([]string)[0], thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)) &&
													betCount.In_array(ContentAll[i].([]string)[1], thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)) &&
													betCount.In_array(ContentAll[i].([]string)[2], thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)) {
													Inspect["Ratio"].([]float64)[0]++
													Inspect["Message"] = 1
												}
											}
										}

									}
									//玩法 (三不同号胆拖)(停用)
								case 2:
									{
										// //依下注的項目列出全部的組合
										// $MI_ContentAll = $MI_Module->KuaiThreeThreeDifferDanTuBet($MI_Content->{"1"},$MI_Content->{"2"});
										// //先檢查開獎結果是否符合目前的玩法規則
										// if (count(array_unique($MI_Result->{$MI_LotteryPlayMode}->{$MI_LotteryPlayGroup}->{$MI_LotteryPlay})) >= 3)
										// {l
										// 	//使用迴圈對獎，依下注的組合逐筆驗證
										// 	for ($i = 0; $i < count($MI_ContentAll); $i++)
										// 	{
										// 		if (in_array($MI_ContentAll[$i][0],$MI_Result->{$MI_LotteryPlayMode}->{$MI_LotteryPlayGroup}->{$MI_LotteryPlay})
										// 		 && in_array($MI_ContentAll[$i][1],$MI_Result->{$MI_LotteryPlayMode}->{$MI_LotteryPlayGroup}->{$MI_LotteryPlay})
										// 		 && in_array($MI_ContentAll[$i][2],$MI_Result->{$MI_LotteryPlayMode}->{$MI_LotteryPlayGroup}->{$MI_LotteryPlay})) { $MI_Inspect->{'Ratio'}[0] += 1; $MI_Inspect->{'Message'} = 1; }
										// 	}
										// }
									}

									//玩法 (二不同号标准)	//玩法 (二不同号单式)
								case 3, 6:
									{
										var ContentAll []interface{}
										if LotteryPlay == 3 {
											ContentAll = betCount.KuaiThreeTwoDifferBet(thisBOContent.(map[string]interface{})["1"].([]interface{}))
										}
										if LotteryPlay == 6 {
											ContentAll = betCount.KuaiThreeTwoDifferSimpleBet(thisBOContent.(map[string]interface{})["text"].(string))

										}

										if len(betCount.Array_unique_str(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))) >= 2 {
											for i := 0; i < len(ContentAll); i++ {
												if betCount.In_array(ContentAll[i].([]string)[0], thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)) &&
													betCount.In_array(ContentAll[i].([]string)[1], thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)) {
													Inspect["Ratio"].([]float64)[0]++
													Inspect["Message"] = 1
												}
											}
										}

									}
									//玩法 (二不同号胆拖)(停用)
								case 4:
									{
										ContentAll := betCount.KuaiThreeTwoDifferDanTuBet(thisBOContent.(map[string]interface{})["1"].([]interface{}), thisBOContent.(map[string]interface{})["2"].([]interface{}))

										if len(betCount.Array_unique_str(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))) >= 2 {
											for i := 0; i < len(ContentAll); i++ {
												if betCount.In_array(ContentAll[i].([]string)[0], thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)) &&
													betCount.In_array(ContentAll[i].([]string)[1], thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)) {
													Inspect["Ratio"].([]float64)[0]++
													Inspect["Message"] = 1
												}
											}

										}
									}
								//玩法錯誤
								default:
									Inspect["State"] = false
									Inspect["Info"] = "Lottery Play  Error"
								}
							}
						default:
							Inspect["State"] = false
							Inspect["Info"] = "Lottery Play Group Error"
							break
						}
					}
				}

			default:
				Inspect["State"] = false
				Inspect["Info"] = "Lottery Play Mode Group Error"
			}
		}

	default:
		Inspect["State"] = false
	}

	return Inspect
}

func InspectModulePKTen(LotteryTypeGroup int,
	LotteryType int,
	LotteryMode int,
	LotteryPlayGroup int,
	LotteryPlay int,
	thisBOContent interface{},
	thisFullResult map[string]interface{},
	config map[string]interface{}) map[string]interface{} {

	//初始化
	Inspect := make(map[string]interface{})
	//對獎結果 (0：沒中獎|1:中獎|2:和局退本金)
	Inspect["Message"] = 0
	//中獎金額倍數 (先預設為0)(簡單的說，就是嬴的注數…)
	Inspect["Ratio"] = []float64{0, 0, 0, 0}
	// //中獎賠率 (預設為1，即使用第一個)
	// CWCheck["Odds"] = 1;
	//執行狀態
	Inspect["State"] = true
	// //執行訊息
	Inspect["Info"] = ""
	//最大下注內容上限 (每一注單/每一玩法)(如果投注項目超過上限，超出的項目全部無視掉 (σ′▽‵)′▽‵)σ)
	ContentMax := int(config["LTR_ContentMax"].(map[string]interface{})[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].(float64))

	switch LotteryTypeGroup {
	//北京PK10 //幸运飞艇
	case 1, 3:
		{
			switch LotteryType {
			case 1, 5, 31, 34, 35: //北京PK10 //幸运飞艇 // 區塊鏈 PK10 // 1分PK10 // 3分PK10
				switch LotteryMode {
				case 1: //传统模式 (信用模式)
					{
						//玩法群組 (因為有對應 TS_LotteryPlayGroup.LPG_Code，所以需要跟資料庫一起變動...ry)
						switch LotteryPlayGroup {
						//玩法群組 (两面) //玩法群組 (冠亚和)
						case 1, 2:
							{
								//如果投注項目比規則中的項目少，就只比對投注的項目就好 (σﾟ∀ﾟ)σﾟ∀ﾟ)σ
								if len(thisBOContent.(map[string]interface{})) < ContentMax {
									ContentMax = len(thisBOContent.(map[string]interface{}))
								}
								for i := 0; i < ContentMax; i++ {
									if betCount.In_array(thisBOContent.(map[string]interface{})[strconv.Itoa(i)], thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)]) {
										//	Inspect["Ratio"].(int)[0] = Inspect["Ratio"].(int)[0] + 1
										// fmt.Println("<====> betCount  ")
										Inspect["Ratio"].([]float64)[0]++
										Inspect["Message"] = 1
									}
								}

							}
						case 3, 4: //玩法群組 (1-5名) 	//玩法群組 (6-10名)
							{
								//如果投注項目比規則中的項目少，就只比對投注的項目就好 (σﾟ∀ﾟ)σﾟ∀ﾟ)σ
								if len(thisBOContent.(map[string]interface{})) < ContentMax {
									ContentMax = len(thisBOContent.(map[string]interface{}))
								}
								// fmt.Println("<====> LotteryPlayGroup  ", LotteryPlayGroup)
								// fmt.Println("<====> LotteryPlay  ", LotteryPlay)
								// fmt.Println("<====> thisBOContent  ", thisBOContent)
								// fmt.Println("<====> betCount  ", thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{}))
								//fmt.Println("<====> betCount  ", thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)])
								for i := 0; i < ContentMax; i++ {
									//	fmt.Println("<====> thisBOContent.(map[string]interface{})[i+1]  ", thisBOContent.(map[string]interface{})[i+1])

									// fmt.Println("<====> thisBOContent.(map[string]interface{})[strconv.Itoa(i+1)]  ", thisBOContent.(map[string]interface{})[strconv.Itoa(i)], thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)])
									// fmt.Println(reflect.TypeOf(thisBOContent.(map[string]interface{})[strconv.Itoa(i)]))
									//fmt.Printf(reflect.TypeOf(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)]))
									if thisBOContent.(map[string]interface{})[strconv.Itoa(i)] == thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].(string) {
										// fmt.Println("QQQQQQQQQQQQQQQQQQ  ")
										Inspect["Ratio"].([]float64)[0]++
										Inspect["Message"] = 1
									}
								}
							}
						default:
							Inspect["State"] = false
							Inspect["Info"] = "Lottery Play Group Error"
							break

						}
					}

				case 2: //官方模式
					{
						//玩法群組 (因為有對應 TS_LotteryPlayGroup.LPG_Code，所以需要跟資料庫一起變動...ry)
						switch LotteryPlayGroup {
						//玩法群組 (前一)
						case 1:
							{
								//玩法 (因為有對應 TS_LotteryPlay.LP_Code，所以需要跟資料庫一起變動...ry)
								switch LotteryPlay {

								//玩法 (直选复式)
								case 1:
									{
										ContentAll := betCount.PKTenOneStraightBet(thisBOContent.(map[string]interface{})["1"].([]interface{}))
										for i := 0; i < len(ContentAll); i++ {
											tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].(string)

											// fmt.Println("<====> tempResult  ", tempResult)
											// fmt.Println("<====> ContentAll[i].([]string)[0]  ", ContentAll[i].([]string)[0])

											// fmt.Println("<====> ContentAll[i].([]string)[0]  ", ContentAll[i].([]string)[0] == tempResult)
											// fmt.Println("<====> strings.EqualFold(value1, value2)  ", strings.EqualFold(tempResult, ContentAll[i].([]string)[0]))
											// fmt.Println(reflect.TypeOf(tempResult))
											// fmt.Println(reflect.TypeOf(ContentAll[i].([]string)[0]))
											if ContentAll[i].([]string)[0] == tempResult {

												// fmt.Println("QQQQQ ")
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											}
										}
									}
								//玩法錯誤
								default:
									Inspect["State"] = false
									Inspect["Info"] = "Lottery Play  Error"
								}

							}
						//玩法群組 (前二)
						case 2:
							{
								//玩法 (因為有對應 TS_LotteryPlay.LP_Code，所以需要跟資料庫一起變動...ry)
								switch LotteryPlay {
								//玩法 (直选复式) //玩法 (直选单式)
								case 1, 2:
									{
										var ContentAll []interface{}
										if LotteryPlay == 1 {
											ContentAll = betCount.PKTenTwoStraightBet(thisBOContent.(map[string]interface{})["1"].([]interface{}), thisBOContent.(map[string]interface{})["2"].([]interface{}))
										}
										if LotteryPlay == 2 {
											ContentAll = betCount.PKTenTwoStraightSimpleBet(thisBOContent.(map[string]interface{})["text"].(string))
										}

										for i := 0; i < len(ContentAll); i++ {
											tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)

											// // fmt.Println("<====> ContentAll[i].([]string)[0]  ", ContentAll[i].([]string)[0])
											// // fmt.Println("<====> ContentAll[i].([]string)[0]  ", tempResult[0])
											// // fmt.Println("<====> ContentAll[i].([]string)[1]  ", ContentAll[i].([]string)[1])
											// // fmt.Println("<====> ContentAll[i].([]string)[1]  ", tempResult[0])
											// // fmt.Println("  ", ContentAll[i].([]string)[0] == tempResult[0] &&
											// 	ContentAll[i].([]string)[1] == tempResult[1])
											if ContentAll[i].([]string)[0] == tempResult[0] &&
												ContentAll[i].([]string)[1] == tempResult[1] {
												// fmt.Println(" CCCCCCCCCCC ")
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											}
										}
									}
								//玩法錯誤
								default:

									Inspect["State"] = false
									Inspect["Info"] = "Lottery Play  Error"

								}
							}
							//玩法群組 (前三)
						case 3:
							{
								//玩法 (因為有對應 TS_LotteryPlay.LP_Code，所以需要跟資料庫一起變動...ry)
								switch LotteryPlay {

								//玩法 (直选复式) //玩法 (直选单式)
								case 1, 2:
									{
										var ContentAll []interface{}
										if LotteryPlay == 1 {
											ContentAll = betCount.PKTenThreeStraightBet(thisBOContent.(map[string]interface{})["1"].([]interface{}), thisBOContent.(map[string]interface{})["2"].([]interface{}), thisBOContent.(map[string]interface{})["3"].([]interface{}))
										}
										if LotteryPlay == 2 {
											ContentAll = betCount.PKTenThreeStraightSimpleBet(thisBOContent.(map[string]interface{})["text"].(string))
										}

										for i := 0; i < len(ContentAll); i++ {
											tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)
											// fmt.Println("====>  玩法群組 (前三) ", ContentAll[i])
											// fmt.Println("====>  玩法群組 (前三) ", ContentAll[i].([]string)[0])
											// fmt.Println("====>   ", ContentAll[i].([]string)[1])
											// fmt.Println("====>   ", ContentAll[i].([]string)[2])
											// fmt.Println("====>   ContentAll", ContentAll)

											// fmt.Println("====>   tempResult", ContentAll[i].([]string)[0] == tempResult[0])
											// fmt.Println("====>   tempResult", ContentAll[i].([]string)[1] == tempResult[1])
											// fmt.Println("====>   tempResult", ContentAll[i].([]string)[2] == tempResult[2])
											if ContentAll[i].([]string)[0] == tempResult[0] &&
												ContentAll[i].([]string)[1] == tempResult[1] &&
												ContentAll[i].([]string)[2] == tempResult[2] {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											}
										}
									}
								//玩法錯誤
								default:

									Inspect["State"] = false
									Inspect["Info"] = "Lottery Play  Error"

								}
							}
							//玩法群組 (定位胆)
						case 4:
							{
								// fmt.Println("玩法群組 (定位胆)  ")
								//玩法 (因為有對應 TS_LotteryPlay.LP_Code，所以需要跟資料庫一起變動...ry)
								switch LotteryPlay {

								//玩法 (定位胆)
								case 1:
									{
										// fmt.Println("玩法 (定位胆) ")
										//因為不特定的下注項目有可能不存在，所以先預設為空陣列以免被警告或提示訊息噴好噴滿…
										if _, exists := thisBOContent.(map[string]interface{})["1"]; !exists {
											thisBOContent.(map[string]interface{})["1"] = []string{}
										}
										if _, exists := thisBOContent.(map[string]interface{})["2"]; !exists {
											thisBOContent.(map[string]interface{})["2"] = []string{}
										}
										if _, exists := thisBOContent.(map[string]interface{})["3"]; !exists {
											thisBOContent.(map[string]interface{})["3"] = []string{}
										}
										if _, exists := thisBOContent.(map[string]interface{})["4"]; !exists {
											thisBOContent.(map[string]interface{})["4"] = []string{}
										}
										if _, exists := thisBOContent.(map[string]interface{})["5"]; !exists {
											thisBOContent.(map[string]interface{})["5"] = []string{}
										}
										if _, exists := thisBOContent.(map[string]interface{})["6"]; !exists {
											thisBOContent.(map[string]interface{})["6"] = []string{}
										}
										if _, exists := thisBOContent.(map[string]interface{})["7"]; !exists {
											thisBOContent.(map[string]interface{})["7"] = []string{}
										}
										if _, exists := thisBOContent.(map[string]interface{})["8"]; !exists {
											thisBOContent.(map[string]interface{})["8"] = []string{}
										}
										if _, exists := thisBOContent.(map[string]interface{})["9"]; !exists {
											thisBOContent.(map[string]interface{})["9"] = []string{}
										}
										if _, exists := thisBOContent.(map[string]interface{})["10"]; !exists {
											thisBOContent.(map[string]interface{})["10"] = []string{}
										}
										// fmt.Println("<====> ContentAll[i].([]string)[0]  ", thisBOContent.(map[string]interface{})["1"])
										ContentAll := betCount.PKTenAllStraightBet(thisBOContent.(map[string]interface{})["1"].([]interface{}), thisBOContent.(map[string]interface{})["2"].([]interface{}), thisBOContent.(map[string]interface{})["3"].([]interface{}), thisBOContent.(map[string]interface{})["4"].([]interface{}), thisBOContent.(map[string]interface{})["5"].([]interface{}),
											thisBOContent.(map[string]interface{})["6"].([]interface{}), thisBOContent.(map[string]interface{})["7"].([]interface{}), thisBOContent.(map[string]interface{})["8"].([]interface{}), thisBOContent.(map[string]interface{})["9"].([]interface{}), thisBOContent.(map[string]interface{})["10"].([]interface{}))
										// //使用迴圈對獎，依下注的組合逐筆驗證
										// fmt.Println("<====> ContentAll[i].([]string)[0]  ", ContentAll)
										for i := 0; i < len(ContentAll); i++ {
											idx, _ := strconv.Atoi(ContentAll[i].([]string)[1])

											// fmt.Println("<====> ContentAll[i].([]string)[0] => ", ContentAll[i].([]string)[0])
											// fmt.Println("<====> ContentAll[i].([]string)[0]  ", thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)[idx])
											if ContentAll[i].([]string)[0] == thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)[idx] {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											}
										}

									}
								//玩法錯誤
								default:

									Inspect["State"] = false
									Inspect["Info"] = "Lottery Play  Error"

								}
							}
						default:
							Inspect["State"] = false
							Inspect["Info"] = "Lottery Play Group Error"
							break
						}
					}
				}

			default:
				Inspect["State"] = false
				Inspect["Info"] = "Lottery Play Mode Group Error"
			}
		}

	default:
		Inspect["State"] = false
	}

	return Inspect
}
func InspectModuleShiShi(LotteryTypeGroup int,
	LotteryType int,
	LotteryMode int,
	LotteryPlayGroup int,
	LotteryPlay int,
	thisBOContent interface{},
	thisFullResult map[string]interface{},
	config map[string]interface{}) map[string]interface{} {

	//初始化
	Inspect := make(map[string]interface{})
	//對獎結果 (0：沒中獎|1:中獎|2:和局退本金)
	Inspect["Message"] = 0
	//中獎金額倍數 (先預設為0)(簡單的說，就是嬴的注數…)
	Inspect["Ratio"] = []float64{0, 0, 0, 0}
	// //中獎賠率 (預設為1，即使用第一個)
	// CWCheck["Odds"] = 1;
	//執行狀態
	Inspect["State"] = true
	// //執行訊息
	Inspect["Info"] = ""

	ContentMax := int(config["LTR_ContentMax"].(map[string]interface{})[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].(float64))
	// fmt.Println("====>LotteryTypeGroup  ", LotteryTypeGroup)
	// fmt.Println("====>LotteryType  ", LotteryType)
	// fmt.Println("====>LotteryMode  ", LotteryMode)
	// fmt.Println("====>LotteryPlayGroup  ", LotteryPlayGroup)
	// fmt.Println("====>LotteryPlay  ", LotteryPlay)
	// fmt.Println("====>ContentMax  ", config["LTR_ContentMax"].(map[string]interface{})[strconv.Itoa(LotteryMode)])
	// fmt.Println("====>ContentMax  ", ContentMax)

	switch LotteryTypeGroup {
	//时时彩
	case 2:
		{
			switch LotteryType {
			//重庆时时彩 //新疆时时彩 //天津时时彩 //精彩1分彩 //精彩3分彩 //精彩5分彩 //精彩秒秒彩
			case 2, 3, 4, 26, 27, 28, 29, 30: // 30 區塊鏈時時彩

				// fmt.Println("InspectModuleShiShi   ", LotteryMode)
				// fmt.Println("<====> thisBOContent  ", thisBOContent)
				// fmt.Println("<====> thisFullResult  ", thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)])
				// fmt.Println("<====> thisFullResult  ", thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)])
				// fmt.Println("<====> thisFullResult  ", thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)])
				//玩法群組模式 (因為有對應 TS_LotteryPlayGroup.LPG_Mode，所以需要跟資料庫一起變動...ry)
				switch LotteryMode {
				//传统模式 (信用模式)
				case 1:
					{
						//玩法群組 (因為有對應 TS_LotteryPlayGroup.LPG_Code，所以需要跟資料庫一起變動...ry)
						switch LotteryPlayGroup {
						//玩法群組 (两面) //玩法群組 (前中后)

						case 1, 3:
							// fmt.Println("玩法群組 (两面) //玩法群組 (前中后)   ")
							{
								// fmt.Println("<====> len(thisBOContent.(map[string]interface{}))   ", len(thisBOContent.(map[string]interface{})))
								// fmt.Println("<====> ContentMax  ", ContentMax)
								//如果投注項目比規則中的項目少，就只比對投注的項目就好 (σﾟ∀ﾟ)σﾟ∀ﾟ)σ
								if len(thisBOContent.(map[string]interface{})) < ContentMax {
									ContentMax = len(thisBOContent.(map[string]interface{}))
								}
								// 使用迴圈對獎，防止注單中可能不只有一項
								// fmt.Println("<====> 玩法群組 (两面)   ", thisBOContent)

								for i := 0; i < ContentMax; i++ {

									// betCount.BetCountOfficial(1, 1, 1, thisBOContent)
									//fmt.Println("玩法群組 (两面) //玩法群組 (前中后) ", thisBOContent.(map[string]interface{})[strconv.Itoa(i)])
									// fmt.Println("<====> LotteryPlay  ", LotteryPlay)

									// fmt.Println("<====> LotteryPlayGroup  ", LotteryPlayGroup)
									// fmt.Println("<====> thisBOContent  ", thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{}))

									// fmt.Println("<====> thisFullResult ", thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)])
									// betCount.In_array(thisBOContent.(map[string]interface{})[strconv.Itoa(i)], thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)])
									if betCount.In_array(thisBOContent.(map[string]interface{})[strconv.Itoa(i)], thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)]) {
										//	Inspect["Ratio"].(int)[0] = Inspect["Ratio"].(int)[0] + 1
										// fmt.Println("<====> betCount  ")
										Inspect["Ratio"].([]float64)[0]++
										Inspect["Message"] = 1
									}

									//fmt.Println("Inspect ", Inspect)
								}
								// break;
							}
						//玩法群組 (1-5球)
						case 2:
							// fmt.Println("玩法群組 (1-5球)   ")
							{
								// fmt.Println("<====> Conte法群組 (1-5球)   ", ContentMax)
								//fmt.Println("<====> Conte法群組 (1-5球)   ", thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)])
								// fmt.Println("<====> Conte法群組 (1-5球)   ", reflect.TypeOf(thisBOContent.(map[string]interface{})[strconv.Itoa(0)]))
								// fmt.Println("<====> Conte法群組 (1-5球)   ", reflect.TypeOf(strconv.Itoa(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].(int))))

								// fmt.Println("<====> Conte法群組 (1-5球)   ", thisBOContent.(map[string]interface{})[strconv.Itoa(0)])
								// fmt.Println("<====> Conte法群組 (1-5球)   ", strconv.Itoa(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].(int)))
								// fmt.Println("<====> thisBOContent  ", thisBOContent.(map[string]interface{})[strconv.Itoa(i)])
								//如果投注項目比規則中的項目少，就只比對投注的項目就好 (σﾟ∀ﾟ)σﾟ∀ﾟ)σ
								if len(thisBOContent.(map[string]interface{})) < ContentMax {
									ContentMax = len(thisBOContent.(map[string]interface{}))
								}
								for i := 0; i < ContentMax; i++ {
									if thisBOContent.(map[string]interface{})[strconv.Itoa(i)] == thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] {
										//	Inspect["Ratio"].(int)[0] = Inspect["Ratio"].(int)[0] + 1
										// fmt.Println("<====> Conte法群組 (1-5球)   ", ContentMax)
										Inspect["Ratio"].([]float64)[0]++
										Inspect["Message"] = 1
									}
								}

								//fmt.Println("Inspect ", Inspect)

							}
						default:
							Inspect["State"] = false
							Inspect["Info"] = "Lottery Play Group Error"
							break
						}

					}
					//官方模式
				case 2:
					{
						switch LotteryPlayGroup {
						//玩法群組 (定位胆)
						case 1:
							{
								//玩法 (因為有對應 TS_LotteryPlay.LP_Code，所以需要跟資料庫一起變動...ry)
								switch LotteryPlay {
								//玩法 (直选复式)
								case 1:
									{
										//因為不特定的下注項目有可能不存在，所以先預設為空陣列以免被警告或提示訊息噴好噴滿…

										if _, exists := thisBOContent.(map[string]interface{})["1"]; !exists {
											thisBOContent.(map[string]interface{})["1"] = []string{}
										}
										if _, exists := thisBOContent.(map[string]interface{})["2"]; !exists {
											thisBOContent.(map[string]interface{})["2"] = []string{}
										}
										if _, exists := thisBOContent.(map[string]interface{})["3"]; !exists {
											thisBOContent.(map[string]interface{})["3"] = []string{}
										}
										if _, exists := thisBOContent.(map[string]interface{})["4"]; !exists {
											thisBOContent.(map[string]interface{})["4"] = []string{}
										}
										if _, exists := thisBOContent.(map[string]interface{})["5"]; !exists {
											thisBOContent.(map[string]interface{})["5"] = []string{}
										}

										//依下注的項目列出全部的組合
										ContentAll := betCount.ShiShiStraightBet(thisBOContent.(map[string]interface{})["1"].([]interface{}), thisBOContent.(map[string]interface{})["2"].([]interface{}), thisBOContent.(map[string]interface{})["3"].([]interface{}), thisBOContent.(map[string]interface{})["4"].([]interface{}), thisBOContent.(map[string]interface{})["5"].([]interface{}))
										for i := 0; i < len(ContentAll); i++ {

											// fmt.Println(" a* ContentAll[i]    => ", ContentAll[i])

											// fmt.Println(" thisFullResult ", thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)])
											idx, _ := strconv.Atoi(ContentAll[i].([]string)[1])
											if ContentAll[i].([]string)[0] == thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)[idx] {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1

											}
										}
									}
								default:
									Inspect["State"] = false
									Inspect["Info"] = "Lottery Play Group Error"
									break

								}

							}

							//玩法群組 (五星)
						case 2:
							{
								//玩法 (因為有對應 TS_LotteryPlay.LP_Code，所以需要跟資料庫一起變動...ry)
								switch LotteryPlay {
								//玩法 (直选复式) //玩法 (直选单式)
								case 1, 8:
									{
										var ContentAll []interface{}
										//依下注的項目列出全部的組合
										if LotteryPlay == 1 {
											// fmt.Println("BC 依下注的項目列出全部的組合  =  ", reflect.TypeOf(thisBOContent.(map[string]interface{})["1"]))
											// fmt.Println("BC TypeOf  =  ", thisBOContent)
											ContentAll = betCount.ShiShiFiveStarBet(thisBOContent.(map[string]interface{})["1"].([]interface{}), thisBOContent.(map[string]interface{})["2"].([]interface{}), thisBOContent.(map[string]interface{})["3"].([]interface{}), thisBOContent.(map[string]interface{})["4"].([]interface{}), thisBOContent.(map[string]interface{})["5"].([]interface{}))
											// fmt.Println("ContentAll ===> ", ContentAll)
										} else if LotteryPlay == 8 {
											ContentAll = betCount.ShiShiFiveStarSimpleBet(thisBOContent.(map[string]interface{})["text"].(string))
										}

										for i := 0; i < len(ContentAll); i++ {
											tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)

											if ContentAll[i].([]string)[0] == tempResult[0] &&
												ContentAll[i].([]string)[1] == tempResult[1] &&
												ContentAll[i].([]string)[2] == tempResult[2] &&
												ContentAll[i].([]string)[3] == tempResult[3] &&
												ContentAll[i].([]string)[4] == tempResult[4] {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											}
										}

									}
									//玩法 (组选120)
								case 2:
									{
										//先檢查開獎結果是否符合目前的玩法規則
										if len(betCount.Array_unique_str(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))) == 5 {

											//依下注的項目列出全部的組合
											ContentAll := betCount.ShiShiFiveStarCombo120Bet(thisBOContent.(map[string]interface{})["1"].([]interface{}))
											tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)]
											for i := 0; i < len(ContentAll); i++ {
												if len(betCount.Array_unique_str(ContentAll[i].([]string))) == 5 &&
													betCount.In_array(ContentAll[i].([]string)[0], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[1], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[2], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[3], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[4], tempResult) {
													Inspect["Ratio"].([]float64)[0]++
													Inspect["Message"] = 1
												}
											}
										}
									}
									//玩法 (组选60)
								case 3:
									{
										//先檢查開獎結果是否符合目前的玩法規則
										if len(betCount.Array_unique_str(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))) == 4 {

											//依下注的項目列出全部的組合
											ContentAll := betCount.ShiShiFiveStarCombo60Bet(thisBOContent.(map[string]interface{})["1"].([]interface{}), thisBOContent.(map[string]interface{})["2"].([]interface{}))
											for i := 0; i < len(ContentAll); i++ {
												tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)]
												if len(betCount.Array_unique_str(ContentAll[i].([]string))) == 4 &&
													betCount.In_array(ContentAll[i].([]string)[0], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[1], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[2], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[3], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[4], tempResult) {

													ResultTemp := betCount.Dup_count(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))

													ResultSorted := betCount.SortMapByValue(ResultTemp)

													// sort.Sort(sort.Reverse(sort.StringSlice(ResultTemp)))

													for _, v := range ResultSorted {
														if v.Key == ContentAll[i].([]string)[0] && v.Value == 2 {
															Inspect["Ratio"].([]float64)[0]++
															Inspect["Message"] = 1

														}
													}
												}
											}

										}
									}

								//玩法 (组选30)
								case 4:
									{
										//先檢查開獎結果是否符合目前的玩法規則
										if len(betCount.Array_unique_str(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))) == 3 {

											//依下注的項目列出全部的組合
											ContentAll := betCount.ShiShiFiveStarCombo30Bet(thisBOContent.(map[string]interface{})["1"].([]interface{}), thisBOContent.(map[string]interface{})["2"].([]interface{}))

											// fmt.Println("ContentAll  content =  ", ContentAll)
											//使用迴圈對獎，依下注的組合逐筆驗證
											for i := 0; i < len(ContentAll); i++ {
												tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)]
												if len(betCount.Array_unique_str(ContentAll[i].([]string))) == 3 &&
													betCount.In_array(ContentAll[i].([]string)[0], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[1], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[2], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[3], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[4], tempResult) {

													ResultTemp := betCount.Dup_count(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))
													ResultSorted := betCount.SortMapByValue(ResultTemp)

													check1 := false
													check2 := false
													for _, v := range ResultSorted {
														if v.Key == ContentAll[i].([]string)[0] && v.Value == 2 {
															check1 = true
														} else if v.Key == ContentAll[i].([]string)[2] && v.Value == 2 {
															check2 = true
														}

													}
													if check1 && check2 {
														Inspect["Ratio"].([]float64)[0]++
														Inspect["Message"] = 1

													}

												}
											}

										}
									}
								//玩法 (组选20)
								case 5:
									{
										//先檢查開獎結果是否符合目前的玩法規則
										if len(betCount.Array_unique_str(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))) == 3 {
											//依下注的項目列出全部的組合
											ContentAll := betCount.ShiShiFiveStarCombo20Bet(thisBOContent.(map[string]interface{})["1"].([]interface{}), thisBOContent.(map[string]interface{})["2"].([]interface{}))
											//使用迴圈對獎，依下注的組合逐筆驗證
											for i := 0; i < len(ContentAll); i++ {
												tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)]
												if len(betCount.Array_unique_str(ContentAll[i].([]string))) == 3 &&
													betCount.In_array(ContentAll[i].([]string)[0], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[1], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[2], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[3], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[4], tempResult) {

													ResultTemp := betCount.Dup_count(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))
													ResultSorted := betCount.SortMapByValue(ResultTemp)

													for _, v := range ResultSorted {
														if v.Key == ContentAll[i].([]string)[0] && v.Value == 3 {
															Inspect["Ratio"].([]float64)[0]++
															Inspect["Message"] = 1

														}
													}
												}
											}
										}
									}
									//玩法 (组选10)
								case 6:
									{
										//先檢查開獎結果是否符合目前的玩法規則
										if len(betCount.Array_unique_str(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))) == 2 {
											//依下注的項目列出全部的組合
											ContentAll := betCount.ShiShiFiveStarCombo10Bet(thisBOContent.(map[string]interface{})["1"].([]interface{}), thisBOContent.(map[string]interface{})["2"].([]interface{}))
											//使用迴圈對獎，依下注的組合逐筆驗證
											for i := 0; i < len(ContentAll); i++ {
												tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)]
												if len(betCount.Array_unique_str(ContentAll[i].([]string))) == 2 &&
													betCount.In_array(ContentAll[i].([]string)[0], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[1], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[2], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[3], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[4], tempResult) {

													ResultTemp := betCount.Dup_count(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))
													ResultSorted := betCount.SortMapByValue(ResultTemp)
													check1 := false
													check2 := false
													for _, v := range ResultSorted {
														if v.Key == ContentAll[i].([]string)[0] && v.Value == 3 {
															check1 = true
														} else if v.Key == ContentAll[i].([]string)[4] && v.Value == 2 {
															check2 = true
														}

													}
													if check1 && check2 {
														Inspect["Ratio"].([]float64)[0]++
														Inspect["Message"] = 1

													}

												}
											}
										}
									}
								//玩法 (组选5)
								case 7:
									{
										//先檢查開獎結果是否符合目前的玩法規則
										if len(betCount.Array_unique_str(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))) == 2 {
											//依下注的項目列出全部的組合
											ContentAll := betCount.ShiShiFiveStarCombo5Bet(thisBOContent.(map[string]interface{})["1"].([]interface{}), thisBOContent.(map[string]interface{})["2"].([]interface{}))
											//使用迴圈對獎，依下注的組合逐筆驗證
											for i := 0; i < len(ContentAll); i++ {
												tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)]
												if len(betCount.Array_unique_str(ContentAll[i].([]string))) == 2 &&
													betCount.In_array(ContentAll[i].([]string)[0], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[1], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[2], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[3], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[4], tempResult) {

													ResultTemp := betCount.Dup_count(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))
													ResultSorted := betCount.SortMapByValue(ResultTemp)

													for _, v := range ResultSorted {
														if v.Key == ContentAll[i].([]string)[0] && v.Value == 4 {
															Inspect["Ratio"].([]float64)[0]++
															Inspect["Message"] = 1

														}
													}

												}
											}
										}
									}
								default:
									Inspect["State"] = false
									Inspect["Info"] = "Lottery Play Group Mode Error"

								}

							}
							//玩法群組 (四星)
						case 3:
							{
								//玩法 (因為有對應 TS_LotteryPlay.LP_Code，所以需要跟資料庫一起變動...ry)
								switch LotteryPlay {
								//玩法 (直选复式) //玩法 (直选单式)
								case 1, 6:
									{
										var ContentAll []interface{}
										//依下注的項目列出全部的組合
										if LotteryPlay == 1 {
											ContentAll = betCount.ShiShiFourStarBet(thisBOContent.(map[string]interface{})["1"].([]interface{}), thisBOContent.(map[string]interface{})["2"].([]interface{}), thisBOContent.(map[string]interface{})["3"].([]interface{}), thisBOContent.(map[string]interface{})["4"].([]interface{}))
										} else if LotteryPlay == 6 {
											ContentAll = betCount.ShiShiFourStarSimpleBet(thisBOContent.(map[string]interface{})["text"].(string))
										}
										for i := 0; i < len(ContentAll); i++ {
											tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)
											if ContentAll[i].([]string)[0] == tempResult[0] &&
												ContentAll[i].([]string)[1] == tempResult[1] &&
												ContentAll[i].([]string)[2] == tempResult[2] &&
												ContentAll[i].([]string)[3] == tempResult[3] {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											}
										}

									}
								//玩法 (组选24)
								case 2:
									{
										//先檢查開獎結果是否符合目前的玩法規則
										if len(betCount.Array_unique_str(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))) == 4 {
											//依下注的項目列出全部的組合
											ContentAll := betCount.ShiShiFourStarCombo24Bet(thisBOContent.(map[string]interface{})["1"].([]interface{}))
											//使用迴圈對獎，依下注的組合逐筆驗證
											for i := 0; i < len(ContentAll); i++ {

												tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)
												if len(betCount.Array_unique_str(ContentAll[i].([]string))) == 4 &&
													betCount.In_array(ContentAll[i].([]string)[0], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[1], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[2], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[3], tempResult) {
													Inspect["Ratio"].([]float64)[0]++
													Inspect["Message"] = 1
												}

											}

										}
									}
								//玩法 (组选12)
								case 3:
									{
										//先檢查開獎結果是否符合目前的玩法規則
										if len(betCount.Array_unique_str(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))) == 3 {
											//依下注的項目列出全部的組合
											ContentAll := betCount.ShiShiFourStarCombo12Bet(thisBOContent.(map[string]interface{})["1"].([]interface{}), thisBOContent.(map[string]interface{})["2"].([]interface{}))
											//使用迴圈對獎，依下注的組合逐筆驗證
											for i := 0; i < len(ContentAll); i++ {

												tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)
												if betCount.In_array(ContentAll[i].([]string)[0], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[1], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[2], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[3], tempResult) {

													ResultTemp := betCount.Dup_count(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))
													ResultSorted := betCount.SortMapByValue(ResultTemp)

													for _, v := range ResultSorted {
														if v.Key == ContentAll[i].([]string)[0] && v.Value == 2 {
															Inspect["Ratio"].([]float64)[0]++
															Inspect["Message"] = 1

														}
													}

												}

											}

										}
									}
								//玩法 (组选6)
								case 4:
									{
										//先檢查開獎結果是否符合目前的玩法規則
										if len(betCount.Array_unique_str(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))) == 2 {
											//依下注的項目列出全部的組合
											ContentAll := betCount.ShiShiFourStarCombo6Bet(thisBOContent.(map[string]interface{})["1"].([]interface{}))
											tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)]
											//使用迴圈對獎，依下注的組合逐筆驗證
											for i := 0; i < len(ContentAll); i++ {

												if betCount.In_array(ContentAll[i].([]string)[0], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[1], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[2], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[3], tempResult) {

													ResultTemp := betCount.Dup_count(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))
													ResultSorted := betCount.SortMapByValue(ResultTemp)
													check1 := false
													check2 := false
													for _, v := range ResultSorted {
														if v.Key == ContentAll[i].([]string)[0] && v.Value == 2 {
															check1 = true
														} else if v.Key == ContentAll[i].([]string)[2] && v.Value == 2 {
															check2 = true
														}

													}
													if check1 && check2 {
														Inspect["Ratio"].([]float64)[0]++
														Inspect["Message"] = 1

													}

												}
											}
										}

									}
								//玩法 (组选4)
								case 5:
									{
										//先檢查開獎結果是否符合目前的玩法規則
										if len(betCount.Array_unique_str(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))) == 2 {

											//依下注的項目列出全部的組合
											ContentAll := betCount.ShiShiFourStarCombo4Bet(thisBOContent.(map[string]interface{})["1"].([]interface{}), thisBOContent.(map[string]interface{})["2"].([]interface{}))
											tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)]

											//使用迴圈對獎，依下注的組合逐筆驗證
											for i := 0; i < len(ContentAll); i++ {
												if betCount.In_array(ContentAll[i].([]string)[0], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[1], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[2], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[3], tempResult) {

													ResultTemp := betCount.Dup_count(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))
													ResultSorted := betCount.SortMapByValue(ResultTemp)
													for _, v := range ResultSorted {
														if v.Key == ContentAll[i].([]string)[0] && v.Value == 3 {
															Inspect["Ratio"].([]float64)[0]++
															Inspect["Message"] = 1
														}
													}
												}

											}

										}
									}
								default:
									Inspect["State"] = false
									Inspect["Info"] = "Lottery Play Group Mode Error"

								}
							}

						//玩法群組 (后三) 	//玩法群組 (前三) //玩法群組 (中三)
						case 4, 5, 12:
							{
								//玩法 (因為有對應 TS_LotteryPlay.LP_Code，所以需要跟資料庫一起變動...ry)
								switch LotteryPlay {

								//玩法 (直选复式) //玩法 (直选单式)
								case 1, 11:
									{
										var ContentAll []interface{}
										//依下注的項目列出全部的組合
										if LotteryPlay == 1 {
											ContentAll = betCount.ShiShiThreeStarBet(thisBOContent.(map[string]interface{})["1"].([]interface{}), thisBOContent.(map[string]interface{})["2"].([]interface{}), thisBOContent.(map[string]interface{})["3"].([]interface{}))
										} else if LotteryPlay == 11 {
											ContentAll = betCount.ShiShiThreeStarSimpleBet(thisBOContent.(map[string]interface{})["text"].(string))
										}
										tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)

										for i := 0; i < len(ContentAll); i++ {
											if ContentAll[i].([]string)[0] == tempResult[0] &&
												ContentAll[i].([]string)[1] == tempResult[1] &&
												ContentAll[i].([]string)[2] == tempResult[2] {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											}
										}

										//fmt.Println("Inspect       ", Inspect)
									}
								//玩法 (直选和值) //玩法 (直选跨度)
								case 2, 3:
									{
										var ContentAll []interface{}
										if LotteryPlay == 2 {
											ContentAll = betCount.ShiShiThreeStarSumBet(thisBOContent.(map[string]interface{})["1"].([]interface{}))
										}
										if LotteryPlay == 3 {
											ContentAll = betCount.ShiShiThreeStarCutResultBet(thisBOContent.(map[string]interface{})["1"].([]interface{}))
										}
										tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)
										for i := 0; i < len(ContentAll); i++ {

											if ContentAll[i].([]string)[0] == tempResult[0] &&
												ContentAll[i].([]string)[1] == tempResult[1] &&
												ContentAll[i].([]string)[2] == tempResult[2] {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											}
										}

									}
								//玩法 (后三组合|前三组合|中三组合)
								case 4:
									{
										ContentAll := betCount.ShiShiThreeStarComboBet(thisBOContent.(map[string]interface{})["1"].([]interface{}), thisBOContent.(map[string]interface{})["2"].([]interface{}), thisBOContent.(map[string]interface{})["3"].([]interface{}))
										tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)

										for i := 0; i < len(ContentAll); i++ {
											RatioTemp := 0

											if ContentAll[i].([]string)[0] == tempResult[0] {
												RatioTemp = RatioTemp + 1
											}
											if ContentAll[i].([]string)[1] == tempResult[1] {
												RatioTemp = RatioTemp + 1
											}
											if ContentAll[i].([]string)[2] == tempResult[2] {
												RatioTemp = RatioTemp + 1
											}

											if RatioTemp > 0 {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											}

										}
									}
								//玩法 (组三复式) 	//玩法 (組三单式)
								case 5, 12:
									{
										var ContentAll []interface{}
										if LotteryPlay == 5 {
											ContentAll = betCount.ShiShiThreeStarCombo3ComplexBet(thisBOContent.(map[string]interface{})["1"].([]interface{}))
										}
										if LotteryPlay == 12 {
											ContentAll = betCount.ShiShiThreeStarCombo3SimpleBet(thisBOContent.(map[string]interface{})["text"].(string))
										}

										//先檢查開獎結果是否符合目前的玩法規則
										if len(betCount.Array_unique_str(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))) == 2 {
											tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)]

											for i := 0; i < len(ContentAll); i++ {

												if len(betCount.Array_unique_str(ContentAll[i].([]string))) == 2 &&
													betCount.In_array(ContentAll[i].([]string)[0], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[1], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[2], tempResult) {

													ResultTemp := betCount.Dup_count(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))
													ResultSorted := betCount.SortMapByValue(ResultTemp)

													ContentTemp := betCount.Dup_count(ContentAll[i].([]string))
													ContentSorted := betCount.SortMapByValue(ContentTemp)

													for _, v := range ResultSorted {
														if v.Key == ContentSorted[0].Key && v.Value == 2 {
															Inspect["Ratio"].([]float64)[0]++
															Inspect["Message"] = 1
														}

													}

												}

											}

										}

									}
									//玩法 (组六复式) //玩法 (組六单式)
								case 6, 13:
									{
										var ContentAll []interface{}
										if LotteryPlay == 6 {
											ContentAll = betCount.ShiShiThreeStarCombo6ComplexBet(thisBOContent.(map[string]interface{})["1"].([]interface{}))
										}
										if LotteryPlay == 13 {
											ContentAll = betCount.ShiShiThreeStarCombo6SimpleBet(thisBOContent.(map[string]interface{})["text"].(string))
										}

										//先檢查開獎結果是否符合目前的玩法規則
										if len(betCount.Array_unique_str(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))) == 3 {
											tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)]

											for i := 0; i < len(ContentAll); i++ {
												if len(betCount.Array_unique_str(ContentAll[i].([]string))) == 3 &&
													betCount.In_array(ContentAll[i].([]string)[0], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[1], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[2], tempResult) {
													Inspect["Ratio"].([]float64)[0]++
													Inspect["Message"] = 1

												}
											}
										}

									}
									//玩法 (组选和值)
								case 7:
									{
										ContentAll := betCount.ShiShiThreeStarComboSumBet(thisBOContent.(map[string]interface{})["1"].([]interface{}))
										//先檢查開獎結果是否符合目前的玩法規則
										if len(betCount.Array_unique_str(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))) >= 2 {
											tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)]
											for i := 0; i < len(ContentAll); i++ {
												if betCount.ArraySum(ContentAll[i].([]string)) == betCount.ArraySum(tempResult) {

													if len(betCount.Array_unique_str(ContentAll[i].([]string))) == 2 &&
														betCount.In_array(ContentAll[i].([]string)[0], tempResult) &&
														betCount.In_array(ContentAll[i].([]string)[1], tempResult) &&
														betCount.In_array(ContentAll[i].([]string)[2], tempResult) {
														Inspect["Ratio"].([]float64)[0]++
														Inspect["Message"] = 1

													} else if len(betCount.Array_unique_str(ContentAll[i].([]string))) == 3 &&
														betCount.In_array(ContentAll[i].([]string)[0], tempResult) &&
														betCount.In_array(ContentAll[i].([]string)[1], tempResult) &&
														betCount.In_array(ContentAll[i].([]string)[2], tempResult) {
														Inspect["Ratio"].([]float64)[1]++
														Inspect["Message"] = 1

													}
												}

											}
										}

									}
									//玩法 (组选包胆)
								case 8:
									{
										if len(betCount.Array_unique_str(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))) >= 2 {
											//tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)]

											ContentAll := betCount.ShiShiThreeStarComboBaoDnoBet(thisBOContent.(map[string]interface{})["1"].([]interface{}))
											//先檢查開獎結果是否符合目前的玩法規則

											for i := 0; i < len(ContentAll); i++ {
												ResultTemp := betCount.Dup_count(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))
												ResultSorted := betCount.SortMapByValue(ResultTemp)

												ContentTemp := betCount.Dup_count(ContentAll[i].([]string))
												ContentSorted := betCount.SortMapByValue(ContentTemp)

												if reflect.DeepEqual(ResultSorted, ContentSorted) && len(ContentTemp) == 2 {
													Inspect["Ratio"].([]float64)[0]++
													Inspect["Message"] = 1
												} else if reflect.DeepEqual(ResultSorted, ContentSorted) && len(ContentTemp) == 3 {
													Inspect["Ratio"].([]float64)[1]++
													Inspect["Message"] = 1
												}

											}

										}

									}
								//玩法 (和值尾数)
								case 9:
									{
										//依下注的項目列出全部的組合
										ContentAll := thisBOContent.(map[string]interface{})["1"].([]interface{})
										//使用迴圈對獎，依下注的組合逐筆驗證
										for i := 0; i < len(ContentAll); i++ {
											if ContentAll[i].(string) == thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)] {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											}

										}
										break
									}
									//玩法 (特殊号) ********
								case 10:
									{
										//依下注的項目列出全部的組合
										ContentAll := thisBOContent.(map[string]interface{})["1"].([]interface{})
										//使用迴圈對獎，依下注的組合逐筆驗證
										tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)]

										for i := 0; i < len(ContentAll); i++ {
											if betCount.In_array(ContentAll[i].([]string), tempResult) {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											}
										}

									}

									//玩法 (混合组选)
								case 14:
									{
										//依下注的項目列出全部的組合
										ContentAll := betCount.ShiShiThreeStarComboMixBet(thisBOContent.(map[string]interface{})["text"].(string))
										//使用迴圈對獎，依下注的組合逐筆驗證
										tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)]
										for i := 0; i < len(ContentAll); i++ {
											//使用「組三单式」规则

											if len(betCount.Array_unique_str(ContentAll[i].([]string))) == 2 {

												if betCount.In_array(ContentAll[i].([]string)[0], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[1], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[2], tempResult) {

													ResultTemp := betCount.Dup_count(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))
													ResultSorted := betCount.SortMapByValue(ResultTemp)

													ContentTemp := betCount.Dup_count(ContentAll[i].([]string))
													ContentSorted := betCount.SortMapByValue(ContentTemp)

													for _, v := range ResultSorted {
														if v.Key == ContentSorted[0].Key && v.Value == 2 {
															Inspect["Ratio"].([]float64)[0]++
															Inspect["Message"] = 1
														}
													}
												}
											}
											if len(betCount.Array_unique_str(ContentAll[i].([]string))) == 3 {
												tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)]
												if betCount.In_array(ContentAll[i].([]string)[0], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[1], tempResult) &&
													betCount.In_array(ContentAll[i].([]string)[2], tempResult) {
													Inspect["Ratio"].([]float64)[1]++
													Inspect["Message"] = 1

												}
											}
										}
									}
								default:
									Inspect["State"] = false
									Inspect["Info"] = "Lottery Play Group Mode Error"

								}
							}

							//玩法群組 (前二) 	//玩法群組 (后二)
						case 6, 13:
							{
								//玩法 (因為有對應 TS_LotteryPlay.LP_Code，所以需要跟資料庫一起變動...ry)
								switch LotteryPlay {
								//玩法 (直选复式) //玩法 (直选单式)
								case 1, 7:
									{
										var ContentAll []interface{}
										if LotteryPlay == 1 {
											ContentAll = betCount.ShiShiTwoStarBet(thisBOContent.(map[string]interface{})["1"].([]interface{}), thisBOContent.(map[string]interface{})["2"].([]interface{}))
										}
										if LotteryPlay == 7 {
											ContentAll = betCount.ShiShiTwoStarSimpleBet(thisBOContent.(map[string]interface{})["text"].(string))
										}
										tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)
										//fmt.Println("		Inspect[Ratio]  ", tempResult)
										//使用迴圈對獎，依下注的組合逐筆驗證
										for i := 0; i < len(ContentAll); i++ {
											if ContentAll[i].([]string)[0] == tempResult[0] &&
												ContentAll[i].([]string)[1] == tempResult[1] {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											}
										}
										//fmt.Println("		Inspect[Ratio]  ", Inspect["Ratio"])

									}
									//玩法 (直选和值) //玩法 (直选跨度)
								case 2, 3:
									{
										var ContentAll []interface{}
										if LotteryPlay == 2 {
											ContentAll = betCount.ShiShiTwoStarSumBet(thisBOContent.(map[string]interface{})["1"].([]interface{}))
										}
										if LotteryPlay == 3 {
											ContentAll = betCount.ShiShiTwoStarCutResultBet(thisBOContent.(map[string]interface{})["1"].([]interface{}))
										}

										tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)

										//使用迴圈對獎，依下注的組合逐筆驗證
										for i := 0; i < len(ContentAll); i++ {
											if ContentAll[i].([]string)[0] == tempResult[0] &&
												ContentAll[i].([]string)[1] == tempResult[1] {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											}
										}
									}
								//玩法 (组选复式)) //玩法 (組选单式)
								case 4, 8:
									{
										var ContentAll []interface{}
										if LotteryPlay == 4 {
											ContentAll = betCount.ShiShiTwoStarComboComplexBet(thisBOContent.(map[string]interface{})["1"].([]interface{}))
										}
										if LotteryPlay == 8 {
											ContentAll = betCount.ShiShiTwoStarComboSimpleBet(thisBOContent.(map[string]interface{})["text"].(string))
										}

										tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)

										//使用迴圈對獎，依下注的組合逐筆驗證
										for i := 0; i < len(ContentAll); i++ {

											if betCount.In_array(ContentAll[i].([]string)[0], tempResult) &&
												betCount.In_array(ContentAll[i].([]string)[1], tempResult) {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1

											}

										}

									}
								//玩法 (组选和值)
								case 5:
									{
										tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)

										if len(betCount.Array_unique_str(tempResult)) == 2 {
											ContentAll := betCount.ShiShiTwoStarComboSumBet(thisBOContent.(map[string]interface{})["1"].([]interface{}))
											//使用迴圈對獎，依下注的組合逐筆驗證
											for i := 0; i < len(ContentAll); i++ {
												if betCount.ArraySum(ContentAll[i].([]string)) == betCount.ArraySum(tempResult) {
													if betCount.In_array(ContentAll[i].([]string)[0], tempResult) &&
														betCount.In_array(ContentAll[i].([]string)[1], tempResult) {
														Inspect["Ratio"].([]float64)[0]++
														Inspect["Message"] = 1
													}
												}
											}
										}
									}
									//玩法 (组选包胆)
								case 6:
									{
										tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)

										if len(betCount.Array_unique_str(tempResult)) == 2 {
											ContentAll := betCount.ShiShiTwoStarComboBaoDnoBet(thisBOContent.(map[string]interface{})["1"].([]interface{}))
											for i := 0; i < len(ContentAll); i++ {

												ResultTemp := betCount.Dup_count(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))
												ResultSorted := betCount.SortMapByValue(ResultTemp)

												ContentTemp := betCount.Dup_count(ContentAll[i].([]string))
												ContentSorted := betCount.SortMapByValue(ContentTemp)

												if reflect.DeepEqual(ResultSorted, ContentSorted) {
													Inspect["Ratio"].([]float64)[0]++
													Inspect["Message"] = 1

												}

											}
										}
									}
								default:
									Inspect["State"] = false
									Inspect["Info"] = "Lottery Play Group Mode Error"

								}
							}
							//玩法群組 (不定位)
						case 7:
							{

								//玩法 (因為有對應 TS_LotteryPlay.LP_Code，所以需要跟資料庫一起變動...ry)
								switch LotteryPlay {
								//玩法 (前三一码) //玩法 (后三一码) //玩法 (中三一码) //玩法 (前四一码) //玩法 (后四一码) //玩法 (五星一码)
								case 1, 3, 12, 5, 7, 9:
									{
										ContentAll := thisBOContent.(map[string]interface{})["1"].([]interface{})

										// fmt.Println("	ContentAll  )", ContentAll)
										tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)
										for i := 0; i < len(ContentAll); i++ {
											if betCount.In_array2(ContentAll[i].(string), tempResult) {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											}

										}

									}
								//玩法 (前三二码) 	//玩法 (后三二码) //玩法 (中三二码) //玩法 (前四二码) //玩法 (后四二码) //玩法 (五星二码)
								case 2, 4, 13, 6, 8, 10:
									{
										ContentAll := betCount.ShiShiAnyPositionPick2Bet(thisBOContent.(map[string]interface{})["1"].([]interface{}))
										tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)

										for i := 0; i < len(ContentAll); i++ {
											if len(betCount.Array_unique_str(ContentAll[i].([]string))) == 2 &&
												betCount.In_array(ContentAll[i].([]string)[0], tempResult) &&
												betCount.In_array(ContentAll[i].([]string)[1], tempResult) {

												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											}

										}
									}
								//玩法 (五星三码)
								case 11:
									{
										//依下注的項目列出全部的組合
										ContentAll := betCount.ShiShiAnyPositionPick3Bet(thisBOContent.(map[string]interface{})["1"].([]interface{})) //列全部的部份，同三星组六复式
										tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)

										//使用迴圈對獎，依下注的組合逐筆驗證
										for i := 0; i < len(ContentAll); i++ {
											if len(betCount.Array_unique_str(ContentAll[i].([]string))) == 3 &&
												betCount.In_array(ContentAll[i].([]string)[0], tempResult) &&
												betCount.In_array(ContentAll[i].([]string)[1], tempResult) &&
												betCount.In_array(ContentAll[i].([]string)[2], tempResult) {
												Inspect["Ratio"].([]float64)[0] = Inspect["Ratio"].([]float64)[0] + 1
												Inspect["Message"] = 1
											}

										}

									}
								default:
									Inspect["State"] = false
									Inspect["Info"] = "Lottery Play Group Mode Error"

								}
							}
							//玩法群組 (双面/串关)********
						case 8:
							{
								//玩法 (因為有對應 TS_LotteryPlay.LP_Code，所以需要跟資料庫一起變動...ry)
								switch LotteryPlay {
								//玩法 (前二大小单双) 	//玩法 (后二大小单双)
								case 1, 2:
									{
										//依下注的項目列出全部的組合
										ContentAll := betCount.ShiShiTwoStarSideBet(thisBOContent.(map[string]interface{})["1"].([]string), thisBOContent.(map[string]interface{})["2"].([]string))
										tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)
										//使用迴圈對獎，依下注的組合逐筆驗證
										for i := 0; i < len(ContentAll); i++ {
											if betCount.In_array(ContentAll[i].([]string)[0], tempResult) &&
												betCount.In_array(ContentAll[i].([]string)[1], tempResult) {

												Inspect["Ratio"].([]float64)[0] = Inspect["Ratio"].([]float64)[0] + 1
												Inspect["Message"] = 1

											}

										}
									}
								//玩法 (前三大小单双) 	//玩法 (后三大小单双)
								case 3, 4:
									{
										ContentAll := betCount.ShiShiThreeStarSideBet(thisBOContent.(map[string]interface{})["1"].([]string), thisBOContent.(map[string]interface{})["2"].([]string), thisBOContent.(map[string]interface{})["3"].([]string))
										tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)

										//使用迴圈對獎，依下注的組合逐筆驗證
										for i := 0; i < len(ContentAll); i++ {

											if betCount.In_array(ContentAll[i].([]string)[0], tempResult) &&
												betCount.In_array(ContentAll[i].([]string)[1], tempResult) &&
												betCount.In_array(ContentAll[i].([]string)[2], tempResult) {
												Inspect["Ratio"].([]float64)[0] = Inspect["Ratio"].([]float64)[0] + 1
												Inspect["Message"] = 1

											}

										}
									}
								default:
									Inspect["State"] = false
									Inspect["Info"] = "Lottery Play Group Mode Error"

								}
							}

						//玩法群組 (任选二)
						case 9:
							{
								//玩法 (因為有對應 TS_LotteryPlay.LP_Code，所以需要跟資料庫一起變動...ry)
								switch LotteryPlay {
								//玩法 (直选复式)
								case 1:
									{
										//因為不特定的下注項目有可能不存在，所以先預設為空陣列以免被警告或提示訊息噴好噴滿…
										if _, exists := thisBOContent.(map[string]interface{})["1"]; !exists {
											thisBOContent.(map[string]interface{})["1"] = []string{}
										}
										if _, exists := thisBOContent.(map[string]interface{})["2"]; !exists {
											thisBOContent.(map[string]interface{})["2"] = []string{}
										}
										if _, exists := thisBOContent.(map[string]interface{})["3"]; !exists {
											thisBOContent.(map[string]interface{})["3"] = []string{}
										}
										if _, exists := thisBOContent.(map[string]interface{})["4"]; !exists {
											thisBOContent.(map[string]interface{})["4"] = []string{}
										}
										if _, exists := thisBOContent.(map[string]interface{})["5"]; !exists {
											thisBOContent.(map[string]interface{})["5"] = []string{}
										}
										ContentAll := betCount.ShiShiAnyPick2Bet(thisBOContent.(map[string]interface{})["1"].([]interface{}),
											thisBOContent.(map[string]interface{})["2"].([]interface{}),
											thisBOContent.(map[string]interface{})["3"].([]interface{}),
											thisBOContent.(map[string]interface{})["4"].([]interface{}),
											thisBOContent.(map[string]interface{})["5"].([]interface{}))

										//使用迴圈對獎，依下注的組合逐筆驗證
										for i := 0; i < len(ContentAll); i++ {
											num := ContentAll[i].([]interface{})[0].([]string)[0]
											num1 := ContentAll[i].([]interface{})[1].([]string)[0]

											tmpIndex := ContentAll[i].([]interface{})[0].([]string)[1]
											tmpIndex1 := ContentAll[i].([]interface{})[1].([]string)[1]
											index, err := strconv.Atoi(tmpIndex)
											if err != nil {
												panic(err)
											}
											index1, err := strconv.Atoi(tmpIndex1)
											if err != nil {
												panic(err)
											}

											if num == thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)[index] &&
												num1 == thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)[index1] {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											}
										}
									}
									//玩法 (直选单式)
								case 5:
									{
										ContentAll := betCount.ShiShiAnyPick2ComboSimpleBet(thisBOContent.(map[string]interface{})["text"].(string))
										OptionsAll := betCount.AllOptionsCompose(thisBOContent.(map[string]interface{})["option"].([]interface{}), 2)

										tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)

										for i := 0; i < len(OptionsAll); i++ {

											item, err := strconv.Atoi(OptionsAll[i].([]string)[0])
											if err != nil {
												panic(err)
											}
											item1, err := strconv.Atoi(OptionsAll[i].([]string)[1])
											if err != nil {
												panic(err)
											}

											ResultChecker := []string{tempResult[item], tempResult[item1]}

											for j := 0; j < len(ContentAll); j++ {
												if ContentAll[j].([]string)[0] == ResultChecker[0] &&
													ContentAll[j].([]string)[1] == ResultChecker[1] {
													Inspect["Ratio"].([]float64)[0]++
													Inspect["Message"] = 1
												}
											}
										}

									}
									//玩法 (直选和值)
								case 2:
									{
										ContentAll := betCount.ShiShiAnyPick2SumBet(thisBOContent.(map[string]interface{})["1"].([]interface{}))
										OptionsAll := betCount.AllOptionsCompose(thisBOContent.(map[string]interface{})["option"].([]interface{}), 2)
										tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)

										for i := 0; i < len(OptionsAll); i++ {
											item, err := strconv.Atoi(OptionsAll[i].([]string)[0])
											if err != nil {
												panic(err)
											}
											item1, err := strconv.Atoi(OptionsAll[i].([]string)[1])
											if err != nil {
												panic(err)
											}
											ResultChecker := []string{tempResult[item], tempResult[item1]}
											for j := 0; j < len(ContentAll); j++ {
												if betCount.ArraySum(ContentAll[j].([]string)) == betCount.ArraySum(ResultChecker) {
													if ContentAll[j].([]string)[0] == ResultChecker[0] &&
														ContentAll[j].([]string)[1] == ResultChecker[1] {
														Inspect["Ratio"].([]float64)[0]++
														Inspect["Message"] = 1
													}

												}

											}
										}
									}
									//玩法 (組选复式) 	//玩法 (組选单式)
								case 3, 6:
									{

										var ContentAll []interface{}
										//依下注的項目列出全部的組合
										if LotteryPlay == 3 { //列全部的部份，同『 二星 (后二|前二) -> 组选复式 』
											ContentAll = betCount.ShiShiAnyPick2ComboBet(thisBOContent.(map[string]interface{})["1"].([]interface{}))
										} else if LotteryPlay == 6 { //列全部的部份，同『 二星 (后二|前二) -> 組选单式 』
											ContentAll = betCount.ShiShiAnyPick2ComboSimpleBet(thisBOContent.(map[string]interface{})["text"].(string))
										}
										OptionsAll := betCount.AllOptionsCompose(thisBOContent.(map[string]interface{})["option"].([]interface{}), 2)
										tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)
										//使用迴圈逐筆列出需要驗證的位數組合
										for i := 0; i < len(OptionsAll); i++ {
											item, err := strconv.Atoi(OptionsAll[i].([]string)[0])
											if err != nil {
												panic(err)
											}
											item1, err := strconv.Atoi(OptionsAll[i].([]string)[1])
											if err != nil {
												panic(err)
											}
											ResultChecker := []string{tempResult[item], tempResult[item1]}
											//使用迴圈對獎，依下注的組合逐筆驗證
											for j := 0; j < len(ContentAll); j++ {
												if betCount.In_array(ContentAll[j].([]string)[0], ResultChecker) &&
													betCount.In_array(ContentAll[j].([]string)[1], ResultChecker) {
													Inspect["Ratio"].([]float64)[0]++
													Inspect["Message"] = 1
												}
											}
										}

									}
									//玩法 (組选和值)
								case 4:
									{
										ContentAll := betCount.ShiShiAnyPick2ComboSumBet(thisBOContent.(map[string]interface{})["1"].([]interface{})) //列全部的部份，同『 二星 (后二|前二) -> 组选和值 』
										OptionsAll := betCount.AllOptionsCompose(thisBOContent.(map[string]interface{})["option"].([]interface{}), 2)
										tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)
										//使用迴圈逐筆列出需要驗證的位數組合
										for i := 0; i < len(OptionsAll); i++ {
											item, err := strconv.Atoi(OptionsAll[i].([]string)[0])
											if err != nil {
												panic(err)
											}
											item1, err := strconv.Atoi(OptionsAll[i].([]string)[1])
											if err != nil {
												panic(err)
											}
											ResultChecker := []string{tempResult[item], tempResult[item1]}
											//使用迴圈對獎，依下注的組合逐筆驗證
											for j := 0; j < len(ContentAll); j++ {
												if betCount.ArraySum(ContentAll[j].([]string)) == betCount.ArraySum(ResultChecker) {
													if betCount.In_array(ContentAll[j].([]string)[0], ResultChecker) &&
														betCount.In_array(ContentAll[j].([]string)[1], ResultChecker) {
														Inspect["Ratio"].([]float64)[0]++
														Inspect["Message"] = 1
													}
												}
											}

										}
									}
								default:
									Inspect["State"] = false
									Inspect["Info"] = "Lottery Play Group Mode Error"

								}
							}
							//玩法群組 (任选三)
						case 10:
							{
								//玩法 (因為有對應 TS_LotteryPlay.LP_Code，所以需要跟資料庫一起變動...ry)
								switch LotteryPlay {
								//玩法 (直选复式)
								case 1:
									{
										//因為不特定的下注項目有可能不存在，所以先預設為空陣列以免被警告或提示訊息噴好噴滿…
										if _, exists := thisBOContent.(map[string]interface{})["1"]; !exists {
											thisBOContent.(map[string]interface{})["1"] = []string{}
										}
										if _, exists := thisBOContent.(map[string]interface{})["2"]; !exists {
											thisBOContent.(map[string]interface{})["2"] = []string{}
										}
										if _, exists := thisBOContent.(map[string]interface{})["3"]; !exists {
											thisBOContent.(map[string]interface{})["3"] = []string{}
										}
										if _, exists := thisBOContent.(map[string]interface{})["4"]; !exists {
											thisBOContent.(map[string]interface{})["4"] = []string{}
										}
										if _, exists := thisBOContent.(map[string]interface{})["5"]; !exists {
											thisBOContent.(map[string]interface{})["5"] = []string{}
										}

										ContentAll := betCount.ShiShiAnyPick3Bet(thisBOContent.(map[string]interface{})["1"].([]interface{}), thisBOContent.(map[string]interface{})["2"].([]interface{}), thisBOContent.(map[string]interface{})["3"].([]interface{}), thisBOContent.(map[string]interface{})["4"].([]interface{}), thisBOContent.(map[string]interface{})["5"].([]interface{}))
										//使用迴圈對獎，依下注的組合逐筆驗證

										for i := 0; i < len(ContentAll); i++ {

											num1 := ContentAll[i].([]interface{})[0].([]string)[0]
											num2 := ContentAll[i].([]interface{})[1].([]string)[0]
											num3 := ContentAll[i].([]interface{})[2].([]string)[0]
											tmpIndex1 := ContentAll[i].([]interface{})[0].([]string)[1]
											tmpIndex2 := ContentAll[i].([]interface{})[1].([]string)[1]
											tmpIndex3 := ContentAll[i].([]interface{})[2].([]string)[1]
											index1, err := strconv.Atoi(tmpIndex1)
											if err != nil {
												panic(err)
											}
											index2, err := strconv.Atoi(tmpIndex2)
											if err != nil {
												panic(err)
											}
											index3, err := strconv.Atoi(tmpIndex3)
											if err != nil {
												panic(err)
											}
											if num1 == thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)[index1] &&
												num2 == thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)[index2] &&
												num3 == thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)[index3] {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											}
										}
									}
									//玩法 (直选单式)
								case 6:
									{
										ContentAll := betCount.ShiShiAnyPick3SimpleBet(thisBOContent.(map[string]interface{})["text"].(string)) ///列全部的部份，同『 三星 (后三、前三、中三) -> 直选单式 』
										OptionsAll := betCount.AllOptionsCompose(thisBOContent.(map[string]interface{})["option"].([]interface{}), 3)
										tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)
										for i := 0; i < len(OptionsAll); i++ {
											item, err := strconv.Atoi(OptionsAll[i].([]string)[0])
											if err != nil {
												panic(err)
											}
											item1, err := strconv.Atoi(OptionsAll[i].([]string)[1])
											if err != nil {
												panic(err)
											}
											item2, err := strconv.Atoi(OptionsAll[i].([]string)[2])
											if err != nil {
												panic(err)
											}

											ResultChecker := []string{tempResult[item], tempResult[item1], tempResult[item2]}
											for j := 0; j < len(ContentAll); j++ {
												if ContentAll[j].([]string)[0] == ResultChecker[0] &&
													ContentAll[j].([]string)[1] == ResultChecker[1] &&
													ContentAll[j].([]string)[2] == ResultChecker[2] {
													Inspect["Ratio"].([]float64)[0]++
													Inspect["Message"] = 1
												}
											}

										}
									}
									//玩法 (直选和值)
								case 2:
									{
										ContentAll := betCount.ShiShiAnyPick3SumBet(thisBOContent.(map[string]interface{})["1"].([]interface{})) //列全部的部份，同『 三星 (后三、前三、中三) -> 直选和值 』
										OptionsAll := betCount.AllOptionsCompose(thisBOContent.(map[string]interface{})["option"].([]interface{}), 3)
										tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)

										for i := 0; i < len(OptionsAll); i++ {
											item, err := strconv.Atoi(OptionsAll[i].([]string)[0])
											if err != nil {
												panic(err)
											}
											item1, err := strconv.Atoi(OptionsAll[i].([]string)[1])
											if err != nil {
												panic(err)
											}
											item2, err := strconv.Atoi(OptionsAll[i].([]string)[2])
											if err != nil {
												panic(err)
											}
											ResultChecker := []string{tempResult[item], tempResult[item1], tempResult[item2]}

											for j := 0; j < len(ContentAll); j++ {
												if betCount.ArraySum(ContentAll[j].([]string)) == betCount.ArraySum(ResultChecker) {
													if ContentAll[j].([]string)[0] == ResultChecker[0] &&
														ContentAll[j].([]string)[1] == ResultChecker[1] &&
														ContentAll[j].([]string)[2] == ResultChecker[2] {
														Inspect["Ratio"].([]float64)[0]++
														Inspect["Message"] = 1
													}
												}
											}
										}
									}
									//玩法 (组三复式) //玩法 (组三单式)
								case 3, 7:
									{
										var ContentAll []interface{}
										//依下注的項目列出全部的組合
										if LotteryPlay == 3 { //列全部的部份，同『 三星 (后三、前三、中三) -> 组三复式 』
											ContentAll = betCount.ShiShiAnyPick3Combo3ComplexBet(thisBOContent.(map[string]interface{})["1"].([]interface{}))
										} else if LotteryPlay == 7 { //列全部的部份，同『 二星 (后二|前二) -> 組选单式 』
											ContentAll = betCount.ShiShiAnyPick3Combo3SimpleBet(thisBOContent.(map[string]interface{})["text"].(string))
										}
										OptionsAll := betCount.AllOptionsCompose(thisBOContent.(map[string]interface{})["option"].([]interface{}), 3)
										tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)

										for i := 0; i < len(OptionsAll); i++ {
											item, err := strconv.Atoi(OptionsAll[i].([]string)[0])
											if err != nil {
												panic(err)
											}
											item1, err := strconv.Atoi(OptionsAll[i].([]string)[1])
											if err != nil {
												panic(err)
											}
											item2, err := strconv.Atoi(OptionsAll[i].([]string)[2])
											if err != nil {
												panic(err)
											}
											ResultChecker := []string{tempResult[item], tempResult[item1], tempResult[item2]}
											if len(betCount.Array_unique_str(ResultChecker)) == 2 {
												for j := 0; j < len(ContentAll); j++ {
													if len(betCount.Array_unique_str(ContentAll[j].([]string))) == 2 &&
														betCount.In_array(ContentAll[j].([]string)[0], ResultChecker) &&
														betCount.In_array(ContentAll[j].([]string)[1], ResultChecker) &&
														betCount.In_array(ContentAll[j].([]string)[2], ResultChecker) {
														Inspect["Ratio"].([]float64)[0]++
														Inspect["Message"] = 1
													}
												}
											}

										}
									}
									//玩法 (组六复式) 	//玩法 (组六单式)
								case 4, 8:
									{
										var ContentAll []interface{}
										//依下注的項目列出全部的組合
										if LotteryPlay == 4 { //列全部的部份，同『 三星 (后三、前三、中三) -> 组六复式 』
											ContentAll = betCount.ShiShiAnyPick3Combo6ComplexBet(thisBOContent.(map[string]interface{})["1"].([]interface{}))
										} else if LotteryPlay == 8 { //列全部的部份，同『三星 (后三、前三、中三) -> 组六单式 』
											ContentAll = betCount.ShiShiAnyPick3Combo6SimpleBet(thisBOContent.(map[string]interface{})["text"].(string))
										}
										// fmt.Println("Run  ContentAll ", ContentAll)
										OptionsAll := betCount.AllOptionsCompose(thisBOContent.(map[string]interface{})["option"].([]interface{}), 3)
										tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)

										for i := 0; i < len(OptionsAll); i++ {
											item, err := strconv.Atoi(OptionsAll[i].([]string)[0])
											if err != nil {
												panic(err)
											}
											item1, err := strconv.Atoi(OptionsAll[i].([]string)[1])
											if err != nil {
												panic(err)
											}
											item2, err := strconv.Atoi(OptionsAll[i].([]string)[2])
											if err != nil {
												panic(err)
											}
											ResultChecker := []string{tempResult[item], tempResult[item1], tempResult[item2]}
											// fmt.Println("Run  ResultChecker ", ResultChecker)
											if len(betCount.Array_unique_str(ResultChecker)) == 3 {
												for j := 0; j < len(ContentAll); j++ {
													if len(betCount.Array_unique_str(ContentAll[j].([]string))) == 3 &&
														betCount.In_array(ContentAll[j].([]string)[0], ResultChecker) &&
														betCount.In_array(ContentAll[j].([]string)[1], ResultChecker) &&
														betCount.In_array(ContentAll[j].([]string)[2], ResultChecker) {
														Inspect["Ratio"].([]float64)[0]++
														Inspect["Message"] = 1
													}
												}
											}

										}
									}
									//玩法 (組选和值)
								case 5:
									{
										ContentAll := betCount.ShiShiAnyPick3ComboSumBet(thisBOContent.(map[string]interface{})["1"].([]interface{})) //列全部的部份，同『 三星 (后三、前三、中三) -> 组选和值 』
										OptionsAll := betCount.AllOptionsCompose(thisBOContent.(map[string]interface{})["option"].([]interface{}), 3)
										tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)

										for i := 0; i < len(OptionsAll); i++ {
											item, err := strconv.Atoi(OptionsAll[i].([]string)[0])
											if err != nil {
												panic(err)
											}
											item1, err := strconv.Atoi(OptionsAll[i].([]string)[1])
											if err != nil {
												panic(err)
											}
											item2, err := strconv.Atoi(OptionsAll[i].([]string)[2])
											if err != nil {
												panic(err)
											}
											ResultChecker := []string{tempResult[item], tempResult[item1], tempResult[item2]}
											for j := 0; j < len(ContentAll); j++ {
												if betCount.ArraySum(ContentAll[j].([]string)) == betCount.ArraySum(ResultChecker) {
													if len(betCount.Array_unique_str(ContentAll[j].([]string))) == 2 &&
														betCount.In_array(ContentAll[j].([]string)[0], ResultChecker) &&
														betCount.In_array(ContentAll[j].([]string)[1], ResultChecker) &&
														betCount.In_array(ContentAll[j].([]string)[2], ResultChecker) {
														Inspect["Ratio"].([]float64)[0]++
														Inspect["Message"] = 1
													} else if len(betCount.Array_unique_str(ContentAll[j].([]string))) == 3 &&
														betCount.In_array(ContentAll[j].([]string)[0], ResultChecker) &&
														betCount.In_array(ContentAll[j].([]string)[1], ResultChecker) &&
														betCount.In_array(ContentAll[j].([]string)[2], ResultChecker) {
														Inspect["Ratio"].([]float64)[1]++
														Inspect["Message"] = 1

													}
												}

											}
										}

									}
									//玩法 (混合组选)
								case 9:
									{
										ContentAll := betCount.ShiShiAnyPick3ComboMixBet(thisBOContent.(map[string]interface{})["text"].(string)) //列全部的部份，同『 三星 (后三、前三、中三) -> 混合组选 』
										OptionsAll := betCount.AllOptionsCompose(thisBOContent.(map[string]interface{})["option"].([]interface{}), 3)
										tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)

										//使用迴圈逐筆列出需要驗證的位數組合
										for i := 0; i < len(OptionsAll); i++ {
											item, err := strconv.Atoi(OptionsAll[i].([]string)[0])
											if err != nil {
												panic(err)
											}
											item1, err := strconv.Atoi(OptionsAll[i].([]string)[1])
											if err != nil {
												panic(err)
											}
											item2, err := strconv.Atoi(OptionsAll[i].([]string)[2])
											if err != nil {
												panic(err)
											}
											ResultChecker := []string{tempResult[item], tempResult[item1], tempResult[item2]}

											//使用迴圈對獎，依下注的組合逐筆驗證
											for j := 0; j < len(ContentAll); j++ {
												//使用「組三单式」规则
												if len(betCount.Array_unique_str(ContentAll[j].([]string))) == 2 {

													if betCount.In_array(ContentAll[j].([]string)[0], ResultChecker) &&
														betCount.In_array(ContentAll[j].([]string)[1], ResultChecker) &&
														betCount.In_array(ContentAll[j].([]string)[2], ResultChecker) {

														ResultTemp := betCount.Dup_count(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))
														ResultSorted := betCount.SortMapByValue(ResultTemp)

														ContentTemp := betCount.Dup_count(ContentAll[j].([]string))
														ContentSorted := betCount.SortMapByValue(ContentTemp)

														for _, v := range ResultSorted {
															if v.Key == ContentSorted[0].Key && v.Value == 2 {
																Inspect["Ratio"].([]float64)[0]++
																Inspect["Message"] = 1
															}
														}

														// ResultTemp := betCount.Dup_count(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))
														// ResultSorted := betCount.SortMapByValue(ResultTemp)
														// for _, v := range ResultSorted {
														// 	if v.Key == ContentAll[j].([]string)[0] && v.Value == 2 {
														// 		Inspect["Ratio"].([]float64)[0] = Inspect["Ratio"].([]float64)[0] + 1
														// 		Inspect["Message"] = 1
														// 	}
														// }

													}
													//使用「組六单式」规则
												} else if len(betCount.Array_unique_str(ContentAll[j].([]string))) == 3 {
													if betCount.In_array(ContentAll[j].([]string)[0], ResultChecker) &&
														betCount.In_array(ContentAll[j].([]string)[1], ResultChecker) &&
														betCount.In_array(ContentAll[j].([]string)[2], ResultChecker) {
														Inspect["Ratio"].([]float64)[1]++
														Inspect["Message"] = 1

													}
												}

											}
										}
									}
								default:
									Inspect["State"] = false
									Inspect["Info"] = "Lottery Play Group Mode Error"
								}
							}
						//玩法群組 (任选四)
						case 11:
							{
								//玩法 (因為有對應 TS_LotteryPlay.LP_Code，所以需要跟資料庫一起變動...ry)
								switch LotteryPlay {
								//玩法 (直选复式)
								case 1:
									{
										//因為不特定的下注項目有可能不存在，所以先預設為空陣列以免被警告或提示訊息噴好噴滿…
										if _, exists := thisBOContent.(map[string]interface{})["1"]; !exists {
											thisBOContent.(map[string]interface{})["1"] = []string{}
										}
										if _, exists := thisBOContent.(map[string]interface{})["2"]; !exists {
											thisBOContent.(map[string]interface{})["2"] = []string{}
										}
										if _, exists := thisBOContent.(map[string]interface{})["3"]; !exists {
											thisBOContent.(map[string]interface{})["3"] = []string{}
										}
										if _, exists := thisBOContent.(map[string]interface{})["4"]; !exists {
											thisBOContent.(map[string]interface{})["4"] = []string{}
										}
										if _, exists := thisBOContent.(map[string]interface{})["5"]; !exists {
											thisBOContent.(map[string]interface{})["5"] = []string{}
										}

										ContentAll := betCount.ShiShiAnyPick4Bet(thisBOContent.(map[string]interface{})["1"].([]interface{}), thisBOContent.(map[string]interface{})["2"].([]interface{}), thisBOContent.(map[string]interface{})["3"].([]interface{}), thisBOContent.(map[string]interface{})["4"].([]interface{}), thisBOContent.(map[string]interface{})["5"].([]interface{}))
										//使用迴圈對獎，依下注的組合逐筆驗證
										//fmt.Println("Run  ContentAll ", ContentAll)
										for i := 0; i < len(ContentAll); i++ {

											num1 := ContentAll[i].([]interface{})[0].([]string)[0]
											num2 := ContentAll[i].([]interface{})[1].([]string)[0]
											num3 := ContentAll[i].([]interface{})[2].([]string)[0]
											num4 := ContentAll[i].([]interface{})[3].([]string)[0]
											tmpIndex1 := ContentAll[i].([]interface{})[0].([]string)[1]
											tmpIndex2 := ContentAll[i].([]interface{})[1].([]string)[1]
											tmpIndex3 := ContentAll[i].([]interface{})[2].([]string)[1]
											tmpIndex4 := ContentAll[i].([]interface{})[3].([]string)[1]
											index1, err := strconv.Atoi(tmpIndex1)
											if err != nil {
												panic(err)
											}
											index2, err := strconv.Atoi(tmpIndex2)
											if err != nil {
												panic(err)
											}
											index3, err := strconv.Atoi(tmpIndex3)
											if err != nil {
												panic(err)
											}
											index4, err := strconv.Atoi(tmpIndex4)
											if err != nil {
												panic(err)
											}
											if num1 == thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)[index1] &&
												num2 == thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)[index2] &&
												num3 == thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)[index3] &&
												num4 == thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)[index4] {
												Inspect["Ratio"].([]float64)[0]++
												Inspect["Message"] = 1
											}
										}
									}
									//玩法 (直选单式)
								case 6:
									{

										ContentAll := betCount.ShiShiAnyPick4SimpleBet(thisBOContent.(map[string]interface{})["text"].(string)) //列全部的部份，同『 三星 (后三、前三、中三) -> 直选单式 』
										OptionsAll := betCount.AllOptionsCompose(thisBOContent.(map[string]interface{})["option"].([]interface{}), 4)
										tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)
										for i := 0; i < len(OptionsAll); i++ {
											item1, err := strconv.Atoi(OptionsAll[i].([]string)[0])
											if err != nil {
												panic(err)
											}
											item2, err := strconv.Atoi(OptionsAll[i].([]string)[1])
											if err != nil {
												panic(err)
											}
											item3, err := strconv.Atoi(OptionsAll[i].([]string)[2])
											if err != nil {
												panic(err)
											}
											item4, err := strconv.Atoi(OptionsAll[i].([]string)[3])
											if err != nil {
												panic(err)
											}
											ResultChecker := []string{tempResult[item1], tempResult[item2], tempResult[item3], tempResult[item4]}

											for j := 0; j < len(ContentAll); j++ {
												if betCount.ArraySum(ContentAll[j].([]string)) == betCount.ArraySum(ResultChecker) {
													if ContentAll[j].([]string)[0] == ResultChecker[0] &&
														ContentAll[j].([]string)[1] == ResultChecker[1] &&
														ContentAll[j].([]string)[2] == ResultChecker[2] &&
														ContentAll[j].([]string)[3] == ResultChecker[3] {
														Inspect["Ratio"].([]float64)[0]++
														Inspect["Message"] = 1
													}
												}
											}
										}

									}
									//玩法 (组选24)
								case 2:
									{
										ContentAll := betCount.ShiShiAnyPick4Combo24Bet(thisBOContent.(map[string]interface{})["1"].([]interface{})) //列全部的部份，同『 三星 (后三、前三、中三) -> 直选单式 』
										OptionsAll := betCount.AllOptionsCompose(thisBOContent.(map[string]interface{})["option"].([]interface{}), 4)
										tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)

										for i := 0; i < len(OptionsAll); i++ {
											item, err := strconv.Atoi(OptionsAll[i].([]string)[0])
											if err != nil {
												panic(err)
											}
											item1, err := strconv.Atoi(OptionsAll[i].([]string)[1])
											if err != nil {
												panic(err)
											}
											item2, err := strconv.Atoi(OptionsAll[i].([]string)[2])
											if err != nil {
												panic(err)
											}
											item3, err := strconv.Atoi(OptionsAll[i].([]string)[3])
											if err != nil {
												panic(err)
											}
											ResultChecker := []string{tempResult[item], tempResult[item1], tempResult[item2], tempResult[item3]}
											if len(betCount.Array_unique_str(ResultChecker)) == 4 {
												for j := 0; j < len(ContentAll); j++ {
													if betCount.In_array(ContentAll[j].([]string)[0], ResultChecker) &&
														betCount.In_array(ContentAll[j].([]string)[1], ResultChecker) &&
														betCount.In_array(ContentAll[j].([]string)[2], ResultChecker) &&
														betCount.In_array(ContentAll[j].([]string)[3], ResultChecker) {
														Inspect["Ratio"].([]float64)[0]++
														Inspect["Message"] = 1
													}
												}

											}
										}
									}
								//玩法 (组选12)
								case 3:
									{
										ContentAll := betCount.ShiShiAnyPick4Combo12Bet(thisBOContent.(map[string]interface{})["1"].([]interface{}), thisBOContent.(map[string]interface{})["2"].([]interface{})) //列全部的部份，同『 四星 -> 组选12 』
										OptionsAll := betCount.AllOptionsCompose(thisBOContent.(map[string]interface{})["option"].([]interface{}), 4)
										tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)

										for i := 0; i < len(OptionsAll); i++ {
											item, err := strconv.Atoi(OptionsAll[i].([]string)[0])
											if err != nil {
												panic(err)
											}
											item1, err := strconv.Atoi(OptionsAll[i].([]string)[1])
											if err != nil {
												panic(err)
											}
											item2, err := strconv.Atoi(OptionsAll[i].([]string)[2])
											if err != nil {
												panic(err)
											}
											item3, err := strconv.Atoi(OptionsAll[i].([]string)[3])
											if err != nil {
												panic(err)
											}
											ResultChecker := []string{tempResult[item], tempResult[item1], tempResult[item2], tempResult[item3]}
											if len(betCount.Array_unique_str(ResultChecker)) == 3 {
												for j := 0; j < len(ContentAll); j++ {
													if betCount.In_array(ContentAll[j].([]string)[0], ResultChecker) &&
														betCount.In_array(ContentAll[j].([]string)[1], ResultChecker) &&
														betCount.In_array(ContentAll[j].([]string)[2], ResultChecker) &&
														betCount.In_array(ContentAll[j].([]string)[3], ResultChecker) {

														ResultTemp := betCount.Dup_count(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))
														ResultSorted := betCount.SortMapByValue(ResultTemp)

														ContentTemp := betCount.Dup_count(ContentAll[j].([]string))
														ContentSorted := betCount.SortMapByValue(ContentTemp)

														for _, v := range ResultSorted {
															if v.Key == ContentSorted[0].Key && v.Value == 2 {
																Inspect["Ratio"].([]float64)[0]++
																Inspect["Message"] = 1
															}
														}

													}
												}
											}

										}
									}
									//玩法 (组选6)
								case 4:
									{
										ContentAll := betCount.ShiShiAnyPick4Combo6Bet(thisBOContent.(map[string]interface{})["1"].([]interface{})) //列全部的部份，同『 四星 -> 组选6 』
										OptionsAll := betCount.AllOptionsCompose(thisBOContent.(map[string]interface{})["option"].([]interface{}), 4)
										tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)

										for i := 0; i < len(OptionsAll); i++ {
											item, err := strconv.Atoi(OptionsAll[i].([]string)[0])
											if err != nil {
												panic(err)
											}
											item1, err := strconv.Atoi(OptionsAll[i].([]string)[1])
											if err != nil {
												panic(err)
											}
											item2, err := strconv.Atoi(OptionsAll[i].([]string)[2])
											if err != nil {
												panic(err)
											}
											item3, err := strconv.Atoi(OptionsAll[i].([]string)[3])
											if err != nil {
												panic(err)
											}
											ResultChecker := []string{tempResult[item], tempResult[item1], tempResult[item2], tempResult[item3]}
											if len(betCount.Array_unique_str(ResultChecker)) == 2 {
												for j := 0; j < len(ContentAll); j++ {
													if betCount.In_array(ContentAll[j].([]string)[0], ResultChecker) &&
														betCount.In_array(ContentAll[j].([]string)[1], ResultChecker) &&
														betCount.In_array(ContentAll[j].([]string)[2], ResultChecker) &&
														betCount.In_array(ContentAll[j].([]string)[3], ResultChecker) {

														check1 := false
														check2 := false
														ResultTemp := betCount.Dup_count(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))
														ResultSorted := betCount.SortMapByValue(ResultTemp)

														ContentTemp := betCount.Dup_count(ContentAll[j].([]string))
														ContentSorted := betCount.SortMapByValue(ContentTemp)

														fmt.Println("Run  ContentSorted ", ContentSorted)
														for _, v := range ResultSorted {

															if v.Key == ContentSorted[0].Key && v.Value == 2 {
																check1 = true
															} else if v.Key == ContentSorted[1].Key && v.Value == 2 {
																check2 = true
															}
														}

														if check1 && check2 {
															Inspect["Ratio"].([]float64)[0]++
															Inspect["Message"] = 1

														}

														// ResultTemp := betCount.Dup_count(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))
														// ResultSorted := betCount.SortMapByValue(ResultTemp)

														// check1 := false
														// check2 := false
														// for _, v := range ResultSorted {
														// 	if v.Key == ContentAll[i].([]string)[0] && v.Value == 2 {
														// 		check1 = true
														// 	} else if v.Key == ContentAll[i].([]string)[2] && v.Value == 2 {
														// 		check2 = true
														// 	}
														// }
														// if check1 && check2 {
														// 	Inspect["Ratio"].([]float64)[0] = Inspect["Ratio"].([]float64)[0] + 1
														// 	Inspect["Message"] = 1

														// }
													}
												}
											}
										}

									}
									//玩法 (组选4)
								case 5:
									{
										ContentAll := betCount.ShiShiAnyPick4Combo4Bet(thisBOContent.(map[string]interface{})["1"].([]interface{}), thisBOContent.(map[string]interface{})["2"].([]interface{})) //列全部的部份，同『 四星 -> 组选4 』
										OptionsAll := betCount.AllOptionsCompose(thisBOContent.(map[string]interface{})["option"].([]interface{}), 4)
										tempResult := thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string)

										for i := 0; i < len(OptionsAll); i++ {
											item, err := strconv.Atoi(OptionsAll[i].([]string)[0])
											if err != nil {
												panic(err)
											}
											item1, err := strconv.Atoi(OptionsAll[i].([]string)[1])
											if err != nil {
												panic(err)
											}
											item2, err := strconv.Atoi(OptionsAll[i].([]string)[2])
											if err != nil {
												panic(err)
											}
											item3, err := strconv.Atoi(OptionsAll[i].([]string)[3])
											if err != nil {
												panic(err)
											}
											ResultChecker := []string{tempResult[item], tempResult[item1], tempResult[item2], tempResult[item3]}
											// fmt.Println(" ResultChecker ", ResultChecker)
											if len(betCount.Array_unique_str(ResultChecker)) == 3 {
												for j := 0; j < len(ContentAll); j++ {
													if betCount.In_array(ContentAll[j].([]string)[0], ResultChecker) &&
														betCount.In_array(ContentAll[j].([]string)[1], ResultChecker) &&
														betCount.In_array(ContentAll[j].([]string)[2], ResultChecker) &&
														betCount.In_array(ContentAll[j].([]string)[3], ResultChecker) {

														ResultTemp := betCount.Dup_count(thisFullResult[strconv.Itoa(LotteryMode)].(map[string]interface{})[strconv.Itoa(LotteryPlayGroup)].(map[string]interface{})[strconv.Itoa(LotteryPlay)].([]string))
														ResultSorted := betCount.SortMapByValue(ResultTemp)

														ContentTemp := betCount.Dup_count(ContentAll[j].([]string))
														ContentSorted := betCount.SortMapByValue(ContentTemp)

														for _, v := range ResultSorted {
															if v.Key == ContentSorted[0].Key && v.Value == 3 {
																Inspect["Ratio"].([]float64)[0]++
																Inspect["Message"] = 1

															}
														}
													}
												}

											}

										}

									}
								default:
									Inspect["State"] = false
									Inspect["Info"] = "Lottery Play Group Mode Error"

								}
							}
						default:
							Inspect["State"] = false
							Inspect["Info"] = "Lottery Play Group Mode Error"

						}
					}
				//彩種錯誤
				default:
					Inspect["State"] = false
				}
			}
		}
	default:
		Inspect["State"] = false
	}

	return Inspect
}
