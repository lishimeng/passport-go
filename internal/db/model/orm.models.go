package model

func Tables() (t []interface{}) {
	t = append(t,
		new(UserProfile),
		new(Mfa),
		new(MfaItem),
		new(MfaDevice),
		new(SdkConfig),
		new(UserInfo),
		new(Notify),
		new(Application),
		new(Tenant),
		new(UserAppAuth),
	)
	return
}
