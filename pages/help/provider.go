package help

import (
	"golang.org/x/text/message"
	"olmax/router"
)

func init() {
	b := &router.Page{
		Access: router.GuestAuth|router.PatientAuth|router.DoctorAuth,
		Css:    "",
		Path:   "help/provider",
		Data:   Provider,
		Extra:  router.ListServices,
	}
	router.Add(b)
}

func Provider(p *message.Printer) map[string]interface{} {
	return map[string]interface{}{
		"title":              p.Sprintf("Olmax Medical | Become A Provider"),
		"specialtyHeader":    p.Sprintf("Specialty"),
		"mainHeader":         p.Sprintf("Become A Provider"),
		"earn":               p.Sprintf("Find out what you could earn"),
		"getStartedHeader":   p.Sprintf("Get Started"),
		"getStartedlink":     p.Sprintf("Get Started"),
		"providerWhy":        p.Sprintf("Why become a provider on Olmax?"),
		"whyText":            p.Sprintf("No matter what your specialty, Olmax makes it simple and secure to reach millions of patients looking for doctors with unique skills and specialties, just like yours."),
		"control":            p.Sprintf("You're in control"),
		"controlText":        p.Sprintf("With Olmax, you're in full control of your availability, prices, medical management, and how you interact with patients. You can set appointment times and handle the process however you like."),
		"everyStep":          p.Sprintf("We're there at every step"),
		"everyStepText":      p.Sprintf("Olmax offers tools, service tips, 24/7 support, and an on-line community of experienced physicians for questions and sharing ideas for success."),
		"provider":           p.Sprintf("How to become an Olmax Provider"),
		"createProfile":      p.Sprintf("Create your profile"),
		"createProfileText":  p.Sprintf("It's free and easy to create a profile on Olmax. Describe your resume, how many patients you can accomodate, set your own times, and add photos and details about yourself."),
		"welcomePatient":     p.Sprintf("Welcome patients"),
		"welcomePatientText": p.Sprintf("Communicate with patients via 3rd party applications, or personal telephone."),
		"getPaid":            p.Sprintf("Get Paid"),
		"getPaidText":        p.Sprintf("Olmax's secure payment system means you will never see a patient without compensation, or have to deal with money directly. Patienst are charged before appointments, and you are paid after the visit is completed. We embrace the future, therefore payments will be via Bitcoin only."),
		"safety":             p.Sprintf("Safety on Olmax"),
		"trust":              p.Sprintf("Olmax is built on trust"),
		"trustText":          p.Sprintf("All Olmax physicians must: submit a profile photo, medical diplomas, residency certification or equivalent, verify their phone, email, government ID, and background checks. Patients and physicians can each publish reviews after visit conclusions keeping everyone accountable and respectful."),
		"seeingPatients":     p.Sprintf("Start seeing patients"),
	}
}
