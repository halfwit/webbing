package doctor

import (
	"github.com/olmaxmedical/olmax_go/router"
	"golang.org/x/text/message"
)

func init() {
	b := &router.Page{
		Access: router.GuestAuth,
		CSS:    "",
		Path:   "doctor/application",
		Data:   Application,
		Extra:  router.ListCountries | router.ListServices | router.FormErrors | router.FormToken,
	}
	router.Add(b)
}

// Application - olmaxmedical.com/doctor/application.html
func Application(p *message.Printer) map[string]interface{} {
	return map[string]interface{}{
		"fullname":   p.Sprint("Full name"),
		"user":       p.Sprint("Personal"),
		"email":      p.Sprint("Email"),
		"title":      p.Sprint("Olmax Medical | Application"),
		"offer":      p.Sprint("Create A New Offer"),
		"area":       p.Sprint("Location and specialties"),
		"doccountry": p.Sprint("What country or countries do you practice medicine?"),
		"docspecial": p.Sprint("What is your specialty, or specialties?"),
		"gender":     p.Sprint("What is your biological gender?"),
		"male":       p.Sprint("Male"),
		"female":     p.Sprint("Female"),
		"documents":  p.Sprint("Documents"),
		"cv":         p.Sprint("Please copy and paste your CV:"),
		"diploma":    p.Sprint("Please upload a copy of your medical school diploma and copy of your board certification:"),
		"complete":   p.Sprint("Please complete the following"),
		"yes":        p.Sprint("Yes"),
		"no":         p.Sprint("No"),
		"q1":         p.Sprint("Has your license to practive medicine in any jurisdiction, your Drug Enforcement Administration registration or any applicable narcotic registration in any jurisdiction ever been denied, limited, restricted, suspended, revoked, not renewed, or subject to probationary conditions, or have you voluntarily or involuntarily relinquished any such license or registration or voluntarily or involuntarily accepted any such actions or conditions, or have you ever been fined or received a letter of reprimand or is such action pending?"),
		"q2":         p.Sprint("Have you ever been charged, suspended, fined, diciplined, or otherwise sanctioned, subject to probationary conditions, restricted or excluded, or have you voluntarily or involuntarily relinquished eligibility to provide services or accepted conditions on your eligibility to provide services by any form of medical insurance entity, or any public program, or is any such action pending?"),
		"q3":         p.Sprint("Have your clinical privileges, membership, contractual participation or employment by any medical organization (e.g. hospital medical staff, medical group, independent practice association, health plan, health mantenance organization, preferred provider organization, private payer, medical society, professional association, medical school faculty position or other health delivery entity or system), ever been denied, suspended, restricted, reduced, subject to probationairy conditions, revoked or not renewed for any reason, or is any such action pending?"),
		"q4":         p.Sprint("Have you ever voluntarily or involuntarily withdrawn a request for membership or clinical privileges, voluntarily or involuntarily relinquished membership or clinical privileges, allowed such membership or clinical privileges to expire, or resigned from any medical organization (e.g., hospital medical staff, medical group, independent practice association, health plan, health mantenance organization, preferred provider organization, private payer, medical society, professional association, medical school faculty position or other health delivery entity or system), for any reason, or in return for such an investigation not being conducted, or is any such action pending?"),
		"q5":         p.Sprint("Have you ever terminated contractual participation or employment with any medical organization (e.g., hospital medical staff, medical group, independent practice association, health plan, health mantenance organization, preferred provider organization, private payer, medical society, professional association, medical school faculty position or other health delivery entity or system), or allowed such contract or employment to expire, for any reason, or in return for such an investigation not being conducted, or is any such action pending?"),
		"q6":         p.Sprint("Have you ever surrendered, voluntarily withdrawn, or been requested or compelled to relinquish your status as a student in good standing in any internship, residency, fellowship, preceptorship, or other clinical education program, or is any such action pending?"),
		"q7":         p.Sprint("Has your membership or fellowship in any local, county, state, regional, nation, or international professional organization ever been revoked, denied, reduced, limited, subject to probationary conditions, or not renewed, or is any such action pending?"),
		"q8":         p.Sprint("Have you ever been denied certification/re-certification by a specialty board, or has your eligibility certification or re-certification status changed, or is any such action pending?"),
		"q9":         p.Sprint("Have you ever been convicted of or pled <i>nolo contendre</i> to any crime (other than a minor traffic violation), or is any such action pending?"),
		"q10":        p.Sprint("Do you currently use drugs illegally, or in the past five years have you participated in any treatment or diversion program related to drug use, alcohol dependency, or psychiatric problems?"),
		"q11":        p.Sprint("Have any judgments been entered against you, or settlements been agreed to by you (including dismissals) within the last ten (10) years, in professional liability cases, or are there any filed and served professional liability lawsuits/arbitrations against you pending?"),
		"q12":        p.Sprint("I hereby affirm that the information above and any addenda thereto is true, current, correct, and complete to the best of my knowledge and belief and is furnished in good faith, I understand that omissions or misrepresentations may result in denial of my application to participate in Olmax Medical web services."),
		"confirm":    p.Sprint("Confirm"),
		"submit":     p.Sprint("Submit"),
	}
}
