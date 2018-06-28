// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PublicAPISessionModelsSessionModel public Api session models session model
// swagger:model PublicApi.Session.Models.SessionModel
type PublicAPISessionModelsSessionModel struct {

	// Trail of all session activities.
	Activities []*PublicAPISessionModelsSessionModelActivity `json:"activities"`

	// Guest A/V capabilities
	// Enum: [WEB_RTC FLASH_ONLY NONE]
	AvSupport string `json:"avSupport,omitempty"`

	// Whether the guest was blacklisted.
	Blacklisted bool `json:"blacklisted,omitempty"`

	// The channel at the time of session creation
	// Enum: [Chat AV]
	Channel string `json:"channel,omitempty"`

	// company Id
	CompanyID int32 `json:"companyId,omitempty"`

	// department Id
	DepartmentID int32 `json:"departmentId,omitempty"`

	// When guest enters via invitation, this is the generated number.
	DirectCallNumber string `json:"directCallNumber,omitempty"`

	// When operator or guest leaves the session.
	// Format: date-time
	Ended strfmt.DateTime `json:"ended,omitempty"`

	// Session entered queue.
	// Format: date-time
	EnteredQueue strfmt.DateTime `json:"enteredQueue,omitempty"`

	// Guest information
	Guest *PublicAPISessionModelsSessionModelGuestInfo `json:"guest,omitempty"`

	// Contents of feedback form custom fields.
	GuestFeedbackCustomFields map[string]interface{} `json:"guestFeedbackCustomFields,omitempty"`

	// Number of stars given by guest on feedback form.
	// Maximum: 5
	// Minimum: 1
	GuestFeedbackStars int32 `json:"guestFeedbackStars,omitempty"`

	// Freeform textual feedback giben by guest on feedback form.
	GuestFeedbackText string `json:"guestFeedbackText,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// The session was proactively created by operator and offered to a visiting guest.
	IsDirectChat bool `json:"isDirectChat,omitempty"`

	// ISO 8601 duration. Length of session or null if session never started.
	Length string `json:"length,omitempty"`

	// Operators who participated in the session
	Operators []*PublicAPISessionModelsSessionModelOperatorInfo `json:"operators"`

	// For callback session only.
	// Time when operator dialed callback number.
	// Format: date-time
	PhoneCallbackDialed strfmt.DateTime `json:"phoneCallbackDialed,omitempty"`

	// When guest enters via phone callback, this is the phone number.
	PhoneCallbackNumber string `json:"phoneCallbackNumber,omitempty"`

	// Reviewers
	Reviewers []*PublicAPISessionModelsSessionModelReviewerInfo `json:"reviewers"`

	// The origin of the session creation
	// Enum: [Default Callback Invitation ChatInvitation API Facebook]
	Source string `json:"source,omitempty"`

	// When first operator accepts the session.
	// Format: date-time
	Started strfmt.DateTime `json:"started,omitempty"`

	// Tags associated with this session.
	Tags []*PublicAPISessionModelsSessionModelAssignedTag `json:"tags"`

	// Entry URL at the time of session creation
	URL string `json:"url,omitempty"`

	// ISO 8601 duration. From entering queue to start of session or leaving the queue.
	// Whichever comes first.
	Waited string `json:"waited,omitempty"`

	// Name of entry widget
	Widget string `json:"widget,omitempty"`
}

// Validate validates this public Api session models session model
func (m *PublicAPISessionModelsSessionModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAvSupport(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChannel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnded(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnteredQueue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGuest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGuestFeedbackStars(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperators(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhoneCallbackDialed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReviewers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStarted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicAPISessionModelsSessionModel) validateActivities(formats strfmt.Registry) error {

	if swag.IsZero(m.Activities) { // not required
		return nil
	}

	for i := 0; i < len(m.Activities); i++ {
		if swag.IsZero(m.Activities[i]) { // not required
			continue
		}

		if m.Activities[i] != nil {
			if err := m.Activities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("activities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var publicApiSessionModelsSessionModelTypeAvSupportPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["WEB_RTC","FLASH_ONLY","NONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		publicApiSessionModelsSessionModelTypeAvSupportPropEnum = append(publicApiSessionModelsSessionModelTypeAvSupportPropEnum, v)
	}
}

const (

	// PublicAPISessionModelsSessionModelAvSupportWEBRTC captures enum value "WEB_RTC"
	PublicAPISessionModelsSessionModelAvSupportWEBRTC string = "WEB_RTC"

	// PublicAPISessionModelsSessionModelAvSupportFLASHONLY captures enum value "FLASH_ONLY"
	PublicAPISessionModelsSessionModelAvSupportFLASHONLY string = "FLASH_ONLY"

	// PublicAPISessionModelsSessionModelAvSupportNONE captures enum value "NONE"
	PublicAPISessionModelsSessionModelAvSupportNONE string = "NONE"
)

// prop value enum
func (m *PublicAPISessionModelsSessionModel) validateAvSupportEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, publicApiSessionModelsSessionModelTypeAvSupportPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PublicAPISessionModelsSessionModel) validateAvSupport(formats strfmt.Registry) error {

	if swag.IsZero(m.AvSupport) { // not required
		return nil
	}

	// value enum
	if err := m.validateAvSupportEnum("avSupport", "body", m.AvSupport); err != nil {
		return err
	}

	return nil
}

var publicApiSessionModelsSessionModelTypeChannelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Chat","AV"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		publicApiSessionModelsSessionModelTypeChannelPropEnum = append(publicApiSessionModelsSessionModelTypeChannelPropEnum, v)
	}
}

const (

	// PublicAPISessionModelsSessionModelChannelChat captures enum value "Chat"
	PublicAPISessionModelsSessionModelChannelChat string = "Chat"

	// PublicAPISessionModelsSessionModelChannelAV captures enum value "AV"
	PublicAPISessionModelsSessionModelChannelAV string = "AV"
)

// prop value enum
func (m *PublicAPISessionModelsSessionModel) validateChannelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, publicApiSessionModelsSessionModelTypeChannelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PublicAPISessionModelsSessionModel) validateChannel(formats strfmt.Registry) error {

	if swag.IsZero(m.Channel) { // not required
		return nil
	}

	// value enum
	if err := m.validateChannelEnum("channel", "body", m.Channel); err != nil {
		return err
	}

	return nil
}

func (m *PublicAPISessionModelsSessionModel) validateEnded(formats strfmt.Registry) error {

	if swag.IsZero(m.Ended) { // not required
		return nil
	}

	if err := validate.FormatOf("ended", "body", "date-time", m.Ended.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PublicAPISessionModelsSessionModel) validateEnteredQueue(formats strfmt.Registry) error {

	if swag.IsZero(m.EnteredQueue) { // not required
		return nil
	}

	if err := validate.FormatOf("enteredQueue", "body", "date-time", m.EnteredQueue.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PublicAPISessionModelsSessionModel) validateGuest(formats strfmt.Registry) error {

	if swag.IsZero(m.Guest) { // not required
		return nil
	}

	if m.Guest != nil {
		if err := m.Guest.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("guest")
			}
			return err
		}
	}

	return nil
}

func (m *PublicAPISessionModelsSessionModel) validateGuestFeedbackStars(formats strfmt.Registry) error {

	if swag.IsZero(m.GuestFeedbackStars) { // not required
		return nil
	}

	if err := validate.MinimumInt("guestFeedbackStars", "body", int64(m.GuestFeedbackStars), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("guestFeedbackStars", "body", int64(m.GuestFeedbackStars), 5, false); err != nil {
		return err
	}

	return nil
}

func (m *PublicAPISessionModelsSessionModel) validateOperators(formats strfmt.Registry) error {

	if swag.IsZero(m.Operators) { // not required
		return nil
	}

	for i := 0; i < len(m.Operators); i++ {
		if swag.IsZero(m.Operators[i]) { // not required
			continue
		}

		if m.Operators[i] != nil {
			if err := m.Operators[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("operators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PublicAPISessionModelsSessionModel) validatePhoneCallbackDialed(formats strfmt.Registry) error {

	if swag.IsZero(m.PhoneCallbackDialed) { // not required
		return nil
	}

	if err := validate.FormatOf("phoneCallbackDialed", "body", "date-time", m.PhoneCallbackDialed.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PublicAPISessionModelsSessionModel) validateReviewers(formats strfmt.Registry) error {

	if swag.IsZero(m.Reviewers) { // not required
		return nil
	}

	for i := 0; i < len(m.Reviewers); i++ {
		if swag.IsZero(m.Reviewers[i]) { // not required
			continue
		}

		if m.Reviewers[i] != nil {
			if err := m.Reviewers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("reviewers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var publicApiSessionModelsSessionModelTypeSourcePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Default","Callback","Invitation","ChatInvitation","API","Facebook"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		publicApiSessionModelsSessionModelTypeSourcePropEnum = append(publicApiSessionModelsSessionModelTypeSourcePropEnum, v)
	}
}

const (

	// PublicAPISessionModelsSessionModelSourceDefault captures enum value "Default"
	PublicAPISessionModelsSessionModelSourceDefault string = "Default"

	// PublicAPISessionModelsSessionModelSourceCallback captures enum value "Callback"
	PublicAPISessionModelsSessionModelSourceCallback string = "Callback"

	// PublicAPISessionModelsSessionModelSourceInvitation captures enum value "Invitation"
	PublicAPISessionModelsSessionModelSourceInvitation string = "Invitation"

	// PublicAPISessionModelsSessionModelSourceChatInvitation captures enum value "ChatInvitation"
	PublicAPISessionModelsSessionModelSourceChatInvitation string = "ChatInvitation"

	// PublicAPISessionModelsSessionModelSourceAPI captures enum value "API"
	PublicAPISessionModelsSessionModelSourceAPI string = "API"

	// PublicAPISessionModelsSessionModelSourceFacebook captures enum value "Facebook"
	PublicAPISessionModelsSessionModelSourceFacebook string = "Facebook"
)

// prop value enum
func (m *PublicAPISessionModelsSessionModel) validateSourceEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, publicApiSessionModelsSessionModelTypeSourcePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PublicAPISessionModelsSessionModel) validateSource(formats strfmt.Registry) error {

	if swag.IsZero(m.Source) { // not required
		return nil
	}

	// value enum
	if err := m.validateSourceEnum("source", "body", m.Source); err != nil {
		return err
	}

	return nil
}

func (m *PublicAPISessionModelsSessionModel) validateStarted(formats strfmt.Registry) error {

	if swag.IsZero(m.Started) { // not required
		return nil
	}

	if err := validate.FormatOf("started", "body", "date-time", m.Started.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PublicAPISessionModelsSessionModel) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PublicAPISessionModelsSessionModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicAPISessionModelsSessionModel) UnmarshalBinary(b []byte) error {
	var res PublicAPISessionModelsSessionModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
