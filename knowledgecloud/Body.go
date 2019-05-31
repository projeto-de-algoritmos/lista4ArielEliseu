package knowledgecloud

/*
Data Body allows you to build knowledge on some task.
Based on the behavior determined on the Head the Samples are analyzed in order to fill up results arrays.
The analysis is explored in the interface file.
*/
type DataBody struct {
	ranks  []int //samples ranks
	points []int //samples points

	idealrankscenter []int //ideal "center of the ranks" for each label
	realrankcenter   []int //real "center of the ranks" for each label

	idealmasscenter []int //ideal center of mass
	realmasspoints  []int //real center of mass

	distrank    []int //results based on ranks
	distclassic []int //results based on points
}
