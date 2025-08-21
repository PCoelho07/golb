# GoLB 🌀

*A simple load balancer written in Go*

![Go](https://img.shields.io/badge/Go-1.22-blue)
![Status](https://img.shields.io/badge/status-learning-green)

---

## 📌 About

**GoLB** is a lightweight HTTP load balancer built in Go, created for study purposes. It aims to be small, readable, and easy to extend.

**Current features:**

* ✅ **Round-robin** load balancing
* ✅ **Reverse proxying** to multiple backends
* ✅ **Basic health checks** (unhealthy backends are temporarily skipped)

**Planned:** random, least‑connections, better logging/metrics, Docker support.

---

## ⚡ How it Works

Incoming requests to the balancer are forwarded to backends using a simple **round‑robin** strategy. If a backend fails a health check, it won’t receive traffic until it becomes healthy again.

```
Client  →  GoLB (8080)  →  [Backend1:3000, Backend2:3001]
```

### Architecture (ASCII)

```
           ┌───────────────────────┐
           │        Client         │
           └───────────┬───────────┘
                       │ HTTP
                       ▼
             ┌───────────────────┐
             │       GoLB        │
             │  (Reverse Proxy)  │
             ├───────────────────┤
             │ Round‑robin       │
             │ Health checks     │
             └─┬───────────────┬─┘
               │               │
         HTTP  ▼               ▼  HTTP
      ┌────────────┐     ┌────────────┐
      │ Backend :3000│   │ Backend :3001│
      └────────────┘     └────────────┘
```

---

## 🚀 Getting Started

### Requirements

* Go 1.22+
* Make

### Run

Clone the repository and start the balancer:

```bash
make run
```

This will start:

* Load balancer on `:8080`
* Two example backends on `:3000` and `:3001` (adjust as needed)

> You can also bring your own backends—any HTTP service will do.

---

## 🔎 Quick Demo (with simple backends)

Start two quick static servers (example using Python):

```bash
# Terminal 1
python3 -m http.server 3000

# Terminal 2
python3 -m http.server 3001
```

Then start GoLB:

```bash
make run
```

Now, hitting `http://localhost:8080` will alternate between `:3000` and `:3001`:

```bash
curl http://localhost:8080
curl http://localhost:8080
curl http://localhost:8080
```

---

## 🩺 Health Checks

GoLB performs basic health checks to avoid sending traffic to failing backends. When a backend is marked unhealthy, it’s skipped until it becomes healthy again. (See code for exact intervals and criteria.)

---

## 🛠️ Roadmap

* [ ] Add **least‑connections** algorithm
* [ ] Add **random** algorithm
* [ ] Structured logging & metrics
* [ ] Docker support

---

## 🤝 Contributing

This is a learning project—fork it, tweak it, and try new balancing strategies! PRs and ideas are welcome.

---

## 📜 License

MIT License – free to use and modify.

