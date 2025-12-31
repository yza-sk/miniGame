package svc

import (
	"fmt"
	"rank_list/internal/config"
	"rank_list/internal/entity"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"gopkg.in/yaml.v2"
)

type NacosService interface {
	GetMysqlConfig() (config.MysqlConfig, error)
}

func NewNacosClient() (nacosService, error) {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("8.148.21.13", 8848, constant.WithContextPath("/nacos")),
	}

	//create ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId("8a323c73-74bb-4890-95ed-f911becd031b"),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)

	// create config client
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		fmt.Printf("Error creating Nacos client: %v\n", err)
		return nacosService{}, fmt.Errorf("Error creating Nacos client: %v", err)
	}
	return nacosService{client: client}, nil
}

type nacosService struct {
	client config_client.IConfigClient
}

func (s *nacosService) GetConfig(cp entity.ConfigParam) (interface{}, error) {
	configContent, err := s.client.GetConfig(vo.ConfigParam{
		DataId: cp.DataId,
		Group:  cp.Group,
	})
	if err != nil {
		return config.MysqlConfig{}, fmt.Errorf("Error getting config from Nacos: %v", err)
	}
	// fmt.Printf("%s", configContent)
	err = yaml.Unmarshal([]byte(configContent), &cp.ConfigType)
	return cp.ConfigType, nil
}
