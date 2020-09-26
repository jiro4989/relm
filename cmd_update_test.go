package main

import (
	"path/filepath"
	"testing"

	"github.com/google/go-github/v32/github"
	"github.com/stretchr/testify/assert"
)

func TestCmdUpdate(t *testing.T) {
	tests := []struct {
		desc    string
		app     App
		rel     Releases
		wantErr bool
	}{
		{
			desc: "ok: update",
			app: App{
				Config: Config{
					RelmaRoot: filepath.Join(testOutputDir, "test_cmd_update"),
				},
			},
			rel: Releases{
				{
					URL:           "https://github.com/jiro4989/nimjson/releases/download/v1.2.6/nimjson_linux.tar.gz",
					Owner:         "jiro4989",
					Repo:          "nimjson",
					Version:       "v1.2.6",
					AssetFileName: "nimjson_linux.tar.gz",
					InstalledFiles: InstalledFiles{
						{},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			// assert := assert.New(t)
			//
			// dir := tt.app.Config.RelmaRoot
			// err := os.MkdirAll(dir, os.ModePerm)
			// assert.NoError(err)
			//
			// f := tt.app.Config.ReleasesFile()
			// b, err := json.Marshal(tt.rel)
			// assert.NoError(err)
			// err = ioutil.WriteFile(f, b, os.ModePerm)
			// assert.NoError(err)
			//
			// err = tt.app.CmdUpdate(nil)
			// if tt.wantErr {
			// 	assert.Error(err)
			// 	return
			// }
			// assert.NoError(err)
		})
	}
}

func TestFetchLatestTag(t *testing.T) {
	client := github.NewClient(nil)
	tests := []struct {
		desc    string
		owner   string
		repo    string
		want    string
		wantErr bool
	}{
		{
			desc:    "ok: fetch latest tag",
			owner:   "jiro4989",
			repo:    "nimjson",
			want:    "v1.2.7",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			got, err := fetchLatestTag(client, tt.owner, tt.repo)
			if tt.wantErr {
				assert.Error(err)
				return
			}
			assert.NoError(err)
			assert.Equal(tt.want, got)
		})
	}
}
