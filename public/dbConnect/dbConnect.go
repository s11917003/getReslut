package dbConnect

import (
	"database/sql"
	"fmt"
	"getReslut/config"
	"getReslut/public"
	"log"
	"strconv"
	"strings"

	_ "github.com/mysql-master"
)

var db *sql.DB

// var connStr = "mem:1qaz!QAZ@tcp(192.168.18.3:3306)/MEM?charset=utf8"
var ConfigData map[string]interface{}

//var dbConfigData map[string]interface{}

// func getDBConfig() string {
// 	dbConfigData := make(map[string]interface{})

// 	dbConfigData["ip"] = "192.168.18.3"
// 	dbConfigData["poolname"] = "mysqlMEM"
// 	dbConfigData["port"] = "3306"
// 	dbConfigData["user"] = "mem"
// 	dbConfigData["password"] = "1qaz!QAZ"
// 	dbConfigData["database"] = "MEM"
// 	dbConfigData["charset"] = "utf8"

// 	connStr = dbConfigData["user"].(string) + ":" + dbConfigData["password"].(string) +
// 		"@tcp(" + dbConfigData["ip"].(string) + ":" + dbConfigData["port"].(string) + ")/" +
// 		dbConfigData["database"].(string) + "?charset=" + dbConfigData["charset"].(string)
// 	return connStr
// }
func SetLotteryDrawChainCode(thisLotteryType string, thisLotteryIssue string, chainCode string, thisOpenResult []string, datedrawTime int64) bool {

	dbCongig := config.GetDBConfig()
	sqlResult := make(map[string]interface{})
	sqlResult["result"] = ""
	sqlResult["count"] = 0
	sqlResult["state"] = 0

	// thisOpenResult1 := strings.Split(thisOpenResult, ";")[0]
	// thisOpenResult2 := strings.Split(thisOpenResult, ";")[1]

	//1. 连接数据库
	db, err := sql.Open("mysql", dbCongig)
	if err != nil {
		log.Fatal("Failed to open database:", err)
		sqlResult["result"] = err
		sqlResult["count"] = 0
		sqlResult["state"] = 0
		return false
	}
	sql := ""

	thisLotteryTypeInt, err := strconv.Atoi(thisLotteryType)
	if public.Contains(thisLotteryTypeInt, config.GetgameCodeMap()["pk10"].(map[string]interface{})["lt"].([]int)) {

		var thisOpenResultInt = []int{}

		for _, i := range thisOpenResult {
			j, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}
			thisOpenResultInt = append(thisOpenResultInt, j)

		}
		sql = fmt.Sprintf(
			"INSERT INTO TS_LotteryDrawChainCode (LDC_LotteryType,LDC_Issue,LDC_LotteryDrawChainCode,LDC_OpenResult,LDC_CreateTime,LDC_DataDrawTime) value (%s,%s,'%s','[%s]',UNIX_TIMESTAMP(),%d)",
			thisLotteryType,
			thisLotteryIssue,
			chainCode,
			strings.Trim(strings.Replace(fmt.Sprint(thisOpenResultInt), " ", ",", -1), "[]"),
			datedrawTime,
		)
	} else {
		sql = fmt.Sprintf(
			"INSERT INTO TS_LotteryDrawChainCode (LDC_LotteryType,LDC_Issue,LDC_LotteryDrawChainCode,LDC_OpenResult,LDC_CreateTime,LDC_DataDrawTime) value (%s,%s,'%s','[%s]',UNIX_TIMESTAMP(),%d)",
			thisLotteryType,
			thisLotteryIssue,
			chainCode,
			strings.Join(thisOpenResult, ","),
			datedrawTime,
		)
	}

	fmt.Println("SetDrawChainCode sql ", sql)

	res, err := db.Exec(sql)
	if err != nil {
		log.Println("exec failed:", err, ", sql:", sql)
		return false
	}

	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	fmt.Println("SetRtpRecord ", num)
	return true
}
func SetRtpRecord(amount float64, bonus float64, thisLotteryTypeGroup string, thisLotteryType string, thisRatio float64, now string, thisLotteryIssue string, rtpNow float64) bool {

	dbCongig := config.GetDBConfig()
	sqlResult := make(map[string]interface{})
	sqlResult["result"] = ""
	sqlResult["count"] = 0
	sqlResult["state"] = 0
	//1. 连接数据库
	db, err := sql.Open("mysql", dbCongig)
	if err != nil {
		log.Fatal("Failed to open database:", err)
		sqlResult["result"] = err
		sqlResult["count"] = 0
		sqlResult["state"] = 0
		return false
	}

	sql := fmt.Sprintf(
		"INSERT INTO M_RtpRecordDetail (LTG_Code,Gid,Issue,RtpSetting,RTPNow, Amount,Bonus,AddDate) value (%s,%s,%s,%f,%f,%f,%f,'%s')",
		thisLotteryTypeGroup,
		thisLotteryType,
		thisLotteryIssue,
		thisRatio,
		rtpNow,
		amount,
		bonus,
		now,
	)

	fmt.Println("SetRtpRecord sql ", sql)
	res, err := db.Exec(sql)
	if err != nil {
		log.Println("exec failed:", err, ", sql:", sql)
		return false
	}

	fmt.Println("UpdateAmount   res", sql)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	fmt.Println("SetRtpRecord ", num)
	return true
}

func SetAmount(amount float64, bonus float64, thisLotteryTypeGroup string, thisLotteryType string, thisRatio float64, now string) bool {

	dbCongig := config.GetDBConfig()
	sqlResult := make(map[string]interface{})
	sqlResult["result"] = ""
	sqlResult["count"] = 0
	sqlResult["state"] = 0
	//1. 连接数据库
	db, err := sql.Open("mysql", dbCongig)
	if err != nil {
		log.Fatal("Failed to open database:", err)
		sqlResult["result"] = err
		sqlResult["count"] = 0
		sqlResult["state"] = 0
		return false
	}

	currec := 0.0
	if bonus == 0 {
		currec = -(1 - thisRatio)
	} else {
		currec = thisRatio - bonus/amount

	}
	sql := fmt.Sprintf(
		"UPDATE M_RtpSetting2 SET Exprec=1-RTP,Currec=%f, Amount=%f,Bonus=%f,Issue=Issue+1,DomputeDate='%s' WHERE LTG_Code=%s and Gid=%s and RTP=%f",
		currec,
		amount,
		bonus,
		now,
		thisLotteryTypeGroup,
		thisLotteryType,
		thisRatio,
	)

	fmt.Println("setAmount sql ", sql)

	res, err := db.Exec(sql)
	if err != nil {
		log.Println("exec failed:", err, ", sql:", sql)
		return false
	}

	// fmt.Println("UpdateAmount   res", sql)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	fmt.Println("UpdateAmount", num)

	return true
}

/**
**/

func UpdateAmount(amount float64, bonus float64, thisLotteryTypeGroup string, thisLotteryType string, thisRatio float64, now string) bool {

	dbCongig := config.GetDBConfig()
	sqlResult := make(map[string]interface{})
	sqlResult["result"] = ""
	sqlResult["count"] = 0
	sqlResult["state"] = 0
	//1. 连接数据库
	db, err := sql.Open("mysql", dbCongig)
	if err != nil {
		log.Fatal("Failed to open database:", err)
		sqlResult["result"] = err
		sqlResult["count"] = 0
		sqlResult["state"] = 0
		return false
	}

	sql := fmt.Sprintf(
		"UPDATE M_RtpSetting2 SET Exprec=1-RTP,Currec=1-(%f/%f),Amount=%f,Bonus=%f,Issue=Issue+1,DomputeDate='%s' WHERE LTG_Code=%s and Gid=%s and State = 1 AND Type = 1 ",
		bonus,
		amount,
		amount,
		bonus,
		now,
		thisLotteryTypeGroup,
		thisLotteryType,
	)

	fmt.Println(sql)

	res, err := db.Exec(sql)
	if err != nil {
		log.Println("exec failed:", err, ", sql:", sql)
		return false
	}

	// fmt.Println("UpdateAmount   res", sql)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	fmt.Println("UpdateAmount", num)

	return true
}
func GetLotteryDrawTime(thisLotteryType string, thisLotteryIssue string) map[string]interface{} {
	dbCongig := config.GetDBConfig()

	sqlResult := make(map[string]interface{})
	sqlResult["datedraw"] = "0"

	SqlString := "SELECT LD_LotteryType,LD_Issue,LD_DateDraw " +
		" FROM TS_LotteryDraw  WHERE  " +
		" LD_LotteryType = " + thisLotteryType +
		" AND LD_Issue = " + thisLotteryIssue
	fmt.Println("SqlString ", SqlString)

	//1. 连接数据库
	db, err := sql.Open("mysql", dbCongig)

	if err != nil {
		fmt.Println("Failed to open database:")
		log.Fatal("Failed to open database:", err)
		sqlResult["datedraw"] = "0"
	}
	rows, err := db.Query(SqlString)

	checkErr(err)

	columns, _ := rows.Columns()

	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))

	resultArr := make([]interface{}, 0)
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		//將行資料儲存到record字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
				//fmt.Println("columns  ", columns[i])
			}
		}
		resultArr = append(resultArr, record)
	}
	db.Close()

	if len(resultArr) == 1 {
		sqlResult["datedraw"] = resultArr[0].(map[string]string)["LD_DateDraw"]
	}
	fmt.Println("sqlResult  ", sqlResult)
	return sqlResult
}
func GetRtpSetting(thisLotteryTypeGroup string, thisLotteryType string) map[string]interface{} {
	dbCongig := config.GetDBConfig()

	sqlResult := make(map[string]interface{})
	sqlResult["result"] = ""
	sqlResult["count"] = 0
	sqlResult["state"] = 0
	SqlString := "SELECT RTP,Amount,Bonus " +
		" FROM M_RtpSetting2  WHERE Sid = 1   AND State = 1 AND Type = 1  " +
		" AND LTG_Code = " + thisLotteryTypeGroup +
		" AND Gid = " + thisLotteryType
	fmt.Println("SqlString ", SqlString)

	//1. 连接数据库
	db, err := sql.Open("mysql", dbCongig)

	if err != nil {
		fmt.Println("Failed to open database:")
		log.Fatal("Failed to open database:", err)
		sqlResult["result"] = err
		sqlResult["count"] = 0
		sqlResult["state"] = 0
	}
	rows, err := db.Query(SqlString)

	checkErr(err)

	columns, _ := rows.Columns()

	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))

	resultArr := make([]interface{}, 0)
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		//將行資料儲存到record字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
				//fmt.Println("columns  ", columns[i])
			}
		}
		resultArr = append(resultArr, record)
	}
	db.Close()

	sqlResult["count"] = len(resultArr)
	if len(resultArr) == 1 {
		sqlResult["result"] = resultArr[0].(map[string]string)["RTP"]
		amount, _ := strconv.ParseFloat(resultArr[0].(map[string]string)["Amount"], 64)
		bonus, _ := strconv.ParseFloat(resultArr[0].(map[string]string)["Bonus"], 64)
		sqlResult["amount"] = amount
		sqlResult["bonus"] = bonus
		sqlResult["state"] = 1
	}
	fmt.Println("sqlResult  ", sqlResult)
	return sqlResult
}

func Run(thisLotteryTypeGroup string, thisLotteryType string, thisLotteryIssue string) map[string]interface{} {
	return SelectBetsList(thisLotteryTypeGroup, thisLotteryType, thisLotteryIssue, "0", "0")
}

func SelectBetsList(LotteryTypeGroup string, LotteryType string, Issue string, Collect string, QueryLimitint string) map[string]interface{} {

	dbCongig := config.GetDBConfig()

	sqlResult := make(map[string]interface{})
	sqlResult["result"] = ""
	sqlResult["count"] = 0
	sqlResult["state"] = 0

	SqlString := "SELECT BO_No, BO_OrderNo , BO_User , BO_Issue " +
		" , BO_LotteryTypeGroup, BO_LotteryType, BO_Mode, BO_LotteryPlayGroup, BO_LotteryPlay , BO_LotteryContent" +
		", BO_Price, BO_RealPrice, BO_Unit, BO_Multiple , BO_BetCount " +
		" , BO_Winnings, BO_Winnings2, BO_Winnings3, BO_WinningLimit " +
		" , BO_Water , BO_BW_SourceUserId, BO_BW_BackWater, BO_BW_LimitBetCount, BO_BW_MaxAmount, BO_BW_AuditMultiple " +
		" , BO_Odds, BO_Odds2, BO_Odds3 , BO_Type, BO_Chase , BO_OpenResult, BO_OpenResult2, BO_DateDraw, BO_DateDraw2, BO_DateUpdate " +
		" FROM TS_BetOrder  WHERE BO_Cancel = 0   AND BO_CancelBackend = 0  " +
		" AND BO_LotteryTypeGroup = " + LotteryTypeGroup +
		//" AND BO_LotteryPlayGroup = 11 " +
		// " AND BO_LotteryPlay = 1" +
		// " AND BO_Mode = 1 " +
		" AND BO_LotteryType = " + LotteryType + //到時候 要拿掉註解
		" AND BO_Issue = " + Issue +
		" AND BO_Collect = " + Collect + //到時候 要拿掉註解
		" ORDER BY BO_DateDraw ASC, BO_DateCreate ASC, BO_DateUpdate ASC, BO_No ASC "

	if QueryLimitint == "0" {
		SqlString += ";"
	} else {
		SqlString += " LIMIT $SBL_QueryLimit ; "

	}
	fmt.Println(SqlString)

	//1. 连接数据库
	db, err := sql.Open("mysql", dbCongig)
	if err != nil {
		log.Fatal("Failed to open database:", err)
		sqlResult["result"] = err
		sqlResult["count"] = 0
		sqlResult["state"] = 0
	}
	rows, err := db.Query(SqlString)
	checkErr(err)

	columns, _ := rows.Columns()

	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))

	resultArr := make([]interface{}, 0)
	for i := range values {
		scanArgs[i] = &values[i]

	}

	for rows.Next() {
		//將行資料儲存到record字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
				// fmt.Println("columns  ", columns[i])
			}
		}

		resultArr = append(resultArr, record)
	}
	db.Close()
	sqlResult["result"] = resultArr
	sqlResult["count"] = len(resultArr)
	sqlResult["state"] = 1

	return sqlResult
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
