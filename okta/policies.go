package okta

import (
	"fmt"
	"time"
)

// PoliciesService handles communication with the Policy data related
// methods of the OKTA API.
type PoliciesService service

// POLICIES

// Return the PasswordPolicy object. Used to create & update the password policy
func (p *PoliciesService) PasswordPolicy() PasswordPolicy {
	return PasswordPolicy{}
}

// Return the SignOnPolicy object. Used to create & update the signon policy
func (p *PoliciesService) SignOnPolicy() SignOnPolicy {
	return SignOnPolicy{}
}

// Return the MfaPolicy object. Used to create & update the mfa policy
func (p *PoliciesService) MfaPolicy() MfaPolicy {
	return MfaPolicy{}
}

// PasswordPolicy represents the Policy Object from the OKTA API
// used to create or update a password policy
type PasswordPolicy struct {
	ID          string    `json:"id,omitempty"`
	Type        string    `json:"type,omitempty"`
	Name        string    `json:"name,omitempty"`
	System      bool      `json:"system,omitempty"`
	Description string    `json:"description,omitempty"`
	Priority    int       `json:"priority,omitempty"`
	Status      string    `json:"status,omitempty"`
	Created     time.Time `json:"created,omitempty"`
	LastUpdated time.Time `json:"lastUpdated,omitempty"`
	Conditions  struct {
		People       `json:"people,omitempty"`
		AuthProvider `json:"authProvider,omitempty"`
	} `json:"conditions,omitempty"`
	Settings struct {
		Password   `json:"password,omitempty"`
		Recovery   `json:"recovery,omitempty"`
		Delegation `json:"delegation,omitempty"`
	} `json:"settings,omitempty"`
}

// SignOnPolicy represents the Policy Object from the OKTA API
// used to create or update a signon policy
type SignOnPolicy struct {
	ID          string    `json:"id,omitempty"`
	Type        string    `json:"type,omitempty"`
	Name        string    `json:"name,omitempty"`
	System      bool      `json:"system,omitempty"`
	Description string    `json:"description,omitempty"`
	Priority    int       `json:"priority,omitempty"`
	Status      string    `json:"status,omitempty"`
	Created     time.Time `json:"created,omitempty"`
	LastUpdated time.Time `json:"lastUpdated,omitempty"`
	Conditions  struct {
		People `json:"people,omitempty"`
	} `json:"conditions,omitempty"`
}

// MfaPolicy represents the Policy Object from the OKTA API
// used to create or update a mfa policy
type MfaPolicy struct {
	ID          string    `json:"id,omitempty"`
	Type        string    `json:"type,omitempty"`
	Name        string    `json:"name,omitempty"`
	System      bool      `json:"system,omitempty"`
	Description string    `json:"description,omitempty"`
	Priority    int       `json:"priority,omitempty"`
	Status      string    `json:"status,omitempty"`
	Created     time.Time `json:"created,omitempty"`
	LastUpdated time.Time `json:"lastUpdated,omitempty"`
	Conditions  struct {
		People `json:"people,omitempty"`
	} `json:"conditions,omitempty"`
	Settings struct {
		Factors `json:"factors,omitempty"`
	} `json:"settings,omitempty"`
}

// Policy represents the complete Policy Object from the OKTA API
// used to return policy data from a GET request
type Policy struct {
	ID          string    `json:"id,omitempty"`
	Type        string    `json:"type,omitempty"`
	Name        string    `json:"name,omitempty"`
	System      bool      `json:"system,omitempty"`
	Description string    `json:"description,omitempty"`
	Priority    int       `json:"priority,omitempty"`
	Status      string    `json:"status,omitempty"`
	Created     time.Time `json:"created,omitempty"`
	LastUpdated time.Time `json:"lastUpdated,omitempty"`
	Conditions  struct {
		People struct {
			Groups struct {
				Include []string `json:"include,omitempty"`
				Exclude []string `json:"exclude,omitempty"`
			} `json:"groups,omitempty"`
			Users struct {
				Include []string `json:"include,omitempty"`
				Exclude []string `json:"exclude,omitempty"`
			} `json:"users,omitempty"`
		} `json:"people,omitempty"`
		AuthType string `json:"authType,omitempty"`
		Network  struct {
			Connection string   `json:"connection,omitempty"`
			Include    []string `json:"include,omitempty"`
			Exclude    []string `json:"exclude,omitempty"`
		} `json:"network,omitempty"`
		AuthProvider `json:"authProvider,omitempty"`
	} `json:"conditions,omitempty"`
	Settings struct {
		Factors    `json:"factors,omitempty"`
		Password   `json:"password,omitempty"`
		Recovery   `json:"recovery,omitempty"`
		Delegation `json:"delegation,omitempty"`
	} `json:"settings,omitempty"`
	Links `json:"_links,omitempty"`
}

// Mfa policy settings factors obj
type Factors struct {
	GoogleOtp struct {
		Consent `json:"consent,omitempty"`
		Enroll  `json:"enroll,omitempty"`
	} `json:"google_otp,omitempty"`
	OktaOtp struct {
		Consent `json:"consent,omitempty"`
		Enroll  `json:"enroll,omitempty"`
	} `json:"okta_otp,omitempty"`
	OktaPush struct {
		Consent `json:"consent,omitempty"`
		Enroll  `json:"enroll,omitempty"`
	} `json:"okta_push,omitempty"`
	OktaQuestion struct {
		Consent `json:"consent,omitempty"`
		Enroll  `json:"enroll,omitempty"`
	} `json:"okta_question,omitempty"`
	OktaSms struct {
		Consent `json:"consent,omitempty"`
		Enroll  `json:"enroll,omitempty"`
	} `json:"okta_sms,omitempty"`
	RsaToken struct {
		Consent `json:"consent,omitempty"`
		Enroll  `json:"enroll,omitempty"`
	} `json:"rsa_token,omitempty"`
	SymantecVip struct {
		Consent `json:"consent,omitempty"`
		Enroll  `json:"enroll,omitempty"`
	} `json:"symantec_vip,omitempty"`
}

// Mfa policy factors consent obj
type Consent struct {
	Terms struct {
		Format string `json:"format,omitempty"`
		Value  string `json:"value,omitempty"`
	} `json:"terms,omitempty"`
	Type string `json:"type,omitempty"`
}

// Mfa policy & rule factors enroll obj
type Enroll struct {
	Self string `json:"self,omitempty"`
}

// Password policy settings password obj
type Password struct {
	Complexity struct {
		MinLength         int      `json:"minLength,omitempty"`
		MinLowerCase      int      `json:"minLowerCase,omitempty"`
		MinUpperCase      int      `json:"minUpperCase,omitempty"`
		MinNumber         int      `json:"minNumber,omitempty"`
		MinSymbol         int      `json:"minSymbol,omitempty"`
		ExcludeUsername   bool     `json:"excludeUsername,omitempty"`
		ExcludeAttributes []string `json:"excludeAttributes,omitempty"`
		Dictionary        struct {
			Common struct {
				Exclude bool `json:"exclude,omitempty"`
			} `json:"common,omitempty"`
		} `json:"dictionary,omitempty"`
	} `json:"complexity,omitempty"`
	Age struct {
		MaxAgeDays     int `json:"maxAgeDays,omitempty"`
		ExpireWarnDays int `json:"expireWarnDays,omitempty"`
		MinAgeMinutes  int `json:"minAgeMinutes,omitempty"`
		HistoryCount   int `json:"historyCount,omitempty"`
	} `json:"age,omitempty"`
	Lockout struct {
		MaxAttempts         int  `json:"maxAttempts,omitempty"`
		AutoUnlockMinutes   int  `json:"autoUnlockMinutes,omitempty"`
		ShowLockoutFailures bool `json:"showLockoutFailures,omitempty"`
	} `json:"lockout,omitempty"`
}

// Passowrd policy settings recover obj
type Recovery struct {
	Factors struct {
		RecoveryQuestion struct {
			Status     string `json:"status,omitempty"`
			Properties struct {
				Complexity struct {
					MinLength int `json:"minLength,omitempty"`
				} `json:"complexity,omitempty"`
			} `json:"properties,omitempty"`
		} `json:"recovery_question,omitempty"`
		OktaEmail struct {
			Status     string `json:"status,omitempty"`
			Properties struct {
				RecoveryToken struct {
					TokenLifetimeMinutes int `json:"tokenLifetimeMinutes,omitempty"`
				} `json:"recoveryToken,omitempty"`
			} `json:"properties,omitempty"`
		} `json:"okta_email,omitempty"`
		OktaSms struct {
			Status string `json:"status,omitempty"`
		} `json:"okta_sms,omitempty"`
	} `json:"factors,omitempty"`
}

// password policy settings delegation obj
type Delegation struct {
	Options struct {
		SkipUnlock bool `json:"skipUnlock,omitempty"`
	} `json:"options,omitempty"`
}

// policy & rule conditions people obj
// when creating an obj, Groups & Users are exclusive
type People struct {
	*Groups `json:"groups,omitempty"`
	*Users  `json:"users,omitempty"`
}

// policy & rule conditions people groups obj
// when creating an obj, Include & Exclude are exclusive
type Groups struct {
	Include *[]string `json:"include,omitempty"`
	Exclude *[]string `json:"exclude,omitempty"`
}

// policy & rule conditions people users obj
// when creating an obj, Include & Exclude are exclusive
type Users struct {
	Include *[]string `json:"include,omitempty"`
	Exclude *[]string `json:"exclude,omitempty"`
}

// policy & rule conditions network obj
// when creating an obj, Include & Exclude are exclusive
// TODO: Include & Exclude not supported as only needed when
// Connection is "ZONE". zone requires the zone api (not implemented atm)
type Network struct {
	Connection string    `json:"connection,omitempty"`
	Include    *[]string `json:"include,omitempty"`
	Exclude    *[]string `json:"exclude,omitempty"`
}

// policy & rule conditions authProvider obj
type AuthProvider struct {
	Provider string   `json:"provider,omitempty"`
	Include  []string `json:"include,omitempty"`
}

// a slice of Policy objs
// used by GetPolicesByType
type policies struct {
	Policies []Policy `json:",-"`
}

// Policy & Rule obj use the same links obj
type Links struct {
	Self struct {
		Href  string `json:"href,omitempty"`
		Hints struct {
			Allow []string `json:"allow,omitempty"`
		} `json:"hints,omitempty"`
	} `json:"self,omitempty"`
	Activate struct {
		Href  string `json:"href,omitempty"`
		Hints struct {
			Allow []string `json:"allow,omitempty"`
		} `json:"hints,omitempty"`
	} `json:"activate",omitempty`
	Deactivate struct {
		Href  string `json:"href,omitempty"`
		Hints struct {
			Allow []string `json:"allow,omitempty"`
		} `json:"hints,omitempty"`
	} `json:"deactivate,omitempty"`
	Rules struct {
		Href  string `json:"href,omitempty"`
		Hints struct {
			Allow []string `json:"allow,omitempty"`
		} `json:"hints,omitempty"`
	} `json:"rules,omitempty"`
}

//RULES

// Return the PasswordRule object. Used to create & update the password rule
func (p *PoliciesService) PasswordRule() PasswordRule {
	return PasswordRule{}
}

// Return the SignOnRule object. Used to create & update the signon rule
func (p *PoliciesService) SignOnRule() SignOnRule {
	return SignOnRule{}
}

// Return the MfaRule object. Used to create & update the mfa rule
func (p *PoliciesService) MfaRule() MfaRule {
	return MfaRule{}
}

// PasswordRule represents the Rule Object from the OKTA API
// used to create or update a password rule
type PasswordRule struct {
	ID          string    `json:"id,omitempty"`
	Type        string    `json:"type,omitempty"`
	Status      string    `json:"status,omitempty"`
	Name        string    `json:"name,omitempty"`
	Priority    int       `json:"priority,omitempty"`
	System      bool      `json:"system,omitempty"`
	Created     time.Time `json:"created,omitempty"`
	LastUpdated time.Time `json:"lastUpdated,omitempty"`
	Conditions  struct {
		People  `json:"people,omitempty"`
		Network `json:"network,omitempty"`
	} `json:"conditions,omitempty"`
	Actions struct {
		PasswordChange           PasswordAction `json:"passwordChange,omitempty"`
		SelfServicePasswordReset PasswordAction `json:"selfServicePasswordReset,omitempty"`
		SelfServiceUnlock        PasswordAction `json:"selfServiceUnlock,omitempty"`
	} `json:"actions,omitempty"`
}

// SignOnRule represents the Rule Object from the OKTA API
// used to create or update a signon rule
type SignOnRule struct {
	ID          string    `json:"id,omitempty"`
	Type        string    `json:"type,omitempty"`
	Status      string    `json:"status,omitempty"`
	Name        string    `json:"name,omitempty"`
	Priority    int       `json:"priority,omitempty"`
	System      bool      `json:"system,omitempty"`
	Created     time.Time `json:"created,omitempty"`
	LastUpdated time.Time `json:"lastUpdated,omitempty"`
	Conditions  struct {
		People   `json:"people,omitempty"`
		Network  `json:"network,omitempty"`
		AuthType string `json:"authType,omitempty"`
	} `json:"conditions,omitempty"`
	Actions struct {
		SignOn `json:"signon,omitempty"`
	} `json:"actions,omitempty"`
}

// MfaRule represents the Rule Object from the OKTA API
// used to create or update a mfa rule
type MfaRule struct {
	ID          string    `json:"id,omitempty"`
	Type        string    `json:"type,omitempty"`
	Status      string    `json:"status,omitempty"`
	Name        string    `json:"name,omitempty"`
	Priority    int       `json:"priority,omitempty"`
	System      bool      `json:"system,omitempty"`
	Created     time.Time `json:"created,omitempty"`
	LastUpdated time.Time `json:"lastUpdated,omitempty"`
	Conditions  struct {
		People  `json:"people,omitempty"`
		Network `json:"network,omitempty"`
	} `json:"conditions,omitempty"`
	Actions struct {
		Enroll `json:"enroll,omitempty"`
	} `json:"actions,omitempty"`
}

// Rule represents the complete Rule Object from the OKTA API
// used to return rule data from a GET request
type Rule struct {
	ID          string    `json:"id,omitempty"`
	Type        string    `json:"type,omitempty"`
	Status      string    `json:"status,omitempty"`
	Priority    int       `json:"priority,omitempty"`
	Name        string    `json:"name,omitempty"`
	System      bool      `json:"system,omitempty"`
	Created     time.Time `json:"created,omitempty"`
	LastUpdated time.Time `json:"lastUpdated,omitempty"`
	Conditions  struct {
		People struct {
			Groups struct {
				Include []string `json:"include,omitempty"`
				Exclude []string `json:"exclude,omitempty"`
			} `json:"groups,omitempty"`
			Users struct {
				Include []string `json:"include,omitempty"`
				Exclude []string `json:"exclude,omitempty"`
			} `json:"users,omitempty"`
		} `json:"people,omitempty"`
		AuthType string `json:"authType,omitempty"`
		Network  struct {
			Connection string   `json:"connection,omitempty"`
			Include    []string `json:"include,omitempty"`
			Exclude    []string `json:"exclude,omitempty"`
		} `json:"network,omitempty"`
		AuthProvider `json:"authProvider,omitempty"`
	} `json:"conditions,omitempty"`
	Actions struct {
		SignOn                   `json:"signon,omitempty"`
		Enroll                   `json:"enroll,omitempty"`
		PasswordChange           PasswordAction `json:"passwordChange,omitempty"`
		SelfServicePasswordReset PasswordAction `json:"selfServicePasswordReset,omitempty"`
		SelfServiceUnlock        PasswordAction `json:"selfServiceUnlock,omitempty"`
	} `json:"actions,omitempty"`
	Links `json:"_links,omitempty"`
}

// signon rule actions signon obj
type SignOn struct {
	Access                  string `json:"access,omitempty"`
	RequireFactor           bool   `json:"requireFactor,omitempty"`
	FactorPromptMode        string `json:"factorPromptMode,omitempty"`
	RememberDeviceByDefault bool   `json:"rememberDeviceByDefault,omitempty"`
	FactorLifetime          int    `json:"factorLifetime,omitempty"`
	Session                 struct {
		MaxSessionIdleMinutes     int  `json:"maxSessionIdleMinutes,omitempty"`
		MaxSessionLifetimeMinutes int  `json:"maxSessionLifetimeMinutes,omitempty"`
		UsePersistentCookie       bool `json:"usePersistentCookie,omitempty"`
	} `json:"session,omitempty"`
}

// rule actions for passwords use the same passwordAction obj
type PasswordAction struct {
	Access string `json:"access,omitempty"`
}

// a slice of Rule objs
// used by GetPolicyRules
type rules struct {
	Rules []Rule `json:",-"`
}

// API FUNCTIONS

// UsersCondition updates the People Users condition for the input policy or rule
// requires inputs string "include" or "exclude" & a string slice of Okta user IDs
func UsersCondition(clude string, values []string) (*Users, error) {
	var pop *Users
	switch {
	case clude == "include":
		pop = &Users{Include: &values}
	case clude == "exclude":
		pop = &Users{Exclude: &values}
	default:
		return nil, fmt.Errorf("[ERROR] UsersCondition input string var supports values \"include\" or \"exclude\"")
	}
	return pop, nil
}

// GroupsCondition updates the People Groups condition for the input policy or rule
// requires inputs string "include" or "exclude" & a string slice of Okta group IDs
func GroupsCondition(clude string, values []string) (*Groups, error) {
	var pop *Groups
	switch {
	case clude == "include":
		pop = &Groups{Include: &values}
	case clude == "exclude":
		pop = &Groups{Exclude: &values}
	default:
		return nil, fmt.Errorf("[ERROR] GroupsCondition input string var supports values \"include\" or \"exclude\"")
	}
	return pop, nil
}

// PeopleCondition updates the People condition for the input policy or rule
// requires inputs string "users" or "groups & "include" or "exclude"
// plus a string slice of Okta group or user IDs
func PeopleCondition(condition string, clude string, values []string) (*People, error) {
	var pop *People
	switch {
	case condition == "users":
		var users *Users
		users, err := UsersCondition(clude, values)
		if err != nil {
			return nil, err
		}
		pop = &People{Users: users}
	case condition == "groups":
		var groups *Groups
		groups, err := GroupsCondition(clude, values)
		if err != nil {
			return nil, err
		}
		pop = &People{Groups: groups}
	default:
		return nil, fmt.Errorf("[ERROR] PeopleCondition input string var supports values \"users\" or \"groups\"")
	}
	return pop, nil
}

// PasswordPolicy PeopleCondition updates the People condition for the input password policy
// requires inputs string "users" or "grooups & "include" or "exclude"
// plus a string slice of Okta group or user IDs
func (p *PasswordPolicy) PeopleCondition(condition string, clude string, values []string) error {
	pop, err := PeopleCondition(condition, clude, values)
	if err != nil {
		return err
	}
	p.Conditions.People = *pop
	return nil
}

// SignOnPolicy PeopleCondition updates the People condition for the input signon policy
// requires inputs string "users" or "grooups & "include" or "exclude"
// plus a string slice of Okta group or user IDs
func (p *SignOnPolicy) PeopleCondition(condition string, clude string, values []string) error {
	pop, err := PeopleCondition(condition, clude, values)
	if err != nil {
		return err
	}
	p.Conditions.People = *pop
	return nil
}

// MfaPolicy PeopleCondition updates the People condition for the input mfa policy
// requires inputs string "users" or "grooups & "include" or "exclude"
// plus a string slice of Okta group or user IDs
func (p *MfaPolicy) PeopleCondition(condition string, clude string, values []string) error {
	pop, err := PeopleCondition(condition, clude, values)
	if err != nil {
		return err
	}
	p.Conditions.People = *pop
	return nil
}

// PasswordRule PeopleCondition updates the People condition for the input password rule
// requires inputs string "users" or "grooups & "include" or "exclude"
// plus a string slice of Okta group or user IDs
func (p *PasswordRule) PeopleCondition(condition string, clude string, values []string) error {
	pop, err := PeopleCondition(condition, clude, values)
	if err != nil {
		return err
	}
	p.Conditions.People = *pop
	return nil
}

// SIgnOnRule PeopleCondition updates the People condition for the input signon rule
// requires inputs string "users" or "grooups & "include" or "exclude"
// plus a string slice of Okta group or user IDs
func (p *SignOnRule) PeopleCondition(condition string, clude string, values []string) error {
	pop, err := PeopleCondition(condition, clude, values)
	if err != nil {
		return err
	}
	p.Conditions.People = *pop
	return nil
}

// MfaRule PeopleCondition updates the People condition for the input mfa rule
// requires inputs string "users" or "grooups & "include" or "exclude"
// plus a string slice of Okta group or user IDs
func (p *MfaRule) PeopleCondition(condition string, clude string, values []string) error {
	pop, err := PeopleCondition(condition, clude, values)
	if err != nil {
		return err
	}
	p.Conditions.People = *pop
	return nil
}

// GetPolicy: Get a policy
// Requires Policy ID from Policy object
func (p *PoliciesService) GetPolicy(id string) (*Policy, *Response, error) {
	u := fmt.Sprintf("policies/%v", id)
	req, err := p.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	policy := new(Policy)
	resp, err := p.client.Do(req, policy)
	if err != nil {
		return nil, resp, err
	}

	return policy, resp, err
}

// GetPolicesByType: Get all policies by type
// Allowed types are OKTA_SIGN_ON, PASSWORD, MFA_ENROLL, or OAUTH_AUTHORIZATION_POLICY
func (p *PoliciesService) GetPolicesByType(policyType string) (*policies, *Response, error) {
	u := fmt.Sprintf("policies?type=%v", policyType)
	req, err := p.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	policy := make([]Policy, 0)
	resp, err := p.client.Do(req, &policy)
	if err != nil {
		return nil, resp, err
	}
	if len(policy) > 0 {
		myPolicies := new(policies)
		for _, v := range policy {
			myPolicies.Policies = append(myPolicies.Policies, v)
		}
		return myPolicies, resp, err
	}

	return nil, resp, err
}

// DeletePolicy: Delete a policy
// Requires Policy ID from Policy object
func (p *PoliciesService) DeletePolicy(id string) (*Response, error) {
	u := fmt.Sprintf("policies/%v", id)
	req, err := p.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// CreatePolicy: Create a policy
// You must pass in the Policy object created from the desired input policy
func (p *PoliciesService) CreatePolicy(policy interface{}) (*Policy, *Response, error) {
	u := fmt.Sprintf("policies")
	req, err := p.client.NewRequest("POST", u, policy)
	if err != nil {
		return nil, nil, err
	}

	newPolicy := new(Policy)
	resp, err := p.client.Do(req, newPolicy)
	if err != nil {
		return nil, resp, err
	}

	return newPolicy, resp, err
}

// UpdatePolicy: Update a policy
// Requires Policy ID from Policy object & Policy object from the desired input policy
func (p *PoliciesService) UpdatePolicy(id string, policy interface{}) (*Policy, *Response, error) {
	u := fmt.Sprintf("policies/%v", id)
	req, err := p.client.NewRequest("PUT", u, policy)
	if err != nil {
		return nil, nil, err
	}

	updatePolicy := new(Policy)
	resp, err := p.client.Do(req, updatePolicy)
	if err != nil {
		return nil, resp, err
	}

	return updatePolicy, resp, err
}

// ActivatePolicy: Activate a policy
// Requires Policy ID from Policy object
func (p *PoliciesService) ActivatePolicy(id string) (*Response, error) {
	u := fmt.Sprintf("policies/%v/lifecycle/activate", id)
	req, err := p.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// DeactivatePolicy: Deactivate a policy
// Requires Policy ID from Policy object
func (p *PoliciesService) DeactivatePolicy(id string) (*Response, error) {
	u := fmt.Sprintf("policies/%v/lifecycle/deactivate", id)
	req, err := p.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// GetPolicyRules: Get policy rules
// Requires Policy ID from Policy object
func (p *PoliciesService) GetPolicyRules(id string) (*rules, *Response, error) {
	u := fmt.Sprintf("policies/%v/rules", id)
	req, err := p.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	rule := make([]Rule, 0)
	resp, err := p.client.Do(req, &rule)
	if err != nil {
		return nil, resp, err
	}
	if len(rule) > 0 {
		myRules := new(rules)
		for _, v := range rule {
			myRules.Rules = append(myRules.Rules, v)
		}
		return myRules, resp, err
	}

	return nil, resp, err
}

// CreatePolicyRule: Create a policy rule
// Requires Policy ID from Policy object
// You must pass in the Rule object created from the desired input rule
func (p *PoliciesService) CreatePolicyRule(id string, rule interface{}) (*Rule, *Response, error) {
	u := fmt.Sprintf("policies/%v/rules", id)
	req, err := p.client.NewRequest("POST", u, rule)
	if err != nil {
		return nil, nil, err
	}

	newRule := new(Rule)
	resp, err := p.client.Do(req, newRule)
	if err != nil {
		return nil, resp, err
	}

	return newRule, resp, err
}

// DeletePolicyRule: Delete a rule
// Requires Policy ID from Policy object and Rule ID from Rule object
func (p *PoliciesService) DeletePolicyRule(policyId string, ruleId string) (*Response, error) {
	u := fmt.Sprintf("policies/%v/rules/%v", policyId, ruleId)
	req, err := p.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// GetPolicyRule: Get a policy rule
// Requires Policy ID from Policy object and Rule ID from Rule object
func (p *PoliciesService) GetPolicyRule(policyId string, ruleId string) (*Rule, *Response, error) {
	u := fmt.Sprintf("policies/%v/rules/%v", policyId, ruleId)
	req, err := p.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	rule := new(Rule)
	resp, err := p.client.Do(req, rule)
	if err != nil {
		return nil, resp, err
	}

	return rule, resp, err
}

// UpdatePolicyRule: Update a policy rule
// Requires Policy ID from Policy object and Rule ID from Rule object
// You must pass in the Rule object from the desited input rule
func (p *PoliciesService) UpdatePolicyRule(policyId string, ruleId string, rule interface{}) (*Rule, *Response, error) {
	u := fmt.Sprintf("policies/%v/rules/%v", policyId, ruleId)
	req, err := p.client.NewRequest("PUT", u, rule)
	if err != nil {
		return nil, nil, err
	}

	updateRule := new(Rule)
	resp, err := p.client.Do(req, updateRule)
	if err != nil {
		return nil, resp, err
	}

	return updateRule, resp, err
}

// ActivatePolicyRule: Activate a policy rule
// Requires Policy ID from Policy object and Rule ID from Rule object
func (p *PoliciesService) ActivatePolicyRule(policyId string, ruleId string) (*Response, error) {
	u := fmt.Sprintf("policies/%v/rules/%v/lifecycle/activate", policyId, ruleId)
	req, err := p.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// DeactivatePolicyRule: Deactivate a policy rule
// Requires Policy ID from Policy object and Rule ID from Rule object
func (p *PoliciesService) DeactivatePolicyRule(policyId string, ruleId string) (*Response, error) {
	u := fmt.Sprintf("policies/%v/rules/%v/lifecycle/deactivate", policyId, ruleId)
	req, err := p.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}
