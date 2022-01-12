package structure_decorator

/*
	举例不好

	装饰器的目的，给被装饰对象增强功能，解决继承关系过于复杂问题，通过组合替代继承

	代码形式和代理模式没区别

	代理模式		主要是给被代理类 添加无关功能
	装饰器模式	主要给被被装饰类 增强功能，存在关联性
*/

type WeaponImp interface {
	Fire() int
}

type Gun struct {
	damage int
}

func (g *Gun) Fire() int {

	return g.damage
}

// Pistol 手枪
type Pistol struct {
	g       WeaponImp
	Enhance func(options ...interface{})
}

func NewPistol(pistol WeaponImp, enhance func(options ...interface{})) *Pistol {
	// init logic

	return &Pistol{
		g:       pistol,
		Enhance: enhance,
	}
}

func (pistol *Pistol) Fire() int {
	//

	return pistol.g.Fire()
}
