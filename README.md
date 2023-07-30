<a name="readme-top"></a>

<!-- PROJECT SHIELDS -->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]

<!-- PROJECT TITLE -->
<div>
  <h3 align="center">provider-installer</h3>

  <p align="center">Golang tool to download a Terraform Provider to a local location.</p>
</div>


<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#pre-requisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project
I haven't touched naitive golang in a while and was missing it. 

I've been heavily interacting with Terraform and different providers, including third-party providers not included in the registry.

This project serves as an exercise to template a tool to easily onboard a new developer to a third-party provider. This installation tool using the Hasicorp Terraform Null Provider as the example of a third-party provider.


<p align="right">(<a href="#readme-top">back to top</a>)</p>


### Built With

* [![Golang][Golang]][Golang-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- GETTING STARTED -->
## Getting Started

To get a local clone up and running, take a look at the following steps...

### Pre-Requisites

#### golang
This was built using `go version go1.20.6 darwin/arm64` but should be able to be built using any verison of golang.

Install golang using your preferred package-manager e.g.
  ```sh
  brew install go
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

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

1. Build the executable of the provider-installer and place it in the target directory.
````bash
# location: project root
go build -o ./target provider-installer
````
2. Run `provider-installer`
````bash
./target/provider-installer
````
3. Make sure provider requirement snipper is defined within your terraform configs
````terraform
terraform {
  required_providers {
    null = {
      source = "terraform/abf/null"
    }
  }
}
````

4. Test via Terraform Configuration
````bash
# location: terraformTesting
terraform init
````

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

- [ ] Add support for specific version installation

See the [open issues](https://github.com/anthonyfreay/provider-installer/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>



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

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Anthony Freay - [@anthonyfreay](https://www.linkedin.com/in/anthonyfreay/) - [anthonyfreay.com](https://anthonyfreay.com)

Project Link: [https://github.com/anthonyfreay/provider-installer](https://github.com/anthonyfreay/provider-installer)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
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
[linkedin-url]: https://linkedin.com/in/anthonyfreay
[Golang]: https://img.shields.io/badge/golang-000000&logo=nextdotjs&logoColor=white
[Golang-url]: https://go.dev/
