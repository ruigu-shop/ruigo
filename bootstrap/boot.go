package bootstrap

import "github.com/kataras/iris/v12"

type Configurator func(*Bootstrapper)

type Bootstrapper struct {
	*iris.Application
	AppName string
}

func New(appName string, cfgs ...Configurator) *Bootstrapper {
	b := &Bootstrapper{
		Application: iris.New(),
		AppName:     appName,
	}
	for _, cfg := range cfgs {
		cfg(b)
	}
	return b
}

func (b *Bootstrapper) SetupViews(viewDir string) {
	b.RegisterView(iris.HTML(viewDir, ".html"))
}

func (b *Bootstrapper) Configure(cs ...Configurator) {
	for _, c := range cs {
		c(b)
	}
}

func (b *Bootstrapper) Listen(addr string, cfgs ...iris.Configurator) {
	b.Run(iris.Addr(addr), cfgs...)
}
