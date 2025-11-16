package api_test

import (
	"errors"
	"testing"

	"github.com/base-go/GoFlow/pkg/api"
)

func TestResourceInitialState(t *testing.T) {
	resource := api.NewResource[string]()
	state := resource.Get()

	if state.Loading {
		t.Error("Expected Loading to be false initially")
	}
	if state.Error != nil {
		t.Error("Expected Error to be nil initially")
	}
	if state.Data != nil {
		t.Error("Expected Data to be nil initially")
	}
}

func TestResourceLoadingState(t *testing.T) {
	resource := api.NewLoadingResource[string]()
	state := resource.Get()

	if !state.Loading {
		t.Error("Expected Loading to be true")
	}
	if state.Error != nil {
		t.Error("Expected Error to be nil")
	}
	if state.Data != nil {
		t.Error("Expected Data to be nil")
	}
}

func TestResourceSetData(t *testing.T) {
	resource := api.NewResource[string]()
	resource.SetData("test data")

	state := resource.Get()

	if state.Loading {
		t.Error("Expected Loading to be false after SetData")
	}
	if state.Error != nil {
		t.Error("Expected Error to be nil after SetData")
	}
	if state.Data == nil {
		t.Fatal("Expected Data to be non-nil after SetData")
	}
	if *state.Data != "test data" {
		t.Errorf("Expected Data to be 'test data', got '%s'", *state.Data)
	}
}

func TestResourceSetError(t *testing.T) {
	resource := api.NewResource[string]()
	testError := errors.New("test error")
	resource.SetError(testError)

	state := resource.Get()

	if state.Loading {
		t.Error("Expected Loading to be false after SetError")
	}
	if state.Error == nil {
		t.Fatal("Expected Error to be non-nil after SetError")
	}
	if state.Error.Error() != "test error" {
		t.Errorf("Expected Error to be 'test error', got '%s'", state.Error.Error())
	}
}

func TestResourceSetLoading(t *testing.T) {
	resource := api.NewResource[string]()
	resource.SetData("existing data")
	resource.SetLoading()

	state := resource.Get()

	if !state.Loading {
		t.Error("Expected Loading to be true after SetLoading")
	}
	if state.Error != nil {
		t.Error("Expected Error to be nil after SetLoading")
	}
	// Data should be preserved
	if state.Data == nil {
		t.Fatal("Expected Data to be preserved after SetLoading")
	}
	if *state.Data != "existing data" {
		t.Errorf("Expected Data to be 'existing data', got '%s'", *state.Data)
	}
}

func TestResourceUpdate(t *testing.T) {
	resource := api.NewResource[string]()
	data := "updated data"
	resource.Update(&data, nil)

	state := resource.Get()

	if state.Loading {
		t.Error("Expected Loading to be false after successful Update")
	}
	if state.Error != nil {
		t.Error("Expected Error to be nil after successful Update")
	}
	if state.Data == nil {
		t.Fatal("Expected Data to be non-nil after successful Update")
	}
	if *state.Data != "updated data" {
		t.Errorf("Expected Data to be 'updated data', got '%s'", *state.Data)
	}
}

func TestResourceUpdateWithError(t *testing.T) {
	resource := api.NewResource[string]()
	testError := errors.New("update error")
	resource.Update(nil, testError)

	state := resource.Get()

	if state.Loading {
		t.Error("Expected Loading to be false after error Update")
	}
	if state.Error == nil {
		t.Fatal("Expected Error to be non-nil after error Update")
	}
	if state.Error.Error() != "update error" {
		t.Errorf("Expected Error to be 'update error', got '%s'", state.Error.Error())
	}
}

func TestResourceHelperMethods(t *testing.T) {
	resource := api.NewResource[string]()

	if resource.IsLoading() {
		t.Error("Expected IsLoading to be false initially")
	}
	if resource.HasError() {
		t.Error("Expected HasError to be false initially")
	}
	if resource.GetData() != nil {
		t.Error("Expected GetData to be nil initially")
	}
	if resource.GetError() != nil {
		t.Error("Expected GetError to be nil initially")
	}

	resource.SetLoading()
	if !resource.IsLoading() {
		t.Error("Expected IsLoading to be true after SetLoading")
	}

	testError := errors.New("test error")
	resource.SetError(testError)
	if !resource.HasError() {
		t.Error("Expected HasError to be true after SetError")
	}
	if resource.GetError() != testError {
		t.Error("Expected GetError to return the test error")
	}

	resource.SetData("test")
	if resource.GetData() == nil {
		t.Fatal("Expected GetData to be non-nil after SetData")
	}
	if *resource.GetData() != "test" {
		t.Errorf("Expected GetData to return 'test', got '%s'", *resource.GetData())
	}
}

func TestResourceListSetData(t *testing.T) {
	resource := api.NewResourceList[string]()
	data := []string{"a", "b", "c"}
	resource.SetData(data)

	state := resource.Get()

	if state.Loading {
		t.Error("Expected Loading to be false after SetData")
	}
	if state.Error != nil {
		t.Error("Expected Error to be nil after SetData")
	}
	if state.Data == nil {
		t.Fatal("Expected Data to be non-nil after SetData")
	}

	result := *state.Data
	if len(result) != 3 {
		t.Errorf("Expected Data length to be 3, got %d", len(result))
	}
	if result[0] != "a" || result[1] != "b" || result[2] != "c" {
		t.Errorf("Expected Data to be [a, b, c], got %v", result)
	}
}

func TestResourceListGetData(t *testing.T) {
	resource := api.NewResourceList[int]()

	// Initially nil
	if resource.GetData() != nil {
		t.Error("Expected GetData to return nil initially")
	}

	// After setting data
	data := []int{1, 2, 3}
	resource.SetData(data)

	result := resource.GetData()
	if result == nil {
		t.Fatal("Expected GetData to return non-nil after SetData")
	}
	if len(result) != 3 {
		t.Errorf("Expected length 3, got %d", len(result))
	}
	if result[0] != 1 || result[1] != 2 || result[2] != 3 {
		t.Errorf("Expected [1, 2, 3], got %v", result)
	}
}
