package k8s_client

import (
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/typed/apps/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/retry"
)

var (
	deploymentsClient v1.DeploymentInterface
	err error
)

const k8s_api = "http://k8s.****.com"

func init() {
	//var kubeconfig *string
	config, err := clientcmd.BuildConfigFromFlags(k8s_api, "")
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	deploymentsClient = clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
}

func Create(name, app, image string, pods, port int32) error {
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(pods),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": app,
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": app,
					},
				},
				Spec: apiv1.PodSpec{
					ImagePullSecrets: []apiv1.LocalObjectReference{
						{
							Name: "Test",
						},
					},
					HostAliases: []apiv1.HostAlias{
						{
							IP: "10.0.0.11",
							Hostnames:[]string{
								"baoxian.abc.cn",
								"collection.abc.cn",
								"fr-customer.abc.cn",
								"order-miniprogram.abc.cn",
								"api-filter.abc.cn",
								"release.order.abc.cn",
								"release.openapi.abc.cn",
								"release.backend.abc.cn",
							},
						},
						{
							IP: "10.0.2.1",
							Hostnames:[]string{
								"beanstalkd.abc.cn",
							},
						},
						{
							IP: "10.0.2.30",
							Hostnames:[]string{
								"ebusiness.abc.cn",
								"ebusiness.abc.cn",
								"publicfund.abc.cn",
							},
						},
						{
							IP: "10.0.2.50",
							Hostnames:[]string{
								"bankcardcheck.abc.cn",
							},
						},
						{
							IP: "10.0.0.70",
							Hostnames:[]string{
								"bigdata.abc.cn",
							},
						},
						{
							IP: "10.0.2.80",
							Hostnames:[]string{
								"elk-redis.abc.cn",
							},
						},
						{
							IP: "10.0.100.230",
							Hostnames:[]string{
								"order.abc.cn",
							},
						},
						{
							IP: "10.0.100.231",
							Hostnames:[]string{
								"jieba-api.abc.cn",
							},
						},
						{
							IP: "10.0.100.232",
							Hostnames:[]string{
								"data.abc.cn",
							},
						},
						{
							IP: "10.0.100.233",
							Hostnames:[]string{
								"qianba.abc.cn",
							},
						},
						{
							IP: "10.0.100.234",
							Hostnames:[]string{
								"marketing.abc.cn",
							},
						},
						{
							IP: "10.0.100.235",
							Hostnames:[]string{
								"jjcardgo.abc.cn",
							},
						},
						{
							IP: "10.0.100.236",
							Hostnames:[]string{
								"qingsuan.abc.cn",
							},
						},
						{
							IP: "10.0.100.237",
							Hostnames:[]string{
								"blackcard.abc.cn",
							},
						},
						{
							IP: "10.0.100.238",
							Hostnames:[]string{
								"frio.abc.cn",
							},
						},
						{
							IP: "10.0.100.239",
							Hostnames:[]string{
								"basic_info.abc.cn",
							},
						},
					},
					Containers: []apiv1.Container{
						{
							Name:  name,
							Image: image,
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: port,
								},
							},
						},
					},
				},
			},
		},
	}
	_, err = deploymentsClient.Create(deployment)
	return err
}

func Update(name, image string, pods int32) error {
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		result, getErr := deploymentsClient.Get(name, metav1.GetOptions{})
		if getErr != nil {
			return getErr
		}

		result.Spec.Replicas = int32Ptr(pods)                   // reduce replica count
		result.Spec.Template.Spec.Containers[0].Image = image 	// change nginx version
		_, updateErr := deploymentsClient.Update(result)
		return updateErr
	})
	if retryErr != nil {
		return retryErr
	}
	return err
}

func int32Ptr(i int32) *int32 { return &i }
