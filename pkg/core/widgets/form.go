package widgets

import (
	"sync"

	goflow "github.com/base-go/GoFlow/pkg/core/framework"
)

// FormState manages the state of a Form
type FormState struct {
	fields         []FormFieldState
	autoValidate   bool
	autoValidateMode AutoValidateMode
	mutex          sync.RWMutex
}

// AutoValidateMode defines when to auto-validate
type AutoValidateMode int

const (
	AutoValidateDisabled AutoValidateMode = iota
	AutoValidateAlways
	AutoValidateOnUserInteraction
)

// NewFormState creates a new form state
func NewFormState() *FormState {
	return &FormState{
		fields:         make([]FormFieldState, 0),
		autoValidate:   false,
		autoValidateMode: AutoValidateDisabled,
	}
}

// RegisterField registers a form field
func (fs *FormState) RegisterField(field FormFieldState) {
	fs.mutex.Lock()
	defer fs.mutex.Unlock()
	fs.fields = append(fs.fields, field)
}

// UnregisterField unregisters a form field
func (fs *FormState) UnregisterField(field FormFieldState) {
	fs.mutex.Lock()
	defer fs.mutex.Unlock()
	for i, f := range fs.fields {
		if f == field {
			fs.fields = append(fs.fields[:i], fs.fields[i+1:]...)
			return
		}
	}
}

// Validate validates all fields in the form
func (fs *FormState) Validate() bool {
	fs.mutex.RLock()
	defer fs.mutex.RUnlock()

	valid := true
	for _, field := range fs.fields {
		if !field.Validate() {
			valid = false
		}
	}
	return valid
}

// Save saves all fields in the form
func (fs *FormState) Save() {
	fs.mutex.RLock()
	defer fs.mutex.RUnlock()

	for _, field := range fs.fields {
		field.Save()
	}
}

// Reset resets all fields in the form
func (fs *FormState) Reset() {
	fs.mutex.RLock()
	defer fs.mutex.RUnlock()

	for _, field := range fs.fields {
		field.Reset()
	}
}

// SetAutoValidateMode sets the auto-validate mode
func (fs *FormState) SetAutoValidateMode(mode AutoValidateMode) {
	fs.mutex.Lock()
	defer fs.mutex.Unlock()
	fs.autoValidateMode = mode
	fs.autoValidate = mode != AutoValidateDisabled
}

// FormFieldState defines the interface for form field state
type FormFieldState interface {
	Validate() bool
	Save()
	Reset()
	GetValue() interface{}
	SetValue(value interface{})
}

// Form widget for managing form validation
type Form struct {
	goflow.BaseWidget
	Key              *FormKey
	Child            goflow.Widget
	AutoValidateMode AutoValidateMode
	OnChanged        func()
}

// FormKey provides access to the form state
type FormKey struct {
	state *FormState
}

// NewFormKey creates a new form key
func NewFormKey() *FormKey {
	return &FormKey{
		state: NewFormState(),
	}
}

// CurrentState returns the current form state
func (fk *FormKey) CurrentState() *FormState {
	return fk.state
}

// NewForm creates a new Form widget
func NewForm(key *FormKey, child goflow.Widget) *Form {
	if key == nil {
		key = NewFormKey()
	}
	return &Form{
		Key:              key,
		Child:            child,
		AutoValidateMode: AutoValidateDisabled,
	}
}

// Build creates the widget tree
func (f *Form) Build(context goflow.BuildContext) goflow.Widget {
	// Set auto-validate mode
	f.Key.state.SetAutoValidateMode(f.AutoValidateMode)

	// Wrap in a FormScope to provide form state to children
	return &FormScope{
		FormState: f.Key.state,
		Child:     f.Child,
	}
}

// FormScope provides form state to descendants
type FormScope struct {
	goflow.BaseWidget
	FormState *FormState
	Child     goflow.Widget
}

// Build creates the widget tree
func (fs *FormScope) Build(context goflow.BuildContext) goflow.Widget {
	return fs.Child
}

// FormField is a base widget for form fields
type FormField struct {
	goflow.BaseWidget
	Validator    func(value string) *string // Returns error message or nil
	OnSaved      func(value string)
	InitialValue string
	AutoValidate bool
	Enabled      bool
}

// FormFieldStateImpl implements FormFieldState
type FormFieldStateImpl struct {
	value        string
	errorText    *string
	validator    func(value string) *string
	onSaved      func(value string)
	hasInteracted bool
	mutex        sync.RWMutex
}

// NewFormFieldState creates a new form field state
func NewFormFieldState(initialValue string, validator func(value string) *string, onSaved func(value string)) *FormFieldStateImpl {
	return &FormFieldStateImpl{
		value:        initialValue,
		errorText:    nil,
		validator:    validator,
		onSaved:      onSaved,
		hasInteracted: false,
	}
}

// Validate validates the field
func (ffs *FormFieldStateImpl) Validate() bool {
	ffs.mutex.Lock()
	defer ffs.mutex.Unlock()

	if ffs.validator != nil {
		ffs.errorText = ffs.validator(ffs.value)
		return ffs.errorText == nil
	}
	ffs.errorText = nil
	return true
}

// Save saves the field value
func (ffs *FormFieldStateImpl) Save() {
	ffs.mutex.RLock()
	defer ffs.mutex.RUnlock()

	if ffs.onSaved != nil {
		ffs.onSaved(ffs.value)
	}
}

// Reset resets the field
func (ffs *FormFieldStateImpl) Reset() {
	ffs.mutex.Lock()
	defer ffs.mutex.Unlock()

	ffs.value = ""
	ffs.errorText = nil
	ffs.hasInteracted = false
}

// GetValue returns the current value
func (ffs *FormFieldStateImpl) GetValue() interface{} {
	ffs.mutex.RLock()
	defer ffs.mutex.RUnlock()
	return ffs.value
}

// SetValue sets the current value
func (ffs *FormFieldStateImpl) SetValue(value interface{}) {
	ffs.mutex.Lock()
	defer ffs.mutex.Unlock()

	if strValue, ok := value.(string); ok {
		ffs.value = strValue
		ffs.hasInteracted = true
	}
}

// GetErrorText returns the error text
func (ffs *FormFieldStateImpl) GetErrorText() *string {
	ffs.mutex.RLock()
	defer ffs.mutex.RUnlock()
	return ffs.errorText
}

// SetErrorText sets the error text
func (ffs *FormFieldStateImpl) SetErrorText(errorText *string) {
	ffs.mutex.Lock()
	defer ffs.mutex.Unlock()
	ffs.errorText = errorText
}

// HasInteracted returns whether the user has interacted with the field
func (ffs *FormFieldStateImpl) HasInteracted() bool {
	ffs.mutex.RLock()
	defer ffs.mutex.RUnlock()
	return ffs.hasInteracted
}

// Common validators
var (
	// RequiredValidator validates that a field is not empty
	RequiredValidator = func(errorMessage string) func(string) *string {
		return func(value string) *string {
			if value == "" {
				msg := errorMessage
				if msg == "" {
					msg = "This field is required"
				}
				return &msg
			}
			return nil
		}
	}

	// EmailValidator validates email format
	EmailValidator = func(errorMessage string) func(string) *string {
		return func(value string) *string {
			if value == "" {
				return nil
			}
			// Simple email validation (can be enhanced)
			hasAt := false
			hasDot := false
			for i, c := range value {
				if c == '@' {
					if hasAt || i == 0 || i == len(value)-1 {
						msg := errorMessage
						if msg == "" {
							msg = "Invalid email format"
						}
						return &msg
					}
					hasAt = true
				}
				if c == '.' && hasAt {
					hasDot = true
				}
			}
			if !hasAt || !hasDot {
				msg := errorMessage
				if msg == "" {
					msg = "Invalid email format"
				}
				return &msg
			}
			return nil
		}
	}

	// MinLengthValidator validates minimum length
	MinLengthValidator = func(minLength int, errorMessage string) func(string) *string {
		return func(value string) *string {
			if len(value) < minLength {
				msg := errorMessage
				if msg == "" {
					msg = "Minimum length is " + string(rune(minLength+'0'))
				}
				return &msg
			}
			return nil
		}
	}

	// MaxLengthValidator validates maximum length
	MaxLengthValidator = func(maxLength int, errorMessage string) func(string) *string {
		return func(value string) *string {
			if len(value) > maxLength {
				msg := errorMessage
				if msg == "" {
					msg = "Maximum length is " + string(rune(maxLength+'0'))
				}
				return &msg
			}
			return nil
		}
	}
)

// ComposeValidators combines multiple validators
func ComposeValidators(validators ...func(string) *string) func(string) *string {
	return func(value string) *string {
		for _, validator := range validators {
			if err := validator(value); err != nil {
				return err
			}
		}
		return nil
	}
}
