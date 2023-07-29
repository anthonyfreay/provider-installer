package util

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// Determine local Operating System, log and return the result.
func GetOS() string {
	os := runtime.GOOS
	log.Printf("Operating System is: %s", os)
	return os
}

// Determine local System Architecture, log and return the result.
func GetArch() string {
	arch := runtime.GOARCH
	log.Printf("System Architecture is: %s", arch)
	return arch
}

// Run `terraform verison` locally to determine the plugins directory location and structure.
func GetLocalTfVersion() string {
	// Execute local `terraform version` command
	cmd := exec.Command("terraform", "version")
	out, err := cmd.Output()
	if err != nil {
		log.Fatal("could not run command: ", err)
	}

	// Parse output of command, log the local version
	tfVersion := []rune(strings.Split(strings.ReplaceAll(string(out), "\n", " "), " ")[1])
	log.Println("Output: ", string(tfVersion))

	// Return local TF Version
	return string(tfVersion[1])
}

// Create the `plugins` directory on local depending on the `terraform version` and the system OS.
// Returns the path of the plugins directory.
func CreatePluginsDirectory(osType string, tfVersion string, arch string, pluginVersion string) string {

	var pluginDirPath string
	var tfdir string
	var homeDir string
	var err error

	homeDir, err = os.UserHomeDir()
	log.Printf("$HOME Directory is: %s", homeDir)

	log.Printf("osType is %s", osType)
	if osType == "darwin" {
		tfdir = ".terraform.d"
		if string(tfVersion) == "1" {
			pluginDirPath = homeDir + "/" + tfdir + "/plugins/terraform/abf/null/" + pluginVersion + "/" + osType + "_" + arch
		}
	} else if osType == "windows" {
		tfdir = "terraform.d"
		if string(tfVersion) == "1" {
			pluginDirPath = homeDir + "/" + tfdir + "/plugins/terraform/abf/null/" + pluginVersion + "/" + osType + "_" + arch
		}
	}
	log.Printf("Plugin Directory Path is set to: %s", pluginDirPath)
	err = os.MkdirAll(pluginDirPath, 0777)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	log.Printf("Plugin Directory has been created/updated")
	return pluginDirPath
}

// Unzip the source artifact and iterate over the contents
func UnzipSource(source, destination string) error {
	// Open the zip file
	reader, err := zip.OpenReader(source)
	if err != nil {
		return err
	}
	defer reader.Close()

	// Get the absolute destination path
	destination, err = filepath.Abs(destination)
	if err != nil {
		return err
	}

	// Iterate over zip files inside the archive and unzip each of them
	for _, f := range reader.File {
		err := unzipFile(f, destination)
		if err != nil {
			return err
		}
	}

	return nil
}

// Unzip individual files and check for Zip Slip vulnerabilites
func unzipFile(f *zip.File, destination string) error {
	// Check if file paths are not vulnerable to Zip Slip
	filePath := filepath.Join(destination, f.Name)
	if !strings.HasPrefix(filePath, filepath.Clean(destination)+string(os.PathSeparator)) {
		return fmt.Errorf("invalid file path: %s", filePath)
	}

	// Create directory tree
	if f.FileInfo().IsDir() {
		if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
			return err
		}
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}

	// Create a destination file for unzipped content
	destinationFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	// Unzip the content of a file and copy it to the destination file
	zippedFile, err := f.Open()
	if err != nil {
		return err
	}
	defer zippedFile.Close()

	if _, err := io.Copy(destinationFile, zippedFile); err != nil {
		return err
	}
	return nil
}
