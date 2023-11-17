package main

import (
	"botline/Library/linethrift"
	"botline/Library/oop"
	"encoding/json"
	"fmt"
	"io/ioutil"
	// "math/rand"
	"os"
	"runtime"
	// "strconv"
	"database/sql"
	// "github.com/kardianos/osext"
	_ "github.com/mattn/go-sqlite3"
	"strings"
	"sync"
	// "syscall"
	"time"
)

type Bots struct {
	bots []string `json:"bots"`
}
type User struct {
	Owner               []string             `json:"owner"`
	Admin               []string             `json:"admin"`
	God                 []string             `json:"god"`
	Staff               []string             `json:"staff"`
	Osquad              []string             `json:"osquad"`
	Ban                 []string             `json:"ban"`
	TargetSpam          []string             `json:"targetspam"`
	LimitStatus         map[string]bool      `json:"limitstatus"`
	LimitTime           map[string]time.Time `json:"limittime"`
	ProKick             map[string]bool      `json:"prokick"`
	ProInvite           map[string]bool      `json:"proinvite"`
	ProCancel           map[string]bool      `json:"procancel"`
	ProJoin             map[string]bool      `json:"projoin"`
	ProQr               map[string]bool      `json:"proqr"`
	ProLINK             map[string]bool      `json:"ProLINK"`
	ProFLEX             map[string]bool      `json:"ProFLEX"`
	ProDelAlbum         map[string]bool      `json:"ProDelAlbum"`
	ProSTICKER          map[string]bool      `json:"ProSTICKER"`
	ProCALL             map[string]bool      `json:"ProCALL"`
	ProFILE             map[string]bool      `json:"ProFILE"`
	ProPOSTNOTIFICATION map[string]bool      `json:"ProPOSTNOTIFICATION"`
	ProVIDEO            map[string]bool      `json:"ProVIDEO"`
	ProAUDIO            map[string]bool      `json:"ProAUDIO"`
	ProIMAGE            map[string]bool      `json:"ProIMAGE"`
	StayGroup           map[string][]string  `json:"staygroup"`
	StayPending         map[string][]string  `json:"staypending"`
}

type changeVideo struct {
	Tipe        string
	Mid         map[string]bool
	PictPath    string
	VideoPath   string
	PictStatus  bool
	VideoStatus bool
}

var arr_data = []string{}
var GroupMemberList = []string{}
var duedatecount int = 0
var (
	data      User
	dataPath  = fmt.Sprintf("data.json")
	dataPathB = fmt.Sprintf("bots.json")
	Squads    = []string{}
	Maker     = []string{
		"u53ab6fa03c2838678a07a10fd142eb81",
	}
	Tokens    = []string{}
	banlist   = []string{}
	Freeze    = []string{}
	KillMod   = true
	MkGroup   = true
	GroupList = []string{}
	Botlist   []*oop.Account
	// sqliteDatabase   *sql.DB
	WarTime          = make(map[string]time.Time)
	TimeJoin         = make(map[string]time.Time)
	msgTemp          = make(map[string][]string)
	TempJoin         = "0"
	TimeSet          = make(map[string]time.Time)
	pala             = ""
	palaMsg          = ""
	PalaSpam         = ""
	NameGroupSpam    = "test spam"
	BotKor           = map[string]bool{}
	lock             = ""
	pTime            = map[int64]bool{}
	cTime            = map[int64]bool{}
	aTime            = map[int64]bool{}
	ModTicket        = ""
	lockMod          = ""
	Token            = ""
	Multy            = true
	Loop             = true
	Scont            = false
	AutoClearban     = false
	AutoMulty        = true
	PromoteSpam      = false
	PromoteBlacklist = false
	delBlacklist     = false
	PromoteStaff     = false
	PromoteAdmin     = false
	PromoteOwner     = false
	DemoteStaff      = false
	DemoteAdmin      = false
	DemoteOwner      = false
	notiFadd         = true
	kickban          = true
	Spamlimit        int
	LimiterJoin      int
	LimiterKick      int
	timeStart        = time.Now()
	CUnsend          string
	updateImage      = map[string]bool{}
	updateCover      = map[string]bool{}
	updateVideo      = &changeVideo{
		Tipe:        "",
		Mid:         map[string]bool{},
		PictPath:    "",
		VideoPath:   "",
		PictStatus:  false,
		VideoStatus: false,
	}
)

func main() {
	_, err := ioutil.ReadFile("sqlite-database.db")
	if err != nil {
		file, err := os.Create("sqlite-database.db") // Create SQLite file
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		file.Close()
	}
	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	defer sqliteDatabase.Close()                                     // Defer Closing the database
	if err != nil {
		createTableBot(sqliteDatabase) // Create Database Tables bot
		createTableBan(sqliteDatabase)
		// INSERT RECORDS
		insertBot(sqliteDatabase, "uca7eaad872ecf76e1af7da4431c72fa6", "uca7eaad872ecf76e1af7da4431c72fa6:aWF0OiAxNjk5OTUyMjAxMTkwCg==..rkqew1/j7HoohEsbn5BIHMaj5lw=")
		insertBot(sqliteDatabase, "u4c2db07fea7f4c775ee1ad393a3cb84c", "u4c2db07fea7f4c775ee1ad393a3cb84c:aWF0OiAxNjk5ODkxMTc5ODQyCg==..On+aFAYH3RCR7PCjRsko6t6/RR4=")

	}
	createTableStatus(sqliteDatabase)
	insertStatus(sqliteDatabase,"msgban",true)
	insertStatus(sqliteDatabase,"kickban",true)
	insertStatus(sqliteDatabase,"ProInvite",true)
	insertStatus(sqliteDatabase,"ProQr",true)
	insertStatus(sqliteDatabase,"ProDelAlbum",true)
	insertStatus(sqliteDatabase,"ProFLEX",true)
	insertStatus(sqliteDatabase,"ProCALL",true)
	// insertStatus(sqliteDatabase,"kickban",true)

	






	// DISPLAY INSERTED RECORDS
	get_bots(sqliteDatabase)
	fmt.Println("Squads", Squads)
	fmt.Println("Tokens", Tokens)
	// fmt.Println("sqlite-database.db created")
	// fileName := fmt.Sprintf("%v.txt", os.Args[1])
	// fileBytes, err := ioutil.ReadFile(fileName)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// Token := strings.Split(string(fileBytes), "\r\n")

	fmt.Println("Token", Tokens[0])
	auth := Tokens[0]
	xcl := oop.Connect(1, auth)
	err_ := xcl.LoginWithAuthToken()
	if err_ == nil {
		fmt.Println("start bot root")
		perBots(xcl)
	} else {
		fmt.Sprintf("%s", err)
		if strings.Contains(fmt.Sprintf("%s", err), "suspended") {
			fmt.Println(auth[:33], "FREEZE OR SUSPEND")
		} else {
			fmt.Println(auth[:33], err)
		}
	}
	// SaveData()

}
func createTableBot(db *sql.DB) {
	createBotTableSQL := `CREATE TABLE bots (
		"id" TEXT NOT NULL PRIMARY KEY,		
		"token" TEXT	
	  );` // SQL Statement for Create Table

	fmt.Println("Create Bot table...")
	statement, err := db.Prepare(createBotTableSQL) // Prepare SQL Statement
	if err != nil {
		fmt.Println(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	fmt.Println("Bot table created")
}
func createTableStatus(db *sql.DB) {
	createBotTableSQL := `CREATE TABLE status (
		"name" TEXT NOT NULL PRIMARY KEY,		
		"status" BOOLEAN	
	  );` // SQL Statement for Create Table

	fmt.Println("Create status table...")
	statement, err := db.Prepare(createBotTableSQL) // Prepare SQL Statement
	if err != nil {
		fmt.Println(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	fmt.Println("status table created")
}
func createTableBan(db *sql.DB) {

	createBanTableSQL := `CREATE TABLE ban ("id" TEXT NOT NULL PRIMARY KEY);` // SQL Statement for Create Table

	fmt.Println("Create Ban table...")
	statement, err := db.Prepare(createBanTableSQL) // Prepare SQL Statement
	if err != nil {
		fmt.Println(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	fmt.Println("Ban table created")
}

// We are passing db reference connection from main to our method with other parameters
func insertBot(db *sql.DB, id string, token string) {
	fmt.Println("Inserting Bot record ...")
	insertBotSQL := `INSERT INTO bots(id, token) VALUES (?, ?)`
	statement, err := db.Prepare(insertBotSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = statement.Exec(id, token)
	if err != nil {
		fmt.Println(err.Error())
	}
}
func insertStatus(db *sql.DB, name string, status bool) {
	fmt.Println("Inserting Bot record ...")
	insertBotSQL := `INSERT INTO status(name, status) VALUES (?, ?)`
	statement, err := db.Prepare(insertBotSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = statement.Exec(name, status)
	if err != nil {
		fmt.Println(err.Error())
	}
}
func get_status(name  string) bool {
	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	defer sqliteDatabase.Close()
	row, err := sqliteDatabase.Query("SELECT * FROM status  WHERE name = ?", name)
	 
	if err != nil {
		fmt.Println(err)
	}
	defer row.Close()

	for row.Next() { // Iterate and fetch the records from result cursor
		var name_ string
		var status_ bool
		row.Scan(&name_, &status_)
		if name_ == name {
			return status_
		}

		// fmt.Println("Tokens",Tokens)
	}
	return false

}
func insertStayGroup(db *sql.DB, id string, target string) {
	fmt.Println("Inserting Bot record ...")
	insertBotSQL := `INSERT INTO StayGroup(id, target) VALUES (?, ?)`
	statement, err := db.Prepare(insertBotSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = statement.Exec(id, target)
	if err != nil {
		fmt.Println(err.Error())
	}
}
func get_bots(db *sql.DB) {
	row, err := db.Query("SELECT * FROM bots ORDER BY id")
	if err != nil {
		fmt.Println(err)
	}
	defer row.Close()

	for row.Next() { // Iterate and fetch the records from result cursor
		var id string
		var token string
		row.Scan(&id, &token)
		Squads = append(Squads, id)
		Tokens = append(Tokens, token)

		// fmt.Println("Tokens",Tokens)
	}

}
func get_ban(cid string) bool {
	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	defer sqliteDatabase.Close()
	row, err := sqliteDatabase.Query("SELECT * FROM ban  WHERE id = ?", cid)
	if err != nil {
		fmt.Println(err)
	}
	defer row.Close()

	var id string
	for row.Next() { // Iterate and fetch the records from result cursor
		row.Scan(&id)
		if cid == id {
			return cid == id
		}
	}
	// return cid == id
	return false
}
func remove_ban(cid string) bool {
	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	defer sqliteDatabase.Close()
	res, err := sqliteDatabase.Exec("DELETE FROM  ban  WHERE id = ?", cid)
	if err != nil {
		fmt.Println(err)
	}

	n, err_ := res.RowsAffected()
	if err != nil {
		fmt.Println(err_)
		return false
	}
	fmt.Printf("The statement has affected %d rows\n", n)
	return true

	// return false
}
func ban_list() []string {
	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	defer sqliteDatabase.Close()
	row, err := sqliteDatabase.Query("SELECT * FROM ban")
	if err != nil {
		fmt.Println(err)
	}
	defer row.Close()
	banlist_data := []string{}
	for row.Next() { // Iterate and fetch the records from result cursor
		var id string
		row.Scan(&id)
		banlist_data = append(banlist_data, id)
	}
	return banlist_data
	// return false
}
func get_StayGroup(db *sql.DB) {
	row, err := db.Query("SELECT * FROM StayGroup WHERE id")
	if err != nil {
		fmt.Println(err)
	}
	defer row.Close()

	for row.Next() { // Iterate and fetch the records from result cursor
		var id string
		var token string
		row.Scan(&id, &token)
		Squads = append(Squads, id)
		Tokens = append(Tokens, token)
	}

}
func perBots(cl *oop.Account) {
	runtime.GOMAXPROCS(2)
	// Botlist = append(Botlist, cl)
	// allgrup, _ := cl.GetAllChatMids(true, false)
	// for g := range allgrup.MemberChatMids {
	// 	putSquad(cl, allgrup.MemberChatMids[g])
	// }
	// if !oop.Contains(Squads, cl.Mid) {
	// 	if !oop.Contains(data.Osquad, cl.Mid) {
	// 		// getBots()
	// 	}
	// }
	// fmt.Println("allgrup", allgrup)
	fmt.Println("data", Squads)
	fmt.Println(" cl.Mid", cl.Mid)
	for {
		ops, _ := cl.FetchOps()

		if len(ops) == 0 {
			continue
		}
		go func(fetch []*linethrift.Operation) {
			for _, op := range fetch {
				if op.Type == 0 {
					cl.CorrectRevision(op, false, true, true)
				} else {
					if op.Type == 133 || op.Type == 124 || op.Type == 123 || op.Type == 122 || op.Type == 126 || op.Type == 26 || op.Type == 5 || op.Type == 130 || op.Type == 132 || op.Type == 55 {
						switch op.Type {
						/*
							OperationType_NOTIFIED_INVITE_INTO_CHAT
							แจ้งเตือนคำเชิญเข้าสู่การแชท
						*/
						case 124:
							fmt.Println("เชิญเข้าสู่การแชท 124")
							// op1, ไอดีห้อง
							// op2, ผู้เชิญ
							// op3 คนถูกเชิญ
							op1, op2, op3, ctime := op.Param1, op.Param2, strings.Split(op.Param3, "\x1e"), op.CreatedTime
							// fmt.Println(op)
							if getAccessForCancel(cl, op2, op3) {
								fmt.Println("getAccessForCancel")
								CclList(cl, op1, op3)
								// go cl.CancelChatInvitation(cl, op1, op2)
							} else if oop.Contains(op3, cl.Mid) {
								fmt.Println("else if 2")

								go func(op1 string) {
									go cl.AcceptChatInvitation(op1)
									// go cl.DeleteOtherFromChat(op1, []string{op2})
									// 	go KickBl(cl, op1)

								}(op1)

							} else if fullAccess(op2) {
								fmt.Println("fullAccess")
								continue
							} else if oop.Contains(data.Ban, op2) || oop.CheckEqual(data.Ban, op3) {
								fmt.Println("else if 3")
								go killAll(cl, op1, op2, op3)
							} else if _, cek := data.ProInvite[op1]; cek {
								fmt.Println("else if 4")
								if getWarAccess(cl, ctime, op1, "", cl.Mid, false) {
									go cl.DeleteOtherFromChat(op1, []string{op2})
									go func() { CclList(cl, op1, op3) }()
									go BanAll(op2, op3)

								}
							} else if kickban == true {
								fmt.Println("kickban")
								 
								// go cl.DeleteOtherFromChat(op1, []string{op1})
								go killAll(cl, op1, op2, op3)
								go func() { CclList(cl, op1, op3) }()
								// go func() { CclList(cl, op1, op3) }()

								// WarTime[op1] = time.Now()
							}
							fmt.Println("********************************")

							/*
								OperationType_INVITE_INTO_CHAT
								เชิญเข้าสู่การแชท
							*/
						case 123:
							fmt.Println("เชิญเข้าสู่การแชท")
							op1 := op.Param1
							op3 := op.Param3
							go KickBl(cl, op1) // ตรวจสอบคนถูกแบนและให้ออกจากลุ่ม
							go InviteMem(cl, op1, op3)
							
						case 133:
							fmt.Println("*************133****************")
							//Kicked
							// op1  คนถูกเตะ
							// op2  คนเตะ
							// op1, ไอดีห้อง
							// op2, ผู้เชิญ
							// op3 คนถูกเชิญ
							op1, op2, op3, ctime := op.Param1, op.Param2, op.Param3, op.CreatedTime

							fmt.Println("ctime", ctime)
							// fmt.Println("ctime", op1, op2, op3)
							if fullAccess(op2) {
								fmt.Println("********fullAccess********")
								continue
							}
							fmt.Println("********Ban********")
							go cl.DeleteOtherFromChat(op1, []string{op2})
							if fullAccess(op3) {
								go cl.FindAndAddContactsByMid(op3)
								go cl.InviteIntoChat(op1, []string{op3})
								go InviteMem(cl, op1, op3)
							}
							// cl.InviteIntoChat(to, []string{usr})
							go Ban(op2)
							fmt.Println("*************133****************")
						case 55:
							// op1  ห้อง
							// op2  คนอ่าน
							fmt.Println("*************55**อ่าน**************")
							op1, op2 := op.Param1, op.Param2
							if get_ban(op2) {
								go cl.DeleteOtherFromChat(op1, []string{op2})
							}
							//Join
						case 130:
							// op1  ห้อง
							// op2  คนเข้า
							fmt.Println("*************130**Join**************")
							// op1, op2, op3 := op.Param1, op.Param2, op.Param3
							op1, op2 := op.Param1, op.Param2
							if fullAccess(op2) {
								fmt.Println("********fullAccess********")
								continue
							} else if kickban == true {
								fmt.Println("*************130**kickban**************")
								if get_ban(op2) {
									go cl.DeleteOtherFromChat(op1, []string{op2})
								}
								// WarTime[op1] = time.Now()
							}
							fmt.Println("*************end****************")
						case 26:
							cl.Rev = -1
							// ctime := op.CreatedTime
							msg := op.Message
							text := op.Message.Text
							sender := msg.From_
							var to = msg.To
							var pesan = strings.ToLower(text)

							if (op.Message.ContentType).String() == "NONE" {
								if get_status("ProLINK") && strings.Contains(pesan, "http") || strings.Contains(pesan, "lin") {
									if !fullAccess(sender) {
										cl.DeleteOtherFromChat(to, []string{sender})
										go Ban(sender)
										cl.SendMessage(msg.To, "❌กันลิ้งค์มิจฉาชีพ❌")
										continue
									}
								}
							}

							if msg.ContentType == 0 {
								Msg := string(msg.Text)
								box := strings.Split((Msg), ",")
								for TX := range box {
									if TX != 0 {
										time.Sleep(1 * time.Second)
									}
									text := string(box[TX])
									txt := strings.ToLower(text)
									var dataMention = []string{}
									if _, cek := msg.ContentMetadata["MENTION"]; cek {
										mentions := oop.Mentions{}
										json.Unmarshal([]byte(msg.ContentMetadata["MENTION"]), &mentions)
										for _, v := range mentions.MENTIONEES {
											if !oop.Contains(dataMention, v.M) {
												dataMention = append(dataMention, v.M)
											}
										}
									}
									if !fullAccess(sender) {
										if get_status("msgban") {
											cl.SendMessage(msg.To, "❌กันข้อความมิจฉาชีพ❌")
											cl.DeleteOtherFromChat(to, []string{sender})
											go Ban(sender)
											// cl.SendMessage(msg.To, "❌กันข้อความมิจฉาชีพ❌")
										}
										continue
									}
									// bot root
									if Squads[0] != cl.Mid {
										continue
									}
									if txt == "ดูดำ" {
										banlist := ban_list()
										if len(banlist) != 0 {
											tx := "• Banlist\n\n"
											target := []string{}
											for x := range banlist {
												if banlist[x] != "" {
													tx += fmt.Sprintf("	%v. @!\n", x+1)
													target = append(target, banlist[x])
												}
											}
											cl.SendMention(to, tx, target)
										} else {
											cl.SendMessage(to, "Not have banlist")
										}
									} else if txt == "ลบดำเปิด" {
										delBlacklist = true
										cl.SendMessage(to, "เปิดส่งคอนแทคคนที่ลงดำ")

									} else if txt == "ลบดำปิด" {
										delBlacklist = false
										cl.SendMessage(to, "ปิดส่งคอนแทคคนที่ลงดำ")
									}
								}
							} else if (op.Message.ContentType).String() == "CONTACT" {
								name := op.Message.ContentMetadata["displayName"]
								mid := op.Message.ContentMetadata["mid"]
								if PromoteBlacklist == true {
									if fullAccess(sender) {
										if !get_ban(mid) { 
											go Ban(mid)
											cl.SendMessage(to, "เพิ่ม "+name+" เข้าบัญชีดำเรียบร้อย")
										}
									}
								} else if delBlacklist == true {
									if fullAccess(sender) {
										if remove_ban(mid) { 
											cl.SendMessage(to, "ลบ "+name+" ออกจากบัญชีดำเรียบร้อย")
										}else {
											cl.SendMessage(to, "ลบ "+name+" ออกจากบัญชีไม่สำเร็จ")
										}
										// data.Ban = oop.Remove(data.Ban, mid)
									}
								}
							}else if (op.Message.ContentType).String() == "FLEX" {
								if get_status("ProFLEX") {
									if fullAccess(  cl.Mid) {
										if !fullAccess(sender) {
											cl.DeleteOtherFromChat(to, []string{sender})
											Ban(sender)
											cl.SendMessage(to, "❌💫ห้าม💫โฆษณาflex❌")
										}
									}
								}
							} else if (op.Message.ContentType).String() == "CHATEVENT" {
								if get_status("ProDelAlbum") && op.Message.ContentMetadata["LOC_KEY"] == "BD" {
									if fullAccess(cl.Mid) {
										if !fullAccess(sender) {
											cl.DeleteOtherFromChat(to, []string{sender})
											Ban(sender)
											cl.SendMessage(to, "🪶💫ห้าม💫ลบอัลบั้ม🪶")
										}
									}
								}
							} else if (op.Message.ContentType).String() == "STICKER" {
								if get_status("ProSTICKER"){
									if fullAccess(cl.Mid) {
										if !fullAccess(sender) {
											cl.DeleteOtherFromChat(to, []string{sender})
											Ban(sender)
											cl.SendMessage(to, "🪶💫ห้าม💫ส่งสติ้กเกอร์🪶")
										}
									}
								}
							} else if (op.Message.ContentType).String() == "CALL" {
								if get_status("ProCALL") && op.Message.ContentMetadata["GC_MEDIA_TYPE"] == "AUDIO" {
									if fullAccess( cl.Mid) {
										if !fullAccess(sender) {
											cl.DeleteOtherFromChat(to, []string{sender})
											Ban(sender)
											cl.SendMessage(to, "🪶💫ห้าม💫โทรกลุ่ม🪶")
										}
									}
								}
							} else if (op.Message.ContentType).String() == "FILE" {
								if get_status ("ProFILE"){
									if fullAccess(  cl.Mid) {
										if !fullAccess(sender) {
											cl.DeleteOtherFromChat(to, []string{sender})
											Ban(sender)
											cl.SendMessage(to, "🪶💫ห้าม💫ส่งไฟล์🪶")
										}
									}
								}
							} else if (op.Message.ContentType).String() == "POSTNOTIFICATION" {
								if get_status("ProPOSTNOTIFICATION") {
									if fullAccess( cl.Mid) {
										if !fullAccess(sender) {
											cl.DeleteOtherFromChat(to, []string{sender})
											Ban(sender)
											cl.SendMessage(to, "🪶💫ห้าม💫สมาชิกโน้ต&&แชร์โพส🪶")
										}
									}
								}
							} else if (op.Message.ContentType).String() == "AUDIO" {
								if get_status("ProAUDIO") {
									if fullAccess( cl.Mid) {
										if !fullAccess(sender) {
											cl.DeleteOtherFromChat(to, []string{sender})
											Ban(sender)
											cl.SendMessage(to, "🪶💫ห้าม💫ส่งคลิปเสียง🪶")
										}
									}
								}
							}else if (op.Message.ContentType).String() == "IMAGE" {
								if fullAccess(sender) {
									if get_status("ProIMAGE") {
										if fullAccess(  cl.Mid) {
											cl.DeleteOtherFromChat(to, []string{sender})
											Ban(sender)
											cl.SendMessage(to, "🪶💫ห้าม💫ส่งรูปภาพ🪶")
										}
									}
									if _, cek := updateImage[cl.Mid]; cek {
										time.Sleep(10 * time.Second)
										path, err := cl.DownloadObjectMsg(msg.ID, "bin")
										if err != nil {
											cl.SendMessage(to, "Error download pict.")
											return
										}
										cl.UpdateProfilePicture(path, "p")
										delete(updateImage, cl.Mid)
										cl.SendMessage(to, "Picture updated")
									} else if _, cek := updateCover[cl.Mid]; cek {
										time.Sleep(10 * time.Second)
										delete(updateCover, cl.Mid)
										cl.SendMessage(to, "Cover updated")
									} else if _, cek := updateVideo.Mid[cl.Mid]; cek {
										if updateVideo.PictStatus {
											time.Sleep(10 * time.Second)
											path, err := cl.DownloadObjectMsg(msg.ID, "bin")
											if err != nil {
												cl.SendMessage(to, "Error download pict.")
												return
											}
											updateVideo.PictPath = path
											if updateVideo.Tipe == "cvp" {
												cl.UpdateProfilePictureVideo(updateVideo.PictPath, updateVideo.VideoPath)
												delete(updateVideo.Mid, cl.Mid)
												cl.SendMessage(to, "Picture video updated")
											}
											updateVideo.Tipe = ""
											updateVideo.PictStatus = false
											updateVideo.PictPath = ""
											updateVideo.VideoPath = ""
										}
									}
								}
							} else if (op.Message.ContentType).String() == "VIDEO" {
								if get_status("ProVIDEO")  {
									if fullAccess(cl.Mid) {
										if !fullAccess(sender) {
											cl.DeleteOtherFromChat(to, []string{sender})
											Ban(sender)
											cl.SendMessage(to, "🪶💫ห้าม💫ส่งวีดีโอ🪶")
										}
									}
									if _, cek := updateVideo.Mid[cl.Mid]; cek {
										if updateVideo.VideoStatus {
											time.Sleep(10 * time.Second)
											path, err := cl.DownloadObjectMsg(msg.ID, "bin")
											if err != nil {
												cl.SendMessage(to, "Error download video.")
												return
											}
											updateVideo.VideoPath = path
											updateVideo.VideoStatus = false
											updateVideo.PictStatus = true
											cl.SendMessage(to, "Please send image !.")
										}
									}
								}
							}
						}
					}
				}
			}
		}(ops)
		cl.Rev = -1
		for _, op := range ops {
			if op.Revision != -1 {
				cl.CorrectRevision(op, true, false, false)
			} else {
				cl.CorrectRevision(op, false, true, true)
			}
		}
	}

}
func putSquad(cl *oop.Account, to string) {
	chat, _ := cl.GetChats([]string{to}, true, true)
	// fmt.Println(" chat", chat)
	if chat != nil {
		members := chat.Chats[0].Extra.GroupExtra.MemberMids
		Pending := chat.Chats[0].Extra.GroupExtra.InviteeMids
		if _, cek := data.StayGroup[to]; cek {
			for b := range Squads {
				if _, cek := members[Squads[b]]; cek {
					if !oop.Contains(data.StayGroup[to], Squads[b]) {
						data.StayGroup[to] = append(data.StayGroup[to], Squads[b])
					}
				} else if _, cek := Pending[Squads[b]]; cek {
					if !oop.Contains(data.StayPending[to], Squads[b]) {
						data.StayPending[to] = append(data.StayPending[to], Squads[b])
					}
				} else {
					if oop.Contains(data.StayGroup[to], Squads[b]) {
						data.StayGroup[to] = oop.Remove(data.StayGroup[to], Squads[b])
					}
					if oop.Contains(data.StayPending[to], Squads[b]) {
						data.StayPending[to] = oop.Remove(data.StayPending[to], Squads[b])
					}
				}
			}
		} else {
			for b := range Squads {
				if _, cek := members[Squads[b]]; cek {
					data.StayGroup[to] = append(data.StayGroup[to], Squads[b])
				} else if _, cek := Pending[Squads[b]]; cek {
					if !oop.Contains(data.StayPending[to], Squads[b]) {
						data.StayPending[to] = append(data.StayPending[to], Squads[b])
					}
				} else {
					if oop.Contains(data.StayGroup[to], Squads[b]) {
						data.StayGroup[to] = oop.Remove(data.StayGroup[to], Squads[b])
					}
					if oop.Contains(data.StayPending[to], Squads[b]) {
						data.StayPending[to] = oop.Remove(data.StayPending[to], Squads[b])
					}
				}
			}
		}
	}
}
func SaveData() {
	file, _ := json.MarshalIndent(data, "", "    ")
	ioutil.WriteFile(dataPath, file, 0644)
}
func getBots() { //update bot
	fmt.Println("Read Data bot")
	dataRead_, err := ioutil.ReadFile(dataPathB)
	if err != nil {
		fmt.Println("Read Data", err)
	}
	json.Unmarshal(dataRead_, &Squads)
	fmt.Println("Read Data bot", Squads)
}
func getAccessForCancel(cl *oop.Account, op2 string, op3 []string) bool {
	if oop.Contains(op3, cl.Mid) {
		return false
	} else if _, cek := BotKor[cl.Mid]; cek {
		return true
	}
	return false
}
func CclList(cl *oop.Account, to string, target []string) {
	for i := 0; i < len(target); i++ {
		go cl.CancelChatInvitation(to, []string{target[i]})
	}
}
func KickBl(cl *oop.Account, to string) {
	chat, _ := cl.GetChats([]string{to}, true, true)
	if chat != nil {
		members := chat.Chats[0].Extra.GroupExtra.MemberMids
		Invitee := chat.Chats[0].Extra.GroupExtra.InviteeMids
		if len(members) != 0 {
			target := []string{}
			target2 := []string{}
			banlist := getBan()
			for x := range getBan() {
				if _, cek := members[banlist[x]]; cek {
					target = append(target, banlist[x])
				} else if _, cek := Invitee[banlist[x]]; cek {
					target2 = append(target2, banlist[x])
				}
			}
			go KickAndCancelByList(cl, to, target, target2)
		}
	}
}
func KickAndCancelByList(cl *oop.Account, to string, targetMem []string, targetInv []string) {
	if len(targetInv) != 0 { 
		for i := 0; i < len(targetInv); i++ {
			go func(i int) { 
				go cl.CancelChatInvitation(to, []string{targetInv[i]})
			}(i)
		}
	}
	if len(targetMem) != 0 {
		
		for i := 0; i < len(targetMem); i++ {
			go func(i int) {
				go cl.DeleteOtherFromChat(to, []string{targetMem[i]})
			}(i)
		}
	}
}
func fullAccess(target string) bool {
	// Maker dev
	for i := 0; i < len(Squads); i++ {
		if target == Squads[i] {
			return true
		}
	}
	// var arr_data = []string{}
	arr_data = Maker
	// looper := len(arr_data)
	//looper := len(arr_data)
	for i := 0; i < len(arr_data); i++ {
		if target == arr_data[i] {
			return true
		}
	}
	// God เทพ
	arr_data = data.God
	// looper2 := len(arr_data)
	for i := 0; i < len(arr_data); i++ {
		if target == arr_data[i] {
			return true
		}
	}
	// Admin
	arr_data = data.Admin
	// looper4 := len(arr_data)
	for i := 0; i < len(arr_data); i++ {
		if target == arr_data[i] {
			return true
		}
	}

	return false
}
func getWarAccess(cl *oop.Account, ct int64, op1 string, op3 string, mid string, trobos bool) bool {
	if op3 != "" {
		BotKor[op3] = true
	}
	if _, cek := BotKor[cl.Mid]; cek && trobos {
		return true
	} else if pala != "" {
		if pala == mid && pala != op3 {
			return true
		} else if op3 == pala && cl.Mid != lockMod {
			if _, cek := cTime[ct]; cek {
				return false
			}
			if len(cTime) == 500 {
				cTime = map[int64]bool{}
			}
			cTime[ct] = true
			if Loop && Action(cl, op1) {
				return false
			}
			pala = mid
			return true
		} else {
			if lock == "" {
				lock = fmt.Sprintf("%v", ct)
				lockMod = cl.Mid
				ticket, _ := cl.ReissueChatTicket(op1)
				if ticket != nil {
					ModTicket = fmt.Sprintf("%v", ticket.TicketId) //input LockMod
					if AutoTimeSet(op1) {
						go AutoSetUp(cl, op1)
					}
					TimeSet[op1] = time.Now()
				}
			} else {
				if lockMod == cl.Mid {
					//fmt.Println(len(BotKor))
					if len(BotKor) >= len(data.StayGroup[op1])-1 {
						if Loop && Action(cl, op1) {
							return false
						}
						lockMod = ""
						pala = cl.Mid
						lock = ""
						if AutoMulty {
							//time.Sleep(2000 * time.Millisecond)
							Multy = false
						}
						return true
					}
				}
			}
			return false
		}
	} else {
		if _, cek := cTime[ct]; cek {
			return false
		}
		if len(cTime) == 50 {
			cTime = map[int64]bool{}
		}
		cTime[ct] = true
		pala = mid
		return true
	}
}
func BanAll(usr string, sq []string) {
	for x := range sq {
		if !oop.Contains(data.Ban, sq[x]) {
			data.Ban = append(data.Ban, sq[x])
		}
	}
	if !oop.Contains(data.Ban, usr) {
		data.Ban = append(data.Ban, usr)
	}
	SaveData()
}
func Ban(usr string) {
	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	defer sqliteDatabase.Close()
	insertBotSQL := `INSERT INTO ban (id) VALUES (?)`
	statement, err := sqliteDatabase.Prepare(insertBotSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = statement.Exec(usr)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("data.Ban", usr)
}
func getBan() []string {
	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	defer sqliteDatabase.Close()
	row, err := sqliteDatabase.Query("SELECT * FROM ban WHERE id")
	if err != nil {
		fmt.Println(err)
	}
	defer row.Close()
	banlist := []string{}
	for row.Next() { // Iterate and fetch the records from result cursor
		var id string 
		row.Scan(&id,)
		banlist = append(banlist, id) 
	}
	 
	fmt.Println("data.Ban", banlist)
	return banlist
}
func Action(cl *oop.Account, to string) bool {
	//time.Sleep(50 * time.Millisecond)
	runtime.GOMAXPROCS(2)
	chat, err := cl.GetChats([]string{to}, false, false)
	if err != nil {
		fmt.Println(err)
		return true
	}
	if chat != nil {
		if chat.Chats[0].Extra.GroupExtra.PreventedJoinByTicket == true {
			return false
		}
	}
	return true
}
func AutoSetUp(cl *oop.Account, to string) {
	time.Sleep(50 * time.Second)
	if AccessWarTime(to) {
		TimeSet[to] = time.Now()
		AutoSetUp(cl, to)
		println("Playing")
		return
	}

	lock = ""
	TempJoin = "0"
	pala = ""
	if AutoMulty {
		Multy = true
	}
	WarTime = make(map[string]time.Time)
	println("Auto SetUp Success")
	play := 0
	if len(BotKor) <= len(data.StayGroup[to]) {
		play = len(BotKor)
	} else {
		play = len(BotKor) + 1
	}
	BotKor = make(map[string]bool)
	for b := range Botlist {
		Botlist[b].SpamJoin = 0
	}
	inviteSquad(cl, to)
	oop.Clearcache()
	if len(data.Ban) != 0 && AutoClearban {
		ModTicket = ""
		lockMod = ""
		tx := "• Auto ClearBans\n\n"
		target := []string{}
		for x := range data.Ban {
			if data.Ban[x] != "" {
				tx += fmt.Sprintf("	%v. @!\n", x+1)
				target = append(target, data.Ban[x])
			}
		}
		tx += fmt.Sprintf("\n	*%v 🔖", play)
		cl.SendMention(to, tx, target)
		data.Ban = []string{}
		time.Sleep(1 * time.Second)
		for b := range Botlist {
			if oop.Contains(data.StayGroup[to], Botlist[b].Mid) {
				Botlist[b].SendMessage(to, "hello gaes..")
			}
		}
	}
}
func AutoTimeSet(group string) bool {
	if _, ok := TimeSet[group]; ok && time.Since(TimeSet[group]) < 500*time.Millisecond {
		return false
	}
	return true
}
func AccessWarTime(group string) bool {
	if _, ok := WarTime[group]; ok && time.Since(WarTime[group]) < 1000*time.Millisecond {
		return true
	}
	return false
}

func inviteSquad(cl *oop.Account, to string) {
	runtime.GOMAXPROCS(1)
	chat, _ := cl.GetChats([]string{to}, true, false)
	if chat != nil {
		target := []string{}
		members := chat.Chats[0].Extra.GroupExtra.MemberMids
		for b := range data.StayGroup[to] {
			if _, cek := members[data.StayGroup[to][b]]; !cek {
				target = append(target, data.StayGroup[to][b])
			}
		}
		var wg sync.WaitGroup
		wg.Add(len(target))
		for i := 0; i < len(target); i++ {
			go func(i int) {
				defer wg.Done()
				cl.InviteIntoChat(to, []string{target[i]})
			}(i)
		}
		wg.Wait()
	}
}
func killAll(cl *oop.Account, op1, usr string, sq []string) {
	for x := range sq {
		if get_ban(sq[x]) {
			Ban(sq[x])
			go cl.DeleteOtherFromChat(op1, []string{sq[x]})
		}
	}
	if !get_ban(usr) {
		// data.Ban = append(data.Ban, )
		Ban(usr)
		go cl.DeleteOtherFromChat(op1, []string{usr})
	}
	fmt.Println("killAll")
}
func InviteMem(cl *oop.Account, to string, usr string) {
	res, _ := cl.GetAllContactIds()
	if !oop.Contains(res, usr) {
		cl.FindAndAddContactsByMid(usr)
	}
	cl.InviteIntoChat(to, []string{usr})
}
func ModJoin(cl *oop.Account, to string, op2 string) {

	for {
		chat, _ := cl.GetChats([]string{to}, false, false)
		if chat != nil {
			if !chat.Chats[0].Extra.GroupExtra.PreventedJoinByTicket {
				go func(to string) {
					go cl.AcceptChatInvitationByTicket(to, ModTicket)
					go cl.DeleteOtherFromChat(to, []string{op2})
					go KickBl(cl, to)
				}(to)
				return
			} else {
				ticket, _ := cl.ReissueChatTicket(to)
				if ticket != nil {
					ModTicket = fmt.Sprintf("%v", ticket.TicketId)
				}
				continue
			}
		} else {
			return
		}
	}
}
func closeQr(cl *oop.Account, to string) {
	time.Sleep(time.Duration(cl.Count) * time.Millisecond)
	runtime.GOMAXPROCS(1)
	chat, err := cl.GetChats([]string{to}, false, false)
	if err != nil {
		return
	}
	if !chat.Chats[0].Extra.GroupExtra.PreventedJoinByTicket {
		go func() { cl.UpdateChatQr(to, true) }()
	}
}
func appendBl(target string) {
	time.Sleep(1 * time.Millisecond)
	if !oop.Contains(data.Ban, target) {
		if !fullAccess(target) {
			data.Ban = append(data.Ban, target)
		}
	}
}

// for num, auth := range Token {
// 	xcl := oop.Connect(num, auth)
// 	err := xcl.LoginWithAuthToken()
// 	if err == nil {
// 		go func(num int) {
// 			// perBots(xcl)
// 		}(num)
// 		time.Sleep(1 * time.Second)

// 	} else {
// 		if strings.Contains(fmt.Sprintf("%s", err), "suspended") {
// 			fmt.Println(auth[:33], "FREEZE OR SUSPEND")
// 		} else {
// 			fmt.Println(auth[:33], err)
// 		}
// 		continue
// 	}
// }
// SaveData()
// ch := make(chan int, len(Token))
// for v := range ch {

// 	if v == 51 {

// 		break
// 	}
// }
// }
