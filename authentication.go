package libgomessage

import(

)

const(
	// the "U" is cut off!
	// ser: mario
	// Pass: sha256("itsame")
	USER_PASSWORD = "U"

	// Salt: random
	// User: mario
	// Pass: sha256("itsame" + salt)
	USER_PASSWORD_SALT = "A"

	// salt:`sha256( username + password + salt )`
	SHA256 = "B"

	// salt:`sha512( username + password + salt )`
	SHA512 = "C"

	// this password is only valid once and can thus be any string
	ONE_TIME_PASSWORD = "D"
)
