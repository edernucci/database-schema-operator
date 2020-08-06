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
	"fmt"

	"github.com/edernucci/database-schema-operator/pkg/helpers"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	dbv1 "github.com/edernucci/database-schema-operator/api/v1"
)

// TableReconciler reconciles a Table object
type TableReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=db.pedag.io,resources=tables,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=db.pedag.io,resources=tables/status,verbs=get;update;patch

func (r *TableReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("table", req.NamespacedName)

	// your logic here
	var table dbv1.Table
	if err := r.Get(ctx, req.NamespacedName, &table); err != nil {
		log.Error(err, "unable to fetch Table")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	exists, err := helpers.CheckTable(table.Spec.Name)
	if err != nil {
		log.Error(err, "Error checking table")
		return ctrl.Result{}, nil
	}

	if exists {
		log.Info(fmt.Sprintf("The table [%s] exists on database.", table.Spec.Name))
		helpers.UpdateColumns(table.Spec.Name, table.Spec.Columns)
	} else {
		log.Info(fmt.Sprintf("Creating table [%s] on database.", table.Spec.Name))
		_, err := helpers.CreateTable(table.Spec.Name, table.Spec.Columns)
		if err != nil {
			log.Error(err, "Error creating table")
			return ctrl.Result{}, nil
		}
	}

	// for _, column := range table.Spec.Columns {
	// 	log.Info(column.Name)
	// }

	return ctrl.Result{}, nil
}

func (r *TableReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&dbv1.Table{}).
		Complete(r)
}
