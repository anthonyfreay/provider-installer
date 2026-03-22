# provider-installer

<!-- PROJECT SHIELDS -->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]

Golang tool to download a Terraform Provider to a local location.

<!-- TABLE OF CONTENTS -->

## Table of Contents

- [About The Project](#about-the-project)
  - [Features](#features)
  - [Built With](#built-with)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

<!-- ABOUT THE PROJECT -->

## About The Project

I haven't touched native golang in a while and was missing it.

I've been heavily interacting with Terraform and different providers, including third-party providers not included in the registry.

This project serves as an exercise to template a tool to easily onboard a new developer to a third-party provider. **Note:** In its current state, the tool is hardcoded to install the [HashiCorp Terraform Null Provider](https://registry.terraform.io/providers/hashicorp/null/latest) as a demonstration.

### Features

- **Automatic OS/Arch Detection:** Detects `darwin`, `linux`, or `windows` and the system architecture (`amd64`, `arm64`) automatically.
- **Terraform Version Awareness:** Checks the local Terraform version to determine the correct plugin directory structure (supports Terraform 1.x).
- **Dynamic Provider Configuration:** Support for custom provider names and source paths via CLI flags.
- **Automated Installation:** Fetches, downloads, unzips, and places the provider in the local filesystem mirror directory.

[back to top](#provider-installer)

### Built With

- [![Golang][Golang]][Golang-url]

[back to top](#provider-installer)

<!-- GETTING STARTED -->

## Getting Started

To get a local clone up and running, take a look at the following steps...

### Prerequisites

#### Golang

This was built using `go version go1.20.6 darwin/arm64` but should be able to be built using any version of golang.

Install golang using your preferred package-manager e.g.

```sh
brew install go
```

#### Terraform

The tool requires `terraform` to be installed and available in your PATH to determine the correct plugin directory structure.

```sh
brew install terraform
```

### Installation

1. Clone the repo

   ```sh
   git clone https://github.com/anthonyfreay/provider-installer.git
   ```

2. Run immediately out-of-box

   ```sh
   go run main.go
   ```

[back to top](#provider-installer)

<!-- USAGE EXAMPLES -->

## Usage

1. Build the executable of the provider-installer and place it in the target directory.

   ```bash
   # location: project root
   go build -o ./target provider-installer
   ```

2. Run `provider-installer`

   ```bash
   # Install the latest version (defaults to null provider)
   ./target/provider-installer

   # Install a specific version
   ./target/provider-installer -version 3.2.1

   # Install a custom provider
   ./target/provider-installer -name my-provider -source my-org/my-provider -version 1.0.0
   ```

3. Make sure provider requirement snippet is defined within your terraform configs

   ```terraform
   terraform {
     required_providers {
       my-provider = {
         source = "my-org/my-provider"
       }
     }
   }
   ```

4. How it works

   The tool performs the following steps:

   - Detects the local OS and Architecture using `runtime.GOOS` and `runtime.GOARCH`.
   - Runs `terraform version` to determine the major version of Terraform installed.
   - Fetches the version metadata (defaults to `null` provider from Hashicorp's release API if `latest` is requested).
   - Creates the local plugin directory using the standard filesystem mirror layout:
     - Linux/Darwin: `~/.terraform.d/plugins/SOURCE/VERSION/OS_ARCH`
     - Windows: `%APPDATA%\terraform.d\plugins\SOURCE\VERSION\OS_ARCH`
   - Downloads and unzips the provider binary into that directory.

   > [!NOTE]
   > While the name and source are dynamic, the download URL is currently optimized for HashiCorp-style release mirrors (e.g., `releases.hashicorp.com`).

5. Test via Terraform Configuration

   ```bash
   # location: terraformTesting
   terraform init
   ```

[back to top](#provider-installer)

<!-- ROADMAP -->

## Roadmap

- [x] Add support for specific version installation
- [x] Support for dynamic provider name/source via CLI flags
- [ ] Support for custom release API endpoints (non-HashiCorp)

See the [open issues](https://github.com/anthonyfreay/provider-installer/issues) for a full list of proposed features (and known issues).

[back to top](#provider-installer)

<!-- CONTRIBUTING -->

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

[back to top](#provider-installer)

<!-- LICENSE -->

## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

[back to top](#provider-installer)

<!-- CONTACT -->

## Contact

Anthony Freay - [@anthonyfreay](https://www.linkedin.com/in/anthonyfreay/) - [anthonyfreay.com](https://anthonyfreay.com)

Project Link: [https://github.com/anthonyfreay/provider-installer](https://github.com/anthonyfreay/provider-installer)

[back to top](#provider-installer)

<!-- MARKDOWN LINKS & IMAGES -->

[contributors-shield]: https://img.shields.io/github/contributors/anthonyfreay/provider-installer
[contributors-url]: https://github.com/anthonyfreay/provider-installer/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/anthonyfreay/provider-installer
[forks-url]: https://github.com/anthonyfreay/provider-installer/network/members
[stars-shield]: https://img.shields.io/github/stars/anthonyfreay/provider-installer
[stars-url]: https://github.com/anthonyfreay/provider-installer/stargazers
[issues-shield]: https://img.shields.io/github/issues/anthonyfreay/provider-installer
[issues-url]: https://github.com/anthonyfreay/provider-installer/issues
[license-shield]: https://img.shields.io/github/license/anthonyfreay/provider-installer
[license-url]: https://github.com/anthonyfreay/provider-installer/blob/master/LICENSE.txt
[Golang]: https://img.shields.io/badge/golang-000000&logo=nextdotjs&logoColor=white
[Golang-url]: https://go.dev/
