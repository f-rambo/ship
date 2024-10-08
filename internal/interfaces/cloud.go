package interfaces

import (
	"context"

	v1alpha1 "github.com/f-rambo/ship/api/cloud/v1alpha1"
	"github.com/f-rambo/ship/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CloudInterface struct {
	v1alpha1.UnimplementedCloudInterfaceServer
	uc  *biz.CloudUsecase
	log *log.Helper
}

func NewCloudInterface(uc *biz.CloudUsecase, logger log.Logger) *CloudInterface {
	return &CloudInterface{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (c *CloudInterface) Ping(ctx context.Context, _ *emptypb.Empty) (*v1alpha1.Msg, error) {
	return &v1alpha1.Msg{Message: "pong"}, nil
}

// InstallKubeadmKubeletCriO
func (c *CloudInterface) InstallKubeadmKubeletCriO(ctx context.Context, req *v1alpha1.Cloud) (*v1alpha1.Msg, error) {
	if req.Arch == "" {
		return nil, errors.New("arch is empty")
	}
	if req.CrioVersion == "" {
		return nil, errors.New("crio version is empty")
	}
	cloud := c.interfaceToBiz(req)
	err := c.uc.InstallKubeadmKubeletCriO(ctx, cloud)
	if err != nil {
		return nil, err
	}
	return &v1alpha1.Msg{Message: "InstallKubeadmKubeletCriO"}, nil
}

// AddKubeletServiceAndSettingKubeadmConfig
func (c *CloudInterface) AddKubeletServiceAndSettingKubeadmConfig(ctx context.Context, req *v1alpha1.Cloud) (*v1alpha1.Msg, error) {
	if req.KubeadmConfig == "" {
		return nil, errors.New("kubeadm config is empty")
	}
	if req.KubeletService == "" {
		return nil, errors.New("kubelet service is empty")
	}
	cloud := c.interfaceToBiz(req)
	err := c.uc.AddKubeletServiceAndSettingKubeadmConfig(ctx, cloud)
	if err != nil {
		return nil, err
	}
	return &v1alpha1.Msg{Message: "AddKubeletServiceAndSettingKubeadmConfig"}, nil
}

// KubeadmJoin
func (c *CloudInterface) KubeadmJoin(ctx context.Context, req *v1alpha1.Cloud) (*v1alpha1.Msg, error) {
	if req.ControlPlaneEndpoint == "" {
		return nil, errors.New("control plane endpoint is empty")
	}
	if req.Token == "" {
		return nil, errors.New("token is empty")
	}
	if req.DiscoveryTokenCaCertHash == "" {
		return nil, errors.New("discovery token ca cert hash is empty")
	}
	cloud := c.interfaceToBiz(req)
	err := c.uc.KubeadmJoin(ctx, cloud)
	if err != nil {
		return nil, err
	}
	return &v1alpha1.Msg{Message: "KubeadmJoin"}, nil
}

// KubeadmInit
func (c *CloudInterface) KubeadmInit(ctx context.Context, req *v1alpha1.Cloud) (*v1alpha1.Msg, error) {
	if req.KubeadmInitConfig == "" || req.KubeadmInitClusterConfig == "" {
		return nil, errors.New("kubeadm init config or kubeadm init cluster config is empty")
	}
	cloud := c.interfaceToBiz(req)
	err := c.uc.KubeadmInit(ctx, cloud)
	if err != nil {
		return nil, err
	}
	return &v1alpha1.Msg{Message: "KubeadmInit"}, nil
}

// KubeadmReset
func (c *CloudInterface) KubeadmReset(ctx context.Context, req *v1alpha1.Cloud) (*v1alpha1.Msg, error) {
	if req.KubeadmResetConfig == "" {
		return nil, errors.New("kubeadm reset config is empty")
	}
	cloud := c.interfaceToBiz(req)
	err := c.uc.KubeadmReset(ctx, cloud)
	if err != nil {
		return nil, err
	}
	return &v1alpha1.Msg{Message: "KubeadmReset"}, nil
}

// KubeadmUpgrade
func (c *CloudInterface) KubeadmUpgrade(ctx context.Context, req *v1alpha1.Cloud) (*v1alpha1.Msg, error) {
	if req.KubeadmUpgradeConfig == "" {
		return nil, errors.New("kubeadm upgrade config is empty")
	}
	cloud := c.interfaceToBiz(req)
	err := c.uc.KubeadmUpgrade(ctx, cloud)
	if err != nil {
		return nil, err
	}
	return &v1alpha1.Msg{Message: "KubeadmUpgrade"}, nil
}

// SetingIpv4Forward
func (c *CloudInterface) SetingIpv4Forward(ctx context.Context, _ *emptypb.Empty) (*v1alpha1.Msg, error) {
	err := c.uc.SetingIpv4Forward(ctx)
	if err != nil {
		return nil, err
	}
	return &v1alpha1.Msg{Message: "SetingIpv4Forward"}, nil
}

// CloseSwap
func (c *CloudInterface) CloseSwap(ctx context.Context, _ *emptypb.Empty) (*v1alpha1.Msg, error) {
	err := c.uc.CloseSwap(ctx)
	if err != nil {
		return nil, err
	}
	return &v1alpha1.Msg{Message: "CloseSwap"}, nil
}

// CloseFirewall
func (c *CloudInterface) CloseFirewall(ctx context.Context, _ *emptypb.Empty) (*v1alpha1.Msg, error) {
	err := c.uc.CloseFirewall(ctx)
	if err != nil {
		return nil, err
	}
	return &v1alpha1.Msg{Message: "CloseFirewall"}, nil
}

// interface cloud to biz cloud
func (c *CloudInterface) interfaceToBiz(cloud *v1alpha1.Cloud) *biz.Cloud {
	return &biz.Cloud{
		ID:                       cloud.Id,
		KubeadmInitConfig:        cloud.KubeadmInitConfig,
		KubeadmConfig:            cloud.KubeadmConfig,
		KubeletService:           cloud.KubeletService,
		CRIOVersion:              cloud.CrioVersion,
		ARCH:                     cloud.Arch,
		Token:                    cloud.Token,
		DiscoveryTokenCACertHash: cloud.DiscoveryTokenCaCertHash,
		ControlPlaneEndpoint:     cloud.ControlPlaneEndpoint,
		JoinConfig:               cloud.JoinConfig,
	}
}
