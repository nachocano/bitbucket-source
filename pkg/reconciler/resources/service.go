package resources

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	servingv1alpha1 "github.com/knative/serving/pkg/apis/serving/v1alpha1"
	sourcesv1alpha1 "github.com/nachocano/bitbucket-source/pkg/apis/sources/v1alpha1"
)

// MakeService generates, but does not create, a Service for the given BitBucketSource.
func MakeService(source *sourcesv1alpha1.BitBucketSource, receiveAdapterImage string) *servingv1alpha1.Service {
	labels := map[string]string{
		"receive-adapter": "bitbucket",
	}
	sinkURI := source.Status.SinkURI
	env := []corev1.EnvVar{
		// TODO should also pass the UUID of the webhook so that the receive adapter can validate that the
		// incoming events correspond to that particular webhook, and discard them otherwise.
		// There is a chicken and egg problem: in order to create the webhook, we need the service domain, but
		// the service container image needs the webhook UUID that we get when the webhook is created.
		// We need to properly populate this in a reconcile loop.
		{
			Name:  "BITBUCKET_UUID",
			Value: source.Status.WebhookUUIDKey,
		},
		{
			Name:  "SINK",
			Value: sinkURI,
		},
	}
	containerArgs := []string{fmt.Sprintf("--sink=%s", sinkURI)}
	return &servingv1alpha1.Service{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: fmt.Sprintf("%s-", source.Name),
			Namespace:    source.Namespace,
			Labels:       labels,
		},
		Spec: servingv1alpha1.ServiceSpec{
			RunLatest: &servingv1alpha1.RunLatestType{
				Configuration: servingv1alpha1.ConfigurationSpec{
					RevisionTemplate: servingv1alpha1.RevisionTemplateSpec{
						Spec: servingv1alpha1.RevisionSpec{
							ServiceAccountName: source.Spec.ServiceAccountName,
							Container: corev1.Container{
								Image: receiveAdapterImage,
								Env:   env,
								Args:  containerArgs,
							},
						},
					},
				},
			},
		},
	}
}