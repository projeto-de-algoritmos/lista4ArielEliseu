package body

import (
	"./datapoint"
	"./datarank"
)

/*
Data Body allows you to build knowledge on some task.
Based on the behavior determined on the Head the Samples are analyzed in order to fill up results arrays.
The analysis is explored in the interface file.
*/
type databody struct {
	points point_b.datapoint //point
	ranks  rank_b.datarank   //rank
}
