package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"provider-installer/structs"
	"provider-installer/util"
)

func main() {
	desiredVersion := flag.String("version", "latest", "The version of the provider to install")
	providerName := flag.String("name", "null", "The name of the provider")
	providerSource := flag.String("source", "terraform/abf/null", "The source path of the provider (e.g. hashicorp/null or terraform/abf/null)")
	flag.Parse()

	log.Println("Determining OS")
	osType := util.GetOS()

	log.Println("Determining System Architecture")
	systemArch := util.GetArch()

	log.Println("Installing Provider")
	installProvider(*desiredVersion, *providerName, *providerSource, osType, systemArch)
}

func installProvider(desiredVersion string, name string, source string, osType string, arch string) {
	var resp *http.Response
	var err error
	var providerVersion string

	log.Printf("Fetch provider version: %s for %s (%s)", desiredVersion, name, source)
	if desiredVersion == "latest" {
		log.Printf("Determining the latest version of provider")
		// Note: This API endpoint is specific to Hashicorp releases. 
		// For third-party providers, this might need to change or be passed as a flag.
		// For now, keeping it as is but using the 'name' variable.
		apiURL := fmt.Sprintf("https://api.releases.hashicorp.com/v1/releases/terraform-provider-%s", name)
		resp, err = http.Get(apiURL)
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
		log.Printf("Current latest version is: %s", providerVersion)
		resp.Body.Close()
	} else {
		providerVersion = desiredVersion
		log.Printf("Using specified version: %s", providerVersion)
	}

	log.Println("Creating Terraform Plugins Directory")
	pluginDirPath := util.CreatePluginsDirectory(osType, util.GetLocalTfVersion(), util.GetArch(), source, providerVersion)

	// URL format: https://releases.hashicorp.com/terraform-provider-NAME/VERSION/terraform-provider-NAME_VERSION_OS_ARCH.zip
	url := fmt.Sprintf("https://releases.hashicorp.com/terraform-provider-%s/%s/terraform-provider-%s_%s_%s_%s.zip",
		name, providerVersion, name, providerVersion, osType, arch)
	
	log.Printf("Fetch provider version from: %s", url)
	resp, err = http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	executableName := fmt.Sprintf("terraform-provider-%s_v%s.zip", name, providerVersion)
	executableFilePath := filepath.Join(pluginDirPath, filepath.Base(executableName))

	log.Printf("Create local verison of executable at: %s", executableFilePath)
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

	log.Printf("Unzipping executable from: %s, to: %s", executableFilePath, pluginDirPath)
	err = util.UnzipSource(executableFilePath, pluginDirPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Remove source zip: %s", executableFilePath)
	os.Remove(executableFilePath)
}
