package simodel

type MessageModel struct {
	Id      uint32
	Len     uint32
	Content []byte
}

func (m *MessageModel) GetMsgID() uint32 {
	return m.Id
}
func (m *MessageModel) GetMsgLen() uint32 {
	return m.Len
}
func (m *MessageModel) GetData() []byte {
	return m.Content
}
func (m *MessageModel) SetMsgLen(len uint32) {
	m.Len = len
}
func (m *MessageModel) SetData(data []byte) {
	m.Content = data
}
func (m *MessageModel) SetMsgID(id uint32) {
	m.Id = id
}

func NewMessageModel(id uint32, len uint32, content []byte) *MessageModel {
	return &MessageModel{
		Id:      id,
		Len:     len,
		Content: content,
	}
}
