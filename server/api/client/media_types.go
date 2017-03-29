// Code generated by goagen v1.1.0, command line:
// $ main
//
// API "fieldkit": Application Media Types
//
// The content of this file is auto-generated, DO NOT MODIFY

package client

import (
	"github.com/goadesign/goa"
	"net/http"
	"unicode/utf8"
)

// ProjectAdministrator media type (default view)
//
// Identifier: application/vnd.app.administrator+json; view=default
type ProjectAdministrator struct {
	ProjectID int `form:"project_id" json:"project_id" xml:"project_id"`
	UserID    int `form:"user_id" json:"user_id" xml:"user_id"`
}

// Validate validates the ProjectAdministrator media type instance.
func (mt *ProjectAdministrator) Validate() (err error) {

	return
}

// DecodeProjectAdministrator decodes the ProjectAdministrator instance encoded in resp body.
func (c *Client) DecodeProjectAdministrator(resp *http.Response) (*ProjectAdministrator, error) {
	var decoded ProjectAdministrator
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// ProjectAdministratorCollection is the media type for an array of ProjectAdministrator (default view)
//
// Identifier: application/vnd.app.administrator+json; type=collection; view=default
type ProjectAdministratorCollection []*ProjectAdministrator

// Validate validates the ProjectAdministratorCollection media type instance.
func (mt ProjectAdministratorCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeProjectAdministratorCollection decodes the ProjectAdministratorCollection instance encoded in resp body.
func (c *Client) DecodeProjectAdministratorCollection(resp *http.Response) (ProjectAdministratorCollection, error) {
	var decoded ProjectAdministratorCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// ProjectAdministrators media type (default view)
//
// Identifier: application/vnd.app.administrators+json; view=default
type ProjectAdministrators struct {
	Administrators ProjectAdministratorCollection `form:"administrators" json:"administrators" xml:"administrators"`
}

// Validate validates the ProjectAdministrators media type instance.
func (mt *ProjectAdministrators) Validate() (err error) {
	if mt.Administrators == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "administrators"))
	}
	return
}

// DecodeProjectAdministrators decodes the ProjectAdministrators instance encoded in resp body.
func (c *Client) DecodeProjectAdministrators(resp *http.Response) (*ProjectAdministrators, error) {
	var decoded ProjectAdministrators
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Expedition media type (default view)
//
// Identifier: application/vnd.app.expedition+json; view=default
type Expedition struct {
	Description string `form:"description" json:"description" xml:"description"`
	ID          int    `form:"id" json:"id" xml:"id"`
	Name        string `form:"name" json:"name" xml:"name"`
	Slug        string `form:"slug" json:"slug" xml:"slug"`
}

// Validate validates the Expedition media type instance.
func (mt *Expedition) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Slug == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "slug"))
	}
	if mt.Description == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "description"))
	}
	if ok := goa.ValidatePattern(`^[[:alnum:]]+(-[[:alnum:]]+)*$`, mt.Slug); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.slug`, mt.Slug, `^[[:alnum:]]+(-[[:alnum:]]+)*$`))
	}
	if utf8.RuneCountInString(mt.Slug) > 40 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.slug`, mt.Slug, utf8.RuneCountInString(mt.Slug), 40, false))
	}
	return
}

// DecodeExpedition decodes the Expedition instance encoded in resp body.
func (c *Client) DecodeExpedition(resp *http.Response) (*Expedition, error) {
	var decoded Expedition
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// ExpeditionCollection is the media type for an array of Expedition (default view)
//
// Identifier: application/vnd.app.expedition+json; type=collection; view=default
type ExpeditionCollection []*Expedition

// Validate validates the ExpeditionCollection media type instance.
func (mt ExpeditionCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeExpeditionCollection decodes the ExpeditionCollection instance encoded in resp body.
func (c *Client) DecodeExpeditionCollection(resp *http.Response) (ExpeditionCollection, error) {
	var decoded ExpeditionCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// Expeditions media type (default view)
//
// Identifier: application/vnd.app.expeditions+json; view=default
type Expeditions struct {
	Expeditions ExpeditionCollection `form:"expeditions" json:"expeditions" xml:"expeditions"`
}

// Validate validates the Expeditions media type instance.
func (mt *Expeditions) Validate() (err error) {
	if mt.Expeditions == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "expeditions"))
	}
	if err2 := mt.Expeditions.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// DecodeExpeditions decodes the Expeditions instance encoded in resp body.
func (c *Client) DecodeExpeditions(resp *http.Response) (*Expeditions, error) {
	var decoded Expeditions
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Inputs media type (default view)
//
// Identifier: application/vnd.app.inputs+json; view=default
type Inputs struct {
	TwitterAccounts TwitterAccountCollection `form:"twitter_accounts,omitempty" json:"twitter_accounts,omitempty" xml:"twitter_accounts,omitempty"`
}

// Validate validates the Inputs media type instance.
func (mt *Inputs) Validate() (err error) {
	if err2 := mt.TwitterAccounts.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// DecodeInputs decodes the Inputs instance encoded in resp body.
func (c *Client) DecodeInputs(resp *http.Response) (*Inputs, error) {
	var decoded Inputs
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Location media type (default view)
//
// Identifier: application/vnd.app.location+json; view=default
type Location struct {
	Location string `form:"location" json:"location" xml:"location"`
}

// Validate validates the Location media type instance.
func (mt *Location) Validate() (err error) {
	if mt.Location == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "location"))
	}
	if err2 := goa.ValidateFormat(goa.FormatURI, mt.Location); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.location`, mt.Location, goa.FormatURI, err2))
	}
	return
}

// DecodeLocation decodes the Location instance encoded in resp body.
func (c *Client) DecodeLocation(resp *http.Response) (*Location, error) {
	var decoded Location
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// TeamMember media type (default view)
//
// Identifier: application/vnd.app.member+json; view=default
type TeamMember struct {
	Role   string `form:"role" json:"role" xml:"role"`
	TeamID int    `form:"team_id" json:"team_id" xml:"team_id"`
	UserID int    `form:"user_id" json:"user_id" xml:"user_id"`
}

// Validate validates the TeamMember media type instance.
func (mt *TeamMember) Validate() (err error) {

	if mt.Role == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "role"))
	}
	return
}

// DecodeTeamMember decodes the TeamMember instance encoded in resp body.
func (c *Client) DecodeTeamMember(resp *http.Response) (*TeamMember, error) {
	var decoded TeamMember
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// TeamMemberCollection is the media type for an array of TeamMember (default view)
//
// Identifier: application/vnd.app.member+json; type=collection; view=default
type TeamMemberCollection []*TeamMember

// Validate validates the TeamMemberCollection media type instance.
func (mt TeamMemberCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeTeamMemberCollection decodes the TeamMemberCollection instance encoded in resp body.
func (c *Client) DecodeTeamMemberCollection(resp *http.Response) (TeamMemberCollection, error) {
	var decoded TeamMemberCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// TeamMembers media type (default view)
//
// Identifier: application/vnd.app.members+json; view=default
type TeamMembers struct {
	Members TeamMemberCollection `form:"members" json:"members" xml:"members"`
}

// Validate validates the TeamMembers media type instance.
func (mt *TeamMembers) Validate() (err error) {
	if mt.Members == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "members"))
	}
	if err2 := mt.Members.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// DecodeTeamMembers decodes the TeamMembers instance encoded in resp body.
func (c *Client) DecodeTeamMembers(resp *http.Response) (*TeamMembers, error) {
	var decoded TeamMembers
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Project media type (default view)
//
// Identifier: application/vnd.app.project+json; view=default
type Project struct {
	Description string `form:"description" json:"description" xml:"description"`
	ID          int    `form:"id" json:"id" xml:"id"`
	Name        string `form:"name" json:"name" xml:"name"`
	Slug        string `form:"slug" json:"slug" xml:"slug"`
}

// Validate validates the Project media type instance.
func (mt *Project) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Slug == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "slug"))
	}
	if mt.Description == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "description"))
	}
	if ok := goa.ValidatePattern(`^[[:alnum:]]+(-[[:alnum:]]+)*$`, mt.Slug); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.slug`, mt.Slug, `^[[:alnum:]]+(-[[:alnum:]]+)*$`))
	}
	if utf8.RuneCountInString(mt.Slug) > 40 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.slug`, mt.Slug, utf8.RuneCountInString(mt.Slug), 40, false))
	}
	return
}

// DecodeProject decodes the Project instance encoded in resp body.
func (c *Client) DecodeProject(resp *http.Response) (*Project, error) {
	var decoded Project
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// ProjectCollection is the media type for an array of Project (default view)
//
// Identifier: application/vnd.app.project+json; type=collection; view=default
type ProjectCollection []*Project

// Validate validates the ProjectCollection media type instance.
func (mt ProjectCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeProjectCollection decodes the ProjectCollection instance encoded in resp body.
func (c *Client) DecodeProjectCollection(resp *http.Response) (ProjectCollection, error) {
	var decoded ProjectCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// Projects media type (default view)
//
// Identifier: application/vnd.app.projects+json; view=default
type Projects struct {
	Projects ProjectCollection `form:"projects" json:"projects" xml:"projects"`
}

// Validate validates the Projects media type instance.
func (mt *Projects) Validate() (err error) {
	if mt.Projects == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "projects"))
	}
	if err2 := mt.Projects.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// DecodeProjects decodes the Projects instance encoded in resp body.
func (c *Client) DecodeProjects(resp *http.Response) (*Projects, error) {
	var decoded Projects
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Team media type (default view)
//
// Identifier: application/vnd.app.team+json; view=default
type Team struct {
	Description string `form:"description" json:"description" xml:"description"`
	ID          int    `form:"id" json:"id" xml:"id"`
	Name        string `form:"name" json:"name" xml:"name"`
	Slug        string `form:"slug" json:"slug" xml:"slug"`
}

// Validate validates the Team media type instance.
func (mt *Team) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Slug == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "slug"))
	}
	if mt.Description == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "description"))
	}
	if ok := goa.ValidatePattern(`^[[:alnum:]]+(-[[:alnum:]]+)*$`, mt.Slug); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.slug`, mt.Slug, `^[[:alnum:]]+(-[[:alnum:]]+)*$`))
	}
	if utf8.RuneCountInString(mt.Slug) > 40 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.slug`, mt.Slug, utf8.RuneCountInString(mt.Slug), 40, false))
	}
	return
}

// DecodeTeam decodes the Team instance encoded in resp body.
func (c *Client) DecodeTeam(resp *http.Response) (*Team, error) {
	var decoded Team
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// TeamCollection is the media type for an array of Team (default view)
//
// Identifier: application/vnd.app.team+json; type=collection; view=default
type TeamCollection []*Team

// Validate validates the TeamCollection media type instance.
func (mt TeamCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeTeamCollection decodes the TeamCollection instance encoded in resp body.
func (c *Client) DecodeTeamCollection(resp *http.Response) (TeamCollection, error) {
	var decoded TeamCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// Teams media type (default view)
//
// Identifier: application/vnd.app.teams+json; view=default
type Teams struct {
	Teams TeamCollection `form:"teams" json:"teams" xml:"teams"`
}

// Validate validates the Teams media type instance.
func (mt *Teams) Validate() (err error) {
	if mt.Teams == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "teams"))
	}
	if err2 := mt.Teams.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// DecodeTeams decodes the Teams instance encoded in resp body.
func (c *Client) DecodeTeams(resp *http.Response) (*Teams, error) {
	var decoded Teams
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// TwitterAccount media type (default view)
//
// Identifier: application/vnd.app.twitter_account+json; view=default
type TwitterAccount struct {
	ExpeditionID     int    `form:"expedition_id" json:"expedition_id" xml:"expedition_id"`
	ID               int    `form:"id" json:"id" xml:"id"`
	ScreenName       string `form:"screen_name" json:"screen_name" xml:"screen_name"`
	TwitterAccountID int    `form:"twitter_account_id" json:"twitter_account_id" xml:"twitter_account_id"`
}

// Validate validates the TwitterAccount media type instance.
func (mt *TwitterAccount) Validate() (err error) {

	if mt.ScreenName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "screen_name"))
	}
	return
}

// DecodeTwitterAccount decodes the TwitterAccount instance encoded in resp body.
func (c *Client) DecodeTwitterAccount(resp *http.Response) (*TwitterAccount, error) {
	var decoded TwitterAccount
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// TwitterAccountCollection is the media type for an array of TwitterAccount (default view)
//
// Identifier: application/vnd.app.twitter_account+json; type=collection; view=default
type TwitterAccountCollection []*TwitterAccount

// Validate validates the TwitterAccountCollection media type instance.
func (mt TwitterAccountCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeTwitterAccountCollection decodes the TwitterAccountCollection instance encoded in resp body.
func (c *Client) DecodeTwitterAccountCollection(resp *http.Response) (TwitterAccountCollection, error) {
	var decoded TwitterAccountCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// TwitterAccounts media type (default view)
//
// Identifier: application/vnd.app.twitter_accounts+json; view=default
type TwitterAccounts struct {
	TwitterAccounts TwitterAccountCollection `form:"twitter_accounts" json:"twitter_accounts" xml:"twitter_accounts"`
}

// Validate validates the TwitterAccounts media type instance.
func (mt *TwitterAccounts) Validate() (err error) {
	if mt.TwitterAccounts == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "twitter_accounts"))
	}
	if err2 := mt.TwitterAccounts.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// DecodeTwitterAccounts decodes the TwitterAccounts instance encoded in resp body.
func (c *Client) DecodeTwitterAccounts(resp *http.Response) (*TwitterAccounts, error) {
	var decoded TwitterAccounts
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// User media type (default view)
//
// Identifier: application/vnd.app.user+json; view=default
type User struct {
	Bio      string `form:"bio" json:"bio" xml:"bio"`
	Email    string `form:"email" json:"email" xml:"email"`
	ID       int    `form:"id" json:"id" xml:"id"`
	Name     string `form:"name" json:"name" xml:"name"`
	Username string `form:"username" json:"username" xml:"username"`
}

// Validate validates the User media type instance.
func (mt *User) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Username == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "username"))
	}
	if mt.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	if mt.Bio == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "bio"))
	}
	if err2 := goa.ValidateFormat(goa.FormatEmail, mt.Email); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, mt.Email, goa.FormatEmail, err2))
	}
	if ok := goa.ValidatePattern(`\S`, mt.Name); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.name`, mt.Name, `\S`))
	}
	if utf8.RuneCountInString(mt.Name) > 256 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, mt.Name, utf8.RuneCountInString(mt.Name), 256, false))
	}
	if ok := goa.ValidatePattern(`^[[:alnum:]]+(-[[:alnum:]]+)*$`, mt.Username); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.username`, mt.Username, `^[[:alnum:]]+(-[[:alnum:]]+)*$`))
	}
	if utf8.RuneCountInString(mt.Username) > 40 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.username`, mt.Username, utf8.RuneCountInString(mt.Username), 40, false))
	}
	return
}

// DecodeUser decodes the User instance encoded in resp body.
func (c *Client) DecodeUser(resp *http.Response) (*User, error) {
	var decoded User
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// UserCollection is the media type for an array of User (default view)
//
// Identifier: application/vnd.app.user+json; type=collection; view=default
type UserCollection []*User

// Validate validates the UserCollection media type instance.
func (mt UserCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeUserCollection decodes the UserCollection instance encoded in resp body.
func (c *Client) DecodeUserCollection(resp *http.Response) (UserCollection, error) {
	var decoded UserCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// Users media type (default view)
//
// Identifier: application/vnd.app.users+json; view=default
type Users struct {
	Users UserCollection `form:"users" json:"users" xml:"users"`
}

// Validate validates the Users media type instance.
func (mt *Users) Validate() (err error) {
	if mt.Users == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "users"))
	}
	if err2 := mt.Users.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// DecodeUsers decodes the Users instance encoded in resp body.
func (c *Client) DecodeUsers(resp *http.Response) (*Users, error) {
	var decoded Users
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeErrorResponse decodes the ErrorResponse instance encoded in resp body.
func (c *Client) DecodeErrorResponse(resp *http.Response) (*goa.ErrorResponse, error) {
	var decoded goa.ErrorResponse
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
