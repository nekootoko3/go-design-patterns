package framework

type Builder interface {
	MakeTitle(title string)
	MakeString(str string)
	MakeItems(items []string)
	Close()
}

type director struct {
	builder Builder
}

func NewDirector(builder Builder) *director {
	return &director{builder: builder}
}

func (d *director) Construct() {
	d.builder.MakeTitle("Greeting")
	d.builder.MakeString("Morning And Noon")
	d.builder.MakeItems([]string{"Good Morning", "Good Afternoon"})
	d.builder.MakeString("Night")
	d.builder.MakeItems([]string{"Good Night", "Good Bye"})
	d.builder.Close()
}
