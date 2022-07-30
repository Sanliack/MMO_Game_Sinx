package simodel

type SetConnAddrModel struct {
	ConnAddrMap map[string]interface{}
}

func (s *SetConnAddrModel) SetAddr(key string, val interface{}) {
	s.ConnAddrMap[key] = val
}
func (s *SetConnAddrModel) GetAddr(key string) interface{} {
	return s.ConnAddrMap[key]
}

func NewSetAddrModel() *SetConnAddrModel {
	return &SetConnAddrModel{
		ConnAddrMap: make(map[string]interface{}, 10),
	}
}
