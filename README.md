# IcedID-C2-Emulator

Simple Golang webserver to serve a IcedID second stage payload to facilaite dynamic analysis.

Once running the server will listen on port 80 and serve the `FileToServe.dat` file, which should be replaced with a encoded second stage payload.
If you do not have a second stage payload one can be extracted from a PCAP found online.

![image](https://github.com/iamaleks/IcedID-C2-Emulator/assets/52838964/9345edb5-cd9a-485f-9a40-26ee98d94f08)
