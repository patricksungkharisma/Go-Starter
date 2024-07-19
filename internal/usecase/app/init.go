package app

type Usecase struct {
	appRepo appRepo
}

func New(
	appRepo appRepo,
) *Usecase {
	return &Usecase{
		appRepo: appRepo,
	}
}
