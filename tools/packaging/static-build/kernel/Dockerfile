# Copyright (c) 2021 Intel Corporation
#
# SPDX-License-Identifier: Apache-2.0

FROM ubuntu:20.04
ENV DEBIAN_FRONTEND=noninteractive
ENV http_proxy http://proxy.cd.intel.com:911
ENV https_proxy http://proxy.cd.intel.com:911
# kernel deps
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
	    bc \
	    bison \
	    build-essential \
	    ca-certificates \
	    curl \
	    flex \
	    git \
	    iptables \
	    libelf-dev && \
    if [ "$(uname -m)" = "s390x" ]; then apt-get install -y --no-install-recommends libssl-dev; fi && \
    apt-get clean && rm -rf /var/lib/lists/
