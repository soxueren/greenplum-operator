package gpdbresource

import (
	gpv1alpha1 "github.com/soxueren/greenplum-operator/pkg/apis/gp/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewPersistentVolume(cr *gpv1alpha1.GPDBCluster, tag string, suffix string) *corev1.PersistentVolumeClaim {
	labels := map[string]string{
		"app": cr.Name + "-pvc-" + tag + "-" + suffix,
		"tag": tag,
	}
	return &corev1.PersistentVolumeClaim{
		TypeMeta: metav1.TypeMeta{
			Kind:       "PersistentVolumeClaim",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-pvc-" + tag + "-" + suffix,
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
