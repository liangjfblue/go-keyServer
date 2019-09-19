package user

type ListUserReq struct {
    Mid string  `json:"mid"`
    Sn  string  `json:"sn"`
}

type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
