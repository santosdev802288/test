package kubgo

type IKubgoRepository interface {
	Save(kubgo *Kubgo) chan error
	Delete(kubgo *Kubgo) chan error
	Update(kubgo *Kubgo) chan error
}
