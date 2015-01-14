package sms

type Sms interface {
	Get(apikey string) (interface{}, error)
	Send(mobile, code string) error
	TplSend(mobile, code string, tpl_id int) error
}
