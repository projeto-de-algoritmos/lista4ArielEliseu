package knowledgecloud

type Knowledge interface {
	/**
	 * @brief      { function_description }
	 *
	 * @param      _real      Indicates if the data set is real or ideal
	 * @param      _single    Indicates the nnumber of samples to be analyzed
	 * @param      _samples   Number of samples (in the data set)
	 * @param      _labels    Number of labels
	 * @param      _point_rn  Space of the ranks
	 * @param      _rank_rn   Space of the points
	 * @param      _ranksize  Number of itens in the rank
	 *
	 * @return     { description_of_the_return_value }
	 */

	Sethead(_real bool, _single bool, _dataset int, _samples int, _labels int,
		_point_rn int, _rank_rn int, _ranksize int) (bool, bool)
	/**
	 * @brief      { build the data set, cam be used only one time }
	 *
	 * @param      datapoints  indicates if there is a point data
	 * @param      dataranks  indicates if there is a rank data
	 * @param      pointdata     points data set
	 * @param      rankdata     ranks data set
	 *
	 * @return     { the first is the error, the second indicates if you tryed to use it again }
	 */
	Setdata(datapoints bool, dataranks bool, pointdata []int, rankdata []int) (bool, bool)

	/**
	 * @brief      { Uses a variation of K Means with ideas rankings to build distrank}
	 *
	 * @param      analyzed  The samples to be analyzed
	 *
	 * @return     { true if it had success }
	 */
	BworkonRanks(analyzed []int) bool

	// learnonKNN() (bool, bool, int)
	// learnonKmeans() (bool, bool, int)

	// binaryAnswer() (bool, bool, int)
	// piechartAnswer() (bool, bool, int, []int)
}
