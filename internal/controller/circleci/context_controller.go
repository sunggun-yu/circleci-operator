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
	"github.com/sunggun-yu/circleci-operator/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

// ContextReconciler reconciles a Context object
type ContextReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Log    logr.Logger
}

const (
	contextSecretWatcherIndexField = ".spec.env.SecretKeyRef.Name"
)

//+kubebuilder:rbac:groups=circleci.meowhq.dev,resources=contexts,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=circleci.meowhq.dev,resources=contexts/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=circleci.meowhq.dev,resources=contexts/finalizers,verbs=update
//+kubebuilder:rbac:groups=circleci.meowhq.dev,resources=providers,verbs=get;list;watch
//+kubebuilder:rbac:groups=circleci.meowhq.dev,resources=providers/status,verbs=get
//+kubebuilder:rbac:groups="",resources=secrets,verbs=get;list;watch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Context object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *ContextReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	cciContext := &circleciv1.Context{}
	err := r.Get(ctx, req.NamespacedName, cciContext)
	if err != nil {
		return reconcile.Result{}, err
	}

	cciContextProvider := &circleciv1.Provider{}
	err = r.Get(ctx, client.ObjectKey{Namespace: req.Namespace, Name: cciContext.Spec.Provider.Name}, cciContextProvider)
	if err != nil {
		return reconcile.Result{}, err
	}

	cclient, err := apiclient.NewCircleciAPIClient(r.Log, r.Client, &cciContextProvider.Spec)
	if err != nil {
		return ctrl.Result{}, err
	}
	ccontextclient, err := cclient.NewContextClient()
	if err != nil {
		return ctrl.Result{}, err
	}

	ccx, err := ccontextclient.ContextByName(cciContext.Spec.Owner.VCS, cciContext.Spec.Owner.Name, cciContext.Spec.Name)
	if ccx == nil || err != nil {
		err := ccontextclient.CreateContext(cciContext.Spec.Owner.VCS, cciContext.Spec.Owner.Name, cciContext.Spec.Name)
		if err != nil {
			return reconcile.Result{}, err
		}
		ccx, err = ccontextclient.ContextByName(cciContext.Spec.Owner.VCS, cciContext.Spec.Owner.Name, cciContext.Spec.Name)
		if err != nil {
			return reconcile.Result{}, err
		}
		r.Log.Info("context has created", "context", ccx)
	}

	cciContext.Status.Id = ccx.ID
	cciContext.Status.CreatedAt = ccx.CreatedAt.String()

	// prepare map to set the env var from individual envvar and envfrom
	envMap := make(map[string]string)
	for _, e := range cciContext.Spec.Env {
		var val string
		if e.ValueFrom != nil {
			if e.ValueFrom.SecretKeyRef.Namespace == "" {
				e.ValueFrom.SecretKeyRef.Namespace = req.Namespace
			}
			val, err = utils.GetSecretValue(r.Client, e.ValueFrom.SecretKeyRef)
			if err != nil {
				continue
			}
		} else if e.Value != "" {
			val = e.Value
		}
		envMap[e.Name] = val
	}

	// create envMap from EnvFrom
	if cciContext.Spec.EnvFrom != nil {
		if cciContext.Spec.EnvFrom.SecretRef.Namespace == "" {
			cciContext.Spec.EnvFrom.SecretRef.Namespace = req.Namespace
		}
		envFromSecret, err := utils.FindSecret(r.Client, cciContext.Spec.EnvFrom.SecretRef.Namespace, cciContext.Spec.EnvFrom.SecretRef.Name)
		if err != nil {
			r.Log.Error(err, "failed to get secret for envFrom", "envFromSecret", envFromSecret)
		} else {
			for n, v := range envFromSecret.Data {
				envMap[n] = string(v)
			}
		}
	}

	// TODO: figure out why goroutine doesn't work
	// var wg sync.WaitGroup
	// ch := make(chan error)
	// for n, v := range envMap {
	// 	wg.Add(1)
	// 	go func(name, value string) {
	// 		defer wg.Done()
	// 		err := ccontextclient.CreateEnvironmentVariable(cciContext.Status.Id, name, value)
	// 		ch <- err
	// 	}(n, v)
	// }
	// wg.Wait()
	// close(ch)
	// for c := range ch {
	// 	if c != nil {
	// 		r.Log.Error(c, "failed to create/update the environment variable")
	// 	}
	// }

	// remove not changed environment variable from the map to avoid updating.
	for n, v := range envMap {
		if k, ok := cciContext.Status.Env[n]; ok {
			err := bcrypt.CompareHashAndPassword(k, []byte(v))
			if err == nil {
				r.Log.V(3).Info("value not changed. skip update", "name", n)
				delete(envMap, n)
			}
		}
	}

	r.Log.V(3).Info("update", "count", len(envMap))

	// Create/Update environment variable
	for n, v := range envMap {
		err := ccontextclient.CreateEnvironmentVariable(cciContext.Status.Id, n, v)
		if err != nil {
			r.Log.Error(err, "failed to create/update the environment variable", "name", n, "value", v)
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(v), bcrypt.DefaultCost)
		if err != nil {
			r.Log.Error(err, "failed to hashing the environment variable value", "name", n, "value", v)
		}
		if cciContext.Status.Env == nil {
			cciContext.Status.Env = make(map[string][]byte)
		}
		cciContext.Status.Env[n] = hash
	}

	if err := r.Status().Update(ctx, cciContext); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ContextReconciler) SetupWithManager(mgr ctrl.Manager) error {

	// create the index to find the object in the watch list
	if err := mgr.GetFieldIndexer().IndexField(context.Background(), &circleciv1.Context{}, contextSecretWatcherIndexField, func(rawObj client.Object) []string {
		// Extract the Secret name from the Provider Spec, if one is provided
		attached := rawObj.(*circleciv1.Context)
		var names []string
		// loop through the Env and get the secret names to be indexed as `spec.env.valueFrom.SecretKeyRef.Name`
		for _, e := range attached.Spec.Env {
			if e.ValueFrom != nil && e.ValueFrom.SecretKeyRef.Name != "" {
				names = append(names, e.ValueFrom.SecretKeyRef.Name)
			}
		}
		// add secret name for envfrom
		if attached.Spec.EnvFrom != nil && attached.Spec.EnvFrom.SecretRef.Name != "" {
			names = append(names, attached.Spec.EnvFrom.SecretRef.Name)
		}
		return names
	}); err != nil {
		return err
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(&circleciv1.Context{}).
		WithEventFilter(predicate.Funcs{
			CreateFunc: func(e event.CreateEvent) bool {
				r.Log.Info("event filter - create event", "name", e.Object.GetName())
				return true
			},
			UpdateFunc: func(e event.UpdateEvent) bool {
				oldGeneration := e.ObjectOld.GetGeneration()
				newGeneration := e.ObjectNew.GetGeneration()

				var proceed bool
				switch e.ObjectNew.DeepCopyObject().(type) {
				case *corev1.Secret:
					r.Log.Info("event filter - secret update watcher", "name", e.ObjectNew.GetName())
					// pass to reconciler if event is from watcher for secret update
					proceed = true
				default:
					r.Log.Info("event filter - update", "update", oldGeneration != newGeneration, "name", e.ObjectNew.GetName())
					proceed = oldGeneration != newGeneration
				}
				return proceed
			},
			DeleteFunc: func(e event.DeleteEvent) bool {
				r.Log.Info("event filter - delete event", "name", e.Object.GetName())
				// not necessary to proceed to reconcile since downstream resources are associated with previewenv as owner
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
func (r *ContextReconciler) findObjectsForSecret(obj client.Object) []reconcile.Request {
	attached := &circleciv1.ContextList{}
	listOps := &client.ListOptions{
		FieldSelector: fields.OneTermEqualSelector(contextSecretWatcherIndexField, obj.GetName()),
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
