package knowledgecloud

type Knowledge interface {
	/**
	 * @brief      { build the data set, cam be used only one time }
	 *
	 * @param      dataranks  indicates if there is a rank data
	 * @param      datapoints  indicates if there is a point data
	 * @param      pointdata     points data set
	 * @param      rankdata     ranks data set
	 *
	 * @return     { the first is the error, the second indicates if you tryed to use it again }
	 */
	Sethead(_real bool, _single bool, _samples int, _labels int,
		_point_rn int, _rank_rn int, _ranksize int) (bool, bool)
	Setdata(datapoints bool, dataranks bool, pointdata []int, rankdata []int) (bool, bool)
	// workonRanks() bool

	// learnonKNN() (bool, bool, int)
	// learnonKmeans() (bool, bool, int)

	// binaryAnswer() (bool, bool, int)
	// piechartAnswer() (bool, bool, int, []int)
}

// /**
//  * @brief      { function_description }
//  *
//  * @return     { description_of_the_return_value }
//  */
// func workonRanks() bool

// /**
//  * @brief      { function_description }
//  *
//  * @return     { description_of_the_return_value }
//  */
// func learnonKNN() (bool, bool, int)

// *
//  * @brief      { function_description }
//  *
//  * @return     { description_of_the_return_value }

// func learnonKmeans() (bool, bool, int)

// /**
//  * @brief      { function_description }
//  *
//  * @return     { description_of_the_return_value }
//  */
// func binaryAnswer() (bool, bool, int)

// /**
//  * @brief      { function_description }
//  *
//  * @return     { description_of_the_return_value }
//  */
// func piechartAnswer() (bool, bool, int, []int)
