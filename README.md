# tc

Simple project to validate that following tools are working:

- Make
- Go
- Docker/Podman
- testcontainers-go

To validate your environment run:

```bash
git clone https://github.com/heppu/tc
cd tc
make
```

If everything works correctly you should see something like this:

```
go test -count=1 -v .
2025/05/07 11:41:53 github.com/testcontainers/testcontainers-go - Connected to docker:
  Server Version: 28.0.4
  API Version: 1.48
  Operating System: Void Linux
  Total Memory: 31433 MB
  Testcontainers for Go Version: v0.37.0
  Resolved Docker Host: unix:///var/run/docker.sock
  Resolved Docker Socket Path: /var/run/docker.sock
  Test SessionID: 2b49b0502a02c6a823b215e92193c04af3d461cb8673e557041997c699f39c48
  Test ProcessID: d9015f54-2877-45f7-becd-f3553ddeb608
 Network c9d1c143-e629-4e4c-a6e5-454192d2bc2b_default  Creating
 Network c9d1c143-e629-4e4c-a6e5-454192d2bc2b_default  Created
 Container c9d1c143-e629-4e4c-a6e5-454192d2bc2b-postgres-1  Creating
 Container c9d1c143-e629-4e4c-a6e5-454192d2bc2b-postgres-1  Created
 Container c9d1c143-e629-4e4c-a6e5-454192d2bc2b-postgres-1  Starting
 Container c9d1c143-e629-4e4c-a6e5-454192d2bc2b-postgres-1  Started
 Container c9d1c143-e629-4e4c-a6e5-454192d2bc2b-postgres-1  Waiting
 Container c9d1c143-e629-4e4c-a6e5-454192d2bc2b-postgres-1  Healthy
2025/05/07 11:41:55 🐳 Creating container for image testcontainers/ryuk:0.11.0
2025/05/07 11:41:55 ✅ Container created: 27c2810bd706
2025/05/07 11:41:55 🐳 Starting container: 27c2810bd706
2025/05/07 11:41:56 ✅ Container started: 27c2810bd706
2025/05/07 11:41:56 ⏳ Waiting for container id 27c2810bd706 image: testcontainers/ryuk:0.11.0. Waiting for: &{Port:8080/tcp timeout:<nil> PollInterval:100ms skipInternalCheck:false}
2025/05/07 11:41:56 🔔 Container is ready: 27c2810bd706
=== RUN   TestValidate
    tc_test.go:25: Test environment is working correctly
--- PASS: TestValidate (0.00s)
PASS
 Container c9d1c143-e629-4e4c-a6e5-454192d2bc2b-postgres-1  Stopping
 Container c9d1c143-e629-4e4c-a6e5-454192d2bc2b-postgres-1  Stopped
 Container c9d1c143-e629-4e4c-a6e5-454192d2bc2b-postgres-1  Removing
 Container c9d1c143-e629-4e4c-a6e5-454192d2bc2b-postgres-1  Removed
 Network c9d1c143-e629-4e4c-a6e5-454192d2bc2b_default  Removing
 Network c9d1c143-e629-4e4c-a6e5-454192d2bc2b_default  Removed
ok  	github.com/heppu/tc	3.431s
```

## Helpful links

- [https://docs.rancherdesktop.io/how-to-guides/using-testcontainers/](https://docs.rancherdesktop.io/how-to-guides/using-testcontainers/)

- [https://podman-desktop.io/tutorial/testcontainers-with-podman](https://podman-desktop.io/tutorial/testcontainers-with-podman)
