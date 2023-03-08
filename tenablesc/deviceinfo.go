package tenablesc

import (
	"fmt"
	"net/url"
)

const deviceInfoEndpoint = "/deviceInfo"

// DeviceInfo represents response structure for https://docs.tenable.com/tenablesc/api/Device-Information.htm

type DeviceInfoContainer struct {
	DevType          string               `json:"type,omitempty"`
	UUID             string               `json:"uuid,omitempty"`
	Score            string               `json:"score,omitempty"`
	Total            string               `json:"total,omitempty"`
	SeverityInfo     string               `json:"severityInfo,omitempty"`
	SeverityLow      string               `json:"severityLow,omitempty"`
	SeverityMedium   string               `json:"severityMedium,omitempty"`
	SeverityHigh     string               `json:"severityHigh,omitempty"`
	SeverityCritical string               `json:"severityCritical,omitempty"`
	MacAddress       string               `json:"macAddress,omitempty"`
	PolicyName       string               `json:"policyName,omitempty"`
	PluginSet        string               `json:"pluginSet,omitempty"`
	NetbiosName      string               `json:"netbiosName,omitempty"`
	DnsName          string               `json:"dnsName,omitempty"`
	OsCPE            string               `json:"osCPE,omitempty"`
	BiosGUID         string               `json:"biosGUID,omitempty"`
	TpmID            string               `json:"tpmID,omitempty"`
	McafeeGUID       string               `json:"mcafeeGUID,omitempty"`
	LastAuthRun      string               `json:"lastAuthRun,omitempty"`
	LastUnauthRun    string               `json:"lastUnauthRun,omitempty"`
	SeverityAll      string               `json:"severityAll,omitempty"`
	Os               string               `json:"os,omitempty"`
	HasPassive       string               `json:"hasPassive,omitempty"`
	HasCompliance    string               `json:"hasCompliance,omitempty"`
	LastScan         string               `json:"lastScan,omitempty"`
	Links            []DeviceInfoLinks    `json:"links,omitempty"`
	Repository       DeviceInfoRepository `json:"repository,omitempty"`
}

type DeviceInfoLinks struct {
	Name string `json:"name,omitempty"`
	Link string `json:"link,omitempty"`
}

type DeviceInfoRepository struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (c *Client) GetDeviceInfoByMAC(mac string) ([]*DeviceInfoContainer, error) {
	var deviceInfo []*DeviceInfoContainer

	query := url.Values{}
	query.Add("macAddress", mac)

	if _, err := c.getResource(fmt.Sprintf("%s?%s", deviceInfoEndpoint, query.Encode()), &deviceInfo); err != nil {
		return nil, fmt.Errorf("failed to find device with MAC address %s: %w", mac, err)
	}

	return deviceInfo, nil
}

func (c *Client) GetDeviceInfoByDNS(hostname string, domain string) ([]*DeviceInfoContainer, error) {
	var deviceInfo []*DeviceInfoContainer
	fqdn := hostname + domain
	query := url.Values{}
	query.Add("macAddress", fqdn)

	if _, err := c.getResource(fmt.Sprintf("%s?%s", deviceInfoEndpoint, query.Encode()), &deviceInfo); err != nil {
		return nil, fmt.Errorf("failed to find device with DNS name %s: %w", fqdn, err)
	}

	return deviceInfo, nil
}

func (c *Client) GetDeviceInfoByIP(ip string) ([]*DeviceInfoContainer, error) {
	var deviceInfo []*DeviceInfoContainer
	query := url.Values{}
	query.Add("ip", ip)

	if _, err := c.getResource(fmt.Sprintf("%s?%s", deviceInfoEndpoint, query.Encode()), &deviceInfo); err != nil {
		return nil, fmt.Errorf("failed to find device with IP %s: %w", ip, err)
	}

	return deviceInfo, nil
}
