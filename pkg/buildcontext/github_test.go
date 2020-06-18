/*
Copyright 2020 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package buildcontext

//import (
//	"compress/gzip"
//	"github.com/GoogleContainerTools/kaniko/pkg/config"
//	"github.com/GoogleContainerTools/kaniko/pkg/util"
//	"io/ioutil"
//	"net/http"
//	"net/http/httptest"
//	"os"
//	"path/filepath"
//	"sort"
//	"strings"
//	"testing"
//
//	"github.com/GoogleContainerTools/kaniko/testutil"
//)
//
//func writeCompressedTar(fileContents map[string]string, w http.ResponseWriter, testDir string) error {
//	if err := testutil.SetupFiles(testDir, fileContents); err != nil {
//		return err
//	}
//	gw := gzip.NewWriter(w)
//	defer func () {
//		gw.Flush()
//		gw.Close()
//	}()
//	t := util.NewTar(gw)
//	defer t.Close()
//	for file := range fileContents {
//		filePath := filepath.Join(testDir, file)
//		if err := t.AddFileToTar(filePath); err != nil {
//			return err
//		}
//	}
//	return nil
//}
//
//func TestUnpackTarFromHost(t *testing.T) {
//	tests := []struct {
//		name             string
//		setupTarContents map[string]string
//		destination      string
//		prefix      	 bool
//		expectedFileList []string
//		errorExpected    bool
//	}{
//		// Prefix = 0 covers util.unTar
//		{
//			name: "multfile tar",
//			setupTarContents: map[string]string{
//				"prefix/foo/file1": "hello World",
//				"prefix/bar/file1": "hello World",
//				"prefix/bar/file2": "hello World",
//				"prefix/file1":     "hello World",
//			},
//			destination:      "/",
//			prefix:      	  false,
//			expectedFileList: []string{"foo/file1", "bar/file1", "bar/file2", "file1"},
//			errorExpected:    false,
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			testDir, err := ioutil.TempDir("", "")
//			if err != nil {
//				t.Fatal(err)
//			}
//			defer os.RemoveAll(testDir)
//
//			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//				config.RootDir = testDir
//				if err := writeCompressedTar(tt.setupTarContents, w, testDir); err != nil {
//					t.Fatal(err)
//				}
//			}))
//			defer ts.Close()
//
//			srcContextDir, err := ioutil.TempDir("", "")
//			if err != nil {
//				t.Fatal(err)
//			}
//			defer os.RemoveAll(srcContextDir)
//
//			fileList, err := downloadAndUnpackTar("https://codeload.github.com/ryan-gerstenkorn-sp/cbuild/legacy.tar.gz/master?token=AOG46UGLXU4COYYH7I2DFMK6545CY", srcContextDir)
//			if !tt.errorExpected && err != nil {
//				t.Fatal(err)
//			}
//			for i, file := range fileList {
//				fileList[i] = strings.TrimPrefix(file, srcContextDir + "/")
//			}
//
//			// sort both slices to ensure objects are in the same order for deep equals
//			sort.Strings(tt.expectedFileList)
//			sort.Strings(fileList)
//			testutil.CheckErrorAndDeepEqual(t, tt.errorExpected, err, tt.expectedFileList, fileList)
//		})
//	}
//
//}
//
