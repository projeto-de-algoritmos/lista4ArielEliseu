package head

import (
	"./head"
)

/*
Data Head allows you to control the behavior of a Knowledge Cloud built on this Head and on a Data Body based 	in integers arrays that are interpreted as some Rn ranked data or Rn cloud of points.
The control role of the Data Head variables are:
*/
type datahead struct {
	single bool /*Single: Determine the nature of a Data Cloud object to be either exclusive to some given sample or incorporate a group of samples*/

	pointcontrol point_h.headpoint //point_h

	rankcontrol rank_h.headrank //rank_h
}
