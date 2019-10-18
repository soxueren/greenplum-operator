package gpdbsegment

import (
	"context"
	"strconv"

	gpv1alpha1 "github.com/soxueren/greenplum-operator/pkg/apis/gp/v1alpha1"
	"github.com/soxueren/greenplum-operator/pkg/controller/gpdbresource"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var _ reconcile.Reconciler = &ReconcileSegment{}

type ReconcileSegment struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// newReconciler returns a new reconcile.Reconciler
func newSegmentReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileSegment{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// Reconcile reads that state of the cluster for a GPDBCluster object and makes changes based on the state read
// and what is in the GPDBCluster.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileSegment) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling GPDBCluster segment")

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
	
	r.CreatePVCForCR(instance)	
	// //TODO when the number of  pods is  more than Replicas,reduce it ?
	r.CreateNodeForCR(instance)	

	return reconcile.Result{}, nil
}

func (r *ReconcileSegment) CreatePVCForCR(instance *gpv1alpha1.GPDBCluster) (reconcile.Result, error) {
	reqLogger := log.WithValues("PersistentVolumeClaim.initialize")
	size := instance.Spec.Segments.Replicas
	// when GPDBCluster created,create pvc and create gpdb node pod
	for i := 0; i < size; i++ {
		//TODO PersistentVolumeClaim
		pvc := gpdbresource.NewPersistentVolume(instance, tag, strconv.Itoa(i))
		reqLogger.Info("Creating a new pvc", "pvc.Name", pvc.Name)

		// Set GPDBCluster instance as the owner and controller
		if err := controllerutil.SetControllerReference(instance, pvc, r.scheme); err != nil {
			return reconcile.Result{}, err
		}

		//Check if this PersistentVolumeClaim already exists
		found := &corev1.PersistentVolumeClaim{}
		err := r.client.Get(context.TODO(), types.NamespacedName{Name: pvc.Name, Namespace: pvc.Namespace}, found)
		if err != nil && errors.IsNotFound(err) {
			reqLogger.Info("Creating a new PersistentVolumeClaim", "pvc.Namespace", pvc.Namespace, "pvc.Name", pvc.Name)
			err = r.client.Create(context.TODO(), pvc)
			if err != nil {
				return reconcile.Result{}, err
			}

			// Pod created successfully - don't requeue
			return reconcile.Result{}, nil
		} else if err != nil {
			return reconcile.Result{}, err
		}

		// Pod already exists - don't requeue
		reqLogger.Info("Skip reconcile: pvc already exists", "pvc.Namespace", found.Namespace, "pvc.Name", found.Name)

	}
	return reconcile.Result{}, nil
}

func (r *ReconcileSegment) CreateNodeForCR(instance *gpv1alpha1.GPDBCluster) (reconcile.Result, error) {
	reqLogger := log.WithValues("Pod.initialize")
	size := instance.Spec.Segments.Replicas
	// when GPDBCluster created,create pvc and create gpdb node pod
	for i := 0; i < size; i++ {
		//create pod object
		pod := gpdbresource.NewPodForCR(instance, tag, strconv.Itoa(i))

		// Set GPDBCluster instance as the owner and controller
		if err := controllerutil.SetControllerReference(instance, pod, r.scheme); err != nil {
			return reconcile.Result{}, err
		}

		//Check if this Pod already exists
		found := &corev1.Pod{}
		err := r.client.Get(context.TODO(), types.NamespacedName{Name: pod.Name, Namespace: pod.Namespace}, found)
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
