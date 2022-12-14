// Code generated by /internal/packages/generator DO NOT EDIT
package visual

import "github.com/readykit/gd"

type GradientInterpolationMode int64

const (
	GradientGradientInterpolateLinear GradientInterpolationMode = 0
	GradientGradientInterpolateConstant GradientInterpolationMode = 1
	GradientGradientInterpolateCubic GradientInterpolationMode = 2
)

type ImageFormat int64

const (
	ImageFormatL8 ImageFormat = 0
	ImageFormatLa8 ImageFormat = 1
	ImageFormatR8 ImageFormat = 2
	ImageFormatRg8 ImageFormat = 3
	ImageFormatRgb8 ImageFormat = 4
	ImageFormatRgba8 ImageFormat = 5
	ImageFormatRgba4444 ImageFormat = 6
	ImageFormatRgb565 ImageFormat = 7
	ImageFormatRf ImageFormat = 8
	ImageFormatRgf ImageFormat = 9
	ImageFormatRgbf ImageFormat = 10
	ImageFormatRgbaf ImageFormat = 11
	ImageFormatRh ImageFormat = 12
	ImageFormatRgh ImageFormat = 13
	ImageFormatRgbh ImageFormat = 14
	ImageFormatRgbah ImageFormat = 15
	ImageFormatRgbe9995 ImageFormat = 16
	ImageFormatDxt1 ImageFormat = 17
	ImageFormatDxt3 ImageFormat = 18
	ImageFormatDxt5 ImageFormat = 19
	ImageFormatRgtcR ImageFormat = 20
	ImageFormatRgtcRg ImageFormat = 21
	ImageFormatBptcRgba ImageFormat = 22
	ImageFormatBptcRgbf ImageFormat = 23
	ImageFormatBptcRgbfu ImageFormat = 24
	ImageFormatEtc ImageFormat = 25
	ImageFormatEtc2R11 ImageFormat = 26
	ImageFormatEtc2R11s ImageFormat = 27
	ImageFormatEtc2Rg11 ImageFormat = 28
	ImageFormatEtc2Rg11s ImageFormat = 29
	ImageFormatEtc2Rgb8 ImageFormat = 30
	ImageFormatEtc2Rgba8 ImageFormat = 31
	ImageFormatEtc2Rgb8a1 ImageFormat = 32
	ImageFormatEtc2RaAsRg ImageFormat = 33
	ImageFormatDxt5RaAsRg ImageFormat = 34
	ImageFormatMax ImageFormat = 35
)

type ImageInterpolation int64

const (
	ImageInterpolateNearest ImageInterpolation = 0
	ImageInterpolateBilinear ImageInterpolation = 1
	ImageInterpolateCubic ImageInterpolation = 2
	ImageInterpolateTrilinear ImageInterpolation = 3
	ImageInterpolateLanczos ImageInterpolation = 4
)

type ImageAlphaMode int64

const (
	ImageAlphaNone ImageAlphaMode = 0
	ImageAlphaBit ImageAlphaMode = 1
	ImageAlphaBlend ImageAlphaMode = 2
)

type ImageCompressMode int64

const (
	ImageCompressS3tc ImageCompressMode = 0
	ImageCompressEtc ImageCompressMode = 1
	ImageCompressEtc2 ImageCompressMode = 2
	ImageCompressBptc ImageCompressMode = 3
)

type ImageUsedChannels int64

const (
	ImageUsedChannelsL ImageUsedChannels = 0
	ImageUsedChannelsLa ImageUsedChannels = 1
	ImageUsedChannelsR ImageUsedChannels = 2
	ImageUsedChannelsRg ImageUsedChannels = 3
	ImageUsedChannelsRgb ImageUsedChannels = 4
	ImageUsedChannelsRgba ImageUsedChannels = 5
)

type ImageCompressSource int64

const (
	ImageCompressSourceGeneric ImageCompressSource = 0
	ImageCompressSourceSrgb ImageCompressSource = 1
	ImageCompressSourceNormal ImageCompressSource = 2
)

type ImageFormatLoaderLoaderFlags int64

const (
	ImageFormatLoaderFlagNone ImageFormatLoaderLoaderFlags = 0
	ImageFormatLoaderFlagForceLinear ImageFormatLoaderLoaderFlags = 1
	ImageFormatLoaderFlagConvertColors ImageFormatLoaderLoaderFlags = 2
)

type TileMapVisibilityMode int64

const (
	TileMapVisibilityModeDefault TileMapVisibilityMode = 0
	TileMapVisibilityModeForceHide TileMapVisibilityMode = 2
	TileMapVisibilityModeForceShow TileMapVisibilityMode = 1
)

type TileSetTileShape int64

const (
	TileSetTileShapeSquare TileSetTileShape = 0
	TileSetTileShapeIsometric TileSetTileShape = 1
	TileSetTileShapeHalfOffsetSquare TileSetTileShape = 2
	TileSetTileShapeHexagon TileSetTileShape = 3
)

type TileSetTileLayout int64

const (
	TileSetTileLayoutStacked TileSetTileLayout = 0
	TileSetTileLayoutStackedOffset TileSetTileLayout = 1
	TileSetTileLayoutStairsRight TileSetTileLayout = 2
	TileSetTileLayoutStairsDown TileSetTileLayout = 3
	TileSetTileLayoutDiamondRight TileSetTileLayout = 4
	TileSetTileLayoutDiamondDown TileSetTileLayout = 5
)

type TileSetTileOffsetAxis int64

const (
	TileSetTileOffsetAxisHorizontal TileSetTileOffsetAxis = 0
	TileSetTileOffsetAxisVertical TileSetTileOffsetAxis = 1
)

type TileSetCellNeighbor int64

const (
	TileSetCellNeighborRightSide TileSetCellNeighbor = 0
	TileSetCellNeighborRightCorner TileSetCellNeighbor = 1
	TileSetCellNeighborBottomRightSide TileSetCellNeighbor = 2
	TileSetCellNeighborBottomRightCorner TileSetCellNeighbor = 3
	TileSetCellNeighborBottomSide TileSetCellNeighbor = 4
	TileSetCellNeighborBottomCorner TileSetCellNeighbor = 5
	TileSetCellNeighborBottomLeftSide TileSetCellNeighbor = 6
	TileSetCellNeighborBottomLeftCorner TileSetCellNeighbor = 7
	TileSetCellNeighborLeftSide TileSetCellNeighbor = 8
	TileSetCellNeighborLeftCorner TileSetCellNeighbor = 9
	TileSetCellNeighborTopLeftSide TileSetCellNeighbor = 10
	TileSetCellNeighborTopLeftCorner TileSetCellNeighbor = 11
	TileSetCellNeighborTopSide TileSetCellNeighbor = 12
	TileSetCellNeighborTopCorner TileSetCellNeighbor = 13
	TileSetCellNeighborTopRightSide TileSetCellNeighbor = 14
	TileSetCellNeighborTopRightCorner TileSetCellNeighbor = 15
)

type TileSetTerrainMode int64

const (
	TileSetTerrainModeMatchCornersAndSides TileSetTerrainMode = 0
	TileSetTerrainModeMatchCorners TileSetTerrainMode = 1
	TileSetTerrainModeMatchSides TileSetTerrainMode = 2
)

type Color = gd.Color

type Gradient = gd.Gradient

type Image = gd.Image

type ImageFormatLoader = gd.ImageFormatLoader

type ImageFormatLoaderExtension = gd.ImageFormatLoaderExtension

type ImageTexture = gd.ImageTexture

type ImageTexture3D = gd.ImageTexture3D

type ImageTextureLayered = gd.ImageTextureLayered

type MovieWriter = gd.MovieWriter

type MovieWriterMJPEG = gd.MovieWriterMJPEG

type MovieWriterPNGWAV = gd.MovieWriterPNGWAV

type TileData = gd.TileData

type TileMap = gd.TileMap

type TileMapPattern = gd.TileMapPattern

type TileSet = gd.TileSet

type TileSetAtlasSource = gd.TileSetAtlasSource

type TileSetScenesCollectionSource = gd.TileSetScenesCollectionSource

type TileSetSource = gd.TileSetSource

type VideoStream = gd.VideoStream

type VideoStreamTheora = gd.VideoStreamTheora
