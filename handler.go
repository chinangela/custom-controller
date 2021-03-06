package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	networking "istio.io/api/networking/v1alpha3"
	"istio.io/istio/pilot/pkg/config/kube/crd"
	"istio.io/istio/pilot/pkg/model"
	core_v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Handler interface contains the methods that are required
type Handler interface {
	Init() error
	ObjectCreated(obj interface{})
	ObjectDeleted(obj interface{})
	ObjectUpdated(objOld, objNew interface{})
}

// TestHandler is a sample implementation of Handler
type TestHandler struct {
	workerIP string
}

// Init handles any handler initialization
func (t *TestHandler) Init() error {
	log.Info("TestHandler.Init")
	return nil
}

// ObjectCreated is called when an object is created
func (t *TestHandler) ObjectCreated(obj interface{}) {
	log.Info("TestHandler.ObjectCreated")
	// assert the type to a service object to pull out relevant data
	service := obj.(*core_v1.Service)
	log.Infof("    Name: %s", service.ObjectMeta.Name)
	log.Infof("    Annotations: %+v", service.ObjectMeta.Annotations)
	// log.Infof("    NodeName: %s", service.Spec.NodeName)
	log.Infof("    Type: %s", service.Spec.Type)
	log.Infof("    Ports: %+v", service.Spec.Ports)
	log.Infof("    Node!!!!: %s", t.workerIP)

	// create client
	client := getKubernetesClient(os.Getenv("HOME") + "/.kube/istio-config")

	annotations := service.ObjectMeta.Annotations
	if annotations["route"] == "true" && service.Spec.Type == core_v1.ServiceTypeNodePort {
		//create the endpoint
		_, err := client.CoreV1().Endpoints(meta_v1.NamespaceDefault).Create(&core_v1.Endpoints{
			TypeMeta: meta_v1.TypeMeta{
				Kind:       "Endpoints",
				APIVersion: "v1",
			},
			ObjectMeta: meta_v1.ObjectMeta{
				Name:      "neverok",
				Namespace: meta_v1.NamespaceDefault,
			},
			Subsets: []core_v1.EndpointSubset{
				{
					Addresses: []core_v1.EndpointAddress{
						{
							IP: t.workerIP,
						},
					},
					Ports: []core_v1.EndpointPort{
						{
							Port:     service.Spec.Ports[0].NodePort,
							Protocol: core_v1.ProtocolTCP,
						},
					},
				},
			},
		})

		if err != nil {
			panic(err)
		}

		service, err = client.CoreV1().Services(meta_v1.NamespaceDefault).Create(&core_v1.Service{
			TypeMeta: meta_v1.TypeMeta{
				Kind:       "Service",
				APIVersion: "v1",
			},
			ObjectMeta: meta_v1.ObjectMeta{
				Name:      "neverok",
				Namespace: meta_v1.NamespaceDefault,
			},
			Spec: core_v1.ServiceSpec{
				Ports: []core_v1.ServicePort{
					{
						Protocol: core_v1.ProtocolTCP,
						Port:     int32(6789),
					},
				},
				Type: core_v1.ServiceTypeClusterIP,
			},
		})

		if err != nil {
			panic(err)
		}

		istioClient, err := crd.NewClient(os.Getenv("HOME")+"/.kube/istio-config", "", model.IstioConfigTypes, "")

		config := model.Config{
			ConfigMeta: model.ConfigMeta{
				Type:    model.VirtualService.Type,
				Version: model.VirtualService.Version,
				Name:    "neverok",
				Group:   model.VirtualService.Group + ".istio.io",
			},
			Spec: &networking.VirtualService{
				Gateways: []string{"cfcr-gateway"},
				Hosts:    []string{"*"},
				Http: []*networking.HTTPRoute{
					{
						Match: []*networking.HTTPMatchRequest{
							{
								Port: uint32(80),
							},
						},
						Route: []*networking.DestinationWeight{
							{
								Destination: &networking.Destination{
									Host: "neverok",
									Port: &networking.PortSelector{
										Port: &networking.PortSelector_Number{
											Number: 6789,
										},
									},
								},
							},
						},
					},
				},
			}}

		_, err = istioClient.Create(config)
		if err != nil {
			panic(err)
		}

		_, err = istioClient.Create(model.Config{
			ConfigMeta: model.ConfigMeta{
				Type:    model.Gateway.Type,
				Version: model.Gateway.Version,
				Name:    "cfcr-gateway",
				Group:   model.Gateway.Group + ".istio.io",
			},
			Spec: &networking.Gateway{
				Servers: []*networking.Server{
					{
						Port: &networking.Port{
							Number:   uint32(80),
							Name:     "port",
							Protocol: "HTTP",
						},
						Hosts: []string{"*"},
					},
				},
				Selector: map[string]string{
					"istio": "ingressgateway",
				},
			},
		})

		if err != nil {
			panic(err)
		}
		// _, err := client.ExtensionsV1Beta1.VirtualServices(meta_v1.NamespaceDefault).Create(&istio_v1alpha3.VirtualService{
		// 	TypeMeta: meta_v1.TypeMeta{
		// 		Kind:       "VirtualService",
		// 		APIVersion: "networking.istio.io/v1alpha3",
		// 	},
		// 	ObjectMeta: meta_v1.ObjectMeta{
		// 		Name:      string(endpoint.ObjectMeta.Name),
		// 		Namespace: meta_v1.NamespaceDefault,
		// 	},
		// 	Spec: istio_v1alpha3.VirtualServiceSpec{
		// 		Gateways: []istio_v1alpha3.Gateway{
		// 			"cfcr-gateway",
		// 		},
		// 		Hosts: []istio_v1alpha3.Host{
		// 			"'*'",
		// 		},
		// 		HTTP: []istio_v1alpha3.HTTPRoute{
		// 			{
		// 				Match: []istio_v1alpha3.MatchRequest{
		// 					{
		// 						Port: uint32(8000),
		// 					},
		// 				},
		// 				Route: []istio_v1alpha3.DestinationWeight{
		// 					{
		// 						Destination: istio_v1alpha3.Destination{
		// 							Host: service.ObjectMeta.Name,
		// 							Port: istio_v1alpha3.PortSelector{
		// 								Port: istio_v1alpha3.PortSelector_Name{
		// 									Name: "6789",
		// 								},
		// 							},
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// })
	}
}

// ObjectDeleted is called when an object is deleted
func (t *TestHandler) ObjectDeleted(obj interface{}) {
	log.Info("TestHandler.ObjectDeleted")
}

// ObjectUpdated is called when an object is updated
func (t *TestHandler) ObjectUpdated(objOld, objNew interface{}) {
	log.Info("TestHandler.ObjectUpdated")
}
