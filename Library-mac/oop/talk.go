package oop

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"
	"context"

	"botline/Library-mac/linethrift"
	"botline/Library-mac/thrift"
)
var (
	MID 			string = ""
	KickBans    = []*Account{}
)


func (cl *Account) GetChats(chatMid []string, withMembers bool, withInvitees bool) (*linethrift.GetChatsResponse, error) {
	req := linethrift.NewGetChatsRequest()
	req.ChatMids = chatMid
	req.WithMembers = withMembers
	req.WithInvitees = withInvitees
	return cl.Talk.GetChats(cl.Ctx, req)
}

func (cl *Account) KickoutFromGroup(to string, mid []string) error {
	return cl.Talk2.KickoutFromGroup(cl.Ctx, 0, to, mid)
}
func (cl *Account) GetGroupIdsJoined() ([]string, error) {
	req := &linethrift.GetAllChatMidsRequest{
		WithInvitedChats: false,
		WithMemberChats:  true,
	}
	res, err := cl.Talk.GetAllChatMids(context.TODO(), req, linethrift.SyncReason_UNKNOWN)
	return res.MemberChatMids, err
}
/*func (cl *Account) InviteIntoGroup(to string, mid []string) error {
 return cl.Talk2.InviteIntoGroup(cl.Ctx, int32(0), to, mid)
}
func (cl *Account) AcceptGroupInvitationByTicket(to string, ticketID string) error {
	return cl.Talk2.AcceptGroupInvitationByTicket(cl.Ctx, int32(0), to, ticketID)
}*/

func (cl *Account) AcceptChatInvitation(chatMid string) error {
	req := linethrift.NewAcceptChatInvitationRequest()
	req.ReqSeq = 0
	req.ChatMid = chatMid
	_, err := cl.Talk.AcceptChatInvitation(cl.Ctx, req)
	return err
}

func (cl *Account) DeleteOtherFromChat(chatMid string, targetMid []string) error {
	req := linethrift.NewDeleteOtherFromChatRequest()
	req.ReqSeq = 0
	req.ChatMid = chatMid
	req.TargetUserMids = targetMid
	_, err := cl.Talk.DeleteOtherFromChat(cl.Ctx, req)
	return err
}

func (cl *Account) InviteIntoChat(chatMid string, targetMids []string) error {
	req := linethrift.NewInviteIntoChatRequest()
	req.ReqSeq = 0
	req.ChatMid = chatMid
	req.TargetUserMids = targetMids
	_, err := cl.Talk.InviteIntoChat(cl.Ctx, req)
	return err
}

func (cl *Account) UpdateProfile(text string, tipe string) error {
	req := linethrift.NewUpdateProfileAttributesRequest()
	content := linethrift.NewProfileContent()
	content.Value = text
	content.Meta = nil
	if tipe == "name" {
		req.ProfileAttributes = map[linethrift.ProfileAttribute]*linethrift.ProfileContent{linethrift.ProfileAttribute_DISPLAY_NAME: content}
	} else if tipe == "bio" {
		req.ProfileAttributes = map[linethrift.ProfileAttribute]*linethrift.ProfileContent{linethrift.ProfileAttribute_STATUS_MESSAGE: content}
	}
	return cl.Talk.UpdateProfileAttributes(cl.Ctx, 0, req)
}

func (cl *Account) CancelChatInvitation(chatMid string, targetMid []string) error {
	req := linethrift.NewCancelChatInvitationRequest()
	req.ReqSeq = 0
	req.ChatMid = chatMid
	req.TargetUserMids = targetMid
	_, err := cl.Talk.CancelChatInvitation(cl.Ctx, req)
	return err
}

func (cl *Account) GetAllChatMids(member bool, invited bool) (*linethrift.GetAllChatMidsResponse, error) {
	req := linethrift.NewGetAllChatMidsRequest()
	req.WithMemberChats = member
	req.WithInvitedChats = invited
	return cl.Talk.GetAllChatMids(cl.Ctx, req, linethrift.SyncReason_UNKNOWN)
}

func (cl *Account) AcceptChatInvitationByTicket(chatMid string, ticketID string) error {
	req := linethrift.NewAcceptChatInvitationByTicketRequest()
	req.ReqSeq = 0
	req.ChatMid = chatMid
	req.TicketId = ticketID
	_, err := cl.Talk.AcceptChatInvitationByTicket(cl.Ctx, req)
	return err
}

func (cl *Account) FindChatByTicket(ticketID string) (*linethrift.FindChatByTicketResponse, error) {
	req := linethrift.NewFindChatByTicketRequest()
	req.TicketId = ticketID
	return cl.Talk.FindChatByTicket(cl.Ctx, req)
}

func (cl *Account) ReissueChatTicket(chatMid string) (*linethrift.ReissueChatTicketResponse, error) {
	req := linethrift.NewReissueChatTicketRequest()
	req.ReqSeq = 0
	req.GroupMid = chatMid
	return cl.Talk.ReissueChatTicket(cl.Ctx, req)
}

func (cl *Account) RejectChatInvitation(chatMid string) error {
	req := linethrift.NewRejectChatInvitationRequest()
	req.ReqSeq = 0
	req.ChatMid = chatMid
	_, err := cl.Talk.RejectChatInvitation(cl.Ctx, req)
	return err
}

func (cl *Account) DeleteSelfFromChat(chatMid string) error {
	req := linethrift.NewDeleteSelfFromChatRequest()
	req.ReqSeq = 0
	req.ChatMid = chatMid
	req.LastSeenMessageDeliveredTime = 0
	req.LastSeenMessageId = ""
	req.LastMessageDeliveredTime = 0
	req.LastMessageId = ""
	_, err := cl.Talk.DeleteSelfFromChat(cl.Ctx, req)
	return err
}

func (cl *Account) CreateChat(name string, targetMids []string) (*linethrift.CreateChatResponse, error) {
	req := &linethrift.CreateChatRequest{
		ReqSeq:         0,
		Type:           linethrift.ChatType_GROUP,
		Name:           name,
		TargetUserMids: targetMids,
	}
	return cl.Talk.CreateChat(cl.Ctx, req)
}

func (cl *Account) SendMessage(to string, text string) (*linethrift.Message, error) {
	msg := &linethrift.Message{
		To:               to,
		Text:             text,
		ContentType:      linethrift.ContentType_NONE,
		RelatedMessageServiceCode: 1,
		//MessageRelationType: 3,
		RelatedMessageId: "0",
		ContentMetadata: map[string]string{},
	}
	return cl.Talk.SendMessage(cl.Ctx, 0, msg)
}

func (cl *Account) SendMessages(to string, text string, contentMetadata map[string]string) (*linethrift.Message, error) {
	msg := &linethrift.Message{
		To:               to,
		Text:             text,
		ContentType:      linethrift.ContentType_NONE,
		ContentMetadata:  contentMetadata,
		RelatedMessageId: "0",
	}
	return cl.Talk.SendMessage(cl.Ctx, 0, msg)
}

func (cl *Account) SendMention(to string, text string, mids []string) (*linethrift.Message, error) {
	var arr = []*Mention{}
	mentionee := "@rr6gbot"
	texts := strings.Split(text, "@!")
	if len(mids) == 0 || len(texts) < len(mids) {
		return &linethrift.Message{}, fmt.Errorf("invalid mids")
	}
	textx := ""
	for i, mid := range mids {
		textx += texts[i]
		arr = append(arr, &Mention{S: strconv.Itoa(utf8.RuneCountInString(textx)), E: strconv.Itoa(utf8.RuneCountInString(textx) + len(mentionee)), M: mid})
		textx += mentionee
	}
	textx += texts[len(mids)]
	arrData, _ := json.Marshal(arr)
	return cl.SendMessages(to, textx, map[string]string{"MENTION": "{\"MENTIONEES\":" + string(arrData) + "}"})
}

//func (cl *Account) UnsendChatnume(toId string, text string) (err error) {
//	return cl.Talk.UnsendMessage(cl.Ctx, int32(0), text)
//}

func (cl *Account) UnsendMessage(messageId string) (error) {
	return cl.Talk.UnsendMessage(context.TODO(), int32(0), messageId)
}
//func (cl *Account) UnsendChat(toId string) (err error) {
//	Nganu, _ := cl.Talk.GetRecentMessagesV2(cl.Ctx, toId, int32(1001))
//	Mid := []string{}
//	for _, chat := range Nganu {
//		if chat.From_ == MID {
//			Mid = append(Mid, chat.ID)
//		}
//	}
//	for i := 0; i < len(Mid); i++ {
//		err = cl.Talk.UnsendMessage(cl.Ctx, int32(0), Mid[i])
//	}
//	return err
//}

func (cl *Account) GetRecentMessagesV2(to string, count int32) ([]*linethrift.Message, error) {
	return cl.Talk.GetRecentMessagesV2(context.TODO(), to, count)
}

func (cl *Account) GetProfile() (*linethrift.Profile, error) {
	return cl.Talk.GetProfile(cl.Ctx, linethrift.SyncReason_UNKNOWN)
}

func (cl *Account) UpdateProfileAttribute(a linethrift.ProfileAttribute, v string) error {
	return cl.Talk.UpdateProfileAttribute(cl.Ctx, 0, a, v)
}

func (cl *Account) UpdateProfileAttributes(a linethrift.ProfileAttribute, v string) error {
	D := make(map[linethrift.ProfileAttribute]*linethrift.ProfileContent)
	D[a] = &linethrift.ProfileContent{Value: v}
	return cl.Talk.UpdateProfileAttributes(cl.Ctx, 0, &linethrift.UpdateProfileAttributesRequest{ProfileAttributes: D})
}
func (cl *Account) FindAndAddContactsByMid(mid string) (map[string]*linethrift.Contact, error) {
	return cl.Talk2.FindAndAddContactsByMid(cl.Ctx, 0, mid, linethrift.ContactType_MID, `{"screen":"homeTab","spec":"native"}`)
}

func (cl *Account) GetContact(mid string) (*linethrift.Contact, error) {
	return cl.Talk.GetContact(cl.Ctx, mid)
}

func (cl *Account) GetAllContactIds() ([]string, error) {
	return cl.Talk.GetAllContactIds(cl.Ctx, linethrift.SyncReason_UNKNOWN)
}

func (cl *Account) SendContact(to string, mid string) (*linethrift.Message, error) {
	msg := linethrift.NewMessage()
	msg.To = to
	msg.ContentType = linethrift.ContentType_CONTACT
	msg.ContentMetadata = map[string]string{"mid": mid}
	msg.RelatedMessageId = "0"
	return cl.Talk.SendMessage(cl.Ctx, 0, msg)
}

func (cl *Account) UpdateChatQr(chatId string, typevar bool) error {
	option := thrift.THttpClientOptions{
		Client: &http.Client{
			Transport: &http.Transport{},
		},
	}
	HTTP, _ := thrift.NewTHttpClientWithOptions(cl.Host+"/S4", option)
	transport := HTTP.(*thrift.THttpClient)
	transport.SetHeader("user-agent", cl.UserAgent)
	transport.SetHeader("x-line-application", cl.LineApp)
	transport.SetHeader("x-line-access", cl.Authtoken)
	transport.SetHeader("x-lal", "en_US")
	transport.SetHeader("x-lpv", "1")
	//if cl.Proxyip != "" && cl.Port != "" {
	//transport.SetProxy(cl.Proxyip, cl.Port)
	//}
	var x string
	if typevar {
		x = "!"
	}
	transport.Write([]byte("\x82!\x00\nupdateChat\x1c\x15\x00\x1c(!" + chatId + "l\x1c" + x + "\x00\x00\x00\x15\x08\x00\x00"))
	return transport.Flush(cl.Ctx)
}

func (cl *Account) UpdateChatName(chatId string, name string) error {
	var fOB = []byte{130, 33, 1, 10, 117, 112, 100, 97, 116, 101, 67, 104, 97, 116, 28, 21, 0, 28, 21, 0, 24, 33}
	fOB = append(fOB, []byte(chatId)...)
	fOB = append(fOB, []byte{22, 0, 18, 22, 0, 24, byte(len(name))}...)
	fOB = append(fOB, []byte(name)...)
	fOB = append(fOB, []byte{24, 0, 0, 21, 2, 0, 0}...)
	option := thrift.THttpClientOptions{
		Client: &http.Client{
			Transport: &http.Transport{},
		},
	}
	HTTP, _ := thrift.NewTHttpClientWithOptions(cl.Host+"/S4", option)
	transport := HTTP.(*thrift.THttpClient)
	transport.SetHeader("user-agent", cl.UserAgent)
	transport.SetHeader("x-line-application", cl.LineApp)
	transport.SetHeader("x-line-access", cl.Authtoken)
	transport.SetHeader("x-lal", "en_US")
	transport.SetHeader("x-lpv", "1")
	//if cl.Proxyip != "" && cl.Port != "" {
	//transport.SetProxy(cl.Proxyip, cl.Port)
	//}
	transport.Write(fOB)
	return transport.Flush(cl.Ctx)
}
