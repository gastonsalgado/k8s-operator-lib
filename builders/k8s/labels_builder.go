package k8sbuilders

type Labels struct {
	Name      string
	Instance  string
	Version   string
	Component string
	PartOf    string
	ManagedBy string
}

func getCommonLabels(labels Labels) map[string]string {
	return map[string]string{
		"app.kubernetes.io/name":       labels.Name,
		"app.kubernetes.io/instance":   labels.Instance,
		"app.kubernetes.io/version":    labels.Version,
		"app.kubernetes.io/component":  labels.Component,
		"app.kubernetes.io/part-of":    labels.PartOf,
		"app.kubernetes.io/managed-by": labels.ManagedBy,
	}
}
