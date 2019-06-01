package point

import (
	"../body/data"
	"../body/knowledge"
)

type datapoint struct {
	points data.data

	idealmasscenter knowledge.knowledge
	realmasscenter  knowledge.knowledge

	distclassic knowledge.knowledge
}
