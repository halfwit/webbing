// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

package main

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

type dictionary struct {
	index []uint32
	data  string
}

func (d *dictionary) Lookup(key string) (data string, ok bool) {
	p := messageKeyToIndex[key]
	start, end := d.index[p], d.index[p+1]
	if start == end {
		return "", false
	}
	return d.data[start:end], true
}

func init() {
	dict := map[string]catalog.Dictionary{
		"en_US": &dictionary{index: en_USIndex, data: en_USData},
	}
	fallback := language.MustParse("en-US")
	cat, err := catalog.NewFromMap(dict, catalog.Fallback(fallback))
	if err != nil {
		panic(err)
	}
	message.DefaultCatalog = cat
}

var messageKeyToIndex = map[string]int{
	"-View all":                             158,
	"0 BTC":                                 85,
	"1.) Go to Edit Profile.":               114,
	"1.) Go to your Notification Settings.": 119,
	"1.) Remove Olmax messages from your spam list":                                                                  126,
	"2.) Add nemo@olmax.com, halfwit@olmax.com, and services@olmax.com to your personal email address book":          127,
	"2.) Look for the \"Email Settings\" field. Make sure you have chosen the email types you want to receive.":      120,
	"2.) Look for the Email Address field. Make sure your address is correct.":                                       115,
	"3.) After adding or removing checkmarks from the right boxes, scroll to the bottom of the page and click Save.": 121,
	"3.) If it is incorrect, add the correct address and click Save.":                                                116,
	"Access to Physicians from around the world":                                                                     55,
	"Acute Pain Medicine": 190,
	"All Olmax physicians must: submit a profile photo, medical diplomas, residency certification or equivalent, verify their phone, email, government ID, and background checks. Patients and physicians can each publish reviews after visit conclusions keeping everyone accountable and respectful.": 184,
	"All payments will be done via ":                       159,
	"Anesthesiology":                                       191,
	"Anonymity":                                            53,
	"Any changes in vision?":                               26,
	"Any fevers or Chills?":                                24,
	"Any heart problems?":                                  28,
	"Any intestinal problems?":                             29,
	"Any kidney problems?":                                 30,
	"Any lung issues?":                                     27,
	"Any nervous system problmes? <i>Strokes</i>":          32,
	"Any problems with muscles or bones?":                  31,
	"Any psychiatric problems? <i>Depression, anxiety</i>": 33,
	"Any weight gain or weight loss?":                      25,
	"Appointment Dates:":                                   42,
	"Appointment Requests":                                 104,
	"Appointment Times":                                    45,
	"Appointments":                                         96,
	"Available patients":                                   226,
	"Bariatric Surgery":                                    192,
	"Become A Partner":                                     100,
	"Become A Provider":                                    166,
	"Bitcoin must be paid in full upon deployment or acceptance of contract.": 138,
	"Can I pay with any currency?":                                            139,
	"Cardiology":                                                              193,
	"Check your email notification settings":                                  117,
	"Check your spam and other email filters":                                 124,
	"Chiropractics":                                                           194,
	"Chronic Pain":                                                            195,
	"Communicate with patients via 3rd party applications, or personal telephone.": 179,
	"Contacting A Physician":     155,
	"Contacting Physician":       133,
	"Copyright 2017, 2018, 2019": 59,
	"Country":                    39,
	"Create Patient Profile":     12,
	"Create your profile":        176,
	"Critical Care":              196,
	"Depending on your provider, emails can take up to a few hours to be delivered. If undelivered or delayed emails continue to be an issue, check with your provider to see if there are any configuration issues or problems with their network that might be affecting your account.": 129,
	"Deposit Funds": 87,
	"Dermatology":   197,
	"Do I need to pay before scheduling an appointment?":             135,
	"Does your pain travel or radiate to another part of your body?": 22,
	"Ear Nose and Throat":          200,
	"Email:":                       64,
	"Emergency Medicine":           198,
	"End Time:":                    47,
	"Endocrinology":                199,
	"Enter Email":                  71,
	"Enter a valid email":          80,
	"Enter password (8+ chars)":    81,
	"Enter your first name":        77,
	"Enter your last name":         78,
	"FAQ":                          90,
	"Family Medicine":              201,
	"Female":                       35,
	"Find a Doctor":                41,
	"Find out what you could earn": 167,
	"First Name:":                  76,
	"First name must be at least 2 characters": 232,
	"First name required":                      231,
	"Forgot your password?":                    66,
	"From:":                                    43,
	"Full name must be at least 2 characters":  5,
	"Full name required":                       4,
	"Gastrointestinology":                      202,
	"Get Paid":                                 180,
	"Get Started":                              168,
	"Have you taken any medications for these symptoms and how much have they worked?": 23,
	"Head and Neck":           203,
	"Hello ":                  10,
	"Help":                    91,
	"Hematology and Oncology": 204,
	"Hepatology":              205,
	"How can I add another appointment or business address to my receipt?": 149,
	"How can we help?":                                                                   143,
	"How do I edit or remove a payment method?":                                          146,
	"How do I make an appointment on Olmax?":                                             157,
	"How do I use Bitcoin to pay?":                                                       148,
	"How do I verify my phone number?":                                                   150,
	"How is the price determined for my appointment?":                                    154,
	"How long have these symptoms lasted?":                                               19,
	"How to become an Olmax Provider":                                                    175,
	"How would you characterize your symptoms? <i>Sharp, Dull, Ache, twisting, ets.</i>": 20,
	"Hyperbaric": 206,
	"I am a patient. How do I check the status of my appointment?": 105,
	"I did not recieve an email confirming nor denying my request": 111,
	"If you do not recieve a confimation email by 12 hrs, then a full refund will be returned to your bitcoin account along with an email stating that an appointment could not be made":                108,
	"If you have other filters or routing rules in your email account that may have sorted Olmax emails elsewhere, be sure to check those, too.</br>Check for issues with your email service provider.": 128,
	"If you have submitted payment, and do not see appointment scheduled on this page; please refer to the %s section.":                                                                                 8,
	"If you recieve an email confirming a cancelation of decline, bitcoin will be returned to your account infull, then you may seek another appointment":                                               110,
	"If your country blocks Olmax Medical?": 130,
	"Immunology":                            207,
	"Infectious Diseases":                   208,
	"Internal Medicine":                     209,
	"Invalid email":                         3,
	"Invalid selection for %s":              102,
	"Invalid selection for question %d":     1,
	"It's free and easy to create a profile on Olmax. Describe your resume, how many patients you can accomodate, set your own times, and add photos and details about yourself.": 177,
	"It's possible your email provider mistakenly sent our messages to your spam or junk folder. To avoid this:":                                                                  125,
	"Last Name:": 79,
	"Last name must be at least 2 characters": 234,
	"Last name required":                      233,
	"Legal":                                   97,
	"Login":                                   67,
	"Make sure your email address is correct": 112,
	"Male":                              34,
	"NO FUNDS CURRENTLY HELD IN ESCROW": 86,
	"Neonatology":                       210,
	"Nephrology":                        211,
	"Neurology":                         212,
	"Neurosurgery":                      213,
	"No":                                164,
	"No matter what your specialty, Olmax makes it simple and secure to reach millions of patients looking for doctors with unique skills and specialties, just like yours.": 170,
	"No selection for %s":          101,
	"No selection for question %d": 0,
	"No.":                          140,
	"Obstetrics and Gynecology":    214,
	"Occupational Medicine":        215,
	"Olmax Medical":                38,
	"Olmax Medical is a world wide network of physicians and patients that enables them to communicate, meet virtually, and negotiate payment on a peer to peer basis, without the interference of insurance giants. We provide a platform where the economics of <i>laissez-fairedes</i> (free-trade) will allow both physicians and patients to negotiate fee for service. Our website provide a platform where both patients and doctors form around the world can deploy customized contracts describing, in detail, the terms of health care. The cost, time, and duration of virtual clinic visits will be pre-determined on contracts posted on our website. The contracts are written by either doctor or patient. Contracts can be created, bought, and sold by anyone, because we believe health care should be available to everyone. It will be our work to investigate and verify physician status. Once doctors are verified, patients will have the opportunity to rate physician performance and bedside manners.": 50,
	"Olmax Medical | Appointments":      6,
	"Olmax Medical | Become A Provider": 165,
	"Olmax Medical | Bookings":          225,
	"Olmax Medical | Create Profile":    11,
	"Olmax Medical | FAQ":               103,
	"Olmax Medical | Find Patients":     227,
	"Olmax Medical | Help":              142,
	"Olmax Medical | Login":             62,
	"Olmax Medical | Messages":          68,
	"Olmax Medical | Our Doctors":       132,
	"Olmax Medical | Profile":           9,
	"Olmax Medical | Sign Up":           74,
	"Olmax Medical | Wallet":            83,
	"Olmax Medical | Welcome":           37,
	"Olmax is built on trust":           183,
	"Olmax offers tools, service tips, 24/7 support, and an on-line community of experienced physicians for questions and sharing ideas for success.":                                                                                                                                          174,
	"Olmax's secure payment system means you will never see a patient without compensation, or have to deal with money directly. Patienst are charged before appointments, and you are paid after the visit is completed. We embrace the future, therefore payments will be via Bitcoin only.": 181,
	"Once an appointment request is submitted, the physician has 4 to 12 hrs to replay. Depending on the urgency. If you would like a reply within 4 hr for urgent consults, an extra fee can be payed. Otherwise doctors have 12 hrs to reply to appointment request.":                        106,
	"Once you have submitted a phone number, you can either receive a text message or call with a confirmation number.":                                                                                                                                                                        187,
	"Opthamology":                            216,
	"Orthopedic Surgery":                     217,
	"Palliative Care":                        218,
	"Password must be at least 8 characters": 229,
	"Password required":                      228,
	"Password:":                              65,
	"Patients will be encouraged to use anonymous names. Medical records are kept between patients and physicians, they are not stored on our servers.": 54,
	"Payment":         57,
	"Payment Methods": 93,
	"Payments will be made with Bitcoin. Minimal fees will be charged by our website for holding the cryptocurrency until clinical visits are complete.": 58,
	"Pediatrics": 219,
	"Physicians from around the world will be able to join our network, see patients from anywhere at anytime.": 56,
	"Please check any of the following if you have experienced in the last 6 weeks:":                            36,
	"Please click the following link to finalize your account creation ":                                        238,
	"Please click the following link to reset your password ":                                                   236,
	"Please give a brief statement regarding the main reason you would like to see your doctor:":                16,
	"Please refer to the following help page: ":                                                                 189,
	"Please submit some information regarding your consult.":                                                    13,
	"Podiatry":                      220,
	"Previous messages: Click here": 70,
	"Prices & Fees":                 153,
	"Prices and Fees":               94,
	"Prices and fees are for the most part determined by doctors and patients.": 162,
	"Prices are set by who deploys the contract (doctor or patient). Fees are structured in a amount of bitcoin (BTC) per unit. In which a unit of time equals 15 mins. All new consults must be a minimum of 2 units, and repeat visits can be 1 unit.": 163,
	"Privacy Policy": 98,
	"Profile information such as government ID, diplomas, phone numbers, and emails will will be verified before being posted on Olmax Medical website.": 186,
	"Pulmonology":            221,
	"Quality Healthcare":     92,
	"Radiaton Oncology":      223,
	"Radiology":              222,
	"Re-enter same password": 230,
	"Reset":                  73,
	"Safety on Olmax":        182,
	"Search":                 48,
	"Search all messages in your email inbox": 122,
	"Second Opinions":                         51,
	"See More":                                60,
	"Send only Bitcoin (BTC) to this address": 88,
	"Sending any other digital asset, including Bitcoin Cash (BCH), will result in permanent loss.": 89,
	"Should I clear my schedule if I have no heard back from my doctor?":                            107,
	"Should I clear my schedule if I have not heard back from the doctor?":                          145,
	"Sign Up":          82,
	"Sign Up for free": 75,
	"Since the time of Hippocrates, patients and doctors were limited to serving and receiving care from physician in their more local community. With our platform patients will not be tied to HMOs or managed health care. In other words, insurance companies or government decisions will no longer chain patients to the type and quality of health care they receive. Doctors with extremely rare specialties will be able to serve communities thousands of miles away from them, and from the comfort of their home if they so desire": 52,
	"Sometimes emails can get lost in your inbox. In your email account, search for terms like \"Olmax Medical\", \"Appointment\", \"Verification\", or other words related to the email you're looking for.": 123,
	"Specialty":             40,
	"Start Time:":           46,
	"Start seeing patients": 185,
	"Subject: Olmax Medical - Reset Your Password\n\n":     235,
	"Subject: Olmax Medical - Verify your new account\n\n": 237,
	"Suggested Topics":     144,
	"To:":                  44,
	"Transplant Surgery":   224,
	"Valid email required": 2,
	"Verification":         95,
	"Wallet":               84,
	"We may be sending emails to an old or incorrect email address. To see or change the email address associated with your account, log in to your Olmax account from a desktop computer and follow the steps below:": 113,
	"We will send a reset code to the email provided": 72,
	"We'll only send the emails you tell us you want. To check your email notification settings, log in to your Olmax account from a desktop computer and follow the steps below:": 118,
	"We're there at every step": 173,
	"Welcome back!":             63,
	"Welcome patients":          178,
	"What are patients saying about our doctors from":                61,
	"What does each appointment status mean?":                        156,
	"What happens if my appointment request is declined or expires?": 109,
	"What is Bitcoin?":                                            147,
	"What is Bitcoin? ":                                           160,
	"What is a Verified Medical License?":                         152,
	"What is your biological gender?":                             15,
	"What makes your symptoms better, and What makes them worse?": 21,
	"When did your symptoms start?":                               17,
	"When were you born?":                                         14,
	"When will I be charged?":                                     137,
	"Where are your symptoms located? <i>part of your body</i>":   18,
	"Who We Are":                              49,
	"Why become a provider on Olmax?":         169,
	"Why did I not get a notification email?": 188,
	"Why didn't I get my email notification?": 151,
	"With Olmax, you're in full control of your availability, prices, medical management, and how you interact with patients. You can set appointment times and handle the process however you like.": 172,
	"Work with us!": 99,
	"Yes, you must submit payment in order to secure appointment contract. Your payment will be held in escrow until the visit is finalized. Once you submit fees, we will contact the physician and give him or her your medical information. The doctor will then confirm appointment, and an email or text will be sent to you, along with the physicians contact information. Fees are structured in amount of bitcoin (BTC) per unit(U) time (BTC/U). Every unit (U) is equivalent to 15 min, time spent in visit will be pre-determined, and visits going longer that what was agreed upon will not cost extra. All new consults must be a minimum of 2 units, and repeat visits can be a minimum of 1 unit.": 136,
	"You can bypass their firewall using tunnel software such as a VPN or Tor software. See the following for more information: ":   161,
	"You can bypass their firewall using tunnel software such as a VPN,  or Tor software. See the following for more information: ": 131,
	"You can bypass their firewall using tunnel software such as a VPN, or Tor software. See the following for more information: ":  141,
	"You currently have no appointments pending.":                                    7,
	"You currently have no messages.":                                                69,
	"You may make contact with your doctor as soon as the appointment is confirmed.": 134,
	"You're in control": 171,
}

var en_USIndex = []uint32{ // 240 elements
	// Entry 0 - 1F
	0x00000000, 0x00000020, 0x00000045, 0x0000005a,
	0x00000068, 0x0000007b, 0x000000a3, 0x000000c0,
	0x000000ec, 0x00000161, 0x00000179, 0x00000184,
	0x000001a3, 0x000001ba, 0x000001f1, 0x00000205,
	0x00000225, 0x00000280, 0x0000029e, 0x000002d8,
	0x000002fd, 0x00000350, 0x0000038c, 0x000003cb,
	0x0000041c, 0x00000432, 0x00000452, 0x00000469,
	0x0000047a, 0x0000048e, 0x000004a7, 0x000004bc,
	// Entry 20 - 3F
	0x000004e0, 0x0000050c, 0x00000541, 0x00000546,
	0x0000054d, 0x0000059c, 0x000005b4, 0x000005c2,
	0x000005ca, 0x000005d4, 0x000005e2, 0x000005f5,
	0x000005fb, 0x000005ff, 0x00000611, 0x0000061d,
	0x00000627, 0x0000062e, 0x00000639, 0x00000a16,
	0x00000a26, 0x00000c30, 0x00000c3a, 0x00000ccc,
	0x00000cf7, 0x00000d61, 0x00000d69, 0x00000dfc,
	0x00000e17, 0x00000e20, 0x00000e50, 0x00000e66,
	// Entry 40 - 5F
	0x00000e74, 0x00000e7b, 0x00000e85, 0x00000e9b,
	0x00000ea1, 0x00000eba, 0x00000eda, 0x00000ef8,
	0x00000f04, 0x00000f34, 0x00000f3a, 0x00000f52,
	0x00000f63, 0x00000f6f, 0x00000f85, 0x00000f9a,
	0x00000fa5, 0x00000fb9, 0x00000fd3, 0x00000fdb,
	0x00000ff2, 0x00000ff9, 0x00000fff, 0x00001021,
	0x0000102f, 0x00001057, 0x000010b5, 0x000010b9,
	0x000010be, 0x000010d1, 0x000010e1, 0x000010f1,
	// Entry 60 - 7F
	0x000010fe, 0x0000110b, 0x00001111, 0x00001120,
	0x0000112e, 0x0000113f, 0x00001156, 0x00001172,
	0x00001186, 0x0000119b, 0x000011d8, 0x000012da,
	0x0000131d, 0x000013d0, 0x0000140f, 0x000014a3,
	0x000014e0, 0x00001508, 0x000015d9, 0x000015f1,
	0x0000163a, 0x0000167a, 0x000016a1, 0x0000174e,
	0x00001774, 0x000017dc, 0x0000184b, 0x00001873,
	0x00001935, 0x0000195d, 0x000019c8, 0x000019f6,
	// Entry 80 - 9F
	0x00001a5c, 0x00001b1e, 0x00001c32, 0x00001c58,
	0x00001cda, 0x00001cf6, 0x00001d0b, 0x00001d5a,
	0x00001d8d, 0x0000203b, 0x00002053, 0x0000209b,
	0x000020b8, 0x000020bc, 0x0000213d, 0x00002152,
	0x00002163, 0x00002174, 0x000021b9, 0x000021e3,
	0x000021f4, 0x00002211, 0x00002256, 0x00002277,
	0x0000229f, 0x000022c3, 0x000022d1, 0x00002301,
	0x00002318, 0x00002340, 0x00002367, 0x00002371,
	// Entry A0 - BF
	0x00002394, 0x000023aa, 0x0000242a, 0x00002474,
	0x00002567, 0x0000256a, 0x0000258c, 0x0000259e,
	0x000025bb, 0x000025c7, 0x000025e7, 0x0000268e,
	0x000026a0, 0x00002760, 0x0000277a, 0x0000280a,
	0x0000282a, 0x0000283e, 0x000028ea, 0x000028fb,
	0x00002948, 0x00002951, 0x00002a6a, 0x00002a7a,
	0x00002a92, 0x00002bb5, 0x00002bcb, 0x00002c5e,
	0x00002cd0, 0x00002cf8, 0x00002d26, 0x00002d3a,
	// Entry C0 - DF
	0x00002d49, 0x00002d5b, 0x00002d66, 0x00002d74,
	0x00002d81, 0x00002d8f, 0x00002d9b, 0x00002dae,
	0x00002dbc, 0x00002dd0, 0x00002de0, 0x00002df4,
	0x00002e02, 0x00002e1a, 0x00002e25, 0x00002e30,
	0x00002e3b, 0x00002e4f, 0x00002e61, 0x00002e6d,
	0x00002e78, 0x00002e82, 0x00002e8f, 0x00002ea9,
	0x00002ebf, 0x00002ecb, 0x00002ede, 0x00002eee,
	0x00002ef9, 0x00002f02, 0x00002f0e, 0x00002f18,
	// Entry E0 - FF
	0x00002f2a, 0x00002f3d, 0x00002f56, 0x00002f69,
	0x00002f87, 0x00002f99, 0x00002fc0, 0x00002fd7,
	0x00002feb, 0x00003014, 0x00003027, 0x0000304f,
	0x00003082, 0x000030be, 0x000030f5, 0x0000313c,
} // Size: 984 bytes

const en_USData string = "" + // Size: 12604 bytes
	"\x02No selection for question %[1]d\x02Invalid selection for question %[" +
	"1]d\x02Valid email required\x02Invalid email\x02Full name required\x02Fu" +
	"ll name must be at least 2 characters\x02Olmax Medical | Appointments" +
	"\x02You currently have no appointments pending.\x02If you have submitted" +
	" payment, and do not see appointment scheduled on this page; please refe" +
	"r to the %[1]s section.\x02Olmax Medical | Profile\x04\x00\x01 \x06\x02H" +
	"ello\x02Olmax Medical | Create Profile\x02Create Patient Profile\x02Plea" +
	"se submit some information regarding your consult.\x02When were you born" +
	"?\x02What is your biological gender?\x02Please give a brief statement re" +
	"garding the main reason you would like to see your doctor:\x02When did y" +
	"our symptoms start?\x02Where are your symptoms located? <i>part of your " +
	"body</i>\x02How long have these symptoms lasted?\x02How would you charac" +
	"terize your symptoms? <i>Sharp, Dull, Ache, twisting, ets.</i>\x02What m" +
	"akes your symptoms better, and What makes them worse?\x02Does your pain " +
	"travel or radiate to another part of your body?\x02Have you taken any me" +
	"dications for these symptoms and how much have they worked?\x02Any fever" +
	"s or Chills?\x02Any weight gain or weight loss?\x02Any changes in vision" +
	"?\x02Any lung issues?\x02Any heart problems?\x02Any intestinal problems?" +
	"\x02Any kidney problems?\x02Any problems with muscles or bones?\x02Any n" +
	"ervous system problmes? <i>Strokes</i>\x02Any psychiatric problems? <i>D" +
	"epression, anxiety</i>\x02Male\x02Female\x02Please check any of the foll" +
	"owing if you have experienced in the last 6 weeks:\x02Olmax Medical | We" +
	"lcome\x02Olmax Medical\x02Country\x02Specialty\x02Find a Doctor\x02Appoi" +
	"ntment Dates:\x02From:\x02To:\x02Appointment Times\x02Start Time:\x02End" +
	" Time:\x02Search\x02Who We Are\x02Olmax Medical is a world wide network " +
	"of physicians and patients that enables them to communicate, meet virtua" +
	"lly, and negotiate payment on a peer to peer basis, without the interfer" +
	"ence of insurance giants. We provide a platform where the economics of <" +
	"i>laissez-fairedes</i> (free-trade) will allow both physicians and patie" +
	"nts to negotiate fee for service. Our website provide a platform where b" +
	"oth patients and doctors form around the world can deploy customized con" +
	"tracts describing, in detail, the terms of health care. The cost, time, " +
	"and duration of virtual clinic visits will be pre-determined on contract" +
	"s posted on our website. The contracts are written by either doctor or p" +
	"atient. Contracts can be created, bought, and sold by anyone, because we" +
	" believe health care should be available to everyone. It will be our wor" +
	"k to investigate and verify physician status. Once doctors are verified," +
	" patients will have the opportunity to rate physician performance and be" +
	"dside manners.\x02Second Opinions\x02Since the time of Hippocrates, pati" +
	"ents and doctors were limited to serving and receiving care from physici" +
	"an in their more local community. With our platform patients will not be" +
	" tied to HMOs or managed health care. In other words, insurance companie" +
	"s or government decisions will no longer chain patients to the type and " +
	"quality of health care they receive. Doctors with extremely rare special" +
	"ties will be able to serve communities thousands of miles away from them" +
	", and from the comfort of their home if they so desire\x02Anonymity\x02P" +
	"atients will be encouraged to use anonymous names. Medical records are k" +
	"ept between patients and physicians, they are not stored on our servers." +
	"\x02Access to Physicians from around the world\x02Physicians from around" +
	" the world will be able to join our network, see patients from anywhere " +
	"at anytime.\x02Payment\x02Payments will be made with Bitcoin. Minimal fe" +
	"es will be charged by our website for holding the cryptocurrency until c" +
	"linical visits are complete.\x02Copyright 2017, 2018, 2019\x02See More" +
	"\x02What are patients saying about our doctors from\x02Olmax Medical | L" +
	"ogin\x02Welcome back!\x02Email:\x02Password:\x02Forgot your password?" +
	"\x02Login\x02Olmax Medical | Messages\x02You currently have no messages." +
	"\x02Previous messages: Click here\x02Enter Email\x02We will send a reset" +
	" code to the email provided\x02Reset\x02Olmax Medical | Sign Up\x02Sign " +
	"Up for free\x02First Name:\x02Enter your first name\x02Enter your last n" +
	"ame\x02Last Name:\x02Enter a valid email\x02Enter password (8+ chars)" +
	"\x02Sign Up\x02Olmax Medical | Wallet\x02Wallet\x020 BTC\x02NO FUNDS CUR" +
	"RENTLY HELD IN ESCROW\x02Deposit Funds\x02Send only Bitcoin (BTC) to thi" +
	"s address\x02Sending any other digital asset, including Bitcoin Cash (BC" +
	"H), will result in permanent loss.\x02FAQ\x02Help\x02Quality Healthcare" +
	"\x02Payment Methods\x02Prices and Fees\x02Verification\x02Appointments" +
	"\x02Legal\x02Privacy Policy\x02Work with us!\x02Become A Partner\x02No s" +
	"election for %[1]s\x02Invalid selection for %[1]s\x02Olmax Medical | FAQ" +
	"\x02Appointment Requests\x02I am a patient. How do I check the status of" +
	" my appointment?\x02Once an appointment request is submitted, the physic" +
	"ian has 4 to 12 hrs to replay. Depending on the urgency. If you would li" +
	"ke a reply within 4 hr for urgent consults, an extra fee can be payed. O" +
	"therwise doctors have 12 hrs to reply to appointment request.\x02Should " +
	"I clear my schedule if I have no heard back from my doctor?\x02If you do" +
	" not recieve a confimation email by 12 hrs, then a full refund will be r" +
	"eturned to your bitcoin account along with an email stating that an appo" +
	"intment could not be made\x02What happens if my appointment request is d" +
	"eclined or expires?\x02If you recieve an email confirming a cancelation " +
	"of decline, bitcoin will be returned to your account infull, then you ma" +
	"y seek another appointment\x02I did not recieve an email confirming nor " +
	"denying my request\x02Make sure your email address is correct\x02We may " +
	"be sending emails to an old or incorrect email address. To see or change" +
	" the email address associated with your account, log in to your Olmax ac" +
	"count from a desktop computer and follow the steps below:\x021.) Go to E" +
	"dit Profile.\x022.) Look for the Email Address field. Make sure your add" +
	"ress is correct.\x023.) If it is incorrect, add the correct address and " +
	"click Save.\x02Check your email notification settings\x02We'll only send" +
	" the emails you tell us you want. To check your email notification setti" +
	"ngs, log in to your Olmax account from a desktop computer and follow the" +
	" steps below:\x021.) Go to your Notification Settings.\x022.) Look for t" +
	"he \x22Email Settings\x22 field. Make sure you have chosen the email typ" +
	"es you want to receive.\x023.) After adding or removing checkmarks from " +
	"the right boxes, scroll to the bottom of the page and click Save.\x02Sea" +
	"rch all messages in your email inbox\x02Sometimes emails can get lost in" +
	" your inbox. In your email account, search for terms like \x22Olmax Medi" +
	"cal\x22, \x22Appointment\x22, \x22Verification\x22, or other words relat" +
	"ed to the email you're looking for.\x02Check your spam and other email f" +
	"ilters\x02It's possible your email provider mistakenly sent our messages" +
	" to your spam or junk folder. To avoid this:\x021.) Remove Olmax message" +
	"s from your spam list\x022.) Add nemo@olmax.com, halfwit@olmax.com, and " +
	"services@olmax.com to your personal email address book\x02If you have ot" +
	"her filters or routing rules in your email account that may have sorted " +
	"Olmax emails elsewhere, be sure to check those, too.</br>Check for issue" +
	"s with your email service provider.\x02Depending on your provider, email" +
	"s can take up to a few hours to be delivered. If undelivered or delayed " +
	"emails continue to be an issue, check with your provider to see if there" +
	" are any configuration issues or problems with their network that might " +
	"be affecting your account.\x02If your country blocks Olmax Medical?\x04" +
	"\x00\x01 }\x02You can bypass their firewall using tunnel software such a" +
	"s a VPN,  or Tor software. See the following for more information:\x02Ol" +
	"max Medical | Our Doctors\x02Contacting Physician\x02You may make contac" +
	"t with your doctor as soon as the appointment is confirmed.\x02Do I need" +
	" to pay before scheduling an appointment?\x02Yes, you must submit paymen" +
	"t in order to secure appointment contract. Your payment will be held in " +
	"escrow until the visit is finalized. Once you submit fees, we will conta" +
	"ct the physician and give him or her your medical information. The docto" +
	"r will then confirm appointment, and an email or text will be sent to yo" +
	"u, along with the physicians contact information. Fees are structured in" +
	" amount of bitcoin (BTC) per unit(U) time (BTC/U). Every unit (U) is equ" +
	"ivalent to 15 min, time spent in visit will be pre-determined, and visit" +
	"s going longer that what was agreed upon will not cost extra. All new co" +
	"nsults must be a minimum of 2 units, and repeat visits can be a minimum " +
	"of 1 unit.\x02When will I be charged?\x02Bitcoin must be paid in full up" +
	"on deployment or acceptance of contract.\x02Can I pay with any currency?" +
	"\x02No.\x04\x00\x01 |\x02You can bypass their firewall using tunnel soft" +
	"ware such as a VPN, or Tor software. See the following for more informat" +
	"ion:\x02Olmax Medical | Help\x02How can we help?\x02Suggested Topics\x02" +
	"Should I clear my schedule if I have not heard back from the doctor?\x02" +
	"How do I edit or remove a payment method?\x02What is Bitcoin?\x02How do " +
	"I use Bitcoin to pay?\x02How can I add another appointment or business a" +
	"ddress to my receipt?\x02How do I verify my phone number?\x02Why didn't " +
	"I get my email notification?\x02What is a Verified Medical License?\x02P" +
	"rices & Fees\x02How is the price determined for my appointment?\x02Conta" +
	"cting A Physician\x02What does each appointment status mean?\x02How do I" +
	" make an appointment on Olmax?\x02-View all\x04\x00\x01 \x1e\x02All paym" +
	"ents will be done via\x04\x00\x01 \x11\x02What is Bitcoin?\x04\x00\x01 {" +
	"\x02You can bypass their firewall using tunnel software such as a VPN or" +
	" Tor software. See the following for more information:\x02Prices and fee" +
	"s are for the most part determined by doctors and patients.\x02Prices ar" +
	"e set by who deploys the contract (doctor or patient). Fees are structur" +
	"ed in a amount of bitcoin (BTC) per unit. In which a unit of time equals" +
	" 15 mins. All new consults must be a minimum of 2 units, and repeat visi" +
	"ts can be 1 unit.\x02No\x02Olmax Medical | Become A Provider\x02Become A" +
	" Provider\x02Find out what you could earn\x02Get Started\x02Why become a" +
	" provider on Olmax?\x02No matter what your specialty, Olmax makes it sim" +
	"ple and secure to reach millions of patients looking for doctors with un" +
	"ique skills and specialties, just like yours.\x02You're in control\x02Wi" +
	"th Olmax, you're in full control of your availability, prices, medical m" +
	"anagement, and how you interact with patients. You can set appointment t" +
	"imes and handle the process however you like.\x02We're there at every st" +
	"ep\x02Olmax offers tools, service tips, 24/7 support, and an on-line com" +
	"munity of experienced physicians for questions and sharing ideas for suc" +
	"cess.\x02How to become an Olmax Provider\x02Create your profile\x02It's " +
	"free and easy to create a profile on Olmax. Describe your resume, how ma" +
	"ny patients you can accomodate, set your own times, and add photos and d" +
	"etails about yourself.\x02Welcome patients\x02Communicate with patients " +
	"via 3rd party applications, or personal telephone.\x02Get Paid\x02Olmax'" +
	"s secure payment system means you will never see a patient without compe" +
	"nsation, or have to deal with money directly. Patienst are charged befor" +
	"e appointments, and you are paid after the visit is completed. We embrac" +
	"e the future, therefore payments will be via Bitcoin only.\x02Safety on " +
	"Olmax\x02Olmax is built on trust\x02All Olmax physicians must: submit a " +
	"profile photo, medical diplomas, residency certification or equivalent, " +
	"verify their phone, email, government ID, and background checks. Patient" +
	"s and physicians can each publish reviews after visit conclusions keepin" +
	"g everyone accountable and respectful.\x02Start seeing patients\x02Profi" +
	"le information such as government ID, diplomas, phone numbers, and email" +
	"s will will be verified before being posted on Olmax Medical website." +
	"\x02Once you have submitted a phone number, you can either receive a tex" +
	"t message or call with a confirmation number.\x02Why did I not get a not" +
	"ification email?\x04\x00\x01 )\x02Please refer to the following help pag" +
	"e:\x02Acute Pain Medicine\x02Anesthesiology\x02Bariatric Surgery\x02Card" +
	"iology\x02Chiropractics\x02Chronic Pain\x02Critical Care\x02Dermatology" +
	"\x02Emergency Medicine\x02Endocrinology\x02Ear Nose and Throat\x02Family" +
	" Medicine\x02Gastrointestinology\x02Head and Neck\x02Hematology and Onco" +
	"logy\x02Hepatology\x02Hyperbaric\x02Immunology\x02Infectious Diseases" +
	"\x02Internal Medicine\x02Neonatology\x02Nephrology\x02Neurology\x02Neuro" +
	"surgery\x02Obstetrics and Gynecology\x02Occupational Medicine\x02Opthamo" +
	"logy\x02Orthopedic Surgery\x02Palliative Care\x02Pediatrics\x02Podiatry" +
	"\x02Pulmonology\x02Radiology\x02Radiaton Oncology\x02Transplant Surgery" +
	"\x02Olmax Medical | Bookings\x02Available patients\x02Olmax Medical | Fi" +
	"nd Patients\x02Password required\x02Password must be at least 8 characte" +
	"rs\x02Re-enter same password\x02First name required\x02First name must b" +
	"e at least 2 characters\x02Last name required\x02Last name must be at le" +
	"ast 2 characters\x04\x00\x02\x0a\x0a-\x02Subject: Olmax Medical - Reset " +
	"Your Password\x04\x00\x01 7\x02Please click the following link to reset " +
	"your password\x04\x00\x02\x0a\x0a1\x02Subject: Olmax Medical - Verify yo" +
	"ur new account\x04\x00\x01 B\x02Please click the following link to final" +
	"ize your account creation"

	// Total table size 13588 bytes (13KiB); checksum: 2C40C1DF
