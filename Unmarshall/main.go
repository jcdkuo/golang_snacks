package main

import (
	"encoding/json"
	"fmt"
)

// https://github.com/composer/composer/blob/master/res/composer-schema.json
// http://getcomposer.org/doc/04-schema.md#repositories
// IGNORE THIS -- MIGHT NOT BE USED
type Repository struct {
	Type    string
	Url     string
	Package struct {
		Name    string
		Version string
		Dist    struct {
			Url  string
			Type string
		}
		Source struct {
			Url       string
			Type      string
			Reference string
		}
	}

	Options struct {
		Ssl struct {
			VerifyPeer bool
		}
	}
}

type Config struct {
	ProcessTimeout    int    `json:"process-timeout"`
	UseIncludePath    bool   `json:"use-include-path"`
	PreferredInstall  string `json:"preferred-install"`
	NotifyOnInstall   bool   `json:"notify-on-install"`
	GithubProtocols   []string
	GithubOauth       map[string]string `json:"github-oauth"`
	GithubDomains     []string          `json:"github-domains"`
	VendorDir         string            `json:"vendor-dir"`
	BinDir            string            `json:"bin-dir"`
	CacheDir          string            `json:"cache-dir"`
	CacheFilesDir     string            `json:"cache-files-dir"`
	CacheRepoDir      string            `json:"cache-repo-dir"`
	CacheVcsDir       string            `json:"cache-vcs-dir"`
	CacheTtl          int               `json:"cache-ttl"`
	CacheFilesTtl     int               `json:"cache-files-ttl"`
	CacheFilesMaxsize string            `json:"cache-files-maxsize"` // Could be string or int -- Default "300MiB"
	DiscardChanges    bool              `json:"discard-changes"`     // This could be true, false or "stash"
	AutoloaderSuffix  string            `json:"autoloader-suffix"`
	PrependAutoloader bool              `json:"prepend-autoloader"`
}

// Config Initializer
func NewConfig() *Config {
	return &Config{
		ProcessTimeout:    300,
		GithubProtocols:   []string{"git", "https", "http"},
		GithubDomains:     []string{"github.com"},
		VendorDir:         "vendor",
		BinDir:            "vendor/bin",
		CacheDir:          "~/.composer/cache",
		CacheFilesMaxsize: "300MiB",
		NotifyOnInstall:   true,
		PreferredInstall:  "auto",
		PrependAutoloader: true,
		DiscardChanges:    false,
		CacheTtl:          15552000,
		CacheFilesTtl:     15552000,
	}
}

type Author struct {
	Name     string
	Email    string
	Homepage string
	Role     string
}

type AutoloadValue struct {
	Psr0     map[string]string `json:"psr-0"`
	Psr4     map[string]string `json:"psr-4"`
	Classmap []string
	files    []string
}

type Script struct {
	PreInstallCmd          []string `json:"pre-install-cmd"`
	PostInstallCmd         []string `json:"post-install-cmd"`
	PreUpdateCmd           []string `json:"pre-update-cmd"`
	PostUpdateCmd          []string `json:"post-update-cmd"`
	PreStatusCmd           []string `json:"pre-status-cmd"`
	PostStatusCmd          []string `json:"post-status-cmd"`
	PrePackageInstall      []string `json:"pre-package-install"`
	PostPackageInstall     []string `json:"post-package-install"`
	PrePackageUpdate       []string `json:"pre-package-update"`
	PostPackageUpdate      []string `json:"post-package-update"`
	PrePackageUninstall    []string `json:"pre-package-uninstall"`
	PostPackageUninstall   []string `json:"post-package-uninstall"`
	PreAutoloadDump        []string `json:"pre-autoload-dump"`
	PostAutoloadDump       []string `json:"post-autoload-dump"`
	PostRootPackageInstall []string `json:"post-root-package-install"`
	PostCreateProjectCmd   []string `json:"post-create-project-cmd"`
}

type Support struct {
	Email  string
	Issues string
	Forum  string
	Wiki   string
	Irc    string
	Source string
}

// Composer.json top level
type Manifest struct {
	Name             string
	Type             string
	Description      string
	Keywords         []string
	Homepage         string
	Version          string
	Time             string
	License          []string
	Authors          []Author
	Require          map[string]string
	Replace          map[string]string
	Conflict         map[string]string
	Provide          map[string]string
	RequireDev       map[string]string `json:"require-dev"`
	Suggest          map[string]string
	Config           Config
	Extra            map[string]string
	Autoload         AutoloadValue
	Archive          []string
	Repositories     map[string]string // -- Don't know if this needs to be a map[string]string or a Repository type
	MinimumStability string            `json:"minimum-stability"`
	PreferStable     bool              `json:"prefer-stable"`
	Bin              []string
	IncludePath      []string `json:"include-path"` // This is apparently deprecated
	Scripts          []Script
	Support          Support
}

func NewManifest() *Manifest {
	conf := *NewConfig()
	return &Manifest{Config: conf}
}

func main() {
	// This needs more added to it
	test := `{ 
        "name": "TestPackage", 
        "type": "library", 
        "description": "The foo bar lorem ipsum sit dolor.", 
        "keywords": [ "foo", "bar", "baz" ], 
        "homepage": "http://www.example.com", 
        "version": "1.1.1", 
        "time": "2013/01/30", 
        "license": ["MIT"], 
        "Authors": [{ "name": "Jesse O'Brien", "email": "jesse@jesse-obrien.ca", "homepage": "http://jesse-obrien.ca", "role": "Developer" }],
        "require": {
            "foo/bar": "1.0",
            "foo/baz": "2.0"
        }
    }`

	manifest := *NewManifest()
	json.Unmarshal([]byte(test), &manifest)
	fmt.Println(manifest)
	fmt.Println(manifest.Name)
}
