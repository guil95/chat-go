package bot

type Stock struct {
	Code  string `json:"code"`
	Value string `json:"value"`
	Room  string `json:"roomID"`
}
