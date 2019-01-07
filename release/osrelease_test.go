package release

import (
	"testing"
)

func TestParseOSRelease(t *testing.T) {
	testData := []string{"NAME=Fedora",
		"VERSION=\"29 (Twenty Nine)\"",
		"ID=fedora",
		"ID_LIKE=linux",
		"VERSION_ID=29",
		"VERSION_CODENAME=\"\"",
		"PLATFORM_ID=\"platform:f29\"",
		"PRETTY_NAME=\"Fedora 29 (Twenty Nine)\"",
		"ANSI_COLOR=\"0;34\"",
		"LOGO=fedora-logo-icon",
		"BUILD_ID=build",
		"VARIANT=var",
		"VARIANT_ID=varid",
		"CPE_NAME=\"cpe:/o:fedoraproject:fedora:29\"",
		"HOME_URL=\"https://fedoraproject.org/\"",
		"DOCUMENTATION_URL=\"https://docs.fedoraproject.org/en-US/fedora/f29/system-administrators-guide/\"",
		"SUPPORT_URL=\"https://fedoraproject.org/wiki/Communicating_and_getting_help\"",
		"BUG_REPORT_URL=\"https://bugzilla.redhat.com/\"",
		"REDHAT_BUGZILLA_PRODUCT=\"Fedora\"",
		"REDHAT_BUGZILLA_PRODUCT_VERSION=29",
		"REDHAT_SUPPORT_PRODUCT=\"Fedora\"",
		"REDHAT_SUPPORT_PRODUCT_VERSION=29",
		"PRIVACY_POLICY_URL=\"https://fedoraproject.org/wiki/Legal:PrivacyPolicy\""}
	osRelease := ParseOSRelease(testData)
	if osRelease.Name != "Fedora" {
		t.Errorf("Received value was incorrect, got: %s, want: %s", osRelease.Name, "Fedora")
	}
	if osRelease.ID != "fedora" {
		t.Errorf("Received value was incorrect, got: %s, want: %s", osRelease.ID, "fedora")
	}
	if osRelease.IDLike != "linux" {
		t.Errorf("Received value was incorrect, got: %s, want: %s", osRelease.IDLike, "linux")
	}
	if osRelease.VersionID != "29" {
		t.Errorf("Received value was incorrect, got: %s, want: %s", osRelease.VersionID, "29")
	}
	if osRelease.VersionCodename != "\"\"" {
		t.Errorf("Received value was incorrect, got: %s, want: %s", osRelease.VersionCodename, "\"\"")
	}
	if osRelease.PrettyName != "\"Fedora 29 (Twenty Nine)\"" {
		t.Errorf("Received value was incorrect, got: %s, want: %s", osRelease.PrettyName, "\"Fedora 29 (Twenty Nine)\"")
	}
	if osRelease.ANSIColor != "\"0;34\"" {
		t.Errorf("Received value was incorrect, got: %s, want: %s", osRelease.ANSIColor, "\"0;34\"")
	}
	if osRelease.CPEName != "\"cpe:/o:fedoraproject:fedora:29\"" {
		t.Errorf("Received value was incorrect, got: %s, want: %s", osRelease.CPEName, "\"cpe:/o:fedoraproject:fedora:29\"")
	}
	if osRelease.HomeURL != "\"https://fedoraproject.org/\"" {
		t.Errorf("Received value was incorrect, got: %s, want: %s", osRelease.HomeURL, "\"https://fedoraproject.org/\"")
	}
	if osRelease.DocumentationURL != "\"https://docs.fedoraproject.org/en-US/fedora/f29/system-administrators-guide/\"" {
		t.Errorf("Received value was incorrect, got: %s, want: %s", osRelease.DocumentationURL,
			"\"https://docs.fedoraproject.org/en-US/fedora/f29/system-administrators-guide/\"")
	}
	if osRelease.SupportURL != "\"https://fedoraproject.org/wiki/Communicating_and_getting_help\"" {
		t.Errorf("Received value was incorrect, got: %s, want: %s", osRelease.SupportURL,
			"\"https://fedoraproject.org/wiki/Communicating_and_getting_help\"")
	}
	if osRelease.BugReportURL != "\"https://bugzilla.redhat.com/\"" {
		t.Errorf("Received value was incorrect, got: %s, want: %s", osRelease.BugReportURL,
			"\"https://bugzilla.redhat.com/\"")
	}
	if osRelease.PrivacyPolicyURL != "\"https://fedoraproject.org/wiki/Legal:PrivacyPolicy\"" {
		t.Errorf("Received value was incorrect, got: %s, want: %s", osRelease.PrivacyPolicyURL,
			"\"https://fedoraproject.org/wiki/Legal:PrivacyPolicy\"")
	}
	if osRelease.BuildID != "build" {
		t.Errorf("Received value was incorrect, got: %s, want: %s", osRelease.BuildID, "build")
	}
	if osRelease.Variant != "var" {
		t.Errorf("Received value was incorrect, got: %s, want: %s", osRelease.Variant, "var")
	}
	if osRelease.VariantID != "varid" {
		t.Errorf("Received value was incorrect, got: %s, want: %s", osRelease.VariantID, "varid")
	}
	if osRelease.Logo != "fedora-logo-icon" {
		t.Errorf("Received value was incorrect, got: %s, want: %s", osRelease.Logo, "fedora-logo-icon")
	}
}
