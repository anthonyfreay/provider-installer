package structs

import "time"

type ProviderVersions []struct {
	Builds []struct {
		Arch string `json:"arch"`
		Os   string `json:"os"`
		URL  string `json:"url"`
	} `json:"builds"`
	IsPrerelease bool   `json:"is_prerelease"`
	LicenseClass string `json:"license_class"`
	Name         string `json:"name"`
	Status       struct {
		State            string    `json:"state"`
		TimestampUpdated time.Time `json:"timestamp_updated"`
	} `json:"status"`
	TimestampCreated     time.Time `json:"timestamp_created"`
	TimestampUpdated     time.Time `json:"timestamp_updated"`
	URLChangelog         string    `json:"url_changelog,omitempty"`
	URLLicense           string    `json:"url_license,omitempty"`
	URLShasums           string    `json:"url_shasums"`
	URLShasumsSignatures []string  `json:"url_shasums_signatures"`
	URLSourceRepository  string    `json:"url_source_repository"`
	Version              string    `json:"version"`
}
