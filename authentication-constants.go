package libgomessage

const(
	USERNAME =	"U"
	PASSWORD =	"P"
	SALT =		"S"

	/*
	AU: username
	P: sha1( password )
	*/
	AUTHENTICATION_USERNAME_PASSWORD	= "A"

	/*
	AU: username
	P: sha1(concat( password, salt))
	S: $RANDOM_SALT [50]byte
	*/
	AUTHENTICATION_USERNAME_PASSWORD_SALT	= "B"

	/*
	BU: sha1(concat( username, password ))
	*/
	AUTHENTICATION_LESS_METADATA		= "C"

	/*
	CU: sha1(concat( username, password, salt ))
	S: $RANDOM_SALT [50]byte
	*/
	AUTHENTICATION_LESS_METADATA_SALT	= "D"

	/*
	------BEGIN PGP SIGNED MESSAGE-----
	*/
	AUTHENTICATION_GPG			= "-"
)

