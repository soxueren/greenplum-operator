package gpdbmirror

import (
	"context"
	"strconv"

	gpv1alpha1 "github.com/soxueren/greenplum-operator/pkg/apis/gp/v1alpha1"
	"github.com/soxueren/greenplum-operator/pkg/controller/gpdbresource"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var _ reconcile.Reconciler = &ReconcileMirror{}

type ReconcileMirror struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// newReconciler returns a new reconcile.Reconciler
func newMirrorReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileMirror{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// Reconcile reads that state of the cluster for a GPDBCluster object and makes changes based on the state read
// and what is in the GPDBCluster.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileMirror) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling GPDBCluster Mirror")

	// Fetch the GPDBCluster instance
	instance := &gpv1alpha1.GPDBCluster{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}
	// when GPDBCluster created,create pvc and create gpdb node pod
	for i := 0; i < instance.Spec.Mirrors.Replicas; i++ {
		//TODO PersistentVolumeClaim
		pvc := gpdbresource.NewPersistentVolume(instance, "mirror", i)
		reqLogger.Info("Creating a new pvc", "pvc.Name", pvc.Name)
		//create pod object
		pod := newMirrorForCR(instance, i)

		// Set GPDBCluster instance as the owner and controller
		if err := controllerutil.SetControllerReference(instance, pod, r.scheme); err != nil {
			return reconcile.Result{}, err
		}

		//Check if this Pod already exists
		found := &corev1.Pod{}
		err = r.client.Get(context.TODO(), types.NamespacedName{Name: pod.Name, Namespace: pod.Namespace}, found)
		if err != nil && errors.IsNotFound(err) {
			reqLogger.Info("Creating a new Pod", "Pod.Namespace", pod.Namespace, "Pod.Name", pod.Name)
			err = r.client.Create(context.TODO(), pod)
			if err != nil {
				return reconcile.Result{}, err
			}

			// Pod created successfully - don't requeue
			return reconcile.Result{}, nil
		} else if err != nil {
			return reconcile.Result{}, err
		}
		// Pod already exists - don't requeue
		reqLogger.Info("Skip reconcile: Pod already exists", "Pod.Namespace", found.Namespace, "Pod.Name", found.Name)

	}

	return reconcile.Result{}, nil
}

// newPodForCR returns a busybox pod with the same name/namespace as the cr
func newMirrorForCR(cr *gpv1alpha1.GPDBCluster, num int) *corev1.Pod {
	labels := map[string]string{
		"app":  cr.Name,
		"name": cr.Name + "-mirror-" + strconv.Itoa(num),
	}
	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-mirror-" + strconv.Itoa(num),
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    "mirror",
					Image:   "busybox",
					Command: []string{"sleep", "3600"},
					VolumeMounts: []corev1.VolumeMount{
						{
							Name:      "data",
							MountPath: "/home/gpadmin/",
						},
					},
				},
			},
			Volumes: []corev1.Volume{
				{
					Name: "data",
					VolumeSource: corev1.VolumeSource{
						PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
							ClaimName: cr.Name + "-mirror-pvc-" + strconv.Itoa(num),
						},
					},
				},
			},
		},
	}
}
