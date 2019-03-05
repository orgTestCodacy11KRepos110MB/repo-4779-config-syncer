package v1beta1

import (
	"hash/fnv"
	"strconv"

	"github.com/appscode/stash/apis"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	hashutil "k8s.io/kubernetes/pkg/util/hash"
	crdutils "kmodules.xyz/client-go/apiextensions/v1beta1"
)

func (bs BackupSession) GetSpecHash() string {
	hash := fnv.New64a()
	hashutil.DeepHashObject(hash, bs.Spec)
	return strconv.FormatUint(hash.Sum64(), 10)
}

func (bs BackupSession) CustomResourceDefinition() *apiextensions.CustomResourceDefinition {
	return crdutils.NewCustomResourceDefinition(crdutils.Config{
		Group:         SchemeGroupVersion.Group,
		Plural:        ResourcePluralBackupSession,
		Singular:      ResourceSingularBackupSession,
		Kind:          ResourceKindBackupSession,
		ShortNames:    []string{"bs"},
		Categories:    []string{"stash", "backup", "appscode"},
		ResourceScope: string(apiextensions.NamespaceScoped),
		Versions: []apiextensions.CustomResourceDefinitionVersion{
			{
				Name:    SchemeGroupVersion.Version,
				Served:  true,
				Storage: true,
			},
		},
		Labels: crdutils.Labels{
			LabelsMap: map[string]string{"app": "stash"},
		},
		SpecDefinitionName:      "github.com/appscode/stash/apis/stash/v1beta1.BackupSession",
		EnableValidation:        true,
		GetOpenAPIDefinitions:   GetOpenAPIDefinitions,
		EnableStatusSubresource: apis.EnableStatusSubresource,
		AdditionalPrinterColumns: []apiextensions.CustomResourceColumnDefinition{
			{
				Name:     "BackupConfiguration",
				Type:     "string",
				JSONPath: ".spec.backupConfiguration.name",
			},
			{
				Name:     "Phase",
				Type:     "string",
				JSONPath: ".status.phase",
			},
			{
				Name:     "Age",
				Type:     "date",
				JSONPath: ".metadata.creationTimestamp",
			},
		},
	})
}
