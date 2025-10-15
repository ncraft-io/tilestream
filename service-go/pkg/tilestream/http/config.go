package http

type Config struct {
	// use {level},{x},{y} for template
	TileUrl string `json:"tileUrl,omitempty"`
	InfoUrl string `json:"infoUrl,omitempty"`
	Proxy   string `json:"proxy"`
}
