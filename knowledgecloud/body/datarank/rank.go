import (
	"../body/data"
	"../body/knowledge"
)

type datarank struct {
	ranks data.data

	idealrankscenter knowledge.knowledge
	realrankcenter   knowledge.knowledge

	distrank knowledge.knowledge
}
