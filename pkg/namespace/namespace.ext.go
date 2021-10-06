package namespace

func NewNamespaceSet() *Set {
	return &Set{
		Items: []*Namespace{},
	}
}

func (s *Set) Add(item *Namespace) {
	s.Items = append(s.Items, item)
}

func NewDefaultNamespace() *Namespace {
	return &Namespace{
		Enabled: true,
	}
}
