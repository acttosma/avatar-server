package req

// Register as a new account
type ActRegister struct {
	// E-Mail address [required]
	Mail string `json:"mail" binding:"required"`
	// Password, for safety reason,the request should mix this parameter with encryption,such as MD5 and SHA256. [required]
	Password string `json:"password" binding:"required"`
	// invite code [required]
	InviteCode string `json:"inviteCode" binding:"required"`
}

// Login with mail and password
type ActLogin struct {
	// E-Mail address [required]
	Mail string `json:"mail" binding:"required"`
	// Password, for safety reason,the request should mix this parameter with encryption,such as MD5 and SHA256. [required]
	Password string `json:"password" binding:"required"`
}

// Login with mail and password
type ChangePassword struct {
	// PreviousPassword, for safety reason,the request should mix this parameter with encryption,such as MD5 and SHA256. [required]
	PreviousPassword string `json:"prePassword" binding:"required"`
	// Password, for safety reason,the request should mix this parameter with encryption,such as MD5 and SHA256. [required]
	Password string `json:"password" binding:"required"`
}
