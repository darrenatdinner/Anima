package user

// IUser - Interface For A Single User Interacting With The Server Via HTTP
type IUser interface {
	Authenticate() (bool, error)
	//CheckAPIAccess(*persist.IDataContext, string) (bool, error)
	GetStamp() string
}
