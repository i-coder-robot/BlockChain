package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"time"
)


var (
	ResponseHeaderTimeout = 30
	DialTimeout           = 5
)

var httpSerever *http.Client

func init() {
	tr := &http.Transport{
		Dial: func(netw, addr string) (net.Conn, error) {
		c, err := net.DialTimeout(netw, addr, time.Second*time.Duration(DialTimeout))
		if err != nil {
		return nil, err
	}
		return c, nil
	},
		ResponseHeaderTimeout: time.Second * time.Duration(ResponseHeaderTimeout),
	}
	httpSerever = &http.Client{Transport: tr}
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func HttpRequest(method, url string, headers map[string]string, data []byte) ([]byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	if headers != nil && len(headers) != 0 {
		for k, v := range headers{
			req.Header.Add(k, v)
		}
	}
	respBytes, err := HttpDo(req)
	return respBytes, err
}

func HttpGet(url string) ([]byte, string, error) {
	resp, err := httpSerever.Get(url)
	if err != nil {
		return nil, "", errors.New("RetryMarkup" + err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("RetryMarkup http get error : url=%v , statusCode=%v", url, resp.StatusCode)
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	return respBytes, resp.Header.Get("Content-Type"), err
}

func HttpPost(url, content_type string, data []byte) ([]byte, error) {
	resp, err := httpSerever.Post(url, content_type, bytes.NewBuffer(data))
	if err != nil {
		return nil, errors.New("RetryMarkup" + err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("RetryMarkup http get error : url=%v , statusCode=%v", url, resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}

func HttpDo(req *http.Request) ([]byte, error) {
	resp, err := httpSerever.Do(req)
	if err != nil {
		return nil, errors.New("RetryMarkup" + err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("RetryMarkup http get error : url=%v , statusCode=%v", req.URL, resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}

type UploadFile struct {
	FileName string
	FilePath string
	Data     []byte
}

func HttpUploadFiles(url string, fileList []*UploadFile) ([]byte, error) {

	b := new(bytes.Buffer)
	w := multipart.NewWriter(b)
	defer w.Close()
	for _, v := range fileList {
		if v.FilePath != "" {
			err := writeFileData(w, v.FileName, v.FilePath)
			if err != nil {
				return nil, err
			}
		} else if len(v.Data) != 0 {
			fw, err := w.CreateFormField(v.FileName)
			if err != nil {
				return nil, err
			}
			if _, err = fw.Write(v.Data); err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("文件上传必须有内容")
		}
	}
	req, err := http.NewRequest("POST", url, b)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "image/png")
	resp, err := httpSerever.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : url=%v , statusCode=%v", url, resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}

func writeFileData(w *multipart.Writer, name, filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	fw, err := w.CreateFormFile(name, filePath)
	if err != nil {
		return err
	}
	if _, err = io.Copy(fw, f); err != nil {
		return err
	}
	return nil
}

func ToString(v interface{}) string {
	bytes, e := json.Marshal(v)
	if e != nil {
		panic("序列化出错拉")
	}
	result:=string(bytes)
	return result
}
