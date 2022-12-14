// Code generated by /internal/packages/generator DO NOT EDIT
package webrtc

import "github.com/readykit/gd"

type DataChannelWriteMode int64

const (
	DataChannelWriteModeText DataChannelWriteMode = 0
	DataChannelWriteModeBinary DataChannelWriteMode = 1
)

type DataChannelChannelState int64

const (
	DataChannelStateConnecting DataChannelChannelState = 0
	DataChannelStateOpen DataChannelChannelState = 1
	DataChannelStateClosing DataChannelChannelState = 2
	DataChannelStateClosed DataChannelChannelState = 3
)

type PeerConnectionConnectionState int64

const (
	PeerConnectionStateNew PeerConnectionConnectionState = 0
	PeerConnectionStateConnecting PeerConnectionConnectionState = 1
	PeerConnectionStateConnected PeerConnectionConnectionState = 2
	PeerConnectionStateDisconnected PeerConnectionConnectionState = 3
	PeerConnectionStateFailed PeerConnectionConnectionState = 4
	PeerConnectionStateClosed PeerConnectionConnectionState = 5
)

type PeerConnectionGatheringState int64

const (
	PeerConnectionGatheringStateNew PeerConnectionGatheringState = 0
	PeerConnectionGatheringStateGathering PeerConnectionGatheringState = 1
	PeerConnectionGatheringStateComplete PeerConnectionGatheringState = 2
)

type PeerConnectionSignalingState int64

const (
	PeerConnectionSignalingStateStable PeerConnectionSignalingState = 0
	PeerConnectionSignalingStateHaveLocalOffer PeerConnectionSignalingState = 1
	PeerConnectionSignalingStateHaveRemoteOffer PeerConnectionSignalingState = 2
	PeerConnectionSignalingStateHaveLocalPranswer PeerConnectionSignalingState = 3
	PeerConnectionSignalingStateHaveRemotePranswer PeerConnectionSignalingState = 4
	PeerConnectionSignalingStateClosed PeerConnectionSignalingState = 5
)

type DataChannel = gd.WebRTCDataChannel

type DataChannelExtension = gd.WebRTCDataChannelExtension

type MultiplayerPeer = gd.WebRTCMultiplayerPeer

type PeerConnection = gd.WebRTCPeerConnection

type PeerConnectionExtension = gd.WebRTCPeerConnectionExtension
