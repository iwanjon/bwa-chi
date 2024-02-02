package user

type userformater struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
	ImageUrl   string `json:"image_url"`
}

func Formatuser(u User, token string) userformater {
	newformat := userformater{
		ID:         u.ID,
		Name:       u.Name,
		Occupation: u.Occupation,
		Email:      u.Email,
		Token:      token,
		ImageUrl:   u.AvatarFileName,
	}
	return newformat

}
