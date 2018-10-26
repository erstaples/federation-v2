FROM buildpack-deps:stretch-curl

LABEL maintainer="erstaples@gmail.com"
LABEL description="Execute kubefed2 commands. Contains all kubefed2 dependency binaries."

ARG KUBEFED2_VERSION=v0.0.2-rc.1
ARG KUBEBUILDER_VERSION=1.0.4

RUN curl -LO https://github.com/kubernetes-sigs/kubebuilder/releases/download/v${KUBEBUILDER_VERSION}/kubebuilder_${KUBEBUILDER_VERSION}_linux_amd64.tar.gz \
  && tar -xzvf kubebuilder_${KUBEBUILDER_VERSION}_linux_amd64.tar.gz \
  && chmod +x kubebuilder_${KUBEBUILDER_VERSION}_linux_amd64/bin/* \
  && mv kubebuilder_${KUBEBUILDER_VERSION}_linux_amd64/bin/* /usr/local/bin

RUN curl -LO https://github.com/kubernetes-sigs/federation-v2/releases/download/${KUBEFED2_VERSION}/kubefed2.tar.gz \
  && tar -xzvf kubefed2.tar.gz \
  && mv kubefed2 /usr/local/bin

ENTRYPOINT ["kubefed2"]
CMD ["help"]


