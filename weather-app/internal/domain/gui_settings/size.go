package guisettings

// WindowSize структура размеров окна
type WindowSize struct {
    width  int
    height int
}

// NewWS создает новый экземпляр WindowSize
func NewWS(w, h int) WindowSize {
    return WindowSize{width: w, height: h}
}

// IsFull проверяет, полноэкранное ли окно
func (ws WindowSize) IsFull() bool {
    return ws.width == 0 && ws.height == 0
}

// Width возвращает ширину окна
func (ws WindowSize) Width() int {
    return ws.width
}

// Height возвращает высоту окна
func (ws WindowSize) Height() int {
    return ws.height
}
