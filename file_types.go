package media_sidecar

import (
	"github.com/smartmediafiles/media/media/maps"
	"github.com/smartmediafiles/media/media/types"
)

// List of supported media.Sidecar file types.
const (
	SidecarXMP      types.FileType = "xmp"  // Adobe XMP sidecar file (XML)
	SidecarXML      types.FileType = "xml"  // XML metadata / config / sidecar file
	SidecarAAE      types.FileType = "aae"  // Apple image edits sidecar file (based on XML)
	SidecarYAML     types.FileType = "yml"  // YAML metadata / config / sidecar file
	SidecarJSON     types.FileType = "json" // JSON metadata / config / sidecar file
	SidecarText     types.FileType = "txt"  // Text config / sidecar file
	SidecarInfo     types.FileType = "nfo"  // Info text file as used by e.g. Plex Media Server
	SidecarMarkdown types.FileType = "md"   // Markdown text sidecar file
)

// SidecarFileTypesExtensions is a map of media.Sidecar file types to their file extensions.
var SidecarFileTypesExtensions = maps.MapFileTypeExtensions{
	SidecarXMP:      {".xmp"},
	SidecarXML:      {".xml"},
	SidecarAAE:      {".aae"},
	SidecarYAML:     {".yml", ".yaml"},
	SidecarJSON:     {".json"},
	SidecarText:     {".txt"},
	SidecarInfo:     {".nfo"},
	SidecarMarkdown: {".md", ".markdown"},
}
