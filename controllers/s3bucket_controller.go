/*


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

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	awscloudv1 "s3controller/api/v1"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"fmt"
	"os"
)

// S3BucketReconciler reconciles a S3Bucket object
type S3BucketReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

var (
	controllerLog = ctrl.Log.WithName("controller")
)

// +kubebuilder:rbac:groups=awscloud.vlad.example.com,resources=s3buckets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=awscloud.vlad.example.com,resources=s3buckets/status,verbs=get;update;patch

func (r *S3BucketReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("s3bucket", req.NamespacedName)

	controllerLog.Info("Running reconciler loop")

	var s3Bucket awscloudv1.S3Bucket

	_ = s3Bucket

	controllerLog.Info("var s3Bucket awscloudv1.S3Bucket")

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewSharedCredentials("", "default"),
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	svc := s3.New(sess)
	_ = svc

	listBucketInut := &s3.ListBucketsInput{}

	listBucket, err := svc.ListBuckets(listBucketInut)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, s := range listBucket.Buckets {
		fmt.Println(*s.Name)
	}

	return ctrl.Result{}, nil
}

func (r *S3BucketReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&awscloudv1.S3Bucket{}).
		Complete(r)
}
