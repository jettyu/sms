package yunpian

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	host             = "http://yunpian.com"
	user_get_uri     = "/user/get.json?apikey="
	sms_send_uri     = "/sms/send.json"
	sms_tpl_send_uri = "/sms/tpl_send.json"
	sms_ctx          = `apikey=%s&mobile=%s&text=%s`
	sms_tpl_ctx      = `apikey=%s&mobile=%s&tpl_value=%s&tpl_id=%d`
	post_type        = `application/x-www-form-urlencoded`
	def_tpl_value    = `#code#=%s&#company#=%s`
)

type YunPian struct {
	apikey           string
	version          string
	user_get_uri     string
	sms_send_uri     string
	sms_tpl_send_uri string
}

func NewYunPian(apikey, version string) *YunPian {
	y := &YunPian{}
	y.apikey = apikey
	y.version = version
	y.user_get_uri = host + "/" + version + user_get_uri
	y.sms_send_uri = host + "/" + version + sms_send_uri
	y.sms_tpl_send_uri = host + "/" + version + sms_tpl_send_uri
	return y
}

func (this *YunPian) Get() (string, error) {
	str := ""
	uri := this.user_get_uri + this.apikey
	resp, err := http.Get(uri)
	if err != nil {
		fmt.Println("err=" + err.Error())
		return str, err
	}
	defer resp.Body.Close()
	buf, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return str, e
	}
	str = string(buf)
	return str, nil
}

func (this *YunPian) Send(mobile, text string) (string, error) {
	str := ""
	c := &http.Client{}
	resp, err := c.Post(this.sms_send_uri, post_type, strings.NewReader(fmt.Sprintf(sms_ctx, this.apikey, mobile, text)))
	if err != nil {
		return str, err
	}
	defer resp.Body.Close()
	buf, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return str, e
	}
	str = string(buf)

	return str, nil
}

func (this *YunPian) TplSend(mobile, tpl_value string, tpl_id int) (string, error) {
	str := ""
	c := &http.Client{}
	tpl_value = url.QueryEscape(tpl_value)
	body := fmt.Sprintf(sms_tpl_ctx, this.apikey, mobile, tpl_value, tpl_id)
	fmt.Println("body=" + body)
	resp, err := c.Post(this.sms_tpl_send_uri, post_type, strings.NewReader(body))
	if err != nil {
		return str, err
	}
	defer resp.Body.Close()
	buf, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return str, e
	}
	str = string(buf)

	return str, nil
}

func (this *YunPian) DefTplSend(mobile, code, company string, tpl_id int) (string, error) {
	return this.TplSend(mobile, fmt.Sprintf(def_tpl_value, code, company), tpl_id)
}
