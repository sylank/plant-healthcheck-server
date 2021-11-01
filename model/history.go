package model

type History struct {
	size     int
	elements []HistoryElement
}

type HistoryElement struct {
	DateStr      string
	Temperature  float32
	Humidity     float32
	SoilMoisture float32
}

func CreateHistory(maxSize int) *History {
	return &History{
		size:     maxSize,
		elements: []HistoryElement{},
	}
}

func (h *History) Push(element HistoryElement) {
	h.elements = append(h.elements, element)
	if len(h.elements) > h.size {
		h.elements = h.elements[1:]
	}
}

func (h *History) GetElements() []HistoryElement {
	return h.elements
}

func (h *History) GetLatest() HistoryElement {
	if len(h.elements) > 0 {
		return h.elements[len(h.elements)-1]
	}

	return HistoryElement{}
}
