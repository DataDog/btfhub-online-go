package btfhubonline

// BTFRecordIdentifier holds the 4 identifiers of a specific BTF file.
type BTFRecordIdentifier struct {
	Distribution        string `json:"distribution"`
	DistributionVersion string `json:"distribution_version"`
	KernelVersion       string `json:"kernel_version"`
	Arch                string `json:"arch"`
}

func (identifier BTFRecordIdentifier) AsMap() map[string]string {
	return map[string]string{
		"distribution":         identifier.Distribution,
		"distribution_version": identifier.DistributionVersion,
		"kernel_version":       identifier.KernelVersion,
		"arch":                 identifier.Arch,
	}
}
