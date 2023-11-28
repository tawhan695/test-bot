package oop

import (
	"strconv"
	"strings"

	"botline/Library-mac/linethrift"
)

/*func GetIntBytes(n int) (valo []byte) {
	var bits = 64
	zigzag := ((n << 1) ^ (n >> (bits - 1)))
	for {
		if zigzag&-128 == 0 {
			valo = append(valo, byte(zigzag))
			break
		} else {
			valo = append(valo, byte((zigzag&0xff)|0x80))
			zigzag >>= 7
		}
	}
	return valo
}

func GetStringBytes(str string) []byte {
	var va []byte
	for a := range str {
		va = append(va, byte(int(str[a])))
	}
	return va
}

func (cl *Account) FetchOps(count int32) (res []*linethrift.Operation, err *modcompact.ExceptionMod) {
	// Building fetchOps manual bytes
	var fOB = []byte{130, 33, 1, 8, 102, 101, 116, 99, 104, 79, 112, 115, 38}
	fOB = append(fOB, GetIntBytes(int(cl.Rev))...)
	fOB = append(fOB, 21)
	fOB = append(fOB, GetIntBytes(int(count))...)
	fOB = append(fOB, 22)
	fOB = append(fOB, GetIntBytes(int(cl.Grev))...)
	fOB = append(fOB, 22)
	fOB = append(fOB, GetIntBytes(int(cl.Irev))...)
	fOB = append(fOB, 0)
	HTTP, _ := thrift.NewTHttpClient("https://" + cl.Host + ".line.naver.jp/P5")
	transport := HTTP.(*thrift.THttpClient)
	transport.SetHeader("user-agent", cl.UserAgent)
	transport.SetHeader("x-line-application", cl.LineApp)
	transport.SetHeader("x-line-access", cl.Authtoken)
	transport.SetHeader("x-las", "F")
	transport.SetHeader("x-lac", "44105") // need random 30000-60000
	transport.SetHeader("x-lam", "w")
	transport.SetHeader("x-cl", "23")
	transport.SetHeader("x-lal", "en_US")
	transport.SetHeader("accept", "application/x-thrift")
	if cl.Proxyip != "" && cl.Port != "" {
		transport.SetProxy(cl.Proxyip, cl.Port)
	}
	transport.SetMoreCompact(true)
	transport.Write(fOB)
	transport.Flush(cl.Ctx)
	//compact := thrift.NewTCompactProtocolFactory().GetProtocol(transport)
	//linethrift.NewTalkServiceClientProtocol(transport, compact, compact).FetchOps(cl.Ctx, cl.Rev, count, cl.Grev, cl.Irev)
	b := transport.GetBody()
	if len(b) > 0 {
		tmcp := modcompact.TMoreCompactProtocolGoods(b)
		return tmcp.GETOPS()
	}
	return res, err
}*/

func (cl *Account) FetchOps() ([]*linethrift.Operation, error) {
	return cl.Poll.FetchOps(cl.Ctx, cl.Rev, 100, cl.Grev, cl.Irev)
}

func (cl *Account) CorrectRevision(op *linethrift.Operation, local bool, global bool, individual bool) {
	if global {
		if op.Revision == -1 && op.Param2 != "" {
			s := strings.Split(op.Param2, "\x1e")
			cl.Grev, _ = strconv.ParseInt(s[0], 10, 64)
		}
	}
	if individual {
		if op.Revision == -1 && op.Param1 != "" {
			s := strings.Split(op.Param1, "\x1e")
			cl.Irev, _ = strconv.ParseInt(s[0], 10, 64)
			
		}
	}
	if local {
		if op.Revision > cl.Rev {
			cl.Rev = op.Revision
			
		}
		
	}
}

func (cl *Account) GetLastOpRevision() (r int64, err error) {
	return cl.Talk.GetLastOpRevision(cl.Ctx)
}
