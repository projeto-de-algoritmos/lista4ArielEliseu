package knowledgecloud

type Cloud struct {
	manager head.mastercontrol
	head    head.datahead
	body    body.databody
}
