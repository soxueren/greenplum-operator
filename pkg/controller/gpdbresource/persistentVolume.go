package gpdbresource

import (
	"strconv"

	gpv1alpha1 "github.com/soxueren/greenplum-operator/pkg/apis/gp/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewPersistentVolume(cr *gpv1alpha1.GPDBCluster, nodetype string, num int) *corev1.PersistentVolumeClaim {
	labels := map[string]string{
		"app": cr.Name + "-" + nodetype + "-pvc-" + strconv.Itoa(num),
	}
	return &corev1.PersistentVolumeClaim{
		TypeMeta: metav1.TypeMeta{
			Kind:       "PersistentVolumeClaim",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-" + nodetype + "-pvc-" + strconv.Itoa(num),
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: corev1.PersistentVolumeClaimSpec{
			AccessModes: []corev1.PersistentVolumeAccessMode{
				corev1.ReadWriteOnce,
			},
			Resources: corev1.ResourceRequirements{
				Requests: corev1.ResourceList{
					"storage": cr.Spec.MasterAndStandby.Storage,
				},
			},
			StorageClassName: &cr.Spec.MasterAndStandby.StorageClassName,
		},
	}
}
