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
							Name: "mifengkong",
						},
					},
					HostAliases: []apiv1.HostAlias{
						{
							IP: "10.0.0.11",
							Hostnames:[]string{
								"baoxian.mifengkong.cn",
								"collection.mifengkong.cn",
								"fr-customer.mifengkong.cn",
								"order-miniprogram.mifengkong.cn",
								"api-filter.mifengkong.cn",
								"release.order.mifengkong.cn",
								"release.openapi.mifengkong.cn",
								"release.backend.mifengkong.cn",
							},
						},
						{
							IP: "10.0.2.1",
							Hostnames:[]string{
								"beanstalkd.mifengkong.cn",
							},
						},
						{
							IP: "10.0.2.30",
							Hostnames:[]string{
								"ebusiness.mifengkong.cn",
								"ebusiness.mifengkong.cn",
								"publicfund.mifengkong.cn",
								"bill.mifengkong.cn",
								"personcredit.mifengkong.cn",
								"telcos.mifengkong.cn",
								"upload.mifengkong.cn",
								"push.mifengkong.cn",
								"alipay.mifengkong.cn",
								"alimns.mifengkong.cn",
								"addinfo.mifengkong.cn",
								"message.mifengkong.cn",
								"odps2oss.mifengkong.cn",
							},
						},
						{
							IP: "10.0.2.50",
							Hostnames:[]string{
								"bankcardcheck.mifengkong.cn",
								"compact.mifengkong.cn",
								"credit.mifengkong.cn",
							},
						},
						{
							IP: "10.0.0.70",
							Hostnames:[]string{
								"bigdata.mifengkong.cn",
							},
						},
						{
							IP: "10.0.2.80",
							Hostnames:[]string{
								"elk-redis.mifengkong.cn",
							},
						},
						{
							IP: "10.0.100.230",
							Hostnames:[]string{
								"order.mifengkong.cn",
								"openapi.mifengkong.cn",
								"backend.mifengkong.cn",
							},
						},
						{
							IP: "10.0.100.231",
							Hostnames:[]string{
								"jieba-api.mifengkong.cn",
							},
						},
						{
							IP: "10.0.100.232",
							Hostnames:[]string{
								"data.mifengkong.cn",
								"datanew.mifengkong.cn",
								"basicinfo.mifengkong.cn",
							},
						},
						{
							IP: "10.0.100.233",
							Hostnames:[]string{
								"qianba.mifengkong.cn",
								"product.mifengkong.cn",
								"userprofile.mifengkong.cn",
								"nameverification.mifengkong.cn",
							},
						},
						{
							IP: "10.0.100.234",
							Hostnames:[]string{
								"marketing.mifengkong.cn",
								"oauth.mifengkong.cn",
								"package.mifengkong.cn",
								"vaccount.mifengkong.cn",
								"advertise_platform.mifengkong.cn",
								"frhoutai-api.mifengkong.cn",
								"miniprogram.mifengkong.cn",
								"promotion.mifengkong.cn",
								"userdata.mifengkong.cn",
								"orderdata.mifengkong.cn",
								"statistics.mifengkong.cn",
								"release.miniprogram.mifengkong.cn",
								"release.promotion.mifengkong.cn",
								"statement.mifengkong.cn",
							},
						},
						{
							IP: "10.0.100.235",
							Hostnames:[]string{
								"moxie.mifengkong.cn",
								"send_service.mifengkong.cn",
								"credit_card.mifengkong.cn",
								"product_filter.mifengkong.cn",
								"open_platform.mifengkong.cn",
								"h5-server.mifengkong.cn",
								"jjcardgo.mifengkong.cn",
							},
						},
						{
							IP: "10.0.100.236",
							Hostnames:[]string{
								"qingsuan.mifengkong.cn",
							},
						},
						{
							IP: "10.0.100.237",
							Hostnames:[]string{
								"jiexiaoer.mifengkong.cn",
								"release.jiexiaoer.mifengkong.cn",
								"xinshen.mifengkong.cn",
								"pushlumen.mifengkong.cn",
								"exezcc.mifengkong.cn",
								"wechat_admin.mifengkong.cn",
								"api.mifengkong.cn",
								"release.api.mifengkong.cn",
								"kamao.mifengkong.cn",
								"credit_cms_api.mifengkong.cn",
								"blackcard.mifengkong.cn",
							},
						},
						{
							IP: "10.0.100.238",
							Hostnames:[]string{
								"frio.mifengkong.cn",
								"kafka_service.mifengkong.cn",
								"kafka-manager.mifengkong.cn",
								"third-party-service.mifengkong.cn",
							},
						},
						{
							IP: "10.0.100.239",
							Hostnames:[]string{
								"basic_info.mifengkong.cn",
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
