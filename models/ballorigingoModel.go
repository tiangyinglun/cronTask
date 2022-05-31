package models

type BallOriginGo struct {
	//gorm.Model
	Id          int64  `json:"id"`
	CId         string `json:"c_id"`
	Types       string `json:"types"`
	Win         string `json:"win"`
	Flat        string `json:"flat"`
	Loss        string `json:"loss"`
	Probability string `json:"probability"`
	UpdateTime  int64  `json:"update_time"`
	DateTime    string `json:"date_time"`
}

