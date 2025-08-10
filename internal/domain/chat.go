package domain

type Chat struct {
	ID       int  `json:"id"`
	IsActive bool `json:"is_active"`
}

type ChatFilter struct {
	ID       *int  `json:"id"`
	IsActive *bool `json:"is_active"`
}

type ChatCreate struct {
	ID       int  `json:"id"`
	IsActive bool `json:"is_active"`
}

type ChatUpdate struct {
	IsActive *bool `json:"is_active"`
}
