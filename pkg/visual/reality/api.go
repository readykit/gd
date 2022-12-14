// Code generated by /internal/packages/generator DO NOT EDIT
package reality

import "github.com/readykit/gd"

type ActionType int64

const (
	OpenActionBool ActionType = 0
	OpenActionFloat ActionType = 1
	OpenActionVector2 ActionType = 2
	OpenActionPose ActionType = 3
)

type OpenHandHands int64

const (
	OpenHandHandLeft OpenHandHands = 0
	OpenHandHandRight OpenHandHands = 1
	OpenHandHandMax OpenHandHands = 2
)

type OpenHandMotionRange int64

const (
	OpenHandMotionRangeUnobstructed OpenHandMotionRange = 0
	OpenHandMotionRangeConformToController OpenHandMotionRange = 1
	OpenHandMotionRangeMax OpenHandMotionRange = 2
)

type InterfaceCapabilities int64

const (
	InterfaceNone InterfaceCapabilities = 0
	InterfaceMono InterfaceCapabilities = 1
	InterfaceStereo InterfaceCapabilities = 2
	InterfaceQuad InterfaceCapabilities = 4
	InterfaceVr InterfaceCapabilities = 8
	InterfaceAr InterfaceCapabilities = 16
	InterfaceExternal InterfaceCapabilities = 32
)

type InterfaceTrackingStatus int64

const (
	InterfaceNormalTracking InterfaceTrackingStatus = 0
	InterfaceExcessiveMotion InterfaceTrackingStatus = 1
	InterfaceInsufficientFeatures InterfaceTrackingStatus = 2
	InterfaceUnknownTracking InterfaceTrackingStatus = 3
	InterfaceNotTracking InterfaceTrackingStatus = 4
)

type InterfacePlayAreaMode int64

const (
	InterfacePlayAreaUnknown InterfacePlayAreaMode = 0
	InterfacePlayArea3dof InterfacePlayAreaMode = 1
	InterfacePlayAreaSitting InterfacePlayAreaMode = 2
	InterfacePlayAreaRoomscale InterfacePlayAreaMode = 3
	InterfacePlayAreaStage InterfacePlayAreaMode = 4
)

type PoseTrackingConfidence int64

const (
	PoseTrackingConfidenceNone PoseTrackingConfidence = 0
	PoseTrackingConfidenceLow PoseTrackingConfidence = 1
	PoseTrackingConfidenceHigh PoseTrackingConfidence = 2
)

type PositionalTrackerTrackerHand int64

const (
	PositionalTrackerTrackerHandUnknown PositionalTrackerTrackerHand = 0
	PositionalTrackerTrackerHandLeft PositionalTrackerTrackerHand = 1
	PositionalTrackerTrackerHandRight PositionalTrackerTrackerHand = 2
)

type ServerSingletonTrackerType int64

const (
	ServerSingletonTrackerHead ServerSingletonTrackerType = 1
	ServerSingletonTrackerController ServerSingletonTrackerType = 2
	ServerSingletonTrackerBasestation ServerSingletonTrackerType = 4
	ServerSingletonTrackerAnchor ServerSingletonTrackerType = 8
	ServerSingletonTrackerAnyKnown ServerSingletonTrackerType = 127
	ServerSingletonTrackerUnknown ServerSingletonTrackerType = 128
	ServerSingletonTrackerAny ServerSingletonTrackerType = 255
)

type ServerSingletonRotationMode int64

const (
	ServerSingletonResetFullRotation ServerSingletonRotationMode = 0
	ServerSingletonResetButKeepTilt ServerSingletonRotationMode = 1
	ServerSingletonDontResetRotation ServerSingletonRotationMode = 2
)

type MobileVRInterface = gd.MobileVRInterface

type OpenAction = gd.OpenXRAction

type OpenActionMap = gd.OpenXRActionMap

type OpenActionSet = gd.OpenXRActionSet

type OpenHand = gd.OpenXRHand

type OpenInteractionProfile = gd.OpenXRInteractionProfile

type OpenInterface = gd.OpenXRInterface

type OpenIPBinding = gd.OpenXRIPBinding

type WebInterface = gd.WebXRInterface

type Anchor3D = gd.XRAnchor3D

type Camera3D = gd.XRCamera3D

type Controller3D = gd.XRController3D

type Interface = gd.XRInterface

type InterfaceExtension = gd.XRInterfaceExtension

type Node3D = gd.XRNode3D

type Origin3D = gd.XROrigin3D

type Pose = gd.XRPose

type PositionalTracker = gd.XRPositionalTracker

var Server ServerSingleton

type ServerSingleton = gd.XRServerSingleton
