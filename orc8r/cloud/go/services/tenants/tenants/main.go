/*
 * Copyright (c) Facebook, Inc. and its affiliates.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

package main

import (
	"magma/orc8r/cloud/go/blobstore"
	"magma/orc8r/cloud/go/obsidian"
	"magma/orc8r/cloud/go/orc8r"
	"magma/orc8r/cloud/go/service"
	"magma/orc8r/cloud/go/services/tenants"
	"magma/orc8r/cloud/go/services/tenants/obsidian/handlers"
	"magma/orc8r/cloud/go/services/tenants/servicers"
	"magma/orc8r/cloud/go/services/tenants/servicers/storage"
	"magma/orc8r/cloud/go/sqorc"
	storage2 "magma/orc8r/cloud/go/storage"
	"magma/orc8r/lib/go/protos"

	"github.com/golang/glog"
)

func main() {
	srv, err := service.NewOrchestratorService(orc8r.ModuleName, tenants.ServiceName)
	if err != nil {
		glog.Fatalf("Error creating tenants service %s", err)
	}
	db, err := sqorc.Open(storage2.SQLDriver, storage2.DatabaseSource)
	if err != nil {
		glog.Fatalf("Failed to connect to database: %s", err)
	}
	factory := blobstore.NewEntStorage(tenants.DBTableName, db, sqorc.GetSqlBuilder())
	err = factory.InitializeFactory()
	if err != nil {
		glog.Fatalf("Error initializing tenant database: %s", err)
	}
	store := storage.NewBlobstoreStore(factory)

	server, err := servicers.NewTenantsServicer(store)
	if err != nil {
		glog.Fatalf("Error creating tenants server: %s", err)
	}
	protos.RegisterTenantsServiceServer(srv.GrpcServer, server)
	obsidian.AttachHandlers(srv.EchoServer, handlers.GetObsidianHandlers())

	err = srv.Run()
	if err != nil {
		glog.Fatalf("Error running service: %s", err)
	}
}
