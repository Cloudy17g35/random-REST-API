package api_resp


type SingleResult struct {
	Stddev float64 `json:"stddev"`
	Data []float64 `json:"data"`
}

type AllResults struct {
	StDev []SingleResult
}


func New() *AllResults {
	return &AllResults{}
	}

func (a * AllResults) Add(result SingleResult) {
	a.StDev = append(a.StDev, result)
}

func (a *AllResults) GetAll() []SingleResult {
	return a.StDev
}