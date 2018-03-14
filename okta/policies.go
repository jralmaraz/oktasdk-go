package okta

import (
	"time"
)

// PoliciesService handles communication with the Policy data related
// methods of the OKTA API.
type PoliciesService service

// Policy represents the Policy Object from the OKTA API
type Policy struct {
	ID          string      `json:"id,omitempty"`
	Type        string      `json:"type"`
	Name        string      `json:"name"`
	System      string      `json:"system,omitempty"`
	Description string      `json:"description,omitempty"`
	Priority    int         `json:"priority,omitempty"`
	Status      string      `json:"status,omitempty"`
	Conditions  conditions  `json:"conditions,omitempty"`
	Settings    settings    `json:"settings,omitempty"`
	Created     time.Time   `json:"created,omitempty"`
	LastUpdated time.Time   `json:"lastUpdated,omitempty"`
	Links       policyLinks `json:"_links,omitempty"`
}

type conditions struct {
	People struct {
		Groups groups `json:"groups,omitempty"`
		Users  users  `json:"users,omitempty"`
	} `json:"people,omitempty"`
	AuthContext  string       `json:"authType,omitempty"`
	Network      network      `json:"network,omitempty"`
	AuthProvider authProvider `json:"authProvider,omitempty"`
}

type users struct {
	Include []string `json:"include,omitempty"`
	Exclude []string `json:"exclude,omitempty"`
}

type groups struct {
	Include []string `json:"include,omitempty"`
	Exclude []string `json:"exclude,omitempty"`
}

type network struct {
	Connection string   `json:"connection,omitempty"`
	Include    []string `json:"include,omitempty"`
	Exclude    []string `json:"exclude,omitempty"`
}

type authProvider struct {
	Provider string   `json:"provider"`
	Include  []string `json:"include,omitempty"`
}

type settings struct {
	Factors struct {
		GoogleOtp    mfaFactor `json:"google_otp,omitempty"`
		OktaOtp      mfaFactor `json:"okta_otp,omitempty"`
		OktaPush     mfaFactor `json:"okta_push,omitempty"`
		OktaQuestion mfaFactor `json:"okta_question,omitempty"`
		OktaSms      mfaFactor `json:"okta_sms,omitempty"`
		RsaToken     mfaFactor `json:"rsa_token,omitempty"`
		SymantecVip  mfaFactor `json:"symantec_vip,omitempty"`
	} `json:"factors,omitempty"`
	Password struct {
		password   password   `json:"password,omitempty"`
		recovery   recovery   `json:"recovery,omitempty"`
		delegation delegation `json:"lockout,omitempty"`
	} `json:"password,omitempty"`
}

type password struct {
	Complexity complexity `json:"complexity,omitempty"`
	Age        age        `json:"age,omitempty"`
	Lockout    lockout    `json:"lockout,omitempty"`
}

type complexity struct {
	MinLength         int        `json:"minLength,omitempty"`
	MinLowerCase      int        `json:"minLowerCase,omitempty"`
	MinUpperCase      int        `json:"minUpperCase,omitempty"`
	MinNumber         int        `json:"minNumber,omitempty"`
	MinSymbol         int        `json:"minSymbol,omitempty"`
	ExcludeUsername   bool       `json:"excludeUsername,omitempty"`
	ExcludeAttributes []string   `json:"excludeAttributes,omitempty"`
	Dictionary        dictionary `json:"dictionary,omitempty"`
}

type dictionary struct {
	Common common `json:"common,omitempty"`
}

type common struct {
	Exclude bool `json:"excllude,omitempty"`
}

type age struct {
	MaxAgeDays     int `json:"maxAgeDays,omitempty"`
	ExpireWarnDays int `json:"expireWarnDays,omitempty"`
	MinAgeMinutes  int `json:"minAgeMinutes,omitempty"`
	HistoryCount   int `json:"historyCount,omitempty"`
}

type lockout struct {
	MaxAttempts         int  `json:"maxAttempts,omitempty"`
	AutoUnlockMinutes   int  `json:"autoUnlockMinutes,omitempty"`
	ShowLockoutFailures bool `json:"showLockoutFailures,omitempty"`
}

type recovery struct {
	Factors struct {
		RecoveryQuestion policyRecoveryQuestion `json:"recovery_question,omitempty"`
		OktaEmail        oktaEmail              `json:"okta_email,omitempty"`
		OktaSms          oktaSms                `json:"okta_sms,omitempty"`
	} `json:"factors,omitempty"`
}

type policyRecoveryQuestion struct {
	Status     string                     `json:"status"`
	Properties recoveryQuestionProperties `json:"properties,omitempty"`
}

type recoveryQuestionProperties struct {
	Complexity recoveryQuestionPropertiesComplexity `json:"complexity,omitempty"`
}

type recoveryQuestionPropertiesComplexity struct {
	MinLength int `json:"minLength,omitempty"`
}

type oktaEmail struct {
	Status     string              `json:"status"`
	Properties oktaEmailProperties `json:"properties,omitempty"`
}

type oktaEmailProperties struct {
	RecoveryToken oktaEmailPropertiesRecoveryToken `json:"recoveryToken,omitempty"`
}

type oktaEmailPropertiesRecoveryToken struct {
	TokenLifetimeMinutes int `json:"tokenLifetimeMinutes,omitempty"`
}

type oktaSms struct {
	Status string `json:"status,omitempty"`
}

type delegation struct {
	Options options `json:"options,omitempty"`
}

type options struct {
	SkipUnlock bool `json:"skipUnlock,omitempty"`
}

type mfaFactor struct {
	Consent factorConsent `json:"consent,omitempty"`
	Enroll  factorEnroll  `json:"enroll,omitempty"`
}

type factorConsent struct {
	Terms factorConsentTerms `json:"terms,omitempty"`
	Type  string             `json:"type,omitempty"`
}

type factorEnroll struct {
	Self string `json:"self,omitempty"`
}

type factorConsentTerms struct {
	Format string `json:"format,omitempty"`
	Value  string `json:"value,omitempty"`
}

// TODO add MFA conditions objects
// https://developer.okta.com/docs/api/resources/policy#policy-conditions-1

// TODO: policy links object is not complete?
// https://developer.okta.com/docs/api/resources/policy#LinksObject
type policyLinks struct {
	Self       string `json:"self"`
	Activate   string `json:"activate",omitempty`
	Deactivate string `json:"deactivate,omitempty"`
	Rules      string `json:"rules,omitempty"`
}

// Rules represent the Rules Object from the OKTA API
type Rules struct {
	ID          string     `json:"id,omitempty"`
	Type        string     `json:"type"`
	Status      string     `json:"status,omitempty"`
	Priority    int        `json:"priority,omitempty"`
	System      string     `json:"system,omitempty"`
	Created     time.Time  `json:"created,omitempty"`
	LastUpdated time.Time  `json:"lastUpdated,omitempty"`
	Conditions  conditions `json:"conditions,omitempty"`
	Actions     actions    `json:"actions,omitempty"`
	Links       ruleLinks  `json:"_links,omitempty"`
}

type actions struct {
	signon struct {
		Access                  string  `json:"access"`
		RequireFactor           bool    `json:"requireFactor,omitempty"`
		FactorPromptMode        string  `json:"factorPromptMode,omitempty"`
		RememberDeviceByDefault bool    `json:"rememberDeviceByDefault,omitempty"`
		FactorLifetime          int     `json:"factorLifetime,omitempty"`
		Session                 session `json:"session,omitempty"`
	} `json:"signon,omitempty"`
	enroll struct {
		Self string `json:"self,omitempty"`
	} `json:"enroll,omitempty"`
	PasswordChange           passwordAction `json:"passwordChange,omitempty"`
	SelfServicePasswordReset passwordAction `json:"selfServicePasswordReset,omitempty"`
	SelfServiceUnlock        passwordAction `json:"selfServiceUnlock,omitempty"`
}

type passwordAction struct {
	Access string `json:"access,omitempty"`
}

type session struct {
	MaxSessionIdleMinutes     int  `json:"maxSessionIdleMinutes,omitempty"`
	MaxSessionLifetimeMinutes int  `json:"maxSessionLifetimeMinutes,omitempty"`
	UsePersistentCookie       bool `json:"usePersistentCookie,,omitempty"`
}

// TODO: rule links object is not complete
// https://developer.okta.com/docs/api/resources/policy#RulesLinksObject
type ruleLinks struct {
	Self       string `json:"self"`
	Activate   string `json:"activate,omitempty"`
	Deactivate string `json:"deactivate,omitempty"`
}
