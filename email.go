package amply

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

type email struct {
	client      *client
}

type EmailData struct {
	From                 interface{}            `json:"from,omitempty"`
	To                   interface{}            `json:"-"`
	Cc                   interface{}            `json:"-"`
	Bcc                  interface{}            `json:"-"`
	Subject              string                 `json:"subject,omitempty"`
	Text                 string                 `json:"-"`
	Html                 string                 `json:"-"`
	Content              []Content              `json:"content,omitempty"`
	ReplyTo              interface{}            `json:"reply_to,omitempty"`
	Template             interface{}            `json:"template,omitempty"`
	DynamicTemplateData  map[string]string      `json:"dynamic_template_data,omitempty"`
	Substitutions        map[string]string      `json:"substitutions,omitempty"`
	UnsubscribeGroupUuid string                 `json:"unsubscribe_group_uuid,omitempty"`
	IpOrPoolUuid         string                 `json:"ip_or_pool_uuid,omitempty"`
	SendAt               string                 `json:"send_at,omitempty"`
	Attachments          []Attachment           `json:"attachments,omitempty"`
	Headers              map[string]string      `json:"headers,omitempty"`
	Categories           []string               `json:"-"`
	Clicktracking        interface{}            `json:"-"`
	Analytics            map[string]interface{} `json:"analytics,omitempty"`
	Personalizations     []Personalization      `json:"personalizations,omitempty"`
}

type Attachment struct {
	Content     string `json:"content,omitempty"`
	Filename    string `json:"filename,omitempty"`
	Type        string `json:"type,omitempty"`
	Disposition string `json:"disposition,omitempty"`
	ContentId   string `json:"disposition,omitempty"`
}

type Content struct {
	Value string `json:"value,omitempty"`
	Type  string `json:"type,omitempty"`
}

type Personalization struct {
	To  []EmailAddress `json:"to,omitempty"`
	Cc  []EmailAddress `json:"cc,omitempty"`
	Bcc []EmailAddress `json:"bcc,omitempty"`
}

func (e email) Create(data EmailData) (*Response, error) {
	parsedEmailData, err := parseEmailData(data)

	if err != nil {
		return nil, err
	}

	return e.client.post("/email", parsedEmailData)
}

func parseEmailData(data EmailData) (EmailData, error) {
	emailData := EmailData{}

	if err := setFrom(&emailData, data.From); err != nil {
		return emailData, err
	}

	if err := setSubject(&emailData, data.Subject); err != nil {
		return emailData, err
	}

	if err := setText(&emailData, data.Text); err != nil {
		return emailData, err
	}

	if err := setHtml(&emailData, data.Html); err != nil {
		return emailData, err
	}

	if err := setContent(&emailData, data.Content); err != nil {
		return emailData, err
	}

	if err := setReplyTo(&emailData, data.ReplyTo); err != nil {
		return emailData, err
	}

	if err := setTemplate(&emailData, data.Template); err != nil {
		return emailData, err
	}

	if err := setDynamicTemplateData(&emailData, data.DynamicTemplateData); err != nil {
		return emailData, err
	}

	if err := setUnsubscribeGroupUuid(&emailData, data.UnsubscribeGroupUuid); err != nil {
		return emailData, err
	}

	if err := setIpOrPoolUuid(&emailData, data.IpOrPoolUuid); err != nil {
		return emailData, err
	}

	if err := setSendAt(&emailData, data.SendAt); err != nil {
		return emailData, err
	}

	if err := setAttachments(&emailData, data.Attachments); err != nil {
		return emailData, err
	}

	if err := setHeaders(&emailData, data.Headers); err != nil {
		return emailData, err
	}

	if err := setCategories(&emailData, data.Categories); err != nil {
		return emailData, err
	}

	if err := setClicktracking(&emailData, data.Clicktracking); err != nil {
		return emailData, err
	}

	if err := setSubstitutions(&emailData, data.Substitutions); err != nil {
		return emailData, err
	}

	if len(data.Personalizations) > 0 {
		if err := setPersonalizations(&emailData, data.Personalizations); err != nil {
			return emailData, err
		}
	} else {
		if err := setPersonalizationsFromTo(&emailData, data.To, data.Cc, data.Bcc); err != nil {
			return emailData, err
		}
	}

	return emailData, nil
}

func setFrom(e *EmailData, from interface{}) error {
	if from == nil {
		return nil
	}

	formattedFrom, err := formatEmail(from)
	if err != nil {
		return err
	}

	e.From = formattedFrom[0]
	return nil
}

func setSubject(e *EmailData, subject string) error {
	if subject == "" {
		return nil
	}

	e.Subject = subject
	return nil
}

func setText(e *EmailData, text string) error {
	if text == "" {
		return nil
	}

	content := Content{
		Type:  "text/plain",
		Value: text,
	}

	e.Content = append(e.Content, content)
	return nil
}

func setHtml(e *EmailData, html string) error {
	if html == "" {
		return nil
	}

	content := Content{
		Type:  "text/html",
		Value: html,
	}

	e.Content = append(e.Content, content)
	return nil
}

func setContent(e *EmailData, content []Content) error {
	if len(content) == 0 {
		return nil
	}

	for i, part := range content {
		if part.Value == "" {
			msg := fmt.Sprintf("String expected for Content[%d].Value", i)
			return errors.New(msg)
		}

		if part.Type == "" {
			msg := fmt.Sprintf("String expected for Content[%d].Type", i)
			return errors.New(msg)
		}

		e.Content = append(e.Content, part)
	}

	return nil
}

func setReplyTo(e *EmailData, replyTo interface{}) error {
	if replyTo == nil {
		return nil
	}

	formattedReplyTo, err := formatEmail(replyTo)
	if err != nil {
		return err
	}

	e.ReplyTo = formattedReplyTo[0]
	return nil
}

func setTemplate(e *EmailData, template interface{}) error {
	if template == nil {
		return nil
	}

	e.Template = template
	return nil
}

func setDynamicTemplateData(e *EmailData, dynamicTemplateData map[string]string) error {
	if dynamicTemplateData == nil {
		return nil
	}

	e.DynamicTemplateData = map[string]string{}

	for k, v := range dynamicTemplateData {
		e.DynamicTemplateData[k] = v
	}

	return nil
}

func setUnsubscribeGroupUuid(e *EmailData, unsubscribeGroupUuid string) error {
	if unsubscribeGroupUuid == "" {
		return nil
	}

	e.UnsubscribeGroupUuid = unsubscribeGroupUuid
	return nil
}

func setIpOrPoolUuid(e *EmailData, ipOrPoolUuid string) error {
	e.IpOrPoolUuid = ipOrPoolUuid
	return nil
}

func setSendAt(e *EmailData, SendAt string) error {
	e.SendAt = SendAt
	return nil
}

func setAttachments(e *EmailData, attachments []Attachment) error {
	for i, attachment := range attachments {
		if attachment.Content == "" {
			msg := fmt.Sprintf("String expected for Attachment[%d].Content", i)
			return errors.New(msg)
		}

		if attachment.Filename == "" {
			msg := fmt.Sprintf("String expected for Attachment[%d].Filename", i)
			return errors.New(msg)
		}

		e.Attachments = append(e.Attachments, attachment)
	}

	return nil
}

func setHeaders(e *EmailData, headers map[string]string) error {
	if headers == nil {
		return nil
	}

	e.Headers = headers
	return nil
}

func setCategories(e *EmailData, categories []string) error {
	if categories == nil {
		return nil
	}

	if e.Analytics == nil {
		e.Analytics = map[string]interface{}{}
	}

	e.Analytics["categories"] = categories
	return nil
}

func setClicktracking(e *EmailData, clicktracking interface{}) error {
	if clicktracking == nil {
		return nil
	}

	if _, isBool := clicktracking.(bool); !isBool {
		return errors.New("bool expected for `Clicktracking`")
	}

	if e.Analytics == nil {
		e.Analytics = map[string]interface{}{}
	}

	e.Analytics["clicktracking"] = clicktracking.(bool)
	return nil
}

func setPersonalizations(e *EmailData, personalizations []Personalization) error {
	if len(personalizations) == 0 {
		return nil
	}

	e.Personalizations = personalizations
	return nil
}

func setPersonalizationsFromTo(e *EmailData, to interface{}, cc interface{}, bcc interface{}) error {
	formattedTo, toError := formatEmail(to)
	if toError != nil {
		return toError
	}

	formattedCc, ccError := formatEmail(cc)
	if ccError != nil {
		return ccError
	}

	formattedBcc, bccError := formatEmail(bcc)
	if bccError != nil {
		return bccError
	}

	e.Personalizations = []Personalization{
		Personalization{
			To:  formattedTo,
			Cc:  formattedCc,
			Bcc: formattedBcc,
		},
	}
	return nil
}

func setSubstitutions(e *EmailData, substitutions map[string]string) error {
	if substitutions == nil {
		return nil
	}

	if e.Substitutions == nil {
		e.Substitutions = map[string]string{}
	}

	for subFrom, subTo := range substitutions {
		e.Substitutions[subFrom] = subTo
	}

	return nil
}

func formatEmail(data interface{}) ([]EmailAddress, error) {
	if data == nil {
		return nil, nil
	}

	rt := reflect.TypeOf(data)
	emails := make([]EmailAddress, 0)

	switch rt.Kind() {
	case reflect.Slice, reflect.Array:
		toJson, _ := json.Marshal(data)
		parsedData := make([]interface{}, 0)
		json.Unmarshal(toJson, &parsedData)

		for _, email := range parsedData {
			if email != nil {
				formattedEmail, err := NewEmailAddress(email)
				if err != nil {
					return nil, err
				}

				emails = append(emails, formattedEmail)
			}
		}
	case reflect.String:
		email, err := NewEmailAddress(data.(string))

		if err != nil {
			return nil, err
		}

		emails = append(emails, email)
	case reflect.Map:
		_, isMap := data.(map[string]string)
		if !isMap {
			return nil, errors.New("map[string]string expected for email address")
		}

		email, err := NewEmailAddress(data.(map[string]string))
		if err != nil {
			return nil, err
		}

		emails = append(emails, email)
	}

	return emails, nil
}
