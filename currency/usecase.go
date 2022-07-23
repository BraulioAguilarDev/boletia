package currency

type Usecase interface {
	Get() error
	Pull() error
}
