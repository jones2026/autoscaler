// Copyright 2018 Drone.IO Inc
// Use of this software is governed by the Business Source License
// that can be found in the LICENSE file.

package mocks

//go:generate mockgen -package=mocks -destination=mock_engine.go   github.com/drone/autoscaler Engine
//go:generate mockgen -package=mocks -destination=mock_server.go   github.com/drone/autoscaler ServerStore
//go:generate mockgen -package=mocks -destination=mock_provider.go github.com/drone/autoscaler Provider
//go:generate mockgen -package=mocks -destination=mock_drone.go    github.com/drone/drone-go/drone Client
//go:generate mockgen -package=mocks -destination=mock_docker.go   docker.io/go-docker APIClient
