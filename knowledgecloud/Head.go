package knowledgecloud

/*
Data Head allows you to control the behavior of a Knowledge Cloud built on this Head and on a Data Body based 	in integers arrays that are interpreted as some Rn ranked data or Rn cloud of points.
The control role of the Data Head variables are:
*/
type DataHead struct {
	Single bool /*Single: Determine the nature of a Data Cloud object to be either exclusive to some given sample or incorporate a group of samples*/

	Samples int /*Samples: The number of samples, it's 1 if Single is true, 1 <= Samples <= (Max int)/(Rn*Labels)   for the points vector and  1 <= Samples <= (Max int)/(Rn*RankSize*Labels)   for the points vector
	Labels: the number of Labels, it's not limited by the relation with the Max int, but it's a variable that can limit the other Head variables.*/

	Labels int //Rank_Rn: The  space in which the ranks are defined

	Point_Rn int //Point_Rn: The  space in which the points are defined

	Rank_Rn int /*Labels: the number of Labels, it's not limited by the relation with the Max int, but it's a variable that can limit the other Head variables.*/

	RankSize int //RankSize: The number of items that are in the rank
}
