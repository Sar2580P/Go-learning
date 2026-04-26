package types


type Student struct{
	Id int  
	Name string  `validate:"required"`   // validation tag to indicate that this field is required
	Email string  `validate:"required"`
	Age int   `validate:"required"`
}