package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"./Library/linethrift"
	"./Library/oop"
	"github.com/kardianos/osext"

	// "botline/Library-mac/oop"
	// "botline/Library-mac/linethrift"

)

type User struct {
	Squad               []string             `json:"squad"`
	Owner               []string             `json:"owner"`
	Admin               []string             `json:"admin"`
	Staff               []string             `json:"staff"`
	Osquad              []string             `json:"osquad"`
	Ban                 []string             `json:"ban"`
	TargetSpam          []string             `json:"targetspam"`
	Gmember          []string             `json:"Gmember"`
	LimitStatus         map[string]bool      `json:"limitstatus"`
	LimitTime           map[string]time.Time `json:"limittime"`
	ProReadKick             map[string]bool      `json:"proreadkick"`
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
	ProKillMsg          map[string]bool      `json:"ProKillMsg"`
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

var cpu int = 1
var RAM int = 1
var duedatecount int = 0
var (
	data      User
	dataPath  = fmt.Sprintf("data.json")
	toeknPath = fmt.Sprintf("token.txt")
	Maker     = []string{
		"u53ab6fa03c2838678a07a10fd142eb81",
	}
	Freeze           = []string{}
	KillMod          = false
	GroupList        = []string{}
	
	Botlist          []*oop.Account
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
	notiFadd         = false
	kickban          = false
	sleepmode        = true
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

func SaveData() {
	file, _ := json.MarshalIndent(data, "", "    ")
	ioutil.WriteFile(dataPath, file, 0644)
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
func IsBanAddFL(cl *oop.Account, from string) bool {
	if !oop.Contains(data.Ban, from) {
		data.Ban = append(data.Ban, from)
		return true
	}
	return false
}
func getAccessForCancel(cl *oop.Account, op2 string, op3 []string) bool {
	if oop.Contains(op3, cl.Mid) {
		return false
	} else if _, cek := BotKor[cl.Mid]; cek {
		return true
	}
	return false
}
func getAccess(ct int64, mid string) bool {
	if _, cek := cTime[ct]; cek {
		if palaMsg == mid {
			return true
		}
		return false
	}
	if len(cTime) == 50 {
		cTime = map[int64]bool{}
	}
	cTime[ct] = true
	palaMsg = mid
	return true
}
func getAccessAjs(ct int64) bool {
	if _, cek := aTime[ct]; cek {
		return false
	}
	if len(cTime) == 50 {
		aTime = map[int64]bool{}
	}
	aTime[ct] = true
	return true
}
func getAccessJoin(ct int64) bool {
	if _, cek := aTime[ct]; cek {
		return false
	}
	if len(cTime) == 50 {
		aTime = map[int64]bool{}
	}
	aTime[ct] = true
	return true
}

func putAccess(ct int64) bool {
	if _, cek := pTime[ct]; cek {
		return false
	}
	if len(pTime) == 50 {
		pTime = map[int64]bool{}
	}
	pTime[ct] = true
	return true
}
func AccessWarTime(group string) bool {
	if _, ok := WarTime[group]; ok && time.Since(WarTime[group]) < 1000*time.Millisecond {
		return true
	}
	return false
}
func AccessWars(from string, group string) bool {
	if _, ok := WarTime[group]; ok && time.Since(WarTime[group]) < 1000*time.Millisecond {
		return true
	}
	return false
}
func IsBanned(from string) bool {
	if oop.Contains(data.Ban, from) == true {
		return true
	}
	return false
}
func Ban(usr string) {
	if !oop.Contains(data.Ban, usr) {
		data.Ban = append(data.Ban, usr)
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
func appendBl(target string) {
	time.Sleep(1 * time.Millisecond)
	if !oop.Contains(data.Ban, target) {
		data.Ban = append(data.Ban, target)
	}
}
func BanWithList(sq []string) {
	for x := range sq {
		if !oop.Contains(data.Ban, sq[x]) {
			data.Ban = append(data.Ban, sq[x])
		}
	}
	SaveData()
}
func Gasak(cl *oop.Account, to string) {
	for b := range Botlist {
		if oop.Contains(Freeze, Botlist[b].Mid) {
			continue
		}
		if oop.Contains(data.StayGroup[to], data.Squad[b]) {
			if _, cek := BotKor[Botlist[b].Mid]; !cek && Botlist[b].Mid != lockMod && Botlist[b].Mid != cl.Mid {
				time.Sleep(1 * time.Millisecond)
				cl.UpdateChatQr(to, true)
				KickBl(Botlist[b], to)
				inviteSquad(cl, to)
				return
			}
		}
	}
}

func putSquad(cl *oop.Account, to string) {
	chat, _ := cl.GetChats([]string{to}, true, true)
	if chat != nil {
		members := chat.Chats[0].Extra.GroupExtra.MemberMids
		Pending := chat.Chats[0].Extra.GroupExtra.InviteeMids
		if _, cek := data.StayGroup[to]; cek {
			for b := range data.Squad {
				if _, cek := members[data.Squad[b]]; cek {
					if !oop.Contains(data.StayGroup[to], data.Squad[b]) {
						data.StayGroup[to] = append(data.StayGroup[to], data.Squad[b])
					}
				} else if _, cek := Pending[data.Squad[b]]; cek {
					if !oop.Contains(data.StayPending[to], data.Squad[b]) {
						data.StayPending[to] = append(data.StayPending[to], data.Squad[b])
					}
				} else {
					if oop.Contains(data.StayGroup[to], data.Squad[b]) {
						data.StayGroup[to] = oop.Remove(data.StayGroup[to], data.Squad[b])
					}
					if oop.Contains(data.StayPending[to], data.Squad[b]) {
						data.StayPending[to] = oop.Remove(data.StayPending[to], data.Squad[b])
					}
				}
			}
		} else {
			for b := range data.Squad {
				if _, cek := members[data.Squad[b]]; cek {
					data.StayGroup[to] = append(data.StayGroup[to], data.Squad[b])
				} else if _, cek := Pending[data.Squad[b]]; cek {
					if !oop.Contains(data.StayPending[to], data.Squad[b]) {
						data.StayPending[to] = append(data.StayPending[to], data.Squad[b])
					}
				} else {
					if oop.Contains(data.StayGroup[to], data.Squad[b]) {
						data.StayGroup[to] = oop.Remove(data.StayGroup[to], data.Squad[b])
					}
					if oop.Contains(data.StayPending[to], data.Squad[b]) {
						data.StayPending[to] = oop.Remove(data.StayPending[to], data.Squad[b])
					}
				}
			}
		}
	}
}

func kickAndInvite(cl *oop.Account, to string) {
	chat, _ := cl.GetChats([]string{to}, true, false)
	if chat != nil {
		members := chat.Chats[0].Extra.GroupExtra.MemberMids
		Squad := []string{}
		for b := range data.Squad {
			if _, cek := members[data.Squad[b]]; !cek {
				Squad = append(Squad, data.Squad[b])
			}
		}
		go cl.InviteIntoChat(to, Squad)
		for x := range data.Ban {
			if _, cek := members[data.Ban[x]]; cek {
				go cl.DeleteOtherFromChat(to, []string{data.Ban[x]})
			}
		}
	}
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

func inviteSquadRandom(cl *oop.Account, to string) {
	runtime.GOMAXPROCS(1)
	rand.Seed(time.Now().Unix())
	chat, _ := cl.GetChats([]string{to}, true, false)
	if chat != nil {
		members := chat.Chats[0].Extra.GroupExtra.MemberMids
		ListTarget := []string{}
		for b := range data.StayGroup[to] {
			if _, cek := members[data.StayGroup[to][b]]; !cek {
				ListTarget = append(ListTarget, data.StayGroup[to][b])
			}
		}
		n := rand.Int() % len(ListTarget)
		f := fmt.Sprintf("%v", n)
		go cl.InviteIntoChat(to, []string{f})
	}
}

func inviteByList(cl *oop.Account, to string, target []string) {
	runtime.GOMAXPROCS(1)
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

func inviteAllBots(cl *oop.Account, to string) {
	chat, _ := cl.GetChats([]string{to}, true, false)
	if chat != nil {
		members := chat.Chats[0].Extra.GroupExtra.MemberMids
		target1 := []string{}
		target2 := []string{}
		for b := range data.Squad {
			if _, cek := members[data.Squad[b]]; !cek && data.Squad[b] != cl.Mid {
				if oop.Contains(data.StayGroup[to], data.Squad[b]) {
					target1 = append(target1, data.Squad[b])
					continue
				}
				if !oop.Contains(data.StayPending[to], data.Squad[b]) {
					target2 = append(target2, data.Squad[b])
				}
			}
		}
		if len(target2) != 0 {
			cl.InviteIntoChat(to, target2)
		}
		if len(target1) != 0 {
			cl.InviteIntoChat(to, target1)
		} //if len(target2) != 0 {inviteByList(cl,to, target2)}
	}
}

func inviteAllBots2(cl *oop.Account, to string) {
	chat, _ := cl.GetChats([]string{to}, true, false)
	if chat != nil {
		members := chat.Chats[0].Extra.GroupExtra.MemberMids
		for b := range data.Squad {
			if _, cek := members[data.Squad[b]]; !cek {
				if !oop.Contains(data.StayPending[to], data.Squad[b]) {
					cl.InviteIntoChat(to, []string{data.Squad[b]})
					//time.Sleep(1 * time.Millisecond)
				}
			}
		}
	}
}
func Kick(cl *oop.Account, to string, usr string) {
	cl.DeleteOtherFromChat(to, []string{usr})
}
func InviteMem(cl *oop.Account, to string, usr string) {
	res, _ := cl.GetAllContactIds()
	if !oop.Contains(res, usr) {
		cl.FindAndAddContactsByMid(usr)
	}
	cl.InviteIntoChat(to, []string{usr})
}

func ModJoinOld(cl *oop.Account, to string, op2 string) {
	if ModTicket == "" {
		time.Sleep(1 * time.Millisecond)
	}
	runtime.GOMAXPROCS(1)
	for {
		if AccessWarTime(to) && Multy {
			chat, _ := cl.GetChats([]string{to}, false, false)
			if chat != nil {
				if !chat.Chats[0].Extra.GroupExtra.PreventedJoinByTicket {
					var wg sync.WaitGroup
					wg.Add(1)
					go func(to string) {
						defer wg.Done()
						go cl.AcceptChatInvitationByTicket(to, ModTicket)
						go cl.DeleteOtherFromChat(to, []string{op2})

						go KickBl(cl, to)
						cl.SpamJoin++
						//go cl.UpdateChatQr(to, false) //go InvQr(cl, to)
						if cl.SpamJoin >= LimiterJoin {
							go closeQr(cl, to)
							//cl.SpamJoin = 0
						}
					}(to)
					wg.Wait()
					return
				} else {
					time.Sleep(1 * time.Millisecond)
					continue
				}
			} else {
				return
			}
			continue
		} else {
			return
		}
	}
	return
}
func ModJoin(cl *oop.Account, to string, op2 string) {
	if ModTicket == "" {
		time.Sleep(1 * time.Millisecond)
	}
	runtime.GOMAXPROCS(1)
	for {
		if AccessWarTime(to) && Multy {
			chat, _ := cl.GetChats([]string{to}, false, false)
			if chat != nil {
				if !chat.Chats[0].Extra.GroupExtra.PreventedJoinByTicket {
					var wg sync.WaitGroup
					wg.Add(1)
					go func(to string) {
						defer wg.Done()
						go cl.AcceptChatInvitationByTicket(to, ModTicket)
						go cl.DeleteOtherFromChat(to, []string{op2})

						go KickBl(cl, to)
						cl.SpamJoin++
						//go cl.UpdateChatQr(to, false) //go InvQr(cl, to)
						if cl.SpamJoin >= LimiterJoin {
							go closeQr(cl, to)
							//cl.SpamJoin = 0
						}
					}(to)
					wg.Wait()
					return
				} else {
					time.Sleep(1 * time.Millisecond)
					continue
				}
			} else {
				return
			}
			continue
		} else {
			return
		}
	}
	return
}
func KickBl(cl *oop.Account, to string) {
	time.Sleep(time.Duration(cl.Count) * time.Millisecond)
	runtime.GOMAXPROCS(1)
	chat, _ := cl.GetChats([]string{to}, true, true)
	if chat != nil {
		members := chat.Chats[0].Extra.GroupExtra.MemberMids
		Invitee := chat.Chats[0].Extra.GroupExtra.InviteeMids
		if len(members) != 0 {
			target := []string{}
			target2 := []string{}
			for x := range data.Ban {
				if _, cek := members[data.Ban[x]]; cek {
					target = append(target, data.Ban[x])
				} else if _, cek := Invitee[data.Ban[x]]; cek {
					target2 = append(target2, data.Ban[x])
				}
			}
			go KickAndCancelByList(cl, to, target, target2)
		}
	}
}

func InvQr(cl *oop.Account, to string) {
	time.Sleep(time.Duration(cl.Count) * time.Millisecond)
	runtime.GOMAXPROCS(1)
	chat, _ := cl.GetChats([]string{to}, false, false)
	if chat != nil {
		if chat.Chats[0].Extra.GroupExtra.PreventedJoinByTicket {
			go func() { cl.UpdateChatQr(to, false) }()
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

func CclList(cl *oop.Account, to string, target []string) {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(len(target))
	for i := 0; i < len(target); i++ {
		go func(i int) {
			defer wg.Done()
			go cl.CancelChatInvitation(to, []string{target[i]})
		}(i)
	}
	wg.Wait()
}

func CclBan(cl *oop.Account, to string) {
	runtime.GOMAXPROCS(1)
	chat, _ := cl.GetChats([]string{to}, false, true)
	if chat != nil {
		Invitee := chat.Chats[0].Extra.GroupExtra.InviteeMids
		var wg sync.WaitGroup
		wg.Add(len(data.Ban))
		for x := range data.Ban {
			if _, cek := Invitee[data.Ban[x]]; cek {
				go func(x int) {
					defer wg.Done()
					go cl.CancelChatInvitation(to, []string{data.Ban[x]})
				}(x)
			}
		}
		wg.Wait()
	}
}

func CclList2(cl *oop.Account, to string, target []string) {
	runtime.GOMAXPROCS(1)
	//if !AccessWarTime(to) {return}
	var wg sync.WaitGroup
	wg.Add(len(target))
	for i := 0; i < len(target); i++ {
		//if oop.Contains(data.Ban,target[i]) {
		go func(i int) {
			defer wg.Done()
			go cl.CancelChatInvitation(to, []string{target[i]})
		}(i)
		//}
		wg.Wait()
	}
}

func accByInvite(cl *oop.Account, to string) {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(1)
	for i := 0; i < 1; i++ {
		go func(i int) {
			defer wg.Done()
			cl.AcceptChatInvitation(to)
		}(i)
	}
	wg.Wait()
}

func KickByList(cl *oop.Account, to string, target []string) {

	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(len(target))
	for i := 0; i < len(target); i++ {
		go func(i int) {
			defer wg.Done()
			go cl.DeleteOtherFromChat(to, []string{target[i]})
		}(i)
	}
	wg.Wait()
}
func KickByBl(cl *oop.Account, to string) {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(len(data.Ban))
	for i := 0; i < len(data.Ban); i++ {
		go func(i int) {
			defer wg.Done()
			go cl.DeleteOtherFromChat(to, []string{data.Ban[i]})
		}(i)
	}
	wg.Wait()
}

func KickAndCancelByList(cl *oop.Account, to string, targetMem []string, targetInv []string) {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	if len(targetInv) != 0 {
		wg.Add(len(targetInv))
		for i := 0; i < len(targetInv); i++ {
			go func(i int) {
				defer wg.Done()
				go cl.CancelChatInvitation(to, []string{targetInv[i]})
			}(i)
		}
	}
	if len(targetMem) != 0 {
		wg.Add(len(targetMem))
		for i := 0; i < len(targetMem); i++ {
			go func(i int) {
				defer wg.Done()
				go cl.DeleteOtherFromChat(to, []string{targetMem[i]})
			}(i)
		}
	}
	wg.Wait()
}
func ByPass(cl *oop.Account, to string) {
	time.Sleep(time.Duration(cl.Count) * time.Millisecond)
	time.Sleep(1 * time.Second)
	runtime.GOMAXPROCS(1)
	chat, err := cl.GetChats([]string{to}, true, true)
	if err != nil {
		fmt.Println(chat, err)
		return
	}
	members := chat.Chats[0].Extra.GroupExtra.MemberMids
	Invitee := chat.Chats[0].Extra.GroupExtra.InviteeMids
	allbot := len(data.StayGroup[to])
	allclientGetMem := map[int][]string{}
	allclientGetPen := map[int][]string{}
	Mnum := 1
	Pnum := 1
	client := map[string]int{}
	count := 1
	for x := range data.StayGroup[to] {
		client[data.StayGroup[to][x]] = count
		count++
	}
	for m := range members {
		if !fullAccess(m) {
			allclientGetMem[Mnum] = append(allclientGetMem[Mnum], m)
			if allbot == Mnum {
				Mnum = 1
			}
			Mnum++
		}
	}
	for p := range Invitee {
		if !fullAccess(p) {
			allclientGetPen[Pnum] = append(allclientGetPen[Pnum], p)
			if allbot == Pnum {
				Pnum = 1
			}
			Pnum++
		}
	}
	for c := range Botlist {
		if oop.Contains(data.StayGroup[to], Botlist[c].Mid) {
			if _, cek := allclientGetMem[client[Botlist[c].Mid]]; cek {
				go KickAndCancelByList(Botlist[c], to, allclientGetMem[client[Botlist[c].Mid]], allclientGetPen[client[Botlist[c].Mid]])
				//Botlist[c].SendMessage(to,fmt.Sprintf("mem : %v\npen : %v", len(allclientGetMem[client[Botlist[c].Mid]]),len(allclientGetPen[client[Botlist[c].Mid]])))
			}
		}
	}
}
func IsFriends(cl *oop.Account, from string) bool {
	friendsip, _ := cl.GetAllContactIds()
	for _, a := range friendsip {
		if a == from {
			return true
			break
		}
	}
	return false
}
func Promax(to string) {
	// if _, cek := data.ProKillMsg[to]; !cek{
	// 	data.ProKillMsg[to] = true
	// }
	if _, cek := data.ProKick[to]; !cek {
		data.ProKick[to] = true
	}
	if _, cek := data.ProInvite[to]; !cek {
		data.ProInvite[to] = true
	}
	if _, cek := data.ProCancel[to]; !cek {
		data.ProCancel[to] = true
	}
	if _, cek := data.ProJoin[to]; !cek {
		data.ProJoin[to] = true
	}
	if _, cek := data.ProQr[to]; !cek {
		data.ProQr[to] = true
	}
	if _, cek := data.ProDelAlbum[to]; !cek {
		data.ProDelAlbum[to] = true
	}
	if _, cek := data.ProFLEX[to]; !cek {
		data.ProFLEX[to] = true
	}
	if _, cek := data.ProCALL[to]; !cek {
		data.ProCALL[to] = true
	}
	if _, cek := data.ProVIDEO[to]; !cek {
		data.ProVIDEO[to] = true
	}
	// if _, cek := data.ProIMAGE[to]; !cek{
	// 	data.ProIMAGE[to] = true
	// }
	if _, cek := data.ProAUDIO[to]; !cek {
		data.ProAUDIO[to] = true
	}
	if _, cek := data.ProPOSTNOTIFICATION[to]; !cek {
		data.ProPOSTNOTIFICATION[to] = true
	}
	if _, cek := data.ProFILE[to]; !cek {
		data.ProFILE[to] = true
	}
	// if _, cek := data.ProSTICKER[to]; !cek{
	// 	data.ProSTICKER[to] = true
	// }
	// if _, cek := data.ProLINK[to]; !cek{
	// 	data.ProLINK[to] = true
	// }
}

func Pronull(to string) {
	delete(data.ProKick, to)
	delete(data.ProInvite, to)
	delete(data.ProCancel, to)
	delete(data.ProJoin, to)
	delete(data.ProQr, to)
	delete(data.ProLINK, to)
	delete(data.ProFLEX, to)
	delete(data.ProSTICKER, to)
	delete(data.ProFILE, to)
	delete(data.ProPOSTNOTIFICATION, to)
	delete(data.ProAUDIO, to)
	delete(data.ProIMAGE, to)
	delete(data.ProVIDEO, to)
	delete(data.ProCALL, to)
	delete(data.ProDelAlbum, to)
	delete(data.ProKillMsg, to)
}

func ProkickOff(to string) {
	if _, cek := data.ProKick[to]; cek {
		delete(data.ProKick, to)
	}
}

func ProkickOn(to string) {
	if _, cek := data.ProKick[to]; !cek {
		data.ProKick[to] = true
	}
}

func ProinviteOff(to string) {
	if _, cek := data.ProInvite[to]; cek {
		delete(data.ProInvite, to)
	}
}

func ProinviteOn(to string) {
	if _, cek := data.ProInvite[to]; !cek {
		data.ProInvite[to] = true
	}
}

func ProcancelOff(to string) {
	if _, cek := data.ProCancel[to]; cek {
		delete(data.ProCancel, to)
	}
}

func ProcancelOn(to string) {
	if _, cek := data.ProCancel[to]; !cek {
		data.ProCancel[to] = true
	}
}
func ProjoinOff(to string) {
	if _, cek := data.ProJoin[to]; cek {
		delete(data.ProJoin, to)
	}
}

func ProjoinOn(to string) {
	if _, cek := data.ProJoin[to]; !cek {
		data.ProJoin[to] = true
	}
}
func ProqrOff(to string) {
	if _, cek := data.ProQr[to]; cek {
		delete(data.ProQr, to)
	}
}

func ProqrOn(to string) {
	if _, cek := data.ProQr[to]; !cek {
		data.ProQr[to] = true
	}
}
func ProFLEXOff(to string) {
	if _, cek := data.ProFLEX[to]; cek {
		delete(data.ProFLEX, to)
	}
}

func ProFLEXOn(to string) {
	if _, cek := data.ProFLEX[to]; !cek {
		data.ProFLEX[to] = true
	}
}
func DelAlbumOff(to string) {
	if _, cek := data.ProDelAlbum[to]; cek {
		delete(data.ProDelAlbum, to)
	}
}

func DelAlbumOn(to string) {
	if _, cek := data.ProDelAlbum[to]; !cek {
		data.ProDelAlbum[to] = true
	}
}
func ProLINKOff(to string) {
	if _, cek := data.ProLINK[to]; cek {
		delete(data.ProLINK, to)
	}
}
func ProKillMsgOff(to string) {
	if _, cek := data.ProKillMsg[to]; cek {
		delete(data.ProKillMsg, to)
	}
}
func ProKillMsgOn(to string) {
	if _, cek := data.ProKillMsg[to]; !cek {
		data.ProKillMsg[to] = true
	}
}

func ProLINKOn(to string) {
	if _, cek := data.ProLINK[to]; !cek {
		data.ProLINK[to] = true
	}
}
func ProSTICKEROff(to string) {
	if _, cek := data.ProSTICKER[to]; cek {
		delete(data.ProSTICKER, to)
	}
}

func ProSTICKEROn(to string) {
	if _, cek := data.ProSTICKER[to]; !cek {
		data.ProSTICKER[to] = true
	}
}
func ProCALLOff(to string) {
	if _, cek := data.ProCALL[to]; cek {
		delete(data.ProCALL, to)
	}
}

func ProCALLOn(to string) {
	if _, cek := data.ProCALL[to]; !cek {
		data.ProCALL[to] = true
	}
}
func ProFILEOff(to string) {
	if _, cek := data.ProFILE[to]; cek {
		delete(data.ProFILE, to)
	}
}

func ProFILEOn(to string) {
	if _, cek := data.ProFILE[to]; !cek {
		data.ProFILE[to] = true
	}
}
func ProPOSTNOTIFICATIONOff(to string) {
	if _, cek := data.ProPOSTNOTIFICATION[to]; cek {
		delete(data.ProPOSTNOTIFICATION, to)
	}
}

func ProPOSTNOTIFICATIONOn(to string) {
	if _, cek := data.ProPOSTNOTIFICATION[to]; !cek {
		data.ProPOSTNOTIFICATION[to] = true
	}
}
func ProVIDEOOff(to string) {
	if _, cek := data.ProVIDEO[to]; cek {
		delete(data.ProVIDEO, to)
	}
}

func ProVIDEOOn(to string) {
	if _, cek := data.ProVIDEO[to]; !cek {
		data.ProVIDEO[to] = true
	}
}
func ProAUDIOOff(to string) {
	if _, cek := data.ProAUDIO[to]; cek {
		delete(data.ProAUDIO, to)
	}
}

func ProAUDIOOn(to string) {
	if _, cek := data.ProAUDIO[to]; !cek {
		data.ProAUDIO[to] = true
	}
}
func ProIMAGEOff(to string) {
	if _, cek := data.ProIMAGE[to]; cek {
		delete(data.ProIMAGE, to)
	}
}

func ProIMAGEOn(to string) {
	if _, cek := data.ProIMAGE[to]; !cek {
		data.ProIMAGE[to] = true
	}
}
func ProReadKickOn(to string) {
	if _, cek := data.ProReadKick[to]; !cek {
		data.ProReadKick[to] = true
	}
}
func ProReadKickOff(to string) {
	if _, cek := data.ProReadKick[to]; cek {
		delete(data.ProReadKick, to)
	}
}
func fullAccess2(target string) bool {
	Menej := []string{}
	Menej = append(Menej, Maker...)
	Menej = append(Menej, data.Owner...)
	Menej = append(Menej, data.Admin...)
	Menej = append(Menej, data.Staff...)
	looper := len(Menej)
	for i := 0; i < looper; i++ {
		if target == Menej[i] {
			return true
		}
	}
	return false
}
func fullAccess(target string) bool {
	Menej := []string{}
	Menej = append(Menej, data.Squad...)
	Menej = append(Menej, Maker...)
	Menej = append(Menej, data.Owner...)
	Menej = append(Menej, data.Admin...)
	Menej = append(Menej, data.Staff...)
	looper := len(Menej)
	for i := 0; i < looper; i++ {
		if target == Menej[i] {
			return true
		}
	}
	return false
}
func fullAccessd(to string, target string) bool {
	Menej := []string{}
	Menej = append(Menej, data.Squad...)
	Menej = append(Menej, Maker...)
	Menej = append(Menej, data.Owner...)
	Menej = append(Menej, data.Admin...)
	Menej = append(Menej, data.Staff...)
	looper := len(Menej)
	for i := 0; i < looper; i++ {
		if target == Menej[i] {
			return true
		}
	}
	return false
}

// func fullAccess(target string) bool {
// 	Menej := []string{}
// 	Menej = append(Menej, Maker...)
// 	Menej = append(Menej, data.Squad...)
// 	Menej = append(Menej, data.Owner...)
// 	Menej = append(Menej, data.Admin...)
// 	Menej = append(Menej, data.Staff...)
// 	looper := len(Menej)
// 	for i := 0; i < looper; i++ {
// 		if target == Menej[i] {
// 			return true
// 		}
// 	}
// 	return false
// }
// func fullAccess(to string) bool {
// 	Menej := []string{}
// 	Menej = append(Menej, Maker...)
// 	Menej = append(Menej, data.Squad...)
// 	Menej = append(Menej, data.Owner...)
// 	Menej = append(Menej, data.Admin...)
// 	Menej = append(Menej, data.Staff...)
// 	looper := len(Menej)
// 	for i := 0; i < looper; i++ {
// 		if to == Menej[i] {
// 			return true
// 		}
// 	}
// 	return false
// }
// func fullAccess(target string) bool {
// 	Menej := []string{}
// 	Menej = append(Menej, Maker...)
// 	Menej = append(Menej, data.Squad...)
// 	looper := len(Menej)
// 	for i := 0; i < looper; i++ {
// 		if target == Menej[i] {
// 			return true
// 		}
// 	}
// 	return false
// }

// func fullAccess(target string) bool {
// 	Menej := []string{}
// 	Menej = append(Menej, Maker...)
// 	Menej = append(Menej, data.Squad...)
// 	Menej = append(Menej, data.Owner...)
// 	looper := len(Menej)
// 	for i := 0; i < looper; i++ {
// 		if target == Menej[i] {
// 			return true
// 		}
// 	}
// 	return false
// }

// func fullAccess(target string) bool {
// 	Menej := []string{}
// 	Menej = append(Menej, Maker...)
// 	Menej = append(Menej, data.Squad...)
// 	Menej = append(Menej, data.Owner...)
// 	Menej = append(Menej, data.Admin...)
// 	looper := len(Menej)
// 	for i := 0; i < looper; i++ {
// 		if target == Menej[i] {
// 			return true
// 		}
// 	}
// 	return false
// }

// func fullAccess(target string) bool {
// 	Menej := []string{}
// 	Menej = append(Menej, Maker...)
// 	Menej = append(Menej, data.Squad...)
// 	Menej = append(Menej, data.Owner...)
// 	Menej = append(Menej, data.Admin...)
// 	Menej = append(Menej, data.Staff...)
// 	looper := len(Menej)
// 	for i := 0; i < looper; i++ {
// 		if target == Menej[i] {
// 			return true
// 		}
// 	}
// 	return false
// }

func SmartKick(cl *oop.Account, op1 string, op2 string) {
	if _, ok := TimeJoin[op1]; ok && time.Since(TimeJoin[op1]) < 50*time.Millisecond {
		if TempJoin != "0" {
			go Kick(cl, op1, TempJoin)
			TempJoin = "0"
		}
		go Kick(cl, op1, op2)
	} else {
		TempJoin = op2
	}
	TimeJoin[op1] = time.Now()
}

func SetLimit(mid string) {
	/*for x := range data.StayGroup {
		if oop.Contains(data.StayGroup[x], mid) {
			data.StayGroup[x] = oop.Remove(data.StayGroup[x], mid)
		}
	}
	if oop.Contains(data.Squad, mid) {
		data.Squad = oop.Remove(data.Squad, mid)
	}
	if !oop.Contains(data.Osquad, mid) {
		data.Osquad = append(data.Osquad, mid)
	}*/
	if !data.LimitStatus[mid] {
		data.LimitStatus[mid] = true
		now := time.Now()
		timeDate := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 100, time.Local)
		add := 1 * time.Hour
		new := timeDate.Add(add)
		data.LimitTime[mid] = new
	}
	SaveData()
}

func SetNormal(mid string) {
	if !oop.Contains(data.Squad, mid) {
		data.Squad = append(data.Squad, mid)
	}
	if oop.Contains(data.Osquad, mid) {
		data.Osquad = oop.Remove(data.Osquad, mid)
	}
	if data.LimitStatus[mid] {
		data.LimitStatus[mid] = false
		data.LimitTime[mid] = time.Now()
	}
	SaveData()
}

func CheckAcceptChatByTicket(cl *oop.Account, ticketID string) {
	link, err := cl.FindChatByTicket(ticketID)
	if err != nil {
		return
	}
	if link == nil {
		return
	}
	chatMids, _ := cl.GetAllChatMids(true, false)
	if chatMids == nil {
		return
	}
	if !oop.Contains(chatMids.MemberChatMids, link.Chat.ChatMid) {
		err := cl.AcceptChatInvitationByTicket(link.Chat.ChatMid, ticketID)
		if err != nil {
			if strings.Contains(fmt.Sprintf("%s", err), "request blocked") {
				SetLimit(cl.Mid)
			}
		}
	}
}

func limitDuration(d time.Duration) string {
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	return fmt.Sprintf("%02dชั่วโมง %02dนาที %02dวินาที", h%24, m, s)
}

func SmartAjs(cl *oop.Account, op *linethrift.Operation) {
	op1 := op.Param1
	op2 := op.Param2
	op3 := op.Param3
	if !oop.Contains(data.Squad, op3) {
		return
	}
	time.Sleep(1 * time.Millisecond)
	runtime.GOMAXPROCS(2)
	go Ban(op2)
	if len(BotKor) >= len(data.StayGroup[op1])-1 {
		tx := fmt.Sprintf("%v00", cl.Count)
		tm, _ := strconv.Atoi(tx)
		time.Sleep(time.Duration(tm) * time.Millisecond)
		time.Sleep(1 * time.Millisecond)
		data.StayPending[op1] = oop.Remove(data.StayPending[op1], cl.Mid)
		Multy = false
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			go cl.AcceptChatInvitation(op1)
			go inviteAllBots(cl, op1)
			go KickBl(cl, op1)
		}()
		wg.Wait()

	} else {
		return
	}
}

func KillMode(cl *oop.Account, to string, target string) map[string][]string {
	chat, _ := cl.GetChats([]string{to}, true, true)
	go Ban(target)
	Target := make(map[string][]string)
	if chat != nil {
		members := chat.Chats[0].Extra.GroupExtra.MemberMids
		invitee := chat.Chats[0].Extra.GroupExtra.InviteeMids
		if _, cek := members[target]; cek {
			timeJoin := strconv.FormatInt(members[target], 10)[0:8]
			for t := range members {
				Filter := strconv.FormatInt(members[t], 10)[0:8]
				if Filter == timeJoin {
					if !fullAccess(t) {
						go Ban(t)
						Target["targetMember"] = append(Target["targetMember"], t)
					}
				}
			}
			for p := range invitee {
				Filter := strconv.FormatInt(invitee[p], 10)[0:8]
				if Filter == timeJoin {
					if !fullAccess(p) {
						go Ban(p)
						Target["targetInvitee"] = append(Target["targetInvitee"], p)
					}
				}
			}
		}
	}
	return Target
}

func AutoTimeSet(group string) bool {
	if _, ok := TimeSet[group]; ok && time.Since(TimeSet[group]) < 500*time.Millisecond {
		return false
	}
	return true
}
func AutoTimeSets(group string, to string) bool {
	if _, ok := TimeSet[group]; ok {
		return false
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
func ClientAcc(to string, op3 string) {
	for x := range Botlist {
		if Botlist[x].Mid == op3 {
			go Botlist[x].AcceptChatInvitationByTicket(to, ModTicket)
			go KickBl(Botlist[x], to)
			break
		}
	}
}

func perBots(cl *oop.Account) {
	runtime.GOMAXPROCS(2)
	Botlist = append(Botlist, cl)
	allgrup, _ := cl.GetAllChatMids(true, false)
	for g := range allgrup.MemberChatMids {
		putSquad(cl, allgrup.MemberChatMids[g])
	}
	if !oop.Contains(data.Squad, cl.Mid) {
		if !oop.Contains(data.Osquad, cl.Mid) {
			data.Squad = append(data.Squad, cl.Mid)
		}
	}
	if _, cek := data.LimitStatus[cl.Mid]; !cek {
		data.LimitStatus[cl.Mid] = false
		data.LimitTime[cl.Mid] = time.Now()
	}
	kickd := ""
	cancld := ""
	invtd := false
	  
	for {
		ops, err := cl.FetchOps()
		if err != nil {
			if strings.Contains(fmt.Sprintf("%s", err), "suspended") {
				data.Squad = oop.Remove(data.Squad, cl.Mid)
				for to := range data.StayGroup {
					if oop.Contains(data.StayGroup[to], cl.Mid) {
						data.StayGroup[to] = oop.Remove(data.StayGroup[to], cl.Mid)
					}
				}
				for to := range data.StayPending {
					if oop.Contains(data.StayPending[to], cl.Mid) {
						data.StayPending[to] = oop.Remove(data.StayPending[to], cl.Mid)
					}
				}
				if !oop.Contains(Freeze, cl.Mid) {
					Freeze = append(Freeze, cl.Mid)
				}
				fmt.Println("TOKEN FREZEE")
				time.Sleep(3600 * time.Second)
				file, err := osext.Executable()
				if err != nil {
					fmt.Println("Reboot", err)
				}
				err = syscall.Exec(file, os.Args, os.Environ())
				if err != nil {
					fmt.Println("Reboot", err)
				}
			}
			continue
		}
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
						case 124:
							if !sleepmode {
								// fmt.Println("เชิญเข้าสู่การแชท 124")
								op1, op2, op3, ctime := op.Param1, op.Param2, strings.Split(op.Param3, "\x1e"), op.CreatedTime
								if getAccessForCancel(cl, op2, op3) {
									CclList(cl, op1, op3)
									fmt.Println("getAccessForCancel")
									//go cl.CancelChatInvitation(cl, op1, op2)
								} else if oop.Contains(op3, cl.Mid) && !oop.Contains(data.StayPending[op1], cl.Mid) {
									var wg sync.WaitGroup
									fmt.Println("WaitGroup")
									wg.Add(1)
									go func(op1 string) {
										defer wg.Done()
										go cl.AcceptChatInvitation(op1)
										go KickBl(cl, op1)
									}(op1)
									wg.Wait()
								} else if fullAccess(op2) {
									continue
								} else if oop.Contains(data.Ban, op2) || oop.CheckEqual(data.Ban, op3) {
									if getWarAccess(cl, ctime, op1, "", cl.Mid, true) {
										BanAll(op2, op3)
									}
								} else if _, cek := data.ProInvite[op1]; cek {
									if getWarAccess(cl, ctime, op1, "", cl.Mid, false) {
										go cl.DeleteOtherFromChat(op1, []string{op2})
										go func() { CclList(cl, op1, op3) }()
										go BanAll(op2, op3)
									}
								} else if kickban == true {
									go Ban(op2)
									WarTime[op1] = time.Now()
								}
							}
						case 123:
							if !sleepmode {
								// ÷continue
								op1 := op.Param1
								op3 := op.Param3
								if AccessWarTime(op1) {
									go KickBl(cl, op1)
									go InviteMem(cl, op1, op3)
								}
							}
							//Kicked
						case 133:
							if !sleepmode {
								// continue
								op1, op2, op3, ctime := op.Param1, op.Param2, op.Param3, op.CreatedTime
								if fullAccess(op2) {
									continue
								} else if op3 == cl.Mid {
									WarTime[op1] = time.Now()
									Ban(op2)
									if Multy {
										ModJoin(cl, op1, op2)

									}
									continue
								} else if oop.Contains(data.StayPending[op1], cl.Mid) {
									if getAccessAjs(ctime) {
										go SmartAjs(cl, op)
									}
								} else if AccessWarTime(op1) {
									if getWarAccess(cl, ctime, op1, op3, cl.Mid, false) {
										if oop.Contains(data.Squad, op3) {
											Ban(op2)
											var wg sync.WaitGroup
											wg.Add(1)
											go func(op3 string) {
												defer wg.Done()
												if op2 != kickd {
													kickd = op2
													invtd = true

													if KillMod {
														res := KillMode(cl, op1, op2)
														go KickAndCancelByList(cl, op1, res["targetMember"], res["targetInvitee"])
													} else {
														//time.Sleep(1000 * time.Millisecond)
														KickBl(cl, op1)

													}
												}
												go cl.DeleteOtherFromChat(op1, []string{op2})
												go cl.FindAndAddContactsByMid(op3)
												go cl.InviteIntoChat(op1, []string{op3})
												if Multy {
													go InvQr(cl, op1)
												}
											}(op3)
											wg.Wait()
										}
									}
								} else if oop.Contains(data.Squad, op3) {
									if getWarAccess(cl, ctime, op1, op3, cl.Mid, false) {
										go Ban(op2)
										go cl.DeleteOtherFromChat(op1, []string{op2})
										if Multy {
											go cl.UpdateChatQr(op1, false)
										}
										go cl.InviteIntoChat(op1, []string{op3})
										WarTime[op1] = time.Now()
									}
								} else if _, cek := data.ProKick[op1]; cek || fullAccess(op3) {
									if getWarAccess(cl, ctime, op1, op3, cl.Mid, false) {
										res := KillMode(cl, op1, op2)
										go KickAndCancelByList(cl, op1, res["targetMember"], res["targetInvitee"])
										go cl.FindAndAddContactsByMid(op3)
										go cl.InviteIntoChat(op1, []string{op3})
										if Multy {
											InviteMem(cl, op1, op3)
										}
									}
								}
							}
							//    client kicked
						case 132:
							if !sleepmode {
								op1 := op.Param1
								if invtd {
									invtd = false
									var wg sync.WaitGroup
									wg.Add(1)
									go func(op1 string) {
										defer wg.Done()
										go inviteSquad(cl, op1)
										if Multy {
											go InvQr(cl, op1)
										}
									}(op1)
								}
							}
						case 55:
							if !sleepmode {
								op1, op2 := op.Param1, op.Param2
								if oop.Contains(data.Ban, op2) {
									go cl.DeleteOtherFromChat(op1, []string{op2})
									//	cl.SendMessage(op1, "ไม่อนุญาติบัญชีดำอ่าน ‶⍵″")
									WarTime[op1] = time.Now()
								}else if _, cek := data.ProReadKick[op1]; cek {
									if !fullAccess(op2) {
										go cl.DeleteOtherFromChat(op1, []string{op2})
									}
								}
							}
							//Join
						case 130:
							if !sleepmode {
								// continue
								op1, op2, op3 := op.Param1, op.Param2, op.Param3
								if fullAccess(op2) {
									continue
								} else if kickban == true {
									go Ban(op2)
									WarTime[op1] = time.Now()
								} else if _, cek := data.ProJoin[op1]; cek {
									go Ban(op2)
									go cl.DeleteOtherFromChat(op1, []string{op2})
									WarTime[op1] = time.Now()
								} else if oop.Contains(data.Ban, op2) || fullAccess(op3) {
									go cl.DeleteOtherFromChat(op1, []string{op2})
									WarTime[op1] = time.Now()
								}
							}
							//Cancel
						case 126:
							if !sleepmode {
								// continue
								op1, op2, op3, ctime := op.Param1, op.Param2, op.Param3, op.CreatedTime

								if fullAccess(op2) {
									continue
								} else if op3 == cl.Mid {
									if Multy {
										ModJoin(cl, op1, op2)

									}
									Ban(op2)
									WarTime[op1] = time.Now()
								} else if AccessWarTime(op1) {
									if getWarAccess(cl, ctime, op1, op3, cl.Mid, false) {
										if oop.Contains(data.Squad, op3) {
											go Ban(op2)
											if op2 != cancld {
												cancld = op2
												go KickBl(cl, op1)

											}
											if Multy {
												go InvQr(cl, op1)

											}
											var wg sync.WaitGroup
											wg.Add(1)
											go func(op1 string) {
												defer wg.Done()
												go cl.InviteIntoChat(op1, []string{op3})
											}(op1)
											wg.Wait()
											WarTime[op1] = time.Now()
										}
									}
								} else if oop.Contains(data.Squad, op3) {
									if getWarAccess(cl, ctime, op1, op3, cl.Mid, false) {
										go cl.DeleteOtherFromChat(op1, []string{op2})
										go inviteSquad(cl, op1)
										go Ban(op2)
										WarTime[op1] = time.Now()
									}
								} else if _, cek := data.ProCancel[op1]; cek || fullAccess(op3) {
									if getWarAccess(cl, ctime, op1, op3, cl.Mid, false) {
										go cl.DeleteOtherFromChat(op1, []string{op2})
										go cl.FindAndAddContactsByMid(op3)
										go InviteMem(cl, op1, op3)
										go Ban(op2)
									}
								}
							}

							//Qr
						case 122:
							if !sleepmode {
								// continue
								op1, op2, ctime := op.Param1, op.Param2, op.CreatedTime
								if fullAccess(op2) {
									continue
								} else if AccessWarTime(op1) {
									go Ban(op2)
									if _, cek := BotKor[cl.Mid]; cek && Multy {
										go cl.UpdateChatQr(op1, false)
										go cl.DeleteOtherFromChat(op1, []string{op2})
									}
								} else if _, cek := data.ProKick[op1]; cek || oop.Contains(data.Ban, op2) {
									if getWarAccess(cl, ctime, op1, "", cl.Mid, false) {
										go cl.DeleteOtherFromChat(op1, []string{op2})
										chat, _ := cl.GetChats([]string{op1}, false, false)
										if chat != nil {
											if chat.Chats[0].Extra.GroupExtra.PreventedJoinByTicket == false {
												go cl.UpdateChatQr(op1, true)
											}
										}
									}
								}
							}

						case 129:
							if !sleepmode {
								// continue
								op1 := op.Param1
								if AccessWarTime(op1) && !Multy {
									go inviteSquad(cl, op1)
								}
							}

						//Notif Add
						case 5:
							if !sleepmode {
								// continue
								adders := op.Param1
								if IsBanned(adders) {
									return
								}
								allgrup, _ := cl.GetGroupIdsJoined()
								CanAdd := false
								for _, v := range allgrup {
									if AutoTimeSets(v, adders) && !fullAccessd(v, adders) && notiFadd == true {
										CanAdd = true
										anu := cl.DeleteOtherFromChat(v, []string{adders})
										if anu != nil {
											break
										}
									}
									if CanAdd {
										appendBl(adders)
									}
								}
							}
						case 26:
							cl.Rev = -1
							ctime := op.CreatedTime
							msg := op.Message
							text := op.Message.Text
							sender := msg.From_
							var to = msg.To
							// fmt.Println([]*.GetChunks)
							// fmt.Println(op)
							// fmt.Println("++++++++++++++++")
							// cl.SendMessage(msg.To, "❌กันลิ้งค์มิจฉาชีพ❌")
							var pesan = strings.ToLower(text)
							if (op.Message.ContentType).String() == "NONE" {
								if _, cek := data.ProLINK[to]; cek {
									if strings.Contains(pesan, "http") || strings.Contains(pesan, "lin") {
										if getAccess(ctime, cl.Mid) {
											if !fullAccess(sender) {
												cl.DeleteOtherFromChat(to, []string{sender})
												appendBl(sender)
												cl.SendMessage(msg.To, "❌กันลิ้งค์มิจฉาชีพ❌")
											}
										}
									}
								}
							}
							if msg.ContentType == 0 {
								// kill คนส่งขอความ
								if _, cek := data.ProKillMsg[to]; cek && !sleepmode {
									if getAccess(ctime, cl.Mid) {
										if !fullAccess(sender) {
											cl.DeleteOtherFromChat(to, []string{sender})
											appendBl(sender)
											// cl.SendMessage(msg.To, "❌กันลิ้งค์มิจฉาชีพ❌")
										}
									}
								}
								Msg := string(msg.Text)

								if !fullAccess2(sender) {
									continue
								}

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
									// fmt.Println(txt)
									// fmt.Println("++++++++++++++++", cl.Mid)
									if txt == "กันหมด เปิด" {
										if getAccess(ctime, cl.Mid) {
											Promax(to)
											SaveData()
											putSquad(cl, to)
											cl.SendMessage(to, "กันหมด เปิดสำเร็จ")

										}
									} else if txt == "กันหมด ปิด" {
										if getAccess(ctime, cl.Mid) {
											Pronull(to)
											SaveData()
											putSquad(cl, to)
											cl.SendMessage(to, "กันหมด ปิดสำเร็จ")

										}
									} else if txt == "กันอ่าน ปิด" {
										if getAccess(ctime, cl.Mid) {
											ProReadKickOff(to)
											SaveData()
											cl.SendMessage(to, "กันอ่าน ปิดสำเร็จ")

										}
									}  else if txt == "กันอ่าน เปิด" {
										if getAccess(ctime, cl.Mid) {
											ProReadKickOn(to)
											SaveData()
											cl.SendMessage(to, "กันอ่าน เปิดสำเร็จ")

										}
									} else if txt == "แทค" {
										if getAccess(ctime, cl.Mid) { 
											chat, _ := cl.GetChats([]string{to}, true, true)
											if chat != nil {
												members := chat.Chats[0].Extra.GroupExtra.MemberMids 
												num := 1
												for b := range members {
													if !fullAccess(b){
														tx := fmt.Sprintf("%v. @!", num)
														num += 1 
														cl.SendMention(to, tx, []string{b}) 
													}
												}
											}
										}
									} else if strings.HasPrefix(txt, "เพิ่มเพื่อน ") {

										result := strings.Split((text), " ")
										// fmt.Println(result)
										for v := range result {
											if v > 0 {
												_, err := cl.FindAndAddContactsByMid(result[v])
												if err != nil {
													// fmt.Println(err)​​
													if getAccess(ctime, cl.Mid) {
														putSquad(cl, to)
														cl.SendMessage(to, "มีเป็นเพื่อนแล้ว")
														// break
													}
												}
											}
										}
										cl.SendMessage(to, "ok ..")

									} else if txt == "mymid" {
										if getAccess(ctime, cl.Mid) {

											cl.SendMessage(to, sender)

										}
									} else if strings.HasPrefix(txt, "ออกกลุ่ม ") {
										result := strings.Split((text), " ")
										// fmt.Println(result)
										for v := range result {
											if v == 1 {
												gc := GroupList[v-1]
												if len(gc) > 5 {
													cl.SendMessage(gc, " bye bye....")
													cl.DeleteSelfFromChat(gc)
													if getAccess(ctime, cl.Mid) {
														cl.SendMessage(to, " ออกกลุ่ม ok")
													}

												}

											}
										}
									} else if strings.HasPrefix(txt, "addbot ") {
										if getAccess(ctime, cl.Mid) {
											result := strings.Split((text), " ")
											fileName := fmt.Sprintf("token.txt")
											fileBytes, err := ioutil.ReadFile(fileName)
											if err != nil {
												fmt.Println(err)
												os.Exit(1)
											}
											Token := "" + string(fileBytes)
											//  index, _ := strconv.Atoi(result[1])
											Token += result[1] + ","
											ioutil.WriteFile(toeknPath, []byte(Token), 0644)
											if getAccess(ctime, cl.Mid) {
												cl.SendMessage(to, " เพิ่มโทนเค่นสำเร็จ รีบูต  server ก่อนใช้งาน")
											}

										}
									} else if txt == "newtoken" {
										if getAccess(ctime, cl.Mid) {
											// fileName := fmt.Sprintf("token.txt")
											ioutil.WriteFile(toeknPath, []byte(""), 0644)
											cl.SendMessage(to, "set null token ok")

										}
									} else if strings.HasPrefix(txt, "สมาชิกกลุ่ม ") {
										if getAccess(ctime, cl.Mid) {
											result := strings.Split((text), " ")
											index, _ := strconv.Atoi(result[1])
											// cl.SendMention(to, tx, bots)
											gc := GroupList[index-1]
											chat, _ := cl.GetChats([]string{gc}, true, true)
											data.Gmember  = []string{}
											if chat != nil {
												members := chat.Chats[0].Extra.GroupExtra.MemberMids
												// name := chat.Chats[0].ChatName
												// tx := "รายชื่อ\n"
												num := 1
												for b := range members {
													tx := fmt.Sprintf("%v. @!", num)
													num += 1
													cl.SendMention(to, tx, []string{b})
													data.Gmember = append(data.Gmember, b)
													// time.Sleep(0.7 * time.Second) 
													time.Sleep(100 * time.Millisecond)
												}
												tx := fmt.Sprintf("  จำนวน :%v ", len(data.Gmember))
												cl.SendMessage(to, tx)
												SaveData()
											}
										}
									} else if strings.HasPrefix(txt, "addban ") {
										if getAccess(ctime, cl.Mid) {
											result := strings.Split((text), " ")
											index, _ := strconv.Atoi(result[1])
											cl.SendMessage(to, data.Gmember[index-1])
											if !oop.Contains(data.Ban, data.Gmember[index-1]) {
												data.Ban = append(data.Ban, data.Gmember[index-1])
												SaveData()
											cl.SendMessage(to, "เพิ่มดำเรียบร้อย !.")
											}
										}
									} else if txt ==  "memberlen" {
										if getAccess(ctime, cl.Mid) {
											tx := fmt.Sprintf("  จำนวน :%v ", len(data.Gmember))
											cl.SendMessage(to, tx)
											}
											
											
									} else if txt == "help" {
										if getAccess(ctime, cl.Mid) {
											tx := "┏เมนูคำสั่งบอท━━\n"
											tx += "┃-help\n"
											tx += "┃-help2(ดูคำสั่งป้องกัน)\n"
											tx += "┃━━Admins━━\n"
											tx += "┃-เชคบัค\n"
											tx += "┃-แทค\n"
											tx += "┃-ออน\n"
											tx += "┃-เพิ่มสตาฟ\n"
											tx += "┃-ลบสตาฟ\n"
											tx += "┃-ค่ะ @เตะคน\n"
											tx += "┃-ป้องกัน * ดูป้องกัน\n"
											tx += "┃-ไรคะ @เตะคน\n"
											tx += "┃-. *เชคบอท\n"
											tx += "┃-เชคบอท *เชคบอท\n"
											tx += "┃-นับบอท *นับบอท\n"
											tx += "┃-count *นับบอท\n"
											tx += "┃-เชคดำ\n"
											tx += "┃-ล้างดำ\n"
											tx += "┃-ยัดดำ เปิด/ปิด\n"
											tx += "┃-ลบดำ เปิด/ปิด\n"
											tx += "┃-เพิ่มดำ @\n"
											tx += "┃-addban (ลำดับจากกลุ่ม)\n"
											tx += "┃-ลบดำ @\n"
											tx += "┃-กันหมด เปิด\n"
											tx += "┃-กันหมด ปิด\n"
											tx += "┃-เปิดลิ้ง\n"
											tx += "┃-ปิดลิ้ง\n"
											tx += "┃-แอดเพื่อนบอท\n"
											tx += "┃-เชคบัค\n"
											tx += "┃-เชคกัน\n"
											tx += "┃-เชคแอดมิน\n"
											tx += "┃-เชคเพื่อน\n"
											tx += "┃-ยึด\n"
											tx += "┃-ดึง @\n"
											tx += "┃-บัคออก\n"
											tx += "┃-join (เชิญแบบ ลิ้งค์)\n"
											tx += "┃-join2 (เชิญแบบ เพิ่มเข้ากลุ่ม )\n"
											tx += "┃-here\n"
											tx += "┃-stay\n"
											tx += "┃-setanti\n"
											tx += "┃-bye\n"
											tx += "┃-rot\n"
											tx += "┃-bm\n"
											tx += "┃-บัญชีแอดมิน\n"
											tx += "┃-gurl\n"
											tx += "┃-เชิญห้อง @\n"
											tx += "┃-เพิ่มแอดมิน\n"
											tx += "┃-ลบแอดมิน\n"
											tx += "┃-กลุ่ม\n"
											tx += "┃-ลิ้งกลุ่ม (เลขห้อง)\n"
											tx += "┃-อัพรูป\n"
											tx += "┃-อัพรูปวีดีโอ\n"
											tx += "┃-อัพชื่อ\n"
											tx += "┃-อัพตัส\n"
											tx += "┃-goto (ตามด้วยลิ้ง)\n"
											tx += "┃-มุดลิ้ง (ตามด้วยลิ้ง)\n"
											tx += "┃-รายชื่อดำ (ดึงไอดีดำทั้งหมด)\n"
											tx += "┃-เพิ่มรายชื่อดำ (เพิ่มไอดีดำทั้งหมด)\n"
											tx += "┃-bypass\n"
											tx += "┃-ac\n"
											// tx += "┃-add admin\n"
											// tx += "┃-del admin\n"
											tx += "┃-เพิ่มแอดใหญ่\n"
											tx += "┃-ลบแอดใหญ่\n"
											tx += "┃-แอดเบอร์\n"
											tx += "┃-app\n"
											tx += "┃-rest\n"
											tx += "┃-newtoken (ล้างโทเค่นออกจากไฟล์)\n"
											tx += "┃-addbot (ใส่โทเค่น)\n"
											tx += "┃-เพิ่มบอท (@)\n"
											tx += "┃-ออกทุกกลุ่ม\n"
											tx += "┃-limiter kick/join\n"
											tx += "┃-fix\n"
											cl.SendMessage(msg.To, tx)
										}
									} else if txt == "help2" {
										if getAccess(ctime, cl.Mid) {
											tx := "┏คำสั่งป้องกันกลุ่ม\n"
											tx += "┃-กันหมด เปิด/ปิด\n"
											tx += "┃-กันแอด เปิด/ปิด\n"
											tx += "┃-กันส่งข้อความ เปิด/ปิด\n"
											tx += "┃-เตะดำ เปิด/ปิด\n"
											tx += "┃-กันวางลิ้ง เปิด/ปิด\n"
											tx += "┃-กันเฟค เปิด/ปิด\n"
											tx += "┃-กันอัลบั้ม เปิด/ปิด\n"
											tx += "┃-กันสติ๊กเกอร์ เปิด/ปิด\n"
											tx += "┃-กันโทรกลุ่ม เปิด/ปิด\n"
											tx += "┃-กันส่งไฟล์ เปิด/ปิด\n"
											tx += "┃-กันโพส เปิด/ปิด\n"
											tx += "┃-กันส่งวีดีโอ เปิด/ปิด\n"
											tx += "┃-กันส่งคลิปเสียง เปิด/ปิด\n"
											tx += "┃-กันส่งรูป เปิด/ปิด\n"
											tx += "┃-กันเตะ เปิด/ปิด\n"
											tx += "┃-กันเชิญ เปิด/ปิด\n"
											tx += "┃-กันคนเข้า เปิด/ปิด\n"
											tx += "┃-กันลิ้ง เปิด/ปิด\n"
											tx += "┃-sleepmode on/off\n"
											tx += "┃-killmode on/off\n"
											tx += "┃-ajs\n"
											tx += "┃-fix\n"
											tx += "┖━━🖤━━━━"
											cl.SendMessage(msg.To, tx)
										}
									} else if strings.HasPrefix(txt, "add ") {
										if getAccess(ctime, cl.Mid) {
											result := strings.Split((text), " ")
											if result[1] == "staff" {
												if fullAccess(sender) {
													PromoteStaff = true
													PromoteAdmin = false
													PromoteOwner = false
													PromoteStaff = false
													DemoteStaff = false
													DemoteAdmin = false
													DemoteOwner = false
													Scont = true
													cl.SendMessage(msg.To, "Please Send contact of prospective Staff !..")
												}
											} else if result[1] == "owner" {
												if fullAccess(sender) {
													PromoteStaff = false
													PromoteAdmin = false
													PromoteOwner = true
													PromoteStaff = false
													DemoteStaff = false
													DemoteAdmin = false
													DemoteOwner = false
													Scont = true
													cl.SendMessage(msg.To, "Please Send contact of prospective Owner !..")
												}
											} else if result[1] == "admin" {
												if fullAccess(sender) {
													PromoteStaff = false
													PromoteAdmin = true
													PromoteOwner = false
													PromoteStaff = false
													DemoteStaff = false
													DemoteAdmin = false
													DemoteOwner = false
													Scont = true
													cl.SendMessage(msg.To, "Please Send contact of prospective Admin !..")
												}
											} else if result[1] == "done" {
												if fullAccess(sender) {
													PromoteStaff = false
													PromoteAdmin = false
													PromoteOwner = false
													PromoteStaff = false
													DemoteStaff = false
													DemoteAdmin = false
													DemoteOwner = false
													Scont = false
													cl.SendMessage(msg.To, "Promote with contact mute !...")
												}
											}
										}
									} else if strings.HasPrefix(txt, "del ") {
										if getAccess(ctime, cl.Mid) {
											result := strings.Split((text), " ")
											if result[1] == "staff" {
												if fullAccess(sender) {
													PromoteStaff = false
													PromoteAdmin = false
													PromoteOwner = false
													PromoteStaff = false
													DemoteStaff = true
													DemoteAdmin = false
													DemoteOwner = false
													Scont = true
													cl.SendMessage(msg.To, "Please Send contact for delete Staff !..")
												}
											} else if result[1] == "owner" {
												if fullAccess(sender) {
													PromoteStaff = false
													PromoteAdmin = false
													PromoteOwner = false
													PromoteStaff = false
													DemoteStaff = false
													DemoteAdmin = false
													DemoteOwner = true
													Scont = true
													cl.SendMessage(msg.To, "Please Send contact for delete Owner !..")
												}
											} else if result[1] == "admin" {
												if fullAccess(sender) {
													PromoteStaff = false
													PromoteAdmin = false
													PromoteOwner = false
													PromoteStaff = false
													DemoteStaff = false
													DemoteAdmin = true
													DemoteOwner = false
													Scont = true
													cl.SendMessage(msg.To, "Please Send contact for delete Admin !..")
												}
											} else if result[1] == "done" {
												if fullAccess(sender) {
													PromoteStaff = false
													PromoteAdmin = false
													PromoteOwner = false
													PromoteStaff = false
													DemoteStaff = false
													DemoteAdmin = false
													DemoteOwner = false
													Scont = false
													cl.SendMessage(msg.To, "Demote with contact mute !...")
												}
											}
										}
									} else if txt == "." {
										// cl.SendMessage(to, "Not have banlist")
										cl.SendMention(to, "ok @!", []string{sender})
									} else if txt == "count" {
										if getAccess(ctime, cl.Mid) {
											chat, _ := cl.GetChats([]string{to}, true, false)
											if chat != nil {
												members := chat.Chats[0].Extra.GroupExtra.MemberMids
												num := 1
												for x := range Botlist {
													if _, cek := members[data.Squad[x]]; cek && oop.Uncontains(Freeze, Botlist[x].Mid) {
														Botlist[x].SendMessage(to, strconv.Itoa(num))
														num += 1
													}
												}
												SaveData()
												putSquad(cl, to)
											}
										}
									} else if txt == "เชคดำ" {
										if getAccess(ctime, cl.Mid) {
											if len(data.Ban) != 0 {
												tx := "• Banlist\n\n"
												target := []string{}
												for x := range data.Ban {
													if data.Ban[x] != "" {
														tx += fmt.Sprintf("	%v. @!\n", x+1)
														target = append(target, data.Ban[x])
													}
												}
												cl.SendMention(to, tx, target)
											} else {
												cl.SendMessage(to, "Not have banlist")
											}
										}
									} else if txt == "รายชื่อดำ" {
										if getAccess(ctime, cl.Mid) {
											if len(data.Ban) != 0 {
												tx := ""
												for x := range data.Ban {
													if data.Ban[x] != "" {
														tx += fmt.Sprintf("%v,", data.Ban[x])
													}
												}
												cl.SendMessage(to, tx)

											} else {
												cl.SendMessage(to, "Not have banlist")

											}
										}
									} else if txt == "bans" {
										if getAccess(ctime, cl.Mid) {
											limitQR := []string{}
											for x := range Botlist {
												if oop.Contains(Freeze, Botlist[x].Mid) {
													continue
												}
												res := Botlist[x].DeleteOtherFromChat(Botlist[x].Mid, []string{Botlist[x].Mid})
												if res != nil {
													if strings.Contains(res.Error(), "request blocked") {
														SetLimit(Botlist[x].Mid)
													}
												}
												SetNormal(Botlist[x].Mid)
												rev := Botlist[x].AcceptChatInvitationByTicket(to, "kntl")
												if strings.Contains(rev.Error(), "request blocked") {
													limitQR = append(limitQR, Botlist[x].Mid)
												}
											}
											limit := []string{}
											for k := range data.LimitStatus {
												if data.LimitStatus[k] {
													limit = append(limit, k)
												}
											}
											if len(limit) != 0 {
												var no = 1
												res := fmt.Sprintf("Bots get Bans %v/%v\n", len(limit), len(data.Squad))
												for _, v := range limit {
													now := time.Now()
													timeDate := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 100, time.Local)
													wkt := data.LimitTime[v]
													wktt := wkt.Sub(timeDate)
													con, _ := cl.GetContact(v)
													cl.SendContact(to, v)
													if oop.Contains(limitQR, v) {
														limitQR = oop.Remove(limitQR, v)
													}
													res += fmt.Sprintf("\n%v. %s", no, con.DisplayName)
													res += fmt.Sprintf("\n	เหลือเวลาบัค: %s", limitDuration(wktt))
													no++
												}

												if len(limitQR) != 0 {
													res += fmt.Sprintf("\n\n\nBlock AcceptTicket")
													num := 1
													for l := range limitQR {
														con, _ := cl.GetContact(limitQR[l])
														res += fmt.Sprintf("\n%v. %s", num, con.DisplayName)
														num++
													}
												}
												cl.SendMessage(to, res)
											} else {
												res := fmt.Sprintf("All bots normal.")
												if len(limitQR) != 0 {
													res += fmt.Sprintf("\n\nBlock Accept Ticket")
													num := 1
													for l := range limitQR {
														con, _ := cl.GetContact(limitQR[l])
														res += fmt.Sprintf("\n	%v. %s", num, con.DisplayName)
														num++
													}
												}
												if len(Freeze) != 0 {
													res += fmt.Sprintf("\n\nAccount Freeze")
													numm := 1
													for l := range Freeze {
														con, _ := cl.GetContact(Freeze[l])
														res += fmt.Sprintf("\n	%v. %s", numm, con.DisplayName)
														numm++
													}
												}
												cl.SendMessage(to, res)
											}
										}
									} else if strings.HasPrefix(txt, "ค่ะ ") {
										if getWarAccess(cl, ctime, to, "", cl.Mid, false) {
											sleepmode = false
											go func() { BanWithList(dataMention) }()
											var wg sync.WaitGroup
											wg.Add(len(dataMention))
											for i := 0; i < len(dataMention); i++ {
												go func(i int) {
													defer wg.Done()
													cl.DeleteOtherFromChat(to, []string{dataMention[i]})
												}(i)
											}
											wg.Wait()
										}
									} else if strings.HasPrefix(txt, "ไรคะ ") {
										if getWarAccess(cl, ctime, to, "", cl.Mid, false) {
											sleepmode = false
											go func() { BanWithList(dataMention) }()
											var wg sync.WaitGroup
											wg.Add(len(dataMention))
											for i := 0; i < len(dataMention); i++ {
												res := KillMode(cl, to, dataMention[i])

												go func(i int) {
													defer wg.Done()
													go KickAndCancelByList(cl, to, res["targetMember"], res["targetInvitee"])
												}(i)
											}
											wg.Wait()
										}
									} else if txt == "ล้างดำ" {
										if getAccess(ctime, cl.Mid) {
											oop.Clearcache()
											ModTicket = ""
											lock = ""
											BotKor = make(map[string]bool)
											TempJoin = "0"
											pala = ""
											lockMod = ""
											if AutoMulty {
												Multy = true
											}
											WarTime = make(map[string]time.Time)
											for b := range Botlist {
												Botlist[b].SpamJoin = 0
											}
											if len(data.Ban) != 0 {
												tx := "• ClearBan\n\n"
												target := []string{}
												for x := range data.Ban {
													if data.Ban[x] != "" {
														tx += fmt.Sprintf("	%v. @!\n", x+1)
														target = append(target, data.Ban[x])
													}
												}
												cl.SendMention(to, tx, target)
											} else {
												cl.SendMessage(to, "Not have banlist")
											}
											data.Ban = []string{}
											SaveData()
											putSquad(cl, to)
										}
									} else if txt == "promax" {
										if getAccess(ctime, cl.Mid) {
											Promax(to)
											SaveData()
											putSquad(cl, to)
											cl.SendMessage(to, "🆗")
										}
									} else if txt == "pronull" {
										if getAccess(ctime, cl.Mid) {
											Pronull(to)
											SaveData()
											putSquad(cl, to)
											cl.SendMessage(to, "All Protect off")
										}
									} else if txt == "เปิดลิ้ง" {
										if getAccess(ctime, cl.Mid) {
											chat, _ := cl.GetChats([]string{to}, false, false)
											if chat != nil {
												if chat.Chats[0].Extra.GroupExtra.PreventedJoinByTicket {
													cl.UpdateChatQr(to, false)
												}
												ticket, _ := cl.ReissueChatTicket(to)
												if ticket != nil {
													cl.SendMessage(to, fmt.Sprintf("https://line.me/R/ti/g/%v", ticket.TicketId))
												}
											}
										}
									} else if txt == "ปิดลิ้ง" {
										if getAccess(ctime, cl.Mid) {
											chat, _ := cl.GetChats([]string{to}, false, false)
											if chat != nil {
												if !chat.Chats[0].Extra.GroupExtra.PreventedJoinByTicket {
													cl.UpdateChatQr(to, true)
												}
											}
										}
									} else if txt == "เชคกัน" {
										if getAccess(ctime, cl.Mid) {
											tx := "┏━━ข้อมูลบอท━━━━━━━━\n"
											tx += "┃-นักพัฒนา : tanongsak695 @!\n"
											tx += fmt.Sprintf("┃-กลุ่มทั้งหมด : %v\n", len(data.StayGroup))
											tx += "┃━━การตั้งค่าการป้องกัน━━\n"
											tx += "┃-ป้องส่งข้อความ : "
											if _, cek := data.ProKillMsg[to]; cek {
												tx += "	🟢\n"
											} else {
												tx += "	🔴\n"
											}
											tx += "┃-ป้องกันวางลิ้ง : "
											if _, cek := data.ProLINK[to]; cek {
												tx += "	🟢\n"
											} else {
												tx += "	🔴\n"
											}
											tx += "┃-ป้องกันโฆษณาflex : "
											if _, cek := data.ProFLEX[to]; cek {
												tx += "	🟢\n"
											} else {
												tx += "	🔴\n"
											}
											tx += "┃-ป้องกันลบอัลบั้ม : "
											if _, cek := data.ProDelAlbum[to]; cek {
												tx += "	🟢\n"
											} else {
												tx += "	🔴\n"
											}
											tx += "┃-ป้องกันส่งสติ๊กเกอร์ : "
											if _, cek := data.ProSTICKER[to]; cek {
												tx += "	🟢\n"
											} else {
												tx += "	🔴\n"
											}
											tx += "┃-ป้องกันโทรกลุ่ม : "
											if _, cek := data.ProCALL[to]; cek {
												tx += "	🟢\n"
											} else {
												tx += "	🔴\n"
											}
											tx += "┃-ป้องกันส่งไฟล์ : "
											if _, cek := data.ProFILE[to]; cek {
												tx += "	🟢\n"
											} else {
												tx += "	🔴\n"
											}
											tx += "┃-ป้องกันแชร์โพส : "
											if _, cek := data.ProPOSTNOTIFICATION[to]; cek {
												tx += "	🟢\n"
											} else {
												tx += "	🔴\n"
											}
											tx += "┃-ป้องกันส่งวีดีโอ : "
											if _, cek := data.ProVIDEO[to]; cek {
												tx += "	🟢\n"
											} else {
												tx += "	🔴\n"
											}
											tx += "┃-ป้องกันส่งคลิปเสียง : "
											if _, cek := data.ProAUDIO[to]; cek {
												tx += "	🟢\n"
											} else {
												tx += "	🔴\n"
											}
											tx += "┃-ป้องกันส่งรูปภาพ : "
											if _, cek := data.ProIMAGE[to]; cek {
												tx += "	??\n"
											} else {
												tx += "	🔴\n"
											}
											tx += "┃-ป้องกันเตะ :  "
											if _, cek := data.ProKick[to]; cek {
												tx += "	🟢\n"
											} else {
												tx += "	🔴\n"
											}
											tx += "┃-ป้องกันยกเลิก : "
											if _, cek := data.ProCancel[to]; cek {
												tx += "	🟢\n"
											} else {
												tx += "	🔴\n"
											}
											tx += "┃-ป้องกันเชิญ :  "
											if _, cek := data.ProInvite[to]; cek {
												tx += "	🟢\n"
											} else {
												tx += "	🔴\n"
											}
											tx += "┃-ป้องกันเปิดลิ้ง : "
											if _, cek := data.ProQr[to]; cek {
												tx += "	🟢\n"
											} else {
												tx += "	🔴\n"
											}
											tx += "┃-ป้องกันคนเข้า : "
											if _, cek := data.ProJoin[to]; cek {
												tx += "	🟢\n"
											} else {
												tx += "	🔴\n"
											}
											tx += "┃━━โหมดการตั้งค่า━━━\n"
											tx += "┃-ป้องกันคนแอดบอท : "
											if notiFadd == true {
												tx += "	🟢\n"
											} else {
												tx += "	🔴\n"
											}
											tx += "┃-ᴋɪᴄᴋʙᴀɴ : "
											if kickban == true {
												tx += "	🟢\n"
											} else {
												tx += "	🔴\n"
											}
											tx += "┃-sleepmode : "
											if sleepmode {
												tx += "	🟢\n"
											} else {
												tx += "	🔴\n"
											}
											tx += "┃-ᴋɪʟʟᴍᴏᴅᴇ : "
											if KillMod {
												tx += "	ᴇɴᴀʙʟᴇ\n"
											} else {
												tx += "	ᴅɪsᴀʙʟᴇ\n"
											}
											tx += "┃-ʀᴏᴛᴀᴛɪᴏɴs : "
											if Loop {
												tx += "	ᴘʀᴇᴠᴇɴᴛᴇᴅᴊᴏɪɴ\n"
											} else {
												tx += "	ᴘᴀʀᴇɴᴛ\n"
											}
											tx += "┃-ʙᴀᴄᴋᴜᴘ : "
											if Multy {
												tx += "	ᴍᴜʟᴛʏ\n"
											} else {
												tx += "	sᴇᴍɪ\n"
											}
											tx += "┃-ᴀᴜᴛᴏᴄʟᴇᴀʀʙᴀɴ : "
											if AutoClearban {
												tx += "	⭕\n"
											} else {
												tx += "	❌\n"
											}
											tx += fmt.Sprintf("┃-ลิมิตการเตะ : %v\n", LimiterKick)
											tx += fmt.Sprintf("┃-ลิมิตการเข้าร่วม : %v\n", LimiterJoin)
											cl.SendMention(to, tx, []string{Maker[0]})
										}
									} else if txt == "join2" {
										if getAccess(ctime, cl.Mid) {
											for c := range data.Squad {
												data.StayGroup[to] = append(data.StayGroup[to], data.Squad[c])
											}
											go inviteAllBots2(cl, to)
										}
									} else if txt == "join" {
										if getAccess(ctime, cl.Mid) {
											chat, _ := cl.GetChats([]string{to}, true, false)
											if chat != nil {
												/*if chat.Chats[0].Extra.GroupExtra.PreventedJoinByTicket {
													cl.UpdateChatQr(to, false)
												}*/
												cl.UpdateChatQr(to, false)
												members := chat.Chats[0].Extra.GroupExtra.MemberMids
												ticket, _ := cl.ReissueChatTicket(to)
												if ticket != nil {
													link := fmt.Sprintf("%v", ticket.TicketId)
													for x := range Botlist {
														if cl.Mid != data.Squad[x] {
															if _, cek := members[data.Squad[x]]; !cek && oop.Uncontains(Freeze, Botlist[x].Mid) {
																err := Botlist[x].AcceptChatInvitationByTicket(to, link)
																if err != nil {
																	fmt.Println("error", err)
																}
															}
														}
													}
												}
											}
											putSquad(cl, to)
										}
									} else if strings.HasPrefix(strings.ToLower(text), "stay ") {
										if getAccess(ctime, cl.Mid) {
											result := strings.Split((text), " ")
											putSquad(cl, to)
											BotStay := data.StayGroup[to]
											Ajs := data.StayPending[to]
											call, _ := strconv.Atoi(result[1])
											grup, _ := cl.GetChats([]string{to}, false, false)
											ticket, err := cl.ReissueChatTicket(to)

											if err != nil {
												continue
											}
											link := fmt.Sprintf("%v", ticket.TicketId)
											if len(data.Squad)-len(Ajs) >= call {
												if len(BotStay) <= call {
													if grup.Chats[0].Extra.GroupExtra.PreventedJoinByTicket == true {
														errup := cl.UpdateChatQr(to, false)
														if errup != nil {
															cl.SendMessage(msg.To, "ขออภัยไม่สามารถนำบอทได้.")
															continue
														}
													}
													for c := range data.Squad {
														if oop.Uncontains(BotStay, data.Squad[c]) && oop.Uncontains(Ajs, data.Squad[c]) && oop.Uncontains(Freeze, Botlist[c].Mid) {
															if len(data.StayGroup[to]) <= call-1 {
																if KillMod {
																	time.Sleep(6000 * time.Millisecond)
																}
																err := Botlist[c].AcceptChatInvitationByTicket(to, link)
																if err != nil {
																	if strings.Contains(err.Error(), "request blocked") {
																		SetLimit(Botlist[c].Mid)

																	}
																} else {
																	data.StayGroup[to] = append(data.StayGroup[to], data.Squad[c])
																	SetNormal(Botlist[c].Mid)
																}
															}
														}
													}
													cl.UpdateChatQr(to, true)
													limit := []string{}
													for k := range data.LimitStatus {
														if data.LimitStatus[k] == true {
															limit = append(limit, k)
														}
													}
													if len(limit) != 0 {
														l := fmt.Sprintf("บอทได้รับการแบน %v/%v", len(limit), len(data.Squad))
														cl.SendMessage(msg.To, l)
													}
												} else {
													for c := range data.Squad {
														if oop.Contains(BotStay, data.Squad[c]) && oop.Uncontains(Freeze, Botlist[c].Mid) {
															if len(data.StayGroup[to])-1 >= call {
																data.StayGroup[to] = oop.Remove(data.StayGroup[to], data.Squad[c])
																Botlist[c].DeleteSelfFromChat(to)
															}
														}
													}
												}
											} else {
												tx := ""
												if len(Ajs) != 0 {
													tx += fmt.Sprintf("คุณสามารถนำมาได้เท่านั้น %v , เพราะ %v อยู่ในระหว่างดำเนินการ ", len(data.Squad)-len(Ajs), len(Ajs))
												} else {
													tx += fmt.Sprintf("คุณบอทเพียงแค่ %v", len(data.Squad))
												}
												cl.SendMessage(msg.To, tx)
											}
										}
									} else if txt == "ยก" {
										msg, _ := cl.GetRecentMessagesV2(to, 9999)
										MED := []string{}
										for _, i := range msg {
											if i.ID != "" {
												if i.From_ == cl.Mid {
													MED = append(MED, i.ID)
												}
											}
										}
										for _, itel := range MED {
											cl.UnsendMessage(itel)
										}
									} else if strings.HasPrefix(strings.ToLower(text), "ajs ") {
										if getAccess(ctime, cl.Mid) {
											result := strings.Split((text), " ")
											putSquad(cl, to)
											BotStay := data.StayGroup[to]
											Ajs := data.StayPending[to]
											call, _ := strconv.Atoi(result[1])
											if len(data.Squad) >= call {
												if len(Ajs) <= call {
													setAjs := []string{}
													for c := range data.Squad {
														if oop.Uncontains(Ajs, data.Squad[c]) && oop.Uncontains(Freeze, Botlist[c].Mid) && data.Squad[c] != cl.Mid {
															if len(data.StayPending[to]) <= call-1 {
																if call > len(data.Squad)-len(BotStay) {
																	data.StayPending[to] = append(data.StayPending[to], data.Squad[c])
																	setAjs = append(setAjs, data.Squad[c])
																	if oop.Contains(data.StayGroup[to], data.Squad[c]) {
																		data.StayGroup[to] = oop.Remove(data.StayGroup[to], data.Squad[c])
																		Botlist[c].DeleteSelfFromChat(msg.To)
																	}
																} else {
																	if oop.Uncontains(BotStay, data.Squad[c]) {
																		data.StayPending[to] = append(data.StayPending[to], data.Squad[c])
																		setAjs = append(setAjs, data.Squad[c])
																	}
																}
															}
														}
													}
													if len(setAjs) != 0 {
														err := cl.InviteIntoChat(to, setAjs)
														if err != nil {
															cl.SendMessage(to, "I got limit !..")
														}
													} else {
														cl.SendMessage(msg.To, fmt.Sprintf("%v Bots Already on Pending !..", len(data.StayPending[to])))
													}
												} else {
													for c := range data.Squad {
														if oop.Contains(Ajs, data.Squad[c]) && oop.Uncontains(Freeze, Botlist[c].Mid) {
															if len(data.StayPending[to])-1 >= call {
																if oop.Contains(data.StayPending[to], data.Squad[c]) {
																	data.StayPending[to] = oop.Remove(data.StayPending[to], data.Squad[c])
																	Botlist[c].AcceptChatInvitation(msg.To)
																}
															}
														}
													}
												}
											} else {
												tx := fmt.Sprintf("You bot just %v", len(data.Squad))
												cl.SendMessage(msg.To, tx)
											}
										}
									} else if txt == "bye" {
										if getAccess(ctime, cl.Mid) {
											continue
										}
										cl.DeleteSelfFromChat(msg.To)
									} else if txt == "out" {
										if getAccess(ctime, cl.Mid) {
											cl.DeleteSelfFromChat(msg.To)
										}
									} else if txt == "แอดเพื่อนบอท" {
										cl.SendMessage(to, "รอแปป")
										time.Sleep(time.Duration(cl.Count) * time.Second)
										time.Sleep(1000 * time.Second)
										if len(data.Squad) != 0 {
											for _, ve := range data.Squad {
												if IsFriends(cl, ve) == false {
													time.Sleep(time.Second * 1)
													_, err := cl.FindAndAddContactsByMid(ve)
													if err != nil {
														fmt.Println(err)
														if getAccess(ctime, cl.Mid) {
															putSquad(cl, to)
															cl.SendMessage(to, "มีบอทเป็นเพื่อนแล้ว")
															break
														}
													}
												}
											}
											if getAccess(ctime, cl.Mid) {
												putSquad(cl, to)
												cl.SendMessage(to, "เพิ่มเพื่อนสำเร็จ")
											}
										}
									} else if txt == "here" {
										if getAccess(ctime, cl.Mid) {
											putSquad(cl, to)
											BotStay := data.StayGroup[to]
											Ajs := data.StayPending[to]
											tx := ""
											if len(Ajs) != 0 {
												tx += fmt.Sprintf("%v/%v 💸\n%v อยู่ค้างเชิญ", len(BotStay), len(data.Squad), len(Ajs))
											} else {
												tx += fmt.Sprintf("%v/%v 💸", len(BotStay), len(data.Squad))
											}
											cl.SendMessage(msg.To, tx)
										}
									} else if txt == "เชคบัค" {
										if getAccess(ctime, cl.Mid) {
											tx := "Statust\n\n"
											for x := range Botlist {
												if oop.Contains(Freeze, Botlist[x].Mid) {
													continue
												}
												bot, _ := Botlist[x].GetProfile()
												res := Botlist[x].DeleteOtherFromChat(Botlist[x].Mid, []string{Botlist[x].Mid})
												if res != nil {
													if strings.Contains(res.Error(), "request blocked") {
														tx += fmt.Sprintf("%v. %s : 🔴บัค\n", x+1, bot.DisplayName)
														SetLimit(Botlist[x].Mid)
														continue
													}
												}
												tx += fmt.Sprintf("%v. %s : 🟢ไม่บัค\n", x+1, bot.DisplayName)
												SetNormal(Botlist[x].Mid)
											}
											cl.SendMessage(to, tx)
										}
									} else if txt == "บัคออก" {
										if getAccess(ctime, cl.Mid) {
											for x := range Botlist {
												if oop.Contains(Freeze, Botlist[x].Mid) {
													continue
												}
												res := Botlist[x].DeleteOtherFromChat(cl.Mid, []string{cl.Mid})
												if strings.Contains(res.Error(), "request blocked") {
													Botlist[x].DeleteSelfFromChat(msg.To)
												}
											}
										}
									} else if txt == "ออน" {
										if getAccess(ctime, cl.Mid) {
											d := time.Since(timeStart)
											d = d.Round(time.Second)
											h := d / time.Hour
											d -= h * time.Hour
											m := d / time.Minute
											d -= m * time.Minute
											s := d / time.Second
											cl.SendMessage(msg.To, fmt.Sprintf("เวลาออนบอท:\n%02d วัน %02d ชั่วโมง %02d นาที %02d วินาที", h/24, h%24, m, s))
										}
									} else if txt == "เชคแอดมิน" {
										if getAccess(ctime, cl.Mid) {
											team := []string{}
											tx := "• ทีมผู้ใช้\n\n"
											if len(Maker) != 0 {
												tx += "	ผู้สร้าง\n"
												for x := range Maker {
													tx += fmt.Sprintf("	%v. @!\n", x+1)
													team = append(team, Maker[x])
												}
											}
											if len(data.Owner) != 0 {
												tx += "\n	แอดมินใหญ่\n"
												for x := range data.Owner {
													tx += fmt.Sprintf("	%v. @!\n", x+1)
													team = append(team, data.Owner[x])
												}
											}
											if len(data.Admin) != 0 {
												tx += "\n	แอดมิน\n"
												for x := range data.Admin {
													tx += fmt.Sprintf("	%v. @!\n", x+1)
													team = append(team, data.Admin[x])
												}
											}
											if len(data.Staff) != 0 {
												tx += "\n	ผู้ช่วยแอดมิน\n"
												for x := range data.Staff {
													tx += fmt.Sprintf("	%v. @!\n", x+1)
													team = append(team, data.Staff[x])
												}
											}
											cl.SendMention(to, tx, team)
										}
									} else if strings.HasPrefix(strings.ToLower(text), "limiter ") {
										if getAccess(ctime, cl.Mid) {
											result := strings.Split((text), " ")
											putSquad(cl, to)
											if result[1] == "kick" {
												kick, _ := strconv.Atoi(result[2])
												LimiterKick = kick
												cl.SendMessage(to, fmt.Sprintf("Limiters Kicked : %v", LimiterKick))
											} else if result[1] == "join" {
												join, _ := strconv.Atoi(result[2])
												LimiterJoin = join
												cl.SendMessage(to, fmt.Sprintf("Limiters Joined : %v", LimiterJoin))
											}
											SaveData()
										}
									} else if strings.HasPrefix(strings.ToLower(text), "กันวางลิ้ง ") {
										if getAccess(ctime, cl.Mid) {
											result := strings.Split((text), " ")
											putSquad(cl, to)
											if result[1] == "เปิด" {
												ProLINKOn(to)
												cl.SendMessage(to, "เปิดป้องกันวางลิ้ง")
											} else if result[1] == "ปิด" {
												ProLINKOff(to)
												cl.SendMessage(to, "ปิดป้องกันวางลิ้ง")
											}
											SaveData()
										}
									} else if strings.HasPrefix(strings.ToLower(text), "กันเฟค ") {
										if getAccess(ctime, cl.Mid) {
											result := strings.Split((text), " ")
											putSquad(cl, to)
											if result[1] == "เปิด" {
												ProFLEXOn(to)
												cl.SendMessage(to, "เปิดป้องกันโฆษาflex")
											} else if result[1] == "ปิด" {
												ProFLEXOff(to)
												cl.SendMessage(to, "ปิดป้องกันโฆษาflex")
											}
											SaveData()
										}
									} else if strings.HasPrefix(strings.ToLower(text), "กันอัลบั้ม ") {
										if getAccess(ctime, cl.Mid) {
											result := strings.Split((text), " ")
											putSquad(cl, to)
											if result[1] == "เปิด" {
												DelAlbumOn(to)
												cl.SendMessage(to, "เปิดป้องกันลบอัลบั้ม")
											} else if result[1] == "ปิด" {
												DelAlbumOff(to)
												cl.SendMessage(to, "ปิดป้องกันลบอัลบั้ม")
											}
											SaveData()
										}
									} else if strings.HasPrefix(strings.ToLower(text), "กันสติ๊กเกอร์ ") {
										if getAccess(ctime, cl.Mid) {
											result := strings.Split((text), " ")
											putSquad(cl, to)
											if result[1] == "เปิด" {
												ProSTICKEROn(to)
												cl.SendMessage(to, "เปิดป้องกันส่งสติ๊กเกอร์")
											} else if result[1] == "ปิด" {
												ProSTICKEROff(to)
												cl.SendMessage(to, "ปิดป้องกันส่งสติ๊กเกอร์")
											}
											SaveData()
										}
									} else if strings.HasPrefix(strings.ToLower(text), "กันโทรกลุ่ม ") {
										if getAccess(ctime, cl.Mid) {
											result := strings.Split((text), " ")
											putSquad(cl, to)
											if result[1] == "เปิด" {
												ProCALLOn(to)
												cl.SendMessage(to, "เปิดป้องกันโทรกลุ่ม")
											} else if result[1] == "ปิด" {
												ProCALLOff(to)
												cl.SendMessage(to, "ปิดป้องกันโทรกลุ่ม")
											}
											SaveData()
										}
									} else if strings.HasPrefix(strings.ToLower(text), "กันส่งไฟล์ ") {
										if getAccess(ctime, cl.Mid) {
											result := strings.Split((text), " ")
											putSquad(cl, to)
											if result[1] == "เปิด" {
												ProFILEOn(to)
												cl.SendMessage(to, "เปิดป้องกันส่งไฟล์")
											} else if result[1] == "ปิด" {
												ProFILEOff(to)
												cl.SendMessage(to, "ปิดป้องกันส่งไฟล์")
											}
											SaveData()
										}
									} else if strings.HasPrefix(strings.ToLower(text), "กันโพส ") {
										if getAccess(ctime, cl.Mid) {
											result := strings.Split((text), " ")
											putSquad(cl, to)
											if result[1] == "เปิด" {
												ProPOSTNOTIFICATIONOn(to)
												cl.SendMessage(to, "เปิดป้องกันโพส")
											} else if result[1] == "ปิด" {
												ProPOSTNOTIFICATIONOff(to)
												cl.SendMessage(to, "ปิดป้องกันโพส")
											}
											SaveData()
										}
									} else if strings.HasPrefix(strings.ToLower(text), "กันส่งวีดีโอ ") {
										if getAccess(ctime, cl.Mid) {
											result := strings.Split((text), " ")
											putSquad(cl, to)
											if result[1] == "เปิด" {
												ProVIDEOOn(to)
												cl.SendMessage(to, "เปิดป้องกันส่งวีดีโอ")
											} else if result[1] == "ปิด" {
												ProVIDEOOff(to)
												cl.SendMessage(to, "ปิดป้องกันส่งวีดีโอ")
											}
											SaveData()
										}
									} else if strings.HasPrefix(strings.ToLower(text), "กันส่งคลิปเสียง ") {
										if getAccess(ctime, cl.Mid) {
											result := strings.Split((text), " ")
											putSquad(cl, to)
											if result[1] == "เปิด" {
												ProAUDIOOn(to)
												cl.SendMessage(to, "เปิดป้องกันส่งคลิปเสียง")
											} else if result[1] == "ปิด" {
												ProAUDIOOff(to)
												cl.SendMessage(to, "ปิดป้องกันส่งคลิปเสียง")
											}
											SaveData()
										}
									} else if strings.HasPrefix(strings.ToLower(text), "กันส่งรูป ") {
										if getAccess(ctime, cl.Mid) {
											result := strings.Split((text), " ")
											putSquad(cl, to)
											if result[1] == "เปิด" {
												ProIMAGEOn(to)
												cl.SendMessage(to, "เปิดป้องกันส่งรูป")
											} else if result[1] == "ปิด" {
												ProIMAGEOff(to)
												cl.SendMessage(to, "ปิดป้องกันส่งรูป")
											}
											SaveData()
										}
									} else if strings.HasPrefix(strings.ToLower(text), "กันเชิญ ") {
										if getAccess(ctime, cl.Mid) {
											result := strings.Split((text), " ")
											putSquad(cl, to)
											if result[1] == "เปิด" {
												ProinviteOn(to)
												cl.SendMessage(to, "🧧ป้องกันสมาชิก🧧\n❌ห้ามเชิญกันเอง❌")
											} else if result[1] == "ปิด" {
												ProinviteOff(to)
												cl.SendMessage(to, "🔓🗡️ปลดล็อค🔓🗡️\n😎ดึงกันเองได้แล้ว😎")
											}
											SaveData()
										}
									} else if strings.HasPrefix(strings.ToLower(text), "กันเตะ ") {
										if getAccess(ctime, cl.Mid) {
											result := strings.Split((text), " ")
											putSquad(cl, to)
											if result[1] == "เปิด" {
												ProkickOn(to)
												cl.SendMessage(to, "ป้องกันสมาชิกเตะเปิด")
											} else if result[1] == "ปิด" {
												ProkickOff(to)
												cl.SendMessage(to, "ป้องกันสมาชิกเตะปิด")
											}
											SaveData()
										}
									} else if strings.HasPrefix(strings.ToLower(text), "กันยกเชิญ ") {
										if getAccess(ctime, cl.Mid) {
											result := strings.Split((text), " ")
											putSquad(cl, to)
											if result[1] == "เปิด" {
												ProcancelOn(to)
												cl.SendMessage(to, "ป้องกันสมาชิกยกเชิญเปิด")
											} else if result[1] == "ปิด" {
												ProcancelOff(to)
												cl.SendMessage(to, "ป้องกันสมาชิกยกเชิญปิด")
											}
											SaveData()
										}
									} else if strings.HasPrefix(strings.ToLower(text), "กันคนเข้า ") {
										if getAccess(ctime, cl.Mid) {
											result := strings.Split((text), " ")
											putSquad(cl, to)
											if result[1] == "เปิด" {
												ProjoinOn(to)
												cl.SendMessage(to, "ป้องกันคนเข้าเปิด")
											} else if result[1] == "ปิด" {
												ProjoinOff(to)
												cl.SendMessage(to, "ป้องกันคนเข้าปิด")
											}
											SaveData()
										}
									} else if strings.HasPrefix(strings.ToLower(text), "กันลิ้ง ") {
										if getAccess(ctime, cl.Mid) {
											result := strings.Split((text), " ")
											putSquad(cl, to)
											if result[1] == "เปิด" {
												ProqrOn(to)
												cl.SendMessage(to, "เปิดป้องกันสมาชิกสมาชิกเปิดลิ้ง")
											} else if result[1] == "ปิด" {
												ProqrOff(to)
												cl.SendMessage(to, "ปิดป้องกันสมาชิกสมาชิกเปิดลิ้ง")
											}
											SaveData()
										}
									} else if strings.HasPrefix(strings.ToLower(text), "killmode ") {
										if getAccess(ctime, cl.Mid) {
											result := strings.Split((text), " ")
											putSquad(cl, to)
											if result[1] == "on" {
												KillMod = true
												cl.SendMessage(to, "Killmode enable")
											} else if result[1] == "off" {
												KillMod = false
												cl.SendMessage(to, "Killmode disable")
											}
											SaveData()
										}
									} else if strings.HasPrefix(strings.ToLower(text), "sleepmode ") {
										if getAccess(ctime, cl.Mid) {
											result := strings.Split((text), " ")
											putSquad(cl, to)
											if result[1] == "on" {
												sleepmode = true
												cl.SendMessage(to, "🟢 เปิดโหดพักผ่อน")
											} else if result[1] == "off" {
												sleepmode = false
												cl.SendMessage(to, "🔴 ปิดโหดพักผ่อน")
											}
											SaveData()
										}
									} else if strings.HasPrefix(strings.ToLower(text), "กันแอด ") {
										if getAccess(ctime, cl.Mid) {
											result := strings.Split((text), " ")
											putSquad(cl, to)
											if result[1] == "เปิด" {
												notiFadd = true
												cl.SendMessage(to, "เปิดกันแอดบอท")
											} else if result[1] == "ปิด" {
												notiFadd = false
												cl.SendMessage(to, "ปิดกันแอดบอท")
											}
											SaveData()
										}
									} else if strings.HasPrefix(strings.ToLower(text), "เตะดำ ") {
										if getAccess(ctime, cl.Mid) {
											result := strings.Split((text), " ")
											putSquad(cl, to)
											if result[1] == "เปิด" {
												kickban = true
												cl.SendMessage(to, "ᴍᴏᴅᴇ ᴋɪᴄᴋʙᴀɴ ᴏɴ")
											} else if result[1] == "ปิด" {
												kickban = false
												cl.SendMessage(to, "ᴍᴏᴅᴇ ᴋɪᴄᴋʙᴀɴ ᴏғғ")
											}
											SaveData()
										}
									} else if strings.HasPrefix(strings.ToLower(text), "กันส่งข้อความ ") {
										if getAccess(ctime, cl.Mid) {
											result := strings.Split((text), " ")
											putSquad(cl, to)
											if result[1] == "เปิด" {
												ProKillMsgOn(to)
												cl.SendMessage(to, "🟢 กันส่งข้อความ เปิดแล้ว")
											} else if result[1] == "ปิด" {
												ProKillMsgOff(to)
												cl.SendMessage(to, "🔴 กันส่งข้อความ ปิดแล้ว")
											}
											SaveData()
										}
									} else if strings.HasPrefix(txt, "เพิ่มสตาฟ ") {
										if getAccess(ctime, cl.Mid) {
											for m := range dataMention {
												if !oop.Contains(data.Staff, dataMention[m]) {
													data.Staff = append(data.Staff, dataMention[m])
												}
											}
											cl.SendMessage(to, "เพิ่มผู้ช่วยแอดมินสำเร็จ !.")
										}
									} else if strings.HasPrefix(txt, "ลบสตาฟ ") {
										if getAccess(ctime, cl.Mid) {
											for m := range dataMention {
												if oop.Contains(data.Staff, dataMention[m]) {
													data.Staff = oop.Remove(data.Staff, dataMention[m])
												}
											}
											cl.SendMessage(to, "ลบผู้ช่วยแอดมินสำเร็จ !.")
										}
									} else if strings.HasPrefix(strings.ToLower(text), "ลิ้งกลุ่ม ") {
										if getAccess(ctime, cl.Mid) {
											result := strings.Split((text), " ")
											num, _ := strconv.Atoi(result[1])
											if len(GroupList) >= num {
												for b := range Botlist {
													if oop.Contains(Freeze, Botlist[b].Mid) {
														continue
													}
													allgc, _ := Botlist[b].GetAllChatMids(true, false)
													if oop.Contains(allgc.MemberChatMids, GroupList[num-1]) {
														chat, _ := Botlist[b].GetChats([]string{GroupList[num-1]}, false, false)
														if chat != nil {
															if chat.Chats[0].Extra.GroupExtra.PreventedJoinByTicket {
																Botlist[b].UpdateChatQr(GroupList[num-1], false)
															}
														}
														ticket, _ := Botlist[b].ReissueChatTicket(GroupList[num-1])
														if ticket != nil {
															cl.SendMessage(to, fmt.Sprintf("https://line.me/R/ti/g/%v", ticket.TicketId))
														}
														break
													}
												}
											} else {
												cl.SendMessage(to, "กรุณาตรวจสอบรายชื่อกลุ่ม..")
											}
										}
									} else if strings.HasPrefix(strings.ToLower(text), "เชิญห้อง ") {
										if getAccess(ctime, cl.Mid) {
											result := strings.Split((text), " ")
											num, _ := strconv.Atoi(result[1])
											if len(GroupList) >= num {
												for b := range Botlist {
													if oop.Contains(Freeze, Botlist[b].Mid) {
														continue
													}
													allgc, _ := Botlist[b].GetAllChatMids(true, false)
													if oop.Contains(allgc.MemberChatMids, GroupList[num-1]) {
														InviteMem(Botlist[b], GroupList[num-1], sender)
														break
													}
												}
												cl.SendMessage(to, "เชิญเสร้จแล้ว !..")
											} else {
												cl.SendMessage(to, "กรุณาตรวจสอบรายชื่อกลุ่ม..")
											}
										}
									} else if strings.HasPrefix(txt, "เพิ่มแอดมิน ") {
										if getAccess(ctime, cl.Mid) {
											for m := range dataMention {
												if !oop.Contains(data.Admin, dataMention[m]) {
													data.Admin = append(data.Admin, dataMention[m])
												}
											}
											cl.SendMessage(to, "เพิ่มแอดมินเรียบร้อย !.")
										}
									} else if strings.HasPrefix(txt, "ลบแอดมิน ") {
										if getAccess(ctime, cl.Mid) {
											for m := range dataMention {
												if oop.Contains(data.Admin, dataMention[m]) {
													data.Admin = oop.Remove(data.Admin, dataMention[m])
												}
											}
											cl.SendMessage(to, "ลบแอดมินเรียบร้อย !.")
										}
									} else if strings.HasPrefix(txt, "เพิ่มดำ ") {
										if getAccess(ctime, cl.Mid) {
											for m := range dataMention {
												if !oop.Contains(data.Ban, dataMention[m]) {
													data.Ban = append(data.Ban, dataMention[m])
												}
											}
											cl.SendMessage(to, "เพิ่มดำเรียบร้อย !.")
										}
									} else if strings.HasPrefix(txt, "เพิ่มรายชื่อดำ ") {
										if getAccess(ctime, cl.Mid) {
											result := strings.Split((text), ",")
											for m := range result {
												if m > 0 {
													if !oop.Contains(data.Ban, result[m]) && len(result[m]) > 3 {
														data.Ban = append(data.Ban, result[m])
													}
												}
											}
											SaveData()
											cl.SendMessage(to, "เพิ่มดำเรียบร้อย !.")
										}
									} else if strings.HasPrefix(txt, "ลบดำ ") {
										if getAccess(ctime, cl.Mid) {
											for m := range dataMention {
												if oop.Contains(data.Ban, dataMention[m]) {
													data.Ban = oop.Remove(data.Ban, dataMention[m])
												}
											}
											cl.SendMessage(to, "ลบดำเรียบร้อย !.")
										}
									} else if txt == "rot" {
										if getAccess(ctime, cl.Mid) {
											if Loop {
												Loop = false
												LimiterJoin = 100
												cl.SendMessage(to, "😎😎")
											} else {
												Loop = true
												LimiterJoin = 100
												cl.SendMessage(to, "😎")
											}
										}
									} else if txt == "bm" {
										if getAccess(ctime, cl.Mid) {
											if Multy {
												Multy = false
												AutoMulty = false
												cl.SendMessage(to, "😎😎")
											} else {
												Multy = true
												AutoMulty = true
												cl.SendMessage(to, "😎")
											}
										}
									} else if txt == "ac" {
										if getAccess(ctime, cl.Mid) {
											if AutoClearban {
												AutoClearban = false
												cl.SendMessage(to, "Auto clear disabled")
											} else {
												AutoClearban = true
												cl.SendMessage(to, "Auto clear enabled")
											}
										}
									} else if strings.HasPrefix(txt, "มุดลิ้ง ") {
										get := strings.Split((text), " ")
										link := strings.Split((get[1]), "https://line.me/R/ti/g/")
										ticket := link[1]
										findGc, err := cl.FindChatByTicket(ticket)
										if err != nil {
											if strings.Contains(err.Error(), "request blocked") {
												cl.SendMessage(to, "เรียบร้อย🤣มีตัวบัค❌")
											}

										}
										gc := fmt.Sprintf("%v", findGc.Chat.ChatMid)
										time.Sleep(time.Duration(cl.Count) * time.Second)
										cl.AcceptChatInvitationByTicket(gc, ticket)
										if getAccess(ctime, cl.Mid) {
											cl.SendMessage(to, "🤏บอทมุดเข้าละคับ🤪")
										}
									} else if txt == "ยัดดำ เปิด" {
										if getAccess(ctime, cl.Mid) {
											PromoteBlacklist = true
											cl.SendMessage(to, "เปิดส่งคอนแทคคนที่ลงดำ")
										}
									} else if txt == "ยัดดำ ปิด" {
										if getAccess(ctime, cl.Mid) {
											PromoteBlacklist = false
											cl.SendMessage(to, "ปิดส่งคอนแทคคนที่ลงดำ")
										}
									} else if txt == "ล้างดำ เปิด" {
										if getAccess(ctime, cl.Mid) {
											delBlacklist = true
											cl.SendMessage(to, "เปิดส่งคอนแทคคนที่ลบดำ")
										}
									} else if txt == "ล้างดำ ปิด" {
										if getAccess(ctime, cl.Mid) {
											delBlacklist = false
											cl.SendMessage(to, "ปิดส่งคอนแทคคนที่ลบดำ")
										}
									} else if txt == "ยึด" {
										if getAccess(ctime, cl.Mid) {
											putSquad(cl, to)
											ByPass(cl, to)
										}
									} else if txt == "กลุ่ม" {
										if getAccess(ctime, cl.Mid) {
											data.StayGroup = map[string][]string{}
											data.StayPending = map[string][]string{}
											for b := range Botlist {
												if oop.Contains(Freeze, Botlist[b].Mid) {
													continue
												}
												allgrup, _ = Botlist[b].GetAllChatMids(true, false)
												for g := range allgrup.MemberChatMids {
													putSquad(Botlist[b], allgrup.MemberChatMids[g])
												}
											}
											tx := "Group List\n\n"
											num := 1
											GroupList = []string{}
											for g := range data.StayGroup {
												if !oop.Contains(GroupList, g) {
													GroupList = append(GroupList, g)
												}
											}

											for g := range GroupList {
												gc := GroupList[g]
												chat, _ := cl.GetChats([]string{gc}, true, true)
												if chat != nil {
													members := chat.Chats[0].Extra.GroupExtra.MemberMids
													pending := chat.Chats[0].Extra.GroupExtra.InviteeMids
													name := chat.Chats[0].ChatName
													if _, cek := data.ProKick[gc]; cek {
														tx += fmt.Sprintf("%v. %v %v/%v 🔒\n", num, name, len(members), len(pending))
													} else {
														tx += fmt.Sprintf("%v. %v %v/%v\n", num, name, len(members), len(pending))
													}
												}
												num += 1
											}
											tx += fmt.Sprintf("Total : %v Group", len(data.StayGroup))
											cl.SendMessage(to, tx)
										}
									} else if strings.HasPrefix(txt, "goto ") {
										get := strings.Split((text), " ")
										link := strings.Split((get[1]), "https://line.me/R/ti/g/")
										ticket := link[1]
										findGc, err := cl.FindChatByTicket(ticket)
										if err != nil {
											if strings.Contains(err.Error(), "request blocked") {
												cl.SendMessage(to, "Im limit !..")
											}
											continue
										}
										gc := fmt.Sprintf("%v", findGc.Chat.ChatMid)
										time.Sleep(time.Duration(cl.Count) * time.Second)
										cl.AcceptChatInvitationByTicket(gc, ticket)
										if getAccess(ctime, cl.Mid) {
											cl.SendMessage(to, "Accept Group by ticket succses !..")
										}
									} else if txt == fmt.Sprintf("%vอัพรูป", cl.Count+1) {
										updateImage[cl.Mid] = true
										cl.SendMessage(to, "โปรดส่งรูปมา !.")
									} else if txt == fmt.Sprintf("%vอัพรูปวีดีโอ", cl.Count+1) {
										updateVideo.Tipe = "cvp"
										updateVideo.Mid[cl.Mid] = true
										updateVideo.VideoStatus = true
										cl.SendMessage(to, "โปรดส่งวีดีโอมา !.")
									} else if txt == fmt.Sprintf("อัพรูป") {
										updateImage[cl.Mid] = true
										cl.SendMessage(to, "โปรดส่งรูปมา !.")
									} else if txt == fmt.Sprintf("อัพรูปวีดีโอ") {
										updateVideo.Tipe = "cvp"
										updateVideo.Mid[cl.Mid] = true
										updateVideo.VideoStatus = true
										cl.SendMessage(to, "โปรดส่งวีดีโอมา !.")
									} else if txt == fmt.Sprintf("%vอัพปก", cl.Count+1) {
										updateCover[cl.Mid] = true
										cl.SendMessage(to, "โปรดส่งรูปมา !.")
									} else if strings.HasPrefix(txt, "อัพชื่อ ") {
										get := strings.Split((text), " ")
										name := ""
										for v := range get {
											if v != 0 {
												name += fmt.Sprintf("%v ", get[v])
											}
										}
										cl.UpdateProfileAttributes(2, name)
										cl.SendMessage(to, "อัพเดทชื่อเป็น "+name)
									} else if strings.HasPrefix(txt, fmt.Sprintf("%vอัพชื่อ ", cl.Count+1)) {
										get := strings.Split((text), " ")
										println("ok")
										cl.UpdateProfileAttributes(2, get[1])
										cl.SendMessage(to, "อัพเดทชื่อเป็น "+get[1])
									} else if strings.HasPrefix(txt, fmt.Sprintf("%vอัพตัส ", cl.Count+1)) {
										get := strings.Split((text), " ")
										println("ok")
										cl.UpdateProfile(get[1], " ")
										cl.SendMessage(to, "อัพเดทสเตตัสเป็น "+get[1])
									} else if txt == "bot" {
										cl.SendMessage(to, cl.Mid)
									} else if txt == "get" {
										if getAccess(ctime, cl.Mid) {
											cl.SendMessage(msg.To, "i Get first")
											SaveData()
											putSquad(cl, to)
										}
										ticket, _ := cl.ReissueChatTicket(to)
										if ticket != nil {
											ModTicket = fmt.Sprintf("%v", ticket.TicketId)
											lock = fmt.Sprintf("%v", ctime)
											fmt.Println(ModTicket)
										}
									} else if txt == "app" {
										cl.SendMessage(msg.To, cl.LineApp+"\n"+cl.UserAgent+"\n"+cl.Host)
									} else if txt == "rest" {
										if getAccess(ctime, cl.Mid) {
											oop.Clearcache()
											cl.SendMessage(to, "รีสตาร์ทระบบ...")
											SaveData()
											putSquad(cl, to)
											file, err := osext.Executable()
											if err != nil {
												fmt.Println("Reboot", err)
											}
											err = syscall.Exec(file, os.Args, os.Environ())
											if err != nil {
												fmt.Println("Reboot", err)
											}
										}
									} else if txt == "เชคบอท" {
										fmt.Println(txt)
										fmt.Println("++++เชคบอท+++++", cl.Mid)
										fmt.Println("++++เชคบอท+++++", getAccess(ctime, cl.Mid))
										if getAccess(ctime, cl.Mid) {
											tx := "• Squad Bots\n\n"
											bots := []string{}
											num := 1
											for b := range data.Squad {
												tx += fmt.Sprintf("%v. @!\n", num)
												num += 1
												bots = append(bots, data.Squad[b])
											}
											cl.SendMention(to, tx, bots)
											fmt.Println(to, tx, bots)
										}
									} else if txt == "เชคเพื่อน" {
										nm := []string{}
										teman, _ := cl.GetAllContactIds()
										for c, v := range teman {
											res, _ := cl.GetContact(v)
											name := res.DisplayName
											c += 1
											name = fmt.Sprintf("%v. %s", c, name)
											nm = append(nm, name)
										}
										stf := "• 𝐟𝐫𝐢𝐞𝐧𝐝𝐥𝐢𝐬𝐭 •\n\n"
										str := strings.Join(nm, "\n")
										cl.SendMessage(to, stf+str)
									} else if txt == "นับบอท" {
										res, _ := cl.GetAllContactIds()
										tx := "Contact\n\n"
										for x := range res {
											get, _ := cl.GetContact(res[x])
											tx += fmt.Sprintf("%v. "+get.DisplayName+" : %v\n", x, res[x])
										}
										cl.SendMessage(to, tx)
									} else if txt == "เพิ่มบอท" {

										cl.SendMessage(to, "เพิ่มบอทok")
										res, _ := cl.GetAllContactIds()
										num := 1
										for m := range data.Squad {
											if !oop.Contains(res, data.Squad[m]) && data.Squad[m] != cl.Mid {
												time.Sleep(time.Duration(cl.Count) * time.Second)
												time.Sleep(1000 * time.Second)
												_, err := cl.FindAndAddContactsByMid(data.Squad[m])
												if err != nil {
													cl.SendMessage(to, "Limit add")
													break
												} else if num == len(data.Squad)-1 {
													cl.SendMessage(to, "Add all success..!")
												}
											} else if num == len(data.Squad)-1 {
												cl.SendMessage(to, "Already in contact..!")
											}
											num += 1
										}
									} else if strings.HasPrefix(txt, "เพิ่มแอดใหญ่ ") {
										if getAccess(ctime, cl.Mid) {
											for m := range dataMention {
												if !oop.Contains(data.Owner, dataMention[m]) {
													data.Owner = append(data.Owner, dataMention[m])
												}
											}
											cl.SendMessage(to, "Promote Owners succes !.")
										}
									} else if strings.HasPrefix(txt, "ลบแอดใหญ่ ") {
										if getAccess(ctime, cl.Mid) {
											for m := range dataMention {
												if oop.Contains(data.Owner, dataMention[m]) {
													data.Owner = oop.Remove(data.Owner, dataMention[m])
												}
											}
											cl.SendMessage(to, "Demote Owners succes !.")
										}
									} else if strings.HasPrefix(txt, "mid ") {
										if getAccess(ctime, cl.Mid) {
											for m := range dataMention {
												cl.SendMessage(to, dataMention[m])
											}
										}
									} else if txt == "ออกทุกกลุ่ม" {
										allgrup, _ := cl.GetAllChatMids(true, false)
										proGroup := 0
										leaveTo := 0
										for g := range allgrup.MemberChatMids {
											if _, cek := data.ProKick[allgrup.MemberChatMids[g]]; !cek && allgrup.MemberChatMids[g] != to {
												cl.DeleteSelfFromChat(allgrup.MemberChatMids[g])
												leaveTo += 1
											} else if allgrup.MemberChatMids[g] != to {
												proGroup += 1
											}
										}
										data.StayGroup = map[string][]string{}
										tx := fmt.Sprintf("ออกจาก %v ความสำเร็จของกลุ่ม !.\nและอยู่ใน %v กลุ่มในการป้องกัน", leaveTo, proGroup)
										allgrup, _ = cl.GetAllChatMids(true, false)
										time.Sleep(time.Duration(cl.Count) * time.Second)
										for g := range allgrup.MemberChatMids {
											putSquad(cl, allgrup.MemberChatMids[g])
										}
										SaveData()
										if getAccess(ctime, cl.Mid) {
											GroupList = []string{}
											cl.SendMessage(to, tx)
										}
									} else if txt == "fix" {
										if getAccess(ctime, cl.Mid) {
											oop.Clearcache()
											data.StayGroup = map[string][]string{}
											data.StayPending = map[string][]string{}
											data.Squad = []string{}
											data.LimitStatus = map[string]bool{}
											data.LimitTime = map[string]time.Time{}
											for x := range Botlist {
												if oop.Contains(Freeze, Botlist[x].Mid) {
													continue
												}
												_, err := Botlist[x].GetProfile()
												if err != nil {
													fmt.Println(err)
													continue
												}
												data.Squad = append(data.Squad, Botlist[x].Mid)
											}
											allgrup, _ := cl.GetAllChatMids(true, false)
											for g := range allgrup.MemberChatMids {
												putSquad(cl, allgrup.MemberChatMids[g])
											}

											SaveData()
											cl.SendMessage(to, "แจ๋วจ้า")
											time.Sleep(1 * time.Second)
											cl.SendMessage(to, "ok")
											file, err := osext.Executable()
											if err != nil {
												fmt.Println("Reboot", err)
											}
											err = syscall.Exec(file, os.Args, os.Environ())
											if err != nil {
												fmt.Println("Reboot", err)
											}
											/*
											 */
										}
									}
								}
							} else if (op.Message.ContentType).String() == "FLEX" {
								if _, cek := data.ProFLEX[to]; cek {
									if getAccess(ctime, cl.Mid) {
										if !sleepmode {
											if !fullAccess(sender) {
												cl.DeleteOtherFromChat(to, []string{sender})
												appendBl(sender)
												cl.SendMessage(to, "❌💫ห้าม💫โฆษณาflex❌")
											}
										}

									}
								}
							} else if (op.Message.ContentType).String() == "CHATEVENT" {
								if _, cek := data.ProDelAlbum[to]; cek && op.Message.ContentMetadata["LOC_KEY"] == "BD" {
									if getAccess(ctime, cl.Mid) {
										if !sleepmode {
											if !fullAccess(sender) {
												cl.DeleteOtherFromChat(to, []string{sender})
												appendBl(sender)
												cl.SendMessage(to, "🪶💫ห้าม💫ลบอัลบั้ม🪶")
											}

										}
									}
								}
							} else if (op.Message.ContentType).String() == "STICKER" {
								if _, cek := data.ProSTICKER[to]; cek {
									if getAccess(ctime, cl.Mid) {
										if !sleepmode {
											if !fullAccess(sender) {
												cl.DeleteOtherFromChat(to, []string{sender})
												appendBl(sender)
												cl.SendMessage(to, "🪶💫ห้าม💫ส่งสติ้กเกอร์🪶")
											}

										}
									}
								}
							} else if (op.Message.ContentType).String() == "CALL" {
								if _, cek := data.ProCALL[to]; cek && op.Message.ContentMetadata["GC_MEDIA_TYPE"] == "AUDIO" {
									if getAccess(ctime, cl.Mid) {
										if !sleepmode {
											if !fullAccess(sender) {
												cl.DeleteOtherFromChat(to, []string{sender})
												appendBl(sender)
												cl.SendMessage(to, "🪶💫ห้าม💫โทรกลุ่ม🪶")
											}
										}

									}
								}
							} else if (op.Message.ContentType).String() == "FILE" {
								if _, cek := data.ProFILE[to]; cek {
									if getAccess(ctime, cl.Mid) {
										if !sleepmode {
											if !fullAccess(sender) {
												cl.DeleteOtherFromChat(to, []string{sender})
												appendBl(sender)
												cl.SendMessage(to, "🪶💫ห้าม💫ส่งไฟล์🪶")
											}
										}

									}
								}
							} else if (op.Message.ContentType).String() == "POSTNOTIFICATION" {
								if _, cek := data.ProPOSTNOTIFICATION[to]; cek {
									if getAccess(ctime, cl.Mid) {
										if !sleepmode {
											if !fullAccess(sender) {
												cl.DeleteOtherFromChat(to, []string{sender})
												appendBl(sender)
												cl.SendMessage(to, "🪶💫ห้าม💫สมาชิกโน้ต&&แชร์โพส🪶")
											}
										}

									}
								}
							} else if (op.Message.ContentType).String() == "AUDIO" {
								if _, cek := data.ProAUDIO[to]; cek {
									if getAccess(ctime, cl.Mid) {
										if !sleepmode {
											if !fullAccess(sender) {
												cl.DeleteOtherFromChat(to, []string{sender})
												appendBl(sender)
												cl.SendMessage(to, "🪶💫ห้าม💫ส่งคลิปเสียง🪶")
											}
										}

									}
								}
							} else if (op.Message.ContentType).String() == "CONTACT" {
								name := op.Message.ContentMetadata["displayName"]
								mid := op.Message.ContentMetadata["mid"]
								if Scont == true && PromoteStaff == true {
									if fullAccess(sender) {
										if !oop.Contains(data.Staff, mid) {
											data.Staff = append(data.Staff, mid)
											cl.SendMessage(to, "Contact Added to list Staff success !..")
										} else {
											cl.SendMessage(to, "Contact Already in Staff list !..")
										}
									}
								} else if Scont == true && PromoteAdmin == true {
									if fullAccess(sender) {
										if !oop.Contains(data.Admin, mid) {
											data.Admin = append(data.Admin, mid)
											cl.SendMessage(to, "Contact Added to list Admin success !..")
										} else {
											cl.SendMessage(to, "Contact Already in Admin list !..")
										}
									}
								} else if Scont == true && PromoteOwner == true {
									if fullAccess(sender) {
										if !oop.Contains(data.Owner, mid) {
											data.Owner = append(data.Owner, mid)
											cl.SendMessage(to, "Contact Added to list Owner success !..")
										} else {
											cl.SendMessage(to, "Contact Already in Owner list !..")
										}
									}
								} else if Scont == true && DemoteStaff == true {
									if fullAccess(sender) {
										if oop.Contains(data.Staff, mid) {
											data.Staff = oop.Remove(data.Staff, mid)
											cl.SendMessage(to, "Contact Remove from list  Staff success !..")
										} else {
											cl.SendMessage(to, "Contact Not have in Staff list !..")
										}
									}
								} else if Scont == true && DemoteAdmin == true {
									if fullAccess(sender) {
										if oop.Contains(data.Admin, mid) {
											data.Admin = oop.Remove(data.Admin, mid)
											cl.SendMessage(to, "Contact Remove from list  Admin success !..")
										} else {
											cl.SendMessage(to, "Contact Not have in Admin list !..")
										}
									}
								} else if Scont == true && DemoteOwner == true {
									if fullAccess(sender) {
										if oop.Contains(data.Owner, mid) {
											data.Owner = oop.Remove(data.Owner, mid)
											cl.SendMessage(to, "Contact Remove from list Owner success !..")
										} else {
											cl.SendMessage(to, "Contact Not have in Owner list !..")
										}
									}
								} else if PromoteBlacklist == true {
									if getAccess(ctime, cl.Mid) {
										if fullAccess(sender) {
											if !oop.Contains(data.Ban, mid) {
												data.Ban = append(data.Ban, mid)
												cl.SendMessage(to, "เพิ่ม "+name+" เข้าบัญชีดำเรียบร้อย")
											}
										}
									}
								} else if delBlacklist == true {
									if getAccess(ctime, cl.Mid) {
										if fullAccess(sender) {
											if getAccess(ctime, cl.Mid) {
												data.Ban = oop.Remove(data.Ban, mid)
												cl.SendMessage(to, "ลบ "+name+" ออกจากบัญชีดำเรียบร้อย")
											}
										}
									}
								}
							} else if (op.Message.ContentType).String() == "IMAGE" {
								if fullAccess(sender) {
									if _, cek := data.ProIMAGE[to]; cek {
										if !sleepmode {
											if getAccess(ctime, cl.Mid) {
												cl.DeleteOtherFromChat(to, []string{sender})
												appendBl(sender)
												cl.SendMessage(to, "🪶💫ห้าม💫ส่งรูปภาพ🪶")
											}
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

								if _, cek := data.ProVIDEO[to]; cek {
									if getAccess(ctime, cl.Mid) {
										if !sleepmode {
											if !fullAccess(sender) {
												cl.DeleteOtherFromChat(to, []string{sender})
												appendBl(sender)
												cl.SendMessage(to, "🪶💫ห้าม💫ส่งวีดีโอ🪶")
											}
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
					// cl.SendMessage(to, "")
					// fmt.Println("CorrectRevision end",cl.Mid)
					cl.CorrectRevision(op, true, false, false)
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
func main() {
	fileName := fmt.Sprintf("%v.txt", os.Args[1])
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	LimiterJoin = 1000
	LimiterKick = 1000
	Token := strings.Split(string(fileBytes), ",")
	dataRead, err := ioutil.ReadFile(dataPath)
	if err != nil {
		fmt.Println("Read Data", err)
	}
	json.Unmarshal(dataRead, &data)
	// fmt.Println(data)
	for num, auth := range Token {
		s := strings.Replace(auth, "\n", "", -1)
		if s == "" {
			continue
		}
		xcl := oop.Connect(num, s)
		err := xcl.LoginWithAuthToken()
		if err == nil {
			//select {
			//case <-quit:
			//return
			//default:
			go func(num int) {
				perBots(xcl)
			}(num)
			//}
			time.Sleep(1 * time.Second)

		} else {
			if strings.Contains(fmt.Sprintf("%s", err), "suspended") {
				fmt.Println(auth[:33], "FREEZE OR SUSPEND")
			} else {
				fmt.Println(auth[:33], err)
			}
			continue
		}
	}
	SaveData()
	ch := make(chan int, len(Token))
	for v := range ch {
		if v == 20 {
			break
		}
	}
}
