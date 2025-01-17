package astilectron

import (
	"testing"

	"os"

	"path/filepath"

	"github.com/stretchr/testify/assert"
)

func TestPaths(t *testing.T) {
	const k = "APPDATA"

	ad := os.Getenv(k)
	os.Setenv(k, "")
	ep, err := os.Executable()
	ep = filepath.Dir(ep)
	assert.NoError(t, err)

	o := Options{VersionAstilectron: DefaultVersionAstilectron, VersionElectron: DefaultVersionElectron}
	p, err := newPaths("linux", "amd64", o)
	assert.NoError(t, err)
	assert.Equal(t, ep+"/vendor/electron-linux-amd64/electron", p.AppExecutable())
	assert.Equal(t, "", p.AppIconDarwinSrc())
	assert.Equal(t, ep, p.BaseDirectory())
	assert.Equal(t, ep, p.DataDirectory())
	assert.Equal(t, ep+"/vendor/astilectron/main.js", p.AstilectronApplication())
	assert.Equal(t, ep+"/vendor/astilectron", p.AstilectronDirectory())
	assert.Equal(t, ep+"/vendor/astilectron-v"+DefaultVersionAstilectron+".zip", p.AstilectronDownloadDst())
	assert.Equal(t, "https://github.com/asticode/astilectron/archive/v"+DefaultVersionAstilectron+".zip", p.AstilectronDownloadSrc())
	assert.Equal(t, ep+"/vendor/astilectron-v"+DefaultVersionAstilectron+".zip/astilectron-"+DefaultVersionAstilectron, p.AstilectronUnzipSrc())
	assert.Equal(t, ep+"/vendor/electron-linux-amd64", p.ElectronDirectory())
	assert.Equal(t, ep+"/vendor/electron-linux-amd64-v"+DefaultVersionElectron+".zip", p.ElectronDownloadDst())
	assert.Equal(t, "https://github.com/electron/electron/releases/download/v"+DefaultVersionElectron+"/electron-v"+DefaultVersionElectron+"-linux-x64.zip", p.ElectronDownloadSrc())
	assert.Equal(t, ep+"/vendor/electron-linux-amd64-v"+DefaultVersionElectron+".zip", p.ElectronUnzipSrc())
	assert.Equal(t, ep+"/vendor/status.json", p.ProvisionStatus())
	assert.Equal(t, ep+"/vendor", p.VendorDirectory())
	p, err = newPaths("linux", "", o)
	assert.NoError(t, err)
	assert.Equal(t, "https://github.com/electron/electron/releases/download/v"+DefaultVersionElectron+"/electron-v"+DefaultVersionElectron+"-linux-ia32.zip", p.ElectronDownloadSrc())
	p, err = newPaths("linux", "arm", o)
	assert.NoError(t, err)
	assert.Equal(t, "https://github.com/electron/electron/releases/download/v"+DefaultVersionElectron+"/electron-v"+DefaultVersionElectron+"-linux-armv7l.zip", p.ElectronDownloadSrc())
	p, err = newPaths("darwin", "", Options{BaseDirectoryPath: "/path/to/base/directory", AppIconDarwinPath: "/path/to/darwin/icon", AppIconDefaultPath: "icon", VersionAstilectron: DefaultVersionAstilectron, VersionElectron: DefaultVersionElectron})
	assert.NoError(t, err)
	assert.Equal(t, "/path/to/base/directory/vendor/electron-darwin-/Electron.app/Contents/MacOS/Electron", p.AppExecutable())
	assert.Equal(t, "/path/to/darwin/icon", p.AppIconDarwinSrc())
	assert.Equal(t, "/path/to/base/directory/icon", p.AppIconDefaultSrc())
	assert.Equal(t, "https://github.com/electron/electron/releases/download/v"+DefaultVersionElectron+"/electron-v"+DefaultVersionElectron+"-darwin-x64.zip", p.ElectronDownloadSrc())
	p, err = newPaths("darwin", "amd64", Options{AppName: "Test app", BaseDirectoryPath: "/path/to/base/directory", DataDirectoryPath: "/path/to/data/directory", VersionAstilectron: DefaultVersionAstilectron, VersionElectron: DefaultVersionElectron})
	assert.NoError(t, err)
	assert.Equal(t, "/path/to/data/directory", p.DataDirectory())
	assert.Equal(t, "/path/to/data/directory/vendor/electron-darwin-amd64/Test app.app/Contents/MacOS/Test app", p.AppExecutable())
	assert.Equal(t, "/path/to/data/directory/vendor/electron-darwin-amd64-v"+DefaultVersionElectron+".zip", p.ElectronDownloadDst())
	assert.Equal(t, "/path/to/data/directory/vendor/electron-darwin-amd64-v"+DefaultVersionElectron+".zip", p.ElectronUnzipSrc())
	const pad = "/path/to/appdata"
	os.Setenv(k, pad)
	p, err = newPaths("windows", "amd64", o)
	assert.NoError(t, err)
	assert.Equal(t, pad, p.DataDirectory())
	assert.Equal(t, pad+"/vendor", p.VendorDirectory())
	assert.Equal(t, pad+"/vendor/electron-windows-amd64/electron.exe", p.AppExecutable())
	assert.Equal(t, "https://github.com/electron/electron/releases/download/v"+DefaultVersionElectron+"/electron-v"+DefaultVersionElectron+"-win32-x64.zip", p.ElectronDownloadSrc())
	assert.Equal(t, pad+"/vendor/electron-windows-amd64-v"+DefaultVersionElectron+".zip", p.ElectronDownloadDst())
	assert.Equal(t, pad+"/vendor/electron-windows-amd64-v"+DefaultVersionElectron+".zip", p.ElectronUnzipSrc())
	p, err = newPaths("windows", "", o)
	assert.NoError(t, err)
	assert.Equal(t, "https://github.com/electron/electron/releases/download/v"+DefaultVersionElectron+"/electron-v"+DefaultVersionElectron+"-win32-ia32.zip", p.ElectronDownloadSrc())
	os.Setenv(k, ad)
}
