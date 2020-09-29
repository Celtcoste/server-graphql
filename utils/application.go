package utils

func GetApplication(application string) *string {
	var table string

	/* Depending the value of application in the header */
	if application == "announcer" {
		table = "announcer_account"
	} else if application == "influencer" {
		table = "influencer_account"
	} else if application == "partner" {
		table = "partner_account"
	} else {
		return nil
	}
	return &table
}