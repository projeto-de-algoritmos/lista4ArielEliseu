package knowledgecloud

import "fmt"

func (c *Cloud) Sethead(_real bool, _single bool, _dataset int, _samples int, _labels int, _point_rn int, _rank_rn int, _ranksize int) (bool, bool) {
	//nc := new(Cloud)
	if c.on {
		fmt.Println("This method can be used only one time, use addsamples() to add more samples,")
		fmt.Println("addrank() to add a rank data set, addpoint() to add a point data set or rebuild() ")
		fmt.Println("to rebuild the data set")
		return false, true
	} else {

		c.head.single = _single
		c.head.dataset = _dataset
		c.head.samples = _samples
		c.head.labels = _labels
		c.head.point_rn = _point_rn
		c.head.rank_rn = _rank_rn
		c.head.ranksize = _ranksize
		c.on = true
		c.real = _real
		return true, false
	}

}

func (c *Cloud) Setdata(datapoints bool, dataranks bool, pointdata []int, rankdata []int) (bool, bool) {
	if datapoints {
		pt := make([]int, (c.head.samples * c.head.point_rn))
		pt = pointdata
		c.body.points = pt
	}
	if dataranks {
		rk := make([]int, (c.head.samples * c.head.rank_rn * c.head.ranksize))
		rk = rankdata
		c.body.ranks = rk
	}
	return true, false
}

func (c *Cloud) BworkonRanks(analyzed []int) bool {
	var auxDists [c.head.samples * c.head.rank_rn]int
	var index int
	for s := 0; s < c.head.samples; s++ {
		for rn := 0; rn < c.head.rank_rn; rn++ {
			//C
		}
	}
}
