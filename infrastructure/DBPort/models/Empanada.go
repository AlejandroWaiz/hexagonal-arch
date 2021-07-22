package models

//Basic struct of an empanada to work on the repository
type Empanada struct {
	ID                  int
	EmpanadaName        string
	EmpanadaValue       int
	EmpanadaDescription string
	Ingredients
}

/*

	Empanada -> table

	EmpanadaID	 PrimaryKey
	EmpanadaName
	EmpanadaValue
	EmpanadaDescription


	Ingredients -> table

	EmpanadaID    UniqueKey
	Cheese
	Mushrooms
	...

	VoucherID 	UniqueKey
	TotalAmmount


*/
