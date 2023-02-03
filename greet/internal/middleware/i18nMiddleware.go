package middleware

import (
	"github.com/beego/i18n"
	"github.com/zeromicro/go-zero/core/logx"
	"greet/internal/config"
	"net/http"
	"net/url"
	"strings"
)

var currentLang string
var langTypes []*langType

type langType struct {
	Lang, Name string
}

func InitLang(c config.Config) {
	langs := strings.Split(c.Lang.Types, "|")
	names := strings.Split(c.Lang.Names, "|")
	langTypes = make([]*langType, 0, len(langs))
	for i, v := range langs {
		langTypes = append(langTypes, &langType{
			Lang: v,
			Name: names[i],
		})
	}
	for _, lang := range langs {
		logx.Debug("Loading language: " + lang)
		if err := i18n.SetMessage(lang, c.Lang.Path+"locale_"+lang+".ini"); err != nil {
			logx.Error("Fail to set message file: " + err.Error())
			return
		}
	}
}

type I18nMiddleware struct {
	conf config.Config
}

func NewI18nMiddleware(config config.Config) *I18nMiddleware {
	return &I18nMiddleware{
		conf: config,
	}
}

func (m *I18nMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		langs := strings.Split(m.conf.Lang.Names, "|")
		names := strings.Split(m.conf.Lang.Types, "|")
		langTypes := make([]*langType, 0, len(langs))
		for i, v := range langs {
			langTypes = append(langTypes, &langType{
				Lang: v,
				Name: names[i],
			})
		}
		hasCookie := false

		// 通过cookie获取语言信息
		var lang string
		langCookie, err := r.Cookie("lang")
		if err != nil {
			// 没有的话默认中文
			lang = "zh-CN"
		} else {
			hasCookie = true

			lang, _ = url.QueryUnescape(langCookie.Value)
			// 看下是否支持语言，默认中文
			if !i18n.IsExist(lang) {
				lang = "zh-CN"
				hasCookie = false
			}
		}

		curLang := langType{
			Lang: lang,
		}

		// 将语言的信息保存在cookie中
		if !hasCookie {
			//c.SetCookie("lang", curLang.Lang, 1<<31-1, "/", viper.GetString("domain"), false, false)
			newLangCookie := http.Cookie{
				Name:  "lang",
				Value: curLang.Lang,
			}
			http.SetCookie(w, &newLangCookie)
		}

		restLangs := make([]*langType, 0, len(langTypes)-1)
		for _, v := range langTypes {
			if lang != v.Lang {
				restLangs = append(restLangs, v)
			} else {
				curLang.Name = v.Name
			}
		}
		// 设置当前语言
		currentLang = curLang.Lang
		// Passthrough to next handler if need
		next(w, r)
	}
}

func GetLang() string {
	return currentLang
}
