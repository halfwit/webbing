package pages

import (
	"github.com/olmaxmedical/olmax_go/plugins"
	"github.com/olmaxmedical/olmax_go/router"
	"golang.org/x/text/message"
)

func init() {
	b := &router.Page{
		Access: router.GuestAuth | router.DoctorAuth | router.PatientAuth,
		CSS:    "index.css",
		Path:   "index",
		Data:   Index,
		Extra:  plugins.ListDoctors | plugins.FormErrors,
	}
	router.AddPage(b)
}

// Index - olmaxmedical.com/index.html
func Index(p *message.Printer) map[string]interface{} {
	return map[string]interface{}{
		"title":            p.Sprintf("Olmax Medical | Welcome"),
		"name":             p.Sprintf("Olmax Medical"),
		"countryText":      p.Sprintf("Country"),
		"specialty":        p.Sprintf("Specialty"),
		"findDoctor":       p.Sprintf("Find a Doctor"),
		"appointmentDates": p.Sprintf("Appointment Dates:"),
		"from":             p.Sprintf("From:"),
		"to":               p.Sprintf("To:"),
		"appointmentTimes": p.Sprintf("Appointment Times"),
		"startTime":        p.Sprintf("Start Time:"),
		"endTime":          p.Sprintf("End Time:"),
		"search":           p.Sprintf("Search"),
		"whoWeAre":         p.Sprintf("Who We Are"),
		"aboutUs":          p.Sprintf("Olmax Medical is a world wide network of physicians and patients that enables them to communicate, meet virtually, and negotiate payment on a peer to peer basis, without the interference of insurance giants. We provide a platform where the economics of <i>laissez-fairedes</i> (free-trade) will allow both physicians and patients to negotiate fee for service. Our website provide a platform where both patients and doctors form around the world can deploy customized contracts describing, in detail, the terms of health care. The cost, time, and duration of virtual clinic visits will be pre-determined on contracts posted on our website. The contracts are written by either doctor or patient. Contracts can be created, bought, and sold by anyone, because we believe health care should be available to everyone. It will be our work to investigate and verify physician status. Once doctors are verified, patients will have the opportunity to rate physician performance and bedside manners."),
		"secondOpinions":   p.Sprintf("Second Opinions"),
		"fromHome":         p.Sprintf("Since the time of Hippocrates, patients and doctors were limited to serving and receiving care from physician in their more local community. With our platform patients will not be tied to HMOs or managed health care. In other words, insurance companies or government decisions will no longer chain patients to the type and quality of health care they receive. Doctors with extremely rare specialties will be able to serve communities thousands of miles away from them, and from the comfort of their home if they so desire"),
		"anonymity":        p.Sprintf("Anonymity"),
		"anonText":         p.Sprintf("Patients will be encouraged to use anonymous names. Medical records are kept between patients and physicians, they are not stored on our servers."),
		"wholeWorld":       p.Sprintf("Access to Physicians from around the world"),
		"wholeWorldText":   p.Sprintf("Physicians from around the world will be able to join our network, see patients from anywhere at anytime."),
		"payment":          p.Sprintf("Payment"),
		"paymentText":      p.Sprintf("Payments will be made with Bitcoin. Minimal fees will be charged by our website for holding the cryptocurrency until clinical visits are complete."),
		"copyright":        p.Sprintf("Copyright 2017, 2018, 2019"),
		"seeMore":          p.Sprintf("See More"),
		"doctorsFrom":      p.Sprintf("What are patients saying about our doctors from"),
	}
}
