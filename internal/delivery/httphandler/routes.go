package httphandler

const (
	// AuthRouteGroup - group route is for authorization and
	// registration new users
	AuthRouteGroup string = "/auth"
	// ApiRouteGroup - group route implements api
	ApiRouteGroup string = "/api"
)

const (
	// RegisterRoute - route is for registration new user
	RegisterRoute string = "/register"
	// LoginRoute - route for registration cew user
	LoginRoute string = "/login"
	// LogoutRoute - route for logout
	LogoutRoute string = "/logout"
	// NotesRoute - route for  notes
	NotesRoute string = "/notes"
	// BankcardsRoute - route for  bank cards
	BankcardsRoute string = "/bankcards"
	// BinariesRoute - route for  binary data
	BinariesRoute string = "/binaries"
	// CredentialsRoute - route for  credentials
	CredentialsRoute string = "/credentials"
	// NoteRoute - route for  note
	NoteRoute string = "/notes/{id}"
	// NotesRoute - route for bank card
	BankcardRoute string = "/bankcards/{id}"
	// BinaryRoute - route for  binary data
	BinaryRoute string = "/binaries/{id}"
	// CredentialRoute - route for credentials
	CredentialRoute string = "/credentials/{id}"
)
