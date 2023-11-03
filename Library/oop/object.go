package oop

import (
	"bufio"
	"bytes"
	b64 "encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func (cl *Account) DownloadObjectMsg(msgid, tipe string) (string, error) {
	hclient := &http.Client{}
	req, _ := http.NewRequest("GET", "https://obs-sg.line-apps.com/talk/m/download.nhn?oid="+msgid, nil)
	for k, v := range map[string]string{
		"User-Agent":         cl.UserAgent,
		"X-Line-Application": cl.LineApp,
		"X-Line-Access":      cl.Authtoken,
		"x-lal":              "en_US",
		"x-lpv":              "1",
	} {
		req.Header.Set(k, v)
	}
	res, _ := hclient.Do(req)
	defer res.Body.Close()
	file, err := os.Create("/tmp/DL-" + msgid + "." + tipe)
	if err != nil {
		return "", err
	}
	io.Copy(file, res.Body)
	file.Close()
	return file.Name(), nil
}

func (cl *Account) UpdateProfilePicture(path, tipe string) error {
	fl, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fl.Close()
	of, err := fl.Stat()
	if err != nil {
		return err
	}
	var size int64 = of.Size()
	bytess := make([]byte, size)
	buffer := bufio.NewReader(fl)
	_, err = buffer.Read(bytess)
	if err != nil {
		return err
	}
	dataa := ""
	nama := filepath.Base(path)
	if tipe == "vp" {
		dataa = fmt.Sprintf(`{"name": "%s", "oid": "%s", "type": "image", "ver": "2.0", "cat": "vp.mp4"}`, nama, cl.Mid)
	} else {
		dataa = fmt.Sprintf(`{"name": "%s", "oid": "%s", "type": "image", "ver": "2.0"}`, nama, cl.Mid)
	}
	sDec := b64.StdEncoding.EncodeToString([]byte(dataa))
	hclient := &http.Client{}
	req, _ := http.NewRequest("POST", "https://obs-sg.line-apps.com/talk/p/upload.nhn", bytes.NewBuffer(bytess))
	for k, v := range map[string]string{
		"User-Agent":         cl.UserAgent,
		"X-Line-Application": cl.LineApp,
		"X-Line-Access":      cl.Authtoken,
		"x-lal":              "en_US",
		"x-lpv":              "1",
	} {
		req.Header.Set(k, v)
	}
	req.Header.Set("x-obs-params", string(sDec))
	req.ContentLength = int64(len(bytess))
	res, err := hclient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}

func (cl *Account) UpdateProfilePictureVideo(pict, vid string) error {
	fl, err := os.Open(vid)
	if err != nil {
		return err
	}
	defer fl.Close()
	of, err := fl.Stat()
	if err != nil {
		return err
	}
	var size int64 = of.Size()
	bytess := make([]byte, size)
	buffer := bufio.NewReader(fl)
	_, err = buffer.Read(bytess)
	if err != nil {
		return err
	}
	dataa := fmt.Sprintf(`{"name": "%s", "oid": "%s", "ver": "2.0", "type": "video", "cat": "vp.mp4"}`, filepath.Base(vid), cl.Mid)
	sDec := b64.StdEncoding.EncodeToString([]byte(dataa))
	hclient := &http.Client{}
	req, _ := http.NewRequest("POST", "https://obs-sg.line-apps.com/talk/vp/upload.nhn", bytes.NewBuffer(bytess))
	for k, v := range map[string]string{
		"User-Agent":         cl.UserAgent,
		"X-Line-Application": cl.LineApp,
		"X-Line-Access":      cl.Authtoken,
		"x-lal":              "en_US",
		"x-lpv":              "1",
	} {
		req.Header.Set(k, v)
	}
	req.Header.Set("x-obs-params", string(sDec))
	req.ContentLength = int64(len(bytess))
	res, err := hclient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return cl.UpdateProfilePicture(pict, "vp")
}
