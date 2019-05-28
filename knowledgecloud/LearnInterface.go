package knowledgecloud

type Knowledge interface {
	buildcloud() (bool, bool)

	workonRanks() bool

	learnonKNN() (bool, bool, int)
	learnonKmeans() (bool, bool, int)

	binaryAnswer() (bool, bool, int)
	piechartAnswer() (bool, bool, int, []int)
}

// /**
//  * @brief      { function_description }
//  *
//  * @return     { description_of_the_return_value }
//  */
// func buildcloud() (bool, bool)

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
