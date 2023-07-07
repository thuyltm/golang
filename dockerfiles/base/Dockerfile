FROM amd64/ubuntu:focal
# Set for all apt-get install, must be at the very beginning of the Dockerfile.
ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update && \
    apt-get install -y software-properties-common
RUN apt-get install build-essential libopus-dev libopusfile-dev -y
RUN apt-get install librdkafka-dev -y
RUN add-apt-repository ppa:savoury1/ffmpeg4 -y
RUN apt-get install libswscale-dev libavcodec-dev libavformat-dev libswresample-dev libavutil-dev -y

# Non-interactive modes get set back.
ENV DEBIAN_FRONTEND newt
CMD echo "This is a test." | wc -