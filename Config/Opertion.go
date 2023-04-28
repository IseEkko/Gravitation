package Config

type Opertion interface {
	Load() ([]*Value, error)
}
