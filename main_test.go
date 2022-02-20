//go:build !windows
// +build !windows

package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainNormal(t *testing.T) {
	testDir := filepath.Join(testOutputDir, "test_main_1_1")
	testConfDir := filepath.Join(testDir, ".config", appName)

	tests := []struct {
		desc    string
		home    string
		confDir string
		args    []string
		wantErr bool
	}{
		{
			desc:    "normal: init",
			home:    testDir,
			confDir: testConfDir,
			args:    []string{"init"},
			wantErr: false,
		},
		{
			desc:    "normal: install",
			home:    testDir,
			confDir: testConfDir,
			args:    []string{"install", "https://github.com/jiro4989/nimjson/releases/download/v1.2.6/nimjson_linux.tar.gz"},
			wantErr: false,
		},
		{
			desc:    "normal: list",
			home:    testDir,
			confDir: testConfDir,
			args:    []string{"list"},
			wantErr: false,
		},
		{
			desc:    "normal: root",
			home:    testDir,
			confDir: testConfDir,
			args:    []string{"root"},
			wantErr: false,
		},
		{
			desc:    "normal: update",
			home:    testDir,
			confDir: testConfDir,
			args:    []string{"update"},
			wantErr: false,
		},
		{
			desc:    "normal: upgrade",
			home:    testDir,
			confDir: testConfDir,
			args:    []string{"upgrade", "-y"},
			wantErr: false,
		},
		{
			desc:    "normal: uninstall",
			home:    testDir,
			confDir: testConfDir,
			args:    []string{"uninstall", "jiro4989/nimjson"},
			wantErr: false,
		},
		{
			desc:    "abnormal: hogefuga command doesn't exist",
			home:    filepath.Join(testOutputDir, "test_main_1_2"),
			confDir: filepath.Join(testOutputDir, "test_main_1_2", ".config", appName),
			args:    []string{"hogefuga"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			assert.NoError(os.MkdirAll(tt.home, os.ModePerm))
			recoverFunc := SetHomeWithRecoverFunc(tt.home)
			defer recoverFunc()
			assert.NoError(os.MkdirAll(tt.confDir, os.ModePerm))

			err := Main(tt.args)
			if tt.wantErr {
				assert.Error(err)
				return
			}
			assert.NoError(err)
		})
	}
}
