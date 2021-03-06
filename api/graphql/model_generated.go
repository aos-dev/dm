// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type CreateIdentity struct {
	Name       string           `json:"name"`
	Type       IdentityType     `json:"type"`
	Credential *CredentialInput `json:"credential"`
	Endpoint   *EndpointInput   `json:"endpoint"`
}

type CreateTask struct {
	Name     string          `json:"name"`
	Type     TaskType        `json:"type"`
	Storages []*StorageInput `json:"storages"`
	Options  []*PairInput    `json:"options"`
	Staffs   []*StaffInput   `json:"staffs"`
}

type Credential struct {
	Protocol string   `json:"protocol"`
	Args     []string `json:"args"`
}

type CredentialInput struct {
	Protocol string   `json:"protocol"`
	Args     []string `json:"args"`
}

type DeleteIdentity struct {
	Name string       `json:"name"`
	Type IdentityType `json:"type"`
}

type DeleteTask struct {
	ID string `json:"id"`
}

type Endpoint struct {
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

type EndpointInput struct {
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

type Identity struct {
	Name       string       `json:"name"`
	Type       IdentityType `json:"type"`
	Credential *Credential  `json:"credential"`
	Endpoint   *Endpoint    `json:"endpoint"`
}

type Pair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type PairInput struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Staff struct {
	ID string `json:"id"`
}

type StaffInput struct {
	ID string `json:"id"`
}

type Storage struct {
	Type    StorageType `json:"type"`
	Options []*Pair     `json:"options"`
}

type StorageInput struct {
	Type    StorageType  `json:"type"`
	Options []*PairInput `json:"options"`
}

type Task struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Type      TaskType   `json:"type"`
	Status    TaskStatus `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Storages  []*Storage `json:"storages"`
	Options   []*Pair    `json:"options"`
	Staffs    []*Staff   `json:"staffs"`
}

type IdentityType string

const (
	IdentityTypeQingstor IdentityType = "Qingstor"
)

var AllIdentityType = []IdentityType{
	IdentityTypeQingstor,
}

func (e IdentityType) IsValid() bool {
	switch e {
	case IdentityTypeQingstor:
		return true
	}
	return false
}

func (e IdentityType) String() string {
	return string(e)
}

func (e *IdentityType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = IdentityType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid IdentityType", str)
	}
	return nil
}

func (e IdentityType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type StorageType string

const (
	StorageTypeFs       StorageType = "Fs"
	StorageTypeQingstor StorageType = "Qingstor"
)

var AllStorageType = []StorageType{
	StorageTypeFs,
	StorageTypeQingstor,
}

func (e StorageType) IsValid() bool {
	switch e {
	case StorageTypeFs, StorageTypeQingstor:
		return true
	}
	return false
}

func (e StorageType) String() string {
	return string(e)
}

func (e *StorageType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = StorageType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid StorageType", str)
	}
	return nil
}

func (e StorageType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TaskStatus string

const (
	TaskStatusCreated  TaskStatus = "Created"
	TaskStatusReady    TaskStatus = "Ready"
	TaskStatusRunning  TaskStatus = "Running"
	TaskStatusFinished TaskStatus = "Finished"
	TaskStatusStopped  TaskStatus = "Stopped"
	TaskStatusError    TaskStatus = "Error"
)

var AllTaskStatus = []TaskStatus{
	TaskStatusCreated,
	TaskStatusReady,
	TaskStatusRunning,
	TaskStatusFinished,
	TaskStatusStopped,
	TaskStatusError,
}

func (e TaskStatus) IsValid() bool {
	switch e {
	case TaskStatusCreated, TaskStatusReady, TaskStatusRunning, TaskStatusFinished, TaskStatusStopped, TaskStatusError:
		return true
	}
	return false
}

func (e TaskStatus) String() string {
	return string(e)
}

func (e *TaskStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TaskStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TaskStatus", str)
	}
	return nil
}

func (e TaskStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TaskType string

const (
	TaskTypeCopyDir TaskType = "copyDir"
)

var AllTaskType = []TaskType{
	TaskTypeCopyDir,
}

func (e TaskType) IsValid() bool {
	switch e {
	case TaskTypeCopyDir:
		return true
	}
	return false
}

func (e TaskType) String() string {
	return string(e)
}

func (e *TaskType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TaskType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TaskType", str)
	}
	return nil
}

func (e TaskType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
