// +build !windows

// Package external provides APIs to operate Menoh model directory.
// API design follows menoh.h interface.
package external

/*
#cgo LDFLAGS: -lmenoh

#include <stdlib.h>
#include <menoh/menoh.h>
*/
import "C"
import (
	"errors"
	"unsafe"
)

// ModelData bind. Required to delete after making, call Delete function.
type ModelData struct {
	h C.menoh_model_data_handle
}

// Delete object.
func (m *ModelData) Delete() {
	C.menoh_delete_model_data(m.h)
	m.h = nil
}

// MakeModelDataFromONNX returns ModelData using ONNX file path.
func MakeModelDataFromONNX(path string) (*ModelData, error) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	var h C.menoh_model_data_handle
	if err := checkError(C.menoh_make_model_data_from_onnx(cPath, &h)); err != nil {
		return nil, err
	}
	return &ModelData{h: h}, nil
}

// MakeModelDataFromONNXBytes return ModelData with ONNX file byte data.
func MakeModelDataFromONNXBytes(data []byte) (*ModelData, error) {
	var h C.menoh_model_data_handle
	if err := checkError(C.menoh_make_model_data_from_onnx_data_on_memory(
			(*C.uchar)(unsafe.Pointer(&data[0])), C.int(len(data)), &h)); err != nil {
		return nil, err
	}
	return &ModelData{h: h}, nil
}

// MakeModelData returns empty ModelData object, to build manually.
func MakeModelData() (*ModelData, error) {
	var h C.menoh_model_data_handle
	if err := checkError(C.menoh_make_model_data(&h)); err != nil {
		return nil, err
	}
	return &ModelData{h: h}, nil
}

// AddParameter adds named parameter.
func (m *ModelData) AddParameter(name string, param Variable) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return checkError(C.menoh_model_data_add_parameter(
		m.h, cName, C.int(param.Dtype), C.int(len(param.Dims)),
		(*C.int)(unsafe.Pointer(&param.Dims[0])), param.BufferHandle))
}

// AddNewNode adds new opType.
func (m *ModelData) AddNewNode(opType string) error {
	cName := C.CString(opType)
	defer C.free(unsafe.Pointer(cName))
	return checkError(C.menoh_model_data_add_new_node(m.h, cName))
}

// AddInputNameToCurrentNode adds input name to current node.
func (m *ModelData) AddInputNameToCurrentNode(inputName string) error {
	cName := C.CString(inputName)
	defer C.free(unsafe.Pointer(cName))
	return checkError(C.menoh_model_data_add_input_name_to_current_node(m.h, cName))
}

// AddOutputNameToCurrentNode adds output name to current node.
func (m *ModelData) AddOutputNameToCurrentNode(outputName string) error {
	cName := C.CString(outputName)
	defer C.free(unsafe.Pointer(cName))
	return checkError(C.menoh_model_data_add_output_name_to_current_node(m.h, cName))
}

// AddAttributeIntToCurrentNode adds integer type attribute to current node.
func (m *ModelData) AddAttributeIntToCurrentNode(attributeName string, value int) error {
	cName := C.CString(attributeName)
	defer C.free(unsafe.Pointer(cName))
	return checkError(C.menoh_model_data_add_attribute_int_to_current_node(m.h, cName, C.int(value)))
}

// AddAttributeFloatToCurrentNode adds float type attribute to current node.
func (m *ModelData) AddAttributeFloatToCurrentNode(attributeName string, value float32) error {
	cName := C.CString(attributeName)
	defer C.free(unsafe.Pointer(cName))
	return checkError(C.menoh_model_data_add_attribute_float_to_current_node(m.h, cName, C.float(value)))
}

// AddAttributeIntsToCurrentNode adds int array type attribute to current node.
func (m *ModelData) AddAttributeIntsToCurrentNode(attributeName string, values []int) error {
	cName := C.CString(attributeName)
	defer C.free(unsafe.Pointer(cName))
	return checkError(C.menoh_model_data_add_attribute_ints_to_current_node(
		m.h, cName, C.int(len(values)), (*C.int)(unsafe.Pointer(&values[0]))))
}

// AddAttributeFloatsToCurrentNode adds float array type attribute to current node.
func (m *ModelData) AddAttributeFloatsToCurrentNode(attributeName string, values []float32) error {
	cName := C.CString(attributeName)
	defer C.free(unsafe.Pointer(cName))
	return checkError(C.menoh_model_data_add_attribute_floats_to_current_node(
		m.h, cName, C.int(len(values)), (*C.float)(unsafe.Pointer(&values[0]))))
}

// Optimize ModelData with profiling table.
func (m *ModelData) Optimize(table VariableProfileTable) error {
	return checkError(C.menoh_model_data_optimize(m.h, table.h))
}

// VariableProfileTableBuilder bind. Required to delete after making, call Delete function.
type VariableProfileTableBuilder struct {
	h C.menoh_variable_profile_table_builder_handle
}

// Delete object.
func (b *VariableProfileTableBuilder) Delete() {
	C.menoh_delete_variable_profile_table_builder(b.h)
	b.h = nil
}

// MakeVariableProfileTableBuilder returns VariableProfileTableBuilder.
func MakeVariableProfileTableBuilder() (*VariableProfileTableBuilder, error) {
	var h C.menoh_variable_profile_table_builder_handle
	if err := checkError(C.menoh_make_variable_profile_table_builder(&h)); err != nil {
		return nil, err
	}
	return &VariableProfileTableBuilder{h: h}, nil
}

// AddInputProfile adds input profile with layer name, data type and dimension size.
func (b *VariableProfileTableBuilder) AddInputProfile(name string, dtype TypeMenohDtype, dims ...int32) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return checkError(
		C.menoh_variable_profile_table_builder_add_input_profile(
			b.h, cName, C.int(dtype), C.int(len(dims)), (*C.int)(unsafe.Pointer(&dims[0]))))
}

// AddOutputProfile adds output profile with layer name and data type.
func (b *VariableProfileTableBuilder) AddOutputProfile(name string, dtype TypeMenohDtype) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return checkError(
		C.menoh_variable_profile_table_builder_add_output_name(b.h, cName))
}

// BuildVariableProfileTable returns VariableProfileTable.
func (b *VariableProfileTableBuilder) BuildVariableProfileTable(md ModelData) (
	*VariableProfileTable, error) {

	var h C.menoh_variable_profile_table_handle
	if err := checkError(C.menoh_build_variable_profile_table(b.h, md.h, &h)); err != nil {
		return nil, err
	}
	return &VariableProfileTable{h: h}, nil
}

// VariableProfile represents profile information to make real variable.
type VariableProfile struct {
	Dtype TypeMenohDtype
	Dims  []int32
}

// VariableProfileTable bind. Required to delete after making, call Delete function.
type VariableProfileTable struct {
	h C.menoh_variable_profile_table_handle
}

// Delete object.
func (t *VariableProfileTable) Delete() {
	C.menoh_delete_variable_profile_table(t.h)
	t.h = nil
}

// GetVariableProfile returns profile which setup variable information includes.
func (t *VariableProfileTable) GetVariableProfile(name string) (*VariableProfile, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	var dtype C.int
	if err := checkError(C.menoh_variable_profile_table_get_dtype(t.h, cName, &dtype)); err != nil {
		return nil, err
	}
	var size C.int
	if err := checkError(C.menoh_variable_profile_table_get_dims_size(t.h, cName, &size)); err != nil {
		return nil, err
	}
	dims := make([]int32, size)
	for i := 0; i < int(size); i++ {
		var dim C.int
		if err := checkError(C.menoh_variable_profile_table_get_dims_at(t.h, cName, C.int(i), &dim)); err != nil {
			return nil, err
		}
		dims[i] = int32(dim)
	}
	return &VariableProfile{
		Dtype: toDtype(dtype),
		Dims:  dims,
	}, nil
}

// ModelBuilder bind. Required to delete after making, call Delete function.
type ModelBuilder struct {
	h C.menoh_model_builder_handle
}

// Delete object.
func (b *ModelBuilder) Delete() {
	C.menoh_delete_model_builder(b.h)
	b.h = nil
}

// MakeModelBuilder returns ModelBuilder.
func MakeModelBuilder(vpt VariableProfileTable) (*ModelBuilder, error) {
	var h C.menoh_model_builder_handle
	if err := checkError(C.menoh_make_model_builder(vpt.h, &h)); err != nil {
		return nil, err
	}
	return &ModelBuilder{h: h}, nil
}

// AttachExternalBuffer attaches data buffer to get data. This process must be done
// before building Model.
func (b *ModelBuilder) AttachExternalBuffer(variableName string, bufferPtr unsafe.Pointer) error {
	cVariableName := C.CString(variableName)
	defer C.free(unsafe.Pointer(cVariableName))
	return checkError(
		C.menoh_model_builder_attach_external_buffer(b.h, cVariableName, bufferPtr))
}

// BuildModel returns Model.
func (b *ModelBuilder) BuildModel(md ModelData, backend, backendConfig string) (*Model, error) {
	cBackend := C.CString(backend)
	defer C.free(unsafe.Pointer(cBackend))
	cBackendConfig := C.CString(backendConfig)
	defer C.free(unsafe.Pointer(cBackendConfig))
	var h C.menoh_model_handle
	if err := checkError(C.menoh_build_model(b.h, md.h, cBackend, cBackendConfig, &h)); err != nil {
		return nil, err
	}
	return &Model{h: h}, nil
}

// Variable represents data include data attribution and pointer.
type Variable struct {
	Dtype        TypeMenohDtype
	Dims         []int32
	BufferHandle unsafe.Pointer
}

// Model bind. Required to delete after making, call Delete function.
type Model struct {
	h C.menoh_model_handle
}

// Delete object.
func (m *Model) Delete() {
	C.menoh_delete_model(m.h)
	m.h = nil
}

// GetVariable returns Variable, which set the target data information.
func (m *Model) GetVariable(name string) (*Variable, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	var dtype C.int
	if err := checkError(C.menoh_model_get_variable_dtype(m.h, cName, &dtype)); err != nil {
		return nil, err
	}
	var size C.int
	if err := checkError(C.menoh_model_get_variable_dims_size(m.h, cName, &size)); err != nil {
		return nil, err
	}
	dims := make([]int32, size)
	for i := 0; i < int(size); i++ {
		var dim C.int
		if err := checkError(C.menoh_model_get_variable_dims_at(m.h, cName, C.int(i), &dim)); err != nil {
			return nil, err
		}
		dims[i] = int32(dim)
	}
	var ptr unsafe.Pointer
	if err := checkError(C.menoh_model_get_variable_buffer_handle(m.h, cName, &ptr)); err != nil {
		return nil, err
	}

	return &Variable{
		Dtype:        toDtype(dtype),
		Dims:         dims,
		BufferHandle: ptr,
	}, nil
}

// Run calculation.
func (m *Model) Run() error {
	return checkError(C.menoh_model_run(m.h))
}

// TypeMenohDtype binds 'menoh_dtype_constant' enum
type TypeMenohDtype C.menoh_dtype

// Dtype
const (
	TypeFloat TypeMenohDtype = iota
)

func toDtype(typeCode C.int) TypeMenohDtype {
	return TypeMenohDtype(int(typeCode))
}

type typeMenohError C.menoh_error_code

// Menoh Error type
const (
	typeSuccess typeMenohError = iota
	typeSTDError
	typeUnknownError
	typeInvalidFilename
	typeUnsupportedONNXOpsetVersion
	typeONNXParseError
	typeInvalidDtype
	typeInvalidAttributeType
	typeUnsupportedOperatorAttribute
	typeDimensionMismatch
	typeVariableNotFound
	typeIndexOutOfRange
	typeJSONParseError
	typeInvalidBackendName
	typeUnsupportedOperator
	typeFailedToConfigureOperator
	typeBackendError
	typeSameNamedVariableAlreadyExist
	typeUnsupportedInputDims
	typeSameNamedParameterAlreadyExist
	typeSameNamedAttributeAlreadyExist
	typeInvalidBackendConfigError
	typeInputNotFoundError
	typeOutputNotFoundError
)

func (e typeMenohError) String() string {
	switch e {
	case typeSuccess:
		return "success"
	case typeSTDError:
		return "std error"
	case typeUnknownError:
		return "unknown error"
	case typeInvalidFilename:
		return "invalid filename"
	case typeUnsupportedONNXOpsetVersion:
		return "unsupported ONNX opset version"
	case typeONNXParseError:
		return "ONNX parse error"
	case typeInvalidDtype:
		return "invalid dtype"
	case typeInvalidAttributeType:
		return "invalid attribute type"
	case typeUnsupportedOperatorAttribute:
		return "unsupported operator attribute"
	case typeDimensionMismatch:
		return "dimension mismatch"
	case typeVariableNotFound:
		return "variable not found"
	case typeIndexOutOfRange:
		return "index out of range"
	case typeJSONParseError:
		return "JSON parse error"
	case typeInvalidBackendName:
		return "invalid backend name"
	case typeUnsupportedOperator:
		return "unsupported operator"
	case typeFailedToConfigureOperator:
		return "failed to configure operator"
	case typeBackendError:
		return "backend error"
	case typeSameNamedVariableAlreadyExist:
		return "same named variable already exist"
	case typeUnsupportedInputDims:
		return "unsupported input dims"
	case typeSameNamedParameterAlreadyExist:
		return "same named parameter already exist"
	case typeSameNamedAttributeAlreadyExist:
		return "same named attribute already exist"
	case typeInvalidBackendConfigError:
		return "invalid backend config"
	case typeInputNotFoundError:
		return "input not found"
	case typeOutputNotFoundError:
		return "output not found"
	default:
		return "unknown type error"
	}
}

func checkError(errCode C.int) error {
	errType := typeMenohError(int(errCode))
	if errType == typeSuccess {
		return nil
	}
	lastMessage := C.GoString(C.menoh_get_last_error_message())
	if lastMessage == "" {
		return errors.New(errType.String())
	}
	// last message includes error type
	return errors.New(lastMessage)
}
