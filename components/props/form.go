package props

type FormValues struct {
	FirstName            string
	LastName             string
	Email                string
	Password             string
	PasswordConfirmation string
	MarketingAccept      string
}

type FormErrors struct {
	Password string
	Email    string
}
