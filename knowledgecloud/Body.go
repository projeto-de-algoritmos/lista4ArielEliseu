package knowledgecloud

/*
Data Body allows you to build knowledge on some task.
Based on the behavior determined on the Head the Samples are analyzed in order to fill up results arrays.
The analysis is explored in the interface file.
*/
type DataBody struct {
	Ranks []int //samples ranks

	Idealrankscenter []int //ideal "center of the ranks" for each label
	Realrankcenter   []int //real "center of the ranks" for each label

	Idealpoints []int //ideal center of mass
	Realpoints  []int //real center of mass

	Distrank    []int //results based on ranks
	Distclassic []int //results based on points
}
