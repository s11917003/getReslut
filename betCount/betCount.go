package betCount

import (
	"reflect"
	"sort"
	"strconv"
	"strings"
)

/***********************************************************************
 *	官方玩法注數計算
 *
 *	輸入：  TypeGroup = 彩種群組 (INT)
 *			PlayGroup  = 玩法群組 (INT)
 *			Play = 玩法 (INT)
 *			BC_Content = 實際下注內容 (map[string]interface{})
 *	輸出：  Count = 已填入實際數值的開獎結果 (INT)
 ***********************************************************************/
func BetCountOfficial(
	TypeGroup int,
	PlayGroup int,
	Play int,
	LotteryContent interface{}) int {

	//return 1
	Count := 0
	content := LotteryContent.(map[string]interface{})["1"].(map[string]interface{})
	switch TypeGroup {
	case 2:
		switch PlayGroup {
		//玩法群組 (定位胆)
		case 1:
			Count = ShiShiStraightBetCount(content["1"].([]interface{}), content["2"].([]interface{}), content["3"].([]interface{}), content["4"].([]interface{}), content["5"].([]interface{}))
		//玩法群組 (五星)
		case 2:
			switch Play {
			case 1:
				Count = ShiShiFiveStarBetCount(content["1"].([]interface{}), content["2"].([]interface{}), content["3"].([]interface{}), content["4"].([]interface{}), content["5"].([]interface{}))
			case 8:
				Count = ShiShiFiveStarSimpleBetCount(content["text"].(string))
			case 2:
				Count = ShiShiFiveStarCombo120BetCount(content["1"].([]interface{}))
			case 3:
				Count = ShiShiFiveStarCombo60BetCount(content["1"].([]interface{}), content["2"].([]interface{}))
			case 4:
				Count = ShiShiFiveStarCombo30BetCount(content["1"].([]interface{}), content["2"].([]interface{}))
			case 5:
				Count = ShiShiFiveStarCombo20BetCount(content["1"].([]interface{}), content["2"].([]interface{}))
			case 6:
				Count = ShiShiFiveStarCombo10BetCount(content["1"].([]interface{}), content["2"].([]interface{}))
			case 7:
				Count = ShiShiFiveStarCombo5BetCount(content["1"].([]interface{}), content["2"].([]interface{}))

			}
		//玩法群組 (四星)
		case 3:
			switch Play {
			case 1:
				Count = ShiShiFourStarBetCount(content["1"].([]interface{}), content["2"].([]interface{}), content["3"].([]interface{}), content["4"].([]interface{}))
			case 6:
				Count = ShiShiFourStarSimpleBetCount(content["text"].(string))
			case 2:
				Count = ShiShiFourStarCombo24BetCount(content["1"].([]interface{}))
			case 3:
				Count = ShiShiFourStarCombo12BetCount(content["1"].([]interface{}), content["2"].([]interface{}))
			case 4:
				Count = ShiShiFourStarCombo6BetCount(content["1"].([]interface{}))
			case 5:
				Count = ShiShiFourStarCombo4BetCount(content["1"].([]interface{}), content["2"].([]interface{}))

			}
		//玩法群組 (后三) //玩法群組 (前三) //玩法群組 (中三)
		case 4, 5, 12:
			switch Play {
			case 1:
				Count = ShiShiThreeStarBetCount(content["1"].([]interface{}), content["2"].([]interface{}), content["3"].([]interface{}))
			case 11:
				Count = ShiShiThreeStarSimpleBetCount(content["text"].(string))
			case 2:
				Count = ShiShiThreeStarSumBetCount(content["1"].([]interface{}))
			case 3:
				Count = ShiShiThreeStarCutResultBetCount(content["1"].([]interface{}))
			case 4:
				Count = ShiShiThreeStarComboBetCount(content["1"].([]interface{}), content["2"].([]interface{}), content["3"].([]interface{}))
			case 5:
				Count = ShiShiThreeStarCombo3ComplexBetCount(content["1"].([]interface{}))
			case 12:
				Count = ShiShiThreeStarCombo3SimpleBetCount(content["text"].(string))
			case 6:
				Count = ShiShiThreeStarCombo6ComplexBetCount(content["1"].([]interface{}))
			case 13:
				Count = ShiShiThreeStarCombo6SimpleBetCount(content["text"].(string))
			case 7:
				Count = ShiShiThreeStarComboSumBetCount(content["1"].([]interface{}))
			case 8:
				Count = ShiShiThreeStarComboBaoDnoBetCount(content["1"].([]interface{}))
			case 9, 10:
				Count = ShiShiThreeStarComboTailBetCount(content["1"].([]interface{}))
			case 14:
				Count = ShiShiThreeStarComboMixBetCount(content["text"].(string)) //目前固定1注

			}

		//玩法群組 (前二) //玩法群組 (后二)
		case 6, 13:
			switch Play {
			case 1:
				Count = ShiShiTwoStarBetCount(content["1"].([]interface{}), content["2"].([]interface{}))
			case 7:
				Count = ShiShiTwoStarSimpleBetCount(content["text"].(string))
			case 2:
				Count = ShiShiTwoStarSumBetCount(content["1"].([]interface{}))
			case 3:
				Count = ShiShiTwoStarCutResultBetCount(content["1"].([]interface{}))
			case 4:
				Count = ShiShiTwoStarComboComplexBetCount(content["1"].([]interface{}))
			case 8:
				Count = ShiShiTwoStarComboSimpleBetCount(content["text"].(string))
			case 5:
				Count = ShiShiTwoStarComboSumBetCount(content["1"].([]interface{}))
			case 6:
				Count = ShiShiTwoStarComboBaoDnoBetCount(content["1"].([]interface{}))
			}

		//玩法群組 (不定位)
		case 7:
			switch Play {
			case 1, 3, 12, 5, 7, 9:
				Count = ShiShiAnyPositionPick1BetCount(content["1"].([]interface{}))
			case 2, 4, 13, 6, 8, 10:
				Count = ShiShiAnyPositionPick2BetCount(content["1"].([]interface{}))
			case 11:
				Count = ShiShiAnyPositionPick3BetCount(content["1"].([]interface{}))
			}
		//玩法群組 (双面/串关)
		case 8:
			switch Play {
			case 1, 2:
				Count = ShiShiTwoStarSideBetCount(content["1"].([]string), content["2"].([]string))
			case 3, 4:
				Count = ShiShiThreeStarSideBetCount(content["1"].([]string), content["2"].([]string), content["3"].([]string))
			case 5:
				break // 20180502 停用，沒有賠率且規則不明確

			}
		//玩法群組 (任选二)
		case 9:
			switch Play {
			case 1:
				Count = ShiShiAnyPick2BetCount(content["1"].([]interface{}), content["2"].([]interface{}), content["3"].([]interface{}), content["4"].([]interface{}), content["5"].([]interface{}))
			case 5:
				Count = ShiShiAnyPick2SimpleBetCount(content["text"].(string), content["option"].([]interface{}))
			case 2:
				Count = ShiShiAnyPick2SumBetCount(content["1"].([]interface{}), content["option"].([]interface{}))
			case 3:
				Count = ShiShiAnyPick2ComboBetCount(content["1"].([]interface{}), content["option"].([]interface{}))
			case 6:
				Count = ShiShiAnyPick2ComboSimpleBetCount(content["text"].(string), content["option"].([]interface{}))
			case 4:
				Count = ShiShiAnyPick2ComboSumBetCount(content["1"].([]interface{}), content["option"].([]interface{}))
			}

		//玩法群組 (任选三)
		case 10:
			switch Play {
			case 1:
				Count = ShiShiAnyPick3BetCount(content["1"].([]interface{}), content["2"].([]interface{}), content["3"].([]interface{}), content["4"].([]interface{}), content["5"].([]interface{}))
			case 6:
				Count = ShiShiAnyPick3SimpleBetCount(content["text"].(string), content["option"].([]interface{}))
			case 2:
				Count = ShiShiAnyPick3SumBetCount(content["1"].([]interface{}), content["option"].([]interface{}))
			case 3:
				Count = ShiShiAnyPick3Combo3ComplexBetCount(content["1"].([]interface{}), content["option"].([]interface{}))
			case 4:
				Count = ShiShiAnyPick3Combo6ComplexBetCount(content["1"].([]interface{}), content["option"].([]interface{}))
			case 5:
				Count = ShiShiAnyPick3ComboSumBetCount(content["1"].([]interface{}), content["option"].([]interface{}))
			case 7:
				Count = ShiShiAnyPick3Combo3SimpleBetCount(content["text"].(string), content["option"].([]interface{}))
			case 8:
				Count = ShiShiAnyPick3Combo6SimpleBetCount(content["text"].(string), content["option"].([]interface{}))
			case 9:
				Count = ShiShiAnyPick3ComboMixBetCount(content["text"].(string), content["option"].([]interface{}))

			}

			//玩法群組 (任选四)
		case 11:
			switch Play {
			case 1:
				Count = ShiShiAnyPick4BetCount(content["1"].([]interface{}), content["2"].([]interface{}), content["3"].([]interface{}), content["4"].([]interface{}), content["5"].([]interface{}))
			case 6:
				Count = ShiShiAnyPick4SimpleBetCount(content["text"].(string), content["option"].([]interface{}))
			case 2:
				Count = ShiShiAnyPick4Combo24BetCount(content["1"].([]interface{}), content["option"].([]interface{}))
			case 3:
				Count = ShiShiAnyPick4Combo12BetCount(content["1"].([]interface{}), content["2"].([]interface{}), content["option"].([]interface{}))
			case 4:
				Count = ShiShiAnyPick4Combo6BetCount(content["1"].([]interface{}), content["option"].([]interface{}))
			case 5:
				Count = ShiShiAnyPick4Combo4BetCount(content["1"].([]interface{}), content["2"].([]interface{}), content["option"].([]interface{}))
			}

		}

	}

	return Count
}

/***********************************************************************
* 	Content1 資料陣列 (Array)
* 	Content2 資料陣列 (Array)
* 	Content3 資料陣列 (Array)
* 	Content4 資料陣列 (Array)
* 	Content5 資料陣列 (Array)
***********************************************************************/
func ShiShiStraightBetCount(Content1 []interface{}, Content2 []interface{}, Content3 []interface{}, Content4 []interface{}, Content5 []interface{}) int {
	Count := 0
	Count = len(ShiShiStraightBet(Content1, Content2, Content3, Content4, Content5))
	return Count
}

/***********************************************************************
* 	Content1 資料陣列 (Array)
* 	Content2 資料陣列 (Array)
* 	Content3 資料陣列 (Array)
* 	Content4 資料陣列 (Array)
* 	Content5 資料陣列 (Array)
 ***********************************************************************/
func ShiShiStraightBet(Content1 []interface{}, Content2 []interface{}, Content3 []interface{}, Content4 []interface{}, Content5 []interface{}) []interface{} {

	var ContentAll []interface{}
	//Count = count($this->ShiShiStraightBet($SBC_Content_1, $SBC_Content_2, $SBC_Content_3, $SBC_Content_4, $SBC_Content_5));

	if len(Content1) > 0 {
		for i := 0; i < len(Content1); i++ {
			ContentAll = append(ContentAll, []string{Content1[i].(string), "0"})
		}

	}
	if len(Content2) > 0 {
		for i := 0; i < len(Content2); i++ {
			ContentAll = append(ContentAll, []string{Content2[i].(string), "1"})
		}

	}
	if len(Content3) > 0 {
		for i := 0; i < len(Content3); i++ {
			ContentAll = append(ContentAll, []string{Content3[i].(string), "2"})
		}

	}
	if len(Content4) > 0 {
		for i := 0; i < len(Content4); i++ {
			ContentAll = append(ContentAll, []string{Content4[i].(string), "3"})
		}

	}
	if len(Content5) > 0 {
		for i := 0; i < len(Content5); i++ {
			ContentAll = append(ContentAll, []string{Content5[i].(string), "4"})
		}

	}
	return ContentAll
}

/***********************************************************************
 *	五星 -> 直选复式 [START] */

/***********************************************************************
 * 	$VSBC_Content_1 資料陣列 (Array)
 * 	$VSBC_Content_2 資料陣列 (Array)
 * 	$VSBC_Content_3 資料陣列 (Array)
 * 	$VSBC_Content_4 資料陣列 (Array)
 * 	$VSBC_Content_5 資料陣列 (Array)
 ***********************************************************************/
func ShiShiFiveStarBetCount(Content1 []interface{}, Content2 []interface{}, Content3 []interface{}, Content4 []interface{}, Content5 []interface{}) int {
	Count := 0
	if len(Content1) > 0 && len(Content2) > 0 && len(Content3) > 0 && len(Content4) > 0 && len(Content5) > 0 {
		Count = len(ShiShiFiveStarBet(Content1, Content2, Content3, Content4, Content5))
	}
	return Count
}

/***********************************************************************
 * 	$SB_Content_1 資料陣列 (Array)
 * 	$SB_Content_2 資料陣列 (Array)
 * 	$SB_Content_3 資料陣列 (Array)
 * 	$SB_Content_4 資料陣列 (Array)
 * 	$SB_Content_5 資料陣列 (Array)
 ***********************************************************************/
func ShiShiFiveStarBet(Content1 []interface{}, Content2 []interface{}, Content3 []interface{}, Content4 []interface{}, Content5 []interface{}) []interface{} {
	var ContentAll []interface{}
	for i1 := 0; i1 < len(Content1); i1++ {
		for i2 := 0; i2 < len(Content2); i2++ {
			for i3 := 0; i3 < len(Content3); i3++ {
				for i4 := 0; i4 < len(Content4); i4++ {
					for i5 := 0; i5 < len(Content5); i5++ {
						ContentAll = append(ContentAll, []string{Content1[i1].(string), Content2[i2].(string), Content3[i3].(string), Content4[i4].(string), Content5[i5].(string)})

					}
				}
			}
		}
	}
	return ContentAll
}

/*	五星 -> 直选复式 [END]
***********************************************************************/

/***********************************************************************
 *	五星 -> 直选单式 [START] */

/***********************************************************************
 * 	Content 資料陣列 (String)
 ***********************************************************************/
func ShiShiFiveStarSimpleBetCount(Content string) int {
	Count := 0
	Count = len(ShiShiFiveStarSimpleBet(Content))
	return Count
}

/***********************************************************************
 * 	Content 資料陣列 (String)
 ***********************************************************************/
func ShiShiFiveStarSimpleBet(Content string) []interface{} {
	var ContentAll []interface{}

	Content1 := strings.Split(Content, ",")
	for i1 := 0; i1 < len(Content1); i1++ {
		Content2 := strings.Split(Content1[i1], "")

		//建立要儲存的組合為陣列
		ContentAll = append(ContentAll, []string{Content2[0], Content2[1], Content2[2], Content2[3], Content2[4]})

	}
	return ContentAll
}

/*	五星 -> 直选单式 [END]
***********************************************************************/

/***********************************************************************
 *	五星 -> 组选120 [START] */

/***********************************************************************
 * 	Content 資料陣列 (Array)
 ***********************************************************************/
func ShiShiFiveStarCombo120BetCount(Content []interface{}) int {
	Count := 0
	if len(Content) > 0 {
		Count = len(ShiShiFiveStarCombo120Bet(Content))
	}
	return Count
}

/***********************************************************************
 * 	Content 資料陣列 (Array)
 ***********************************************************************/
func ShiShiFiveStarCombo120Bet(Content []interface{}) []interface{} { //列全部
	var ContentAll []interface{}
	for i1 := 0; i1 < len(Content); i1++ {
		for i2 := i1 + 1; i2 < len(Content); i2++ {
			for i3 := i2 + 1; i3 < len(Content); i3++ {
				for i4 := i3 + 1; i4 < len(Content); i4++ {
					for i5 := i4 + 1; i5 < len(Content); i5++ {
						ContentAll = append(ContentAll, []string{Content[i1].(string), Content[i2].(string), Content[i3].(string), Content[i4].(string), Content[i5].(string)})

					}
				}
			}
		}
	}
	return ContentAll
}

/*	五星 -> 组选120 [END]
***********************************************************************/

/***********************************************************************
 *	五星 -> 组选60 [START] */

/***********************************************************************
 * 	$VSC60BC_Content_1 資料陣列 (Array)
 * 	$VSC60BC_Content_2 資料陣列 (Array)
 ***********************************************************************/
func ShiShiFiveStarCombo60BetCount(Content1 []interface{}, Content2 []interface{}) int {
	Count := 0
	if len(Content1) > 0 && len(Content2) > 0 {
		Count = len(ShiShiFiveStarCombo60Bet(Content1, Content2))
	}
	return Count
}

/***********************************************************************
 * 	$VSC60B_Content_1 資料陣列 (Array)
 * 	$VSC60B_Content_2 資料陣列 (Array)
 ***********************************************************************/
func ShiShiFiveStarCombo60Bet(Content1 []interface{}, Content2 []interface{}) []interface{} { //列全部
	var ContentAll []interface{}
	for i1 := 0; i1 < len(Content1); i1++ {
		for i2 := 0; i2 < len(Content2); i2++ {
			for i3 := i2 + 1; i3 < len(Content2); i3++ {
				for i4 := i3 + 1; i4 < len(Content2); i4++ {
					if len(Array_unique_str([]string{Content1[i1].(string), Content2[i2].(string), Content2[i3].(string), Content2[i4].(string)})) == 4 { //避免重複的被塞進去
						ContentAll = append(ContentAll, []string{Content1[i1].(string), Content1[i1].(string), Content2[i2].(string), Content2[i3].(string), Content2[i4].(string)})
					}
				}
			}
		}
	}
	return ContentAll
}

/*	五星 -> 组选60 [END]
***********************************************************************/

/***********************************************************************
 *	五星 -> 组选30 [START] */

/***********************************************************************
 * 	$VSC30BC_Content_1 資料陣列 (Array)
 * 	$VSC30BC_Content_2 資料陣列 (Array)
 ***********************************************************************/
func ShiShiFiveStarCombo30BetCount(Content1 []interface{}, Content2 []interface{}) int {
	Count := 0
	if len(Content1) > 0 && len(Content2) > 0 {
		if len(Content1) >= 2 {
			Count = len(ShiShiFiveStarCombo30Bet(Content1, Content2))
		}
	}
	return Count
}

/***********************************************************************
 * 	$VSC30B_Content_1 資料陣列 (Array)
 * 	$VSC30B_Content_2 資料陣列 (Array)
 ***********************************************************************/
func ShiShiFiveStarCombo30Bet(Content1 []interface{}, Content2 []interface{}) []interface{} { //列全部

	var ContentAll []interface{}
	for i1 := 0; i1 < len(Content1); i1++ {
		for i2 := i1 + 1; i2 < len(Content1); i2++ {
			for i3 := 0; i3 < len(Content2); i3++ {
				if len(Array_unique_str([]string{Content1[i1].(string), Content1[i2].(string), Content2[i3].(string)})) == 3 { //避免重複的被塞進去
					ContentAll = append(ContentAll, []string{Content1[i1].(string), Content1[i1].(string), Content1[i2].(string), Content1[i2].(string), Content2[i3].(string)})
				}

			}
		}
	}
	return ContentAll
}

/*	五星 -> 组选30 [END]
***********************************************************************/

/***********************************************************************
 *	五星 -> 组选20 [START] */

/***********************************************************************
 * 	$VSC20BC_Content_1 資料陣列 (Array)
 * 	$VSC20BC_Content_2 資料陣列 (Array)
 ***********************************************************************/
func ShiShiFiveStarCombo20BetCount(Content1 []interface{}, Content2 []interface{}) int {

	Count := 0
	if len(Content1) > 0 && len(Content2) > 0 {
		if len(Content2) >= 2 {
			Count = len(ShiShiFiveStarCombo20Bet(Content1, Content2))
		}
	}
	return Count
}

/***********************************************************************
 * 	$VSC20B_Content_1 資料陣列 (Array)
 * 	$VSC20B_Content_2 資料陣列 (Array)
 ***********************************************************************/
func ShiShiFiveStarCombo20Bet(Content1 []interface{}, Content2 []interface{}) []interface{} { //列全部

	var ContentAll []interface{}

	for i1 := 0; i1 < len(Content1); i1++ {
		for i2 := 0; i2 < len(Content2); i2++ {
			for i3 := i2 + 1; i3 < len(Content2); i3++ {
				if len(Array_unique_str([]string{Content1[i1].(string), Content2[i2].(string), Content2[i3].(string)})) == 3 { //避免重複的被塞進去
					ContentAll = append(ContentAll, []string{Content1[i1].(string), Content1[i1].(string), Content1[i1].(string), Content2[i2].(string), Content2[i3].(string)})

				}
			}
		}
	}
	return ContentAll
}

/*	五星 -> 组选20 [END]
***********************************************************************/

/***********************************************************************
 *	五星 -> 组选10 [START] */

/***********************************************************************
 * 	$VSC10BC_Content_1 資料陣列 (Array)
 * 	$VSC10BC_Content_2 資料陣列 (Array)
 ***********************************************************************/
func ShiShiFiveStarCombo10BetCount(Content1 []interface{}, Content2 []interface{}) int {
	Count := 0
	if len(Content1) > 0 && len(Content2) > 0 {
		Count = len(ShiShiFiveStarCombo10Bet(Content1, Content2))
	}
	return Count
}

/***********************************************************************
 * 	$VSC10B_Content_1 資料陣列 (Array)
 * 	$VSC10B_Content_2 資料陣列 (Array)
 ***********************************************************************/
func ShiShiFiveStarCombo10Bet(Content1 []interface{}, Content2 []interface{}) []interface{} { //列全部
	var ContentAll []interface{}
	for i1 := 0; i1 < len(Content1); i1++ {
		for i2 := 0; i2 < len(Content2); i2++ {
			if Content1[i1] != Content2[i2] { //避免重複的被塞進去
				ContentAll = append(ContentAll, []string{Content1[i1].(string), Content1[i1].(string), Content1[i1].(string), Content2[i2].(string), Content2[i2].(string)})
			}
		}
	}
	return ContentAll
}

/*	五星 -> 组选10 [END]
***********************************************************************/

/***********************************************************************
 *	五星 -> 组选5 [START] */

/***********************************************************************
 * 	$VSC5BC_Content_1 資料陣列 (Array)
 * 	$VSC5BC_Content_2 資料陣列 (Array)
 ***********************************************************************/
func ShiShiFiveStarCombo5BetCount(Content1 []interface{}, Content2 []interface{}) int {
	Count := 0
	if len(Content1) > 0 && len(Content2) > 0 {
		Count = len(ShiShiFiveStarCombo5Bet(Content1, Content2))
	}
	return Count
}

/***********************************************************************
 * 	$VSC5B_Content_1 資料陣列 (Array)
 * 	$VSC5B_Content_2 資料陣列 (Array)
 ***********************************************************************/
func ShiShiFiveStarCombo5Bet(Content1 []interface{}, Content2 []interface{}) []interface{} { //列全部
	var ContentAll []interface{}
	for i1 := 0; i1 < len(Content1); i1++ {
		for i2 := 0; i2 < len(Content2); i2++ {
			if Content1[i1] != Content2[i2] { //避免重複的被塞進去
				ContentAll = append(ContentAll, []string{Content1[i1].(string), Content1[i1].(string), Content1[i1].(string), Content1[i2].(string), Content2[i2].(string)})
			}
		}
	}
	return ContentAll
}

/*	五星 -> 组选5 [END]
 ***********************************************************************/

/*	五星 [END]
 ***********************************************************************/

/***********************************************************************
 *	四星 [START] */

/***********************************************************************
 *	四星 -> 直选复式 [START] */

/***********************************************************************
 * 	$FSBC_Content_1 資料陣列 (Array)
 * 	$FSBC_Content_2 資料陣列 (Array)
 * 	$FSBC_Content_3 資料陣列 (Array)
 * 	$FSBC_Content_4 資料陣列 (Array)
 ***********************************************************************/
func ShiShiFourStarBetCount(Content1 []interface{}, Content2 []interface{}, Content3 []interface{}, Content4 []interface{}) int {
	Count := 0
	if len(Content1) > 0 && len(Content2) > 0 && len(Content3) > 0 && len(Content4) > 0 {
		Count = len(ShiShiFourStarBet(Content1, Content2, Content3, Content4))
	}
	return Count
}

/***********************************************************************
 * 	$FSB_Content_1 資料陣列 (Array)
 * 	$FSB_Content_2 資料陣列 (Array)
 * 	$FSB_Content_3 資料陣列 (Array)
 * 	$FSB_Content_4 資料陣列 (Array)
 *
*	共用：『 任选四 -> 直选单式 』
***********************************************************************/
func ShiShiFourStarBet(Content1 []interface{}, Content2 []interface{}, Content3 []interface{}, Content4 []interface{}) []interface{} { //列全部
	var ContentAll []interface{}

	for i1 := 0; i1 < len(Content1); i1++ {
		for i2 := 0; i2 < len(Content2); i2++ {
			for i3 := 0; i3 < len(Content3); i3++ {
				for i4 := 0; i4 < len(Content4); i4++ {
					ContentAll = append(ContentAll, []string{Content1[i1].(string), Content2[i2].(string), Content3[i3].(string), Content4[i4].(string)})

				}
			}
		}
	}

	return ContentAll
}

/*	四星 -> 直选复式 [END]
 ***********************************************************************/

/***********************************************************************
 *	四星 -> 直选单式 [START] */

/***********************************************************************
 * 	$FSSBC_Content 資料陣列 (String)
 ***********************************************************************/
func ShiShiFourStarSimpleBetCount(Content string) int {
	Count := 0
	if len(Content) > 0 {
		Count = len(ShiShiFourStarSimpleBet(Content))
	}
	return Count
}

/***********************************************************************
 * 	$FSSB_Content 資料陣列 (String)
 ***********************************************************************/
func ShiShiFourStarSimpleBet(Content string) []interface{} { //列全部
	var ContentAll []interface{}

	Content1 := strings.Split(Content, ",")
	for i1 := 0; i1 < len(Content1); i1++ {
		Content2 := strings.Split(Content1[i1], "")

		//建立要儲存的組合為陣列
		ContentAll = append(ContentAll, []string{Content2[0], Content2[1], Content2[2], Content2[3]})

	}
	return ContentAll
}

/*	四星 -> 直选单式 [END]
 ***********************************************************************/

/***********************************************************************
 *	四星 -> 组选24 [START] */

/***********************************************************************
 * 	$FSC24BC_Content 資料陣列 (Array)
 ***********************************************************************/
func ShiShiFourStarCombo24BetCount(Content []interface{}) int {

	Count := 0
	if len(Content) > 0 {
		Count = len(ShiShiFourStarCombo24Bet(Content))
	}
	return Count
}

/***********************************************************************
 * 	$FSC24B_Content 資料陣列 (Array)
 *
 *	共用：	『 任选四 -> 组选24 』
***********************************************************************/
func ShiShiFourStarCombo24Bet(Content []interface{}) []interface{} { //列全部

	var ContentAll []interface{}
	for i1 := 0; i1 < len(Content); i1++ {
		for i2 := i1 + 1; i2 < len(Content); i2++ {
			for i3 := i2 + 1; i3 < len(Content); i3++ {
				for i4 := i3 + 1; i4 < len(Content); i4++ {
					//建立要儲存的組合為陣列
					ContentAll = append(ContentAll, []string{Content[i1].(string), Content[i2].(string), Content[i3].(string), Content[i4].(string)})
				}
			}
		}
	}

	return ContentAll
}

/*	四星 -> 组选24 [END]
 ***********************************************************************/

/***********************************************************************
 *	四星 -> 组选12 [START] */

/***********************************************************************
 * 	$FSC12BC_Content_1 資料陣列 (Array)
 * 	$FSC12BC_Content_2 資料陣列 (Array)
 ***********************************************************************/
func ShiShiFourStarCombo12BetCount(Content1 []interface{}, Content2 []interface{}) int {

	Count := 0
	if len(Content1) > 0 && len(Content2) > 0 {

		Count = len(ShiShiFourStarCombo12Bet(Content1, Content2))
	}
	return Count
}

/***********************************************************************
 * 	$FSC12B_Content_1 資料陣列 (Array)
 * 	$FSC12B_Content_2 資料陣列 (Array)
 *
 *	共用：	『 任选四 -> 组选12 』
***********************************************************************/
func ShiShiFourStarCombo12Bet(Content1 []interface{}, Content2 []interface{}) []interface{} { //列全部
	var ContentAll []interface{}

	for i1 := 0; i1 < len(Content1); i1++ {
		for i2 := 0; i2 < len(Content2); i2++ {
			for i3 := i2 + 1; i3 < len(Content2); i3++ {
				if len(Array_unique_str([]string{Content1[i1].(string), Content2[i2].(string), Content2[i3].(string)})) == 3 { //避免重複的被塞進去
					ContentAll = append(ContentAll, []string{Content1[i1].(string), Content1[i1].(string), Content2[i2].(string), Content2[i3].(string)})
				}
			}
		}
	}

	return ContentAll
}

/*	四星 -> 组选12 [END]
 ***********************************************************************/

/***********************************************************************
 *	四星 -> 组选6 [START] */

/***********************************************************************
 * 	$FSC6BC_Content 資料陣列 (Array)
 ***********************************************************************/
func ShiShiFourStarCombo6BetCount(Content []interface{}) int {
	Count := 0
	if len(Content) > 0 {
		Count = len(ShiShiFourStarCombo6Bet(Content))
	}
	return Count
}

/***********************************************************************
 * 	$FSC6B_Content 資料陣列 (Array)
 *
 *	共用：	『 任选四 -> 组选6 』
***********************************************************************/
func ShiShiFourStarCombo6Bet(Content []interface{}) []interface{} { //列全部
	var ContentAll []interface{}
	for i1 := 0; i1 < len(Content); i1++ {
		for i2 := i1 + 1; i2 < len(Content); i2++ {
			ContentAll = append(ContentAll, []string{Content[i1].(string), Content[i1].(string), Content[i2].(string), Content[i2].(string)})

		}
	}
	return ContentAll
}

/*	四星 -> 组选6 [END]
 ***********************************************************************/

/***********************************************************************
 *	四星 -> 组选4 [START] */

/***********************************************************************
 * 	$FSC4BC_Content_1 資料陣列 (Array)
 * 	$FSC4BC_Content_2 資料陣列 (Array)
 ***********************************************************************/
func ShiShiFourStarCombo4BetCount(Content1 []interface{}, Content2 []interface{}) int {
	Count := 0
	if len(Content1) > 0 && len(Content2) > 0 {
		Count = len(ShiShiFourStarCombo4Bet(Content1, Content2))
	}
	return Count
}

/***********************************************************************
 * 	$FSC4B_Content_1 資料陣列 (Array)
 * 	$FSC4B_Content_2 資料陣列 (Array)
 *
 *	共用：『 任选四 -> 组选4 』
***********************************************************************/
func ShiShiFourStarCombo4Bet(Content1 []interface{}, Content2 []interface{}) []interface{} { //列全部

	var ContentAll []interface{}

	for i1 := 0; i1 < len(Content1); i1++ {
		for i2 := 0; i2 < len(Content2); i2++ {
			if Content1[i1] != Content2[i2] { //避免重複的被塞進去
				ContentAll = append(ContentAll, []string{Content1[i1].(string), Content1[i1].(string), Content1[i1].(string), Content2[i2].(string)})
			}

		}
	}
	return ContentAll
}

/*	四星 -> 组选4 [END]
 ***********************************************************************/

/*	四星 [END]
 ***********************************************************************/

/***********************************************************************
 *	三星 (后三、前三、中三) [START] */

/***********************************************************************
 *	三星 (后三、前三、中三) -> 直选复式 [START] */

/***********************************************************************
 * 	$TSBC_Content_1 資料陣列 (Array)
 * 	$TSBC_Content_2 資料陣列 (Array)
 * 	$TSBC_Content_3 資料陣列 (Array)
 ***********************************************************************/
func ShiShiThreeStarBetCount(Content1 []interface{}, Content2 []interface{}, Content3 []interface{}) int {
	Count := 0
	if len(Content1) > 0 && len(Content2) > 0 && len(Content3) > 0 {
		Count = len(ShiShiThreeStarBet(Content1, Content2, Content3))
	}
	return Count
}

/***********************************************************************
 * 	$TSB_Content_1 資料陣列 (Array)
 * 	$TSB_Content_2 資料陣列 (Array)
 * 	$TSB_Content_3 資料陣列 (Array)
 *
 *	共用：	『 任选三 -> 组六单式 』
*			『 任选三 -> 组六单式 』
***********************************************************************/
func ShiShiThreeStarBet(Content1 []interface{}, Content2 []interface{}, Content3 []interface{}) []interface{} {
	var ContentAll []interface{}
	for i1 := 0; i1 < len(Content1); i1++ {
		for i2 := 0; i2 < len(Content2); i2++ {
			for i3 := 0; i3 < len(Content3); i3++ {
				ContentAll = append(ContentAll, []string{Content1[i1].(string), Content2[i2].(string), Content3[i3].(string)})
			}
		}
	}
	return ContentAll
}

/*	三星 (后三、前三、中三) -> 直选复式 [END]
 ***********************************************************************/

/***********************************************************************
 *	三星 (后三、前三、中三) -> 直选单式 [START] */

/***********************************************************************
 * 	$HSSBC_Content 資料陣列 (String)
 ***********************************************************************/
func ShiShiThreeStarSimpleBetCount(Content string) int {
	Count := 0
	if len(Content) > 0 {
		Count = len(ShiShiThreeStarSimpleBet(Content))
	}
	return Count
}

/***********************************************************************
 * 	$HSSB_Content 資料陣列 (String)
 *
 *	共用：	『 三星 (后三、前三、中三) -> 直选单式 』
***********************************************************************/
func ShiShiThreeStarSimpleBet(Content string) []interface{} {

	var ContentAll []interface{}
	Content1 := strings.Split(Content, ",")
	for i1 := 0; i1 < len(Content1); i1++ {
		Content2 := strings.Split(Content1[i1], "")
		//建立要儲存的組合為陣列
		ContentAll = append(ContentAll, []string{Content2[0], Content2[1], Content2[2]})
	}
	return ContentAll
}

/*	三星 (后三、前三、中三) -> 直选单式 [END]
 ***********************************************************************/

/***********************************************************************
 *	三星 (后三、前三、中三) -> 直选和值 [START] */

/***********************************************************************
 *	$TSBC_Content 資料陣列 (Array)
***********************************************************************/
func ShiShiThreeStarSumBetCount(Content []interface{}) int {
	Count := 0
	if len(Content) > 0 { //如果裡面有值

		Count = len(ShiShiThreeStarSumBet(Content))
	}
	return Count
}

/***********************************************************************
 * 	$TSB_Content 資料陣列 (Array)
 *
 *	共用：	『 任选三 -> 直选和值 』
***********************************************************************/
func ShiShiThreeStarSumBet(Content []interface{}) []interface{} { //列全部
	var ContentAll []interface{}

	for i1 := 0; i1 < len(Content); i1++ {
		for i2 := 0; i2 < 10; i2++ {
			for i3 := 0; i3 < 10; i3++ {
				for i4 := 0; i4 < 10; i4++ {

					item, err := strconv.Atoi(Content[i1].(string))
					if err != nil {
						panic(err)
					}

					if (i2 + i3 + i4) == item {
						ContentAll = append(ContentAll, []string{strconv.Itoa(i2), strconv.Itoa(i3), strconv.Itoa(i4)})
					}
				}
			}
		}
	}
	return ContentAll
}

/*	三星 (后三、前三、中三) -> 直选和值 [END]
***********************************************************************/

/***********************************************************************
 *	三星 (后三、前三、中三) -> 直选跨度 [START] */

/***********************************************************************
 *	$TCRBC_Content 資料陣列 (Array)
***********************************************************************/
func ShiShiThreeStarCutResultBetCount(Content []interface{}) int {
	Count := 0
	if len(Content) > 0 {
		Count = len(ShiShiThreeStarCutResultBet(Content))
	}
	return Count
}

/***********************************************************************
 * 	$TCRB_Content 資料陣列 (Array)
 ***********************************************************************/
func ShiShiThreeStarCutResultBet(Content []interface{}) []interface{} { //列全部
	var ContentAll []interface{}
	for i1 := 0; i1 < len(Content); i1++ {
		for i2 := 0; i2 < 10; i2++ {
			for i3 := 0; i3 < 10; i3++ {
				for i4 := 0; i4 < 10; i4++ {
					item, err := strconv.Atoi(Content[i1].(string))
					if err != nil {
						panic(err)
					}

					if (MaxIntSlice([]int{i2, i3, i4}) - MinIntSlice([]int{i2, i3, i4})) == item {
						ContentAll = append(ContentAll, []string{strconv.Itoa(i2), strconv.Itoa(i3), strconv.Itoa(i4)})
					}
				}
			}
		}
	}
	return ContentAll
}

/*	三星 (后三、前三、中三) -> 直选跨度 [END]
***********************************************************************/

/***********************************************************************
 *	三星 (后三、前三、中三) -> 组选组合 [START] */

/***********************************************************************
 * 	$TSCBC_Content_1 資料陣列 (Array)
 * 	$TSCBC_Content_2 資料陣列 (Array)
 * 	$TSCBC_Content_3 資料陣列 (Array)
 ***********************************************************************/
func ShiShiThreeStarComboBetCount(Content1 []interface{}, Content2 []interface{}, Content3 []interface{}) int {
	Count := 0
	if len(Content1) > 0 && len(Content2) > 0 && len(Content3) > 0 {
		Count = len(ShiShiThreeStarComboBet(Content1, Content2, Content3)) * 3
	}
	return Count
}

/***********************************************************************
 * 	$TSCB_Content_1 資料陣列 (Array)
 * 	$TSCB_Content_2 資料陣列 (Array)
 * 	$TSCB_Content_3 資料陣列 (Array)
 ***********************************************************************/
func ShiShiThreeStarComboBet(Content1 []interface{}, Content2 []interface{}, Content3 []interface{}) []interface{} { //列全部
	var ContentAll []interface{}

	for i1 := 0; i1 < len(Content1); i1++ {
		for i2 := 0; i2 < len(Content2); i2++ {
			for i3 := 0; i3 < len(Content3); i3++ {
				ContentAll = append(ContentAll, []string{Content1[i1].(string), Content2[i2].(string), Content3[i3].(string)})
			}
		}
	}
	return ContentAll
}

/*	三星 (后三、前三、中三) -> 组选组合 [END]
 ***********************************************************************/

/***********************************************************************
 *	三星 (后三、前三、中三) -> 组三复式 [START] */

/***********************************************************************
 * 	$TSC3CBC_Content 資料陣列 (Array)
 ***********************************************************************/
func ShiShiThreeStarCombo3ComplexBetCount(Content []interface{}) int {
	Count := 0
	if len(Content) >= 2 {
		Count = len(ShiShiThreeStarCombo3ComplexBet(Content))
	}
	return Count
}

/***********************************************************************
 * 	$TSC3FB_Content 資料陣列 (Array)
 *
 *	共用：	『 任选三 -> 组三复式 』
***********************************************************************/
func ShiShiThreeStarCombo3ComplexBet(Content []interface{}) []interface{} { //列全部

	var ContentAll []interface{}

	for i1 := 0; i1 < len(Content); i1++ {
		for i2 := 0; i2 < len(Content); i2++ {
			if Content[i1] != Content[i2] { //避免重複的被塞進去
				ContentAll = append(ContentAll, []string{Content[i1].(string), Content[i1].(string), Content[i2].(string)})

			}
		}
	}

	return ContentAll
}

/*	三星 (后三、前三、中三) -> 组三复式 [END]
 ***********************************************************************/

/***********************************************************************
 *	三星 (后三、前三、中三) -> 组三单式 [START] */

/***********************************************************************
 * 	$TSC3SBC_Content 資料陣列 (String)
 ***********************************************************************/
func ShiShiThreeStarCombo3SimpleBetCount(Content string) int {
	Count := 0
	if len(Content) > 0 {
		Count = len(ShiShiThreeStarCombo3SimpleBet(Content))
	}
	return Count
}

/***********************************************************************
 * 	$TSC3SB_Content 資料陣列 (String)
 *
 *	共用：	『 任选三 -> 组三单式 』
***********************************************************************/
func ShiShiThreeStarCombo3SimpleBet(Content string) []interface{} {
	var ContentAll []interface{}
	Content1 := strings.Split(Content, ",")
	for i1 := 0; i1 < len(Content1); i1++ {
		Content2 := strings.Split(Content1[i1], "")

		ContentTemp := []string{Content2[0], Content2[1], Content2[2]}

		sort.Strings(ContentTemp)
		//建立要儲存的組合為陣列

		//如果要儲存的組合符合規則、且沒有在要回傳的列表中，就加進去…
		if !In_array(ContentTemp, ContentAll) && len(Array_unique_str(ContentTemp)) == 2 {
			ContentAll = append(ContentAll, ContentTemp)
		}
	}

	return ContentAll

}

/*	三星 (后三、前三、中三) -> 组三单式 [END]
 ***********************************************************************/

/***********************************************************************
 *	三星 (后三、前三、中三) -> 组六复式 [START] */

/***********************************************************************
 * 	$TSC6CBC_Content 資料陣列 (Array)
 ***********************************************************************/
func ShiShiThreeStarCombo6ComplexBetCount(Content []interface{}) int {
	Count := 0
	if len(Content) >= 3 {
		Count = len(ShiShiThreeStarCombo6ComplexBet(Content))
	}
	return Count
}

/***********************************************************************
 * 	$TSC6FB_Content 資料陣列 (Array)
 *
 *	共用：	『 不定位 -> 五星三码 』
*			『 任选三 -> 组六复式 』
***********************************************************************/
func ShiShiThreeStarCombo6ComplexBet(Content []interface{}) []interface{} { //列全部

	var ContentAll []interface{}
	for i1 := 0; i1 < len(Content); i1++ {
		for i2 := i1 + 1; i2 < len(Content); i2++ {
			for i3 := i2 + 1; i3 < len(Content); i3++ {
				ContentAll = append(ContentAll, []string{Content[i1].(string), Content[i2].(string), Content[i3].(string)})
			}
		}
	}

	return ContentAll
}

/*	三星 (后三、前三、中三) -> 组六复式 [END]
 ***********************************************************************/

/***********************************************************************
 *	三星 (后三、前三、中三) -> 组六单式 [START] */

/***********************************************************************
 * 	$TSC6SBC_Content 資料陣列 (String)
 ***********************************************************************/
func ShiShiThreeStarCombo6SimpleBetCount(Content string) int {
	Count := 0
	if len(Content) > 0 {
		Count = len(ShiShiThreeStarCombo6SimpleBet(Content))
	}
	return Count
}

/***********************************************************************
 * 	$TSC6SB_Content 資料陣列 (String)
 *
 *	共用：	『 任选三 -> 组六单式 』
***********************************************************************/
func ShiShiThreeStarCombo6SimpleBet(Content string) []interface{} {
	var ContentAll []interface{}
	Content1 := strings.Split(Content, ",")
	for i1 := 0; i1 < len(Content1); i1++ {
		Content2 := strings.Split(Content1[i1], "")
		ContentTemp := []string{Content2[0], Content2[1], Content2[2]}
		sort.Strings(ContentTemp)
		//如果要儲存的組合符合規則、且沒有在要回傳的列表中，就加進去…
		if !In_array(ContentTemp, ContentAll) && len(Array_unique_str(ContentTemp)) == 3 {
			ContentAll = append(ContentAll, ContentTemp)
		}

	}
	return ContentAll

}

/*	三星 (后三、前三、中三) -> 组六单式 [END]
 ***********************************************************************/

/***********************************************************************
 *	三星 (后三、前三、中三) -> 组选和值 [START] */

/***********************************************************************
 * 	$TSCSBC_Content 資料陣列 (Array)
 ***********************************************************************/
func ShiShiThreeStarComboSumBetCount(Content []interface{}) int {
	Count := 0
	if len(Content) > 0 {
		Count = len(ShiShiThreeStarComboSumBet(Content))
	}
	return Count
}

/***********************************************************************
 * 	$TSCSB_Content 資料陣列 (Array)
 *
 *	共用：	『 任选三 -> 組选和值  』
***********************************************************************/
func ShiShiThreeStarComboSumBet(Content []interface{}) []interface{} { //列全部
	var ContentAll []interface{}
	for i1 := 0; i1 < len(Content); i1++ {
		for i2 := 0; i2 < 10; i2++ {
			for i3 := 0; i3 < 10; i3++ {
				for i4 := 0; i4 < 10; i4++ {

					if len(Array_unique_str([]string{strconv.Itoa(i2), strconv.Itoa(i3), strconv.Itoa(i4)})) > 1 { //避免豹子被塞進去
						item, err := strconv.Atoi(Content[i1].(string))
						if err != nil {
							panic(err)
						}

						if (i2 + i3 + i4) == item {
							Temp := []string{strconv.Itoa(i2), strconv.Itoa(i3), strconv.Itoa(i4)}
							sort.Strings(Temp)
							if !In_array2(Temp, ContentAll) {
								ContentAll = append(ContentAll, Temp)

							}
							// all, _ := json.Marshal(ContentAll)
							// tmp, _ := json.Marshal(Temp)
							// if strings.Index(string(all), string(tmp)) > -1 {

							// }
						}
					}
				}
			}
		}
	}

	return ContentAll
}

/*	三星 (后三、前三、中三) -> 组选和值 [END]
 ***********************************************************************/

/***********************************************************************
 *	三星 (后三、前三、中三) -> 组选包胆 [START] */

/***********************************************************************
 * 	$HSCBDBC_Content 資料陣列 (Array)
 ***********************************************************************/
func ShiShiThreeStarComboBaoDnoBetCount(Content []interface{}) int {
	Count := 0
	if len(Content) == 1 {
		Count = len(ShiShiThreeStarComboBaoDnoBet(Content))
	}
	return Count
}

/***********************************************************************
 * 	$HSCBDB_Content 資料陣列 (Array)
 ***********************************************************************/
func ShiShiThreeStarComboBaoDnoBet(Content []interface{}) []interface{} {
	var ContentAll []interface{}
	for i1 := 0; i1 < len(Content); i1++ {
		for i2 := 0; i2 < 10; i2++ {
			for i3 := 0; i3 < 10; i3++ {
				for i4 := 0; i4 < 10; i4++ {
					Temp := []string{strconv.Itoa(i2), strconv.Itoa(i3), strconv.Itoa(i4)}
					if len(Array_unique_str(Temp)) > 1 { //避免豹子被塞進去
						if In_array(Content[i1], Temp) {

							if !In_array2(Temp, ContentAll) {
								ContentAll = append(ContentAll, Temp)

							}
							// sort.Strings(Temp)
							// all, _ := json.Marshal(ContentAll)
							// tmp, _ := json.Marshal(Temp)

							// if strings.Index(string(all), string(tmp)) > -1 {
							// 	ContentAll = append(ContentAll, Temp)
							// }
						}
					}
				}
			}
		}
	}

	return ContentAll
}

/*	三星 (后三、前三、中三) -> 组选包胆 [END]
 ***********************************************************************/

/***********************************************************************
 *	三星 (后三、前三、中三) -> 和值尾数|特殊号 [START] */

/***********************************************************************
 * 	$TSCTBC_Content 資料陣列 (Array)
 ***********************************************************************/
func ShiShiThreeStarComboTailBetCount(Content []interface{}) int {
	Count := 0
	if len(Content) > 0 {
		Count = len(Content)
	}
	return Count
}

/*	三星 (后三、前三、中三) -> 和值尾数|特殊号 [END]
 ***********************************************************************/

/***********************************************************************
 *	三星 (后三、前三、中三) -> 混合组选 [START] */

/***********************************************************************
 * 	$TSCMBC_Content 資料陣列 (Array)
 ***********************************************************************/
func ShiShiThreeStarComboMixBetCount(Content string) int {
	Count := 0
	if len(Content) > 0 {
		Count = len(ShiShiThreeStarComboMixBet(Content))
	}
	return Count
}

/***********************************************************************
 * 	$TSCMB_Content (String)
 *
 *	共用：	『 任选三 -> 混合组选 』
***********************************************************************/
func ShiShiThreeStarComboMixBet(Content string) []interface{} {
	var ContentAll []interface{}

	Content1 := strings.Split(Content, ",")
	for i1 := 0; i1 < len(Content1); i1++ {
		Content2 := strings.Split(Content1[i1], "")
		ContentTemp := []string{Content2[0], Content2[1], Content2[2]} //建立要儲存的組合為陣列
		sort.Strings(ContentTemp)                                      //排序要儲存的組合
		//如果要儲存的組合符合規則、且沒有在要回傳的列表中，就加進去…
		if !In_array2(ContentTemp, ContentAll) {
			ContentAll = append(ContentAll, ContentTemp)
		}

	}
	return ContentAll
}

/*	三星 (后三、前三、中三) -> 混合组选 [END]
 ***********************************************************************/

/*	三星 (后三、前三、中三) [END]
 ***********************************************************************/

/***********************************************************************
 *	二星 (后二|前二) [START] */

/***********************************************************************
 *	二星 (后二|前二) -> 直选复式 [START] */

/***********************************************************************
 * 	$TSBC_Content_1 資料陣列 (Array)
 * 	$TSBC_Content_2 資料陣列 (Array)
 ***********************************************************************/
func ShiShiTwoStarBetCount(Content1 []interface{}, Content2 []interface{}) int {
	Count := 0
	if len(Content1) > 0 && len(Content2) > 0 {
		Count = len(ShiShiTwoStarBet(Content1, Content2))
	}
	return Count
}

/***********************************************************************
 * 	$TSB_Content_1 資料陣列 (Array)
 * 	$TSB_Content_2 資料陣列 (Array)
 *
 *	共用：	『 二星 (后二|前二) -> 组选单式 』
*			『 任选二 -> 直选单式 』
*			『 任选二 -> 组选单式 』
***********************************************************************/
func ShiShiTwoStarBet(Content1 []interface{}, Content2 []interface{}) []interface{} {
	var ContentAll []interface{}
	for i1 := 0; i1 < len(Content1); i1++ {
		for i2 := 0; i2 < len(Content2); i2++ {
			ContentAll = append(ContentAll, []string{Content1[i1].(string), Content2[i2].(string)})

		}
	}
	return ContentAll
}

/*	二星 (后二|前二) -> 直选复式 [END]
 ***********************************************************************/

/***********************************************************************
 *	二星 (后二|前二) -> 直选单式 [START] */

/***********************************************************************
 * 	$TSSBC_Content 資料陣列 (String)
 ***********************************************************************/
func ShiShiTwoStarSimpleBetCount(Content string) int {
	Count := 0
	if len(Content) > 0 {
		Count = len(ShiShiTwoStarSimpleBet(Content))
	}
	return Count
}

/***********************************************************************
 * 	$TSSB_Content 資料陣列 (String)
 *
 *	共用：	『 任选二 -> 直选单式 』
***********************************************************************/
func ShiShiTwoStarSimpleBet(Content string) []interface{} {
	var ContentAll []interface{}
	Content1 := strings.Split(Content, ",")
	for i1 := 0; i1 < len(Content1); i1++ {
		Content2 := strings.Split(Content1[i1], "")
		ContentTemp := []string{Content2[0], Content2[1]} //建立要儲存的組合為陣列
		ContentAll = append(ContentAll, ContentTemp)
	}
	return ContentAll
}

/*	二星 (后二|前二) -> 直选单式 [END]
 ***********************************************************************/

/***********************************************************************
 *	二星 (后二|前二) -> 直选和值 [START] */

/***********************************************************************
 *	$TSSBC_Count 資料陣列 (Array)
***********************************************************************/
func ShiShiTwoStarSumBetCount(Content []interface{}) int {
	Count := 0
	if len(Content) > 0 { //裡面必須要有值
		Count = len(ShiShiTwoStarSumBet(Content))
	}
	return Count
}

/***********************************************************************
 * 	$TSSB_Content 資料陣列 (Array)
 *
 *	共用：	『 任选二 -> 直选和值 』
***********************************************************************/
func ShiShiTwoStarSumBet(Content []interface{}) []interface{} { //列全部
	var ContentAll []interface{}
	for i1 := 0; i1 < len(Content); i1++ {
		for i2 := 0; i2 < 10; i2++ {
			for i3 := 0; i3 < 10; i3++ {
				item, err := strconv.Atoi(Content[i1].(string))
				if err != nil {
					panic(err)
				}
				if (i2 + i3) == item {
					ContentAll = append(ContentAll, []string{strconv.Itoa(i2), strconv.Itoa(i3)})

				}
			}
		}
	}

	return ContentAll
}

/*	二星 (后二|前二) -> 直选和值 [END]
***********************************************************************/

/***********************************************************************
 *	二星 (后二|前二) -> 直选跨度 [START] */

/***********************************************************************
 *	$TSCRBC_Content 資料陣列 (Array)
***********************************************************************/
func ShiShiTwoStarCutResultBetCount(Content []interface{}) int {
	Count := 0
	if len(Content) > 0 { //裡面必須要有值
		Count = len(ShiShiTwoStarCutResultBet(Content))
	}
	return Count
}

/***********************************************************************
 * 	$TSCRB_Content 資料陣列 (Array)
 ***********************************************************************/
func ShiShiTwoStarCutResultBet(Content []interface{}) []interface{} { //列全部

	var ContentAll []interface{}
	for i1 := 0; i1 < len(Content); i1++ {
		for i2 := 0; i2 < 10; i2++ {
			for i3 := 0; i3 < 10; i3++ {
				item, err := strconv.Atoi(Content[i1].(string))
				if err != nil {
					panic(err)
				}

				if MaxIntSlice([]int{i2, i3})-MinIntSlice([]int{i2, i3}) == item {
					ContentAll = append(ContentAll, []string{strconv.Itoa(i2), strconv.Itoa(i3)})
				}
			}
		}
	}

	return ContentAll
}

/*	二星 (后二|前二) -> 直选跨度 [END]
***********************************************************************/

/***********************************************************************
 *	二星 (后二|前二) -> 组选复式 [START] */

/***********************************************************************
 * 	$TSCCBC_Content 資料陣列 (Array)
 ***********************************************************************/
func ShiShiTwoStarComboComplexBetCount(Content []interface{}) int {
	Count := 0
	if len(Content) >= 2 { //裡面必須要有值
		Count = len(ShiShiTwoStarComboComplexBet(Content))
	}
	return Count
}

/***********************************************************************
 * 	$TSCFB_Content 資料陣列 (Array)
 *
 *	共用：	『 不定位 -> 二码 (前三|后三|中三|前四|后四|五星) 』
*			『 任选二 -> 組选复式 』
***********************************************************************/
func ShiShiTwoStarComboComplexBet(Content []interface{}) []interface{} { //列全部

	var ContentAll []interface{}
	for i1 := 0; i1 < len(Content); i1++ {
		for i2 := i1 + 1; i2 < len(Content); i2++ {
			ContentAll = append(ContentAll, []string{Content[i1].(string), Content[i2].(string)})
		}
	}

	return ContentAll
}

/*	二星 (后二|前二) -> 组选复式 [END]
 ***********************************************************************/

/***********************************************************************
 *	二星 (后二|前二) -> 组选单式 [START] */

/***********************************************************************
 * 	$TSCBC_Content 資料陣列 (String)
 ***********************************************************************/
func ShiShiTwoStarComboSimpleBetCount(Content string) int {
	Count := 0
	if len(Content) > 0 {
		Count = len(ShiShiTwoStarComboSimpleBet(Content))
	}
	return Count
}

/***********************************************************************
 * 	$TSCB_Content 資料陣列 (String)
 *
 *	共用：	『 任选二 -> 直选单式 』
***********************************************************************/
func ShiShiTwoStarComboSimpleBet(Content string) []interface{} {
	var ContentAll []interface{}

	Content1 := strings.Split(Content, ",")

	for i1 := 0; i1 < len(Content1); i1++ {
		Content2 := strings.Split(Content1[i1], "")
		ContentTemp := []string{Content2[0], Content2[1]} //建立要儲存的組合為陣列
		sort.Strings(ContentTemp)
		//如果要儲存的組合符合規則、且沒有在要回傳的列表中，就加進去…
		if !In_array(ContentTemp, ContentAll) {
			ContentAll = append(ContentAll, ContentTemp)
		}
	}
	return ContentAll
}

/*	二星 (后二|前二) -> 组选单式 [END]
 ***********************************************************************/

/***********************************************************************
 *	二星 (后二|前二) -> 组选和值 [START] */

/***********************************************************************
 * 	$TSCSBC_Content 資料陣列 (Array)
 *
 *	共用：	『 任选二 -> 组选和值 』
***********************************************************************/
func ShiShiTwoStarComboSumBetCount(Content []interface{}) int {
	Count := 0
	if len(Content) > 0 {
		Count = len(ShiShiTwoStarComboSumBet(Content))
	}
	return Count
}

/***********************************************************************
 * 	$TSCSB_Content 資料陣列 (Array)
 ***********************************************************************/
func ShiShiTwoStarComboSumBet(Content []interface{}) []interface{} { //列全部
	var ContentAll []interface{}
	for i1 := 0; i1 < len(Content); i1++ {
		for i2 := 0; i2 < 10; i2++ {
			for i3 := 0; i3 < 10; i3++ {
				if i2 != i3 { //避免对子被塞進去ㄠ
					item, err := strconv.Atoi(Content[i1].(string))
					if err != nil {
						panic(err)
					}

					if (i2 + i3) == item {
						ContentTemp := []string{strconv.Itoa(i2), strconv.Itoa(i3)}

						if !In_array2(ContentTemp, ContentAll) {
							ContentAll = append(ContentAll, ContentTemp)

						}
						// sort.Strings(Temp)
						// all, _ := json.Marshal(ContentAll)
						// tmp, _ := json.Marshal(Temp)

						// if strings.Index(string(all), string(tmp)) > -1 {
						// 	ContentAll = append(ContentAll, Temp)
						// }

					}
				}
			}
		}
	}
	return ContentAll
}

/*	二星 (后二|前二) -> 组选和值 [END]
***********************************************************************/

/***********************************************************************
 *	二星 (后二|前二) -> 组选包胆 [START] */

/***********************************************************************
 * 	$TSCBDBC_Content 資料陣列 (Array)
 ***********************************************************************/
func ShiShiTwoStarComboBaoDnoBetCount(Content []interface{}) int {
	Count := 0
	if len(Content) == 1 {

		Count = len(ShiShiTwoStarComboBaoDnoBet(Content))
	}

	return Count
}

/***********************************************************************
 * 	$TSCBDB_Content 資料陣列 (Array)
 ***********************************************************************/
func ShiShiTwoStarComboBaoDnoBet(Content []interface{}) []interface{} { //列全部

	var ContentAll []interface{}
	for i1 := 0; i1 < len(Content); i1++ {
		for i2 := 0; i2 < 10; i2++ {
			for i3 := 0; i3 < 10; i3++ {
				if i2 != i3 { //避免对子被塞進去
					if len(Array_unique_str([]string{strconv.Itoa(i2), strconv.Itoa(i3)})) == 2 {
						ContentTemp := []string{strconv.Itoa(i2), strconv.Itoa(i3)}
						if In_array(Content[i1], ContentTemp) {
							if !In_array2(ContentTemp, ContentAll) {
								ContentAll = append(ContentAll, ContentTemp)

							}
						}
						// sort.Strings(ContentTemp)
						// all, _ := json.Marshal(ContentAll)
						// tmp, _ := json.Marshal(ContentTemp)
						// if strings.Index(string(all), string(tmp)) > -1 {
						// 	ContentAll = append(ContentAll, ContentTemp)
						// }

					}
				}
			}
		}
	}

	return ContentAll
}

/*	二星 (后二|前二) -> 组选包胆 [END]
 ***********************************************************************/

/*	二星 (后二|前二) [END]
***********************************************************************/

/***********************************************************************
 *	不定位 [START] */

/***********************************************************************
 *	不定位 -> 一码 (前三|后三|中三|前四|后四|五星) [START] */

/***********************************************************************
 * 	$APP1BC_Content 資料陣列 (Array)
 ***********************************************************************/
func ShiShiAnyPositionPick1BetCount(Content []interface{}) int {
	Count := 0
	if len(Content) > 0 {
		Count = len(Content)
	}
	return Count
}

/*	不定位 -> 一码 (前三|后三|中三|前四|后四|五星) [END]
 ***********************************************************************/

/***********************************************************************
 *	不定位 -> 二码 (前三|后三|中三|前四|后四|五星) [START] */

/***********************************************************************
 * 	$APP2BC_Content 資料陣列 (Array)
 ***********************************************************************/
func ShiShiAnyPositionPick2BetCount(Content []interface{}) int {
	Count := 0
	if len(Content) >= 2 { //至少要有二個

		Count = len(ShiShiAnyPositionPick2Bet(Content))
	}
	return Count
}

/***********************************************************************
 * 	列全部的部份，同『 二星 -> 组选复式 』
 ***********************************************************************/
func ShiShiAnyPositionPick2Bet(Content []interface{}) []interface{} {
	return ShiShiTwoStarComboComplexBet(Content)
}

/*	不定位 -> 二码 (前三|后三|中三|前四|后四|五星) [END]
 ***********************************************************************/

/***********************************************************************
 *	不定位 -> 五星三码 [START] */

/***********************************************************************
 * 	$APP3BC_Content 資料陣列 (Array)
 ***********************************************************************/
func ShiShiAnyPositionPick3BetCount(Content []interface{}) int {
	Count := 0
	if len(Content) >= 3 {
		Count = len(ShiShiAnyPositionPick3Bet(Content))
	}
	return Count
}

/***********************************************************************
 * 	列全部的部份，同『 三星 -> 组六复式 』
 ***********************************************************************/
func ShiShiAnyPositionPick3Bet(Content []interface{}) []interface{} {
	return ShiShiThreeStarCombo6ComplexBet(Content)
}

/*	不定位 -> 五星三码 [END]
 ***********************************************************************/

/*	不定位 [END]
 ***********************************************************************/

/***********************************************************************
*	双面/串关 [START] */

/***********************************************************************
 *	双面/串关 -> 前二、后二大小单双 [START] */

/***********************************************************************
 * 	$TSSBC_Content_1 資料陣列 (Array)
 * 	$TSSBC_Content_2 資料陣列 (Array)
 ***********************************************************************/
func ShiShiTwoStarSideBetCount(Content1 []string, Content2 []string) int {
	Count := 0
	//判斷有沒有值
	if len(Content1) > 0 && len(Content2) > 0 {
		Count = len(ShiShiTwoStarSideBet(Content1, Content2))
	}
	return Count
}

/***********************************************************************
 * 	$TSSB_Content_1 資料陣列 (Array)
 * 	$TSSB_Content_2 資料陣列 (Array)
 ***********************************************************************/
func ShiShiTwoStarSideBet(Content1 []string, Content2 []string) []interface{} {
	var ContentAll []interface{}
	for i1 := 0; i1 < len(Content1); i1++ {
		for i2 := 0; i2 < len(Content2); i2++ {
			ContentAll = append(ContentAll, []string{Content1[i1], Content2[i2]})
		}
	}
	return ContentAll
}

/*	双面/串关 -> 前二、后二大小单双 [END]
 ***********************************************************************/

/***********************************************************************
 *	双面/串关 -> 前三、后三大小单双 [START] */

/***********************************************************************
 * 	$HSSBC_Content_1 資料陣列 (Array)
 * 	$HSSBC_Content_2 資料陣列 (Array)
 ***********************************************************************/
func ShiShiThreeStarSideBetCount(Content1 []string, Content2 []string, Content3 []string) int {
	Count := 0
	//判斷有沒有值
	if len(Content1) > 0 && len(Content2) > 0 && len(Content3) > 0 {
		Count = len(ShiShiThreeStarSideBet(Content1, Content2, Content3))
	}
	return Count
}

/***********************************************************************
 * 	$HSSB_Content_1 資料陣列 (Array)
 * 	$HSSB_Content_2 資料陣列 (Array)
 ***********************************************************************/
func ShiShiThreeStarSideBet(Content1 []string, Content2 []string, Content3 []string) []interface{} {
	var ContentAll []interface{}
	for i1 := 0; i1 < len(Content1); i1++ {
		for i2 := 0; i2 < len(Content2); i2++ {
			for i3 := 0; i3 < len(Content3); i3++ {
				ContentAll = append(ContentAll, []string{Content1[i1], Content2[i2], Content3[i3]})
			}
		}
	}
	return ContentAll
}

/*	双面/串关 -> 前三、后三大小单双 [END]
 ***********************************************************************/

/*	双面/串关 [END]
 ***********************************************************************/

/***********************************************************************
 *	任选二 [START] */

/***********************************************************************
 *	任选二 -> 直选复式 [START] */

/***********************************************************************
 * 	$AP2BC_Content_1 資料陣列 (Array)
 * 	$AP2BC_Content_2 資料陣列 (Array)
 * 	$AP2BC_Content_3 資料陣列 (Array)
 * 	$AP2BC_Content_4 資料陣列 (Array)
 * 	$AP2BC_Content_5 資料陣列 (Array)
 ***********************************************************************/
func ShiShiAnyPick2BetCount(Content1 []interface{}, Content2 []interface{}, Content3 []interface{}, Content4 []interface{}, Content5 []interface{}) int {
	Count := 0
	Count = len(ShiShiAnyPick2Bet(Content1, Content2, Content3, Content4, Content5))

	return Count
}

/***********************************************************************
 * 	$AP2B_Content_1 資料陣列 (Array)
 * 	$AP2B_Content_2 資料陣列 (Array)
 * 	$AP2B_Content_3 資料陣列 (Array)
 * 	$AP2B_Content_4 資料陣列 (Array)
 * 	$AP2B_Content_5 資料陣列 (Array)
 ***********************************************************************/
func ShiShiAnyPick2Bet(Content1 []interface{}, Content2 []interface{}, Content3 []interface{}, Content4 []interface{}, Content5 []interface{}) []interface{} {
	var ContentAll []interface{}
	var ContentTemp []interface{}

	if len(Content1) > 0 {
		for i := 0; i < len(Content1); i++ {
			ContentTemp = append(ContentTemp, []string{Content1[i].(string), "0"})
		}

	}

	if len(Content2) > 0 {
		for i := 0; i < len(Content2); i++ {
			ContentTemp = append(ContentTemp, []string{Content2[i].(string), "1"})
		}

	}
	if len(Content3) > 0 {
		for i := 0; i < len(Content3); i++ {
			ContentTemp = append(ContentTemp, []string{Content3[i].(string), "2"})
		}

	}
	if len(Content4) > 0 {
		for i := 0; i < len(Content4); i++ {
			ContentTemp = append(ContentTemp, []string{Content4[i].(string), "3"})
		}

	}
	if len(Content5) > 0 {
		for i := 0; i < len(Content5); i++ {
			ContentTemp = append(ContentTemp, []string{Content5[i].(string), "4"})
		}

	}

	for i1 := 0; i1 < len(ContentTemp); i1++ {
		for i2 := i1 + 1; i2 < len(ContentTemp); i2++ {
			if len(Array_unique_str([]string{ContentTemp[i1].([]string)[1], ContentTemp[i2].([]string)[1]})) == 2 { //位數不可以相同
				var tmp []interface{}
				tmp = append(tmp, ContentTemp[i1])
				tmp = append(tmp, ContentTemp[i2])
				ContentAll = append(ContentAll, tmp)
			}
		}
	}
	return ContentAll
}

/*	任选二 -> 直选复式 [END]
 ***********************************************************************/

/***********************************************************************
 *	任选二 -> 直选单式 [START] */

/***********************************************************************
 * 	$AP2SBC_Content 資料陣列 (String)
 * 	$AP2SBC_Options 資料陣列 (Array)
 ***********************************************************************/
func ShiShiAnyPick2SimpleBetCount(Content string, Options []interface{}) int {
	Count := 0
	if len(Content) > 0 {
		if len(Options) >= 2 {
			Count = len(ShiShiAnyPick2SimpleBet(Content))
		}
	}
	return Count * len(AllOptionsCompose(Options, 2))
}

/***********************************************************************
 * 	列全部的部份，同『 二星 (后二|前二) -> 直选单式 』
 ***********************************************************************/
func ShiShiAnyPick2SimpleBet(Content string) []interface{} {
	return ShiShiTwoStarSimpleBet(Content)
}

/*	任选二 -> 直选单式 [END]
 ***********************************************************************/

/***********************************************************************
 *	任选二 -> 直选和值 [START] */

/***********************************************************************
 * 	$AP2SBC_Content 資料陣列 (Array)
 * 	$AP2SBC_Options 資料陣列 (Array)
 ***********************************************************************/
func ShiShiAnyPick2SumBetCount(Content []interface{}, Options []interface{}) int {
	Count := 0
	if len(Content) > 0 && len(Options) > 0 {
		Count = len(ShiShiAnyPick2SumBet(Content)) * len(AllOptionsCompose(Options, 2))
	}
	return Count
}

/***********************************************************************
 * 	列全部的部份，同『 二星 (后二|前二) -> 直选和值 』
 ***********************************************************************/
func ShiShiAnyPick2SumBet(Content []interface{}) []interface{} {
	return ShiShiTwoStarSumBet(Content)
}

/*	任选二 -> 直选和值 [END]
 ***********************************************************************/

/***********************************************************************
 *	任选二 -> 組选复式 [START] */

/***********************************************************************
 * 	$AP2CBC_Content 資料陣列 (Array)
 * 	$AP2CBC_Options 資料陣列 (Array)
 ***********************************************************************/
func ShiShiAnyPick2ComboBetCount(Content []interface{}, Options []interface{}) int {
	Count := 0
	if len(Content) > 0 && len(Options) > 0 {
		if len(Options) >= 2 {
			Count = len(ShiShiAnyPick2ComboBet(Content)) * len(AllOptionsCompose(Options, 2))
		}
	}
	return Count
}

/***********************************************************************
 * 	列全部的部份，同『 二星 (后二|前二) -> 组选复式 』
 ***********************************************************************/
func ShiShiAnyPick2ComboBet(Content []interface{}) []interface{} {
	return ShiShiTwoStarComboComplexBet(Content)
}

/*	任选二 -> 組选复式 [END]
 ***********************************************************************/

/***********************************************************************
 *	任选二 -> 组选单式 [START] */

/***********************************************************************
 * 	$AP2CLBC_Content 資料陣列 (Array)
 * 	$AP2CLBC_Options 資料陣列 (Array)
 ***********************************************************************/
func ShiShiAnyPick2ComboSimpleBetCount(Content string, Options []interface{}) int {
	Count := 0
	if len(Content) > 0 {
		if len(Options) >= 2 {
			Count = len(ShiShiAnyPick2ComboSimpleBet(Content)) * len(AllOptionsCompose(Options, 2))
		}
	}
	return Count
}

/***********************************************************************
 * 	列全部的部份，同『 二星 (后二|前二) -> 组选单式 』
 ***********************************************************************/
func ShiShiAnyPick2ComboSimpleBet(Content string) []interface{} {
	return ShiShiTwoStarComboSimpleBet(Content)
}

/*	任选二 -> 组选单式 [END]
 ***********************************************************************/
/***********************************************************************
 *	任选二 -> 组选和值 [START] */

/***********************************************************************
 * 	$AP2CSBC_Content 資料陣列 (Array)
 * 	$AP2CSBC_Options 資料陣列 (Array)
 ***********************************************************************/
func ShiShiAnyPick2ComboSumBetCount(Content []interface{}, Options []interface{}) int {
	Count := 0
	if len(Content) > 0 && len(Options) > 0 {
		Count = len(ShiShiAnyPick2ComboSumBet(Content)) * len(AllOptionsCompose(Options, 2))
	}
	return Count
}

/***********************************************************************
 * 	列全部的部份，同『 二星 (后二|前二) -> 组选和值 』
 ***********************************************************************/
func ShiShiAnyPick2ComboSumBet(Content []interface{}) []interface{} {

	return ShiShiTwoStarComboSumBet(Content)
}

/*	任选二 -> 组选和值 [END]
 ***********************************************************************/

/*	任选二 [END]
 ***********************************************************************/

/***********************************************************************
 *	任选三 [START] */

/***********************************************************************
 *	任选三 -> 直选复式 [START] */

/***********************************************************************
 * 	$AP3BC_Content_1 資料陣列 (Array)
 * 	$AP3BC_Content_2 資料陣列 (Array)
 * 	$AP3BC_Content_3 資料陣列 (Array)
 * 	$AP3BC_Content_4 資料陣列 (Array)
 * 	$AP3BC_Content_5 資料陣列 (Array)
 ***********************************************************************/
func ShiShiAnyPick3BetCount(Content1 []interface{}, Content2 []interface{}, Content3 []interface{}, Content4 []interface{}, Content5 []interface{}) int {
	Count := 0
	Count = len(ShiShiAnyPick3Bet(Content1, Content2, Content3, Content4, Content5))
	return Count
}

/***********************************************************************
 * 	$AP3B_Content_1 資料陣列 (Array)
 * 	$AP3B_Content_2 資料陣列 (Array)
 * 	$AP3B_Content_3 資料陣列 (Array)
 * 	$AP3B_Content_4 資料陣列 (Array)
 * 	$AP3B_Content_5 資料陣列 (Array)
 ***********************************************************************/
func ShiShiAnyPick3Bet(Content1 []interface{}, Content2 []interface{}, Content3 []interface{}, Content4 []interface{}, Content5 []interface{}) []interface{} {
	var ContentAll []interface{}
	var ContentTemp []interface{}

	if len(Content1) > 0 {
		for i := 0; i < len(Content1); i++ {
			ContentTemp = append(ContentTemp, []string{Content1[i].(string), "0"})
		}

	}

	if len(Content2) > 0 {
		for i := 0; i < len(Content2); i++ {
			ContentTemp = append(ContentTemp, []string{Content2[i].(string), "1"})
		}

	}
	if len(Content3) > 0 {
		for i := 0; i < len(Content3); i++ {
			ContentTemp = append(ContentTemp, []string{Content3[i].(string), "2"})
		}

	}
	if len(Content4) > 0 {
		for i := 0; i < len(Content4); i++ {
			ContentTemp = append(ContentTemp, []string{Content4[i].(string), "3"})
		}

	}
	if len(Content5) > 0 {
		for i := 0; i < len(Content5); i++ {
			ContentTemp = append(ContentTemp, []string{Content5[i].(string), "4"})
		}

	}

	for i1 := 0; i1 < len(ContentTemp); i1++ {
		for i2 := i1 + 1; i2 < len(ContentTemp); i2++ {
			for i3 := i2 + 1; i3 < len(ContentTemp); i3++ {

				if len(Array_unique_str([]string{ContentTemp[i1].([]string)[1], ContentTemp[i2].([]string)[1], ContentTemp[i3].([]string)[1]})) == 3 { //位數不可以相同
					var tmp []interface{}
					tmp = append(tmp, ContentTemp[i1])
					tmp = append(tmp, ContentTemp[i2])
					tmp = append(tmp, ContentTemp[i3])
					ContentAll = append(ContentAll, tmp)
				}
			}
		}
	}
	return ContentAll
}

/*	任选三 -> 直选复式 [END]
 ***********************************************************************/

/***********************************************************************
 *	任选三 -> 直选单式 [START] */

/***********************************************************************
 * 	$AP3SBC_Content_1 資料陣列 (Array)
 * 	$AP3SBC_Content_2 資料陣列 (Array)
 * 	$AP3SBC_Content_3 資料陣列 (Array)
 * 	$AP3SBC_Options 資料陣列 (Array)
 ***********************************************************************/
func ShiShiAnyPick3SimpleBetCount(Content string, Options []interface{}) int {

	Count := 0

	if len(Content) > 0 {
		if len(Options) >= 3 {
			Count = len(ShiShiAnyPick3SimpleBet(Content))
		}
	}
	return Count * len(AllOptionsCompose(Options, 3))

}

/***********************************************************************
 * 	列全部的部份，同『 三星 (后三、前三、中三) -> 直选单式 』
 ***********************************************************************/
func ShiShiAnyPick3SimpleBet(Content string) []interface{} {
	return ShiShiThreeStarSimpleBet(Content)
}

/*	任选三 -> 直选单式 [END]
 ***********************************************************************/
/***********************************************************************
 *	任选三 -> 直选和值 [START] */

/***********************************************************************
 *	$AP3SBC_Content 資料陣列 (Array)
*	$AP3SBC_Options 資料陣列 (Array)
***********************************************************************/
func ShiShiAnyPick3SumBetCount(Content []interface{}, Options []interface{}) int {
	Count := 0
	if len(Content) > 0 { //如果裡面有值
		if len(Options) >= 3 {
			Count = len(ShiShiAnyPick3SumBet(Content)) * len(AllOptionsCompose(Options, 3))
		}
	}
	return Count
}

/***********************************************************************
 * 	列全部的部份，同『 三星 (后三、前三、中三) -> 直选和值 』
 ***********************************************************************/
func ShiShiAnyPick3SumBet(Content []interface{}) []interface{} {
	return ShiShiThreeStarSumBet(Content)
}

/*	任选三 -> 直选和值 [END]
***********************************************************************/
/***********************************************************************
 *	任选三 -> 组三复式 [START] */

/***********************************************************************
 * 	$AP3C3CBC_Content 資料陣列 (Array)
 * 	$AP3C3CBC_Options 資料陣列 (Array)
 ***********************************************************************/
func ShiShiAnyPick3Combo3ComplexBetCount(Content []interface{}, Options []interface{}) int {

	Count := 0
	if len(Content) > 0 && len(Options) > 0 {
		if len(Content) >= 2 && len(Options) >= 3 {
			Count = len(ShiShiAnyPick3Combo3ComplexBet(Content)) * len(AllOptionsCompose(Options, 3))
		}
	}
	return Count
}

/***********************************************************************
 * 	列全部的部份，同『 三星 (后三、前三、中三) -> 组三复式 』
 ***********************************************************************/
func ShiShiAnyPick3Combo3ComplexBet(Content []interface{}) []interface{} {
	return ShiShiThreeStarCombo3ComplexBet(Content)
}

/*	任选三 -> 组三复式 [END]
 ***********************************************************************/
/***********************************************************************
 *	任选三 -> 组六复式 [START] */

/***********************************************************************
 * 	$AP3C6CBC_Content 資料陣列 (Array)
 ***********************************************************************/
func ShiShiAnyPick3Combo6ComplexBetCount(Content []interface{}, Options []interface{}) int {

	Count := 0
	if len(Content) > 0 && len(Options) > 0 {
		if len(Content) >= 2 && len(Options) >= 3 {
			Count = len(ShiShiAnyPick3Combo6ComplexBet(Content)) * len(AllOptionsCompose(Options, 3))

		}
	}
	return Count
}

/***********************************************************************
 * 	列全部的部份，同『 三星 (后三、前三、中三) -> 组六复式 』
 ***********************************************************************/
func ShiShiAnyPick3Combo6ComplexBet(Content []interface{}) []interface{} {

	return ShiShiThreeStarCombo6ComplexBet(Content)
}

/*	任选三 -> 组六复式 [END]
 ***********************************************************************/
/***********************************************************************
 *	任选三 -> 組选和值 [START] */

/***********************************************************************
 * 	$AP3CSBC_Content 資料陣列 (Array)
 ***********************************************************************/
func ShiShiAnyPick3ComboSumBetCount(Content []interface{}, Options []interface{}) int {
	Count := 0
	if len(Content) > 0 {
		if len(Options) >= 0 {
			Count = len(ShiShiAnyPick3ComboSumBet(Content)) * len(AllOptionsCompose(Options, 3))
		}
	}
	return Count
}

/***********************************************************************
 * 	列全部的部份，同『 三星 (后三、前三、中三) -> 组选和值 』
 ***********************************************************************/
func ShiShiAnyPick3ComboSumBet(Content []interface{}) []interface{} {
	return ShiShiThreeStarComboSumBet(Content)
}

/*	任选三 -> 組选和值 [END]
 ***********************************************************************/

/***********************************************************************
 *	任选三 -> 組三单式 [START] */

/***********************************************************************
 * 	$AP3C3SBC_Content 資料陣列 (Array)
 * 	$AP3C3SBC_Options 資料陣列 (Array)
 ***********************************************************************/
func ShiShiAnyPick3Combo3SimpleBetCount(Content string, Options []interface{}) int {

	Count := 0
	if len(Content) > 0 {
		Count = len(ShiShiAnyPick3Combo3SimpleBet(Content)) * len(AllOptionsCompose(Options, 3))
	}
	return Count
}

/***********************************************************************
 * 	列全部的部份，同『三星 (后三、前三、中三) -> 組三单式 』
 ***********************************************************************/
func ShiShiAnyPick3Combo3SimpleBet(Content string) []interface{} {

	return ShiShiThreeStarCombo3SimpleBet(Content)
}

/*	任选三 -> 組三单式 [END]
 ***********************************************************************/
/***********************************************************************
 *	任选三 -> 组六单式 [START] */

/***********************************************************************
 * 	$AP3C6SBC_Content 資料陣列 (Array)
 ***********************************************************************/
func ShiShiAnyPick3Combo6SimpleBetCount(Content string, Options []interface{}) int {
	Count := 0
	if len(Content) > 0 {
		Count = len(ShiShiAnyPick3Combo6SimpleBet(Content)) * len(AllOptionsCompose(Options, 3))
	}
	return Count
}

/***********************************************************************
 * 	列全部的部份，同『三星 (后三、前三、中三) -> 组六单式 』
 ***********************************************************************/
func ShiShiAnyPick3Combo6SimpleBet(Content string) []interface{} {
	return ShiShiThreeStarCombo6SimpleBet(Content)
}

/*	任选三 -> 组六单式 [END]
 ***********************************************************************/

/***********************************************************************
 *	任选三 -> 混合组选 [START] */

/***********************************************************************
 * 	$AP3CMBC_Content 資料陣列 (Array)
 ***********************************************************************/
func ShiShiAnyPick3ComboMixBetCount(Content string, Options []interface{}) int {
	Count := 0
	if len(Content) > 0 {
		Count = len(ShiShiAnyPick3ComboMixBet(Content)) * len(AllOptionsCompose(Options, 3))
	}
	return Count
}

/***********************************************************************
 * 	列全部的部份，同『 三星 (后三、前三、中三) -> 混合组选 』
 ***********************************************************************/
func ShiShiAnyPick3ComboMixBet(Content string) []interface{} {

	return ShiShiThreeStarComboMixBet(Content)
}

/*	任选三 -> 混合组选 [END]
 ***********************************************************************/
/*	任选三 [END]
***********************************************************************/

/***********************************************************************
 *	任选四 [START] */

/***********************************************************************
 *	任选四 -> 直选复式 [START] */

/***********************************************************************
 * 	$AP4BC_Content_1 資料陣列 (Array)
 * 	$AP4BC_Content_2 資料陣列 (Array)
 * 	$AP4BC_Content_3 資料陣列 (Array)
 * 	$AP4BC_Content_4 資料陣列 (Array)
 * 	$AP4BC_Content_5 資料陣列 (Array)
 ***********************************************************************/
func ShiShiAnyPick4BetCount(Content1 []interface{}, Content2 []interface{}, Content3 []interface{}, Content4 []interface{}, Content5 []interface{}) int {
	Count := 0
	Count = len(ShiShiAnyPick4Bet(Content1, Content2, Content3, Content4, Content5))
	return Count
}

/***********************************************************************
 * 	$AP4B_Content_1 資料陣列 (Array)
 * 	$AP4B_Content_2 資料陣列 (Array)
 * 	$AP4B_Content_3 資料陣列 (Array)
 * 	$AP4B_Content_4 資料陣列 (Array)
 * 	$AP4B_Content_5 資料陣列 (Array)
 ***********************************************************************/
func ShiShiAnyPick4Bet(Content1 []interface{}, Content2 []interface{}, Content3 []interface{}, Content4 []interface{}, Content5 []interface{}) []interface{} {

	var ContentAll []interface{}
	var ContentTemp []interface{}

	if len(Content1) > 0 {
		for i := 0; i < len(Content1); i++ {
			ContentTemp = append(ContentTemp, []string{Content1[i].(string), "0"})
		}

	}

	if len(Content2) > 0 {
		for i := 0; i < len(Content2); i++ {
			ContentTemp = append(ContentTemp, []string{Content2[i].(string), "1"})
		}

	}
	if len(Content3) > 0 {
		for i := 0; i < len(Content3); i++ {
			ContentTemp = append(ContentTemp, []string{Content3[i].(string), "2"})
		}

	}
	if len(Content4) > 0 {
		for i := 0; i < len(Content4); i++ {
			ContentTemp = append(ContentTemp, []string{Content4[i].(string), "3"})
		}

	}
	if len(Content5) > 0 {
		for i := 0; i < len(Content5); i++ {
			ContentTemp = append(ContentTemp, []string{Content5[i].(string), "4"})
		}

	}

	for i1 := 0; i1 < len(ContentTemp); i1++ {
		for i2 := i1 + 1; i2 < len(ContentTemp); i2++ {
			for i3 := i2 + 1; i3 < len(ContentTemp); i3++ {
				for i4 := i3 + 1; i4 < len(ContentTemp); i4++ {

					if len(Array_unique_str([]string{ContentTemp[i1].([]string)[1], ContentTemp[i2].([]string)[1], ContentTemp[i3].([]string)[1], ContentTemp[i4].([]string)[1]})) == 4 { //位數不可以相同
						var tmp []interface{}
						tmp = append(tmp, ContentTemp[i1])
						tmp = append(tmp, ContentTemp[i2])
						tmp = append(tmp, ContentTemp[i3])
						tmp = append(tmp, ContentTemp[i4])
						ContentAll = append(ContentAll, tmp)
					}
				}
			}
		}
	}
	return ContentAll

}

/*	任选四 -> 直选复式 [END]
 ***********************************************************************/

/***********************************************************************
 *	任选四 -> 直选单式 [START] */

/***********************************************************************
 * 	$AP4SBC_Content 資料陣列 (String)
 * 	$AP4SBC_Options 資料陣列 (Array)
 ***********************************************************************/
func ShiShiAnyPick4SimpleBetCount(Content string, Options []interface{}) int {
	Count := 0
	if len(Content) > 0 {
		if len(Options) >= 4 {
			Count = len(ShiShiAnyPick4SimpleBet(Content))
		}
	}
	return Count * len(AllOptionsCompose(Options, 4))
}

/***********************************************************************
 * 	列全部的部份，同『 四星 -> 直选单式 』
 ***********************************************************************/
func ShiShiAnyPick4SimpleBet(Content string) []interface{} {

	return ShiShiFourStarSimpleBet(Content)
}

/*	任选四 -> 直选单式 [END]
 ***********************************************************************/
/***********************************************************************
 *	任选四 -> 组选24 [START] */

/***********************************************************************
 * 	$AP4C24BC_Content 資料陣列 (Array)
 * 	$AP4C24BC_Options 資料陣列 (Array)
 ***********************************************************************/
func ShiShiAnyPick4Combo24BetCount(Content []interface{}, Options []interface{}) int {
	Count := 0
	if len(Content) > 0 && len(Options) > 0 {
		if len(Content) >= 4 && len(Options) >= 4 {
			Count = len(ShiShiFourStarCombo24Bet(Content)) * len(AllOptionsCompose(Options, 4))
		}
	}
	return Count
}

/***********************************************************************
 * 	列全部的部份，同『 四星 -> 组选24 』
 ***********************************************************************/
func ShiShiAnyPick4Combo24Bet(Content []interface{}) []interface{} {
	return ShiShiFourStarCombo24Bet(Content)
}

/*	任选四 -> 组选24 [END]
 ***********************************************************************/
/***********************************************************************
 *	任选四 -> 组选12 [START] */

/***********************************************************************
 * 	$AP4C12BC_Content_1 資料陣列 (Array) 雙重號
 * 	$AP4C12BC_Content_2 資料陣列 (Array) 單號
 * 	$AP4C12BC_Options   資料陣列 (Array) 位置
 ***********************************************************************/
func ShiShiAnyPick4Combo12BetCount(Content1 []interface{}, Content2 []interface{}, Options []interface{}) int {

	Count := 0
	if len(Content1) > 0 && len(Content2) > 0 && len(Options) > 0 {
		if len(Content1) >= 1 && len(Content2) >= 2 && len(Options) >= 4 {

			Count = len(ShiShiAnyPick4Combo12Bet(Content1, Content2)) * len(AllOptionsCompose(Options, 4))
		}
	}
	return Count
}

/***********************************************************************
 * 	列全部的部份，同『 四星 -> 组选12 』
 ***********************************************************************/
func ShiShiAnyPick4Combo12Bet(Content1 []interface{}, Content2 []interface{}) []interface{} {
	return ShiShiFourStarCombo12Bet(Content1, Content2)
}

/*	任选四 -> 组选12 [END]
 ***********************************************************************/
/***********************************************************************
 *	任选四 -> 组选6 [START] */

/***********************************************************************
 * 	$AP4C6BC_Content 資料陣列 (Array) 雙重號
 * 	$AP4C6BC_Options 資料陣列 (Array) 位置
 ***********************************************************************/
func ShiShiAnyPick4Combo6BetCount(Content1 []interface{}, Options []interface{}) int {
	Count := 0
	if len(Content1) > 0 && len(Options) > 0 {
		if len(Content1) >= 2 && len(Options) >= 4 {
			Count = len(ShiShiAnyPick4Combo6Bet(Content1)) * len(AllOptionsCompose(Options, 4))
		}
	}
	return Count
}

/***********************************************************************
 * 	列全部的部份，同『 四星 -> 组选6 』
 ***********************************************************************/
func ShiShiAnyPick4Combo6Bet(Content1 []interface{}) []interface{} {
	return ShiShiFourStarCombo6Bet(Content1)
}

/*	任选四 -> 组选6 [END]
 ***********************************************************************/

/***********************************************************************
 *	任选四 -> 组选4 [START] */

/***********************************************************************
 * 	$AP4C4BC_Content_1 資料陣列 (Array) 三重號
 * 	$AP4C4BC_Content_2 資料陣列 (Array) 單號
 * 	$AP4C4BC_Options   資料陣列 (Array) 位置
 ***********************************************************************/
func ShiShiAnyPick4Combo4BetCount(Content1 []interface{}, Content2 []interface{}, Options []interface{}) int {

	Count := 0
	if len(Content1) > 0 && len(Content2) > 0 {
		if len(Content1) >= 1 && len(Content2) >= 1 && len(Options) >= 4 {
			Count = len(ShiShiAnyPick4Combo4Bet(Content1, Content2)) * len(AllOptionsCompose(Options, 4))
		}
	}

	return Count
}

/***********************************************************************
 * 	列全部的部份，同『 四星 -> 组选4 』
 ***********************************************************************/
func ShiShiAnyPick4Combo4Bet(Content1 []interface{}, Content2 []interface{}) []interface{} {
	return ShiShiFourStarCombo4Bet(Content1, Content2)
}

/*	任选四 -> 组选4 [END]
 ***********************************************************************/

/*	任选四 [END]
 ***********************************************************************/
/***********************************************************************
 *	任选ｎ -> 位數計算 [START] */

func AllOptionsCompose(Options []interface{}, Pick int) []interface{} {
	var OptionsAll []interface{}
	if len(Options) >= Pick {
		if Pick == 2 {
			for i1 := 0; i1 < len(Options); i1++ {
				for i2 := i1 + 1; i2 < len(Options); i2++ {
					OptionsAll = append(OptionsAll, []string{Options[i1].(string), Options[i2].(string)})
					//array_push($AOC_OptionsAll,[(int)$AOC_Options[$i1],(int)$AOC_Options[$i2]]);
				}
			}
		} else if Pick == 3 {
			for i1 := 0; i1 < len(Options); i1++ {
				for i2 := i1 + 1; i2 < len(Options); i2++ {
					for i3 := i2 + 1; i3 < len(Options); i3++ {
						OptionsAll = append(OptionsAll, []string{Options[i1].(string), Options[i2].(string), Options[i3].(string)})
						//	array_push($AOC_OptionsAll,[(int)$AOC_Options[$i1],(int)$AOC_Options[$i2],(int)$AOC_Options[$i3]]);
					}
				}
			}
		} else if Pick == 4 {
			for i1 := 0; i1 < len(Options); i1++ {
				for i2 := i1 + 1; i2 < len(Options); i2++ {
					for i3 := i2 + 1; i3 < len(Options); i3++ {
						for i4 := i3 + 1; i4 < len(Options); i4++ {
							OptionsAll = append(OptionsAll, []string{Options[i1].(string), Options[i2].(string), Options[i3].(string), Options[i4].(string)})
						}
					}
				}
			}
		}
	}
	return OptionsAll
}

/***********************************************************************
 * 	北京PK10 幸运飞艇 官方玩法的的算注數方式
 ***********************************************************************/

/***********************************************************************
 *	前1玩法 -> 直选复式 [START] */

/***********************************************************************
 *	$OSBC_Content 資料陣列 (Array)
 ***********************************************************************/
//  public function PKTenOneStraightBetCount($OSBC_Content)
//  {
// 	 $OSBC_Count = 0;
// 	 if(isset($OSBC_Content))
// 	 {
// 		 $OSBC_Count = count($this->PKTenOneStraightBet($OSBC_Content));
// 	 }
// 	 return $OSBC_Count;
//  }
/***********************************************************************
 *	$OSB_Content 資料陣列 (Array)
 ***********************************************************************/
func PKTenOneStraightBet(Content []interface{}) []interface{} { //列全部
	var OptionsAll []interface{}
	for i := 0; i < len(Content); i++ {
		OptionsAll = append(OptionsAll, []string{Content[i].(string)})
		// array_push($OSB_ContentAll,[$OSB_Content[$i]]);
	}
	return OptionsAll
}

/*	前1玩法 -> 直选复式 [END]
 ***********************************************************************/

/***********************************************************************
 *	前2玩法 -> 直选复式 [START] */

/***********************************************************************
 *	$TSBC_Content_1 資料陣列 (Array)
 *	$TSBC_Content_2 資料陣列 (Array)
 ***********************************************************************/
//  public function PKTenTwoStraightBetCount($TSBC_Content_1,$TSBC_Content_2)
//  {
// 	 $OSBC_Count = 0;
// 	 if(isset($TSBC_Content_1) && isset($TSBC_Content_2))
// 	 {
// 		 $OSBC_Count = count($this->PKTenTwoStraightBet($TSBC_Content_1,$TSBC_Content_2));
// 	 }
// 	 return $OSBC_Count;
//  }
/***********************************************************************
 *	$TSB_Content_1 資料陣列 (Array)
 *	$TSB_Content_2 資料陣列 (Array)
 ***********************************************************************/
func PKTenTwoStraightBet(Content1 []interface{}, Content2 []interface{}) []interface{} { //列全部
	var ContentAll []interface{}
	for i1 := 0; i1 < len(Content1); i1++ {
		for i2 := 0; i2 < len(Content2); i2++ {
			if Content1[i1] != Content2[i2] { //避免重複的被塞進去
				ContentAll = append(ContentAll, []string{Content1[i1].(string), Content2[i2].(string)})
			}

		}
	}
	return ContentAll
}

/*	前2玩法 -> 直选复式 [END]
 ***********************************************************************/

/***********************************************************************
 *	前2玩法 -> 直选單式 [START] */

/***********************************************************************
 *	$TSBC_Content 資料陣列 (Array)
 ***********************************************************************/
//  public function PKTenTwoStraightSimpleBetCount($TSBC_Content)
//  {
// 	 $OSBC_Count = 0;
// 	 if(isset($TSBC_Content))
// 	 {
// 		 $OSBC_Count = count($this->PKTenTwoStraightSimpleBet($TSBC_Content));
// 	 }
// 	 return $OSBC_Count;
//  }
/***********************************************************************
 *	$TSB_Content 資料陣列 (Array)
 ***********************************************************************/
func PKTenTwoStraightSimpleBet(Content string) []interface{} { //列全部
	var ContentAll []interface{}

	Content1 := strings.Split(Content, ",")
	for i1 := 0; i1 < len(Content1); i1++ {
		Content2 := strings.Split(Content1[i1], " ")
		for i2 := 0; i2 < len(Content2); i2++ {
			if Content2[i2] != "10" {
				Content2[i2] = "0" + Content2[i2]
			}
		}
		//建立要儲存的組合為陣列
		ContentAll = append(ContentAll, []string{Content2[0], Content2[1]})

	}
	return ContentAll
}

/*	前2玩法 -> 直选單式 [END]
 ***********************************************************************/
/***********************************************************************
 *	前3玩法 -> 直选复式 [START] */

/***********************************************************************
 *	$HSBC_Content_1 資料陣列 (Array)
 *	$HSBC_Content_2 資料陣列 (Array)
 *	$HSBC_Content_3 資料陣列 (Array)
 ***********************************************************************/
// func PKTenThreeStraightBetCount($HSBC_Content_1,$HSBC_Content_2 ,$HSBC_Content_3) {
// 	$HSBC_Count = 0;
// 	if(isset($HSBC_Content_1) && isset($HSBC_Content_2) && isset($HSBC_Content_3))
// 	{
// 		$HSBC_Count = count($this->PKTenThreeStraightBet($HSBC_Content_1,$HSBC_Content_2,$HSBC_Content_3));
// 	}
// 	return $HSBC_Count;
// }
/***********************************************************************
 *	$HSB_Content_1 資料陣列 (Array)
 *	$HSB_Content_2 資料陣列 (Array)
 *	$HSB_Content_3 資料陣列 (Array)
 ***********************************************************************/
func PKTenThreeStraightBet(Content1 []interface{}, Content2 []interface{}, Content3 []interface{}) []interface{} { //列全部
	var ContentAll []interface{}
	for i1 := 0; i1 < len(Content1); i1++ {
		for i2 := 0; i2 < len(Content2); i2++ {
			for i3 := 0; i3 < len(Content3); i3++ {
				ContentTemp := []string{Content1[i1].(string), Content2[i2].(string), Content3[i3].(string)}
				if len(Array_unique_str(ContentTemp)) == 3 { //避免重複的被塞進去
					ContentAll = append(ContentAll, ContentTemp)
				}
			}
		}
	}
	return ContentAll
}

/*	前3玩法 -> 直选复式 [END]
***********************************************************************/

/***********************************************************************
 *	前3玩法 -> 直选單式 [START] */

/***********************************************************************
 *	$HSBC_Content 資料陣列 (Array)
 ***********************************************************************/
//   func PKTenThreeStraightSimpleBetCount($HSBC_Content)
//  {
// 	 $HSBC_Count = 0;
// 	 if(isset($HSBC_Content) )
// 	 {
// 		 $HSBC_Count = count($this->PKTenThreeStraightSimpleBet($HSBC_Content));
// 	 }
// 	 return $HSBC_Count;
//  }
/***********************************************************************
 *	$HSB_Content 資料陣列 (Array)
 ***********************************************************************/
func PKTenThreeStraightSimpleBet(Content string) []interface{} { //列全部
	var ContentAll []interface{}
	Content1 := strings.Split(Content, ",")
	for i1 := 0; i1 < len(Content1); i1++ {
		Content2 := strings.Split(Content1[i1], " ")
		for i2 := 0; i2 < len(Content2); i2++ {
			if Content2[i2] != "10" {
				Content2[i2] = "0" + Content2[i2]
			}
		}
		//建立要儲存的組合為陣列
		ContentAll = append(ContentAll, []string{Content2[0], Content2[1], Content2[2]})

		//  $HSB_Content_2 = explode(" ", trim($HSB_Content_1[$i1]));
		//  //建立要儲存的組合為陣列
		//  $HSB_ContentTemp = [$HSB_Content_2[0],$HSB_Content_2[1],$HSB_Content_2[2]];

		//  array_push($HSB_ContentAll,$HSB_ContentTemp);
	}
	return ContentAll
}

/*	前3玩法 -> 直选單式 [END]
 ***********************************************************************/

/***********************************************************************
 *	定位胆 -> 直选复式 [START] */

/***********************************************************************
 * 	$ASBC_Content_1 資料陣列 (Array)
 * 	$ASBC_Content_2 資料陣列 (Array)
 * 	$ASBC_Content_3 資料陣列 (Array)
 * 	$ASBC_Content_4 資料陣列 (Array)
 * 	$ASBC_Content_5 資料陣列 (Array)
 * 	$ASBC_Content_6 資料陣列 (Array)
 * 	$ASBC_Content_7 資料陣列 (Array)
 * 	$ASBC_Content_8 資料陣列 (Array)
 * 	$ASBC_Content_9 資料陣列 (Array)
 * 	$ASBC_Content_10 資料陣列 (Array)
 ***********************************************************************/
//  public function PKTenAllStraightBetCount($ASBC_Content_1, $ASBC_Content_2, $ASBC_Content_3, $ASBC_Content_4, $ASBC_Content_5, $ASBC_Content_6, $ASBC_Content_7, $ASBC_Content_8, $ASBC_Content_9, $ASBC_Content_10)
//  {
// 	 $ASBC_Count = 0;

// 	 $ASBC_Count = count($this->PKTenAllStraightBet($ASBC_Content_1, $ASBC_Content_2, $ASBC_Content_3, $ASBC_Content_4, $ASBC_Content_5, $ASBC_Content_6, $ASBC_Content_7, $ASBC_Content_8, $ASBC_Content_9, $ASBC_Content_10));

// 	 return $ASBC_Count;
//  }
/***********************************************************************
 * 	$ASB_Content_1 資料陣列 (Array)
 * 	$ASB_Content_2 資料陣列 (Array)
 * 	$ASB_Content_3 資料陣列 (Array)
 * 	$ASB_Content_4 資料陣列 (Array)
 * 	$ASB_Content_5 資料陣列 (Array)
 * 	$ASB_Content_6 資料陣列 (Array)
 * 	$ASB_Content_7 資料陣列 (Array)
 * 	$ASB_Content_8 資料陣列 (Array)
 * 	$ASB_Content_9 資料陣列 (Array)
 * 	$ASB_Content_10 資料陣列 (Array)
 ***********************************************************************/
func PKTenAllStraightBet(Content1 []interface{}, Content2 []interface{}, Content3 []interface{}, Content4 []interface{}, Content5 []interface{}, Content6 []interface{}, Content7 []interface{}, Content8 []interface{}, Content9 []interface{}, Content10 []interface{}) []interface{} {

	// var ContentAll []interface{}
	var ContentTemp []interface{}
	if len(Content1) > 0 {
		for i := 0; i < len(Content1); i++ {
			ContentTemp = append(ContentTemp, []string{Content1[i].(string), "0"})
		}
	}
	if len(Content2) > 0 {
		for i := 0; i < len(Content2); i++ {
			ContentTemp = append(ContentTemp, []string{Content2[i].(string), "1"})
		}
	}
	if len(Content3) > 0 {
		for i := 0; i < len(Content3); i++ {
			ContentTemp = append(ContentTemp, []string{Content3[i].(string), "2"})
		}

	}
	if len(Content4) > 0 {
		for i := 0; i < len(Content4); i++ {
			ContentTemp = append(ContentTemp, []string{Content4[i].(string), "3"})
		}

	}
	if len(Content5) > 0 {
		for i := 0; i < len(Content5); i++ {
			ContentTemp = append(ContentTemp, []string{Content5[i].(string), "4"})
		}
	}
	if len(Content6) > 0 {
		for i := 0; i < len(Content6); i++ {
			ContentTemp = append(ContentTemp, []string{Content6[i].(string), "5"})
		}
	}

	if len(Content7) > 0 {
		for i := 0; i < len(Content7); i++ {
			ContentTemp = append(ContentTemp, []string{Content7[i].(string), "6"})
		}

	}
	if len(Content8) > 0 {
		for i := 0; i < len(Content8); i++ {
			ContentTemp = append(ContentTemp, []string{Content8[i].(string), "7"})
		}

	}
	if len(Content9) > 0 {
		for i := 0; i < len(Content9); i++ {
			ContentTemp = append(ContentTemp, []string{Content9[i].(string), "8"})
		}

	}
	if len(Content10) > 0 {
		for i := 0; i < len(Content10); i++ {
			ContentTemp = append(ContentTemp, []string{Content10[i].(string), "9"})
		}

	}
	return ContentTemp
}

/*	定位胆 -> 直选复式 [END]
 ***********************************************************************/
/***********************************************************************
 *	同号 -> 二同号单选 [START] */

/***********************************************************************
 * 	$TKSBC_Content_1 資料陣列 (Array)
 * 	$TKSBC_Content_2 資料陣列 (Array)
 ***********************************************************************/
//  public function KuaiThreeTwoKindSingleBetCount($TKSBC_Content_1,$TKSBC_Content_2)
//  {
// 	 $TKSBC_Count = 0;
// 	 if(isset($TKSBC_Content_1) && isset($TKSBC_Content_2))
// 	 {
// 		 $TKSBC_Count = count($this->KuaiThreeTwoKindSingleBet($TKSBC_Content_1,$TKSBC_Content_2));
// 	 }
// 	 return $TKSBC_Count;
//  }
/***********************************************************************
 * 	$TKSB_Content_1 資料陣列 (Array)
 * 	$TKSB_Content_2 資料陣列 (Array)
 ***********************************************************************/
func KuaiThreeTwoKindSingleBet(Content1 []interface{}, Content2 []interface{}) []interface{} { //列全部
	var ContentAll []interface{}
	if len(Content1) > 0 && len(Content2) > 0 {
		for i1 := 0; i1 < len(Content1); i1++ {
			for i2 := 0; i2 < len(Content2); i2++ {
				//	ContentTemp := []string{strconv.Itoa(i2), strconv.Itoa(i3)}
				Content1Temp := ""
				if reflect.TypeOf(Content1).Kind() == reflect.Slice {
					Content1Temp = Content1[i1].([]interface{})[0].(string)
				} else {
					Content1Temp = Content1[i1].(string)[0 : 0+1]
				}

				if Content1Temp != Content2[i2].(string) {
					ContentAll = append(ContentAll, []string{Content1Temp, Content2[i2].(string)})
				}

			}
		}
	}
	return ContentAll
}

/*	同号 -> 二同号单选 [END]
***********************************************************************/

/***********************************************************************
 *	同号 -> 二同号单式 [START] */

/***********************************************************************
 * 	$TKLBC_Content 資料陣列 (String)
 ***********************************************************************/
//  public function KuaiThreeTwoKindSimpleBetCount($TKLBC_Content)
//  {
// 	 $TKLBC_Count = 0;
// 	 if(isset($TKLBC_Content))
// 	 {

// 		 $TKLBC_Count = count($this->KuaiThreeTwoKindSimpleBet($TKLBC_Content));

// 	 }
// 	 return $TKLBC_Count;
//  }
/***********************************************************************
 * 	$TKLB_Content 資料陣列 (String)
 ***********************************************************************/
func KuaiThreeTwoKindSimpleBet(Content string) []interface{} { //列全部
	var ContentAll []interface{}
	Content1 := strings.Split(Content, ",")

	for i1 := 0; i1 < len(Content1); i1++ {

		ResultTemp := Dup_count(strings.Split(Content1[i1], ""))
		content1Value := ""
		content2Value := ""
		for k, v := range ResultTemp {

			if v == 2 {
				content1Value = k
			} else if v == 1 {
				content2Value = k
			}
		}
		ContentTemp := []string{content1Value, content2Value}
		if !In_array(ContentTemp, ContentAll) {

			ContentAll = append(ContentAll, ContentTemp)
		}
	}
	return ContentAll
}

/*	同号 -> 二同号单式 [END]
 ***********************************************************************/

/*	同号 [END]
 ***********************************************************************/

/*	三连号 [END]
***********************************************************************/

/***********************************************************************
 *	不同号 [START] */

/***********************************************************************
 *	不同号 -> 三不同号标准 [START] */

/***********************************************************************
 * 	$HDBC_Content 資料陣列 (Array)
 ***********************************************************************/
//  public function KuaiThreeThreeDifferBetCount($HDBC_Content)
//  {
// 	 $HDBC_Count = 0;
// 	 if(isset($HDBC_Content))
// 	 {
// 		 $HDBC_Count = count($this->KuaiThreeThreeDifferBet($HDBC_Content));
// 	 }
// 	 return $HDBC_Count;
//  }
/***********************************************************************
 * 	$HDB_Content 資料陣列 (Array)
 ***********************************************************************/
func KuaiThreeThreeDifferBet(Content1 []interface{}) []interface{} { //列全部

	var ContentAll []interface{}
	// if len(Content1) > 0 && len(Content2) > 0 {
	for i1 := 0; i1 < len(Content1); i1++ {
		for i2 := i1 + 1; i2 < len(Content1); i2++ {
			for i3 := i2 + 1; i3 < len(Content1); i3++ {
				ContentAll = append(ContentAll, []string{Content1[i1].(string), Content1[i2].(string), Content1[i3].(string)})
			}
		}
	}
	// }
	return ContentAll
}

/*	不同号 -> 三不同号标准 [END]
 ***********************************************************************/

/***********************************************************************
 *	不同号 -> 三不同号胆拖 [START] */

/***********************************************************************
 * 	$HDDTBC_Content_1 資料陣列 (Array)
 * 	$HDDTBC_Content_2 資料陣列 (Array)
 ***********************************************************************/
//  public function KuaiThreeThreeDifferDanTuBetCount($HDDTBC_Content_1,$HDDTBC_Content_2)
//  {
// 	 $HDDTBC_Count = 0;
// 	 if(isset($HDDTBC_Content_1) && isset($HDDTBC_Content_2))
// 	 {
// 		 if(count($HDDTBC_Content_2) >= 2)
// 		 {
// 			 $HDDTBC_Count = count($this->KuaiThreeThreeDifferDanTuBet($HDDTBC_Content_1,$HDDTBC_Content_2));
// 		 }
// 	 }
// 	 return $HDDTBC_Count;
//  }
/***********************************************************************
  * 	$HDDTB_Content_1 資料陣列 (Array)
  * 	$HDDTB_Content_2 資料陣列 (Array)
//   ***********************************************************************/
//   func KuaiThreeThreeDifferDanTuBet(Content string) []interface{} { //列全部
// 	var ContentAll []interface{}
// 	for i1 := 0; i1 < len(Content1); i1++ {
// 		for i2 := 0; i2 < len(Content1); i2++ {
// 			for i3 :=  i2+1; i3 < len(Content1); i3++ {

// 	 for($i1 = 0; $i1 < count($HDDTB_Content_1); $i1++)
// 	 {
// 		 for($i2 = 0; $i2 < count($HDDTB_Content_2); $i2++)
// 		 {
// 			 for($i3 = $i2 + 1; $i3 < count($HDDTB_Content_2); $i3++)
// 			 {
// 				 if(count(array_unique([$HDDTB_Content_1[$i1],$HDDTB_Content_2[$i2],$HDDTB_Content_2[$i3]]))) //胆码和拖码不可相同
// 				 {
// 					 array_push($HDDTB_ContentAll,[$HDDTB_Content_1[$i1],$HDDTB_Content_2[$i2],$HDDTB_Content_2[$i3]]);
// 				 }
// 			 }
// 		 }
// 	 }
// 	 return $HDDTB_ContentAll;
//  }

/*	不同号 -> 三不同号胆拖 [END]
 ***********************************************************************/

/***********************************************************************
 *	不同号 -> 二不同号标准 [START] */

/***********************************************************************
 * 	$TDBC_Content 資料陣列 (Array)
 ***********************************************************************/
//  public function KuaiThreeTwoDifferBetCount($TDBC_Content)
//  {
// 	 $TDBC_Count = 0;
// 	 if(isset($TDBC_Content))
// 	 {
// 		 $TDBC_Count = count($this->KuaiThreeTwoDifferBet($TDBC_Content));
// 	 }
// 	 return $TDBC_Count;
//  }
/***********************************************************************
 * 	$TDB_Content 資料陣列 (Array)
 ***********************************************************************/
func KuaiThreeTwoDifferBet(Content1 []interface{}) []interface{} { //列全部
	var ContentAll []interface{}
	// if len(Content1) > 0 && len(Content2) > 0 {
	for i1 := 0; i1 < len(Content1); i1++ {
		for i2 := i1 + 1; i2 < len(Content1); i2++ {
			ContentAll = append(ContentAll, []string{Content1[i1].(string), Content1[i2].(string)})
		}
	}
	return ContentAll

}

/*	不同号 -> 二不同号标准 [END]
 ***********************************************************************/

/***********************************************************************
 *	不同号 -> 二不同号胆拖 [START] */

/***********************************************************************
 * 	$TDDTBC_Content_1 資料陣列 (Array)
 * 	$TDDTBC_Content_2 資料陣列 (Array)
 ***********************************************************************/
//  public function KuaiThreeTwoDifferDanTuBetCount($TDDTBC_Content_1,$TDDTBC_Content_2)
//  {
// 	 $TDDTBC_Count = 0;
// 	 if(isset($TDDTBC_Content_1) && isset($TDDTBC_Content_2))
// 	 {
// 		 $TDDTBC_Count = count($this->KuaiThreeTwoDifferDanTuBet($TDDTBC_Content_1,$TDDTBC_Content_2));
// 	 }
// 	 return $TDDTBC_Count;
//  }
/***********************************************************************
 * 	$TDDTB_Content_1 資料陣列 (Array)
 * 	$TDDTB_Content_2 資料陣列 (Array)
 ***********************************************************************/
func KuaiThreeTwoDifferDanTuBet(Content1 []interface{}, Content2 []interface{}) []interface{} { //列全部
	var ContentAll []interface{}
	// if len(Content1) > 0 && len(Content2) > 0 {
	for i1 := 0; i1 < len(Content1); i1++ {
		for i2 := 0; i2 < len(Content2); i2++ {

			if Content1[i1].(string) != Content1[i2].(string) {

				ContentAll = append(ContentAll, []string{Content1[i1].(string), Content1[i2].(string)})
			}

		}
	}
	return ContentAll

}

/*	不同号 -> 二不同号胆拖 [END]
 ***********************************************************************/

/***********************************************************************
 *	不同号 -> 三不同号单式 [START] */

/***********************************************************************
 * 	$HDLBC_Content 資料陣列 (Array)
 ***********************************************************************/
//  func KuaiThreeThreeDifferSimpleBetCount($HDLBC_Content)
//  {
// 	 $HDLBC_Count = 0;
// 	 if(isset($HDLBC_Content))
// 	 {
// 		 $HDLBC_Count = count($this->KuaiThreeThreeDifferSimpleBet($HDLBC_Content));
// 	 }
// 	 return $HDLBC_Count;
//  }
/***********************************************************************
 * 	$HDLB_Content 資料陣列 (Array)
 ***********************************************************************/
func KuaiThreeThreeDifferSimpleBet(Content string) []interface{} { //列全部

	var ContentAll []interface{}
	Content1 := strings.Split(Content, ",")

	for i1 := 0; i1 < len(Content1); i1++ {
		Content2 := strings.Split(Content1[i1], "")
		ContentTemp := []string{Content2[0], Content2[1], Content2[2]}

		sort.Strings(ContentTemp)
		if !In_array(ContentTemp, ContentAll) {
			ContentAll = append(ContentAll, ContentTemp)
		}
	}
	return ContentAll
}

/*	不同号 -> 三不同号单式 [END]
 ***********************************************************************/
/***********************************************************************
 *	不同号 -> 二不同号单式 [START] */

/***********************************************************************
 * 	$TDLBC_Content 資料陣列 (Array)
 ***********************************************************************/
//  public function KuaiThreeTwoDifferSimpleBetCount($TDLBC_Content)
//  {
// 	 $TDLBC_Count = 0;
// 	 if(isset($TDLBC_Content))
// 	 {
// 		 $TDLBC_Count = count($this->KuaiThreeTwoDifferSimpleBet($TDLBC_Content));
// 	 }
// 	 return $TDLBC_Count;
//  }
/***********************************************************************
 * 	$TDLB_Content 資料陣列 (Array)
 ***********************************************************************/
func KuaiThreeTwoDifferSimpleBet(Content string) []interface{} { //列全部

	var ContentAll []interface{}
	Content1 := strings.Split(Content, ",")

	for i1 := 0; i1 < len(Content1); i1++ {
		Content2 := strings.Split(Content1[i1], "")
		//建立要儲存的組合為陣列
		ContentTemp := []string{Content2[0], Content2[1]}
		//排序要儲存的組合
		sort.Strings(ContentTemp)

		//如果要儲存的組合符合規則、且沒有在要回傳的列表中，就加進去…
		if !In_array(ContentTemp, ContentAll) {
			ContentAll = append(ContentAll, ContentTemp)
		}
	}
	return ContentAll
}

/*	不同号 -> 二不同号单式 [END]
 ***********************************************************************/

/*	不同号 [END]
 ***********************************************************************/

/***********************************************************************
 *	连码 -> 三全中 [START] */

/***********************************************************************
 * 	$SHGABC_Content 資料陣列 (Array)
 ***********************************************************************/
//  func MarkSixStraightThreeGetAllBetCount(Content) {
// 	 $SHGABC_Count = 0;
// 	 if(isset($SHGABC_Content))
// 	 {
// 		 if (count($SHGABC_Content) >= 3)
// 		 {
// 			 $SHGABC_Count = count($this->MarkSixStraightThreeGetAllBet($SHGABC_Content));
// 		 }
// 	 }
// 	 return $SHGABC_Count;
//  }
/***********************************************************************
 * 	$SHGAB_Content 資料陣列 (Array)
 *
 *	共用：	『 连码 -> 三中二 』
 *			『 连尾 -> 三肖连中 / 三肖连不中 』
 ***********************************************************************/
func MarkSixStraightThreeGetAllBet(Content []interface{}) []interface{} { //列全部

	var ContentAll []interface{}
	for i := 0; i < len(Content); i++ {
		for j := i + 1; j < len(Content); j++ {
			for k := j + 1; k < len(Content); k++ {
				ContentAll = append(ContentAll, []string{Content[i].(string), Content[j].(string), Content[k].(string)})
			}
		}
	}
	return ContentAll
}

/*	连码 -> 三全中 [END]
 ***********************************************************************/
/***********************************************************************
 *	连码 -> 三中二 [START] */

/***********************************************************************
 * 	$STGABC_Content 資料陣列 (Array)
 ***********************************************************************/
func MarkSixStraightThreeGetTwoBetCount(Content []interface{}) int {
	Count := 0

	if len(Content) >= 3 {
		Count = len(MarkSixStraightThreeGetTwoBet(Content))
	}

	return Count
}

/***********************************************************************
 * 	列全部的部份，同『连码 -> 三全中』
 ***********************************************************************/
func MarkSixStraightThreeGetTwoBet(Content []interface{}) []interface{} {
	return MarkSixStraightThreeGetAllBet(Content)
}

/*	连码 -> 三中二 [END]
 ***********************************************************************/

/***********************************************************************
 *	连码 -> 二全中 [START] */

/***********************************************************************
 * 	$STGABC_Content 資料陣列 (Array)
 ***********************************************************************/
func MarkSixStraightTwoGetAllBetCount(Content []interface{}) int {
	Count := 0

	if len(Content) >= 2 {
		Count = len(MarkSixStraightTwoGetAllBet(Content))
	}

	return Count
}

/***********************************************************************
 * 	$STGAB_Content 資料陣列 (Array)
 *
 *	共用：	『 连码 -> 二中特 』
*			『 连码 -> 特串 』
*			『 连尾 -> 二尾连中 / 二尾连不中 』
***********************************************************************/
func MarkSixStraightTwoGetAllBet(Content []interface{}) []interface{} { //列全部

	var ContentAll []interface{}
	for i := 0; i < len(Content); i++ {
		for j := i + 1; j < len(Content); j++ {
			ContentAll = append(ContentAll, []string{Content[i].(string), Content[j].(string)})

		}
	}
	return ContentAll
}

/*	连码 -> 二全中 [END]
 ***********************************************************************/
/***********************************************************************
 *	连码 -> 二中特 [START] */

/***********************************************************************
 * 	$STGTBC_Content 資料陣列 (Array)
 ***********************************************************************/
func MarkSixStraightTwoGetUniqueBetCount(Content []interface{}) int {
	Count := 0
	if len(Content) >= 2 {
		Count = len(MarkSixStraightTwoGetUniqueBet(Content))
	}

	return Count
}

/***********************************************************************
 * 	列全部的部份，同『连码 -> 二全中』
 ***********************************************************************/
func MarkSixStraightTwoGetUniqueBet(Content []interface{}) []interface{} {
	return MarkSixStraightTwoGetAllBet(Content)
}

/*	连码 -> 二中特 [END]
 ***********************************************************************/

/***********************************************************************
 *	连码 -> 特串 [START] */

/***********************************************************************
 * 	$SUTBC_Content 資料陣列 (Array)
 ***********************************************************************/
func MarkSixStraightUniqueThreadBetCount(Content []interface{}) int {
	Count := 0

	if len(Content) >= 2 {
		Count = len(MarkSixStraightUniqueThreadBet(Content))
	}

	return Count
}

/***********************************************************************
 * 	列全部的部份，同『连码 -> 二全中』
 ***********************************************************************/
func MarkSixStraightUniqueThreadBet(Content []interface{}) []interface{} {

	return MarkSixStraightTwoGetAllBet(Content)
}

/*	连码 -> 特串 [END]
 ***********************************************************************/
/***********************************************************************
 *	连码 -> 四全中 [START] */

/***********************************************************************
 * 	$SFGABC_Content 資料陣列 (Array)
 ***********************************************************************/
func MarkSixStraightFourGetAllBetCount(Content []interface{}) int {

	Count := 0
	if len(Content) >= 4 {
		Count = len(MarkSixStraightFourGetAllBet(Content))
	}

	return Count
}

/***********************************************************************
 * 	$SFGAB_Content 資料陣列 (Array)
 *
 *	共用：	『 连码 -> 四肖连中 / 四肖连不中  』
 *			『 连尾 -> 四尾连中 / 四尾连不中 』
 ***********************************************************************/
func MarkSixStraightFourGetAllBet(Content []interface{}) []interface{} { //列全部
	var ContentAll []interface{}
	for i1 := 0; i1 < len(Content); i1++ {
		for i2 := i1 + 1; i2 < len(Content); i2++ {
			for i3 := i2 + 1; i3 < len(Content); i3++ {
				for i4 := i3 + 1; i4 < len(Content); i4++ {
					ContentAll = append(ContentAll, []string{Content[i1].(string), Content[i2].(string), Content[i3].(string), Content[i4].(string)})
				}
			}
			// C

		}
	}

	return ContentAll
}

/*	连码 -> 四全中 [END]
 ***********************************************************************/
/*	连码 [END]
 ***********************************************************************/

/***********************************************************************
 *	连肖 [START] */

/***********************************************************************
 *	连肖 -> 二肖连中 / 二肖连不中 [START] */

/***********************************************************************
 * 	$STTBC_Content 資料陣列 (Array)
 ***********************************************************************/
func MarkSixStraightTwoZodiacBetCount(Content []interface{}) int {
	Count := 0

	if len(Content) >= 2 {
		Count = len(MarkSixStraightTwoZodiacBet(Content))
	}

	return Count
}

/***********************************************************************
 * 	列全部的部份，同『连码 -> 二全中』
 ***********************************************************************/
func MarkSixStraightTwoZodiacBet(Content []interface{}) []interface{} {
	return MarkSixStraightTwoGetAllBet(Content)
}

/*	连肖 -> 二肖连中 / 二肖连不中 [END]
 ***********************************************************************/

/***********************************************************************
 *	连肖 -> 三肖连中 / 三肖连不中 [START] */

/***********************************************************************
 * 	$STHBC_Content 資料陣列 (Array)
 ***********************************************************************/
func MarkSixStraightThreeZodiacBetCount(Content []interface{}) int {
	Count := 0
	if len(Content) >= 3 {
		Count = len(MarkSixStraightThreeZodiacBet(Content))
	}
	return Count
}

/***********************************************************************
 * 	列全部的部份，同『连码 -> 三全中』
 ***********************************************************************/
func MarkSixStraightThreeZodiacBet(Content []interface{}) []interface{} {
	return MarkSixStraightThreeGetAllBet(Content)
}

/*	连肖 -> 三肖连中 / 三肖连不中 [END]
 ***********************************************************************/
func Array_unique(OpenResultTemp []int) []int {
	resultTemp := make([]int, 0, len(OpenResultTemp))
	temp := map[int]struct{}{}
	for _, item := range OpenResultTemp {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			resultTemp = append(resultTemp, item)
		}
	}
	return resultTemp
}
func Array_unique_str(OpenResultTemp []string) []string {
	resultTemp := make([]string, 0, len(OpenResultTemp))
	temp := map[string]struct{}{}
	for _, item := range OpenResultTemp {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			resultTemp = append(resultTemp, item)
		}
	}
	return resultTemp
}

func MinIntSlice(v []int) int {
	sort.Ints(v)
	return v[0]
}

func MaxIntSlice(v []int) int {
	sort.Ints(v)
	return v[len(v)-1]
}

func Dup_count(list []string) map[string]int {

	duplicate_frequency := make(map[string]int)

	for _, item := range list {
		// check if the item/element exist in the duplicate_frequency map

		_, exist := duplicate_frequency[item]

		if exist {
			duplicate_frequency[item] += 1 // increase counter by 1 if already in the map
		} else {
			duplicate_frequency[item] = 1 // else start counting from 1
		}
	}
	return duplicate_frequency
}
func In_array2(values interface{}, array interface{}) bool {
	exists := false
	switch reflect.TypeOf(values).Kind() {
	case reflect.Slice:
		k := reflect.ValueOf(values)

		switch reflect.TypeOf(array).Kind() {

		case reflect.Slice:
			s := reflect.ValueOf(array)
			for i := 0; i < s.Len(); i++ {

				if reflect.DeepEqual(k.Interface(), s.Index(i).Interface()) == true {
					exists = true
					return exists
				}
			}
		}

	case reflect.String:
		switch reflect.TypeOf(array).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(array)
			for i := 0; i < s.Len(); i++ {
				if reflect.DeepEqual(values, s.Index(i).Interface()) == true {
					exists = true
					return exists
				}
			}
		}
	}
	return exists

}
func In_array(values interface{}, array interface{}) bool {
	exists := false

	switch reflect.TypeOf(values).Kind() {
	case reflect.Slice:
		k := reflect.ValueOf(values)
		for j := 0; j < k.Len(); j++ {

			switch reflect.TypeOf(array).Kind() {
			case reflect.Slice:
				s := reflect.ValueOf(array)
				for i := 0; i < s.Len(); i++ {
					if reflect.DeepEqual(k.Index(j).Interface(), s.Index(i).Interface()) == true {

						exists = true
						return exists
					}
				}
			}
		}

	case reflect.String:
		switch reflect.TypeOf(array).Kind() {
		case reflect.Slice:

			s := reflect.ValueOf(array)

			for i := 0; i < s.Len(); i++ {
				if reflect.DeepEqual(values, s.Index(i).Interface()) == true {
					exists = true
					return exists
				}
			}
		}
	case reflect.Int:
		switch reflect.TypeOf(array).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(array)
			for i := 0; i < s.Len(); i++ {
				if reflect.DeepEqual(values, s.Index(i).Interface()) == true {
					exists = true
					return exists
				}
			}
		}
	case reflect.Float64:
		switch reflect.TypeOf(array).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(array)
			for i := 0; i < s.Len(); i++ {
				if reflect.DeepEqual(values, s.Index(i).Interface()) == true {
					exists = true
					return exists
				}
			}
		}
	}
	return exists
}

type Pair struct {
	Key   string
	Value int
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

// func flipMapByValue(m map[string]int) v map[string]int{

// 	for k, v := range m {
// 		p[i] = Pair{k, v}
// 		i++
// 	}

// 	// p := make(PairList, len(m))
// 	// i := 0
// 	// for k, v := range m {
// 	// 	p[i] = Pair{k, v}
// 	// 	i++
// 	// }
// 	// sort.Sort(p)
// 	return v
// }

func SortMapByValue(m map[string]int) PairList {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)
	return p
}

func ArraySum(input interface{}) int {
	sum := 0
	for _, i := range input.([]string) {
		num, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		sum += num
	}
	return sum
}
