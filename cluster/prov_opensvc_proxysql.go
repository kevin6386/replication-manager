// replication-manager - Replication Manager Monitoring and CLI for MariaDB and MySQL
// Copyright 2017 Signal 18 SARL
// Authors: Guillaume Lefranc <guillaume@signal18.io>
//          Stephane Varoqui  <svaroqui@gmail.com>
// This source code is licensed under the GNU General Public License, version 3.

package cluster

import (
	"fmt"
	"log"

	"github.com/signal18/replication-manager/opensvc"
)

func (cluster *Cluster) GetProxysqlTemplate(collector opensvc.Collector, servers string, agent opensvc.Host, prx *Proxy) (string, error) {

	conf := `
[DEFAULT]
nodes = {env.nodes}
flex_primary = {env.nodes[0]}
topology = flex
rollback = false
orchestrate = start
`
	conf += "app = " + cluster.Conf.ProvCodeApp
	conf = conf + cluster.GetDockerDiskTemplate(collector)
	i := 0
	pod := fmt.Sprintf("%02d", i+1)
	conf = conf + cluster.GetPodDiskTemplate(collector, pod, agent.Node_name)
	conf = conf + prx.GetInitContainer(collector)
	//conf = conf + `post_provision = {svcmgr} -s {svcpath} push status;{svcmgr} -s {svcpath} compliance fix --attach --moduleset mariadb.svc.mrm.proxy
	//`
	conf = conf + cluster.GetPodNetTemplate(collector, pod, i)
	conf = conf + cluster.GetPodDockerProxysqlTemplate(collector, pod)
	conf = conf + cluster.GetPodPackageTemplate(collector, pod)
	conf = conf + cluster.GetProxiesEnv(collector, servers, agent, prx)
	log.Println(conf)
	return conf, nil
}

func (cluster *Cluster) GetPodDockerProxysqlTemplate(collector opensvc.Collector, pod string) string {
	var vm string
	if collector.ProvProxMicroSrv == "docker" {
		vm = vm + `
[container#00` + pod + `]
type = docker
hostname = {svcname}.{namespace}.svc.{clustername}
image = google/pause
rm = true


[container#20` + pod + `]
tags = pod` + pod + `
type = docker
run_image = {env.proxysql_img}
rm = true
netns = container#00` + pod + `
run_args = --ulimit nofile=262144:262144
    -v /etc/localtime:/etc/localtime:ro
    -v {env.base_dir}/pod` + pod + `/conf/proxysql.cnf:/etc/proxysql.cnf:rw
		-v {env.base_dir}/pod` + pod + `/data:/var/lib/proxysql:rw
run_command = proxysql --initial -f -c /conf/proxysql.cnf
`
		if dockerMinusRm {
			vm = vm + ` --rm
`
		}
	}
	return vm
}
