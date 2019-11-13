package help

import (
	"github.com/olmaxmedical/olmax_go/router"
	"golang.org/x/text/message"
)

func init() {
	b := &router.Page{
		Access: router.GuestAuth | router.PatientAuth | router.DoctorAuth,
		CSS:    "",
		Path:   "help/appointments",
		Data:   appointments,
		Extra:  0,
	}
	router.AddPage(b)
}

func appointments(p *message.Printer) map[string]interface{} {
	return map[string]interface{}{
		"title":             p.Sprintf("Olmax Medical | FAQ"),
		"mainHeader":        p.Sprintf("Olmax Medical"),
		"requestHeader":     p.Sprintf("Appointment Requests"),
		"statusHeader":      p.Sprintf("I am a patient. How do I check the status of my appointment?"),
		"statusBody":        p.Sprintf("Once an appointment request is submitted, the physician has 4 to 12 hrs to replay. Depending on the urgency. If you would like a reply within 4 hr for urgent consults, an extra fee can be payed. Otherwise doctors have 12 hrs to reply to appointment request."),
		"scheduleHeader":    p.Sprintf("Should I clear my schedule if I have no heard back from my doctor?"),
		"scheduleBody":      p.Sprintf("If you do not recieve a confimation email by 12 hrs, then a full refund will be returned to your bitcoin account along with an email stating that an appointment could not be made"),
		"expiresHeader":     p.Sprintf("What happens if my appointment request is declined or expires?"),
		"expiresBody":       p.Sprintf("If you recieve an email confirming a cancelation of decline, bitcoin will be returned to your account infull, then you may seek another appointment"),
		"emailHeader":       p.Sprintf("I did not recieve an email confirming nor denying my request"),
		"emailBody1":        p.Sprintf("Make sure your email address is correct"),
		"emailBody2":        p.Sprintf("We may be sending emails to an old or incorrect email address. To see or change the email address associated with your account, log in to your Olmax account from a desktop computer and follow the steps below:"),
		"emailBody3":        p.Sprintf("1.) Go to Edit Profile."),
		"emailBody4":        p.Sprintf("2.) Look for the Email Address field. Make sure your address is correct."),
		"emailBody5":        p.Sprintf("3.) If it is incorrect, add the correct address and click Save."),
		"notifyHeader":      p.Sprintf("Check your email notification settings"),
		"notifyBody":        p.Sprintf("We'll only send the emails you tell us you want. To check your email notification settings, log in to your Olmax account from a desktop computer and follow the steps below:"),
		"notifyBody1":       p.Sprintf("1.) Go to your Notification Settings."),
		"notifyBody2":       p.Sprintf("2.) Look for the \"Email Settings\" field. Make sure you have chosen the email types you want to receive."),
		"notifyBody3":       p.Sprintf("3.) After adding or removing checkmarks from the right boxes, scroll to the bottom of the page and click Save."),
		"inboxSearchHeader": p.Sprintf("Search all messages in your email inbox"),
		"inboxSearchBody":   p.Sprintf("Sometimes emails can get lost in your inbox. In your email account, search for terms like \"Olmax Medical\", \"Appointment\", \"Verification\", or other words related to the email you're looking for."),
		"checkSpamHeader":   p.Sprintf("Check your spam and other email filters"),
		"checkSpamBody":     p.Sprintf("It's possible your email provider mistakenly sent our messages to your spam or junk folder. To avoid this:"),
		"checkSpamBody1":    p.Sprintf("1.) Remove Olmax messages from your spam list"),
		"checkSpamBody2":    p.Sprintf("2.) Add nemo@olmax.com, halfwit@olmax.com, and services@olmax.com to your personal email address book"),
		"checkSpamAdd":      p.Sprintf("If you have other filters or routing rules in your email account that may have sorted Olmax emails elsewhere, be sure to check those, too.</br>Check for issues with your email service provider."),
		"deliveryHeader":    p.Sprintf("Depending on your provider, emails can take up to a few hours to be delivered. If undelivered or delayed emails continue to be an issue, check with your provider to see if there are any configuration issues or problems with their network that might be affecting your account."),
		"blockHeader":       p.Sprintf("If your country blocks Olmax Medical?"),
		"blockBody":         p.Sprintf("You can bypass their firewall using tunnel software such as a VPN,  or Tor software. See the following for more information: "),
	}
}
