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
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var providerlog = logf.Log.WithName("provider-resource")
var providerWebhookK8sClient k8sclient.Client

func (r *Provider) SetupWebhookWithManager(mgr ctrl.Manager) error {
	providerWebhookK8sClient = mgr.GetClient()
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-circleci-meowhq-dev-v1-provider,mutating=true,failurePolicy=fail,sideEffects=None,groups=circleci.meowhq.dev,resources=providers,verbs=create;update,versions=v1,name=mprovider.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &Provider{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Provider) Default() {
	providerlog.Info("default", "name", r.Name)

	if r.Spec.Host == "" {
		r.Spec.Host = "https://circleci.com"
	}
	if r.Spec.RestEndpoint == "" {
		r.Spec.RestEndpoint = "api/v2"
	}
	if r.Spec.Endpoint == "" {
		r.Spec.Endpoint = "graphql-unstable"
	}
	if r.Spec.Auth.Token.SecretKeyRef.Namespace == "" {
		r.Spec.Auth.Token.SecretKeyRef.Namespace = r.Namespace
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-circleci-meowhq-dev-v1-provider,mutating=false,failurePolicy=fail,sideEffects=None,groups=circleci.meowhq.dev,resources=providers,verbs=create;update,versions=v1,name=vprovider.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &Provider{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Provider) ValidateCreate() error {
	providerlog.Info("validate create", "name", r.Name)

	return r.validateProviderTokenSecretRef(providerWebhookK8sClient)
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Provider) ValidateUpdate(old runtime.Object) error {
	providerlog.Info("validate update", "name", r.Name)

	return r.validateProviderTokenSecretRef(providerWebhookK8sClient)
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Provider) ValidateDelete() error {
	providerlog.Info("validate delete", "name", r.Name)

	// Do Nothing
	return nil
}

// validateProviderTokenSecretRef validate if Secret is existing
func (r *Provider) validateProviderTokenSecretRef(client k8sclient.Client) error {
	ref := &corev1.Secret{}
	err := providerWebhookK8sClient.Get(context.TODO(), k8sclient.ObjectKey{
		Name:      r.Spec.Auth.Token.SecretKeyRef.Name,
		Namespace: r.Namespace,
	}, ref)

	if err != nil {
		// error including not found error
		return err
	}

	// return error if secret key is not found
	if _, ok := ref.Data[r.Spec.Auth.Token.SecretKeyRef.Key]; !ok {
		return fmt.Errorf("secret key '%s' is not existing", r.Spec.Auth.Token.SecretKeyRef.Key)
	}

	return nil
}
