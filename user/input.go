package user

//input dari user
//handler, mapping input dari user -> struck input
//service : melakukan mapping dari struck input ke struck user
//repository
//db

type RegisterUserInput struct {
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
}
