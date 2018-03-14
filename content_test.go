package content

import (
	"testing"
	"os"
	"io/ioutil"
	"fmt"
	"path/filepath"
)

func TestGenerateContent(t *testing.T) {
	// write some temporary test files
	tfilesContent := []string{"Static content one", "Static content two"}
	tfilesName := []string{"staticOne", "staticTwo"}
	tFilesDir := "static-content"
	tOutputDir := "generated-content"
	packageName := "content"
	tOutputName := "static.go"

	tFilesDir, err := ioutil.TempDir("", tFilesDir)
	if err != nil {
		t.Errorf("error creating temp directory: %v", err)
	}
	tOutputDir, err = ioutil.TempDir("", tOutputDir)
	if err != nil {
		t.Errorf("error creating temp directory: %v", err)
	}

	for i, tfContent := range tfilesContent {
		tf, err := os.Create(filepath.Join(tFilesDir, tfilesName[i]))
		if err != nil {
			t.Errorf("error creating a test file: %v\n", err)
		}
		_, err = tf.WriteString(tfContent)
		if err != nil {
			t.Errorf("error writing a test file: %v\n", err)
		}
		err = tf.Close()
		if err != nil {
			t.Errorf("error closing a test file: %v\n", err)
		}
		tfilesName = append(tfilesName, tf.Name())
	}

	// generate the Go source file with the static content from the files

	//func GenerateContent(staticDir string, outputDir string, outputFile string, packageName string) error {

	err = GenerateContent(tFilesDir, tOutputDir, tOutputName, packageName)
	if err != nil {
		t.Errorf("error running GenerateContent: %v\n", err)
	}

	// Erase the input temp files

	err = os.RemoveAll(tFilesDir)
	if err != nil {
		t.Errorf("error removing the test input files: %v\n", err)
	}

	// Read and erase the generated file

	tgf, err := os.Open(filepath.Join(tOutputDir, tOutputName))
	if err != nil {
		t.Errorf("error opening generated file: %v\n", err)
	}
	tgfContent, err := ioutil.ReadAll(tgf)
	if err != nil {
		t.Errorf("error reading generated file: %v\n", err)
	}
	err = tgf.Close()
	if err != nil {
		t.Errorf("error closing generated file: %v\n", err)
	}
	err = os.RemoveAll(tOutputDir)
	if err != nil {
		t.Errorf("error removing generated file: %v\n", err)
	}

	// just print out the file until we check it automatically
	fmt.Print(string(tgfContent))
}
