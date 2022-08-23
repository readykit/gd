package gdnative

// #include <gdnative_interface.h>
import "C"

import (
	"fmt"
	"reflect"
	"runtime/cgo"
	"strconv"
	"unsafe"
)

//export loadExtension
func loadExtension(api, library, init unsafe.Pointer) uint8 {
	load(api, library, init)
	return 1
}

//export goInitialise
func goInitialise(userdata unsafe.Pointer, level InitializationLevel) {
	fmt.Println("initialising ", level)
	if level == 3 {
		for _, f := range onload {
			f()
		}
		for _, f := range oninit {
			f()
		}
	}
}

//export goDeinitialize
func goDeinitialize(userdata unsafe.Pointer, level InitializationLevel) {
	fmt.Println("deinitialize ", level)
}

//export goClassCreateInstance
func goClassCreateInstance(userdata uintptr) uintptr {
	return uintptr(cgo.Handle(userdata).Value().(Class).Create())
}

//export goClassGetVirtual
func goClassGetVirtual(userdata, name *C.char) {
	fmt.Println(C.GoString(name))
}

//export goClassFreeInstance
func goClassFreeInstance(userdata, instance uintptr) {
	cgo.Handle(userdata).Value().(Class).Delete(InstanceID(instance))
}

var (
	typeVector2            = reflect.TypeOf(Vector2{})
	typeVector2i           = reflect.TypeOf(Vector2i{})
	typeRect2              = reflect.TypeOf(Rect2{})
	typeRect2i             = reflect.TypeOf(Rect2i{})
	typeVector3            = reflect.TypeOf(Vector3{})
	typeVector3i           = reflect.TypeOf(Vector3i{})
	typeTransform2D        = reflect.TypeOf(Transform2D{})
	typeVector4            = reflect.TypeOf(Vector4{})
	typeVector4i           = reflect.TypeOf(Vector4i{})
	typePlane              = reflect.TypeOf(Plane{})
	typeQuaternion         = reflect.TypeOf(Quaternion{})
	typeAABB               = reflect.TypeOf(AABB{})
	typeBasis              = reflect.TypeOf(Basis{})
	typeTransform3D        = reflect.TypeOf(Transform3D{})
	typeProjection         = reflect.TypeOf(Projection{})
	typeColor              = reflect.TypeOf(Color{})
	typeStringName         = reflect.TypeOf(StringName{})
	typeNodePath           = reflect.TypeOf(NodePath{})
	typeRID                = reflect.TypeOf(RID{})
	typeObject             = reflect.TypeOf(Object(0))
	typeCallable           = reflect.TypeOf(Callable{})
	typeSignal             = reflect.TypeOf(Signal{})
	typeDictionary         = reflect.TypeOf(Dictionary{})
	typeArray              = reflect.TypeOf(Array{})
	typePackedByteArray    = reflect.TypeOf(PackedByteArray{})
	typePackedInt32Array   = reflect.TypeOf(PackedInt32Array{})
	typePackedInt64Array   = reflect.TypeOf(PackedInt64Array{})
	typePackedFloat32Array = reflect.TypeOf(PackedFloat32Array{})
	typePackedFloat64Array = reflect.TypeOf(PackedFloat64Array{})
	typePackedStringArray  = reflect.TypeOf(PackedStringArray{})
	typePackedVector2Array = reflect.TypeOf(PackedVector2Array{})
	typePackedVector3Array = reflect.TypeOf(PackedVector3Array{})
	typePackedColorArray   = reflect.TypeOf(PackedColorArray{})
)

//export goMethodGetArgumentInfo
func goMethodGetArgumentInfo(userdata uintptr, argument int32, info *C.GDNativePropertyInfo) {
	info._type = C.uint32_t(goMethodGetArgumentType(userdata, argument))
	info.name = cString("arg" + strconv.Itoa(int(argument))) // FIXME does godot take ownership of this?
}

//export goMethodGetArgumentMetadata
func goMethodGetArgumentMetadata(userdata uintptr, argument int32) C.GDNativeExtensionClassMethodArgumentMetadata {
	caller := cgo.Handle(userdata).Value().(methodCaller)

	var rtype reflect.Type
	if argument == -1 {
		if caller.method.Type.NumOut() == 0 {
			return C.GDNATIVE_VARIANT_TYPE_NIL
		}
		rtype = caller.method.Type.Out(0) // return value
	} else {
		rtype = caller.method.Type.In(int(argument) + 1) // skip receiver
	}

	switch rtype.Kind() {
	case reflect.Int:
		if rtype.Size() == 4 {
			return C.GDNATIVE_EXTENSION_METHOD_ARGUMENT_METADATA_INT_IS_INT32
		}
		return C.GDNATIVE_EXTENSION_METHOD_ARGUMENT_METADATA_INT_IS_INT64
	case reflect.Int8:
		return C.GDNATIVE_EXTENSION_METHOD_ARGUMENT_METADATA_INT_IS_INT8
	case reflect.Int16:
		return C.GDNATIVE_EXTENSION_METHOD_ARGUMENT_METADATA_INT_IS_INT16
	case reflect.Int32:
		return C.GDNATIVE_EXTENSION_METHOD_ARGUMENT_METADATA_INT_IS_INT32
	case reflect.Int64:
		return C.GDNATIVE_EXTENSION_METHOD_ARGUMENT_METADATA_INT_IS_INT64
	case reflect.Uint:
		if rtype.Size() == 4 {
			return C.GDNATIVE_EXTENSION_METHOD_ARGUMENT_METADATA_INT_IS_UINT32
		}
		return C.GDNATIVE_EXTENSION_METHOD_ARGUMENT_METADATA_INT_IS_UINT64
	case reflect.Uint8:
		return C.GDNATIVE_EXTENSION_METHOD_ARGUMENT_METADATA_INT_IS_UINT8
	case reflect.Uint16:
		return C.GDNATIVE_EXTENSION_METHOD_ARGUMENT_METADATA_INT_IS_UINT16
	case reflect.Uint32:
		return C.GDNATIVE_EXTENSION_METHOD_ARGUMENT_METADATA_INT_IS_UINT32
	case reflect.Uint64:
		return C.GDNATIVE_EXTENSION_METHOD_ARGUMENT_METADATA_INT_IS_UINT64
	case reflect.Float32:
		return C.GDNATIVE_EXTENSION_METHOD_ARGUMENT_METADATA_REAL_IS_FLOAT
	case reflect.Float64:
		return C.GDNATIVE_EXTENSION_METHOD_ARGUMENT_METADATA_REAL_IS_DOUBLE
	default:
		return C.GDNATIVE_EXTENSION_METHOD_ARGUMENT_METADATA_NONE
	}
}

//export goMethodGetArgumentType
func goMethodGetArgumentType(userdata uintptr, argument int32) C.GDNativeVariantType {
	caller := cgo.Handle(userdata).Value().(methodCaller)

	var rtype reflect.Type
	if argument == -1 {
		if caller.method.Type.NumOut() == 0 {
			return C.GDNATIVE_VARIANT_TYPE_NIL
		}
		rtype = caller.method.Type.Out(0) // return value
	} else {
		rtype = caller.method.Type.In(int(argument) + 1) // skip receiver
	}

	switch rtype.Kind() {
	case reflect.Bool:
		return C.GDNATIVE_VARIANT_TYPE_BOOL
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return C.GDNATIVE_VARIANT_TYPE_INT
	case reflect.Float32, reflect.Float64:
		return C.GDNATIVE_VARIANT_TYPE_FLOAT
	case reflect.String:
		return C.GDNATIVE_VARIANT_TYPE_STRING
	default:
		switch rtype {
		case typeVector2:
			return C.GDNATIVE_VARIANT_TYPE_VECTOR2
		case typeVector2i:
			return C.GDNATIVE_VARIANT_TYPE_VECTOR2I
		case typeRect2:
			return C.GDNATIVE_VARIANT_TYPE_RECT2
		case typeRect2i:
			return C.GDNATIVE_VARIANT_TYPE_RECT2I
		case typeVector3:
			return C.GDNATIVE_VARIANT_TYPE_VECTOR3
		case typeVector3i:
			return C.GDNATIVE_VARIANT_TYPE_VECTOR3I
		case typeTransform2D:
			return C.GDNATIVE_VARIANT_TYPE_TRANSFORM2D
		case typeVector4:
			return C.GDNATIVE_VARIANT_TYPE_VECTOR4
		case typeVector4i:
			return C.GDNATIVE_VARIANT_TYPE_VECTOR4I
		case typePlane:
			return C.GDNATIVE_VARIANT_TYPE_PLANE
		case typeQuaternion:
			return C.GDNATIVE_VARIANT_TYPE_QUATERNION
		case typeAABB:
			return C.GDNATIVE_VARIANT_TYPE_AABB
		case typeBasis:
			return C.GDNATIVE_VARIANT_TYPE_BASIS
		case typeTransform3D:
			return C.GDNATIVE_VARIANT_TYPE_TRANSFORM3D
		case typeProjection:
			return C.GDNATIVE_VARIANT_TYPE_PROJECTION
		case typeColor:
			return C.GDNATIVE_VARIANT_TYPE_COLOR
		case typeStringName:
			return C.GDNATIVE_VARIANT_TYPE_STRING_NAME
		case typeNodePath:
			return C.GDNATIVE_VARIANT_TYPE_NODE_PATH
		case typeRID:
			return C.GDNATIVE_VARIANT_TYPE_RID
		case typeObject:
			return C.GDNATIVE_VARIANT_TYPE_OBJECT
		case typeCallable:
			return C.GDNATIVE_VARIANT_TYPE_CALLABLE
		case typeSignal:
			return C.GDNATIVE_VARIANT_TYPE_SIGNAL
		case typeDictionary:
			return C.GDNATIVE_VARIANT_TYPE_DICTIONARY
		case typeArray:
			return C.GDNATIVE_VARIANT_TYPE_ARRAY
		case typePackedByteArray:
			return C.GDNATIVE_VARIANT_TYPE_PACKED_BYTE_ARRAY
		case typePackedInt32Array:
			return C.GDNATIVE_VARIANT_TYPE_PACKED_INT32_ARRAY
		case typePackedInt64Array:
			return C.GDNATIVE_VARIANT_TYPE_PACKED_INT64_ARRAY
		case typePackedFloat32Array:
			return C.GDNATIVE_VARIANT_TYPE_PACKED_FLOAT32_ARRAY
		case typePackedFloat64Array:
			return C.GDNATIVE_VARIANT_TYPE_PACKED_FLOAT64_ARRAY
		case typePackedStringArray:
			return C.GDNATIVE_VARIANT_TYPE_PACKED_STRING_ARRAY
		case typePackedVector2Array:
			return C.GDNATIVE_VARIANT_TYPE_PACKED_VECTOR2_ARRAY
		case typePackedVector3Array:
			return C.GDNATIVE_VARIANT_TYPE_PACKED_VECTOR3_ARRAY
		case typePackedColorArray:
			return C.GDNATIVE_VARIANT_TYPE_PACKED_COLOR_ARRAY
		default:
			panic(fmt.Sprintf("unsupported godot argument type: %s", rtype))
		}
	}
}

//export goMethodCall
func goMethodCall(userdata, instance uintptr, args *C.GDNativeVariantPtr, count C.GDNativeInt, result C.GDNativeVariantPtr, err *C.GDNativeCallError) {
	caller := cgo.Handle(userdata).Value().(methodCaller)

	fmt.Println(caller.method.Name, " METHOD CALLED")

	caller.cached[0] = caller.lookup(InstanceID(instance))
	if caller.method.Type.NumIn() > 1 {
		slice := unsafe.Slice(args, count)
		value := caller.cached[1:]
		for i := range value {
			value[i] = loadArgFromVariant(caller.method.Type.In(i+1), slice[i])
		}
	}

	fmt.Println(caller.cached)

	caller.method.Func.Call(caller.cached)
}

//export goMethodCallDirect
func goMethodCallDirect(userdata, instance uintptr, ptrs *C.GDNativeTypePtr, result C.GDNativeTypePtr) {
	caller := cgo.Handle(userdata).Value().(methodCaller)

	fmt.Println(caller.method.Name, " METHOD CALLED DIRECTLY")

	caller.cached[0] = caller.lookup(InstanceID(instance))
	if caller.method.Type.NumIn() > 1 {
		slice := unsafe.Slice(ptrs, caller.method.Type.NumIn()-1)
		value := caller.cached[1:]
		for i := range value {
			value[i] = loadArgFromNativeType(caller.method.Type.In(i+1), slice[i])
		}
	}

	fmt.Println(caller.cached)

	caller.method.Func.Call(caller.cached)
}

// Main should be called in your extension's main function.
func Main() {}

var oninit []func()

func OnLoad(fn func()) {
	oninit = append(oninit, fn)
}
