package gpdbresource

import (
	gpv1alpha1 "github.com/soxueren/greenplum-operator/pkg/apis/gp/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func NewService(cr *gpv1alpha1.GPDBCluster) *corev1.Service {
	
	labels := map[string]string{
		"app": cr.Name,		
	}

	selector := make(map[string]string)
	selector["app"] = cr.Name
	selector["name"] = cr.Name + "-" + cr.Spec.MasterSelector

	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name,
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: corev1.ServiceSpec{
			Type:     corev1.ServiceTypeNodePort,
			Selector: selector,
			Ports: []corev1.ServicePort{
				{
					Protocol:   corev1.ProtocolTCP,
					Name:       "default-postgres",
					Port:       5432,
					TargetPort: intstr.FromInt(5432),
				},
			},
		},
	}
}
