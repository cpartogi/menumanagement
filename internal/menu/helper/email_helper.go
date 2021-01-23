package helper

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

type SendEmail struct {
	FromName   string
	FromEmail  string
	ToName     string
	ToEmail    string
	ReplyTo    string
	Subject    string
	TemplateId string
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func SendInBlueEmail(emailData SendEmail, subData map[string]string) {
	var respBody UrlHttpResponse

	header := map[string]string{
		"accept":  "application/json",
		"api-key": "a",
	}

	formParam := map[string]interface{}{
		"name":       emailData.ToName,
		"email":      emailData.ToEmail,
		"templateId": 1,
	}
	body, _ := json.Marshal(formParam)

	apiCall := APICall{
		URL:       "https://api.sendinblue.com/v3/smtp/email",
		Method:    http.MethodPost,
		FormParam: string(body),
		Header:    header,
	}
	resp, err := apiCall.Call()
	if err == nil {
		json.Unmarshal([]byte(resp.Body), &respBody)
	}

	fmt.Printf("%+v", resp.Body)
	//reqByte, _ := json.Marshal(apiCall)
	//resByte, _ := json.Marshal(resp)

}

func SendVerificationEmail(emailData SendEmail, verificationLink string) {
	var emailBody = fmt.Sprintf(`<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html>

<head>
    <meta charset="UTF-8">
    <meta content="width=device-width, initial-scale=1" name="viewport">
    <meta name="x-apple-disable-message-reformatting">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta content="telephone=no" name="format-detection">
    <title></title>
    <!--[if (mso 16)]>
    <style type="text/css">
    a {text-decoration: none;}
    </style>
    <![endif]-->
    <!--[if gte mso 9]><style>sup { font-size: 100%% !important; }</style><![endif]-->
</head>

<body>
    <div class="es-wrapper-color">
        <!--[if gte mso 9]>
			<v:background xmlns:v="urn:schemas-microsoft-com:vml" fill="t">
				<v:fill type="tile" color="#ffffff"></v:fill>
			</v:background>
		<![endif]-->
        <table class="es-wrapper" width="100%%" cellspacing="0" cellpadding="0">
            <tbody>
                <tr>
                    <td class="esd-email-paddings" valign="top">
                        <table class="es-content esd-header-popover" align="center" cellspacing="0" cellpadding="0">
                            <tbody>
                                <tr>
                                    <td class="esd-stripe" align="center">
                                        <table class="es-content-body" align="center" width="600" cellspacing="0" cellpadding="0" bgcolor="#ffffff">
                                            <tbody>
                                                <tr>
                                                    <td class="esd-structure es-p20t" esd-custom-block-id="8430" align="left">
                                                        <table width="100%%" cellspacing="0" cellpadding="0">
                                                            <tbody>
                                                                <tr>
                                                                    <td class="esd-container-frame" align="center" width="600" valign="top">
                                                                        <table width="100%%" cellspacing="0" cellpadding="0">
                                                                            <tbody>
                                                                                <tr>
                                                                                    <td class="esd-block-image es-p20r es-p20l" align="center" style="font-size:0"><a target="_blank"><img class="adapt-img" src="https://hipffk.stripocdn.email/content/guids/CABINET_57116f9afe83495a646cd7734bc77d26/images/39641523866414281.jpg" alt="Image" style="display: block;" title="Image" width="260"></a></td>
                                                                                </tr>
                                                                            </tbody>
                                                                        </table>
                                                                    </td>
                                                                </tr>
                                                            </tbody>
                                                        </table>
                                                    </td>
                                                </tr>
                                            </tbody>
                                        </table>
                                    </td>
                                </tr>
                            </tbody>
                        </table>
                        <table class="es-content" align="center" cellspacing="0" cellpadding="0">
                            <tbody>
                                <tr>
                                    <td class="esd-stripe" align="center">
                                        <table class="es-content-body" style="border-left:1px solid transparent;border-right:1px solid transparent;border-top:1px solid transparent;border-bottom:1px solid transparent;" align="center" width="600" cellspacing="0" cellpadding="0" bgcolor="#ffffff">
                                            <tbody>
                                                <tr>
                                                    <td class="esd-structure es-p20t es-p40b es-p40r es-p40l" esd-custom-block-id="8537" align="left">
                                                        <table width="100%%" cellspacing="0" cellpadding="0">
                                                            <tbody>
                                                                <tr>
                                                                    <td class="esd-container-frame" align="left" width="518">
                                                                        <table width="100%%" cellspacing="0" cellpadding="0">
                                                                            <tbody>
                                                                                <tr>
                                                                                    <td class="esd-block-image es-m-txt-c es-p5b" align="center" style="font-size:0"><a target="_blank"><img src="https://hipffk.stripocdn.email/content/guids/CABINET_3b108160279226d19ea952ffd5948152/images/92241523451206218.png" alt="icon" style="display: block;" title="icon" width="30"></a></td>
                                                                                </tr>
                                                                                <tr>
                                                                                    <td class="esd-block-text es-m-txt-c" align="center">
                                                                                        <h2>Hi %s <br></h2>
                                                                                    </td>
                                                                                </tr>
                                                                                <tr>
                                                                                    <td class="esd-block-text es-m-txt-c es-p15t" align="center">
                                                                                        <p>Thanks for signing up with clientname!
																							Please click the button below to validate your account and confirm we've got the right email address. If you don’t know why you got this email, please tell us straight away so we can fix this for you.
																						</p>
                                                                                    </td>
                                                                                </tr>
                                                                                <tr>
                                                                                    <td class="esd-block-button es-p20t es-p15b es-p10r es-p10l" align="center"><span class="es-button-border"><a href="%s" class="es-button" target="_blank">Confirm Email</a></span></td>
                                                                                </tr>
                                                                            </tbody>
                                                                        </table>
                                                                    </td>
                                                                </tr>
                                                            </tbody>
                                                        </table>
                                                    </td>
                                                </tr>
                                            </tbody>
                                        </table>
                                    </td>
                                </tr>
                            </tbody>
                        </table>
                        <table class="esd-footer-popover es-content" align="center" cellspacing="0" cellpadding="0">
                            <tbody>
                                <tr>
                                    <td class="esd-stripe" align="center">
                                        <table class="es-content-body" style="background-color: transparent;" align="center" width="600" cellspacing="0" cellpadding="0">
                                            <tbody>
                                                <tr>
                                                    <td class="esd-structure es-p30t es-p30b es-p20r es-p20l" align="left">
                                                        <table width="100%%" cellspacing="0" cellpadding="0">
                                                            <tbody>
                                                                <tr>
                                                                    <td class="esd-container-frame" align="center" width="560" valign="top">
                                                                        <table width="100%%" cellspacing="0" cellpadding="0">
                                                                            <tbody>
                                                                                <tr>
                                                                                    <td align="center" class="esd-empty-container" style="display: none;"></td>
                                                                                </tr>
                                                                            </tbody>
                                                                        </table>
                                                                    </td>
                                                                </tr>
                                                            </tbody>
                                                        </table>
                                                    </td>
                                                </tr>
                                            </tbody>
                                        </table>
                                    </td>
                                </tr>
                            </tbody>
                        </table>
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
    <div style="position: absolute; left: -9999px; top: -9999px; margin: 0px;"></div>
</body>

</html>`, emailData.ToName, verificationLink)
	smtpPort, _ := strconv.Atoi(os.Getenv("GMAIL_SMTP_PORT"))
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("VERIFICATION_EMAIL"))
	m.SetHeader("To", emailData.ToEmail)
	m.SetHeader("Subject", emailData.Subject)
	m.SetBody("text/html", emailBody)

	d := gomail.NewDialer(os.Getenv("GMAIL_SMTP"), smtpPort, os.Getenv("VERIFICATION_EMAIL"), os.Getenv("VERIFICATION_EMAIL_PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
	}
}

func SendResetPasswordToken(emailData SendEmail, verificationLink string) {
	var emailBody = fmt.Sprintf(`<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html>

<head>
    <meta charset="UTF-8">
    <meta content="width=device-width, initial-scale=1" name="viewport">
    <meta name="x-apple-disable-message-reformatting">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta content="telephone=no" name="format-detection">
    <title></title>
    <!--[if (mso 16)]>
    <style type="text/css">
    a {text-decoration: none;}
    </style>
    <![endif]-->
    <!--[if gte mso 9]><style>sup { font-size: 100%% !important; }</style><![endif]-->
</head>

<body>
    <div class="es-wrapper-color">
        <!--[if gte mso 9]>
			<v:background xmlns:v="urn:schemas-microsoft-com:vml" fill="t">
				<v:fill type="tile" color="#ffffff"></v:fill>
			</v:background>
		<![endif]-->
        <table class="es-wrapper" width="100%%" cellspacing="0" cellpadding="0">
            <tbody>
                <tr>
                    <td class="esd-email-paddings" valign="top">
                        <table class="es-content esd-header-popover" align="center" cellspacing="0" cellpadding="0">
                            <tbody>
                                <tr>
                                    <td class="esd-stripe" align="center">
                                        <table class="es-content-body" align="center" width="600" cellspacing="0" cellpadding="0" bgcolor="#ffffff">
                                            <tbody>
                                                <tr>
                                                    <td class="esd-structure es-p20t" esd-custom-block-id="8430" align="left">
                                                        <table width="100%%" cellspacing="0" cellpadding="0">
                                                            <tbody>
                                                                <tr>
                                                                    <td class="esd-container-frame" align="center" width="600" valign="top">
                                                                        <table width="100%%" cellspacing="0" cellpadding="0">
                                                                            <tbody>
                                                                                <tr>
                                                                                    <td class="esd-block-image es-p20r es-p20l" align="center" style="font-size:0"><a target="_blank"><img class="adapt-img" src="https://hipffk.stripocdn.email/content/guids/CABINET_57116f9afe83495a646cd7734bc77d26/images/39641523866414281.jpg" alt="Image" style="display: block;" title="Image" width="260"></a></td>
                                                                                </tr>
                                                                            </tbody>
                                                                        </table>
                                                                    </td>
                                                                </tr>
                                                            </tbody>
                                                        </table>
                                                    </td>
                                                </tr>
                                            </tbody>
                                        </table>
                                    </td>
                                </tr>
                            </tbody>
                        </table>
                        <table class="es-content" align="center" cellspacing="0" cellpadding="0">
                            <tbody>
                                <tr>
                                    <td class="esd-stripe" align="center">
                                        <table class="es-content-body" style="border-left:1px solid transparent;border-right:1px solid transparent;border-top:1px solid transparent;border-bottom:1px solid transparent;" align="center" width="600" cellspacing="0" cellpadding="0" bgcolor="#ffffff">
                                            <tbody>
                                                <tr>
                                                    <td class="esd-structure es-p20t es-p40b es-p40r es-p40l" esd-custom-block-id="8537" align="left">
                                                        <table width="100%%" cellspacing="0" cellpadding="0">
                                                            <tbody>
                                                                <tr>
                                                                    <td class="esd-container-frame" align="left" width="518">
                                                                        <table width="100%%" cellspacing="0" cellpadding="0">
                                                                            <tbody>
                                                                                <tr>
                                                                                    <td class="esd-block-image es-m-txt-c es-p5b" align="center" style="font-size:0"><a target="_blank"><img src="https://hipffk.stripocdn.email/content/guids/CABINET_3b108160279226d19ea952ffd5948152/images/92241523451206218.png" alt="icon" style="display: block;" title="icon" width="30"></a></td>
                                                                                </tr>
                                                                                <tr>
                                                                                    <td class="esd-block-text es-m-txt-c" align="center">
                                                                                        <h2>Hi %s <br></h2>
                                                                                    </td>
                                                                                </tr>
                                                                                <tr>
                                                                                    <td class="esd-block-text es-m-txt-c es-p15t" align="center">
                                                                                        <p>This is your reset password link!
																							Please click the button below to validate your reset password
																						</p>
                                                                                    </td>
                                                                                </tr>
                                                                                <tr>
                                                                                    <td class="esd-block-button es-p20t es-p15b es-p10r es-p10l" align="center"><span class="es-button-border"><a href="%s" class="es-button" target="_blank">Reset Password</a></span></td>
                                                                                </tr>
                                                                            </tbody>
                                                                        </table>
                                                                    </td>
                                                                </tr>
                                                            </tbody>
                                                        </table>
                                                    </td>
                                                </tr>
                                            </tbody>
                                        </table>
                                    </td>
                                </tr>
                            </tbody>
                        </table>
                        <table class="esd-footer-popover es-content" align="center" cellspacing="0" cellpadding="0">
                            <tbody>
                                <tr>
                                    <td class="esd-stripe" align="center">
                                        <table class="es-content-body" style="background-color: transparent;" align="center" width="600" cellspacing="0" cellpadding="0">
                                            <tbody>
                                                <tr>
                                                    <td class="esd-structure es-p30t es-p30b es-p20r es-p20l" align="left">
                                                        <table width="100%%" cellspacing="0" cellpadding="0">
                                                            <tbody>
                                                                <tr>
                                                                    <td class="esd-container-frame" align="center" width="560" valign="top">
                                                                        <table width="100%%" cellspacing="0" cellpadding="0">
                                                                            <tbody>
                                                                                <tr>
                                                                                    <td align="center" class="esd-empty-container" style="display: none;"></td>
                                                                                </tr>
                                                                            </tbody>
                                                                        </table>
                                                                    </td>
                                                                </tr>
                                                            </tbody>
                                                        </table>
                                                    </td>
                                                </tr>
                                            </tbody>
                                        </table>
                                    </td>
                                </tr>
                            </tbody>
                        </table>
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
    <div style="position: absolute; left: -9999px; top: -9999px; margin: 0px;"></div>
</body>

</html>`, emailData.ToName, verificationLink)
	smtpPort, _ := strconv.Atoi(os.Getenv("GMAIL_SMTP_PORT"))
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("VERIFICATION_EMAIL"))
	m.SetHeader("To", emailData.ToEmail)
	m.SetHeader("Subject", emailData.Subject)
	m.SetBody("text/html", emailBody)

	d := gomail.NewDialer(os.Getenv("GMAIL_SMTP"), smtpPort, os.Getenv("VERIFICATION_EMAIL"), os.Getenv("VERIFICATION_EMAIL_PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
	}
}
