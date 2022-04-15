package models

// TemplateData holds data sent from handlers to template
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{} // interface object is used for any type of data
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
