package app

type Handler struct {
	appUsecase appUsecase
}

func New(appUsecase appUsecase) *Handler {
	return &Handler{
		appUsecase: appUsecase,
	}
}