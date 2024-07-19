package app

type Handler struct {
	AppUsecase appUsecase
}

func New(
	AppUsecase appUsecase,
) *Handler {
	return &Handler{
		AppUsecase: AppUsecase,
	}
}
