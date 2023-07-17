package pkg

import (
	informerLister "k8s.io/client-go/informers/core/v1"
	networkingLister "k8s.io/client-go/informers/networking/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/listers/core/v1"
	v12 "k8s.io/client-go/listers/networking/v1"
	"k8s.io/client-go/tools/cache"
)

type Controller struct {
	client        kubernetes.Interface
	serviceLister v1.ServiceLister
	ingressLister v12.IngressLister
}

func (c *Controller) Run(stopCh chan struct{}) {
	<-stopCh
}

func (c *Controller) addService(obj interface{}) {

}

func (c *Controller) updateService(oldObj interface{}, newObj interface{}) {

}

func (c *Controller) deleteIngress(obj interface{}) {

}

func NewController(clientSet kubernetes.Interface, serviceInformer informerLister.ServiceInformer, ingressInformer networkingLister.IngressInformer) Controller {
	c := Controller{
		client:        clientSet,
		serviceLister: serviceInformer.Lister(),
		ingressLister: ingressInformer.Lister(),
	}

	serviceInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    c.addService,
		UpdateFunc: c.updateService,
	})

	ingressInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		DeleteFunc: c.deleteIngress,
	})
	return c
}
