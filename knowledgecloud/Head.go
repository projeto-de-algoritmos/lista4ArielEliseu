package knowledgecloud

import (
	"./head"
)

/*
Data Head allows you to control the behavior of a Knowledge Cloud built on this Head and on a Data Body based 	in integers arrays that are interpreted as some Rn ranked data or Rn cloud of points.
The control role of the Data Head variables are:
*/
type DataHead struct {
	control   head.datacontrol
	headpoint head.pointcontrol
	headrank  head.rankcontrol
}
