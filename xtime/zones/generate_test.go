package zones

import (
	"archive/zip"
	"bytes"
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xstring"
	"github.com/Aoi-hosizora/ahlib/xtime"
	"go/format"
	"io"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"
)

func TestGenerate(t *testing.T) {
	_ = generate
	// err := generate()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
}

func generate() error {
	nameList, setList, err := getList()
	if err != nil {
		return fmt.Errorf("getList: %w", err)
	}

	// 1. code
	sb := bytes.Buffer{}
	sb.WriteString(`package zones

// Time zones are referred from "${runtime.GOROOT()}/lib/time/zoneinfo.zip".

const (
` + strings.Join(setList, "\n") + `
)`)
	err = formatAndWrite(sb.Bytes(), "_generate/zones.go")
	if err != nil {
		return fmt.Errorf("code formatAndWrite: %w", err)
	}

	// 2. test code
	sb.Reset()
	sb.WriteString(`package zones

import (
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"testing"
	"time"
)

func TestZones(t *testing.T) {
	for _, zone := range []string{
` + strings.Join(nameList, ",\n") + `,
	} {
		_, err := time.LoadLocation(zone)
		xtesting.Nil(t, err)
	}
}`)
	err = formatAndWrite(sb.Bytes(), "_generate/zones_test.go")
	if err != nil {
		return fmt.Errorf("test code formatAndWrite: %w", err)
	}

	return nil
}

func formatAndWrite(bs []byte, filename string) error {
	bs, err := format.Source(bs)
	if err != nil {
		return err
	}
	err = os.MkdirAll(filepath.Dir(filename), 0644)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(bs)
	if err != nil {
		return err
	}
	return nil
}

func unzip(srcFile string, dstDir string) error {
	archive, err := zip.OpenReader(srcFile)
	if err != nil {
		return err
	}
	defer archive.Close()

	for _, f := range archive.File {
		path := filepath.Join(dstDir, f.Name)
		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(path, 0644); err != nil {
				return err
			}
			continue
		}
		if err := os.MkdirAll(filepath.Dir(path), 0644); err != nil {
			return err
		}
		fileInArchive, err := f.Open()
		if err != nil {
			return err
		}
		dstFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			return err
		}
		_ = dstFile.Close()
		_ = fileInArchive.Close()
	}

	return nil
}

func getList() (varNameList []string, varSetList []string, err error) {
	zoneSource := filepath.Join(runtime.GOROOT(), "lib/time/zoneinfo.zip")
	tempDir := "./temp"
	err = unzip(zoneSource, tempDir)
	if err != nil {
		return nil, nil, fmt.Errorf("unzip: %w", err)
	}
	defer os.RemoveAll(tempDir)

	// read zoneinfo folder
	result := make([]string, 0)
	var readDir func(path string) error
	readDir = func(path string) error {
		files, err := os.ReadDir(filepath.Join(tempDir, path))
		if err != nil {
			return err
		}
		for _, entry := range files {
			name := entry.Name()
			if name != strings.ToUpper(name[:1])+name[1:] {
				continue
			}
			if entry.IsDir() {
				err = readDir(path + "/" + name)
				if err != nil {
					return err
				}
			} else {
				result = append(result, (path + "/" + name)[1:])
			}
		}
		return nil
	}
	err = readDir("")
	if err != nil {
		return nil, nil, fmt.Errorf("readDir: %w", err)
	}

	// generate final list
	varNameList = make([]string, 0, len(result)+2)
	varSetList = make([]string, 0, len(result)+2)
	varNameList = append(varNameList, "UTC", "Local")
	varSetList = append(varSetList, `UTC = "UTC" // time.UTC`, `Local = "Local" // time.Local`)
	for _, r := range result {
		loc, err := time.LoadLocation(r)
		if err != nil {
			return nil, nil, fmt.Errorf("LoadLocation: %w", err)
		}
		words := xstring.SplitToWords(loc.String(), "/")
		for i, w := range words {
			if strings.HasPrefix(w, "Etc") || strings.HasPrefix(w, "GMT") {
				// skip
			} else {
				w = xstring.PascalCase(w)
			}
			words[i] = w
		}
		varName := strings.Join(words, "_")

		du := xtime.LocationDuration(loc)
		flag := "+"
		if du < 0 {
			flag = "-"
		}
		comment := fmt.Sprintf("%s%02d:%02d", flag, int(math.Abs(du.Hours())), xtime.DurationMinuteComponent(du))

		if varName == "UTC" {
			continue
		} else if strings.HasPrefix(varName, "Etc") || strings.HasPrefix(varName, "GMT") {
			varName = strings.Replace(varName, "+", "_P", 1)
			varName = strings.Replace(varName, "-", "_M", 1)
		}
		varNameList = append(varNameList, varName)
		varSetList = append(varSetList, fmt.Sprintf("%s = \"%s\" // %s", varName, loc.String(), comment))
	}
	return varNameList, varSetList, nil
}
