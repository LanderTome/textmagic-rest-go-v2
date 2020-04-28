/*
 * TextMagic API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

type SendPhoneVerificationCodeTfaInputObject struct {
	// Use the phone number in international E.164 format. If you need to pass a phone number in the local format, please use it with the **country** parameter to specify the origin country of the phone number. 
	Phone string `json:"phone,omitempty"`
	// **Workflows**  The Verify API allows you to select the best workflow for your use case. This might depend on the type of verification taking place, your users' preference, or their geographical location. You can specify which workflow to use for each Verify API request by setting the workflowId field to an integer value 1-7. The details of each of these preset workflows are detailed below.  <br />  **Workflow 1 (Default Workflow): SMS -> TTS -> TTS**  Send PIN code by text message, follow up with two subsequent voice calls if the request wasn't already verified.  Send SMS to user with PIN code Wait for 60 seconds Call user and give TTS PIN code Wait for 60 seconds Call user and give TTS PIN code  Request expires after 300 seconds  <br />  **Workflow 2: SMS -> SMS -> TTS**   Send PIN code by text message, follow up with a second text message and finally a voice call if the request has not been verified.  Send SMS to user with PIN code Wait for 60 seconds Send SMS to user with PIN code Wait for 60 seconds Call user and give TTS PIN code  Request expires after 300 seconds  <br />  **Workflow 3: TTS -> TTS**  Call the user and speak a PIN code, follow up with a second call if the request wasn't already verified.  Call user and give TTS PIN code Wait for 60 seconds Call user and give TTS PIN code  Request expires after 300 seconds  <br />  **Workflow 4: SMS -> SMS**   Send PIN code by text message, follow up with a second text message if the code hasn't been verified.  Send SMS to user with PIN code Wait for 60 seconds Send SMS to user with PIN code Wait for 60 seconds  Request expires after 300 seconds  <br />  **Workflow 5: SMS -> TTS**   Send PIN code by text message, follow up with a voice call if the code hasn't been verified.  Send SMS to user with PIN code Wait for 60 seconds Call user and give TTS PIN code Wait for 60 seconds  Request expires after 300 seconds  <br />  **Workflow 6: SMS**   Send PIN code by text message once only.  Send SMS to user with PIN code Request expires after 300 seconds  <br />  **Workflow 7: TTS**  Call the user and speak a PIN code once only.  Call user and give TTS PIN code  Request expires after 300 seconds 
	WorkflowId string `json:"workflowId,omitempty"`
	// An alphanumeric string with up to 18 characters you can use to personalize the verification text message body, to help users identify your company or application name. For example: “Your TextMagic PIN is …” 
	Brand string `json:"brand,omitempty"`
	// The length of the verification code. The value can be 4 or 6 characters. 
	CodeLength int32 `json:"codeLength,omitempty"`
	// By default, the SMS or text-to-speech (TTS) voice message is generated in the locale that matches the number. For example, the text message or TTS message for a 33\\* number is sent in French. Use this parameter to explicitly control the language, accent, and gender used for the verification request. Choosing one of the following: `de-de`, `en-au`, `en-gb`, `en-us`, `en-in`, `es-es`, `es-mx`, `es-us`, `fr-ca`, `fr-fr`, `is-is`, `it-it`, `ja-jp`, `ko-kr`, `nl-nl`, `pl-pl`, `pt-pt`, `pt-br`, `ro-ro`, `ru-ru`, `sv-se`, `tr-tr`, `zh-cn` or `zh-tw`. 
	Language string `json:"language,omitempty"`
	// One of the available [sender settings](https://my.textmagic.com/online/reply-options/) on your TextMagic account. If the specified sender setting type is not allowed for some destinations, a fallback default sender will be used to ensure message delivery. More info about known restrictions can be found [here](https://support.textmagic.com/article/how-to-understand-sender-setting-restrictions/). 
	SenderId string `json:"senderId,omitempty"`
	// The 2-letter ISO country code for the local phone number.
	Country string `json:"country,omitempty"`
}
