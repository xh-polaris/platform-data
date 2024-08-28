package handler

type Handler struct {
	IInsertHandler
}

func NewHandler() *Handler {
	return &Handler{
		IInsertHandler: NewInsertHandler(),
	}
}
