
package api_resp

import ("testing")


func TestAddOneItem(t *testing.T) {
	all_res := New()
	single_res := SingleResult{
		Stddev:1.41421356237,
		Data: []float64{1, 2, 3, 4, 5},
	}
	all_res.Add(single_res)
	if len(all_res.StDev) != 1{
		t.Errorf("Item was not added!")
	}
	
}


func TestAddTwoItems(t *testing.T) {
	all_res := New()
	single_res := SingleResult{
		Stddev:1.41421356237,
		Data: []float64{1, 2, 3, 4, 5},
	}
	single_res_2 := SingleResult{
		Stddev:1.41421356237,
		Data: []float64{1, 2, 3, 4, 5},
	}
	all_res.Add(single_res)
	all_res.Add(single_res_2)
	if len(all_res.StDev) != 2{
		t.Errorf("Items were not added!")
	}
}


func TestGetAll(t *testing.T) {
	all_res := New()
	single_res := SingleResult{
		Stddev:1.41421356237,
		Data: []float64{1, 2, 3, 4, 5},
	}
	single_res_2 := SingleResult{
		Stddev:1.41421356237,
		Data: []float64{1, 2, 3, 4, 5},
	}
	all_res.Add(single_res)
	all_res.Add(single_res_2)
	all := all_res.GetAll()
	if len(all) != 2{
		t.Errorf("Expected length 2!")
	}
}