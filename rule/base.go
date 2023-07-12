package rule

type Rule struct{}

func (r Rule) New() *Rule {
	return &r
}

func (r *Rule) Execute() int {
	return r.validateExpired()
}
