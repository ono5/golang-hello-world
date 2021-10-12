package models

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMMap map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flaash    string
	Warning   string
	Error     string
}
