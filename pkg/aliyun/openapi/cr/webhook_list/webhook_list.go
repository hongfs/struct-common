package webhook_list

type Webhook struct {
	WebhookID   int    `json:"webhookId,omitempty"`
	TriggerTag  string `json:"triggerTag,omitempty"`
	WebhookName string `json:"webhookName,omitempty"`
	TriggerType string `json:"triggerType,omitempty"`
	WebhookURL  string `json:"webhookUrl,omitempty"`
}

type Data struct {
	Webhooks []Webhook `json:"webhooks,omitempty"`
}

type ResponseData struct {
	Data Data `json:"data,omitempty"`
}
