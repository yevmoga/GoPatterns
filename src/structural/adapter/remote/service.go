package remote

type CustomerInfo struct {
	FirstName  string
	SecondName string
	BirthYear  uint16
}

func GetUsefulInfo() CustomerInfo {
	return CustomerInfo{
		FirstName:  "Michael",
		SecondName: "Jordan",
		BirthYear:  1963,
	}
}
