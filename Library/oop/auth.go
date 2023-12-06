package oop

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"../linethrift"
	"../thrift"
)

type Account struct {
	Mid       string
	Authtoken string
	Host      string
	UserAgent string
	LineApp   string
	LineApp2  string
	Appver    string
	Proxyip   string
	Port      string
	Rev       int64
	Grev      int64
	Irev      int64
	Sort      int
	Count     int
	SpamJoin  int
	Ctx       context.Context
	Talk      *linethrift.TalkServiceClient
	Talk2     *linethrift.TalkServiceClient
	Poll      *linethrift.TalkServiceClient
}

var (
	Proxylist = []string{
		"172.105.199.175:8080",
		"172.104.113.207:8080",
	}
	SUBDOMAIN = []string{
		"legy-jp",
		"legy-jp-addr",
		"legy-sg-addr",
		"legy-jp-addr-long",
		"legy-jp-addr-short",
		"legy-jp-short",
		"legy-jp-long",
		"ga2",
		"gd2",
		"gwz",
		"gws",
		"gf",
	}
	ListApp  = []string{"ANDROID\t12.4.0\tAndroid OS\t11.0.0", "ANDROID\t12.0.0\tAndroid OS\t9.0.0", "ANDROID\t12.3.0\tAndroid OS\t8.0.0", "ANDROID\t11.20.0\tAndroid OS\t7.1.2", "ANDROID\t11.19.0\tAndroid OS\t8.1.0", "ANDROID\t12.2.0\tAndroid OS\t9.0.0", "ANDROID\t11.20.0\tAndroid OS\t9.0.0", "ANDROID\t11.21.1\tAndroid OS\t8.1.0", "ANDROID\t11.19.0\tAndroid OS\t9.0.0", "ANDROID\t12.2.0\tAndroid OS\t8.0.0", "ANDROID\t11.22.1\tAndroid OS\t10.0.0", "ANDROID\t12.4.0\tAndroid OS\t10.0.0", "ANDROID\t12.3.0\tAndroid OS\t7.1.2", "ANDROID\t12.1.0\tAndroid OS\t9.0.0", "ANDROID\t11.19.0\tAndroid OS\t7.1.2", "ANDROID\t12.1.0\tAndroid OS\t8.1.0", "ANDROID\t12.1.0\tAndroid OS\t11.0.0", "ANDROID\t11.20.0\tAndroid OS\t10.0.0", "ANDROID\t12.4.0\tAndroid OS\t7.0.0", "ANDROID\t12.0.0\tAndroid OS\t7.1.1", "ANDROID\t12.2.0\tAndroid OS\t8.1.0", "ANDROID\t11.22.1\tAndroid OS\t7.1.2", "ANDROID\t11.20.0\tAndroid OS\t11.0.0", "ANDROID\t12.4.0\tAndroid OS\t8.1.0", "ANDROID\t12.5.0\tAndroid OS\t12.0.0", "ANDROID\t12.2.0\tAndroid OS\t7.1.1", "ANDROID\t12.1.0\tAndroid OS\t7.0.0", "ANDROID\t12.6.0\tAndroid OS\t12.0.0", "ANDROID\t11.21.1\tAndroid OS\t7.0.0", "ANDROID\t12.0.0\tAndroid OS\t7.0.0", "ANDROID\t12.5.0\tAndroid OS\t7.1.1", "ANDROID\t11.22.1\tAndroid OS\t12.0.0", "ANDROID\t11.19.0\tAndroid OS\t11.0.0", "ANDROID\t12.4.0\tAndroid OS\t12.0.0", "ANDROID\t12.3.0\tAndroid OS\t7.1.1", "ANDROID\t12.4.0\tAndroid OS\t7.1.1", "ANDROID\t12.0.0\tAndroid OS\t10.0.0", "ANDROID\t11.19.0\tAndroid OS\t8.0.0", "ANDROID\t11.20.0\tAndroid OS\t12.0.0", "ANDROID\t12.1.0\tAndroid OS\t8.0.0", "ANDROID\t11.19.0\tAndroid OS\t10.0.0", "ANDROID\t11.21.1\tAndroid OS\t7.1.1", "ANDROID\t12.3.0\tAndroid OS\t11.0.0", "ANDROID\t12.0.0\tAndroid OS\t8.1.0", "ANDROID\t12.0.0\tAndroid OS\t7.1.2", "ANDROID\t12.2.0\tAndroid OS\t7.1.2", "ANDROID\t12.5.0\tAndroid OS\t7.0.0", "ANDROID\t12.6.0\tAndroid OS\t9.0.0", "ANDROID\t12.3.0\tAndroid OS\t8.1.0", "ANDROID\t12.5.0\tAndroid OS\t10.0.0", "ANDROID\t12.3.0\tAndroid OS\t10.0.0", "ANDROID\t11.21.1\tAndroid OS\t8.0.0", "ANDROID\t12.6.0\tAndroid OS\t10.0.0", "ANDROID\t12.0.0\tAndroid OS\t12.0.0", "ANDROID\t12.1.0\tAndroid OS\t7.1.2", "ANDROID\t12.6.0\tAndroid OS\t8.0.0", "ANDROID\t11.20.0\tAndroid OS\t7.0.0", "ANDROID\t11.21.1\tAndroid OS\t7.1.2", "ANDROID\t12.6.0\tAndroid OS\t7.1.1", "ANDROID\t12.2.0\tAndroid OS\t10.0.0", "ANDROID\t11.21.1\tAndroid OS\t12.0.0", "ANDROID\t12.3.0\tAndroid OS\t9.0.0", "ANDROID\t12.4.0\tAndroid OS\t9.0.0", "ANDROID\t12.2.0\tAndroid OS\t7.0.0", "ANDROID\t11.22.1\tAndroid OS\t11.0.0", "ANDROID\t12.2.0\tAndroid OS\t11.0.0", "ANDROID\t12.6.0\tAndroid OS\t7.1.2", "ANDROID\t12.4.0\tAndroid OS\t7.1.2", "ANDROID\t12.5.0\tAndroid OS\t8.1.0", "ANDROID\t11.22.1\tAndroid OS\t7.0.0", "ANDROID\t12.5.0\tAndroid OS\t11.0.0", "ANDROID\t12.5.0\tAndroid OS\t9.0.0", "ANDROID\t12.1.0\tAndroid OS\t7.1.1", "ANDROID\t11.19.0\tAndroid OS\t7.1.1", "ANDROID\t11.22.1\tAndroid OS\t8.1.0", "ANDROID\t11.22.1\tAndroid OS\t7.1.1", "ANDROID\t11.19.0\tAndroid OS\t7.0.0", "ANDROID\t12.6.0\tAndroid OS\t11.0.0", "ANDROID\t12.6.0\tAndroid OS\t7.0.0", "ANDROID\t12.1.0\tAndroid OS\t12.0.0", "ANDROID\t11.22.1\tAndroid OS\t9.0.0", "ANDROID\t11.19.0\tAndroid OS\t12.0.0", "ANDROID\t12.1.0\tAndroid OS\t10.0.0", "ANDROID\t11.21.1\tAndroid OS\t10.0.0", "ANDROID\t12.5.0\tAndroid OS\t8.0.0", "ANDROID\t12.0.0\tAndroid OS\t11.0.0", "ANDROID\t12.3.0\tAndroid OS\t12.0.0", "ANDROID\t12.0.0\tAndroid OS\t8.0.0", "ANDROID\t12.4.0\tAndroid OS\t8.0.0", "ANDROID\t11.20.0\tAndroid OS\t8.0.0"}
	ListApp2 = []string{"CHANNELCP\t12.4.0\tAndroid OS\t11.0.0", "CHANNELCP\t12.0.0\tAndroid OS\t9.0.0", "CHANNELCP\t12.3.0\tAndroid OS\t8.0.0", "CHANNELCP\t11.20.0\tAndroid OS\t7.1.2", "CHANNELCP\t11.19.0\tAndroid OS\t8.1.0", "CHANNELCP\t12.2.0\tAndroid OS\t9.0.0", "CHANNELCP\t11.20.0\tAndroid OS\t9.0.0", "CHANNELCP\t11.21.1\tAndroid OS\t8.1.0", "CHANNELCP\t11.19.0\tAndroid OS\t9.0.0", "CHANNELCP\t12.2.0\tAndroid OS\t8.0.0", "CHANNELCP\t11.22.1\tAndroid OS\t10.0.0", "CHANNELCP\t12.4.0\tAndroid OS\t10.0.0", "CHANNELCP\t12.3.0\tAndroid OS\t7.1.2", "CHANNELCP\t12.1.0\tAndroid OS\t9.0.0", "CHANNELCP\t11.19.0\tAndroid OS\t7.1.2", "CHANNELCP\t12.1.0\tAndroid OS\t8.1.0", "CHANNELCP\t12.1.0\tAndroid OS\t11.0.0", "CHANNELCP\t11.20.0\tAndroid OS\t10.0.0", "CHANNELCP\t12.4.0\tAndroid OS\t7.0.0", "CHANNELCP\t12.0.0\tAndroid OS\t7.1.1", "CHANNELCP\t12.2.0\tAndroid OS\t8.1.0", "CHANNELCP\t11.22.1\tAndroid OS\t7.1.2", "CHANNELCP\t11.20.0\tAndroid OS\t11.0.0", "CHANNELCP\t12.4.0\tAndroid OS\t8.1.0", "CHANNELCP\t12.5.0\tAndroid OS\t12.0.0", "CHANNELCP\t12.2.0\tAndroid OS\t7.1.1", "CHANNELCP\t12.1.0\tAndroid OS\t7.0.0", "CHANNELCP\t12.6.0\tAndroid OS\t12.0.0", "CHANNELCP\t11.21.1\tAndroid OS\t7.0.0", "CHANNELCP\t12.0.0\tAndroid OS\t7.0.0", "CHANNELCP\t12.5.0\tAndroid OS\t7.1.1", "CHANNELCP\t11.22.1\tAndroid OS\t12.0.0", "CHANNELCP\t11.19.0\tAndroid OS\t11.0.0", "CHANNELCP\t12.4.0\tAndroid OS\t12.0.0", "CHANNELCP\t12.3.0\tAndroid OS\t7.1.1", "CHANNELCP\t12.4.0\tAndroid OS\t7.1.1", "CHANNELCP\t12.0.0\tAndroid OS\t10.0.0", "CHANNELCP\t11.19.0\tAndroid OS\t8.0.0", "CHANNELCP\t11.20.0\tAndroid OS\t12.0.0", "CHANNELCP\t12.1.0\tAndroid OS\t8.0.0", "CHANNELCP\t11.19.0\tAndroid OS\t10.0.0", "CHANNELCP\t11.21.1\tAndroid OS\t7.1.1", "CHANNELCP\t12.3.0\tAndroid OS\t11.0.0", "CHANNELCP\t12.0.0\tAndroid OS\t8.1.0", "CHANNELCP\t12.0.0\tAndroid OS\t7.1.2", "CHANNELCP\t12.2.0\tAndroid OS\t7.1.2", "CHANNELCP\t12.5.0\tAndroid OS\t7.0.0", "CHANNELCP\t12.6.0\tAndroid OS\t9.0.0", "CHANNELCP\t12.3.0\tAndroid OS\t8.1.0", "CHANNELCP\t12.5.0\tAndroid OS\t10.0.0", "CHANNELCP\t12.3.0\tAndroid OS\t10.0.0", "CHANNELCP\t11.21.1\tAndroid OS\t8.0.0", "CHANNELCP\t12.6.0\tAndroid OS\t10.0.0", "CHANNELCP\t12.0.0\tAndroid OS\t12.0.0", "CHANNELCP\t12.1.0\tAndroid OS\t7.1.2", "CHANNELCP\t12.6.0\tAndroid OS\t8.0.0", "CHANNELCP\t11.20.0\tAndroid OS\t7.0.0", "CHANNELCP\t11.21.1\tAndroid OS\t7.1.2", "CHANNELCP\t12.6.0\tAndroid OS\t7.1.1", "CHANNELCP\t12.2.0\tAndroid OS\t10.0.0", "CHANNELCP\t11.21.1\tAndroid OS\t12.0.0", "CHANNELCP\t12.3.0\tAndroid OS\t9.0.0", "CHANNELCP\t12.4.0\tAndroid OS\t9.0.0", "CHANNELCP\t12.2.0\tAndroid OS\t7.0.0", "CHANNELCP\t11.22.1\tAndroid OS\t11.0.0", "CHANNELCP\t12.2.0\tAndroid OS\t11.0.0", "CHANNELCP\t12.6.0\tAndroid OS\t7.1.2", "CHANNELCP\t12.4.0\tAndroid OS\t7.1.2", "CHANNELCP\t12.5.0\tAndroid OS\t8.1.0", "CHANNELCP\t11.22.1\tAndroid OS\t7.0.0", "CHANNELCP\t12.5.0\tAndroid OS\t11.0.0", "CHANNELCP\t12.5.0\tAndroid OS\t9.0.0", "CHANNELCP\t12.1.0\tAndroid OS\t7.1.1", "CHANNELCP\t11.19.0\tAndroid OS\t7.1.1", "CHANNELCP\t11.22.1\tAndroid OS\t8.1.0", "CHANNELCP\t11.22.1\tAndroid OS\t7.1.1", "CHANNELCP\t11.19.0\tAndroid OS\t7.0.0", "CHANNELCP\t12.6.0\tAndroid OS\t11.0.0", "CHANNELCP\t12.6.0\tAndroid OS\t7.0.0", "CHANNELCP\t12.1.0\tAndroid OS\t12.0.0", "CHANNELCP\t11.22.1\tAndroid OS\t9.0.0", "CHANNELCP\t11.19.0\tAndroid OS\t12.0.0", "CHANNELCP\t12.1.0\tAndroid OS\t10.0.0", "CHANNELCP\t11.21.1\tAndroid OS\t10.0.0", "CHANNELCP\t12.5.0\tAndroid OS\t8.0.0", "CHANNELCP\t12.0.0\tAndroid OS\t11.0.0", "CHANNELCP\t12.3.0\tAndroid OS\t12.0.0", "CHANNELCP\t12.0.0\tAndroid OS\t8.0.0", "CHANNELCP\t12.4.0\tAndroid OS\t8.0.0", "CHANNELCP\t11.20.0\tAndroid OS\t8.0.0"}
)

func Connect(num int, authToken string) *Account {
	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(ListApp)
	a := strings.Split(ListApp[n], "\t")
	//fmt.Println(a[0])
	sort := rand.Intn(9999-1000) + 1000
	Splitter := strings.Split(authToken, ":")
	Clients := new(Account)
	Clients.Mid = Splitter[0]
	Clients.Authtoken = authToken
	Clients.Host = "https://" + SUBDOMAIN[3] + ".line.naver.jp"
	Clients.UserAgent = fmt.Sprintf("Line/%v", a[1])
	Clients.LineApp = fmt.Sprintf("%v", ListApp[n])
	Clients.LineApp2 = fmt.Sprintf("%v", ListApp2[n])
	//Clients.UserAgent = fmt.Sprintf("Line/12.4.0 A5012F 14.4098.92939495%v", sort)
	//Clients.LineApp = fmt.Sprintf("IOS_RC\t12.4.0\tiOS\t14.4098.92939495%v", sort)
	Clients.Proxyip = ""
	Clients.Port = ""
	Clients.Count = num
	Clients.Rev = -1
	Clients.Grev = 0
	Clients.Irev = 0
	Clients.Sort = sort
	Clients.Ctx = context.Background()
	return Clients
	//rand.Seed(time.Now().UnixNano())
	//Headers := []string{"ANDROID\t12.4.0\tAndroid OS\t11.0.0", "ANDROID\t12.0.0\tAndroid OS\t9.0.0", "ANDROID\t12.3.0\tAndroid OS\t8.0.0", "ANDROID\t11.20.0\tAndroid OS\t7.1.2", "ANDROID\t11.19.0\tAndroid OS\t8.1.0", "ANDROID\t12.2.0\tAndroid OS\t9.0.0", "ANDROID\t11.20.0\tAndroid OS\t9.0.0", "ANDROID\t11.21.1\tAndroid OS\t8.1.0", "ANDROID\t11.19.0\tAndroid OS\t9.0.0", "ANDROID\t12.2.0\tAndroid OS\t8.0.0", "ANDROID\t11.22.1\tAndroid OS\t10.0.0", "ANDROID\t12.4.0\tAndroid OS\t10.0.0", "ANDROID\t12.3.0\tAndroid OS\t7.1.2", "ANDROID\t12.1.0\tAndroid OS\t9.0.0", "ANDROID\t11.19.0\tAndroid OS\t7.1.2", "ANDROID\t12.1.0\tAndroid OS\t8.1.0", "ANDROID\t12.1.0\tAndroid OS\t11.0.0", "ANDROID\t11.20.0\tAndroid OS\t10.0.0", "ANDROID\t12.4.0\tAndroid OS\t7.0.0", "ANDROID\t12.0.0\tAndroid OS\t7.1.1", "ANDROID\t12.2.0\tAndroid OS\t8.1.0", "ANDROID\t11.22.1\tAndroid OS\t7.1.2", "ANDROID\t11.20.0\tAndroid OS\t11.0.0", "ANDROID\t12.4.0\tAndroid OS\t8.1.0", "ANDROID\t12.5.0\tAndroid OS\t12.0.0", "ANDROID\t12.2.0\tAndroid OS\t7.1.1", "ANDROID\t12.1.0\tAndroid OS\t7.0.0", "ANDROID\t12.6.0\tAndroid OS\t12.0.0", "ANDROID\t11.21.1\tAndroid OS\t7.0.0", "ANDROID\t12.0.0\tAndroid OS\t7.0.0", "ANDROID\t12.5.0\tAndroid OS\t7.1.1", "ANDROID\t11.22.1\tAndroid OS\t12.0.0", "ANDROID\t11.19.0\tAndroid OS\t11.0.0", "ANDROID\t12.4.0\tAndroid OS\t12.0.0", "ANDROID\t12.3.0\tAndroid OS\t7.1.1", "ANDROID\t12.4.0\tAndroid OS\t7.1.1", "ANDROID\t12.0.0\tAndroid OS\t10.0.0", "ANDROID\t11.19.0\tAndroid OS\t8.0.0", "ANDROID\t11.20.0\tAndroid OS\t12.0.0", "ANDROID\t12.1.0\tAndroid OS\t8.0.0", "ANDROID\t11.19.0\tAndroid OS\t10.0.0", "ANDROID\t11.21.1\tAndroid OS\t7.1.1", "ANDROID\t12.3.0\tAndroid OS\t11.0.0", "ANDROID\t12.0.0\tAndroid OS\t8.1.0", "ANDROID\t12.0.0\tAndroid OS\t7.1.2", "ANDROID\t12.2.0\tAndroid OS\t7.1.2", "ANDROID\t12.5.0\tAndroid OS\t7.0.0", "ANDROID\t12.6.0\tAndroid OS\t9.0.0", "ANDROID\t12.3.0\tAndroid OS\t8.1.0", "ANDROID\t12.5.0\tAndroid OS\t10.0.0", "ANDROID\t12.3.0\tAndroid OS\t10.0.0", "ANDROID\t11.21.1\tAndroid OS\t8.0.0", "ANDROID\t12.6.0\tAndroid OS\t10.0.0", "ANDROID\t12.0.0\tAndroid OS\t12.0.0", "ANDROID\t12.1.0\tAndroid OS\t7.1.2", "ANDROID\t12.6.0\tAndroid OS\t8.0.0", "ANDROID\t11.20.0\tAndroid OS\t7.0.0", "ANDROID\t11.21.1\tAndroid OS\t7.1.2", "ANDROID\t12.6.0\tAndroid OS\t7.1.1", "ANDROID\t12.2.0\tAndroid OS\t10.0.0", "ANDROID\t11.21.1\tAndroid OS\t12.0.0", "ANDROID\t12.3.0\tAndroid OS\t9.0.0", "ANDROID\t12.4.0\tAndroid OS\t9.0.0", "ANDROID\t12.2.0\tAndroid OS\t7.0.0", "ANDROID\t11.22.1\tAndroid OS\t11.0.0", "ANDROID\t12.2.0\tAndroid OS\t11.0.0", "ANDROID\t12.6.0\tAndroid OS\t7.1.2", "ANDROID\t12.4.0\tAndroid OS\t7.1.2", "ANDROID\t12.5.0\tAndroid OS\t8.1.0", "ANDROID\t11.22.1\tAndroid OS\t7.0.0", "ANDROID\t12.5.0\tAndroid OS\t11.0.0", "ANDROID\t12.5.0\tAndroid OS\t9.0.0", "ANDROID\t12.1.0\tAndroid OS\t7.1.1", "ANDROID\t11.19.0\tAndroid OS\t7.1.1", "ANDROID\t11.22.1\tAndroid OS\t8.1.0", "ANDROID\t11.22.1\tAndroid OS\t7.1.1", "ANDROID\t11.19.0\tAndroid OS\t7.0.0", "ANDROID\t12.6.0\tAndroid OS\t11.0.0", "ANDROID\t12.6.0\tAndroid OS\t7.0.0", "ANDROID\t12.1.0\tAndroid OS\t12.0.0", "ANDROID\t11.22.1\tAndroid OS\t9.0.0", "ANDROID\t11.19.0\tAndroid OS\t12.0.0", "ANDROID\t12.1.0\tAndroid OS\t10.0.0", "ANDROID\t11.21.1\tAndroid OS\t10.0.0", "ANDROID\t12.5.0\tAndroid OS\t8.0.0", "ANDROID\t12.0.0\tAndroid OS\t11.0.0", "ANDROID\t12.3.0\tAndroid OS\t12.0.0", "ANDROID\t12.0.0\tAndroid OS\t8.0.0", "ANDROID\t12.4.0\tAndroid OS\t8.0.0", "ANDROID\t11.20.0\tAndroid OS\t8.0.0",}
	//choosen := Headers[rand.Int()%len(Headers)]
	//Clients.LineApp = fmt.Sprintf(choosen)
	//Clients.UserAgent = fmt.Sprintf("Line/%v %v\t%v", strings.Split(Clients.LineApp, "\t")[1], strings.Split(Clients.LineApp, "\t")[2], strings.Split(Clients.LineApp, "\t")[3])
	//Clients.UserAgent = "Line/12.1.0"
	//Clients.LineApp = fmt.Sprintf("IOSIPAD\t12.1.0\tiPadOS\t15.3.%v", sort)
	//Clients.UserAgentCP = fmt.Sprintf("LLA/2.17.0 SM-G610F 5.1.%v", sort)
	//Clients.LineAppCP = fmt.Sprintf("CHANNELCP\t2.17.0\tAndroid OS\t5.1.%v", sort)
}

func (cl *Account) LoginWithAuthToken() error {
	cl.Talk = cl.TalkService()
	cl.Talk2 = cl.TalkService2()
	cl.Poll = cl.PollService()
	rev, err := cl.GetLastOpRevision()
	if err != nil {
		return err
	}
	cl.Rev = rev
	profile, _ := cl.GetProfile()
	fmt.Println(profile.DisplayName, cl.LineApp)
	return nil
}

/* Connection */
func (cl *Account) TalkService() *linethrift.TalkServiceClient {
	option := thrift.THttpClientOptions{
		Client: &http.Client{
			Transport: &http.Transport{},
		},
	}
	HTTP, _ := thrift.NewTHttpClientWithOptions(cl.Host+"/S4", option) //F4
	transport := HTTP.(*thrift.THttpClient)
	transport.SetHeader("user-agent", cl.UserAgent)
	transport.SetHeader("x-line-application", cl.LineApp)
	transport.SetHeader("x-line-access", cl.Authtoken)
	transport.SetHeader("x-lal", "en_US")
	transport.SetHeader("x-lpv", "1")
	//transport.SetHeader("accept", "application/x-thrift")
	//if cl.Proxyip != "" && cl.Port != "" {
	//transport.SetProxy(cl.Proxyip, cl.Port)
	//}
	compact := thrift.NewTCompactProtocolFactory().GetProtocol(transport)
	return linethrift.NewTalkServiceClient(thrift.NewTStandardClient(compact, compact))
}

func (cl *Account) TalkService2() *linethrift.TalkServiceClient {
	option := thrift.THttpClientOptions{
		Client: &http.Client{
			Transport: &http.Transport{},
		},
	}
	HTTP, _ := thrift.NewTHttpClientWithOptions(cl.Host+"/S4", option) //F4
	transport := HTTP.(*thrift.THttpClient)
	transport.SetHeader("user-agent", cl.UserAgent)
	transport.SetHeader("x-line-application", cl.LineApp2)
	transport.SetHeader("x-line-access", cl.Authtoken)
	transport.SetHeader("x-lal", "en_US")
	transport.SetHeader("x-lpv", "1")
	//transport.SetHeader("accept", "application/x-thrift")
	//if cl.Proxyip != "" && cl.Port != "" {
	//transport.SetProxy(cl.Proxyip, cl.Port)
	//}
	compact := thrift.NewTCompactProtocolFactory().GetProtocol(transport)
	return linethrift.NewTalkServiceClient(thrift.NewTStandardClient(compact, compact))
}

func (cl *Account) PollService() *linethrift.TalkServiceClient {
	option := thrift.THttpClientOptions{
		Client: &http.Client{
			Transport: &http.Transport{},
		},
	}
	HTTP, _ := thrift.NewTHttpClientWithOptions(cl.Host+"/P4", option) //F4
	transport := HTTP.(*thrift.THttpClient)
	transport.SetHeader("user-agent", cl.UserAgent)
	transport.SetHeader("x-line-application", cl.LineApp)
	transport.SetHeader("x-line-access", cl.Authtoken)
	transport.SetHeader("x-las", "F")
	transport.SetHeader("x-lac", "44105") // need random 30000-60000
	transport.SetHeader("x-lam", "w")
	transport.SetHeader("x-cl", "23")
	transport.SetHeader("x-lal", "en_US")
	//transport.SetHeader("accept", "application/x-thrift")
	//if cl.Proxyip != "" && cl.Port != "" {
	//transport.SetProxy(cl.Proxyip, cl.Port)
	//}
	compact := thrift.NewTCompactProtocolFactory().GetProtocol(transport)
	return linethrift.NewTalkServiceClient(thrift.NewTStandardClient(compact, compact))
	//pcol := thrift.NewTCompactProtocol(HTTP)
	//tstc := thrift.NewTStandardClient(pcol, pcol)
	//return linethrift.NewTalkServiceClient(tstc)
}
