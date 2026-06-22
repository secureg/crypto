package acme

type orderProfileOpt string

func (orderProfileOpt) privateOrderOpt() {}

func WithProfileOption(p string) OrderOption {
	return orderProfileOpt(p)
}
