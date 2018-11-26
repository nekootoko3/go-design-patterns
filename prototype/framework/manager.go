package framework

type manager struct {
	showCase map[string]Producter
}

func NewManager() *manager {
	return &manager{showCase: make(map[string]Producter)}
}

func (m *manager) Register(name string, proto Producter) {
	m.showCase[name] = proto
}

func (m *manager) Clone(protoname string) Producter {
	proto := m.showCase[protoname]
	return proto.CreateClone()
}
