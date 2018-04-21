package main

type plugin struct {
	Information information `json:"information,omitempty"`
	Request     request     `json:"request,omitempty"`
	Status      status      `json:"status,omitempty"`
}

type status struct {
	JudgeYesKeyword string `json:"judge_yes_keyword,omitempty"`
	JudegNoKeyword  string `json:"judeg_no_keyword,omitempty"`
	ProfileURL      string `json:"profile_url,omitempty"`
}

type request struct {
	CellPhoneURL string            `json:"cellphone_url,omitempty"`
	EmailURL     string            `json:"email_url,omitempty"`
	UserURL      string            `json:"user_url,omitempty"`
	Method       string            `json:"method,omitempty"`
	PostFields   map[string]string `json:"post_fields,omitempty"`
}

type information struct {
	Author   string `json:"author,omitempty"`
	Date     string `json:"date,omitempty"`
	Name     string `json:"name,omitempty"`
	WebSite  string `json:"website,omitempty"`
	Category string `json:"category,omitempty"`
	Icon     string `json:"icon,omitempty"`
	Desc     string `json:"desc,omitempty"`
}

func (p *plugin) name() string {
	return p.Information.Name
}

// getURL 安装类型返回 url
func (p *plugin) getURL(checkType string) string {
	if checkType == "phone" {
		return p.Request.CellPhoneURL
	}
	if checkType == "email" {
		return p.Request.EmailURL
	}
	return p.Request.UserURL
}

// judgeYesKeyword 返回正确判断是前缀字符串
func (p *plugin) judgeYesKeyword() []byte {
	return []byte(p.Status.JudgeYesKeyword)
}
