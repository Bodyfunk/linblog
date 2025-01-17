package repository

import (
	"linblog/config"
	"linblog/model"
)

type SiteRepository struct {
}

func (site *SiteRepository) Get() interface{} {
	siteInfo := &model.Site{
		Avatar: config.Avatar,
		Slogan: config.Slogan,
		Name:   config.Name,
		Domain: config.Domain,
		Notice: config.Notice,
		Desc:   config.Desc,
	}
	return siteInfo
}

func (site *SiteRepository) GetSocials() []*model.Social {
	socials := []*model.Social{
		{
			Id:    1,
			Title: "QQ",
			Icon:  "iconqq",
			Color: "#1AB6FF ",
			Href:  "http://wpa.qq.com/msgrd?v=3&uin=874183200&site=qq&menu=yes",
		},
		{
			Id:    2,
			Title: "GitHub",
			Icon:  "icongithub",
			Color: "",
			Href:  "https://github.com/Codexiaoyi",
		},
	}
	return socials
}
