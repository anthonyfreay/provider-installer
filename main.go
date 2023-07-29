package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"provider-installer/structs"
	"provider-installer/util"
)

func main() {
	log.Println("Determining OS")
	osType := util.GetOS()

	log.Println("Determining System Architecture")
	systemArch := util.GetArch()

	log.Println("Installing Provider")
	installProvider("latest", osType, systemArch)
}

func installProvider(desiredVersion string, osType string, arch string) {
	var resp *http.Response
	var err error
	var providerVersion string

	log.Println("Place Proper Provider in directory with new name")

	if desiredVersion == "latest" {
		resp, err = http.Get("https://api.releases.hashicorp.com/v1/releases/terraform-provider-null")
		if err != nil {
			log.Fatal(err)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		var result structs.ProviderVersions
		if err := json.Unmarshal(body, &result); err != nil {
			log.Fatal("Can not unmarshal JSON")
		}
		providerVersion = result[0].Version
		log.Printf("Current latest verison is: %s", providerVersion)
		resp.Body.Close()
	} else {
		// GET SPECIFIC VERSION
	}

	log.Println("Create Proper Terraform Plugins Directory")
	pluginDirPath := util.CreatePluginsDirectory(osType, util.GetLocalTfVersion(), util.GetArch(), providerVersion)

	url := `https://releases.hashicorp.com/terraform-provider-null/` + providerVersion + `/terraform-provider-null_` + providerVersion + `_` + osType + `_` + arch + `.zip`
	resp, err = http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	executableName := `terraform-provider-null_v` + providerVersion + `.zip`
	executableFilePath := filepath.Join(pluginDirPath, filepath.Base(executableName))
	out, err := os.Create(executableFilePath)
	os.Chmod(executableFilePath, 0777)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()

	err = util.UnzipSource(executableFilePath, pluginDirPath)
	if err != nil {
		log.Fatal(err)
	}

	os.Remove(executableFilePath)
}
