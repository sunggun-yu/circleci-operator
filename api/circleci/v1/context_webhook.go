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

package v1

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var contextlog = logf.Log.WithName("context-resource")
var contextWebhookK8sClient k8sclient.Client

func (r *Context) SetupWebhookWithManager(mgr ctrl.Manager) error {
	contextWebhookK8sClient = mgr.GetClient()
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-circleci-meowhq-dev-v1-context,mutating=true,failurePolicy=fail,sideEffects=None,groups=circleci.meowhq.dev,resources=contexts,verbs=create;update,versions=v1,name=mcontext.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &Context{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Context) Default() {
	contextlog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-circleci-meowhq-dev-v1-context,mutating=false,failurePolicy=fail,sideEffects=None,groups=circleci.meowhq.dev,resources=contexts,verbs=create;update,versions=v1,name=vcontext.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &Context{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Context) ValidateCreate() error {
	contextlog.Info("validate create", "name", r.Name)

	return r.validateContextProviderRef(contextWebhookK8sClient)
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Context) ValidateUpdate(old runtime.Object) error {
	contextlog.Info("validate update", "name", r.Name)

	return r.validateContextProviderRef(contextWebhookK8sClient)
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Context) ValidateDelete() error {
	contextlog.Info("validate delete", "name", r.Name)

	// Do Nothing
	return nil
}

// validateProviderTokenSecretRef validate if Secret is existing
func (r *Context) validateContextProviderRef(client k8sclient.Client) error {

	ref := &Provider{}
	err := client.Get(context.TODO(), k8sclient.ObjectKey{
		Name:      r.Spec.Provider.Name,
		Namespace: r.Namespace,
	}, ref)

	if err != nil {
		// error including not found error
		return err

	}
	return nil
}
