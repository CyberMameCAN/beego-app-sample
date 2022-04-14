package controllers

import (
	"app/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/astaxie/beego"
)

// Controllerを登録
type BbsController struct {
	beego.Controller
}

type BbsLog models.BbsLog

type BbsLogs []BbsLog

var bbslogs []BbsLog

const logFile = "logs.json" // データの保存先

func init() {
}

// 書き込みログを画面に表示する
func (this *BbsController) Show() {
	bbslogs := loadLogs() // 既存のデータを読み出し
	this.Data["logs"] = bbslogs
	this.TplName = "bbs.tpl.html"
}

// フォームから送信された内容を書き込み
func (this *BbsController) Post() {
	var bbslog BbsLog

	// 名前の取得
	name := this.GetStrings("name")
	if name != nil {
		bbslog.Name = name[0]
	} else {
		msg := fmt.Sprintln("Name is not set.")
		beego.Debug(msg)
		bbslog.Name = "名前無し"
	}
	// 本文の取得
	body := this.GetStrings("body")
	if body != nil {
		bbslog.Body = body[0]
	} else {
		msg := fmt.Sprintln("Body is not set.")
		beego.Debug(msg)
		bbslog.Body = "本文無し"
	}
	bbslog.CTime = time.Now().Unix()

	bbslogs := loadLogs() // 既存のデータを読み出し
	bbslog.ID = len(bbslogs) + 1
	bbslog.CTime = time.Now().Unix()
	bbslogs = append(bbslogs, bbslog) // 追記 --- (*12)
	saveLogs(bbslogs)                 // 保存
	this.Ctx.Redirect(302, "/bbs")
}

// ファイルからログファイルの読み込み --- (*15)
func loadLogs() []BbsLog {
	// ファイルを開く
	text, err := ioutil.ReadFile(logFile)
	if err != nil {
		return make([]BbsLog, 0)
	}
	// JSONをパース --- (*16)
	var bbslogs []BbsLog
	json.Unmarshal([]byte(text), &bbslogs)
	return bbslogs
}

// ログファイルの書き込み --- (*17)
func saveLogs(bbslogs []BbsLog) {
	// JSONにエンコード
	bytes, _ := json.Marshal(bbslogs)
	// ファイルへ書き込む
	ioutil.WriteFile(logFile, bytes, 0644)
}
