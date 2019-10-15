package gpdbresource

import (
	gpv1alpha1 "github.com/soxueren/greenplum-operator/pkg/apis/gp/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	MOUNT_NAME = "data"
	MOUNT_PATH = "/home/gpadmin/"
	NODE_TAGS  = []string{"master", "segment", "mirror"}
)

func NewPodForCR(cr *gpv1alpha1.GPDBCluster, tag string, suffix string) *corev1.Pod {
	labels := map[string]string{
		"app":  cr.Name,
		"name": cr.Name + "-" + tag + "-" + suffix,
		"tag": tag,
	}
	image := cr.Spec.MasterAndStandby.Image
	commond := []string{"sleep", "3600"}
	switch tag {
	case NODE_TAGS[0]:
		image = cr.Spec.MasterAndStandby.Image
	case NODE_TAGS[1]:
		image = cr.Spec.Segments.Image
	case NODE_TAGS[2]:
		image = cr.Spec.Mirrors.Image
	default:
	}
	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-" + tag + "-" + suffix,
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    tag,
					Image:   image,
					Command: commond,
					VolumeMounts: []corev1.VolumeMount{
						{
							Name:      MOUNT_NAME,
							MountPath: MOUNT_PATH,
						},
					},
				},
			},
			Volumes: []corev1.Volume{
				{
					Name: MOUNT_NAME,
					VolumeSource: corev1.VolumeSource{
						PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
							ClaimName: cr.Name + "-pvc-" + tag + "-" + suffix,
						},
					},
				},
			},
		},
	}
}
