package message

type (
	Message struct {
		Aps Aps `json:"aps"`
	}

	Aps struct {
		Alert Alert `json:"alert"`
	}

	Alert struct {
		Title string `json:"title"`

		Subtitle string `json:"subtitle"`

		Body string `json:"body"`
	}
)

func New(title, subtitle, message string) *Message {
	return &Message{
		Aps: Aps{
			Alert: Alert{
				Title:    title,
				Subtitle: subtitle,
				Body:     message,
			},
		},
	}
}
