package logs

import (
	"strings"

	utils "../../utils"
	domaind "../domaind"
)

func IsDomainContains(detectedDomains []domaind.DdData, domain string) bool {
	for i := 0; i < len(detectedDomains); i++ {
		if strings.Contains(detectedDomains[i].Domains, domain) {
			return true
		}
	}
	return false
}

func DetectDomains(ddRule []domaind.DdData, pass_links string) (detectedDomains []domaind.DdData) {
	pass_links = strings.TrimRight(pass_links, ",")
	logLinks := strings.Split(pass_links, ",")
	if len(pass_links) == 0 || len(logLinks) == 0 {
		return detectedDomains
	}

	for _, rule := range ddRule {
		ruleLinks := strings.Split(utils.DelSpace(rule.Domains), ",")
		for _, link := range ruleLinks {
			if strings.Contains(pass_links, link) {
				detectedLink := domaind.DdData{
					Color:   rule.Color,
					Domains: link,
				}

				if !IsDomainContains(detectedDomains, detectedLink.Domains) {
					detectedDomains = append(detectedDomains, detectedLink)
				}
			}
		}
	}

	return detectedDomains
}

func SetDetectedDomains(logsData []LogsData) {
	domains := domaind.GetDd()
	if len(domains) == 0 {
		return
	}

	for i := 0; i < len(logsData); i++ {
		if len(logsData[i].Domains) == 0 {
			continue
		}

		logsData[i].DetectedDomains = DetectDomains(domains, logsData[i].Domains)
	}
}
