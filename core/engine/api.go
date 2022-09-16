// Code generated by /internal/packages/generator DO NOT EDIT
package engine

import "github.com/readykit/gd"

type PerformanceSingletonMonitor int64

const (
	PerformanceSingletonTimeFps PerformanceSingletonMonitor = 0
	PerformanceSingletonTimeProcess PerformanceSingletonMonitor = 1
	PerformanceSingletonTimePhysicsProcess PerformanceSingletonMonitor = 2
	PerformanceSingletonMemoryStatic PerformanceSingletonMonitor = 3
	PerformanceSingletonMemoryStaticMax PerformanceSingletonMonitor = 4
	PerformanceSingletonMemoryMessageBufferMax PerformanceSingletonMonitor = 5
	PerformanceSingletonObjectCount PerformanceSingletonMonitor = 6
	PerformanceSingletonObjectResourceCount PerformanceSingletonMonitor = 7
	PerformanceSingletonObjectNodeCount PerformanceSingletonMonitor = 8
	PerformanceSingletonObjectOrphanNodeCount PerformanceSingletonMonitor = 9
	PerformanceSingletonRenderTotalObjectsInFrame PerformanceSingletonMonitor = 10
	PerformanceSingletonRenderTotalPrimitivesInFrame PerformanceSingletonMonitor = 11
	PerformanceSingletonRenderTotalDrawCallsInFrame PerformanceSingletonMonitor = 12
	PerformanceSingletonRenderVideoMemUsed PerformanceSingletonMonitor = 13
	PerformanceSingletonRenderTextureMemUsed PerformanceSingletonMonitor = 14
	PerformanceSingletonRenderBufferMemUsed PerformanceSingletonMonitor = 15
	PerformanceSingletonPhysics2dActiveObjects PerformanceSingletonMonitor = 16
	PerformanceSingletonPhysics2dCollisionPairs PerformanceSingletonMonitor = 17
	PerformanceSingletonPhysics2dIslandCount PerformanceSingletonMonitor = 18
	PerformanceSingletonPhysics3dActiveObjects PerformanceSingletonMonitor = 19
	PerformanceSingletonPhysics3dCollisionPairs PerformanceSingletonMonitor = 20
	PerformanceSingletonPhysics3dIslandCount PerformanceSingletonMonitor = 21
	PerformanceSingletonAudioOutputLatency PerformanceSingletonMonitor = 22
	PerformanceSingletonMonitorMax PerformanceSingletonMonitor = 23
)

var Class ClassSingleton

type ClassSingleton = gd.EngineSingleton

var Debugger DebuggerSingleton

type DebuggerSingleton = gd.EngineDebuggerSingleton

type Profiler = gd.EngineProfiler

type MainLoop = gd.MainLoop

var Performance PerformanceSingleton

type PerformanceSingleton = gd.PerformanceSingleton