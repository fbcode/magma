/*
 * Copyright (c) Facebook, Inc. and its affiliates.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

//go:generate bash -c "protoc -I . -I /usr/include --proto_path=$MAGMA_ROOT --go_out=plugins=grpc:. *.proto"
package storage
