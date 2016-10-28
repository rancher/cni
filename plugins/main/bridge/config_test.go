package main

import (
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/containernetworking/cni/pkg/types"
)

func TestErrorNetworkConfigMissingSubnet(t *testing.T) {

	c := &NetConf{
		NetConf: types.NetConf{
			Name: "rancher-network",
			Type: "rancher-bridge",
		},
		BrName:   "docker0",
		BrSubnet: "",
		BrIP:     "",
	}

	_, err := calculateBridgeIP(c)
	if err == nil {
		t.Fatalf("Expecting error, didn't get any")
	}
}

func TestErrorNetworkConfigSubnetIPWithNoMask(t *testing.T) {

	c := &NetConf{
		NetConf: types.NetConf{
			Name: "rancher-network",
			Type: "rancher-bridge",
		},
		BrName:   "docker0",
		BrSubnet: "10.42.0.0",
		BrIP:     "",
	}

	_, err := calculateBridgeIP(c)
	if err == nil {
		t.Fatalf("Expecting error, didn't get any")
	}
}

func TestErrorNetworkConfigInvalidSubnetIP(t *testing.T) {

	c := &NetConf{
		NetConf: types.NetConf{
			Name: "rancher-network",
			Type: "rancher-bridge",
		},
		BrName:   "docker0",
		BrSubnet: "a.b.c.d",
		BrIP:     "",
	}

	_, err := calculateBridgeIP(c)
	if err == nil {
		t.Fatalf("Expecting error, didn't get any")
	}
}

func TestErrorNetworkConfigInvalidBridgeIP(t *testing.T) {

	c := &NetConf{
		NetConf: types.NetConf{
			Name: "rancher-network",
			Type: "rancher-bridge",
		},
		BrName:   "docker0",
		BrSubnet: "10.42.0.0/16",
		BrIP:     "a.b.c.d",
	}

	_, err := calculateBridgeIP(c)
	if err == nil {
		t.Fatalf("Expecting error, didn't get any")
	}
}

func TestNetworkConfigWithValidSubnetWithoutBrIP(t *testing.T) {

	conf := &NetConf{
		NetConf: types.NetConf{
			Name: "rancher-network",
			Type: "rancher-bridge",
		},
		BrName:   "docker0",
		BrSubnet: "10.42.0.0/16",
	}

	ip, err := calculateBridgeIP(conf)
	if err != nil {
		t.Fatalf("not expecting error: %v", err)
	}

	logrus.Infof("got ip: %v", ip)
	expected := "10.42.0.1/16"
	actual := ip.String()
	if actual != expected {
		t.Fatalf("expected: %v, got: %v", expected, actual)
	}
}

func TestNetworkConfigWithSubnetIPWithoutBrIP(t *testing.T) {

	conf := &NetConf{
		NetConf: types.NetConf{
			Name: "rancher-network",
			Type: "rancher-bridge",
		},
		BrName:   "docker0",
		BrSubnet: "10.42.0.11/16",
	}

	ip, err := calculateBridgeIP(conf)
	if err != nil {
		t.Fatalf("not expecting error: %v", err)
	}

	logrus.Infof("got ip: %v", ip)
	expected := "10.42.0.1/16"
	actual := ip.String()
	if actual != expected {
		t.Fatalf("expected: %v, got: %v", expected, actual)
	}
}

func TestNetworkConfigWithValidBrIPWithNoMask(t *testing.T) {

	conf := &NetConf{
		NetConf: types.NetConf{
			Name: "rancher-network",
			Type: "rancher-bridge",
		},
		BrName:   "docker0",
		BrSubnet: "10.42.0.0/16",
		BrIP:     "10.42.0.5",
	}

	ip, err := calculateBridgeIP(conf)
	if err != nil {
		t.Fatalf("not expecting error: %v", err)
	}

	logrus.Infof("got ip: %v", ip)
	expected := "10.42.0.5/16"
	actual := ip.String()
	if actual != expected {
		t.Fatalf("expected: %v, got: %v", expected, actual)
	}
}

func TestNetworkConfigWithValidBrIPWithValidMask(t *testing.T) {

	conf := &NetConf{
		NetConf: types.NetConf{
			Name: "rancher-network",
			Type: "rancher-bridge",
		},
		BrName:   "docker0",
		BrSubnet: "10.42.0.0/16",
		BrIP:     "10.42.0.5/16",
	}

	ip, err := calculateBridgeIP(conf)
	if err != nil {
		t.Fatalf("not expecting error: %v", err)
	}

	logrus.Infof("got ip: %v", ip)
	expected := "10.42.0.5/16"
	actual := ip.String()
	if actual != expected {
		t.Fatalf("expected: %v, got: %v", expected, actual)
	}
}

func TestNetworkConfigWithInValidBrIPWithNoMask(t *testing.T) {

	conf := &NetConf{
		NetConf: types.NetConf{
			Name: "rancher-network",
			Type: "rancher-bridge",
		},
		BrName:   "docker0",
		BrSubnet: "10.42.0.0/16",
		BrIP:     "10.41.0.5",
	}

	_, err := calculateBridgeIP(conf)
	if err == nil {
		t.Fatalf("Expecting error, didn't get any")
	}
}
