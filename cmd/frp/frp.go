// Copyright 2018 fatedier, fatedier@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package frp

import (
	"github.com/HaidyCao/frp_0420/cmd/frpc/sub"
	"github.com/HaidyCao/frp_0420/cmd/frps/frps"
	"github.com/HaidyCao/frp_0420/pkg/util/version"
	"github.com/HaidyCao/golib/crypto"
	"runtime/debug"
)

func init() {
	debug.SetGCPercent(5)
}

func RunFrpc(cfgFilePath string) (err error) {
	crypto.DefaultSalt = "frp"

	return sub.RunFrpc(cfgFilePath)
}

func StopFrpc() (err error) {
	defer debug.FreeOSMemory()
	return sub.StopFrp()
}

func IsFrpcRunning() bool {
	return sub.IsFrpRunning()
}

func RunFrps(cfgFilePath string) (err error) {
	crypto.DefaultSalt = "frp"

	return frps.RunFrps(cfgFilePath)
}

// StopFrps 停止frps服务
func StopFrps() error {
	defer debug.FreeOSMemory()
	return frps.StopFrps()
}

// IsFrpsRunning 是否还在运行
func IsFrpsRunning() bool {
	return frps.IsFrpsRunning()
}

// Version frp 版本号
func Version() string {
	return version.Full()
}
