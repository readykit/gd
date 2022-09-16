// Code generated by /internal/packages/generator DO NOT EDIT
package editor

import "github.com/readykit/gd"

type FeatureProfileFeature int64

const (
	FeatureProfileFeature3d FeatureProfileFeature = 0
	FeatureProfileFeatureScript FeatureProfileFeature = 1
	FeatureProfileFeatureAssetLib FeatureProfileFeature = 2
	FeatureProfileFeatureSceneTree FeatureProfileFeature = 3
	FeatureProfileFeatureNodeDock FeatureProfileFeature = 4
	FeatureProfileFeatureFilesystemDock FeatureProfileFeature = 5
	FeatureProfileFeatureImportDock FeatureProfileFeature = 6
	FeatureProfileFeatureMax FeatureProfileFeature = 7
)

type FileDialogFileMode int64

const (
	FileDialogFileModeOpenFile FileDialogFileMode = 0
	FileDialogFileModeOpenFiles FileDialogFileMode = 1
	FileDialogFileModeOpenDir FileDialogFileMode = 2
	FileDialogFileModeOpenAny FileDialogFileMode = 3
	FileDialogFileModeSaveFile FileDialogFileMode = 4
)

type FileDialogAccess int64

const (
	FileDialogAccessResources FileDialogAccess = 0
	FileDialogAccessUserdata FileDialogAccess = 1
	FileDialogAccessFilesystem FileDialogAccess = 2
)

type FileDialogDisplayMode int64

const (
	FileDialogDisplayThumbnails FileDialogDisplayMode = 0
	FileDialogDisplayList FileDialogDisplayMode = 1
)

type PluginCustomControlContainer int64

const (
	PluginContainerToolbar PluginCustomControlContainer = 0
	PluginContainerSpatialEditorMenu PluginCustomControlContainer = 1
	PluginContainerSpatialEditorSideLeft PluginCustomControlContainer = 2
	PluginContainerSpatialEditorSideRight PluginCustomControlContainer = 3
	PluginContainerSpatialEditorBottom PluginCustomControlContainer = 4
	PluginContainerCanvasEditorMenu PluginCustomControlContainer = 5
	PluginContainerCanvasEditorSideLeft PluginCustomControlContainer = 6
	PluginContainerCanvasEditorSideRight PluginCustomControlContainer = 7
	PluginContainerCanvasEditorBottom PluginCustomControlContainer = 8
	PluginContainerInspectorBottom PluginCustomControlContainer = 9
	PluginContainerProjectSettingTabLeft PluginCustomControlContainer = 10
	PluginContainerProjectSettingTabRight PluginCustomControlContainer = 11
)

type PluginDockSlot int64

const (
	PluginDockSlotLeftUl PluginDockSlot = 0
	PluginDockSlotLeftBl PluginDockSlot = 1
	PluginDockSlotLeftUr PluginDockSlot = 2
	PluginDockSlotLeftBr PluginDockSlot = 3
	PluginDockSlotRightUl PluginDockSlot = 4
	PluginDockSlotRightBl PluginDockSlot = 5
	PluginDockSlotRightUr PluginDockSlot = 6
	PluginDockSlotRightBr PluginDockSlot = 7
	PluginDockSlotMax PluginDockSlot = 8
)

type ScenePostImportPluginInternalImportCategory int64

const (
	ScenePostImportPluginInternalImportCategoryNode ScenePostImportPluginInternalImportCategory = 0
	ScenePostImportPluginInternalImportCategoryMesh3dNode ScenePostImportPluginInternalImportCategory = 1
	ScenePostImportPluginInternalImportCategoryMesh ScenePostImportPluginInternalImportCategory = 2
	ScenePostImportPluginInternalImportCategoryMaterial ScenePostImportPluginInternalImportCategory = 3
	ScenePostImportPluginInternalImportCategoryAnimation ScenePostImportPluginInternalImportCategory = 4
	ScenePostImportPluginInternalImportCategoryAnimationNode ScenePostImportPluginInternalImportCategory = 5
	ScenePostImportPluginInternalImportCategorySkeleton3dNode ScenePostImportPluginInternalImportCategory = 6
	ScenePostImportPluginInternalImportCategoryMax ScenePostImportPluginInternalImportCategory = 7
)

type UndoRedoManagerSpecialHistory int64

const (
	UndoRedoManagerGlobalHistory UndoRedoManagerSpecialHistory = 0
	UndoRedoManagerInvalidHistory UndoRedoManagerSpecialHistory = -99
)

type VCSInterfaceChangeType int64

const (
	VCSInterfaceChangeTypeNew VCSInterfaceChangeType = 0
	VCSInterfaceChangeTypeModified VCSInterfaceChangeType = 1
	VCSInterfaceChangeTypeRenamed VCSInterfaceChangeType = 2
	VCSInterfaceChangeTypeDeleted VCSInterfaceChangeType = 3
	VCSInterfaceChangeTypeTypechange VCSInterfaceChangeType = 4
	VCSInterfaceChangeTypeUnmerged VCSInterfaceChangeType = 5
)

type VCSInterfaceTreeArea int64

const (
	VCSInterfaceTreeAreaCommit VCSInterfaceTreeArea = 0
	VCSInterfaceTreeAreaStaged VCSInterfaceTreeArea = 1
	VCSInterfaceTreeAreaUnstaged VCSInterfaceTreeArea = 2
)

type SurfaceToolCustomFormat int64

const (
	SurfaceToolCustomRgba8Unorm SurfaceToolCustomFormat = 0
	SurfaceToolCustomRgba8Snorm SurfaceToolCustomFormat = 1
	SurfaceToolCustomRgHalf SurfaceToolCustomFormat = 2
	SurfaceToolCustomRgbaHalf SurfaceToolCustomFormat = 3
	SurfaceToolCustomRFloat SurfaceToolCustomFormat = 4
	SurfaceToolCustomRgFloat SurfaceToolCustomFormat = 5
	SurfaceToolCustomRgbFloat SurfaceToolCustomFormat = 6
	SurfaceToolCustomRgbaFloat SurfaceToolCustomFormat = 7
	SurfaceToolCustomMax SurfaceToolCustomFormat = 8
)

type SurfaceToolSkinWeightCount int64

const (
	SurfaceToolSkin4Weights SurfaceToolSkinWeightCount = 0
	SurfaceToolSkin8Weights SurfaceToolSkinWeightCount = 1
)

type UndoRedoMergeMode int64

const (
	UndoRedoMergeDisable UndoRedoMergeMode = 0
	UndoRedoMergeEnds UndoRedoMergeMode = 1
	UndoRedoMergeAll UndoRedoMergeMode = 2
)

type CommandPalette = gd.EditorCommandPalette

type DebuggerPlugin = gd.EditorDebuggerPlugin

type ExportPlatform = gd.EditorExportPlatform

type ExportPlugin = gd.EditorExportPlugin

type FeatureProfile = gd.EditorFeatureProfile

type FileDialog = gd.EditorFileDialog

type FileSystem = gd.EditorFileSystem

type FileSystemDirectory = gd.EditorFileSystemDirectory

type FileSystemImportFormatSupportQuery = gd.EditorFileSystemImportFormatSupportQuery

type ImportPlugin = gd.EditorImportPlugin

type Inspector = gd.EditorInspector

type InspectorPlugin = gd.EditorInspectorPlugin

type Interface = gd.EditorInterface

type Node3DGizmo = gd.EditorNode3DGizmo

type Node3DGizmoPlugin = gd.EditorNode3DGizmoPlugin

type Paths = gd.EditorPaths

type Plugin = gd.EditorPlugin

type Property = gd.EditorProperty

type ResourceConversionPlugin = gd.EditorResourceConversionPlugin

type ResourcePicker = gd.EditorResourcePicker

type ResourcePreview = gd.EditorResourcePreview

type ResourcePreviewGenerator = gd.EditorResourcePreviewGenerator

type SceneFormatImporter = gd.EditorSceneFormatImporter

type SceneFormatImporterBlend = gd.EditorSceneFormatImporterBlend

type SceneFormatImporterFBX = gd.EditorSceneFormatImporterFBX

type SceneFormatImporterGLTF = gd.EditorSceneFormatImporterGLTF

type ScenePostImport = gd.EditorScenePostImport

type ScenePostImportPlugin = gd.EditorScenePostImportPlugin

type Script = gd.EditorScript

type ScriptPicker = gd.EditorScriptPicker

type Selection = gd.EditorSelection

type Settings = gd.EditorSettings

type SpinSlider = gd.EditorSpinSlider

type SyntaxHighlighter = gd.EditorSyntaxHighlighter

type TranslationParserPlugin = gd.EditorTranslationParserPlugin

type UndoRedoManager = gd.EditorUndoRedoManager

type VCSInterface = gd.EditorVCSInterface

type EncodedObjectAsID = gd.EncodedObjectAsID

type GridMap = gd.GridMap

type Marker2D = gd.Marker2D

type Marker3D = gd.Marker3D

type MissingNode = gd.MissingNode

type MissingResource = gd.MissingResource

type SurfaceTool = gd.SurfaceTool

type UndoRedo = gd.UndoRedo