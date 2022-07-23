package currency

type Repository interface {
	Get() error
	Pull() error
}
