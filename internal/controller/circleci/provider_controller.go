/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package circleci

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	"github.com/go-logr/logr"
	circleciv1 "github.com/sunggun-yu/circleci-operator/api/circleci/v1"
	"github.com/sunggun-yu/circleci-operator/internal/circleci/apiclient"
)

// ProviderReconciler reconciles a Provider object
type ProviderReconciler struct {
	client.Client
	Scheme *runtime.Scheme

	Log logr.Logger
}

const (
	providerSecretWatcherIndexField = ".spec.auth.token.secretKeyRef.name"
)

//+kubebuilder:rbac:groups=circleci.meowhq.dev,resources=providers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=circleci.meowhq.dev,resources=providers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=circleci.meowhq.dev,resources=providers/finalizers,verbs=update
//+kubebuilder:rbac:groups="",resources=secrets,verbs=get;list;watch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Provider object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *ProviderReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	p := &circleciv1.Provider{}
	err := r.Get(ctx, req.NamespacedName, p)
	if err != nil {
		r.Log.Error(err, "cannot get provider")
		return ctrl.Result{}, err
	}

	client, err := apiclient.NewCircleciAPIClient(r.Log, r.Client, &p.Spec)
	if client == nil || err != nil {
		p.Status.Valid = false
	} else {
		p.Status.Valid = true
	}
	// update status of provider
	if err := r.Status().Update(ctx, p); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ProviderReconciler) SetupWithManager(mgr ctrl.Manager) error {

	// create the index to find the object in the watch list
	if err := mgr.GetFieldIndexer().IndexField(context.Background(), &circleciv1.Provider{}, providerSecretWatcherIndexField, func(rawObj client.Object) []string {
		// Extract the Secret name from the Provider Spec, if one is provided
		attached := rawObj.(*circleciv1.Provider)
		if attached.Spec.Auth.Token.SecretKeyRef.Name == "" {
			return nil
		}
		return []string{attached.Spec.Auth.Token.SecretKeyRef.Name}
	}); err != nil {
		return err
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(&circleciv1.Provider{}).
		WithEventFilter(predicate.Funcs{
			CreateFunc: func(e event.CreateEvent) bool {
				r.Log.Info("event filter - create event", "name", e.Object.GetName())
				return true
			},
			DeleteFunc: func(e event.DeleteEvent) bool {
				r.Log.Info("event filter - delete event", "name", e.Object.GetName())
				return false
			},
			UpdateFunc: func(e event.UpdateEvent) bool {
				oldGeneration := e.ObjectOld.GetGeneration()
				newGeneration := e.ObjectNew.GetGeneration()
				r.Log.Info("event filter - update event", "update", oldGeneration != newGeneration, "name", e.ObjectNew.GetName())
				return true
			},
			GenericFunc: func(event.GenericEvent) bool {
				return false
			},
		}).
		Watches(
			&source.Kind{Type: &corev1.Secret{}},
			handler.EnqueueRequestsFromMapFunc(r.findObjectsForSecret),
			builder.WithPredicates(predicate.ResourceVersionChangedPredicate{}),
		).
		Complete(r)
}

// findObjectsForSecret finds the objects that match the given predicate from indexed fields
func (r *ProviderReconciler) findObjectsForSecret(obj client.Object) []reconcile.Request {
	attached := &circleciv1.ProviderList{}
	listOps := &client.ListOptions{
		FieldSelector: fields.OneTermEqualSelector(providerSecretWatcherIndexField, obj.GetName()),
		Namespace:     obj.GetNamespace(),
	}
	err := r.List(context.TODO(), attached, listOps)
	if err != nil {
		return []reconcile.Request{}
	}

	requests := make([]reconcile.Request, len(attached.Items))
	for i, item := range attached.Items {
		requests[i] = reconcile.Request{
			NamespacedName: types.NamespacedName{
				Name:      item.GetName(),
				Namespace: item.GetNamespace(),
			},
		}
	}
	return requests
}
