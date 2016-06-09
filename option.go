package config

type Option struct {
	key      string
	iv       string
	value    []string
	comments []string
}

func NewOption(key, iv string, value []string) *Option {
	var opt = &Option{}
	opt.key = key
	opt.iv = iv
	opt.value = value
	return opt
}

func (this *Option) Key() string {
	return this.key
}

func (this *Option) Comments() []string {
	return this.comments
}

func (this *Option) Comment() string {
	if len(this.comments) > 0 {
		return this.comments[0]
	}
	return ""
}

func (this *Option) AddComment(comment string) {
	this.comments = append(this.comments, comment)
}

func (this *Option) Value() string {
	if len(this.value[0]) > 0 {
		return this.value[0]
	}
	return ""
}

func (this *Option) SetValue(v string) {
	this.value = []string{v}
}

func (this *Option) AppendValue(v ...string) {
	this.value = append(this.value, v...)
}

func (this *Option) ListValue() []string {
	return this.value
}

func (this *Option) String() string {
	return this.Value()
}
