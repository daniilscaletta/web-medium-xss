package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProfilePage(ctx *gin.Context) {

	isAuthenticated := ctx.GetBool("isAuthenticated")
	if isAuthenticated {

		user := CurrentUser
		if user == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		appointments := []map[string]string{
			{"Date": "2024-11-15", "Doctor": "Dr. Smith", "Specialty": "Cardiology"},
			{"Date": "2024-11-20", "Doctor": "Dr. Jones", "Specialty": "Dermatology"},
			{"Date": "2024-11-20", "Doctor": "Dr. Jones", "Specialty": "Dermatology"},
			{"Date": "2024-11-20", "Doctor": "Dr. Jones", "Specialty": "Dermatology"},
			{"Date": "2024-11-20", "Doctor": "Dr. Jones", "Specialty": "Dermatology"},
			{"Date": "2024-11-20", "Doctor": "Dr. Jones", "Specialty": "Dermatology"},
			{"Date": "2024-11-20", "Doctor": "Dr. Jones", "Specialty": "Dermatology"},
			{"Date": "2024-11-20", "Doctor": "Dr. Jones", "Specialty": "Dermatology"},
			{"Date": "2024-11-20", "Doctor": "Dr. Jones", "Specialty": "Dermatology"},
			{"Date": "2024-11-20", "Doctor": "Dr. Jones", "Specialty": "DermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatologyDermatology"},
		}

		ctx.HTML(http.StatusOK, "profile.html", gin.H{
			"Name":         user.Name,
			"Surname":      user.Surname,
			"DateOfBirth":  user.DateOfBirthday,
			"Email":        user.Email,
			"PhoneNumber":  user.PhoneNumber,
			"Passport":     user.Passport,
			"Login":        user.Login,
			"AvatarURL":    "https://ui-avatars.com/api/?name=" + user.Name + "+" + user.Surname + "&background=random",
			"Appointments": appointments,
		})
	} else {
		ctx.Redirect(http.StatusTemporaryRedirect, "/login")
	}

}
