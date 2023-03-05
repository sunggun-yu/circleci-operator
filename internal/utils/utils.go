package utils

import (
	"context"
	"fmt"
	"strings"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	circleciv1 "github.com/sunggun-yu/circleci-operator/api/circleci/v1"
)

// FindSecret finds a corev1 secret
func FindSecret(c client.Client, namespace, name string) (*corev1.Secret, error) {
	secret := &corev1.Secret{}
	ref := types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}

	err := c.Get(context.TODO(), ref, secret)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", ref.Name, err)
	}
	return secret, nil
}

// GetSecretValue
func GetSecretValue(c client.Client, secretRef *circleciv1.SecretKeySelector) (string, error) {
	secret, err := FindSecret(c, secretRef.Namespace, secretRef.Name)
	if err != nil {
		return "", err
	}

	keyBytes, ok := secret.Data[secretRef.Key]
	if !ok {
		return "", fmt.Errorf("%s is not found", secretRef.Key)
	}

	value := string(keyBytes)
	valueStr := strings.TrimSpace(value)
	return valueStr, nil
}
