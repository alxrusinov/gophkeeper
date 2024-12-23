package httphandler

const (
	// AuthRouteGroup - group route is for authorization and
	// registration new users
	AuthRouteGroup = "/auth"
	// ApiRouteGroup - group route implements api
	ApiRouteGroup = "/api"
)

const (
	// RegisterRoute - route is for registration new user
	RegisterRoute = "/register"
	// LoginRoute - route for registration cew user
	LoginRoute = "/login"
	// LogoutRoute - route for logout
	LogoutRoute = "/logout"
	// NotesRoute - route for  notes
	NotesRoute = "/notes"
	// BankcardsRoute - route for  bank cards
	BankcardsRoute = "/bankcards"
	// BinariesRoute - route for  binary data
	BinariesRoute = "/binaries"
	// CredentialsRoute - route for  credentials
	CredentialsRoute = "/credentials"
	// NoteRoute - route for  note
	NoteRoute = "/notes/{id}"
	// NotesRoute - route for bank card
	BankcardRoute = "/bankcards/{id}"
	// BinaryRoute - route for  binary data
	BinaryRoute = "/binaries/{id}"
	// CredentialRoute - route for credentials
	CredentialRoute = "/credentials/{id}"
)
