package types

type AlertTrigger struct {
	SiteName  string `json:"sitename"`
	SiteURL   string `json:"siteurl"`
	AlertType string `json:"alerttype"`
}

type ProjectDetailsResponse struct {
	Username     string        `json:"username"`
	ProjectName  string        `json:"projectname"`
	AlertTriggers []AlertTrigger `json:"alertTriggers"`
}