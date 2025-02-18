
IMAGE_NAME = quay.io/skupper/tcp-go-echo
PLATFORM = linux/amd64,linux/arm64,linux/s390x

# The option below creates the images in the docker format.
# That is required for their use with Openshift 3.11
FORMAT_OPTIONS = --format docker

.DEFAULT_GOAL := build

build: clean
	podman build --no-cache --platform $(PLATFORM) --manifest $(IMAGE_NAME) $(FORMAT_OPTIONS) .

push:
	podman manifest push $(IMAGE_NAME)

clean:
	-podman manifest rm $(IMAGE_NAME)

check:
	podman manifest inspect $(IMAGE_NAME)
