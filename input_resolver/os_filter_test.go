package input_resolver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TEST1020: macOS .DS_Store is excluded
func Test1020_ds_store_excluded(t *testing.T) {
	assert.True(t, ShouldExclude("/some/path/.DS_Store"))
	assert.True(t, ShouldExclude(".DS_Store"))
}

// TEST1021: Windows Thumbs.db is excluded
func Test1021_thumbs_db_excluded(t *testing.T) {
	assert.True(t, ShouldExclude("/some/path/Thumbs.db"))
	assert.True(t, ShouldExclude("Thumbs.db"))
}

// TEST1022: macOS resource fork files are excluded
func Test1022_resource_fork_excluded(t *testing.T) {
	assert.True(t, ShouldExclude("/path/._file.txt"))
	assert.True(t, ShouldExclude("._anything"))
}

// TEST1023: Office lock files are excluded
func Test1023_office_lock_excluded(t *testing.T) {
	assert.True(t, ShouldExclude("/path/~$document.docx"))
	assert.True(t, ShouldExclude("~$spreadsheet.xlsx"))
}

// TEST1024: .git directory is excluded
func Test1024_git_dir_excluded(t *testing.T) {
	assert.True(t, ShouldExcludeDir("/repo/.git"))
	assert.True(t, ShouldExcludeDir(".git"))
}

// TEST1025: __MACOSX archive artifact is excluded
func Test1025_macosx_dir_excluded(t *testing.T) {
	assert.True(t, ShouldExcludeDir("/extracted/__MACOSX"))
	assert.True(t, ShouldExcludeDir("__MACOSX"))
}

// TEST1026: Temp files are excluded
func Test1026_temp_files_excluded(t *testing.T) {
	assert.True(t, ShouldExclude("/path/file.tmp"))
	assert.True(t, ShouldExclude("/path/file.temp"))
	assert.True(t, ShouldExclude("/path/file.swp"))
	assert.True(t, ShouldExclude("/path/file.bak"))
}

// TEST1027: .localized is excluded
func Test1027_localized_excluded(t *testing.T) {
	assert.True(t, ShouldExclude("/path/.localized"))
}

// TEST1028: desktop.ini is excluded
func Test1028_desktop_ini_excluded(t *testing.T) {
	assert.True(t, ShouldExclude("/path/desktop.ini"))
}

// TEST1029: Normal files are NOT excluded
func Test1029_normal_files_not_excluded(t *testing.T) {
	assert.False(t, ShouldExclude("/path/file.txt"))
	assert.False(t, ShouldExclude("/path/data.json"))
	assert.False(t, ShouldExclude("/path/notes.md"))
	assert.False(t, ShouldExclude("/path/.gitignore")) // Config file, keep
	assert.False(t, ShouldExclude("/path/.env"))       // Config file, keep
	assert.False(t, ShouldExclude("/path/README.md"))
}
